package testutil

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"testing"
)

func OpenRedisForTest(t *testing.T) *redis.Client {
	t.Helper()

	host := "127.0.0.1"
	port := 36379
	if _, defined := os.LookupEnv("CI"); defined {
		port = 6379
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		t.Fatalf("failed to connect to redis: %v", err)
	}

	return client

}
