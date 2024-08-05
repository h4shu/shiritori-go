package inputs

import (
	"github.com/h4shu/shiritori-go/domain/entities"
)

type (
	WordchainListInputData struct {
		limit int
	}
	WordchainAppendInputData struct {
		w entities.IWord
	}
)

func NewWordchainListInputData(limit int) *WordchainListInputData {
	return &WordchainListInputData{
		limit: limit,
	}
}

func (i *WordchainListInputData) GetLimit() int {
	return i.limit
}

func NewWordchainAppendInputData(w entities.IWord) *WordchainAppendInputData {
	return &WordchainAppendInputData{
		w: w,
	}
}

func (i *WordchainAppendInputData) GetWord() entities.IWord {
	return i.w
}
