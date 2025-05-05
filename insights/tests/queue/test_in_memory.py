import pytest
from insights.queue.in_memory import InMemoryQueue


@pytest.mark.asyncio
async def test_in_memory_queue():
    queue = InMemoryQueue()
    assert queue.is_empty()

    await queue.push("Hello, world!")
    assert not queue.is_empty()
    assert await queue.pop() == "Hello, world!"
    assert queue.is_empty()
