package model

import (
	"fmt"
	"github.com/colebasaur/numz/internal/storage"
)

type Number int
type NumberResult struct {
	Number
	New bool
}

func (nr NumberResult) String() string {
	if nr.New {
		return fmt.Sprintf("%d (New!)", nr.Number)
	}
	return fmt.Sprintf("%d", nr.Number)
}

// AddToCollection adds a number to the collection in the database.
// Returns true if the number was inserted, false if updated.
func (number Number) AddToCollection() (bool, error) {
	db, err := storage.GetDB()
	if err != nil {
		return false, err
	}

	result, err := db.Conn.Exec(`
		INSERT INTO collection
			(number_id, quantity)
		VALUES
			(?, 1)
		ON CONFLICT(number_id)
		DO UPDATE SET
			quantity = quantity + 1
		`, number)
	if err != nil {
		return false, err
	}

	lastId, _ := result.LastInsertId()
	if lastId == 0 {
		return false, nil // Updated
	}

	return true, nil // Inserted
}
