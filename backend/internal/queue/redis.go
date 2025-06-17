package queue

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisQueue struct {
	client *redis.Client
}

func NewRedisQueue() (*RedisQueue, error) {
	db, err := strconv.Atoi(os.Getenv("REDIS_DEAULT_DB"))
	if err != nil {
		return nil, err
	}

	addr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return &RedisQueue{client: client}, nil
}

func (q *RedisQueue) PushPendingProject(ctx context.Context, key string, projectID string) error {
	_, err := q.client.ZAdd(ctx, key, redis.Z{
		Member: projectID,
		Score:  float64(time.Now().Unix()),
	}).Result()

	return err
}

func (q *RedisQueue) SubscribeForDocumentTokens(ctx context.Context, projectID string) <-chan string {
	streamNamePrefix := os.Getenv("REDIS_LLM_STREAM_CHANNEL_PREFIX")
	streamName := fmt.Sprintf("%s:%s", streamNamePrefix, projectID)

	out := make(chan string, 100)

	go func() {
		defer close(out)
		lastID := "0"

		for {
			select {
			case <-ctx.Done():
				return
			default:
				streams, err := q.client.XRead(ctx, &redis.XReadArgs{
					Streams: []string{streamName, lastID},
					Block:   1 * time.Second,
				}).Result()

				if err != nil && err != redis.Nil {
					slog.Error("XREAD error", "error", err)
					time.Sleep(500 * time.Millisecond)
					continue
				}

				if len(streams) == 0 {
					continue
				}

				for _, stream := range streams {
					for _, msg := range stream.Messages {
						token, ok := msg.Values["token"].(string)
						if !ok {
							continue
						}
						out <- token
						lastID = msg.ID
					}
				}
			}
		}
	}()

	return out
}

func (q *RedisQueue) Close() error {
	return q.client.Close()
}
