import pytest
from apps.queue.interface import QueueMessage
from apps.queue.redis import RedisQueue


@pytest.mark.asyncio
@pytest.mark.queue
async def test_redis_queue():
    stream_name = "test_stream"
    async with RedisQueue.create(stream=stream_name) as queue:
        message = QueueMessage(type="test", payload={"test": "test"})
        await queue.redis_client.xadd(
            stream_name, fields={"message": message.to_json()}
        )
        assert await queue.pop() == message
