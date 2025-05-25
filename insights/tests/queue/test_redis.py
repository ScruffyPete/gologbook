import os
import pytest
from apps.queue.redis import RedisQueue

@pytest.fixture
def key() -> str:
    return os.getenv("REDIS_PENDING_PROJECTS_KEY")


@pytest.mark.asyncio
@pytest.mark.queue
async def test_redis_queue(key):
    mapping = {
        "one": 1,
        "two": 2, 
        "three": 3
    }

    async with RedisQueue.create() as queue:
        await queue.redis_client.zadd(name=key, mapping=mapping)
        
        project_ids = await queue.pop_ready_projects(
            cuttoff_time=5,
            batch_size=10,
        )
        
        expected = ["one", "two", "three"]
        assert list(project_ids) == expected
