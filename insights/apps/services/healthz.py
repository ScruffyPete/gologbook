import asyncio
import logging
import os

from aiohttp import web
from contextlib import asynccontextmanager
from dataclasses import dataclass


logger = logging.getLogger(__name__)


@dataclass
class ServiceState:
    has_started: bool = False
    last_processed: float | None = None

    def mark_started(self):
        self.has_started = True

    def mark_processed(self):
        self.last_processed = asyncio.get_running_loop().time()


@asynccontextmanager
async def health_service(state: ServiceState):
    app = web.Application()

    async def healthz(request):
        if not state.has_started:
            return web.Response(status=503, text="Service not started")
        return web.Response(text="OK")

    app.add_routes([web.get("/healthz", healthz)])

    runner = web.AppRunner(app)
    await runner.setup()
    site = web.TCPSite(runner, port=int(os.environ["PORT"]))
    await site.start()

    logger.info("Healthz service started at %s", site.name)

    try:
        yield
    finally:
        await runner.cleanup()
