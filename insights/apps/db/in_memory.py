from contextlib import asynccontextmanager
from apps.domain.entities import Entry, Insight
from uuid import UUID


class InMemoryEntryRepository:
    def __init__(self, entries: list[Entry] = []):
        self.entries = {entry.id: entry for entry in entries}

    async def get_entry(self, entry_id: UUID) -> Entry | None:
        return self.entries.get(entry_id, None)


class InMemoryInsightRepository:
    def __init__(self, insights: list[Insight] = []):
        self.insights = {insight.id: insight for insight in insights}

    async def get_insights_by_entry_id(self, entry_id: UUID) -> list[Insight]:
        return [
            insight
            for insight in self.insights.values()
            if entry_id in insight.entry_ids
        ]

    async def create(self, insight: Insight) -> None:
        self.insights[insight.id] = insight


class InMemoryRepositoryUnit:
    def __init__(self, entries: list[Entry] = [], insights: list[Insight] = []):
        self.entry_repo = InMemoryEntryRepository(entries)
        self.insight_repo = InMemoryInsightRepository(insights)

    @classmethod
    @asynccontextmanager
    async def create(cls):
        repo = cls()
        yield repo
