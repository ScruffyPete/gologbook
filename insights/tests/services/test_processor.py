import pytest
import uuid
from datetime import datetime

import pytest_asyncio

from apps.db.in_memory import InMemoryUnitOfWork
from apps.domain.entities import Entry
from apps.db.interface import UnitOfWorkInterface
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


@pytest_asyncio.fixture
async def uow(entry: Entry): 
    async with InMemoryUnitOfWork.create(entries=[entry]) as _uow: 
        yield _uow
        

@pytest.mark.asyncio
async def test_processor(uow: UnitOfWorkInterface, queue: QueueInterface, llm: LLMInterface, entry: Entry):
    await process_project(entry.project_id, uow, queue, llm)

    documents = await uow.document_repo.get_documents_by_entry_id(entry.id)
    assert len(documents) == 1

    document = documents[0]
    assert document.entry_ids == [entry.id]
