package model

import (
	"github.com/mroth/weightedrand/v2"
)

type Pack int

// PackOpen opens a pack and returns the numbers inside.
func PackOpen(pack Pack) ([]Number, error) {
	packNumbers, err := GetPackNumbers(pack)
	if err != nil {
		return nil, err
	}

	choices := make([]weightedrand.Choice[Number, int], len(packNumbers))
	for i, pn := range packNumbers {
		choices[i] = weightedrand.NewChoice(pn.Number, pn.Rarity)
	}

	picker, err := weightedrand.NewChooser(choices...)
	if err != nil {
		return nil, err
	}

	var results []Number
	for range 5 {
		number := picker.Pick()
		number.AddToCollection()
		results = append(results, number)
	}

	return results, nil
}
