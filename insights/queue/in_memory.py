import asyncio
from typing import Any


class InMemoryQueue:
    def __init__(self):
        self.queue = asyncio.Queue()

    def is_empty(self) -> bool:
        return self.queue.empty()

    async def push(self, item: Any) -> None:
        await self.queue.put(item)

    async def pop(self) -> Any:
        try:
            return await self.queue.get()
        except asyncio.CancelledError:
            return None
        except Exception as e:
            print(f"Error popping from queue: {e}")
            return None
