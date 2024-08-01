package entities_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/h4shu/shiritori-go/domain/entities"
)

func TestNewHiraganaWord(t *testing.T) {
	word := "あいうえおabcdefg"
	got, err := entities.NewHiraganaWord(word)
	e := &entities.ErrNotHiragana{word}
	assert.EqualErrorf(t, err, e.Error(), "invalid error: %v", err)
	assert.Emptyf(t, got, "got '%s'; want ''", got)

	word = "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん"
	got, err = entities.NewHiraganaWord(word)
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.Equalf(t, got.String(), word, "got '%s'; want '%s'", got, word)
}

func TestHiraganaWordString(t *testing.T) {
	str := "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん"
	w, err := entities.NewHiraganaWord(str)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got := w.String()
	assert.Equalf(t, got, str, "got '%s'; want '%s'", got, str)
}

func TestHiraganaWordMarshalBinary(t *testing.T) {
	str := "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん"
	w, err := entities.NewHiraganaWord(str)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got, err := w.MarshalBinary()
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.Equalf(t, string(got), str, "got %v; want %v", got, []byte(str))
}

func TestHiraganaWordUnmarshalBinary(t *testing.T) {
	var got entities.Word
	str := "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん"
	err := got.UnmarshalBinary([]byte(str))
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.Equalf(t, got.String(), str, "got %v; want %v", got, str)
}

func TestHiraganaWordFirstChr(t *testing.T) {
	str := "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん"
	w, err := entities.NewHiraganaWord(str)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got := w.FirstChr()
	assert.Equalf(t, got, rune('あ'), "got %v; want 'あ'", got)
}

func TestHiraganaWordLastChr(t *testing.T) {
	str := "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん"
	w, err := entities.NewHiraganaWord(str)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got := w.LastChr()
	assert.Equalf(t, got, rune('ん'), "got %v; want 'ん'", got)
}

func TestHiraganaWordValidateChain(t *testing.T) {
	s1 := "ごりら"
	s2 := "らいおん"
	w1, err := entities.NewHiraganaWord(s1)
	assert.Nilf(t, err, "unexpected error: %v", err)
	w2, err := entities.NewHiraganaWord(s2)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got := w1.ValidateChain(w2)
	assert.Falsef(t, got, "'%s' -> '%s' is invalid chain", s1, s2)

	s1 = "かもめ"
	s2 = "め"
	w1, err = entities.NewHiraganaWord(s1)
	assert.Nilf(t, err, "unexpected error: %v", err)
	w2, err = entities.NewHiraganaWord(s2)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got = w1.ValidateChain(w2)
	assert.Falsef(t, got, "'%s' -> '%s' is invalid chain", s1, s2)

	s1 = "しりとり"
	s2 = "りんご"
	w1, err = entities.NewHiraganaWord(s1)
	assert.Nilf(t, err, "unexpected error: %v", err)
	w2, err = entities.NewHiraganaWord(s2)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got = w1.ValidateChain(w2)
	assert.Truef(t, got, "'%s' -> '%s' is valid chain", s1, s2)
}
