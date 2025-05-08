package postgres

import (
	"database/sql"
	"encoding/json"

	"github.com/ScruffyPete/gologbook/internal/domain"
	_ "github.com/lib/pq"
)

type InsightRepository struct {
	db *sql.DB
}

func NewInsightRepository(db *sql.DB) *InsightRepository {
	return &InsightRepository{db: db}
}

func (repo *InsightRepository) ListInsights(projectID string) ([]*domain.Insight, error) {
	rows, err := repo.db.Query(
		"SELECT id, created_at, project_id, entry_ids, body FROM insights WHERE project_id = $1",
		projectID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	insights := make([]*domain.Insight, 0)
	for rows.Next() {
		var insight domain.Insight
		var entryIDsRaw []byte
		if err := rows.Scan(&insight.ID, &insight.CreatedAt, &insight.ProjectID, &entryIDsRaw, &insight.Body); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(entryIDsRaw, &insight.EntryIDs); err != nil {
			return nil, err
		}
		insights = append(insights, &insight)
	}
	return insights, nil
}
