package internal

import (
	"github.com/gookit/color"
	"github.com/muesli/termenv"
)

func overline(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		out := termenv.String(t)
		out = out.Overline()
		t = out.String()
		return t
	}
}

func hexForegroundColor(cl string, f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		out := termenv.String(t)
		p := termenv.ColorProfile()
		out = out.Foreground(p.Color(cl))
		t = out.String()
		return t
	}
}

func hexBackgroundColor(cl string, f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		out := termenv.String(t)
		p := termenv.ColorProfile()
		out = out.Background(p.Color(cl))
		t = out.String()
		return t
	}
}

func colorTemplate(cl color.Color) func(f func(text string) string) func(text string) string {
	return func(f func(text string) string) func(text string) string {
		return func(text string) string {
			t := f(text)
			t = cl.Sprint(t)
			return t
		}
	}
}
