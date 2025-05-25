import pytest
import uuid
from datetime import datetime

from apps.domain.entities import Entry
from apps.llm.in_memory import InMemoryLLM


@pytest.mark.asyncio
async def test_in_memory_llm():
    llm = InMemoryLLM()
    project_id = uuid.uuid4()
    entry = Entry(
        id=uuid.uuid4(),
        project_id=project_id,
        body="Hello, world!",
        created_at=datetime.now(),
    )
    insight = await llm.compile_messages(project_id, [entry])
    assert insight.entry_ids == [entry.id]
    assert insight.body == f"Insight for entry {entry.id}: {entry.body[:100]}"
