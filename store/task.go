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
	//goland:noinspection ALL
	sql := `SELECT id, title, status, created, modified FROM task;`
	if err := db.SelectContext(ctx, &tasks, sql); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) AddTask(
	ctx context.Context, db Execer, t *entity.Task,
) error {
	t.Created = r.Clocker.Now()
	t.Modified = r.Clocker.Now()
	//goland:noinspection ALL
	sql := `INSERT INTO task
		(title, status, created, modified) 
		VALUES (?,?,?,?);`
	id, err := db.ExecContext(
		ctx, sql, t.Title, t.Status,
		t.Created, t.Modified,
	)
	if err != nil {
		return err
	}
	t.ID = entity.TaskID(id)
	return nil
}
