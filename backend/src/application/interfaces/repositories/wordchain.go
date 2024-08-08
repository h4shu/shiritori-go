package repositories

import (
	"context"

	"github.com/h4shu/shiritori-go/domain/entities"
)

type IWordchainRepository interface {
	GetLast(ctx context.Context) (entities.IWord, error)
	List(ctx context.Context, limit int) (*entities.Wordchain, error)
	Append(ctx context.Context, w entities.IWord) error
}
