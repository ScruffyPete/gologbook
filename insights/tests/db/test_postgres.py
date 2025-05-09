import uuid
import pytest
import pytest_asyncio
from apps.db.postgres import PGRepositoryUnit
from datetime import datetime
from apps.domain.entities import Entry, Insight, Project


@pytest_asyncio.fixture
async def repo():
    async with PGRepositoryUnit.create() as repo:
        yield repo


@pytest_asyncio.fixture
async def entry(repo: PGRepositoryUnit):
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
async def insights(repo: PGRepositoryUnit, entry: Entry):
    insights = [
        Insight(
            id=uuid.uuid4(),
            project_id=entry.project_id,
            entry_ids=[entry.id],
            body="Hello, world!",
            created_at=datetime.now(),
        )
        for _ in range(3)
    ]
    await repo.insight_repo.create(insights[0])
    await repo.insight_repo.create(insights[1])
    await repo.insight_repo.create(insights[2])

    yield insights


@pytest.mark.db
@pytest.mark.asyncio
async def test_get_entry__valid_entry(repo: PGRepositoryUnit, entry: Entry):
    db_entry = await repo.entry_repo.get_entry(entry.id)
    assert db_entry.id == entry.id
    assert db_entry.project_id == entry.project_id
    assert db_entry.body == entry.body
    assert db_entry.created_at == entry.created_at


@pytest.mark.db
@pytest.mark.asyncio
async def test_get_entry__invalid_entry(repo: PGRepositoryUnit):
    db_entry = await repo.entry_repo.get_entry(uuid.uuid4())
    assert db_entry is None


@pytest.mark.db
@pytest.mark.asyncio
async def test_get_insights_by_entry_id__valid_entry(
    repo: PGRepositoryUnit, entry: Entry, insights: list[Insight]
):
    db_insights = await repo.insight_repo.get_insights_by_entry_id(entry.id)
    assert len(db_insights) == 3
    assert db_insights[0].id == insights[0].id
    assert db_insights[1].id == insights[1].id
    assert db_insights[2].id == insights[2].id


@pytest.mark.db
@pytest.mark.asyncio
async def test_get_insights_by_entry_id__invalid_entry(repo: PGRepositoryUnit):
    db_insights = await repo.insight_repo.get_insights_by_entry_id(uuid.uuid4())
    assert len(db_insights) == 0


@pytest.mark.db
@pytest.mark.asyncio
async def test_create_insight__valid_insight(repo: PGRepositoryUnit, entry: Entry):
    insight = Insight(
        id=uuid.uuid4(),
        project_id=entry.project_id,
        entry_ids=[entry.id],
        body="Hello, world!",
        created_at=datetime.now(),
    )
    await repo.insight_repo.create(insight)

    db_insights = await repo.insight_repo.get_insights_by_entry_id(entry.id)
    assert len(db_insights) == 1
    assert db_insights[0].id == insight.id
    assert db_insights[0].project_id == insight.project_id
    assert db_insights[0].entry_ids == insight.entry_ids
    assert db_insights[0].body == insight.body


@pytest.mark.db
@pytest.mark.asyncio
async def test_create_insight__invalid_insight(repo: PGRepositoryUnit):
    insight = Insight(
        id=uuid.uuid4(),
        project_id=uuid.uuid4(),
        entry_ids=[uuid.uuid4()],
        body="Hello, world!",
        created_at=datetime.now(),
    )
    with pytest.raises(Exception):
        await repo.insight_repo.create(insight)
