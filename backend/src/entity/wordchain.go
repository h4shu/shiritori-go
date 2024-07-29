package entity

type (
	Wordchain            []Word
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

func (wc *Wordchain) Append(w Word) *Wordchain {
	nwc := append(*wc, w)
	return &nwc
}

func (wc *Wordchain) Len() int {
	return len(*wc)
}

func (wc *Wordchain) ToStrSlice() []string {
	var s []string
	for _, w := range *wc {
		s = append(s, w.String())
	}
	return s
}
