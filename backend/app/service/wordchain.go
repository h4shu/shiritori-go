package service

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/h4shu/shiritori-go/model"
)

const (
	WordchainKey   = "wordchain"
	WordchainLimit = 10
)

type WordchainService struct {
	rdb *redis.Client
}

func NewWordchainService(rdb *redis.Client) *WordchainService {
	return &WordchainService{
		rdb: rdb,
	}
}

func (s *WordchainService) GetLastWord(ctx context.Context) (*model.Word, error) {
	val, err := s.rdb.LRange(ctx, WordchainKey, 0, 0).Result()
	if err != nil {
		return nil, err
	}
	var w model.Word
	if len(val) == 0 {
		w = model.NewWord("しりとり")
	} else {
		w = model.NewWord(val[0])
	}
	return &w, nil
}

func (s *WordchainService) GetWordchain(ctx context.Context) (*model.Wordchain, error) {
	val, err := s.rdb.LRange(ctx, WordchainKey, 0, WordchainLimit-1).Result()
	if err != nil {
		return nil, err
	}
	var wc model.Wordchain
	for _, v := range val {
		wc.Append(model.NewWord(v))
	}
	return &wc, nil
}

func (s *WordchainService) pushWord(ctx context.Context, word *model.Word) error {
	_, err := s.rdb.LPush(ctx, WordchainKey, word).Result()
	if err != nil {
		return err
	}
	return nil
}

func (s *WordchainService) TryAddWord(ctx context.Context, word *model.Word) error {
	w, err := s.GetLastWord(ctx)
	if err != nil {
		return err
	} else if !word.ValidateChain(w) {
		return &model.ErrWordInvalid{
			Word: word,
		}
	}
	err = s.pushWord(ctx, word)
	if err != nil {
		return err
	}
	return nil
}
