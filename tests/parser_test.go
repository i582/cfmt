package tests

import (
	"testing"

	"github.com/i582/cfmt"
)

type TestStruct struct {
	field1 string
	field2 int
}

func TestParse(t *testing.T) {
	cfmt.RegisterStyle("code", func(s string) string {
		return cfmt.Sprintf("{{%s}}::red|underline", s)
	})

	cfmt.Println("{{こんにちは, correct group}}::code  sdas")
	cfmt.Println("{{привет, correct group}}::red|underline and {{other}}::red")
	cfmt.Print("{{error group}} \n")
	cfmt.Print("{{overline group}}::overline\n")
	cfmt.Print("{{reverse group, こんにちは}}::reverse\n")
	cfmt.Print(cfmt.Sprintln("{{faint group}}::faint sfafs"))
	cfmt.Println(cfmt.Sprint("{{blink group}}::blink"))
	cfmt.Printf("{{hex %s}}::#ff00ff sfas\n", "color group")
	cfmt.Printf(cfmt.Sprintf("{{background color %s}}::bg#ffff00\n", "hex color"))
	cfmt.Printf("{{{hello}}}::red|underline\n")
	cfmt.Printf("{{some test struct: %v}}::red|underline\n", TestStruct{"hello", 1})
}
