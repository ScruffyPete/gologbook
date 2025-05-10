from apps.db.interface import RepositoryUnit
from apps.llm.interface import LLMInterface
from apps.queue.interface import QueueMessage


async def process_entry(repo: RepositoryUnit, message: QueueMessage, llm: LLMInterface):
    """Process an entry and generate an insight.

    Args:
        entry_id (str): The ID of the entry to process.
        db (DBInterface): The database interface.
        llm (LLMInterface): The LLM interface.
    """
    try:
        try:
            entry_id = message.payload["entry_id"]
        except KeyError:
            print(f"No entry_id in message: {message.payload}")
            return

        entry = await repo.entry_repo.get_entry(entry_id)
        if entry is None:
            return

        insight = await llm.generate_insight(entry)
        await repo.insight_repo.create(insight)
    except Exception as e:
        print(f"Error processing entry: {e}")
