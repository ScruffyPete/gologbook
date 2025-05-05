import asyncio
import logging

from queue.in_memory import InMemoryQueue
from db.in_memory import InMemoryDB
from llm.in_memory import InMemoryLLM
from services.processor import process_entry

logger = logging.getLogger(__name__)


async def main():
    db = InMemoryDB()
    llm = InMemoryLLM()
    queue = InMemoryQueue()

    while True:
        entry_id = await queue.pop()
        if entry_id is None:
            logger.info("No entry to process")
            continue
        await process_entry(entry_id, db, llm)


if __name__ == "__main__":
    asyncio.run(main())
