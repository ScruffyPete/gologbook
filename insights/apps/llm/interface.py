from typing import Protocol
import uuid
from apps.domain.entities import Entry, Document


class LLMInterface(Protocol):
    async def compile_messages(self, project_id: uuid.UUID, entries: list[Entry]) -> Document: ...
