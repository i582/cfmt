package internal

import (
	"log"
	"strings"
)

// Parse parses the passed format string and applies styles, returning a styled string.
func Parse(format string) string {
	var tempToken = make([]rune, 0, 20)
	var resParts = make([]string, 0, strings.Count(format, "::"))
	var lastToken string

	var inText bool

	var inFormat bool
	var inStartFormat bool

	var inFormatGroup bool
	var inStartFormatGroup bool
	var inEndFormatGroup bool

	var lastIsFormatGroup bool
	var needCheckFormatGroupStyle bool

	inText = true

	for index, s := range format {
		if inText {
			// find {{
			if s == '{' && index+1 < len(format) && format[index+1] == '{' {
				if len(tempToken) != 0 {
					resParts = append(resParts, string(tempToken))
					tempToken = nil
				}
				inFormatGroup = true
				inFormat = false
				inText = false

				inStartFormatGroup = true
				continue
			}

			// skip }
			if inEndFormatGroup && s == '}' {
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

			// ::
			if inEndFormatGroup && s == ':' && index+1 < len(format) && format[index+1] == ':' && index+2 < len(format) && format[index+2] != ' ' {
				// lastToken = string(tempToken)
				// tempToken = nil

				inFormatGroup = false
				inFormat = true
				inText = false

				inStartFormat = true
				continue
			}

			tempToken = append(tempToken, s)
			continue
		}

		if inFormatGroup {
			// skip {
			if inStartFormatGroup && s == '{' {
				inStartFormatGroup = false
				continue
			}

			if s == '}' && index+1 < len(format) && format[index+1] == '}' {

				inFormatGroup = false
				inFormat = false
				inText = true

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

		if inFormat {
			// skip :
			if inStartFormat && s == ':' {
				inStartFormat = false
				continue
			}

			if !(s >= 'a' && s <= 'z' || s >= 'A' && s <= 'Z') && s != '|' && !(s >= '0' && s <= '9') && s != '#' {
				lastToken = groupStyle(format, lastToken, string(tempToken))

				resParts = append(resParts, lastToken)

				tempToken = make([]rune, 0, 20)
				tempToken = append(tempToken, s)

				inFormatGroup = false
				inFormat = false
				inText = true

				inEndFormatGroup = false
				lastIsFormatGroup = false
				continue
			}

			tempToken = append(tempToken, s)
			continue
		}
	}

	if len(lastToken) != 0 {
		if lastIsFormatGroup {
			lastToken = groupStyle(format, lastToken, string(tempToken))
			resParts = append(resParts, lastToken)
		} else {
			resParts = append(resParts, string(tempToken))
		}
	}

	return strings.Join(resParts, "")
}

func groupStyle(format string, token string, style string) string {
	styler, err := styleBuilder(style)
	if err != nil {
		log.Fatalf("Error parse style string in '%s' format string: %v", format, err)
	}

	return styler(token)
}
