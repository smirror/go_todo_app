package store

import (
	"context"
	"github.com/jmoiron/sqlx"
	"todo_app/entity"
)

func (r *Repository) ListTasks(
	ctx context.Context, db *sqlx.DB,
) (entity.Tasks, error) {
	tasks := make(entity.Tasks, 0)
	sql := `SELECT id, title, status, created, modified FROM task;`
	if err := db.SelectContext(ctx, &tasks, sql); err != nil {
		return nil, err
	}
	return tasks, nil
}
