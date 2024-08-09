package repositories

import (
	"context"

	"github.com/h4shu/shiritori-go/domain/entities"
)

type (
	IWorddictRepository interface {
		Exist(ctx context.Context, word entities.IWord) error
	}
)
