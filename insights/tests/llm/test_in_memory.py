import pytest
from insights.llm.in_memory import InMemoryLLM
from insights.domain.entities import Entry
import uuid
from datetime import datetime

@pytest.mark.asyncio
async def test_in_memory_llm():
    llm = InMemoryLLM()
    entry = Entry(id=uuid.uuid4(), project_id=uuid.uuid4(), text="Hello, world!", created_at=datetime.now())
    insight = await llm.generate_insight(entry)
    assert insight.entry_id == entry.id
    assert insight.text == f"Insight for entry {entry.id}: {entry.text[:100]}"
