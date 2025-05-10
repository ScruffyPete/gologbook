import os
import pytest
import aiohttp
import pytest_asyncio

from apps.services.state import ServiceState, state_service


@pytest_asyncio.fixture
async def client():
    async with aiohttp.ClientSession() as client:
        yield client


@pytest.mark.asyncio
async def test_state_service__started(client: aiohttp.ClientSession):
    state = ServiceState()
    state.mark_started()
    async with state_service(state):
        response = await client.get(f"http://localhost:{os.environ['PORT']}/healthz")
        assert response.status == 200

        response = await client.get(f"http://localhost:{os.environ['PORT']}/status")
        assert response.status == 200
        data = await response.json()
        assert data["has_started"]
        assert data["last_processed"] is None


@pytest.mark.asyncio
async def test_state_service__not_started(client: aiohttp.ClientSession):
    state = ServiceState()
    async with state_service(state):
        response = await client.get(f"http://localhost:{os.environ['PORT']}/healthz")
        assert response.status == 503

        response = await client.get(f"http://localhost:{os.environ['PORT']}/status")
        assert response.status == 200
        data = await response.json()
        assert not data["has_started"]
        assert data["last_processed"] is None
