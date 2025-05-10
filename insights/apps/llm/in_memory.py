import uuid
from contextlib import asynccontextmanager
from datetime import datetime

from apps.domain.entities import Entry, Insight


class InMemoryLLM:
    @classmethod
    @asynccontextmanager
    async def create(cls):
        llm = cls()
        yield llm

    async def generate_insight(self, entry: Entry) -> Insight:
        insight_body = f"Insight for entry {entry.id}: {entry.body[:100]}"
        return Insight(
            id=uuid.uuid4(),
            project_id=entry.project_id,
            entry_ids=[entry.id],
            body=insight_body,
            created_at=datetime.now(),
        )
