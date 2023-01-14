package objective

import (
	"fmt"
)

func ReplaceAtIndexTo(str string, replacement interface{}, startIndex int, endIndex int) string {
	if endIndex < startIndex {
		return fmt.Sprintf("%v%v", str[:startIndex], replacement)
	}
	return fmt.Sprintf("%v%v%v", str[:startIndex], replacement, str[endIndex:])
}
