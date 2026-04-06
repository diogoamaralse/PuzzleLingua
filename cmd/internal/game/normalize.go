package game

import "strings"

func Normalize(s string) string {
	s = strings.TrimSpace(strings.ToLower(s))

	replacer := strings.NewReplacer(
		"á", "a",
		"à", "a",
		"ã", "a",
		"â", "a",
		"ä", "a",
		"é", "e",
		"ê", "e",
		"è", "e",
		"ë", "e",
		"í", "i",
		"ì", "i",
		"ï", "i",
		"ó", "o",
		"ô", "o",
		"õ", "o",
		"ò", "o",
		"ö", "o",
		"ú", "u",
		"ù", "u",
		"ü", "u",
		"ç", "c",
	)
	return replacer.Replace(s)
}
