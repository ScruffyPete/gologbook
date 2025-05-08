import os
import pytest
from apps.queue.interface import QueueMessage
from apps.queue.redis import Redis

import redis


def check_redis():
    port = os.getenv("REDIS_PORT")
    host = os.getenv("REDIS_HOST")
    db = os.getenv("REDIS_DEAULT_DB")

    redis_client = redis.Redis(host=host, port=port, db=db, decode_responses=True)

    try:
        result = redis_client.ping()
        redis_client.flushall()
        print("Redis responded:", result)
        return True
    except Exception as e:
        print("Redis check failed:", e)
        raise e
        return False


@pytest.mark.asyncio
@pytest.mark.skipif(not check_redis(), reason="Redis is not running")
async def test_redis_queue():
    stream = "test_stream"
    queue = Redis(stream)
    message = QueueMessage(type="test", payload={"test": "test"})
    await queue.redis.xadd(stream, fields={"message": message.to_json()})
    assert await queue.pop() == message
