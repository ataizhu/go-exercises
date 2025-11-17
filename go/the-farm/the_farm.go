package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
    cows    int
    message string
}

func DivideFood(fc FodderCalculator, numCows int) (float64, error) {
	fa, err := fc.FodderAmount(numCows)
	if err != nil{
		return 0.0, err
	}
	ff, err := fc.FatteningFactor()
	if err != nil{
		return 0.0, err
	}
	result := (fa * ff) / float64(numCows)

	return result, nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, numCows int) (float64, error) {
	if numCows <= 0{
		return 0.0, errors.New("invalid number of cows")
	}

	return	DivideFood(fc, numCows)

}

func (e *InvalidCowsError) Error() string{
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0{
		return &InvalidCowsError{
			cows: cows,
			message: "there are no negative cows",
		}
	}
	if cows == 0 {
		return &InvalidCowsError{
			cows: cows,
			message: "no cows don't need food",
		}
	}
	return nil
}