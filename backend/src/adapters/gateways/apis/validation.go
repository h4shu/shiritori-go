package apis

import "github.com/h4shu/shiritori-go/domain/entities"

type (
	IValidateWordClient interface {
		ValidateJapanese(word string) (bool, error)
	}
	ValidateWordApi struct {
		client IValidateWordClient
	}
)

func NewValidateWordApi(client IValidateWordClient) *ValidateWordApi {
	return &ValidateWordApi{
		client: client,
	}
}

func (a *ValidateWordApi) Validate(word entities.IWord, t entities.WordType) (ret bool, err error) {
	switch t {
	case entities.WordTypeHiragana:
		ret, err = a.client.ValidateJapanese(word.String())
	default:
		ret, err = false, &entities.ErrWordTypeInvalid{
			Word:     word,
			WordType: t,
		}
	}
	return
}
