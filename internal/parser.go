package internal

import (
	"log"
	"strings"
)

// Parse parses the passed format string and applies styles, returning a styled string.
func Parse(format string) string {
	var resParts = make([]string, 0, 50)

	var err error

	var tempTokenStartIndex int
	var tempTokenEndIndex int

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
	formatLen := len(format)

	for index := 0; index < formatLen; index++ {
		s := format[index]

		if inText {
			// find {{
			if s == '{' && index+1 < formatLen && format[index+1] == '{' {
				tempTokenEndIndex = index
				resParts = append(resParts, format[tempTokenStartIndex:tempTokenEndIndex])

				tempTokenStartIndex = index + 2

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

				if !(s == ':' && index+1 < formatLen && format[index+1] == ':' && index+2 < formatLen && format[index+2] != ' ') {
					lastIsFormatGroup = false

					if lastToken != "" {
						resParts = append(resParts, "{{"+lastToken+"}}")
					}
				}
			}

			// ::
			if inEndFormatGroup && s == ':' && index+1 < formatLen && format[index+1] == ':' && index+2 < formatLen && format[index+2] != ' ' {
				inFormatGroup = false
				inFormat = true
				inText = false

				inStartFormat = true
				tempTokenStartIndex = index + 2
				continue
			}

			continue
		}

		if inFormatGroup {
			// skip {
			if inStartFormatGroup && s == '{' {
				inStartFormatGroup = false
				continue
			}

			if s == '}' && index+1 < formatLen && format[index+1] == '}' {
				inFormatGroup = false
				inFormat = false
				inText = true

				inEndFormatGroup = true
				lastIsFormatGroup = true
				needCheckFormatGroupStyle = true

				tempTokenEndIndex = index

				lastToken = format[tempTokenStartIndex:tempTokenEndIndex]
				continue
			}
			continue
		}

		if inFormat {
			// skip :
			if inStartFormat && s == ':' {
				inStartFormat = false
				continue
			}

			if s == '|' {
				tempTokenEndIndex = index

				singleFormat := format[tempTokenStartIndex:tempTokenEndIndex]
				lastToken, err = ApplyStyle(lastToken, singleFormat)
				if err != nil {
					log.Fatalf("Error parse style string in '%s' format string: %v", format, err)
				}

				tempTokenStartIndex = index + 1
				continue
			}

			if !(s >= 'a' && s <= 'z' || s >= 'A' && s <= 'Z') && s != '|' && !(s >= '0' && s <= '9') && s != '#' {
				tempTokenEndIndex = index

				singleFormat := format[tempTokenStartIndex:tempTokenEndIndex]
				lastToken, err = ApplyStyle(lastToken, singleFormat)
				if err != nil {
					log.Fatalf("Error parse style string in '%s' format string: %v", format, err)
				}

				tempTokenStartIndex = index

				resParts = append(resParts, lastToken)

				inFormatGroup = false
				inFormat = false
				inText = true

				inEndFormatGroup = false
				lastIsFormatGroup = false
				continue
			}
			continue
		}
	}

	if lastIsFormatGroup {
		singleFormat := format[tempTokenStartIndex:]
		lastToken, _ = ApplyStyle(lastToken, singleFormat)
		resParts = append(resParts, lastToken)
	} else {
		resParts = append(resParts, format[tempTokenStartIndex:])
	}

	return strings.Join(resParts, "")
}
