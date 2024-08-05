package presenters

import (
	"context"

	"github.com/h4shu/shiritori-go/adapters/models"
	"github.com/h4shu/shiritori-go/application/outputs"
)

type (
	WordchainPresenter struct{}
)

func NewWordchainPresenter() *WordchainPresenter {
	return &WordchainPresenter{}
}

func (c *WordchainPresenter) GetLast(ctx context.Context, o *outputs.WordchainGetLastOutputData) *models.WordchainGetLastModel {
	w := o.GetWord()
	m := models.NewWordchainGetLastModel(w.String())
	return m
}

func (c *WordchainPresenter) List(ctx context.Context, o *outputs.WordchainListOutputData) *models.WordchainListModel {
	wc := o.GetWordchain()
	m := models.NewWordchainListModel(
		wc.ToStrSlice(),
		wc.Len(),
	)
	return m
}
