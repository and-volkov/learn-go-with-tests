package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	if repeatCount <= 0 {
		repeatCount = 5
	}
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
