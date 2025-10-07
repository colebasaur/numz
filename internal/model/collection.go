package model

import (
	"github.com/colebasaur/numz/internal/storage"
)

type CollectionItem struct {
	Number   Number
	Quantity int
}

// CollectionItemsGet retrieves all items in the collection from the database.
// Should we ever move to shared database, this will need to be updated to include user context.
func CollectionItemsGet() ([]CollectionItem, error) {
	db, err := storage.GetDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Conn.Query("SELECT number_id, quantity FROM collection")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []CollectionItem
	for rows.Next() {
		var item CollectionItem
		if err := rows.Scan(&item.Number, &item.Quantity); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// CollectionItemAdd adds a number to the collection in the database.
func CollectionItemAdd(number Number) error {
	db, err := storage.GetDB()
	if err != nil {
		return err
	}

	_, err = db.Conn.Exec(`
		INSERT INTO collection
			(number_id, quantity)
		VALUES
			(?, ?)
		ON CONFLICT(number_id)
		DO UPDATE SET
			quantity = quantity + ?
		`, number)
	return err
}
