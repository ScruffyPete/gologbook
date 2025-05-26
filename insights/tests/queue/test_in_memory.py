import datetime
import os
import uuid
import pytest
import pytest_asyncio
from apps.queue.in_memory import InMemoryQueue

@pytest.fixture
def project_id() -> uuid.UUID:
    return uuid.uuid4()

@pytest_asyncio.fixture
async def queue(project_id: uuid.UUID):
    queue = InMemoryQueue()
    key = os.getenv("REDIS_PENDING_PROJECTS_KEY")
    now = datetime.datetime.now().timestamp()
    queue.pending_projects_zset[key] = {project_id: now}
    return queue


@pytest.mark.asyncio
async def test_pop_ready_projects(
    queue: InMemoryQueue, 
    project_id: uuid.UUID
):
    results = await queue.pop_ready_projects(0,0)  # params don't affect the in memory result
    assert project_id in results
    timestamp = results[project_id]
    assert timestamp > 0

@pytest.mark.asyncio
async def test_remove_processed_projects(
    queue: InMemoryQueue,
    project_id: uuid.UUID,
):
    results = await queue.pop_ready_projects(0,0)
    assert len(results) > 0
    
    await queue.remove_processed_projects([project_id])
    
    results = await queue.pop_ready_projects(0,0)  # params don't affect the in memory result
    assert results == {}
    
    
@pytest.mark.asyncio
async def test_publish_project_token(
    queue: InMemoryQueue,
    project_id: uuid.UUID,
):
    token = "test_token"
    await queue.publish_project_token(project_id, token)
    
    q_token = await queue.project_token_channels[project_id].get()
    assert token == q_token
    