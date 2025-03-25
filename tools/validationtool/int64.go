package validationtool

import (
	"fmt"
)

func Int64MinAndMax(value int64, name string, min int64, max int64) error {

	if err := Int64Min(value, name, min); err != nil {

		return err
	}

	return Int64Max(value, name, max)
}

func Int64Min(value int64, name string, min int64) error {

	if value < min {

		return fmt.Errorf("min value for %s is %d", name, min)
	}

	return nil
}

func Int64Max(value int64, name string, max int64) error {

	if value > max {

		return fmt.Errorf("max value for %s is %d", name, max)
	}

	return nil
}
