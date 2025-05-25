import asyncio
import logging
import os
from uuid import UUID

from aiohttp import web
from contextlib import asynccontextmanager
from dataclasses import dataclass

logger = logging.getLogger(__name__)


@dataclass
class ServiceState:
    has_started: bool = False
    last_processed: float | None = None
    last_processed_project_id: UUID | None = None
    failure_count: int = 0
    last_failure: float | None = None
    last_error: str | None = None

    def mark_started(self):
        self.has_started = True

    def mark_processed(self, project_id: UUID):
        self.last_processed = asyncio.get_running_loop().time()
        self.last_processed_project_id = project_id

    def mark_failed(self, error: str):
        self.failure_count += 1
        self.last_failure = asyncio.get_running_loop().time()
        self.last_error = error

    def to_json(self) -> dict:
        return {
            "has_started": self.has_started,
            "last_processed": self.last_processed,
            "last_processed_project_id": self.last_processed_project_id,
            "failure_count": self.failure_count,
            "last_failure": self.last_failure,
            "last_error": self.last_error,
        }


@asynccontextmanager
async def state_service(state: ServiceState):
    app = web.Application()

    async def healthz(request):
        if not state.has_started:
            return web.Response(status=503, text="Service not started")
        return web.Response(text="OK")

    async def status(request):
        return web.json_response(state.to_json())

    app.add_routes([web.get("/healthz", healthz), web.get("/status", status)])

    runner = web.AppRunner(app)
    await runner.setup()
    site = web.TCPSite(runner, port=int(os.environ["PORT"]))
    await site.start()

    logger.info("Healthz service started at %s", site.name)

    try:
        yield
    finally:
        await runner.cleanup()
