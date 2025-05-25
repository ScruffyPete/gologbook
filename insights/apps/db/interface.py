from typing import Protocol
from apps.domain.entities import Entry, Document
from uuid import UUID


class EntryRepository(Protocol):
    async def get_project_entries(self, project_id: UUID) -> list[Entry]: ...


class DocumentRepository(Protocol):
    async def get_documents_by_entry_id(self, entry_id: UUID) -> list[Document]: ...

    async def create(self, document: Document) -> None: ...


class RepositoryBundleInterface(Protocol):
    entry_repo: EntryRepository
    document_repo: DocumentRepository
