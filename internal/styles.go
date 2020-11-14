package internal

import (
	"github.com/gookit/color"
	"github.com/muesli/termenv"
)

// Bold provides a bold style.
func Bold(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.Bold.Sprint(t)
		return t
	}
}

// Italic provides a italic style.
func Italic(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.OpItalic.Sprint(t)
		return t
	}
}

// CrossOut provides a crossout style.
func CrossOut(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.OpStrikethrough.Sprint(t)
		return t
	}
}

// Underline provides a underline style.
func Underline(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.OpUnderscore.Sprint(t)
		return t
	}
}

// Overline provides a overline style.
func Overline(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		out := termenv.String(t)
		out = out.Overline()
		t = out.String()
		return t
	}
}
