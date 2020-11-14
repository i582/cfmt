package internal

import (
	"log"
	"strings"
	"unicode"
)

// Parse parses the passed format string and applies styles, returning a styled string.
func Parse(format string) string {
	var tempToken = make([]rune, 0, 20)
	var resParts = make([]string, 0, strings.Count(format, "::"))
	var lastToken string

	var inFormat bool
	var inStartFormat bool

	var inFormatGroup bool
	var inStartFormatGroup bool
	var inEndFormatGroup bool

	var lastIsFormatGroup bool
	var needCheckFormatGroupStyle bool

	for index, s := range format {
		// find {{
		if s == '{' && index+1 < len(format) && format[index+1] == '{' {
			if len(tempToken) != 0 {
				resParts = append(resParts, string(tempToken))
				tempToken = nil
			}
			inFormatGroup = true
			inStartFormatGroup = true
			continue
		}

		// skip {
		if inStartFormatGroup && s == '{' {
			inStartFormat = false
			continue
		}

		// skip }
		if inEndFormatGroup && s == '}' {
			inEndFormatGroup = false
			continue
		}

		// if after {{}} no :: style
		if needCheckFormatGroupStyle {
			needCheckFormatGroupStyle = false

			if !(s == ':' && index+1 < len(format) && format[index+1] == ':' && index+2 < len(format) && format[index+2] != ' ') {
				lastIsFormatGroup = false

				if lastToken != "" {
					resParts = append(resParts, "{{"+lastToken+"}}")
				}
			}
		}

		if inFormatGroup {
			if s == '}' && index+1 < len(format) && format[index+1] == '}' {
				inFormatGroup = false
				inStartFormatGroup = false
				inEndFormatGroup = true
				lastIsFormatGroup = true
				needCheckFormatGroupStyle = true

				lastToken = string(tempToken)

				tempToken = nil
				continue
			}

			tempToken = append(tempToken, s)
			continue
		}

		// ::
		if s == ':' && index+1 < len(format) && format[index+1] == ':' && index+2 < len(format) && format[index+2] != ' ' {
			if !lastIsFormatGroup {
				lastToken = string(tempToken)
				tempToken = nil
			}

			inFormat = true
			inStartFormat = true
			continue
		}

		// skip :
		if inStartFormat && s == ':' {
			inStartFormat = false
			continue
		}

		if inFormat {
			if !(s >= 'a' && s <= 'z' || s >= 'A' && s <= 'Z') && s != '|' && !(s >= '0' && s <= '9') && s != '#' {
				if lastIsFormatGroup {
					lastIsFormatGroup = false
					lastToken = groupStyle(format, lastToken, string(tempToken))
				} else {
					lastToken = singleStyle(format, lastToken, string(tempToken))
				}

				resParts = append(resParts, lastToken)

				tempToken = make([]rune, 0, 20)
				tempToken = append(tempToken, s)
				inFormat = false
				continue
			}
		}

		tempToken = append(tempToken, s)
	}

	if len(lastToken) != 0 {
		if lastIsFormatGroup {
			lastToken = groupStyle(format, lastToken, string(tempToken))
			resParts = append(resParts, lastToken)
		} else if inFormat {
			lastToken = singleStyle(format, lastToken, string(tempToken))
			resParts = append(resParts, lastToken)
		} else {
			resParts = append(resParts, string(tempToken))
		}
	}

	return strings.Join(resParts, "")
}

func singleStyle(format string, token string, style string) string {
	styler, err := styleBuilder(style)
	if err != nil {
		log.Fatalf("Error parse style string in '%s' format string: %v", format, err)
	}
	rem, last, oneWord := lastWord(token)

	if oneWord {
		rem = styler(rem)
	} else {
		last = styler(last)
	}

	return rem + last
}

func groupStyle(format string, token string, style string) string {
	styler, err := styleBuilder(style)
	if err != nil {
		log.Fatalf("Error parse style string in '%s' format string: %v", format, err)
	}

	return styler(token)
}

func lastWord(text string) (remainingText string, word string, oneWord bool) {
	runes := []rune(text)

	for i := len(runes) - 1; i >= 0; i-- {
		if isWordSeparator(runes[i]) {
			startIndex := i + 1
			word := string(runes[startIndex:])
			remainingText := string(runes[0:startIndex])

			return remainingText, word, false
		}
	}

	return text, "", true
}

func isWordSeparator(r rune) bool {
	return !unicode.IsDigit(r) && !unicode.IsLetter(r) && r != '_' && r != '-'
}
