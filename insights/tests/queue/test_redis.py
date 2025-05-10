import pytest
from apps.queue.interface import QueueMessage
from apps.queue.redis import RedisQueue


@pytest.mark.asyncio
@pytest.mark.queue
async def test_redis_queue():
    async with RedisQueue.create() as queue:
        message = QueueMessage(type="test", payload={"test": "test"})
        await queue.redis_client.xadd(
            queue.stream, fields={"message": message.to_json()}
        )
        assert await queue.pop() == message
