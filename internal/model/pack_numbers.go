package model

import (
	"github.com/colebasaur/numz/internal/storage"
)

type PackNumber struct {
	Pack   Pack
	Number Number
	Rarity int
}

// PacksGet retrieves all packs from the database.
func PacksGet() ([]Pack, error) {
	db, err := storage.GetDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Conn.Query("SELECT DISTINCT pack_id FROM pack_numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packs []Pack
	for rows.Next() {
		var pack Pack
		if err := rows.Scan(&pack); err != nil {
			return nil, err
		}
		packs = append(packs, pack)
	}

	return packs, nil
}

// GetPackNumbers retrieves all numbers associated with a pack from the database.
func GetPackNumbers(packID Pack) ([]PackNumber, error) {
	db, err := storage.GetDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Conn.Query(`
		SELECT pack_id, number_id, rarity
			FROM pack_numbers
			WHERE pack_id = ?
	`, packID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packNumbers []PackNumber
	for rows.Next() {
		var pn PackNumber
		if err := rows.Scan(&pn.Pack, &pn.Number, &pn.Rarity); err != nil {
			return nil, err
		}
		packNumbers = append(packNumbers, pn)
	}

	return packNumbers, nil
}
