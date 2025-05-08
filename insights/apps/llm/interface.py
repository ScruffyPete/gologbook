from typing import Protocol
from apps.domain.entities import Entry, Insight


class LLMInterface(Protocol):
    async def generate_insight(self, entry: Entry) -> Insight: ...
