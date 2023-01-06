package store

import (
	"golang.org/x/net/context"
	"testing"
	"todo_app/entity"
	"todo_app/testutil"
)

func TestKVS_Save(t *testing.T) {
	t.Parallel()

	cli := testutil.OpenRedisForTest(t)

	sut := &KVS{Cli: cli}
	key := "TestKVS_Save"
	uid := entity.UserID(1234)
	ctx := context.Background()
	t.Cleanup(func() {
		cli.Del(ctx, key)
	})

	if err := sut.Save(ctx, key, uid); err != nil {
		t.Fatalf("failed to save: %v", err)
	}
}
