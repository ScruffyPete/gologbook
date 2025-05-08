import asyncio
import logging

from apps.queue.in_memory import InMemory
from apps.db.in_memory import InMemoryDB
from apps.llm.in_memory import InMemoryLLM
from apps.services.processor import process_entry

logger = logging.getLogger(__name__)


async def main():
    db = InMemoryDB()
    llm = InMemoryLLM()
    queue = InMemory()

    while True:
        entry_id = await queue.pop()
        if entry_id is None:
            logger.info("No entry to process")
            continue
        await process_entry(entry_id, db, llm)


if __name__ == "__main__":
    asyncio.run(main())
