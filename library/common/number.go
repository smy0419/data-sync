package common

import "strconv"

func IntToString(val int64) string {
	return strconv.FormatInt(val, 10)
}

func UIntToString(val uint64) string {
	return strconv.FormatUint(val, 10)
}

func FloatToString(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}

func CalculateEfficiency(actualBlocks int64, planedBlocks int64) int32 {
	return int32(float64(actualBlocks) / (float64(planedBlocks)) * 100)
}
