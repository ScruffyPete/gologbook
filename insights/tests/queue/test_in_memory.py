import pytest
import pytest_asyncio
from apps.queue.in_memory import InMemoryQueue


@pytest_asyncio.fixture
async def queue():
    async with InMemoryQueue.create() as queue:
        await queue.queue.put("Hello, world!")
        yield queue


@pytest.mark.asyncio
async def test_in_memory_queue(queue):
    assert await queue.pop() == "Hello, world!"
