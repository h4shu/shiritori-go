package entities

type (
	Wordchain struct {
		words    []IWord
		wordType WordType
	}
)

func NewWordchain(t WordType) *Wordchain {
	return &Wordchain{
		words:    []IWord{},
		wordType: t,
	}
}

func (wc *Wordchain) Len() int {
	return len(wc.words)
}

func (wc *Wordchain) GetLast() IWord {
	if wc.Len() == 0 {
		return nil
	}
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
	if lw != nil {
		err := lw.ValidateChain(w)
		if err != nil {
			return nil, err
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
	if lw != nil {
		err := lw.ValidateChain(nw)
		if err != nil {
			return nil, err
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
