package util

import (
	"strconv"
	"strings"
)

func CompareString(f, s string) bool {
	arr1, arr2 := strings.Split(f, " "), strings.Split(s, " ")
	for _, v1 := range arr1 {
		for _, v2 := range arr2 {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}

func StringToInt(val string) int {
	out, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return out
}

func FloatToString(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}
