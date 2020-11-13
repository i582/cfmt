package main

import (
	"github.com/gookit/color"

	cfmt "github.com/i582/cfmt/cmd"
)

func main() {

	cfmt.RegisterStyle("code", func(s string) string {
		return cfmt.Sprintf("{{%s}}::red|underline", s)
	})

	// cfmt.Printf("{{Count}} {{sss}}:: ")

	cfmt.Printf("{{Count of lines}}::#ff0|bold|underline: {{%s}}::#E06C75|bold\n", "Hello World")
	cfmt.Printf("This is red::red color")
	cfmt.Printf(`
	
    {{Example of reports}}::bold

    {{                                                                            }}::bgRed
    {{                            Critical errors found                           }}::bgRed|#ffffff
    {{                                                                            }}::bgRed

    {{100}}::#ffffff myStyle := color.{{New(color.FgWhite, color.BgBlack, color.OpBold)}}::code
        {{[100, 17]}}::blue Undefined function New at {{~/projects/test}}::underline:100

    {{101}}::#ffffff {{myStyle}}::code.Print("t")
        {{[101, 0]}}::blue Undefined variable myStyle at {{~/projects/test}}::underline:101

`)

	myStyle := color.New(color.FgWhite, color.BgBlack, color.OpBold)
	myStyle.Print("t")

	// 	cfmt.Printf(`
	//    {{              }}::bgRed           {{              }}::bgRed           {{              }}::bgRed
	//    {{              }}::bgRed           {{              }}::bgRed           {{              }}::bgRed
	//    {{    About     }}::bgRed           {{     This     }}::bgRed           {{     Tool     }}::bgRed
	//    {{              }}::bgRed           {{              }}::bgRed           {{              }}::bgRed
	//    {{              }}::bgRed           {{              }}::bgRed           {{              }}::bgRed
	//
	// `)

	// 	code := `
	//     {{About PHPStats v0.1.0}}::#ff0000|bold|bgBlack
	//
	//   PHPStats::red|bgBlack is a tool for {{collecting project statistics}}::italic|grey|bgBlack and {{building}}::italic|grey|bgBlack
	//   {{dependency graphs}}::italic|grey|bgBlack for PHP, that allows you to find places in the code
	//   that can be improved.
	//
	//   It tries to be fast, {{~150k LOC/s}}::red|italic|bgBlack ({{lines of code per second}}::italic) on Core i5
	//   with SSD with ~3500Mb/s for reading.
	//
	//   This tool is written in Go and uses NoVerify::underline({{https://github.com/VKCOM/noverify}}::blue|underline).
	//
	//   Author::yellow: {{Petr Makhnev}}::red|bgBlack (tg: {{@petr_makhnev}}::blue|underline)
	//
	//   MIT (c) 2020
	//
	// `
	//
	// 	cfmt.Printf(code)
}
