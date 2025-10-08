package model

import (
	"github.com/mroth/weightedrand/v2"
)

type Pack int

// PackOpen opens a pack and returns the numbers inside.
func PackOpen(pack Pack) ([]NumberResult, error) {
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

	var results []NumberResult
	for range 5 {
		number := picker.Pick()
		isNew, err := number.AddToCollection()
		if err != nil {
			return nil, err
		}
		results = append(results, NumberResult{Number: number, New: isNew})
	}

	return results, nil
}
