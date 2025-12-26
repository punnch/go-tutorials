package sentence

import (
	"strings"
)

// JoinWithCom joins words with ", " and last word with ", and"
func JoinWithCommas(phrase []string) string {
	switch len(phrase) {
	case 0:
		return ""
	case 1:
		return phrase[0]
	case 2:
		return phrase[0] + " and " + phrase[1]
	default:
		return strings.Join(phrase[:len(phrase)-1], ", ") + ", and " + phrase[len(phrase)-1]
	}
}
