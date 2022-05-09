package error

import "fmt"

type Ninja04Error struct {
	lat  string
	long string
	err  error
}

func (err Ninja04Error) Error() string {
	return fmt.Sprintf("%v %v: %v", err.lat, err.long, err.err)
}

func Ninja04() {
	_, err := ninja04Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	}
}

func ninja04Sqrt(number float64) (float64, error) {
	if number < 0 {
		// return 0, errors.New("cannot be negative")
		return 0, Ninja04Error{"50.2289N", "99.4656W", fmt.Errorf("cannot be negative %v", number)}
	}

	return 42, nil
}
