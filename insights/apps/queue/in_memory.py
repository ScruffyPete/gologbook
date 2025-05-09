import asyncio
from contextlib import asynccontextmanager
from typing import Any


class InMemoryQueue:
    def __init__(self):
        self.queue = asyncio.Queue()

    @classmethod
    @asynccontextmanager
    async def create(cls):
        queue = cls()
        yield queue

    async def pop(self) -> Any:
        try:
            return await self.queue.get()
        except asyncio.CancelledError:
            return None
        except Exception as e:
            print(f"Error popping from app.queue: {e}")
            return None
