package gozmo

import (
	"fmt"
)

func IsTrue(value interface{}, err error) bool {
	if err != nil {
		fmt.Println(err)
		return false
	}
	isTrue, ok := value.(bool)
	if !ok {
		return false
	}
	return isTrue
}

func CastBool(value interface{}) (bool, error) {
	flag, ok := value.(bool)
	if ok {
		return flag, nil
	}
	// special case for allowing numbers as boolean
	num, err := CastFloat32(value)
	if err == nil {
		return num != 0, nil
	}
        return false, fmt.Errorf("expects a bool")
}
