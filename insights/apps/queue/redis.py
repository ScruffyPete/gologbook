import os

import redis.asyncio as aioredis
from typing import Any

from apps.queue.interface import QueueMessage


class Redis:
    def __init__(self, stream: str):
        redis_host = os.getenv("REDIS_HOST")
        redis_port = os.getenv("REDIS_PORT")
        redis_db = os.getenv("REDIS_DB")
        self.redis = aioredis.Redis.from_url(
            f"redis://{redis_host}:{redis_port}/{redis_db}"
        )
        self.stream = stream
        self.last_id = "0"

    async def pop(self) -> Any:
        try:
            entries = await self.redis.xread(
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
