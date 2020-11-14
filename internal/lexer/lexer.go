package lexer

import (
	"strings"
)

type tokenType uint8

const (
	Text tokenType = iota
	FormatWord
	FormatGroup
	InvalidGroup
)

type token struct {
	Value string
	Type  tokenType
}

type Lexer struct{}

func (l *Lexer) Tokenize(format string) []token {
	var tokens = make([]token, 0, strings.Count(format, "::"))
	var tempToken = make([]rune, 0, 20)

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
				tokens = append(tokens, token{
					Value: string(tempToken),
					Type:  Text,
				})
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
				if len(tokens) != 0 {
					tokens[len(tokens)-1].Type = InvalidGroup
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
				tokens = append(tokens, token{
					Value: string(tempToken),
					Type:  Text,
				})
				tempToken = nil
				continue
			}

			tempToken = append(tempToken, s)
			continue
		}

		// ::
		if s == ':' && index+1 < len(format) && format[index+1] == ':' && index+2 < len(format) && format[index+2] != ' ' {
			if len(tempToken) != 0 {
				tokens = append(tokens, token{
					Value: string(tempToken),
					Type:  Text,
				})
			}

			tempToken = nil
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
				tp := FormatWord
				if lastIsFormatGroup {
					tp = FormatGroup
					lastIsFormatGroup = false
				}
				tokens = append(tokens, token{
					Value: string(tempToken),
					Type:  tp,
				})
				tempToken = make([]rune, 0, 20)
				tempToken = append(tempToken, s)
				inFormat = false
				continue
			}
		}

		tempToken = append(tempToken, s)
	}

	if len(tempToken) != 0 {
		tp := Text
		if lastIsFormatGroup {
			tp = FormatGroup
		} else if inFormat {
			tp = FormatWord
		}
		tokens = append(tokens, token{
			Value: string(tempToken),
			Type:  tp,
		})
	}

	return tokens
}
