import datetime
import time
import uuid
import pytest
import pytest_asyncio
from apps.queue.in_memory import InMemoryQueue

@pytest.fixture
def key() -> str:
    return "test_key"

@pytest.fixture
def project_id() -> uuid.UUID:
    return uuid.uuid4()

@pytest_asyncio.fixture
async def queue(key: str, project_id: uuid.UUID):
    async with InMemoryQueue.create() as queue:
        queue.items[key][project_id] = datetime.datetime.now()
        yield queue


@pytest.mark.asyncio
async def test_in_memory_queue(
    queue: InMemoryQueue, 
    key: str, 
    project_id: uuid.UUID
):
    timestamp = queue.items[key][project_id]
    assert timestamp is not None
