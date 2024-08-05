package outputs

import (
	"github.com/h4shu/shiritori-go/domain/entities"
)

type (
	WordchainGetLastOutputData struct {
		w entities.IWord
	}
	WordchainListOutputData struct {
		wc *entities.Wordchain
	}
)

func NewWordchainGetLastOutputData(w entities.IWord) *WordchainGetLastOutputData {
	return &WordchainGetLastOutputData{
		w: w,
	}
}

func (o *WordchainGetLastOutputData) GetWord() entities.IWord {
	return o.w
}

func NewWordchainListOutputData(wc *entities.Wordchain) *WordchainListOutputData {
	return &WordchainListOutputData{
		wc: wc,
	}
}

func (o *WordchainListOutputData) GetWordchain() *entities.Wordchain {
	return o.wc
}
