package entity

import (
	"fmt"
	"unicode"
)

type (
	Word           []rune
	ErrWordInvalid struct {
		Word *Word
	}
)

func NewWord(word string) Word {
	return []rune(word)
}

func (w *Word) String() string {
	return string(*w)
}

func (w *Word) MarshalBinary() ([]byte, error) {
	return []byte(w.String()), nil
}

func (w *Word) UnmarshalBinary(data []byte) error {
	*w = NewWord(string(data))
	return nil
}

func (w *Word) FirstChr() rune {
	if len(*w) == 0 {
		return 0
	}
	return (*w)[0]
}

func (w *Word) LastChr() rune {
	if len(*w) == 0 {
		return 0
	}
	return (*w)[len(*w)-1]
}

func (w *Word) Validate() bool {
	for _, r := range *w {
		if !unicode.In(r, unicode.Hiragana) {
			return false
		}
	}
	if len(*w) < 2 || w.LastChr() == 'ん' {
		return false
	}
	return true
}

func (w *Word) ValidateChain(pre *Word) bool {
	if !w.Validate() || w.FirstChr() != pre.LastChr() {
		return false
	}
	return true
}

func (e *ErrWordInvalid) Error() string {
	return fmt.Sprintf("'%s' は正しくありません", e.Word)
}
