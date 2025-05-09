from apps.db.interface import RepositoryUnit
from apps.llm.interface import LLMInterface
from uuid import UUID


async def process_entry(repo: RepositoryUnit, entry_id: UUID, llm: LLMInterface):
    """Process an entry and generate an insight.

    Args:
        entry_id (str): The ID of the entry to process.
        db (DBInterface): The database interface.
        llm (LLMInterface): The LLM interface.
    """
    entry = await repo.entry_repo.get_entry(entry_id)
    if entry is None:
        return

    insight = await llm.generate_insight(entry)
    await repo.insight_repo.create(insight)
