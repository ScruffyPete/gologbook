import uuid
import pytest
import pytest_asyncio
from apps.db.postgres import PGRepositoryBundle
from datetime import datetime
from apps.domain.entities import Entry, Document, Project


@pytest_asyncio.fixture
async def repo():
    async with PGRepositoryBundle.create() as repo:
        yield repo


@pytest_asyncio.fixture
async def entry(repo: PGRepositoryBundle):
    project_id = uuid.uuid4()
    project = Project(
        id=project_id,
        title="Test Project",
        created_at=datetime.now(),
    )
    await repo.project_repo._create_project(project)

    entry = Entry(
        id=uuid.uuid4(),
        project_id=project_id,
        body="Hello, world!",
        created_at=datetime.now(),
    )
    await repo.entry_repo._create_entry(entry)

    yield entry


@pytest_asyncio.fixture
async def insights(repo: PGRepositoryBundle, entry: Entry):
    insights = [
        Document(
            id=uuid.uuid4(),
            project_id=entry.project_id,
            entry_ids=[entry.id],
            body="Hello, world!",
            created_at=datetime.now(),
        )
        for _ in range(3)
    ]
    await repo.document_repo.create(insights[0])
    await repo.document_repo.create(insights[1])
    await repo.document_repo.create(insights[2])

    yield insights


@pytest.mark.db
@pytest.mark.asyncio
async def test_get_project_entries__valid_project(repo: PGRepositoryBundle, entry: Entry):
    db_entries = await repo.entry_repo.get_project_entries(entry.project_id)
    assert len(db_entries) == 1
    assert db_entries[0].id == entry.id
    assert db_entries[0].project_id == entry.project_id
    assert db_entries[0].body == entry.body
    assert db_entries[0].created_at == entry.created_at


@pytest.mark.db
@pytest.mark.asyncio
async def test_get_project_entries__invalid_project(repo: PGRepositoryBundle):
    db_entry = await repo.entry_repo.get_project_entries(uuid.uuid4())
    assert len(db_entry) == 0


@pytest.mark.db
@pytest.mark.asyncio
async def test_get_insights_by_entry_id__valid_entry(
    repo: PGRepositoryBundle, entry: Entry, insights: list[Document]
):
    db_insights = await repo.document_repo.get_documents_by_entry_id(entry.id)
    assert len(db_insights) == 3
    assert db_insights[0].id == insights[0].id
    assert db_insights[1].id == insights[1].id
    assert db_insights[2].id == insights[2].id


@pytest.mark.db
@pytest.mark.asyncio
async def test_get_insights_by_entry_id__invalid_entry(repo: PGRepositoryBundle):
    db_insights = await repo.document_repo.get_documents_by_entry_id(uuid.uuid4())
    assert len(db_insights) == 0


@pytest.mark.db
@pytest.mark.asyncio
async def test_create_insight__valid_insight(repo: PGRepositoryBundle, entry: Entry):
    insight = Document(
        id=uuid.uuid4(),
        project_id=entry.project_id,
        entry_ids=[entry.id],
        body="Hello, world!",
        created_at=datetime.now(),
    )
    await repo.document_repo.create(insight)

    db_insights = await repo.document_repo.get_documents_by_entry_id(entry.id)
    assert len(db_insights) == 1
    assert db_insights[0].id == insight.id
    assert db_insights[0].project_id == insight.project_id
    assert db_insights[0].entry_ids == insight.entry_ids
    assert db_insights[0].body == insight.body


@pytest.mark.db
@pytest.mark.asyncio
async def test_create_insight__invalid_insight(repo: PGRepositoryBundle):
    insight = Document(
        id=uuid.uuid4(),
        project_id=uuid.uuid4(),
        entry_ids=[uuid.uuid4()],
        body="Hello, world!",
        created_at=datetime.now(),
    )
    with pytest.raises(Exception):
        await repo.document_repo.create(insight)
