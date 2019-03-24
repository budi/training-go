package stringsplosion

import (
	"strings"
)

// StringSplosion will blow up your string.
func StringSplosion(s string) string {
	var str strings.Builder
	for i := 0; i < len(s); i++ {
		str.WriteString(s[:i+1])
	}
	return str.String()
}
