package tests

import (
	"testing"

	"github.com/i582/cfmt/cmd/cfmt"
)

func TestDisableColors(t *testing.T) {
	cfmt.DisableColors()
	cfmt.Println("{{привет, correct group}}::red|underline and {{other}}::red")

	cfmt.EnableColors()
	cfmt.Println("{{привет, correct group}}::red|underline and {{other}}::red")
}
