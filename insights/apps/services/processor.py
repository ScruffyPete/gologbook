from uuid import UUID
from apps.db.interface import RepositoryBundleInterface
from apps.llm.interface import LLMInterface


async def process_entry(repo: RepositoryBundleInterface, project_id: UUID, llm: LLMInterface):
    entries = await repo.entry_repo.get_project_entries(project_id)
    if not len(entries) > 0:
        return

    document = await llm.compile_messages(project_id, entries)
    await repo.document_repo.create(document)
