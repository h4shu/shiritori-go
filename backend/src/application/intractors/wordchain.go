package intractors

import (
	"context"

	"github.com/h4shu/shiritori-go/application/inputs"
	"github.com/h4shu/shiritori-go/application/outputs"
	"github.com/h4shu/shiritori-go/application/repositories"
	"github.com/h4shu/shiritori-go/domain/entities"
)

type (
	WordchainUsecase struct {
		wcr   repositories.IWordchainRepository
		wdr   repositories.IWorddictRepository
		t     entities.WordType
		limit int
	}
)

func NewWordchainUsecase(
	wcr repositories.IWordchainRepository,
	wdr repositories.IWorddictRepository,
	t entities.WordType,
	limit int,
) *WordchainUsecase {
	return &WordchainUsecase{
		wcr:   wcr,
		wdr:   wdr,
		t:     t,
		limit: limit,
	}
}

func (u *WordchainUsecase) GetLast(ctx context.Context) (*outputs.WordchainGetLastOutputData, error) {
	w, err := u.wcr.GetLast(ctx)
	if err != nil {
		return nil, err
	} else if w == nil {
		w, err = entities.GetFirstWordForType(u.t)
		if err != nil {
			return nil, err
		}
	}
	o := outputs.NewWordchainGetLastOutputData(w)
	return o, nil
}

func (u *WordchainUsecase) List(ctx context.Context, i *inputs.WordchainListInputData) (*outputs.WordchainListOutputData, error) {
	wc, err := u.wcr.List(ctx, i.GetLimit())
	if err != nil {
		return nil, err
	}
	o := outputs.NewWordchainListOutputData(wc)
	return o, nil
}

func (u *WordchainUsecase) Append(ctx context.Context, i *inputs.WordchainAppendInputData) error {
	wc, err := u.wcr.List(ctx, u.limit)
	if err != nil {
		return err
	} else if wc.Len() == 0 {
		fw, err := entities.GetFirstWordForType(u.t)
		if err != nil {
			return err
		}
		wc, err = wc.Append(fw)
		if err != nil {
			return err
		}
	}
	w := i.GetWord()
	if _, err = wc.Append(w); err != nil {
		return err
	}
	if err = u.wdr.Exist(ctx, w); err != nil {
		return err
	}
	err = u.wcr.Append(ctx, w)
	if err != nil {
		return err
	}
	return nil
}
