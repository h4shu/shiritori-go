package service

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/h4shu/shiritori-go/entity"
)

const (
	WordchainKey   = "wordchain"
	WordchainLimit = 1000
)

type WordchainService struct {
	rdb *redis.Client
}

func NewWordchainService(rdb *redis.Client) *WordchainService {
	return &WordchainService{
		rdb: rdb,
	}
}

func (s *WordchainService) GetLastWord(ctx context.Context) (*entity.Word, error) {
	val, err := s.rdb.LRange(ctx, WordchainKey, 0, 0).Result()
	if err != nil {
		return nil, err
	}
	var w entity.Word
	if len(val) == 0 {
		w = entity.NewWord("しりとり")
	} else {
		w = entity.NewWord(val[0])
	}
	return &w, nil
}

func (s *WordchainService) GetWordchain(ctx context.Context) (*entity.Wordchain, error) {
	val, err := s.rdb.LRange(ctx, WordchainKey, 0, WordchainLimit-1).Result()
	if err != nil {
		return nil, err
	}
	var wc entity.Wordchain
	for _, v := range val {
		wc = *wc.Append(entity.NewWord(v))
	}
	return &wc, nil
}

func (s *WordchainService) pushWord(ctx context.Context, word *entity.Word) error {
	_, err := s.rdb.LPush(ctx, WordchainKey, word).Result()
	if err != nil {
		return err
	}
	return nil
}

func (s *WordchainService) TryAddWord(ctx context.Context, word *entity.Word) error {
	w, err := s.GetLastWord(ctx)
	if err != nil {
		return err
	} else if !word.ValidateChain(w) {
		return &entity.ErrWordInvalid{
			Word: word,
		}
	}
	err = s.pushWord(ctx, word)
	if err != nil {
		return err
	}
	return nil
}
