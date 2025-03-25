package validationtool

import (
	"fmt"
)

func Int32MinAndMax(value int32, name string, min int32, max int32) error {

	if err := Int32Min(value, name, min); err != nil {

		return err
	}

	return Int32Max(value, name, max)
}

func Int32Min(value int32, name string, min int32) error {

	if value < min {

		return fmt.Errorf("min value for %s is %d", name, min)
	}

	return nil
}

func Int32Max(value int32, name string, max int32) error {

	if value > max {

		return fmt.Errorf("max value for %s is %d", name, max)
	}

	return nil
}
