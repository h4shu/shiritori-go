package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/h4shu/shiritori-go/entity"
)

func TestAppend(t *testing.T) {
	var wc entity.Wordchain
	w := entity.NewWord("あいうえお")
	wc = *wc.Append(w)
	got := wc[0]
	assert.Equalf(t, got.String(), w.String(), "got '%s'; want '%s'", got, w)
}

func TestLen(t *testing.T) {
	len := 10
	var wc entity.Wordchain
	for i := 0; i < len; i++ {
		wc = *wc.Append(entity.NewWord("あいうえお"))
	}
	got := wc.Len()
	assert.Equalf(t, got, len, "got %d; want %d", got, len)
}

func TestToStrSlice(t *testing.T) {
	s := []string{"あいうえお", "かきくけこ", "さしすせそ"}
	var wc entity.Wordchain
	for _, v := range s {
		wc = *wc.Append(entity.NewWord(v))
	}
	got := wc.ToStrSlice()
	for i, v := range s {
		assert.Equalf(t, got[i], v, "got '%s'; want '%s'", got[i], v)
	}
}
