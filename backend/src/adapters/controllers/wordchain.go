package controllers

import (
	"context"

	"github.com/h4shu/shiritori-go/adapters/models"
	"github.com/h4shu/shiritori-go/adapters/presenters"
	"github.com/h4shu/shiritori-go/application/inputs"
	"github.com/h4shu/shiritori-go/application/usecases"
	"github.com/h4shu/shiritori-go/domain/entities"
)

type (
	WordchainController struct {
		u usecases.IWordchainUsecase
		p *presenters.WordchainPresenter
		t entities.WordType
	}
)

func NewWordchainController(u usecases.IWordchainUsecase, p *presenters.WordchainPresenter, t entities.WordType) *WordchainController {
	return &WordchainController{
		u: u,
		p: p,
		t: t,
	}
}

func (c *WordchainController) GetLast(ctx context.Context) (*models.WordchainGetLastModel, error) {
	o, err := c.u.GetLast(ctx)
	if err != nil {
		return nil, err
	}
	m := c.p.GetLast(ctx, o)
	return m, nil
}

func (c *WordchainController) List(ctx context.Context, limit int) (*models.WordchainListModel, error) {
	i := inputs.NewWordchainListInputData(limit)
	o, err := c.u.List(ctx, i)
	if err != nil {
		return nil, err
	}
	m := c.p.List(ctx, o)
	return m, nil
}

func (c *WordchainController) Append(ctx context.Context, word string) error {
	w, err := entities.NewWordWithType(word, c.t)
	if err != nil {
		return err
	}
	i := inputs.NewWordchainAppendInputData(w)
	err = c.u.Append(ctx, i)
	if err != nil {
		return err
	}
	return nil
}
