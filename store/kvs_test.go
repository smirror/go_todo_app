package store

import (
	"errors"
	"golang.org/x/net/context"
	"testing"
	"time"
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

func TestKVS_Load(t *testing.T) {
	t.Parallel()

	cli := testutil.OpenRedisForTest(t)
	sut := &KVS{Cli: cli}

	t.Run("not found", func(t *testing.T) {
		t.Parallel()

		key := "TestKVS_Load"
		uid := entity.UserID(1234)
		ctx := context.Background()
		cli.Set(ctx, key, int64(uid), 30*time.Minute)
		t.Cleanup(func() {
			cli.Del(ctx, key)
		})

		got, err := sut.Load(ctx, key)
		if err != nil {
			t.Fatalf("want no error, but got %v", err)
		}
		if got != uid {
			t.Errorf("want %d, but got %d", uid, got)
		}

	})

	t.Run("not found", func(t *testing.T) {
		t.Parallel()

		key := "TestKVS_Save_not_found"
		ctx := context.Background()
		got, err := sut.Load(ctx, key)
		if err == nil || !errors.Is(err, ErrNotFound) {
			t.Errorf("want %v, but got %v(value = %d", ErrNotFound, err, got)
		}
	})
}
