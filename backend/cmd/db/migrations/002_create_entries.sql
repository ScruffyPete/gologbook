CREATE TABLE IF NOT EXISTS entries (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    project_id UUID NOT NULL,
    body TEXT NOT NULL,

    CONSTRAINT fk_project
        FOREIGN KEY(project_id) 
        REFERENCES projects(id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_entries_project_id ON entries(project_id);

CREATE INDEX IF NOT EXISTS idx_entries_created_at ON entries(created_at DESC);

INSERT INTO migrations (name) VALUES ('002_create_entries.sql')
ON CONFLICT (name) DO NOTHING; 