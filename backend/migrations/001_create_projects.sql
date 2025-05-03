CREATE TABLE IF NOT EXISTS projects (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_projects_created_at ON projects(created_at DESC);

INSERT INTO migrations (name) VALUES ('001_create_projects.sql')
ON CONFLICT (name) DO NOTHING; 