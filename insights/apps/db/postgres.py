from contextlib import asynccontextmanager
from uuid import UUID
import asyncpg
import json
import os

from apps.domain.entities import Entry, Document, Project


@asynccontextmanager
async def db_connection(dsn: str | None = None):
    """
    Context manager for a database connection.
    """
    if dsn is None:
        dsn = os.getenv("DATABASE_URL")
    if dsn is None:
        raise ValueError("DATABASE_URL is not set")
    conn = await asyncpg.connect(dsn)
    try:
        yield conn
    finally:
        await conn.close()


class PGBaseRepository:
    def __init__(self, conn: asyncpg.Connection):
        self.conn = conn


class PGProjectRepository(PGBaseRepository):
    async def _create_project(self, project: Project) -> None:
        query = """
            INSERT INTO projects (id, title, created_at) VALUES ($1, $2, $3)
        """
        await self.conn.execute(query, project.id, project.title, project.created_at)


class PGEntryRepository(PGBaseRepository):
    async def get_project_entries(self, project_id: UUID) -> list[Entry]:
        query = """
            SELECT * FROM entries WHERE project_id = $1
        """
        results = await self.conn.fetch(query, project_id)
        return [Entry(**row) for row in results]

    async def _create_entry(self, entry: Entry) -> None:
        query = """
            INSERT INTO entries (id, created_at, project_id, body) VALUES ($1, $2, $3, $4)
        """
        await self.conn.execute(
            query, entry.id, entry.created_at, entry.project_id, entry.body
        )


class PGDocumentRepository(PGBaseRepository):
    async def get_documents_by_entry_id(self, entry_id: UUID) -> list[Document]:
        query = """
            SELECT * FROM insights WHERE entry_ids @> to_jsonb(ARRAY[$1::UUID])
        """
        results = await self.conn.fetch(query, entry_id)
        return [
            Document(
                id=row["id"],
                project_id=row["project_id"],
                entry_ids=[UUID(eid) for eid in json.loads(row["entry_ids"])],
                body=row["body"],
                created_at=row["created_at"],
            )
            for row in results
        ]

    async def create(self, document: Document) -> None:
        query = """
            INSERT INTO insights (id, created_at, project_id, entry_ids, body) VALUES ($1, $2, $3, $4, $5)
        """
        await self.conn.execute(
            query,
            document.id,
            document.created_at,
            document.project_id,
            json.dumps([str(entry_id) for entry_id in document.entry_ids]),
            document.body,
        )


class PGRepositoryBundle:
    def __init__(self, conn: asyncpg.Connection):
        self.project_repo = PGProjectRepository(conn)
        self.entry_repo = PGEntryRepository(conn)
        self.document_repo = PGDocumentRepository(conn)

    @classmethod
    @asynccontextmanager
    async def create(cls):
        dsn = os.getenv("DATABASE_URL")
        if dsn is None:
            raise ValueError("DATABASE_URL is not set")
        async with db_connection(dsn) as conn:
            yield cls(conn)
