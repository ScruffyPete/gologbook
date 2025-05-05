from typing import Protocol
from insights.domain.entities import Entry, Insight


class LLMInterface(Protocol):
    async def generate_insight(self, entry: Entry) -> Insight: ...
