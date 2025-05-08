from dataclasses import dataclass
from datetime import datetime


@dataclass
class Entry:
    id: str
    project_id: str
    text: str
    created_at: datetime


@dataclass
class Insight:
    id: str
    project_id: str
    entry_id: str | None
    text: str
    created_at: datetime
