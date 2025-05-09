import pytest
import uuid
from datetime import datetime

from apps.db.in_memory import InMemoryRepositoryUnit
from apps.domain.entities import Entry
from apps.services.processor import process_entry
from apps.llm.in_memory import InMemoryLLM


@pytest.fixture
def entry():
    return Entry(
        id=uuid.uuid4(),
        project_id=uuid.uuid4(),
        body="Hello, world!",
        created_at=datetime.now(),
    )


@pytest.fixture
def llm():
    return InMemoryLLM()


@pytest.fixture
def repo(entry: Entry):
    return InMemoryRepositoryUnit(entries=[entry])


@pytest.mark.asyncio
async def test_processor(repo: InMemoryRepositoryUnit, llm: InMemoryLLM, entry: Entry):
    await process_entry(repo, entry.id, llm)

    insights = await repo.insight_repo.get_insights_by_entry_id(entry.id)
    assert len(insights) == 1
    insight = insights[0]
    assert insight.entry_ids == [entry.id]
