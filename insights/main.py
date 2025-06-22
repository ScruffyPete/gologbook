import argparse
import asyncio
from contextlib import AsyncExitStack
import logging
import os
import time
import signal
from typing import Type

from apps.db.in_memory import InMemoryUnitOfWork
from apps.db.interface import UnitOfWorkInterface
from apps.db.postgres import PGUnitOfWork
from apps.queue.in_memory import InMemoryQueue
from apps.queue.interface import QueueInterface
from apps.queue.redis import RedisQueue
from apps.llm.in_memory import InMemoryLLM
from apps.llm.interface import LLMInterface
from apps.services.state import ServiceState, state_service
from apps.services.processor import process_project


shutdown_event = asyncio.Event()
logger = logging.getLogger(__name__)


def handle_shutdown():
    logger.info("Shutdown signal received")
    shutdown_event.set()
    

async def dispatcher(
    queue_interface: QueueInterface,
    work_queue: asyncio.Queue,
    cooldown: int,
    batch_size: int,
):
    while not shutdown_event.is_set():
        try:
            cuttoff_time = time.time() - cooldown
            project_ids = await queue_interface.pop_ready_projects(
                cuttoff_time=cuttoff_time,
                batch_size=batch_size,
            )
            
            for project_id in project_ids:
                await work_queue.put(project_id)
            
            await queue_interface.remove_processed_projects(project_ids)
                
        except Exception as e:
            logger.exception("Dispatcher error: %s", e)
            
        await asyncio.sleep(0.5)
    
    
async def worker(
    worker_id: str,
    queue_interface: QueueInterface,
    work_queue: asyncio.Queue,
    unit_of_work_cls: Type[UnitOfWorkInterface],
    llm: LLMInterface,
    state: ServiceState,
):
    while not shutdown_event.is_set():
        try:
            project_id = await work_queue.get()
            logger.info(f"Worker-{worker_id} processing project {project_id}")
            async with unit_of_work_cls.create() as uow:
                await process_project(project_id, uow, queue_interface, llm)
            state.mark_processed(project_id)
            
        except Exception as e:
            logger.exception(f"Worker-{worker_id} failed: {e}")
            state.mark_failed(str(e))
            
        finally:
            work_queue.task_done()
    
    
async def run_service(state: ServiceState, in_memory: bool):
    Queue = InMemoryQueue if in_memory else RedisQueue
    LLM = InMemoryLLM
    UnitOfWork = InMemoryUnitOfWork if in_memory else PGUnitOfWork

    logger.info(
        "Running insights with %s storage and queue",
        "in-memory" if in_memory else "real",
    )

    async with AsyncExitStack() as stack:
        await stack.enter_async_context(state_service(state))

        llm = await stack.enter_async_context(LLM.create())
        queue = await stack.enter_async_context(Queue.create())
        
        work_queue = asyncio.Queue(maxsize=100)
        cooldown = int(os.getenv("REDIS_PENDING_PROJECTS_COOLDOWN", 10))
        dispatcher_task = asyncio.create_task(dispatcher(queue, work_queue, cooldown=cooldown, batch_size=10))
        worker_count = 5
        worker_tasks = [
            asyncio.create_task(worker(str(n), queue, work_queue, UnitOfWork, llm, state))
            for n in range(worker_count)
        ]
        
        state.mark_started()
        
        await shutdown_event.wait()
        dispatcher_task.cancel()
        for worker_task in worker_tasks:
            worker_task.cancel()
        await asyncio.gather(*worker_tasks, dispatcher_task, return_exceptions=True)


async def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--in-memory", action="store_true", help="Run with in-memory storage")
    args = parser.parse_args()
    
    loop = asyncio.get_running_loop()
    loop.add_signal_handler(signal.SIGINT, handle_shutdown)
    loop.add_signal_handler(signal.SIGTERM, handle_shutdown)

    state = ServiceState()
    logger.info("Starting insights service...")
    await run_service(state, in_memory=args.in_memory)

    
if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    asyncio.run(main())
