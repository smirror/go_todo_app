package store

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"testing"
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
