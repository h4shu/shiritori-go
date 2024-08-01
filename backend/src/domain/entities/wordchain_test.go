package entities_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/h4shu/shiritori-go/domain/entities"
)

func TestNewWordchain(t *testing.T) {
	wt := entities.WordTypeNone
	got, err := entities.NewWordchain(wt)
	assert.Nilf(t, got, "got %v; want nil", got)
	e := &entities.ErrNoFirstWord{wt}
	assert.EqualErrorf(t, err, e.Error(), "invalid error: %v", err)

	wt = entities.WordTypeHiragana
	got, err = entities.NewWordchain(wt)
	assert.NotNil(t, got, "got nil")
	assert.NotZero(t, got.Len(), "got no word")
	assert.Nilf(t, err, "unexpected error: %v", err)

}

func TestAppend(t *testing.T) {
	w, err := entities.NewWord("abcdefg")
	assert.Nilf(t, err, "unexpected error: %v", err)
	wt := entities.WordTypeHiragana
	wc, err := entities.NewWordchain(wt)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got, err := wc.Append(&w)
	et := &entities.ErrWordTypeInvalid{&w, wt}
	assert.EqualErrorf(t, err, et.Error(), "invalid error: %v", err)
	assert.Nilf(t, got, "got %v; want nil", got)

	hw, err := entities.NewHiraganaWord("あいうえお")
	assert.Nilf(t, err, "unexpected error: %v", err)
	got, err = wc.Append(hw)
	ew := &entities.ErrWordchainInvalid{wc, hw}
	assert.EqualErrorf(t, err, ew.Error(), "invalid error: %v", err)
	assert.Nilf(t, got, "got %v; want nil", got)

	hw, err = entities.NewHiraganaWord("りんご")
	assert.Nilf(t, err, "unexpected error: %v", err)
	got, err = wc.Append(hw)
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.NotNil(t, got, "got nil")
}

func TestAppendStr(t *testing.T) {
	wt := entities.WordTypeHiragana
	wc, err := entities.NewWordchain(wt)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got, err := wc.AppendStr("")
	assert.ErrorIsf(t, err, &entities.ErrWordEmpty{}, "unexpected error: %v", err)
	assert.Nilf(t, got, "got %v; want nil", got)

	hw, err := entities.NewHiraganaWord("あいうえお")
	assert.Nilf(t, err, "unexpected error: %v", err)
	ew := &entities.ErrWordchainInvalid{wc, hw}
	got, err = wc.AppendStr("あいうえお")
	assert.EqualErrorf(t, err, ew.Error(), "invalid error: %v", err)
	assert.Nilf(t, got, "got %v; want nil", got)

	got, err = wc.AppendStr("りんご")
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.NotNil(t, got, "got nil")
}

func TestLen(t *testing.T) {
	wt := entities.WordTypeHiragana
	wc, err := entities.NewWordchain(wt)
	assert.Nilf(t, err, "unexpected error: %v", err)
	str := []string{"りんご", "ごりら", "らくだ"}
	for _, s := range str {
		wc, err = wc.AppendStr(s)
		assert.Nilf(t, err, "unexpected error: %v", err)
	}
	got := wc.Len()
	want := 4
	assert.Equalf(t, got, want, "got %d; want %d", got, want)
}

func TestToStrSlice(t *testing.T) {
	s := []string{"りんご", "ごりら", "らくだ"}
	wt := entities.WordTypeHiragana
	wc, err := entities.NewWordchain(wt)
	assert.Nilf(t, err, "unexpected error: %v", err)
	for _, v := range s {
		w, err := entities.NewHiraganaWord(v)
		assert.Nilf(t, err, "unexpected error: %v", err)
		wc, err = wc.Append(w)
		assert.Nilf(t, err, "unexpected error: %v", err)
	}
	got := wc.ToStrSlice()
	assert.NotNil(t, got, "got nil")
	for i, v := range s {
		assert.Equalf(t, got[i+1], v, "got '%s'; want '%s'", got[i+1], v)
	}
}
