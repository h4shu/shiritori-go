package entity

type (
	Wordchain struct {
		words    []IWord
		wordType WordType
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

func NewWordchain(t WordType) *Wordchain {
	return &Wordchain{
		words:    []IWord{},
		wordType: t,
	}
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
	nwc := NewWordchain(wc.wordType)
	nwc.words = append(wc.words, w)
	return nwc, nil
}

func (wc *Wordchain) Len() int {
	return len(wc.words)
}

func (wc *Wordchain) ToStrSlice() []string {
	var s []string
	for _, w := range wc.words {
		s = append(s, w.String())
	}
	return s
}
