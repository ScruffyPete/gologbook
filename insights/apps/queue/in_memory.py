import asyncio
from collections import defaultdict
from contextlib import asynccontextmanager
import os
from typing import Any
from uuid import UUID


class InMemoryQueue:
    pending_projects_zset: dict[str, dict[UUID, float]]
    project_token_channels: defaultdict[asyncio.Queue]
    
    def __init__(self):
        self.pending_projects_zset: dict[str, dict[UUID, float]] = {}
        self.project_token_channels = defaultdict(lambda: asyncio.Queue(maxsize=100))
        self.key = os.getenv("REDIS_PENDING_PROJECTS_KEY")

    @classmethod
    @asynccontextmanager
    async def create(cls):
        queue = cls()
        yield queue

    async def pop_ready_projects(self, cutoff_time: float, batch_size: int) -> dict[UUID, float]:
        try:
            return self.pending_projects_zset[self.key]
        except KeyError:
            return {}

    async def remove_processed_projects(self, project_ids: list[UUID]) -> None:
        for project_id in project_ids:
            del self.pending_projects_zset[self.key][project_id]

    async def publish_project_token(self, project_id: UUID, token: str):
        await self.project_token_channels[project_id].put(token)

    async def clear_project_stream(self, project_id: UUID) -> None:
        if self.key in self.pending_projects_zset and project_id in self.pending_projects_zset[self.key]:
            del self.pending_projects_zset[self.key][project_id]