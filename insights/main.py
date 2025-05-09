import asyncio
import logging

from apps.db.in_memory import InMemoryRepositoryUnit
from apps.queue.in_memory import InMemory
from apps.llm.in_memory import InMemoryLLM
from apps.services.processor import process_entry


logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


async def main():
    repo = InMemoryRepositoryUnit()
    llm = InMemoryLLM()
    queue = InMemory()

    while True:
        logger.info("Waiting for entry to process")
        entry_id = await queue.pop()
        if entry_id is None:
            logger.info("No entry to process")
            continue
        await process_entry(repo, entry_id, llm)


if __name__ == "__main__":
    asyncio.run(main())
