package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/h4shu/shiritori-go/model"
)

func TestNewWord(t *testing.T) {
	word := []rune("")
	got := model.NewWord(string(word))
	assert.Emptyf(t, got, "got '%s'; want ''", got)

	word = []rune("あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん")
	got = model.NewWord(string(word))
	for i, v := range word {
		assert.Equalf(t, got[i], v, "got '%c'; want '%c'", got[i], v)
	}
}

func TestString(t *testing.T) {
	str := ""
	w := model.NewWord(str)
	got := w.String()
	assert.Equalf(t, got, str, "got '%s'; want '%s'", got, str)

	str = "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん"
	w = model.NewWord(str)
	got = w.String()
	assert.Equalf(t, got, str, "got '%s'; want '%s'", got, str)
}

func TestMarshalBinary(t *testing.T) {
	str := ""
	w := model.NewWord(str)
	got, err := w.MarshalBinary()
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.Equalf(t, string(got), str, "got %v; want %v", got, []byte(str))

	str = "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん"
	w = model.NewWord(str)
	got, err = w.MarshalBinary()
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.Equalf(t, string(got), str, "got %v; want %v", got, []byte(str))
}

func TestUnmarshalBinary(t *testing.T) {
	var got model.Word
	str := ""
	err := got.UnmarshalBinary([]byte(str))
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.Equalf(t, got.String(), str, "got %v; want %v", got, model.NewWord(str))

	str = "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん"
	err = got.UnmarshalBinary([]byte(str))
	assert.Nilf(t, err, "unexpected error: %v", err)
	assert.Equalf(t, got.String(), str, "got %v; want %v", got, model.NewWord(str))
}

func TestFirstChr(t *testing.T) {
	str := ""
	w := model.NewWord(str)
	got := w.FirstChr()
	assert.Equalf(t, got, rune(0), "got %v; want 0", got)

	str = "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん"
	w = model.NewWord(str)
	got = w.FirstChr()
	assert.Equalf(t, got, rune('あ'), "got %v; want 'あ'", got)
}

func TestLastChr(t *testing.T) {
	str := ""
	w := model.NewWord(str)
	got := w.LastChr()
	assert.Equalf(t, got, rune(0), "got %v; want 0", got)

	str = "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよわをん"
	w = model.NewWord(str)
	got = w.LastChr()
	assert.Equalf(t, got, rune('ん'), "got %v; want 'ん'", got)
}

func TestValidate(t *testing.T) {
	str := ""
	w := model.NewWord(str)
	got := w.Validate()
	assert.Falsef(t, got, "'%s' is invalid word", str)

	str = "あ"
	w = model.NewWord(str)
	got = w.Validate()
	assert.Falsef(t, got, "'%s' is invalid word", str)

	str = "らいおん"
	w = model.NewWord(str)
	got = w.Validate()
	assert.Falsef(t, got, "'%s' is invalid word", str)

	str = "とり"
	w = model.NewWord(str)
	got = w.Validate()
	assert.Truef(t, got, "'%s' is valid word", str)
}

func TestValidateChain(t *testing.T) {
	s1 := ""
	s2 := ""
	w1 := model.NewWord(s1)
	w2 := model.NewWord(s2)
	got := w2.ValidateChain(&w1)
	assert.Falsef(t, got, "'%s' -> '%s' is invalid chain", s1, s2)

	s1 = "しりとり"
	s2 = ""
	w1 = model.NewWord(s1)
	w2 = model.NewWord(s2)
	got = w2.ValidateChain(&w1)
	assert.Falsef(t, got, "'%s' -> '%s' is invalid chain", s1, s2)

	s1 = ""
	s2 = "りんご"
	w1 = model.NewWord(s1)
	w2 = model.NewWord(s2)
	got = w2.ValidateChain(&w1)
	assert.Falsef(t, got, "'%s' -> '%s' is invalid chain", s1, s2)

	s1 = "ごりら"
	s2 = "らいおん"
	w1 = model.NewWord(s1)
	w2 = model.NewWord(s2)
	got = w2.ValidateChain(&w1)
	assert.Falsef(t, got, "'%s' -> '%s' is invalid chain", s1, s2)

	s1 = "かもめ"
	s2 = "め"
	w1 = model.NewWord(s1)
	w2 = model.NewWord(s2)
	got = w2.ValidateChain(&w1)
	assert.Falsef(t, got, "'%s' -> '%s' is invalid chain", s1, s2)

	s1 = "しりとり"
	s2 = "りんご"
	w1 = model.NewWord(s1)
	w2 = model.NewWord(s2)
	got = w2.ValidateChain(&w1)
	assert.Truef(t, got, "'%s' -> '%s' is valid chain", s1, s2)
}
