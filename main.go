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
//   cfmt.RegisterStyle("code", func(s string) string {
//	  return cfmt.Sprintf("{{%s}}::red|underline", s)
//   })
//
// The first argument is the name by which this style will be used,
// and the second is the styling function itself.
func RegisterStyle(name string, fn func(string) string) {
	internal.CustomMap[name] = fn
}

// Sprint is the same as fmt.
func Sprint(a ...interface{}) string {
	text := fmt.Sprint(a...)
	return internal.ParseAndApply(text)
}

// Fprint is the same as fmt.
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	text := Sprint(a...)
	return fmt.Fprint(w, text)
}

// Print is the same as fmt.
func Print(a ...interface{}) (n int, err error) {
	return Fprint(os.Stdout, a...)
}

// Sprintln is the same as fmt.
func Sprintln(a ...interface{}) string {
	return Sprint(a...) + "\n"
}

// Fprintln is the same as fmt.
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	text := Sprintln(a...)
	return fmt.Fprint(w, text)
}

// Println is the same as fmt.
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}

// Sprintf is the same as fmt.
func Sprintf(format string, a ...interface{}) string {
	text := fmt.Sprintf(format, a...)
	return internal.ParseAndApply(text)
}

// Fprintf is the same as fmt.
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	text := Sprintf(format, a...)
	return fmt.Fprint(w, text)
}

// Printf is the same as fmt.
func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}

// Fatalf is the same as fmt.
func Fatalf(format string, a ...interface{}) {
	_, _ = Printf(format, a...)
	os.Exit(1)
}

// Fatal is the same as fmt.
func Fatal(a ...interface{}) {
	_, _ = Print(a...)
	os.Exit(1)
}

// Fatalln is the same as fmt.
func Fatalln(a ...interface{}) {
	_, _ = Println(a...)
	os.Exit(1)
}
