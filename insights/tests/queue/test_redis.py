import os
import pytest
import pytest_asyncio
from apps.queue.redis import RedisQueue

@pytest.fixture
def key() -> str:
    return os.getenv("REDIS_PENDING_PROJECTS_KEY")


@pytest.fixture
def lock_prefix() -> str:
    return os.getenv("REDIS_PROJECT_LOCK_PREFIX")


@pytest_asyncio.fixture
async def queue(key: str, lock_prefix: str):
    mapping = {
        "one": 1,
        "two": 2, 
        "three": 3
    }
    async with RedisQueue.create() as q:
        await q.redis_client.zadd(name=key, mapping=mapping)
        yield q
        for project_id in mapping.keys():
            await q.redis_client.delete(f"{lock_prefix}:{project_id}")


@pytest.mark.asyncio
@pytest.mark.queue
async def test_pop_ready_projects(queue: RedisQueue):
    project_ids = await queue.pop_ready_projects(
        cuttoff_time=5,
        batch_size=10,
    )
    
    expected = ["one", "two", "three"]
    assert list(project_ids) == expected
    

@pytest.mark.asyncio
@pytest.mark.queue
async def test_remove_processed_projects(queue: RedisQueue):
    project_ids = await queue.pop_ready_projects(
        cuttoff_time=5,
        batch_size=10,
    )
    assert len(project_ids) > 0
    
    await queue.remove_processed_projects(project_ids)
    
    new_project_ids = await queue.pop_ready_projects(
        cuttoff_time=5,
        batch_size=10,
    )
    assert new_project_ids == tuple()
    