package lexer

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type Suite struct {
	Value  string
	Tokens []token
	Errors []error
}

func TestLexer(t *testing.T) {
	suites := []Suite{
		{
			Value: "Test word::red",
			Tokens: []token{
				{
					Value: "Test word", Type: Text,
				},
				{
					Value: "red", Type: FormatWord,
				},
			},
		},
		{
			Value: "{{Test word}}::red",
			Tokens: []token{
				{
					Value: "Test word", Type: Text,
				},
				{
					Value: "red", Type: FormatGroup,
				},
			},
		},
		{
			Value: "{{Test word}}::red|bold|italic",
			Tokens: []token{
				{
					Value: "Test word", Type: Text,
				},
				{
					Value: "red|bold|italic", Type: FormatGroup,
				},
			},
		},
		{
			Value: "{{Test word}}::#ffffff|bold|italic",
			Tokens: []token{
				{
					Value: "Test word", Type: Text,
				},
				{
					Value: "#ffffff|bold|italic", Type: FormatGroup,
				},
			},
		},
		{
			Value: "{{Test word}}::#ffffff|bg#ff0000|bold|italic",
			Tokens: []token{
				{
					Value: "Test word", Type: Text,
				},
				{
					Value: "#ffffff|bg#ff0000|bold|italic", Type: FormatGroup,
				},
			},
		},
		{
			Value: "{{Test word}}::italic {{red color}}::red",
			Tokens: []token{
				{
					Value: "Test word", Type: Text,
				},
				{
					Value: "italic", Type: FormatGroup,
				},
				{
					Value: " ", Type: Text,
				},
				{
					Value: "red color", Type: Text,
				},
				{
					Value: "red", Type: FormatGroup,
				},
			},
		},
		{
			Value: "{{Test word}}::italic red color::red",
			Tokens: []token{
				{
					Value: "Test word", Type: Text,
				},
				{
					Value: "italic", Type: FormatGroup,
				},
				{
					Value: " red color", Type: Text,
				},
				{
					Value: "red", Type: FormatWord,
				},
			},
		},
		{
			Value: "   Test word::red",
			Tokens: []token{
				{
					Value: "   Test word", Type: Text,
				},
				{
					Value: "red", Type: FormatWord,
				},
			},
		},
		{
			Value: "   Test word::red   ",
			Tokens: []token{
				{
					Value: "   Test word", Type: Text,
				},
				{
					Value: "red", Type: FormatWord,
				},
				{
					Value: "   ", Type: Text,
				},
			},
		},
		{
			Value: "   {{    Test word    }}::red   ",
			Tokens: []token{
				{
					Value: "   ", Type: Text,
				},
				{
					Value: "    Test word    ", Type: Text,
				},
				{
					Value: "red", Type: FormatGroup,
				},
				{
					Value: "   ", Type: Text,
				},
			},
		},
		{
			// invalid group
			Value: "{{Test word}} red::red",
			Tokens: []token{
				{
					Value: "Test word", Type: InvalidGroup,
				},
				{
					Value: " red", Type: Text,
				},
				{
					Value: "red", Type: FormatWord,
				},
			},
		},
	}

	for _, suite := range suites {
		l := Lexer{}
		tokens := l.Tokenize(suite.Value)
		if !cmp.Equal(tokens, suite.Tokens) {
			t.Error(cmp.Diff(tokens, suite.Tokens))
		}
	}
}
