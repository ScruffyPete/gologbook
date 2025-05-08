CREATE TABLE insights (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    entry_ids JSONB NOT NULL,
    body TEXT NOT NULL
);

CREATE INDEX idx_insights_project_created_at ON insights(project_id, created_at DESC);

INSERT INTO migrations (name) VALUES ('004_create_insights.sql')
ON CONFLICT (name) DO NOTHING;