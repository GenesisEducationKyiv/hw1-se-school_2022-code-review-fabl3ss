package utils

import (
	"strconv"
)

func StringToFloat64(val string) (float64, error) {
	bitSize := 64
	return strconv.ParseFloat(val, bitSize)
}
