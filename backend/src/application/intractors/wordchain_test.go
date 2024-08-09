package intractors_test

import (
	"context"
	"testing"

	"github.com/h4shu/shiritori-go/application/inputs"
	"github.com/h4shu/shiritori-go/application/intractors"
	"github.com/h4shu/shiritori-go/domain/entities"
	"github.com/stretchr/testify/assert"
)

type (
	WordchainRepositoryMock struct {
		GetLastVal entities.IWord
		GetLastErr error
		ListVal    *entities.Wordchain
		ListErr    error
		AppendErr  error
	}

	WorddictRepositoryMock struct {
		ExistVal bool
		ExitErr  error
	}
)

func NewWordchainRepositoryMock(
	GetLastVal entities.IWord,
	GetLastErr error,
	ListVal *entities.Wordchain,
	ListErr error,
	AppendErr error,
) *WordchainRepositoryMock {
	return &WordchainRepositoryMock{
		GetLastVal,
		GetLastErr,
		ListVal,
		ListErr,
		AppendErr,
	}
}

func NewWorddictRepositoryMock(
	ExitErr error,
) *WorddictRepositoryMock {
	return &WorddictRepositoryMock{
		ExitErr: ExitErr,
	}
}

func (m *WordchainRepositoryMock) GetLast(ctx context.Context) (entities.IWord, error) {
	return m.GetLastVal, m.GetLastErr
}
func (m *WordchainRepositoryMock) List(ctx context.Context, limit int) (*entities.Wordchain, error) {
	return m.ListVal, m.ListErr
}
func (m *WordchainRepositoryMock) Append(ctx context.Context, w entities.IWord) error {
	return m.AppendErr
}

func (m *WorddictRepositoryMock) Exist(ctx context.Context, word entities.IWord) error {
	return m.ExitErr
}

func TestAppend(t *testing.T) {
	ctx := context.Background()
	ty := entities.WordTypeHiragana
	wc := entities.NewWordchain(ty)
	lw, err := entities.NewHiraganaWord("しりとり")
	assert.Nilf(t, err, "unexpected error: %v", err)
	wc.Append(lw)
	mwcr := NewWordchainRepositoryMock(nil, nil, wc, nil, nil)
	mwdr := NewWorddictRepositoryMock(nil)
	u := intractors.NewWordchainUsecase(mwcr, mwdr, entities.WordTypeHiragana, 100)
	w, err := entities.NewHiraganaWord("りんご")
	assert.Nilf(t, err, "unexpected error: %v", err)
	i := inputs.NewWordchainAppendInputData(w)
	err = u.Append(ctx, i)
	assert.Nilf(t, err, "unexpected error: %v", err)

	w, err = entities.NewHiraganaWord("らいおん")
	assert.Nilf(t, err, "unexpected error: %v", err)
	i = inputs.NewWordchainAppendInputData(w)
	err = u.Append(ctx, i)
	assert.Error(t, err, "need error")
}
