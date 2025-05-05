from insights.db.interface import DBInterface
from insights.llm.interface import LLMInterface


async def process_entry(entry_id: str, db: DBInterface, llm: LLMInterface):
    """Process an entry and generate an insight.

    Args:
        entry_id (str): The ID of the entry to process.
        db (DBInterface): The database interface.
        llm (LLMInterface): The LLM interface.
    """
    entry = await db.get_entry(entry_id)
    if entry is None:
        return

    insight = await llm.generate_insight(entry)
    await db.save_insight(insight)
