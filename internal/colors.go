package internal

import (
	"github.com/gookit/color"
	"github.com/muesli/termenv"
)

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

// Provides a color.
func Red(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.Red.Sprint(t)
		return t
	}
}

// Provides a color.
func Cyan(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.Cyan.Sprint(t)
		return t
	}
}

// Provides a color.
func Gray(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.Gray.Sprint(t)
		return t
	}
}

// Provides a color.
func Blue(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.Blue.Sprint(t)
		return t
	}
}

// Provides a color.
func Black(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.Black.Sprint(t)
		return t
	}
}

// Provides a color.
func Green(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.Green.Sprint(t)
		return t
	}
}

// Provides a color.
func White(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.White.Sprint(t)
		return t
	}
}

// Provides a color.
func Yellow(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.Yellow.Sprint(t)
		return t
	}
}

// Provides a color.
func Magenta(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.Magenta.Sprint(t)
		return t
	}
}

// Provides a color.
func LightRed(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.LightRed.Sprint(t)
		return t
	}
}

// Provides a color.
func LightCyan(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.LightCyan.Sprint(t)
		return t
	}
}

// Provides a color.
func LightBlue(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.LightBlue.Sprint(t)
		return t
	}
}

// Provides a color.
func LightGreen(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.LightGreen.Sprint(t)
		return t
	}
}

// Provides a color.
func LightWhite(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.LightWhite.Sprint(t)
		return t
	}
}

// Provides a color.
func LightYellow(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.LightYellow.Sprint(t)
		return t
	}
}

// Provides a color.
func LightMagenta(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.LightMagenta.Sprint(t)
		return t
	}
}

// Provides a color.
func BgBlack(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.BgBlack.Sprint(t)
		return t
	}
}

// Provides a color.
func BgRed(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.BgRed.Sprint(t)
		return t
	}
}

// Provides a color.
func BgGreen(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.BgGreen.Sprint(t)
		return t
	}
}

// Provides a color.
func BgYellow(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.BgYellow.Sprint(t)
		return t
	}
}

// Provides a color.
func BgBlue(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.BgBlue.Sprint(t)
		return t
	}
}

// Provides a color.
func BgMagenta(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.BgMagenta.Sprint(t)
		return t
	}
}

// Provides a color.
func BgWhite(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.BgWhite.Sprint(t)
		return t
	}
}

// Provides a color.
func BgCyan(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.BgCyan.Sprint(t)
		return t
	}
}

// Provides a color.
func BgDefault(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.BgDefault.Sprint(t)
		return t
	}
}

// Provides a color.
func BgDarkGray(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.BgDarkGray.Sprint(t)
		return t
	}
}

// Provides a color.
func BgLightRed(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.BgLightRed.Sprint(t)
		return t
	}
}

// Provides a color.
func BgLightGreen(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.BgLightGreen.Sprint(t)
		return t
	}
}

// Provides a color.
func BgLightYellow(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.BgLightYellow.Sprint(t)
		return t
	}
}

// Provides a color.
func BgLightBlue(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.BgLightBlue.Sprint(t)
		return t
	}
}

// Provides a color.
func BgLightMagenta(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.BgLightMagenta.Sprint(t)
		return t
	}
}

// Provides a color.
func BgLightCyan(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.BgLightCyan.Sprint(t)
		return t
	}
}

// Provides a color.
func BgLightWhite(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.BgLightWhite.Sprint(t)
		return t
	}
}

// Provides a color.
func BgGray(f func(text string) string) func(text string) string {
	return func(text string) string {
		t := f(text)
		t = color.BgGray.Sprint(t)
		return t
	}
}
