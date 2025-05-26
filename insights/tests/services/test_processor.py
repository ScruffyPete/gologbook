import pytest
import uuid
from datetime import datetime

from apps.db.in_memory import InMemoryRepositoryBundle
from apps.domain.entities import Entry
from apps.db.interface import RepositoryBundleInterface
from apps.llm.interface import LLMInterface
from apps.queue.in_memory import InMemoryQueue
from apps.queue.interface import QueueInterface
from apps.services.processor import process_project
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
def queue() -> InMemoryQueue:
    return InMemoryQueue()


@pytest.fixture
def repo(entry: Entry) -> InMemoryRepositoryBundle:
    return InMemoryRepositoryBundle(entries=[entry])


@pytest.mark.asyncio
async def test_processor(
    repo: RepositoryBundleInterface, queue: QueueInterface, llm: LLMInterface, entry: Entry
):
    await process_project(entry.project_id, repo, queue, llm)

    documents = await repo.document_repo.get_documents_by_entry_id(entry.id)
    assert len(documents) == 1

    document = documents[0]
    assert document.entry_ids == [entry.id]
