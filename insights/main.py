import asyncio
from contextlib import AsyncExitStack
import logging

from apps.db.in_memory import InMemoryRepositoryUnit
from apps.queue.in_memory import InMemoryQueue
from apps.llm.in_memory import InMemoryLLM
from apps.services.processor import process_entry


logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


async def main():
    async with AsyncExitStack() as stack:
        repo = await stack.enter_async_context(InMemoryRepositoryUnit.create())
        llm = await stack.enter_async_context(InMemoryLLM.create())
        queue = await stack.enter_async_context(InMemoryQueue.create())

        while True:
            logger.info("Waiting for entry to process")
            entry_id = await queue.pop()
            if entry_id is None:
                logger.info("No entry to process")
                continue
            await process_entry(repo, entry_id, llm)


if __name__ == "__main__":
    asyncio.run(main())
