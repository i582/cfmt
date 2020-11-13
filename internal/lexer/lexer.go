package lexer

import (
	"unicode"
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
	var tokens []token
	var tempToken string

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
			if tempToken != "" {
				tokens = append(tokens, token{
					Value: tempToken,
					Type:  Text,
				})
				tempToken = ""
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
					Value: tempToken,
					Type:  Text,
				})
				tempToken = ""
				continue
			}

			tempToken += string(s)
			continue
		}

		// ::
		if s == ':' && index+1 < len(format) && format[index+1] == ':' && index+2 < len(format) && format[index+2] != ' ' {
			if tempToken != "" {
				tokens = append(tokens, token{
					Value: tempToken,
					Type:  Text,
				})
			}

			tempToken = ""
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
			if !unicode.IsLetter(s) && s != '|' && !unicode.IsDigit(s) && s != '#' {
				tp := FormatWord
				if lastIsFormatGroup {
					tp = FormatGroup
					lastIsFormatGroup = false
				}
				tokens = append(tokens, token{
					Value: tempToken,
					Type:  tp,
				})
				tempToken = string(s)
				inFormat = false
				continue
			}
		}

		tempToken += string(s)
	}

	if tempToken != "" {
		tp := Text
		if lastIsFormatGroup {
			tp = FormatGroup
		} else if inFormat {
			tp = FormatWord
		}
		tokens = append(tokens, token{
			Value: tempToken,
			Type:  tp,
		})
	}

	return tokens
}
