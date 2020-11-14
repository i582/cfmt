package cfmt

import (
	"fmt"
	"io"
	"os"

	"github.com/i582/cfmt/internal"
)

// RegisterStyle registers a new custom style.
//
// It will look like this:
//
// cfmt.RegisterStyle("code", func(s string) string {
//	return cfmt.Sprintf("{{%s}}::red|underline", s)
// })
//
// The first argument is the name by which this style will be used,
// and the second is the styling function itself.
func RegisterStyle(name string, fn func(string) string) {
	internal.Map[name] = func(f func(text string) string) func(text string) string {
		return func(text string) string {
			t := f(text)
			t = fn(t)
			return t
		}
	}
}

// Same as fmt.
func Sprint(a ...interface{}) string {
	text := fmt.Sprint(a...)
	parsed := internal.Parse(text)
	return parsed
}

// Same as fmt.
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	text := Sprint(a...)
	return fmt.Fprint(w, text)
}

// Same as fmt.
func Print(a ...interface{}) (n int, err error) {
	return Fprint(os.Stdout, a...)
}

// Same as fmt.
func Sprintln(a ...interface{}) string {
	return Sprint(a...) + "\n"
}

// Same as fmt.
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	text := Sprintln(a...)
	return fmt.Fprint(w, text)
}

// Same as fmt.
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}

// Same as fmt.
func Sprintf(format string, a ...interface{}) string {
	text := fmt.Sprintf(format, a...)
	parsed := internal.Parse(text)
	return parsed
}

// Same as fmt.
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	text := Sprintf(format, a...)
	return fmt.Fprint(w, text)
}

// Same as fmt.
func Printf(format string, a ...interface{}) (n int, err error) {
	text := Sprintf(format, a...)
	return fmt.Print(text)
}

// Same as fmt.
func Fatalf(format string, a ...interface{}) {
	text := Sprintf(format, a...)
	fmt.Print(text)
	os.Exit(1)
}

// Same as fmt.
func Fatal(a ...interface{}) {
	text := Sprint(a...)
	fmt.Print(text)
	os.Exit(1)
}
