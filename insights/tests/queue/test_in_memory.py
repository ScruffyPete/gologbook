import pytest
import pytest_asyncio
from apps.queue.in_memory import InMemory


@pytest_asyncio.fixture
async def queue():
    queue = InMemory()
    await queue.queue.put("Hello, world!")
    yield queue


@pytest.mark.asyncio
async def test_in_memory_queue(queue):
    assert await queue.pop() == "Hello, world!"
