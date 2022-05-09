package error

import (
	"fmt"
)

func ErrorInfoExample() {
	_, err := sqrt(-1)

	if err != nil {
		fmt.Println(err)
	}
}

func sqrt(number float64) (float64, error) {
	if number < 0 {
		// return 0, errors.New("cannot be negative")
		return 0, fmt.Errorf("cannot be negative %v", number)
	}

	return 42, nil
}
