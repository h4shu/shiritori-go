package entities

import "fmt"

type (
	WordType           int
	ErrWordTypeInvalid struct {
		Word     IWord
		WordType WordType
	}
	ErrNoFirstWord struct {
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

func GetFirstWordForType(t WordType) (IWord, error) {
	switch t {
	case WordTypeHiragana:
		return GetFirstHiraganaWord()
	}
	return nil, &ErrNoFirstWord{t}
}

func (e *ErrWordTypeInvalid) Error() string {
	return fmt.Sprintf("Invalid word type: %v with %v", e.Word, e.WordType)
}

func (e *ErrNoFirstWord) Error() string {
	return fmt.Sprintf("No first word for %v", e.WordType)
}
