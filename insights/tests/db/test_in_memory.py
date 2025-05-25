from apps.db.in_memory import (
    InMemoryEntryRepository,
    InMemoryDocumentRepository,
)
from apps.domain.entities import Entry, Document, Project
import uuid
from datetime import datetime
import pytest


@pytest.fixture
def project():
    return Project(
        id=uuid.uuid4(),
        title="Test Project",
        created_at=datetime.now(),
    )


@pytest.fixture
def entry(project: Project):
    return Entry(
        id=uuid.uuid4(),
        project_id=project.id,
        body="Hello, world!",
        created_at=datetime.now(),
    )


@pytest.fixture
def insights(entry: Entry):
    return [
        Document(
            id=uuid.uuid4(),
            project_id=entry.project_id,
            entry_ids=[entry.id],
            body="Hello!",
            created_at=datetime.now(),
        ),
        Document(
            id=uuid.uuid4(),
            project_id=entry.project_id,
            entry_ids=[entry.id],
            body="World!",
            created_at=datetime.now(),
        ),
    ]


@pytest.mark.asyncio
async def test_entry_repository_get_entry__valid_entry(entry: Entry):
    entry_repository = InMemoryEntryRepository(entries=[entry])
    db_entries = await entry_repository.get_project_entries(entry.project_id)
    assert db_entries == [entry]


@pytest.mark.asyncio
async def test_entry_repository_get_entry__invalid_entry(entry: Entry):
    entry_repository = InMemoryEntryRepository(entries=[])
    db_entries = await entry_repository.get_project_entries(entry.project_id)
    assert db_entries == []


@pytest.mark.asyncio
async def test_insight_repository_get_insights_by_entry_id__valid_entry(
    entry: Entry, insights: list[Document]
):
    insight_repository = InMemoryDocumentRepository(documents=insights)
    db_insights = await insight_repository.get_documents_by_entry_id(entry.id)
    assert db_insights == insights


@pytest.mark.asyncio
async def test_insight_repository_get_insights_by_entry_id__invalid_entry(
    insights: list[Document],
):
    insight_repository = InMemoryDocumentRepository(documents=insights)
    db_insights = await insight_repository.get_documents_by_entry_id(uuid.uuid4())
    assert db_insights == []


@pytest.mark.asyncio
async def test_insight_repository_save_insight__valid_insight(entry: Entry):
    insight_repository = InMemoryDocumentRepository(documents=[])
    new_insight = Document(
        id=uuid.uuid4(),
        project_id=entry.project_id,
        entry_ids=[entry.id],
        body="Hello, world!",
        created_at=datetime.now(),
    )
    await insight_repository.create(new_insight)
    db_insights = await insight_repository.get_documents_by_entry_id(entry.id)
    assert db_insights == [new_insight]
