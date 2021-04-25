package tests

import (
	"testing"

	"github.com/i582/cfmt/cmd/cfmt"
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
			Value: []string{"underline", "reverse"},
			Error: "",
		},
		{
			Value: []string{"bg#ff0", "bold"},
			Error: "",
		},
		{
			Value: []string{"#ff0", "bold"},
			Error: "",
		},
		{
			Value: []string{"#ff01", "bold"},
			Error: "invalid '#ff01' hex color",
		},
		{
			Value: []string{"bg#ff01", "bold"},
			Error: "invalid '#ff01' hex color",
		},
		{
			Value: []string{"some", "bold"},
			Error: "unknown style 'some'",
		},
		{
			Value: []string{},
			Error: "style string is empty",
		},
	}

	for _, suite := range suites {
		_, err := internal.StyleBuilder(suite.Value, "")
		if err == nil && suite.Error != "" {
			t.Errorf("mismatch\nwant: ''\nhave: %s", suite.Error)
			continue
		}

		if err != nil && err.Error() != suite.Error {
			t.Errorf("mismatch\nwant: %s\nhave: %s", suite.Error, err.Error())
		}
	}
}
