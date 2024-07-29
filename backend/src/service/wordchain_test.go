package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"

	"github.com/h4shu/shiritori-go/entity"
	"github.com/h4shu/shiritori-go/service"
)

func TestGetLastWord(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	svc := service.NewWordchainService(rdb)
	ctx := context.Background()

	mock.ExpectLRange(service.WordchainKey, 0, 0).SetVal([]string{})
	got, err := svc.GetLastWord(ctx)
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.Equalf(t, got.String(), "しりとり", "got '%s'; want 'しりとり'", got)

	mock.ClearExpect()
	mock.ExpectLRange(service.WordchainKey, 0, 0).SetVal([]string{"りんご"})
	got, err = svc.GetLastWord(ctx)
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.Equalf(t, got.String(), "りんご", "got '%s'; want 'りんご'", got)

	mock.ClearExpect()
	mock.ExpectLRange(service.WordchainKey, 0, 0).SetErr(errors.New("FAIL"))
	got, err = svc.GetLastWord(ctx)
	assert.EqualErrorf(t, err, "FAIL", "unexpected error: %v", err)
	assert.Nilf(t, got, "got '%s'; want nil", got)
}

func TestGetWordchain(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	svc := service.NewWordchainService(rdb)
	ctx := context.Background()

	val := []string{}
	mock.ExpectLRange(service.WordchainKey, 0, service.WordchainLimit-1).SetVal(val)
	got, err := svc.GetWordchain(ctx)
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.Equalf(t, got.Len(), 0, "got %d; want 0", got.Len())

	mock.ClearExpect()
	word := []rune("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん")
	for i := 0; i < service.WordchainLimit; i++ {
		for j := range word {
			// Hiragana unicode range: U+3041('ぁ')-U+3094('ゔ')
			word[j] = rune('ぁ' + ((i + j) % ('ゔ' - 'ぁ' + 1)))
		}
		val = append(val, string(word))
	}
	mock.ExpectLRange(service.WordchainKey, 0, service.WordchainLimit-1).SetVal(val)
	got, err = svc.GetWordchain(ctx)
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.Equalf(t, got.Len(), service.WordchainLimit, "got %d; want %d", got.Len(), service.WordchainLimit)
	for i, v := range got.ToStrSlice() {
		assert.Equalf(t, v, val[i], "got '%s'; want '%s'", v, val[i])
	}

	mock.ClearExpect()
	mock.ExpectLRange(service.WordchainKey, 0, service.WordchainLimit-1).SetErr(errors.New("FAIL"))
	got, err = svc.GetWordchain(ctx)
	assert.EqualErrorf(t, err, "FAIL", "unexpected error: %v", err)
	assert.Nilf(t, got, "got '%v'; want nil", got)
}

func TestTryAddWord(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	svc := service.NewWordchainService(rdb)
	ctx := context.Background()

	str := "あ"
	w := entity.NewWord(str)
	mock.ExpectLRange(service.WordchainKey, 0, 0).SetVal([]string{})
	err := svc.TryAddWord(ctx, &w)
	var wantErr *entity.ErrWordInvalid
	if assert.Errorf(t, err, "return error when '%s'", w) {
		assert.ErrorAsf(t, err, &wantErr, "got wrong error: %v", err)
	}

	mock.ClearExpect()
	str = "りんご"
	w = entity.NewWord(str)
	mock.ExpectLRange(service.WordchainKey, 0, 0).SetVal([]string{})
	mock.Regexp().ExpectLPush(service.WordchainKey, str).SetErr(errors.New("FAIL"))
	err = svc.TryAddWord(ctx, &w)
	assert.EqualErrorf(t, err, "FAIL", "unexpected error: %v", err)
}
