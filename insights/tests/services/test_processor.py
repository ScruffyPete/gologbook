import pytest
import uuid
from datetime import datetime

from apps.db.in_memory import InMemoryRepositoryBundle
from apps.domain.entities import Entry
from apps.db.interface import RepositoryBundleInterface
from apps.llm.interface import LLMInterface
from apps.services.processor import process_entry
from apps.llm.in_memory import InMemoryLLM


@pytest.fixture
def entry() -> Entry:
    return Entry(
        id=uuid.uuid4(),
        project_id=uuid.uuid4(),
        body="Hello, world!",
        created_at=datetime.now(),
    )


@pytest.fixture
def llm() -> InMemoryLLM:
    return InMemoryLLM()


@pytest.fixture
def repo(entry: Entry) -> InMemoryRepositoryBundle:
    return InMemoryRepositoryBundle(entries=[entry])


@pytest.mark.asyncio
async def test_processor(
    repo: RepositoryBundleInterface, llm: LLMInterface, entry: Entry
):
    await process_entry(repo, entry.project_id, llm)

    documents = await repo.document_repo.get_documents_by_entry_id(entry.id)
    assert len(documents) == 1

    document = documents[0]
    assert document.entry_ids == [entry.id]
