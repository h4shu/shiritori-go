package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type WordchainStore struct {
	c   *redis.Client
	key string
}

func NewWordchainStore(c *redis.Client, key string) *WordchainStore {
	return &WordchainStore{
		c:   c,
		key: key,
	}
}

func (s *WordchainStore) GetLast(ctx context.Context) (string, error) {
	val, err := s.c.LRange(ctx, s.key, 0, 0).Result()
	if err != nil {
		return "", err
	} else if len(val) == 0 {
		return "", nil
	}
	return val[0], nil
}

func (s *WordchainStore) List(ctx context.Context, limit int) ([]string, error) {
	val, err := s.c.LRange(ctx, s.key, 0, int64(limit)-1).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (s *WordchainStore) Append(ctx context.Context, w string) error {
	_, err := s.c.LPush(ctx, s.key, w).Result()
	if err != nil {
		return err
	}
	return nil
}
