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
    buffer = []
    async for token in llm.stream_document(project_id, [entry]):
        buffer.append(token)
        
    body = "".join(buffer)
    assert body == f"Insight for entry {entry.id}: {entry.body[:100]}\n\n"
