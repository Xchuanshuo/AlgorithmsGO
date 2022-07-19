package test804

import "strings"

var tables = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..",
	".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-",
	".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	var codeMap = make(map[string]bool)
	for _, w := range words {
		var sb strings.Builder
		for _, c := range w {
			sb.WriteString(tables[c-'a'])
		}
		codeMap[sb.String()] = true
	}
	return len(codeMap)
}
