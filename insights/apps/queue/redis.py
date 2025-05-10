from contextlib import asynccontextmanager
import os

import redis.asyncio as aioredis
from typing import Any

from apps.queue.interface import QueueMessage


class RedisQueue:
    def __init__(self):
        redis_host = os.getenv("REDIS_HOST")
        redis_port = os.getenv("REDIS_PORT")
        redis_db = os.getenv("REDIS_DB")
        self.redis_client = aioredis.Redis.from_url(
            f"redis://{redis_host}:{redis_port}/{redis_db}"
        )
        self.stream = os.getenv("REDIS_STREAM")
        self.last_id = "0"

    @classmethod
    @asynccontextmanager
    async def create(cls):
        rq = cls()
        yield rq
        await rq.redis_client.aclose()

    async def pop(self) -> Any:
        try:
            entries = await self.redis_client.xread(
                streams={self.stream: self.last_id}, count=1, block=1000
            )
            if not entries:
                return None

            _, messages = entries[0]
            msg_id, fields = messages[0]
            self.last_id = msg_id

            raw_bytes = fields[b"message"]
            json_str = raw_bytes.decode("utf-8")
            return QueueMessage.from_json(json_str)

        except aioredis.TimeoutError:
            return None
