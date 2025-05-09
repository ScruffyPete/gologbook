import pytest
from apps.queue.interface import QueueMessage
from apps.queue.redis import Redis


@pytest.mark.asyncio
@pytest.mark.queue
async def test_redis_queue():
    stream = "test_stream"
    queue = Redis(stream)
    message = QueueMessage(type="test", payload={"test": "test"})
    await queue.redis.xadd(stream, fields={"message": message.to_json()})
    assert await queue.pop() == message
