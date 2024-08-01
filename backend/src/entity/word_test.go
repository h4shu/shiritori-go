package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/h4shu/shiritori-go/entity"
)

func TestNewWord(t *testing.T) {
	word := []rune("")
	got, err := entity.NewWord(string(word))
	assert.ErrorIsf(t, err, &entity.ErrWordEmpty{}, "unexpected error: %v", err)
	assert.Emptyf(t, got, "got '%s'; want ''", got)

	word = []rune("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん")
	got, err = entity.NewWord(string(word))
	assert.Nilf(t, err, "unexpected error: %v", err)
	for i, v := range word {
		assert.Equalf(t, got[i], v, "got '%c'; want '%c'", got[i], v)
	}
}

func TestWordString(t *testing.T) {
	str := "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん"
	w, err := entity.NewWord(str)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got := w.String()
	assert.Equalf(t, got, str, "got '%s'; want '%s'", got, str)
}

func TestWordMarshalBinary(t *testing.T) {
	str := "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん"
	w, err := entity.NewWord(str)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got, err := w.MarshalBinary()
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.Equalf(t, string(got), str, "got %v; want %v", got, []byte(str))
}

func TestWordUnmarshalBinary(t *testing.T) {
	var got entity.Word
	str := "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん"
	err := got.UnmarshalBinary([]byte(str))
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.Equalf(t, got.String(), str, "got %v; want %v", got, str)
}

func TestWordFirstChr(t *testing.T) {
	str := "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん"
	w, err := entity.NewWord(str)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got := w.FirstChr()
	assert.Equalf(t, got, rune('あ'), "got %v; want 'あ'", got)
}

func TestWordLastChr(t *testing.T) {
	str := "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん"
	w, err := entity.NewWord(str)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got := w.LastChr()
	assert.Equalf(t, got, rune('ん'), "got %v; want 'ん'", got)
}

func TestWordValidateChain(t *testing.T) {
	s1 := "alice"
	s2 := "bob"
	w1, err := entity.NewWord(s1)
	assert.Nilf(t, err, "unexpected error: %v", err)
	w2, err := entity.NewWord(s2)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got := w1.ValidateChain(&w2)
	assert.Falsef(t, got, "'%s' -> '%s' is invalid chain", s1, s2)

	s1 = "ab"
	s2 = "b"
	w1, err = entity.NewWord(s1)
	assert.Nilf(t, err, "unexpected error: %v", err)
	w2, err = entity.NewWord(s2)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got = w1.ValidateChain(&w2)
	assert.Falsef(t, got, "'%s' -> '%s' is invalid chain", s1, s2)

	s1 = "かもめ"
	s2 = "め"
	w1, err = entity.NewWord(s1)
	assert.Nilf(t, err, "unexpected error: %v", err)
	w2, err = entity.NewWord(s2)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got = w1.ValidateChain(&w2)
	assert.Falsef(t, got, "'%s' -> '%s' is invalid chain", s1, s2)

	s1 = "apple"
	s2 = "egg"
	w1, err = entity.NewWord(s1)
	assert.Nilf(t, err, "unexpected error: %v", err)
	w2, err = entity.NewWord(s2)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got = w1.ValidateChain(&w2)
	assert.Truef(t, got, "'%s' -> '%s' is valid chain", s1, s2)

	s1 = "しりとり"
	s2 = "りんご"
	w1, err = entity.NewWord(s1)
	assert.Nilf(t, err, "unexpected error: %v", err)
	w2, err = entity.NewWord(s2)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got = w1.ValidateChain(&w2)
	assert.Truef(t, got, "'%s' -> '%s' is valid chain", s1, s2)

	s1 = "ごりら"
	s2 = "らいおん"
	w1, err = entity.NewWord(s1)
	assert.Nilf(t, err, "unexpected error: %v", err)
	w2, err = entity.NewWord(s2)
	assert.Nilf(t, err, "unexpected error: %v", err)
	got = w1.ValidateChain(&w2)
	assert.Truef(t, got, "'%s' -> '%s' is valid chain", s1, s2)
}
