import asyncio
from typing import Any


class InMemory:
    def __init__(self):
        self.queue = asyncio.Queue()

    async def pop(self) -> Any:
        try:
            return await self.queue.get()
        except asyncio.CancelledError:
            return None
        except Exception as e:
            print(f"Error popping from app.queue: {e}")
            return None
