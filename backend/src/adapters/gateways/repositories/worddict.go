package repositories

import (
	"context"
	"database/sql"

	"github.com/h4shu/shiritori-go/domain/entities"
)

type (
	WorddictRepository struct {
		db *sql.DB
		t  entities.WordType
	}
)

func NewWorddictRepository(db *sql.DB, t entities.WordType) *WorddictRepository {
	return &WorddictRepository{
		db: db,
		t:  t,
	}
}

func (r *WorddictRepository) Exist(ctx context.Context, word entities.IWord) error {
	var query string
	switch r.t {
	case entities.WordTypeHiragana:
		query = "SELECT 1 FROM words WHERE hiragana = ? limit 1"
	default:
		return &entities.ErrWordTypeInvalid{
			Word:     word,
			WordType: r.t,
		}
	}
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var dummy string
	err = stmt.QueryRowContext(ctx, word.String()).Scan(&dummy)
	if err == sql.ErrNoRows {
		return &entities.ErrWordNotFound{
			Word: word,
		}
	} else if err == nil {
		return nil
	}
	return err
}
