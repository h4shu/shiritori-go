package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/h4shu/shiritori-go/entity"
)

func TestNewWordchain(t *testing.T) {
	wt := entity.WordTypeNone
	got := entity.NewWordchain(wt)
	assert.NotNil(t, got, "got nil")
}

func TestAppend(t *testing.T) {
	w, err := entity.NewWord("abcdefg")
	assert.Nilf(t, err, "unexpected error: %v", err)
	hw, err := entity.NewHiraganaWord("あいうえお")
	assert.Nilf(t, err, "unexpected error: %v", err)

	wt := entity.WordTypeNone
	wc := entity.NewWordchain(wt)
	got, err := wc.Append(hw)
	e := &entity.ErrWordTypeInvalid{hw, wt}
	assert.EqualErrorf(t, err, e.Error(), "invalid error: %v", err)
	assert.Nilf(t, got, "got %v; want nil", got)

	got, err = wc.Append(&w)
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.NotNil(t, got, "got nil")
}

func TestLen(t *testing.T) {
	w, err := entity.NewWord("あいうえお")
	assert.Nilf(t, err, "unexpected error: %v", err)
	len := 10
	wt := entity.WordTypeNone
	wc := entity.NewWordchain(wt)
	for i := 0; i < len; i++ {
		wc, err = wc.Append(&w)
		assert.Nilf(t, err, "unexpected error: %v", err)
	}
	got := wc.Len()
	assert.Equalf(t, got, len, "got %d; want %d", got, len)
}

func TestToStrSlice(t *testing.T) {
	s := []string{"あいうえお", "かきくけこ", "さしすせそ"}
	wt := entity.WordTypeNone
	wc := entity.NewWordchain(wt)
	for _, v := range s {
		w, err := entity.NewWord(v)
		assert.Nilf(t, err, "unexpected error: %v", err)
		wc, err = wc.Append(&w)
		assert.Nilf(t, err, "unexpected error: %v", err)
	}
	got := wc.ToStrSlice()
	for i, v := range s {
		assert.Equalf(t, got[i], v, "got '%s'; want '%s'", got[i], v)
	}
}
