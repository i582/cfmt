package tests

import (
	"testing"

	"github.com/i582/cfmt"
)

func TestParse(t *testing.T) {
	cfmt.RegisterStyle("code", func(s string) string {
		return cfmt.Sprintf("{{%s}}::red|underline", s)
	})

	cfmt.Println("{{correct group}}::code")
	cfmt.Println("{{correct group}}::red|underline and {{other}}::red")
	cfmt.Print("{{error group}} \n")
	cfmt.Print("{{overline group}}::overline\n")
	cfmt.Print("{{reverse group}}::reverse\n")
	cfmt.Print(cfmt.Sprintln("{{faint group}}::faint"))
	cfmt.Println(cfmt.Sprint("{{blink group}}::blink"))
	cfmt.Printf("{{hex %s}}::#ff00ff\n", "color group")
	cfmt.Printf(cfmt.Sprintf("{{background color %s}}::bg#ffff00", "hex color"))
}
