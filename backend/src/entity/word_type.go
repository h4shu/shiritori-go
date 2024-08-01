package entity

import "fmt"

type (
	WordType           int
	ErrWordTypeInvalid struct {
		Word     IWord
		WordType WordType
	}
)

const (
	WordTypeNone WordType = iota
	WordTypeHiragana
)

func NewWordWithType(word string, t WordType) (IWord, error) {
	switch t {
	case WordTypeHiragana:
		return NewHiraganaWord(word)
	}
	w, err := NewWord(word)
	return &w, err
}

func (e *ErrWordTypeInvalid) Error() string {
	return fmt.Sprintf("WordType が異なります {%v, %v}", e.Word, e.WordType)
}
