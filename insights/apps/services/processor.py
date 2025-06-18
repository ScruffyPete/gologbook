import uuid

from apps.db.interface import RepositoryBundleInterface
from apps.domain.entities import Document
from apps.llm.interface import LLMInterface
from apps.queue.interface import QueueInterface

from datetime import datetime



async def process_project(
    project_id: uuid.UUID, 
    repo: RepositoryBundleInterface, 
    queue: QueueInterface,
    llm: LLMInterface,
):
    entries = await repo.entry_repo.get_project_entries(project_id)
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
        
    await repo.document_repo.create(document)
    await queue.clear_project_stream(project_id)

