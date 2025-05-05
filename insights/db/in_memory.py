from insights.domain.entities import Entry, Insight


class InMemoryDB:
    def __init__(self, entries: list[Entry] = [], insights: list[Insight] = []):
        self.entries = {entry.id: entry for entry in entries}
        self.insights = {insight.id: insight for insight in insights}

    async def get_entry(self, entry_id: str) -> Entry | None:
        return self.entries.get(entry_id, None)

    async def save_entry(self, entry: Entry) -> None:
        self.entries[entry.id] = entry

    async def get_insights_by_entry_id(self, entry_id: str) -> list[Insight]:
        return [
            insight
            for insight in self.insights.values()
            if insight.entry_id == entry_id
        ]

    async def save_insight(self, insight: Insight) -> None:
        self.insights[insight.id] = insight
