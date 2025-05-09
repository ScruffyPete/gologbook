from typing import Protocol
from apps.domain.entities import Entry, Insight
from uuid import UUID


class EntryRepository(Protocol):
    async def get_entry(self, entry_id: UUID) -> Entry | None: ...


class InsightRepository(Protocol):
    async def get_insights_by_entry_id(self, entry_id: UUID) -> list[Insight]: ...

    async def save_insight(self, insight: Insight) -> None: ...


class RepositoryUnit(Protocol):
    entry_repo: EntryRepository
    insight_repo: InsightRepository
