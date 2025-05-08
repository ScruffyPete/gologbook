CREATE TABLE IF NOT EXISTS entries (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    body TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_entries_project_created_at ON entries(project_id, created_at DESC);

INSERT INTO migrations (name) VALUES ('002_create_entries.sql')
ON CONFLICT (name) DO NOTHING; 