package parser

import (
	"log"
	"unicode"

	"github.com/i582/cfmt/internal/lexer"
	"github.com/i582/cfmt/internal/styles"
)

type Parser struct{}

func (p *Parser) Parse(format string) string {
	l := lexer.Lexer{}
	tokens := l.Tokenize(format)

	for index, token := range tokens {
		if token.Type == lexer.FormatWord {
			if index-1 >= 0 {
				text := tokens[index-1].Value
				styler, err := styles.StyleBuilder(token.Value)
				if err != nil {
					log.Fatalf("Error parse style string in '%s' format string: %v", format, err)
				}
				rem, last, oneWord := p.lastWord(text)

				if oneWord {
					rem = styler(rem)
				} else {
					last = styler(last)
				}

				text = rem + last
				tokens[index-1].Value = text
			}
		}

		if token.Type == lexer.FormatGroup {
			if index-1 >= 0 {
				text := tokens[index-1].Value
				styler, err := styles.StyleBuilder(token.Value)
				if err != nil {
					log.Fatalf("Error parse style string in '%s' format string: %v", format, err)
				}

				tokens[index-1].Value = styler(text)
			}
		}
	}

	var res string
	for _, token := range tokens {
		if token.Type == lexer.Text {
			res += token.Value
		}
		if token.Type == lexer.InvalidGroup {
			res += "{{" + token.Value + "}}"
		}
	}

	return res
}

func (p *Parser) lastWord(text string) (remainingText string, word string, oneWord bool) {
	runes := []rune(text)

	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]

		if isWordSeparator(r) {
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
