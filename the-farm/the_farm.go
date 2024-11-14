package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(calc FodderCalculator, cows int) (ret float64, err error) {
	ret = 0
	amount, err := calc.FodderAmount(cows)
	if err == nil {
		factor, err := calc.FatteningFactor()
		if err == nil {
			ret = amount * factor / float64(cows)
		} else {
			return 0, err
		}
	}
	return
}

func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(fodderCalculator, cows)
	}
	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

func ValidateNumberOfCows(cows int) error {
	var err error
	switch {
	case cows < 0:
		err = &InvalidCowsError{
			cows:    cows,
			message: "there are no negative cows",
		}
	case cows == 0:
		err = &InvalidCowsError{
			cows:    cows,
			message: "no cows don't need food",
		}
	}

	return err
}
