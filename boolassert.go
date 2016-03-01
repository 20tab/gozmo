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
	return isTrue;
}
