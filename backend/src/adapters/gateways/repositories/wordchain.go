package repositories

import (
	"context"
	"slices"

	"github.com/h4shu/shiritori-go/domain/entities"
)

type (
	IWordchainStore interface {
		GetLast(ctx context.Context) (string, error)
		List(ctx context.Context, limit int) ([]string, error)
		Append(ctx context.Context, w string) error
	}

	WordchainRepository struct {
		s IWordchainStore
		t entities.WordType
	}
)

func NewWordchainRepository(s IWordchainStore, t entities.WordType) *WordchainRepository {
	return &WordchainRepository{
		s: s,
		t: t,
	}
}

func (r *WordchainRepository) GetLast(ctx context.Context) (entities.IWord, error) {
	w, err := r.s.GetLast(ctx)
	if err != nil {
		return nil, err
	} else if w == "" {
		return entities.GetFirstWordForType(r.t)
	}
	return entities.NewWordWithType(w, r.t)
}

func (r *WordchainRepository) List(ctx context.Context, limit int) (*entities.Wordchain, error) {
	words, err := r.s.List(ctx, limit)
	if err != nil {
		return nil, err
	}
	slices.Reverse(words)
	wc := entities.NewWordchain(r.t)
	for _, w := range words {
		wc, err = wc.AppendStr(w)
		if err != nil {
			return nil, err
		}
	}
	return wc, nil
}

func (r *WordchainRepository) Append(ctx context.Context, w entities.IWord) error {
	return r.s.Append(ctx, w.String())
}
