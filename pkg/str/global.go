package str

import (
	"strconv"
)

// ShowString ...
func ShowString(isShow bool, data string) string {
	if isShow {
		return data
	}

	return ""
}

// StringToBool ...
func StringToBool(data string) bool {
	res, err := strconv.ParseBool(data)
	if err != nil {
		res = false
	}

	return res
}

// StringToBoolString ...
func StringToBoolString(data string) string {
	res, err := strconv.ParseBool(data)
	if err != nil {
		return "false"
	}

	return strconv.FormatBool(res)
}

// StringToInt ...
func StringToInt(data string) int {
	res, err := strconv.Atoi(data)
	if err != nil {
		res = 0
	}

	return res
}

// Contains ...
func Contains(slices []string, comparizon string) bool {
	for _, a := range slices {
		if a == comparizon {
			return true
		}
	}

	return false
}

// DefaultData ...
func DefaultData(data, defaultData string) string {
	if data == "" {
		return defaultData
	}

	return data
}
