package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/redis/go-redis/v9"
)

type RedisQueue struct {
	client *redis.Client
	stream string
}

func NewRedisQueue(stream string) (*RedisQueue, error) {
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
	return &RedisQueue{client: client, stream: stream}, nil
}

func (q *RedisQueue) Push(ctx context.Context, item domain.Message) error {
	payload, err := json.Marshal(item)
	if err != nil {
		return err
	}

	args := &redis.XAddArgs{
		Stream: q.stream,
		Values: map[string]any{
			"message": payload,
		},
	}
	if _, err := q.client.XAdd(ctx, args).Result(); err != nil {
		return fmt.Errorf("failed to push message to redis: %w", err)
	}
	return nil
}

func (q *RedisQueue) Close() error {
	return q.client.Close()
}
