import json
from dataclasses import dataclass
from typing import Any, Protocol, Self
from uuid import UUID


class QueueInterface(Protocol):
    async def pop_ready_projects(self, cuttoff_time: float, batch_size: int) -> tuple[UUID]: ...
    
    async def remove_processed_projects(self, project_ids: list[UUID]) -> None: ...
    
    async def publish_project_token(self, project_id: UUID, token: str) -> None: ...
