package store

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"testing"
	"todo_app/clock"
	"todo_app/entity"
	"todo_app/testutil"
)

func TestRepository_ListTasks(t *testing.T) {
	ctx := context.Background()
	// entity.Taskをさくせいする他のテスト家＾スト混ざるとテストが失敗する
	// そのためトランザクションをはることでこのテストケースの中だけのテーブル状態にする
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	// このテストケースが完了したら元に戻す
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatalf("failed to begin transaction: %v", err)
	}
	wants := prepareTasks(t, tx)

	sut := &Repository{}
	gots, err := sut.ListTasks(ctx, tx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if d := cmp.Diff(gots, wants); len(d) != 0 {
		t.Errorf("ListTasks() mismatch (-got +want):\n%s", d)
	}
}

func prepareTasks(ctx context.Context, t *testing.T, con Execer) entity.Tasks {
	t.Helper()
	// 一度に綺麗にしておく
	if _, err := con.ExecContext(ctx, "DELETE FROM task;"); err != nil {
		t.Fatalf("failed to delete all tasks: %v", err)
	}
	c := clock.FixedClocker{}
	wants := entity.Tasks{
		{
			Title: "want task 1", Status: "todo",
			Created: c.Now(), Modified: c.Now(),
		},
		{
			Title: "want task 2", Status: "doing",
			Created: c.Now(), Modified: c.Now(),
		},
		{
			Title: "want task 3", Status: "done",
			Created: c.Now(), Modified: c.Now(),
		},
	}

	result, err := con.ExecContext(ctx,
		"INSERT INTO task (title, status, created, modified) VALUES (?,?,?,?);",
		wants[0].Title, wants[0].Status, wants[0].Created, wants[0].Modified,
		wants[1].Title, wants[1].Status, wants[1].Created, wants[1].Modified,
		wants[2].Title, wants[2].Status, wants[2].Created, wants[2].Modified
	)
	if err != nil {
		t.Fatalf("failed to insert tasks: %v", err)
	}
	id. err := result.LastInsertId()
	if err != nil {
		t.Fatalf("failed to get last insert id: %v", err)
	}
	wants[0].ID = entity.TaskID(id)
	wants[1].ID = entity.TaskID(id + 1)
	wants[2].ID = entity.TaskID(id + 2)
	return wants
}
