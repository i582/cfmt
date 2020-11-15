package internal

import (
	"github.com/muesli/termenv"
)

var colorProfile termenv.Profile

func init() {
	colorProfile = termenv.ColorProfile()
}

func overline(text string) string {
	out := termenv.String(text)
	out = out.Overline()
	return out.String()
}

func reverse(text string) string {
	out := termenv.String(text)
	out = out.Reverse()
	return out.String()
}

func blink(text string) string {
	out := termenv.String(text)
	out = out.Blink()
	return out.String()
}

func faint(text string) string {
	out := termenv.String(text)
	out = out.Faint()
	return out.String()
}

func hexBackgroundColorFunc(cl string, text string) string {
	out := termenv.String(text)
	out = out.Background(colorProfile.Color(cl))
	return out.String()
}

func hexForegroundColorFunc(cl string, text string) string {
	out := termenv.String(text)
	out = out.Foreground(colorProfile.Color(cl))
	return out.String()
}
