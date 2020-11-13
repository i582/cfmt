package cfmt

import (
	"fmt"
	"io"
	"os"

	"github.com/i582/cfmt/internal/parser"
)

func Sprint(a ...interface{}) string {
	text := fmt.Sprint(a...)
	p := parser.Parser{}
	parsed := p.Parse(text)
	return parsed
}

func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	text := Sprint(a...)
	return fmt.Fprint(w, text)
}

func Print(a ...interface{}) (n int, err error) {
	return Fprint(os.Stdout, a...)
}

func Sprintln(a ...interface{}) string {
	return Sprint(a...) + "\n"
}

func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	text := Sprintln(a...)
	return fmt.Fprint(w, text)
}

func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}

func Sprintf(format string, a ...interface{}) string {
	text := fmt.Sprintf(format, a...)
	p := parser.Parser{}
	parsed := p.Parse(text)
	return parsed
}

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	text := Sprintf(format, a...)
	return fmt.Fprint(w, text)
}

func Printf(format string, a ...interface{}) (n int, err error) {
	text := Sprintf(format, a...)
	return fmt.Print(text)
}

func Fatalf(format string, a ...interface{}) {
	text := Sprintf(format, a...)
	fmt.Print(text)
	os.Exit(1)
}

func Fatal(a ...interface{}) {
	text := Sprint(a...)
	fmt.Print(text)
	os.Exit(1)
}
