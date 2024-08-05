package entities

import (
	"fmt"
	"unicode"
)

type (
	HiraganaWord struct {
		w Word
	}
	ErrNotHiragana struct {
		Word string
	}
	ErrHiraganaWordInvalid struct {
		Word *HiraganaWord
	}
)

func NewHiraganaWord(word string) (*HiraganaWord, error) {
	for _, r := range word {
		if !unicode.In(r, unicode.Hiragana) {
			return nil, &ErrNotHiragana{word}
		}
	}
	w, err := NewWord(word)
	if err != nil {
		return nil, err
	}
	return &HiraganaWord{
		w: w,
	}, nil
}

func GetFirstHiraganaWord() (*HiraganaWord, error) {
	return NewHiraganaWord("しりとり")
}

func (hw *HiraganaWord) String() string {
	return hw.w.String()
}

func (hw *HiraganaWord) MarshalBinary() ([]byte, error) {
	return hw.w.MarshalBinary()
}

func (hw *HiraganaWord) UnmarshalBinary(data []byte) error {
	new, err := NewHiraganaWord(string(data))
	*hw = *new
	return err
}

func (hw *HiraganaWord) FirstChr() rune {
	return hw.w.FirstChr()
}

func (hw *HiraganaWord) LastChr() rune {
	return hw.w.LastChr()
}

func (hw HiraganaWord) ValidateChain(next IWord) error {
	nhw, ok := next.(*HiraganaWord)
	if !ok || (nhw.LastChr() == 'ん') {
		return &ErrWordInvalid{
			Word: next,
		}
	}
	return hw.w.ValidateChain(&(nhw.w))
}

func (e *ErrNotHiragana) Error() string {
	return fmt.Sprintf("'%s' は、ひらがなではありません", e.Word)
}

func (e *ErrHiraganaWordInvalid) Error() string {
	ew := ErrWordInvalid{e.Word}
	return ew.Error()
}
