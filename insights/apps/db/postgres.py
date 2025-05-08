import asyncpg
import os

from typing import Any

from apps.domain.entities import Entry, Insight, Project


class PostgresDB:
    def __init__(self):
        self.database_url = os.getenv("DATABASE_URL")

    async def connect(self):
        self.conn = await asyncpg.connect(self.database_url)

    async def close(self):
        await self.conn.close()

    async def get_entry(self, entry_id: str) -> Entry | None:
        query = """
            SELECT * FROM entries WHERE id = $1
        """
        result = await self._execute(query, entry_id)
        if len(result) == 0:
            return None
        return Entry(**result[0])

    async def save_entry(self, entry: Entry) -> Entry:
        query = """
            INSERT INTO entries (id, created_at, project_id, body) VALUES ($1, $2, $3, $4)
        """
        await self._execute(
            query, entry.id, entry.created_at, entry.project_id, entry.body
        )
        return entry

    async def get_insights_by_entry_id(self, entry_id: str) -> list[Insight]:
        pass

    async def save_insight(self, insight: Insight) -> None:
        pass

    async def _execute(self, query: str, *args) -> list[Any]:
        return await self.conn.fetch(query, *args)

    async def _create_project(self, project: Project) -> None:
        query = """
            INSERT INTO projects (id, title, created_at) VALUES ($1, $2, $3)
        """
        await self._execute(query, project.id, project.title, project.created_at)

    async def _create_entry(self, entry: Entry) -> None:
        query = """
            INSERT INTO entries (id, created_at, project_id, body) VALUES ($1, $2, $3, $4)
        """
        await self._execute(
            query, entry.id, entry.created_at, entry.project_id, entry.body
        )
