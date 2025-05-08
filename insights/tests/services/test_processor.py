import pytest
import uuid
from datetime import datetime

import pytest_asyncio

from apps.domain.entities import Entry
from apps.services.processor import process_entry
from apps.db.in_memory import InMemoryDB
from apps.llm.in_memory import InMemoryLLM


@pytest.fixture
def entry():
    return Entry(
        id=uuid.uuid4(),
        project_id=uuid.uuid4(),
        body="Hello, world!",
        created_at=datetime.now(),
    )


@pytest_asyncio.fixture
async def db(entry):
    db = InMemoryDB(entries=[entry])
    return db


@pytest.fixture
def llm():
    return InMemoryLLM()


@pytest.mark.asyncio
async def test_processor(entry, db, llm):
    await process_entry(entry.id, db, llm)

    insights = await db.get_insights_by_entry_id(entry.id)
    assert len(insights) == 1
    insight = insights[0]
    assert insight.entry_id == entry.id
