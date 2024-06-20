package helper

import (
	"strings"
)

// SplitDate ...
func SplitDate(data string) string {
	dataArr := strings.Split(data, "-")
	if len(dataArr) != 3 {
		return ""
	}

	return dataArr[0] + dataArr[1] + dataArr[2]
}
