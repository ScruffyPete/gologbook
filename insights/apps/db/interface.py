from typing import Protocol
from apps.domain.entities import Entry, Insight


class DBInterface(Protocol):
    async def get_entry(self, entry_id: str) -> Entry | None: ...

    async def get_insights_by_entry_id(self, entry_id: str) -> list[Insight]: ...

    async def save_insight(self, insight: Insight) -> None: ...
