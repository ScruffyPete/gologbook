import json
from dataclasses import dataclass
from typing import Any, Protocol, Self
from uuid import UUID


class QueueInterface(Protocol):
    async def pop_ready_projects(self, cuttoff_time: float, batch_size: int) -> tuple[UUID]: ...
