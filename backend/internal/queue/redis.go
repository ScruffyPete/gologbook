package queue

import (
	"context"
	"fmt"
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

func (q *RedisQueue) SubscribeForDocumentTokens(ctx context.Context, channelName string) <-chan string {
	pubsub := q.client.Subscribe(ctx, channelName)

	out := make(chan string, 100)

	go func() {
		defer close(out)
		defer pubsub.Close()

		for msg := range pubsub.Channel() {
			if msg.Payload == "[[STOP]]" {
				return
			}
			out <- msg.Payload
		}
	}()

	return out
}

func (q *RedisQueue) Close() error {
	return q.client.Close()
}
