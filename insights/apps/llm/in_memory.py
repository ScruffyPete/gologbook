import asyncio
from typing import AsyncIterator
import uuid
from contextlib import asynccontextmanager
from datetime import datetime

from apps.domain.entities import Entry, Document


class InMemoryLLM:
    @classmethod
    @asynccontextmanager
    async def create(cls):
        llm = cls()
        yield llm

    async def compile_document(self, project_id: uuid.UUID, entries: list[Entry]) -> Document:
        document_body = "\n\n".join(
            f"Insight for entry {entry.id}: {entry.body[:100]}" for entry in entries
        )
        return Document(
            id=uuid.uuid4(),
            project_id=project_id,
            entry_ids=[entry.id for entry in entries],
            body=document_body,
            created_at=datetime.now(),
        )

    async def stream_document(self, project_id: uuid.UUID, entries: list[Entry]) -> AsyncIterator:
        for entry in entries:
            await asyncio.sleep(.1)
            text = f"Insight for entry {entry.id}: {entry.body[:100]}\n\n"
            yield text