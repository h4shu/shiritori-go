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
		ValidateChain(next IWord) error
	}
	ErrWordEmpty struct{}
	ErrWordShort struct {
		Word IWord
	}
	ErrWordInvalid struct {
		Word IWord
	}
	ErrWordchainInvalid struct {
		Prev IWord
		Next IWord
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

func (w *Word) ValidateChain(next IWord) error {
	nw, ok := next.(*Word)
	if !ok {
		return &ErrWordInvalid{
			Word: nw,
		}
	}
	if len(*nw) < 2 {
		return &ErrWordShort{
			Word: nw,
		}
	}
	if w.LastChr() != nw.FirstChr() {
		return &ErrWordchainInvalid{
			Prev: w,
			Next: nw,
		}
	}
	return nil
}

func (e *ErrWordEmpty) Error() string {
	return "文字列が空です"
}

func (e *ErrWordShort) Error() string {
	return fmt.Sprintf("'%s' は、文字数が足りません", e.Word)
}

func (e *ErrWordInvalid) Error() string {
	return fmt.Sprintf("'%s' は、正しくありません", e.Word)
}

func (e *ErrWordchainInvalid) Error() string {
	return fmt.Sprintf("しりとりが正しくありません: %s => %s", e.Prev, e.Next)
}
