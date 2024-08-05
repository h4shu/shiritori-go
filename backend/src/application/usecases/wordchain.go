package usecases

import (
	"context"

	"github.com/h4shu/shiritori-go/application/inputs"
	"github.com/h4shu/shiritori-go/application/outputs"
)

type IWordchainUsecase interface {
	GetLast(ctx context.Context) (*outputs.WordchainGetLastOutputData, error)
	List(ctx context.Context, i *inputs.WordchainListInputData) (*outputs.WordchainListOutputData, error)
	Append(ctx context.Context, i *inputs.WordchainAppendInputData) error
}
