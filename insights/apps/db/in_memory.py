from collections import defaultdict
from contextlib import asynccontextmanager
from apps.domain.entities import Entry, Document
from uuid import UUID


class InMemoryUnitOfWork:
    def __init__(self, entries=None, documents=None):
        self.entry_repo = InMemoryEntryRepository(entries or [])
        self.document_repo = InMemoryDocumentRepository(documents or [])

    @classmethod
    @asynccontextmanager
    async def create(cls, *args, **kwargs):
        yield cls(*args, **kwargs)


class InMemoryEntryRepository:
    entries: defaultdict[UUID, list[Entry]]
    
    def __init__(self, entries: list[Entry] = []):
        self.entries = defaultdict(lambda: list())

        for entry in entries:
            self.entries[entry.project_id].append(entry)

    async def get_project_entries(self, project_id: UUID) -> list[Entry]:
        return self.entries.get(project_id, [])


class InMemoryDocumentRepository:
    def __init__(self, documents: list[Document] = []):
        self.documents = {document.id: document for document in documents}

    async def get_documents_by_entry_id(self, entry_id: UUID) -> list[Document]:
        return [
            document
            for document in self.documents.values()
            if entry_id in document.entry_ids
        ]

    async def create(self, document: Document) -> None:
        self.documents[document.id] = document


class InMemoryRepositoryBundle:
    def __init__(self, entries: list[Entry] = [], documents: list[Document] = []):
        self.entry_repo = InMemoryEntryRepository(entries)
        self.document_repo = InMemoryDocumentRepository(documents)

    @classmethod
    @asynccontextmanager
    async def create(cls):
        repo = cls()
        yield repo
