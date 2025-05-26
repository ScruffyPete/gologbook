from typing import AsyncIterator, Protocol
import uuid
from apps.domain.entities import Entry, Document


class LLMInterface(Protocol):
    async def compile_document(self, project_id: uuid.UUID, entries: list[Entry]) -> Document: ...
    
    async def stream_document(self, project_id: uuid.UUID, entries: list[Entry]) -> AsyncIterator: ...
