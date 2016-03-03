package gozmo

import (
	"fmt"
)

func CastFloat32(number interface{}) (float32, error) {
	valueF32, ok := number.(float32)
	if ok {
		return valueF32, nil
	}
	valueF64, ok := number.(float64)
	if ok {
		return float32(valueF64), nil
	}
	valueU32, ok := number.(uint32)
	if ok {
		return float32(valueU32), nil
	}
	valueU64, ok := number.(uint64)
	if ok {
		return float32(valueU64), nil
	}
	valueI32, ok := number.(int32)
	if ok {
		return float32(valueI32), nil
	}
	valueI64, ok := number.(int64)
	if ok {
		return float32(valueI64), nil
	}
	return 0, fmt.Errorf("expects a float32")
}

func CastUInt32(number interface{}) (uint32, error) {
	valueUI32, ok := number.(uint32)
	if ok {
		return valueUI32, nil
	}
	valueUI64, ok := number.(uint64)
	if ok {
		return uint32(valueUI64), nil
	}
	valueF32, ok := number.(float32)
	if ok {
		return uint32(valueF32), nil
	}
	valueF64, ok := number.(float64)
	if ok {
		return uint32(valueF64), nil
	}
	return 0, fmt.Errorf("expects a uint32")
}

func CastInt(number interface{}) (int, error) {
	valueI, ok := number.(int)
	if ok {
		return valueI, nil
	}
	valueUI, ok := number.(uint)
	if ok {
		return int(valueUI), nil
	}
	valueUI32, ok := number.(uint32)
	if ok {
		return int(valueUI32), nil
	}
	valueUI64, ok := number.(uint64)
	if ok {
		return int(valueUI64), nil
	}
	valueF32, ok := number.(float32)
	if ok {
		return int(valueF32), nil
	}
	valueF64, ok := number.(float64)
	if ok {
		return int(valueF64), nil
	}
	return 0, fmt.Errorf("expects a int")
}
