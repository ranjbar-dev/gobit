package validationtool

import (
	"fmt"
)

func IntMinAndMax(value int, name string, min int, max int) error {

	if err := IntMin(value, name, min); err != nil {

		return err
	}

	return IntMax(value, name, max)
}

func IntMin(value int, name string, min int) error {

	if value < min {

		return fmt.Errorf("min value for %s is %d", name, min)
	}

	return nil
}

func IntMax(value int, name string, max int) error {

	if value > max {

		return fmt.Errorf("max value for %s is %d", name, max)
	}

	return nil
}
