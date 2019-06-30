package stringutil

import (
	"strings"
)

func Reverse(str string) string {
	var strCopy string
	sliced := strings.Split(str, "")
	for i := len(str) - 1; i >= 0; i-- {
		strCopy += sliced[i]
	}

	return strCopy
}
