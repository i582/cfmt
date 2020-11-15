package tests

import (
	"testing"

	"github.com/i582/cfmt"
)

func TestParse(t *testing.T) {
	cfmt.Println("{{correct group}}::red|underline")
	cfmt.Println("{{correct group}}::red|underline and {{other}}::red")
	cfmt.Println("{{error group}} ")
	cfmt.Println("{{overline group}}::overline")
	cfmt.Println("{{reverse group}}::reverse")
	cfmt.Println("{{faint group}}::faint")
	cfmt.Println("{{blink group}}::blink")
	cfmt.Println("{{hex color group}}::#ff00ff")
	cfmt.Println("{{background color hex group}}::bg#ffff00")
}
