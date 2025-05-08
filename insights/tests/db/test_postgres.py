import uuid
import pytest
import pytest_asyncio
from apps.db.postgres import PostgresDB
from datetime import datetime

from apps.domain.entities import Entry, Project


@pytest_asyncio.fixture
async def db():
    db = PostgresDB()
    try:
        await db.connect()
    except Exception:
        pytest.skip("Postgres is not running or unreachable")
    yield db
    await db.close()


@pytest_asyncio.fixture
async def entry(db: PostgresDB):
    project_id = uuid.uuid4()
    project = Project(
        id=project_id,
        title="Test Project",
        created_at=datetime.now(),
    )
    await db._create_project(project)

    entry = Entry(
        id=uuid.uuid4(),
        project_id=project_id,
        body="Hello, world!",
        created_at=datetime.now(),
    )
    await db._create_entry(entry)

    return entry


@pytest.mark.asyncio
async def test_get_entry(db: PostgresDB, entry: Entry):
    db_entry = await db.get_entry(entry.id)
    assert db_entry.id == entry.id
    assert db_entry.project_id == entry.project_id
    assert db_entry.body == entry.body
    assert db_entry.created_at == entry.created_at
