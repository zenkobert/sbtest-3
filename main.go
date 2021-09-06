package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(getSubstringInFirstBracket("Luffy is going to be the (PIRATE)(KING)"))
}

// refactor result
func getSubstringInFirstBracket(str string) string {
	openingBracketIdx := strings.Index(str, "(")

	if openingBracketIdx >= 0 {
		stringAfterOpeningBracket := str[openingBracketIdx:]
		closingBracketIdx := strings.Index(stringAfterOpeningBracket, ")")

		if closingBracketIdx >= 0 {
			return stringAfterOpeningBracket[1:closingBracketIdx]
		}
	}

	return ""
}

// old code
func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1 : indexClosingBracketFound-1])
			} else {
				return ""
			}
		} else {
			return ""
		}
	} else {
		return ""
	}

	return ""
}
