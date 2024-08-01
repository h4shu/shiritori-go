package entities

import (
	"fmt"
)

type (
	Wordchain struct {
		words    []IWord
		wordType WordType
	}
	ErrWordchainInvalid struct {
		Wordchain *Wordchain
		Next      IWord
	}
	GetWordchainRequest  struct{}
	GetWordchainResponse struct {
		Wordchain []string `json:"wordchain"`
		Len       int      `json:"len,string"`
	}
	AddWordchainRequest struct {
		Word string `json:"word" form:"word"`
	}
	AddWordchainResponse struct{}
)

func NewWordchain(t WordType) (*Wordchain, error) {
	w, err := GetFirstWordForType(t)
	if err != nil {
		return nil, err
	}
	return &Wordchain{
		words:    []IWord{w},
		wordType: t,
	}, nil
}

func (wc *Wordchain) Len() int {
	return len(wc.words)
}

func (wc *Wordchain) GetLast() IWord {
	return wc.words[len(wc.words)-1]
}

func (wc *Wordchain) Append(w IWord) (*Wordchain, error) {
	ok := false
	switch wc.wordType {
	case WordTypeHiragana:
		_, ok = w.(*HiraganaWord)
	default:
		_, ok = w.(*Word)
	}
	if !ok {
		return nil, &ErrWordTypeInvalid{
			Word:     w,
			WordType: wc.wordType,
		}
	}
	lw := wc.GetLast()
	if !lw.ValidateChain(w) {
		return nil, &ErrWordchainInvalid{
			Wordchain: wc,
			Next:      w,
		}
	}
	return &Wordchain{
		words:    append(wc.words, w),
		wordType: wc.wordType,
	}, nil
}

func (wc *Wordchain) AppendStr(w string) (*Wordchain, error) {
	nw, err := NewWordWithType(w, wc.wordType)
	if err != nil {
		return nil, err
	}
	lw := wc.GetLast()
	if !lw.ValidateChain(nw) {
		return nil, &ErrWordchainInvalid{
			Wordchain: wc,
			Next:      nw,
		}
	}
	return &Wordchain{
		words:    append(wc.words, nw),
		wordType: wc.wordType,
	}, nil
}

func (wc *Wordchain) ToStrSlice() []string {
	var s []string
	for _, w := range wc.words {
		s = append(s, w.String())
	}
	return s
}

func (e *ErrWordchainInvalid) Error() string {
	return fmt.Sprintf("しりとりが正しくありません: %s => %s", e.Wordchain.GetLast(), e.Next)
}
