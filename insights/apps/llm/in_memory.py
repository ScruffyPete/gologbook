import uuid
from datetime import datetime

from apps.domain.entities import Entry, Insight


class InMemoryLLM:
    async def generate_insight(self, entry: Entry) -> Insight:
        insight_text = f"Insight for entry {entry.id}: {entry.text[:100]}"
        return Insight(
            id=uuid.uuid4(),
            project_id=entry.project_id,
            entry_id=entry.id,
            text=insight_text,
            created_at=datetime.now(),
        )
