package apis

import "github.com/h4shu/shiritori-go/domain/entities"

type (
	IValidateWordApi interface {
		Validate(word entities.IWord, t entities.WordType) (bool, error)
	}
)
