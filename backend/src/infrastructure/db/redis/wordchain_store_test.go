package redis_test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"

	"github.com/h4shu/shiritori-go/infrastructure/db/redis"
)

func TestWordchainStoreGetLast(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	key := "TestWordchainStoreGetLast"
	s := redis.NewWordchainStore(rdb, key)
	ctx := context.Background()

	mock.ExpectLRange(key, 0, 0).SetVal([]string{})
	got, err := s.GetLast(ctx)
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.Emptyf(t, got, "got '%s'; want ''", got)

	mock.ClearExpect()
	mock.ExpectLRange(key, 0, 0).SetVal([]string{"りんご"})
	got, err = s.GetLast(ctx)
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.Equalf(t, got, "りんご", "got '%s'; want 'りんご'", got)

	mock.ClearExpect()
	mock.ExpectLRange(key, 0, 0).SetErr(errors.New("FAIL"))
	got, err = s.GetLast(ctx)
	assert.EqualErrorf(t, err, "FAIL", "unexpected error: %v", err)
	assert.Emptyf(t, got, "got '%s'; want ''", got)
}

func TestWordchainStoreList(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	key := "TestWordchainStoreList"
	s := redis.NewWordchainStore(rdb, key)
	ctx := context.Background()

	limit := 100
	val := []string{}
	mock.ExpectLRange(key, 0, int64(limit)-1).SetVal(val)
	got, err := s.List(ctx, limit)
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.Emptyf(t, got, "got %v; want []", got)

	mock.ClearExpect()
	word := []rune("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん")
	for i := 0; i < limit; i++ {
		for j := range word {
			// Hiragana unicode range: U+3041('ぁ')-U+3094('ゔ')
			word[j] = rune('ぁ' + ((i + j) % ('ゔ' - 'ぁ' + 1)))
		}
		val = append(val, string(word))
	}
	mock.ExpectLRange(key, 0, int64(limit)-1).SetVal(val)
	got, err = s.List(ctx, limit)
	assert.Nilf(t, err, "unexpected error: %v", err)
	len := len(got)
	assert.Equalf(t, len, limit, "got %d; want %d", len, limit)
	for i, v := range got {
		assert.Equalf(t, v, val[i], "got '%s'; want '%s'", v, val[i])
	}

	mock.ClearExpect()
	mock.ExpectLRange(key, 0, int64(limit)-1).SetErr(errors.New("FAIL"))
	got, err = s.List(ctx, limit)
	assert.EqualErrorf(t, err, "FAIL", "unexpected error: %v", err)
	assert.Nilf(t, got, "got '%v'; want nil", got)
}

func TestWordchainStoreAppend(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	key := "TestWordchainStoreAppend"
	s := redis.NewWordchainStore(rdb, key)
	ctx := context.Background()

	str := ""
	mock.Regexp().ExpectLPush(key, str).SetVal(0)
	mock.Regexp().ExpectLPush(key, str).SetErr(nil)
	err := s.Append(ctx, str)
	assert.Nilf(t, err, "unexpected error: %v", err)

	mock.ClearExpect()
	str = "りんご"
	mock.Regexp().ExpectLPush(key, str).SetErr(errors.New("FAIL"))
	err = s.Append(ctx, str)
	assert.EqualErrorf(t, err, "FAIL", "unexpected error: %v", err)
}
