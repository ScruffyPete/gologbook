from typing import Type
import uuid

from apps.db.interface import UnitOfWorkInterface
from apps.domain.entities import Document
from apps.llm.interface import LLMInterface
from apps.queue.interface import QueueInterface

from datetime import datetime


class CreateDocumentExeption(Exception):
    def __init__(self, original_exception: Exception):
        self.original_exception = original_exception
        super().__init__(f"Failed to create document: {str(original_exception)}")


async def process_project(
    project_id: uuid.UUID, 
    unit_of_work: UnitOfWorkInterface,
    queue: QueueInterface,
    llm: LLMInterface,
):
    await queue.clear_project_stream(project_id)

    entries = await unit_of_work.entry_repo.get_project_entries(project_id)
    if not entries:
        return

    document_buffer = []
    async for token in llm.stream_document(project_id, entries):
        await queue.publish_project_token(project_id, token)
        document_buffer.append(token)

    document = Document(
        id=uuid.uuid4(),
        project_id=project_id,
        entry_ids=[entry.id for entry in entries],
        body="".join(document_buffer),
        created_at=datetime.now(),
    )
    
    try: 
        await unit_of_work.document_repo.create(document)
    except Exception as e:
        raise CreateDocumentExeption(e)
