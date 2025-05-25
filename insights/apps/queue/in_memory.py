import asyncio
from collections import defaultdict
from contextlib import asynccontextmanager
import os
from typing import Any
from uuid import UUID


class InMemoryQueue:
    items: defaultdict[defaultdict[float]]
    
    def __init__(self):
        self.items = defaultdict(lambda: defaultdict(float))
        self.key = os.getenv("REDIS_PENDING_PROJECTS_KEY")

    @classmethod
    @asynccontextmanager
    async def create(cls):
        queue = cls()
        yield queue

    async def pop_ready_projects(self, cutoff_time: float, batch_size: int) -> tuple[UUID]:
        try:
            return self.items[self.key]
        except asyncio.CancelledError:
            return None
        except Exception as e:
            print(f"Error popping from app.queue: {e}")
            return None
