from contextlib import asynccontextmanager
import os
import uuid

import redis.asyncio as aioredis
from typing import Any

from tests.queue.test_in_memory import key


class RedisQueue:
    def __init__(self):
        redis_host = os.getenv("REDIS_HOST")
        redis_port = os.getenv("REDIS_PORT")
        redis_db = os.getenv("REDIS_DB")
        self.redis_client = aioredis.Redis.from_url(
            f"redis://{redis_host}:{redis_port}/{redis_db}"
        )
        self.pending_projects_key = os.getenv("REDIS_PENDING_PROJECTS_KEY")
        self.lock_prefix = os.getenv("REDIS_PROJECT_LOCK_PREFIX")
        self.lock_ttl = 30

    @classmethod
    @asynccontextmanager
    async def create(cls):
        rq = cls()
        yield rq
        await rq.redis_client.aclose()

    async def pop_ready_projects(self, cuttoff_time: float, batch_size: int) -> tuple[uuid.UUID]:
        raw_project_ids = await self.redis_client.zrangebyscore(
            name=self.pending_projects_key,
            min="-inf",
            max=cuttoff_time,
            start=0,
            num=batch_size * 5  # overfetch to account for locked entries
        )
        if not raw_project_ids:
            return tuple()

        ready_project_ids = []
        for raw_project_id in raw_project_ids:
            project_id = raw_project_id.decode()
            lock_key = f"{self.lock_prefix}:{project_id}"
            locked = await self.redis_client.set(
                name=lock_key,
                value="1",
                nx=True, # Only set the key if it does not already exist
                ex=self.lock_ttl,
            )
            if locked:
                ready_project_ids.append(project_id)
            
            if len(ready_project_ids) >= batch_size:
                # cuttoff the overfetch
                break
            
        return tuple(ready_project_ids)
