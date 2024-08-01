package entities

import (
	"fmt"
)

type (
	Word  []rune
	IWord interface {
		String() string
		MarshalBinary() ([]byte, error)
		UnmarshalBinary(data []byte) error
		FirstChr() rune
		LastChr() rune
		ValidateChain(next IWord) bool
	}
	ErrWordEmpty   struct{}
	ErrWordInvalid struct {
		Word IWord
	}
)

func NewWord(word string) (Word, error) {
	if len(word) == 0 {
		return nil, &ErrWordEmpty{}
	}
	return []rune(word), nil
}

func (w *Word) String() string {
	return string(*w)
}

func (w *Word) MarshalBinary() ([]byte, error) {
	return []byte(w.String()), nil
}

func (w *Word) UnmarshalBinary(data []byte) error {
	new, err := NewWord(string(data))
	*w = new
	return err
}

func (w *Word) FirstChr() rune {
	return (*w)[0]
}

func (w *Word) LastChr() rune {
	return (*w)[len(*w)-1]
}

func (w *Word) ValidateChain(next IWord) bool {
	nw, ok := next.(*Word)
	if !ok || len(*nw) < 2 || w.LastChr() != nw.FirstChr() {
		return false
	}
	return true
}

func (e *ErrWordEmpty) Error() string {
	return "文字列が空です"
}

func (e *ErrWordInvalid) Error() string {
	return fmt.Sprintf("'%s' は、正しくありません", e.Word)
}
