package styles

import (
	"github.com/gookit/color"
	"github.com/muesli/termenv"
)

func Bold(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.Bold.Sprint(t)
		return t
	}
}

func Italic(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.OpItalic.Sprint(t)
		return t
	}
}

func CrossOut(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.OpStrikethrough.Sprint(t)
		return t
	}
}

func Underline(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.OpUnderscore.Sprint(t)
		return t
	}
}

func Overline(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		out := termenv.String(t)
		out = out.Overline()
		t = out.String()
		return t
	}
}
