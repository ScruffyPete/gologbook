import argparse
import asyncio
from contextlib import AsyncExitStack
import logging
import signal

from apps.db.in_memory import InMemoryRepositoryUnit
from apps.db.postgres import PGRepositoryUnit
from apps.queue.in_memory import InMemoryQueue
from apps.queue.redis import RedisQueue
from apps.llm.in_memory import InMemoryLLM
from apps.services.healthz import ServiceState, health_service
from apps.services.processor import process_entry


shutdown_event = asyncio.Event()


def handle_shutdown():
    logger.info("Shutdown signal received")
    shutdown_event.set()


async def main(state: ServiceState):
    loop = asyncio.get_running_loop()
    loop.add_signal_handler(signal.SIGINT, handle_shutdown)
    loop.add_signal_handler(signal.SIGTERM, handle_shutdown)

    logger.info("Starting insights service...")

    parser = argparse.ArgumentParser()
    parser.add_argument(
        "--in-memory", action="store_true", help="Run with in-memory storage"
    )
    args = parser.parse_args()

    RepositoryUnit = InMemoryRepositoryUnit if args.in_memory else PGRepositoryUnit
    Queue = InMemoryQueue if args.in_memory else RedisQueue
    LLM = InMemoryLLM

    logger.info(
        "Running insights with %s storage and queue",
        "in-memory" if args.in_memory else "real",
    )

    async with AsyncExitStack() as stack:
        await stack.enter_async_context(health_service(state))

        repo = await stack.enter_async_context(RepositoryUnit.create())
        llm = await stack.enter_async_context(LLM.create())
        queue = await stack.enter_async_context(Queue.create())

        state.mark_started()
        logger.info("Waiting for entry to process")
        while not shutdown_event.is_set():
            try:
                message = await asyncio.wait_for(queue.pop(), timeout=2)
            except asyncio.TimeoutError:
                continue

            if message is None:
                continue

            logger.info("Processing entry %s", message)
            try:
                await process_entry(repo, message, llm)
            except Exception as e:
                logger.exception("Failed to process entry %s, error: %s", message, e)
            else:
                state.mark_processed()


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    logger = logging.getLogger(__name__)
    state = ServiceState()
    asyncio.run(main(state))
