from dataclasses import dataclass
from datetime import datetime
from uuid import UUID


@dataclass
class Project:
    id: UUID
    title: str
    created_at: datetime


@dataclass
class Entry:
    id: UUID
    project_id: UUID
    body: str
    created_at: datetime


@dataclass
class Insight:
    id: UUID
    project_id: UUID
    entry_ids: list[UUID]
    body: str
    created_at: datetime
