package date

import "strings"

var layoutMap = map[string]TimeLayout{
	"Y": YEAR,
	"m": MONTH,
	"d": DATE,
	"H": HOUR,
	"M": MINUTE,
	"s": SECOND,
}

func MakeLayout(delim string, layouts ...TimeLayout) string {
	var result []string
	for _, layout := range layouts {
		result = append(result, string(layout))
	}

	return strings.Join(result, delim)
}

func ReplaceLayout(layout string) string {
	var replaced string
	for _, c := range layout {
		replaced = replaced + GetLayoutWord(string(c))
	}

	return replaced
}

func GetLayoutWord(word string) string {
	result := layoutMap[word]
	if result == "" {
		return word
	}

	return string(result)
}
