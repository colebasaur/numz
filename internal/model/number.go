package model

import (
	"github.com/colebasaur/numz/internal/storage"
)

type Number int

// AddToCollection adds a number to the collection in the database.
func (number Number) AddToCollection() error {
	db, err := storage.GetDB()
	if err != nil {
		return err
	}

	_, err = db.Conn.Exec(`
		INSERT INTO collection
			(number_id, quantity)
		VALUES
			(?, 1)
		ON CONFLICT(number_id)
		DO UPDATE SET
			quantity = quantity + 1
		`, number)

	return err
}
