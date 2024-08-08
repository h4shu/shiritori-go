package intractors

import (
	"context"

	"github.com/h4shu/shiritori-go/application/inputs"
	"github.com/h4shu/shiritori-go/application/interfaces/apis"
	"github.com/h4shu/shiritori-go/application/interfaces/repositories"
	"github.com/h4shu/shiritori-go/application/outputs"
	"github.com/h4shu/shiritori-go/domain/entities"
)

type (
	WordchainUsecase struct {
		r     repositories.IWordchainRepository
		t     entities.WordType
		limit int
		a     apis.IValidateWordApi
	}
)

func NewWordchainUsecase(a apis.IValidateWordApi, r repositories.IWordchainRepository, t entities.WordType, limit int) *WordchainUsecase {
	return &WordchainUsecase{
		r:     r,
		t:     t,
		limit: limit,
		a:     a,
	}
}

func (u *WordchainUsecase) GetLast(ctx context.Context) (*outputs.WordchainGetLastOutputData, error) {
	w, err := u.r.GetLast(ctx)
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
	wc, err := u.r.List(ctx, i.GetLimit())
	if err != nil {
		return nil, err
	}
	o := outputs.NewWordchainListOutputData(wc)
	return o, nil
}

func (u *WordchainUsecase) Append(ctx context.Context, i *inputs.WordchainAppendInputData) error {
	wc, err := u.r.List(ctx, u.limit)
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
	v, err := u.a.Validate(w, u.t)
	if err != nil {
		return err
	} else if !v {
		return &entities.ErrWordInvalid{
			Word: w,
		}
	}
	err = u.r.Append(ctx, w)
	if err != nil {
		return err
	}
	return nil
}
