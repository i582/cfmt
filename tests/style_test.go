package tests

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/i582/cfmt"
	"github.com/i582/cfmt/internal"
)

type StyleBuilderSuite struct {
	Value []string
	Error string
}

func TestStyleBuilder(t *testing.T) {
	cfmt.RegisterStyle("code", func(s string) string {
		return cfmt.Sprintf("{{%s}}::red|underline", s)
	})

	suites := []StyleBuilderSuite{
		{
			Value: []string{"red"},
			Error: "",
		},
		{
			Value: []string{"red", "bold"},
			Error: "",
		},
		{
			Value: []string{"#ff00ff", "bold"},
			Error: "",
		},
		{
			Value: []string{"bg#ff00ff", "bold"},
			Error: "",
		},
		{
			Value: []string{"code", "underline", "blink"},
			Error: "",
		},
		{
			Value: []string{"overline", "reverse", "faint"},
			Error: "",
		},
		{
			Value: []string{"bg#ff0", "bold"},
			Error: "invalid hex: length of hex color must be 6",
		},
		{
			Value: []string{"#ff0", "bold"},
			Error: "invalid hex: length of hex color must be 6",
		},
		{
			Value: []string{"some", "bold"},
			Error: "unknown style some",
		},
		{
			Value: []string{},
			Error: "style string is empty",
		},
	}

	for _, suite := range suites {
		_, err := internal.StyleBuilder(suite.Value, "")
		if err == nil && suite.Error != "" {
			t.Error(cmp.Diff("", suite.Error))
			continue
		}

		if err != nil && !cmp.Equal(err.Error(), suite.Error) {
			t.Error(cmp.Diff(err.Error(), suite.Error))
		}
	}
}
