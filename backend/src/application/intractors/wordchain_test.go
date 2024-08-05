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

func (m *WordchainRepositoryMock) GetLast(ctx context.Context) (entities.IWord, error) {
	return m.GetLastVal, m.GetLastErr
}
func (m *WordchainRepositoryMock) List(ctx context.Context, limit int) (*entities.Wordchain, error) {
	return m.ListVal, m.ListErr
}
func (m *WordchainRepositoryMock) Append(ctx context.Context, w entities.IWord) error {
	return m.AppendErr
}

func TestAppend(t *testing.T) {
	ctx := context.Background()
	lw, err := entities.NewHiraganaWord("しりとり")
	assert.Nilf(t, err, "unexpected error: %v", err)
	m := NewWordchainRepositoryMock(lw, nil, nil, nil, nil)
	u := intractors.NewWordchainUsecase(m, entities.WordTypeHiragana)
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
