package entities_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/h4shu/shiritori-go/domain/entities"
)

func TestNewWordchain(t *testing.T) {
	wt := entities.WordTypeHiragana
	got := entities.NewWordchain(wt)
	assert.NotNil(t, got, "got nil")
	assert.Zerof(t, got.Len(), "got any word: %v", got)
}

func TestAppend(t *testing.T) {
	w, err := entities.NewWord("abcdefg")
	assert.Nilf(t, err, "unexpected error: %v", err)
	wt := entities.WordTypeHiragana
	wc := entities.NewWordchain(wt)
	got, err := wc.Append(&w)
	et := &entities.ErrWordTypeInvalid{&w, wt}
	assert.EqualErrorf(t, err, et.Error(), "invalid error: %v", err)
	assert.Nilf(t, got, "got %v; want nil", got)

	hw, err := entities.NewHiraganaWord("りんご")
	assert.Nilf(t, err, "unexpected error: %v", err)
	got, err = wc.Append(hw)
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.NotNil(t, got, "got nil")
}

func TestAppendStr(t *testing.T) {
	wt := entities.WordTypeHiragana
	wc := entities.NewWordchain(wt)
	got, err := wc.AppendStr("")
	assert.ErrorIsf(t, err, &entities.ErrWordEmpty{}, "unexpected error: %v", err)
	assert.Nilf(t, got, "got %v; want nil", got)

	got, err = wc.AppendStr("りんご")
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.NotNil(t, got, "got nil")
}

func TestLen(t *testing.T) {
	wt := entities.WordTypeHiragana
	wc := entities.NewWordchain(wt)
	str := []string{"りんご", "ごりら", "らくだ"}
	var err error
	for _, s := range str {
		wc, err = wc.AppendStr(s)
		assert.Nilf(t, err, "unexpected error: %v", err)
	}
	got := wc.Len()
	want := len(str)
	assert.Equalf(t, got, want, "got %d; want %d", got, want)
}

func TestToStrSlice(t *testing.T) {
	s := []string{"りんご", "ごりら", "らくだ"}
	wt := entities.WordTypeHiragana
	wc := entities.NewWordchain(wt)
	for _, v := range s {
		w, err := entities.NewHiraganaWord(v)
		assert.Nilf(t, err, "unexpected error: %v", err)
		wc, err = wc.Append(w)
		assert.Nilf(t, err, "unexpected error: %v", err)
	}
	got := wc.ToStrSlice()
	assert.NotNil(t, got, "got nil")
	for i, v := range s {
		assert.Equalf(t, got[i], v, "got '%s'; want '%s'", got[i], v)
	}
}
