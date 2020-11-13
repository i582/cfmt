package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type LastWordSuite struct {
	Value     string
	Remaining string
	Last      string
	OneWord   bool
}

func TestLastWord(t *testing.T) {
	suites := []LastWordSuite{
		{
			Value:     "text",
			Last:      "",
			Remaining: "text",
			OneWord:   true,
		},
		{
			Value:     "some text",
			Last:      "text",
			Remaining: "some ",
			OneWord:   false,
		},
		{
			Value:     "some, text",
			Last:      "text",
			Remaining: "some, ",
			OneWord:   false,
		},
		{
			Value:     "some, dash-text",
			Last:      "dash-text",
			Remaining: "some, ",
			OneWord:   false,
		},
		{
			Value:     "some, under_text",
			Last:      "under_text",
			Remaining: "some, ",
			OneWord:   false,
		},
		{
			Value:     "some, 123text",
			Last:      "123text",
			Remaining: "some, ",
			OneWord:   false,
		},
	}

	for _, suite := range suites {
		p := Parser{}
		rem, word, oneWord := p.lastWord(suite.Value)
		if rem != suite.Remaining {
			t.Error(cmp.Diff(rem, suite.Remaining))
		}
		if word != suite.Last {
			t.Error(cmp.Diff(word, suite.Last))
		}
		if oneWord != suite.OneWord {
			t.Error(cmp.Diff(oneWord, suite.OneWord))
		}
	}
}
