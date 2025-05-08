from apps.db.in_memory import InMemoryDB
from apps.domain.entities import Entry
import uuid
from datetime import datetime
import pytest
import pytest_asyncio


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
    return InMemoryDB(entries=[entry])


@pytest.mark.asyncio
async def test_save_entry(db: InMemoryDB, entry: Entry):
    entry = Entry(
        id=uuid.uuid4(),
        project_id=uuid.uuid4(),
        body="Hello, world!",
        created_at=datetime.now(),
    )
    await db._save_entry(entry)
    db_entry = await db.get_entry(entry.id)
    assert db_entry == entry


@pytest.mark.asyncio
async def test_get_entry(db, entry):
    db_entry = await db.get_entry(entry.id)
    assert db_entry == entry
