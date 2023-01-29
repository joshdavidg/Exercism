package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	numOfCows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.numOfCows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()

	if fodder < 0 {
		if err == ErrScaleMalfunction || err == nil {
			return 0.0, errors.New("negative fodder")
		}

		return 0.0, err
	}

	if err == ErrScaleMalfunction && fodder > 0 {
		return (fodder * 2) / float64(cows), nil
	}

	if err != nil {
		return 0.0, err
	}

	if cows < 0 {
		return 0.0, &SillyNephewError{
			numOfCows: cows,
		}
	}

	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}

	return fodder / float64(cows), nil
}
