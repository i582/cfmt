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

	cfmt.RegisterStyle("flag", func(s string) string {
		return cfmt.Sprintf("{{--%s}}::green (-%c)", s, s[0])
	})

	flag := "help"
	cfmt.Printf("{{%s}}::flag \n", flag)
	cfmt.Println("{{こんにちは, correct group}}::code  sdas")
	cfmt.Println("{{привет, correct group}}::red|underline and {{other}}::red")
	cfmt.Print("{{error group}} \n")
	cfmt.Print("{{underline group}}::underline\n")
	cfmt.Print("{{reverse group, こんにちは}}::reverse\n")
	cfmt.Print(cfmt.Sprintln("{{some group}}::red sfafs"))
	cfmt.Println(cfmt.Sprint("{{blink group}}::blink"))
	cfmt.Printf("{{hex %s}}::#ff00ff sfas\n", "color group")
	cfmt.Printf("{{3 hex %s}}::#ff0 sfas\n", "color group")
	cfmt.Printf(cfmt.Sprintf("{{background color %s}}::bg#ffff00\n", "hex color"))
	cfmt.Printf("{{{hello}}}::red|underline\n")
	cfmt.Printf("{{some test struct: %v}}::red|underline\n", TestStruct{"hello", 1})
	cfmt.Println("{{hello}}::red{{world}}::green")
}
