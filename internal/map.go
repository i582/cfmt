package internal

import (
	"github.com/gookit/color"
)

type styleCallback func(func(string) string) func(string) string

// Map is the styles map
var Map = map[string]styleCallback{
	"red":          colorTemplate(color.Red),
	"cyan":         colorTemplate(color.Cyan),
	"gray":         colorTemplate(color.Gray),
	"blue":         colorTemplate(color.Blue),
	"black":        colorTemplate(color.Black),
	"green":        colorTemplate(color.Green),
	"white":        colorTemplate(color.White),
	"yellow":       colorTemplate(color.Yellow),
	"magenta":      colorTemplate(color.Magenta),
	"lightRed":     colorTemplate(color.LightRed),
	"lightCyan":    colorTemplate(color.LightCyan),
	"lightBlue":    colorTemplate(color.LightBlue),
	"lightGreen":   colorTemplate(color.LightGreen),
	"lightWhite":   colorTemplate(color.LightWhite),
	"lightYellow":  colorTemplate(color.LightYellow),
	"lightMagenta": colorTemplate(color.LightMagenta),

	"bgRed":          colorTemplate(color.BgRed),
	"bgGray":         colorTemplate(color.BgGray),
	"bgCyan":         colorTemplate(color.BgCyan),
	"bgBlue":         colorTemplate(color.BgBlue),
	"bgWhite":        colorTemplate(color.BgWhite),
	"bgBlack":        colorTemplate(color.BgBlack),
	"bgGreen":        colorTemplate(color.BgGreen),
	"bgYellow":       colorTemplate(color.BgYellow),
	"bgMagenta":      colorTemplate(color.BgMagenta),
	"bgDefault":      colorTemplate(color.BgDefault),
	"bgDarkGray":     colorTemplate(color.BgDarkGray),
	"bgLightRed":     colorTemplate(color.BgLightRed),
	"bgLightCyan":    colorTemplate(color.BgLightCyan),
	"bgLightBlue":    colorTemplate(color.BgLightBlue),
	"bgLightWhite":   colorTemplate(color.BgLightWhite),
	"bgLightGreen":   colorTemplate(color.BgLightGreen),
	"bgLightYellow":  colorTemplate(color.BgLightYellow),
	"bgLightMagenta": colorTemplate(color.BgLightMagenta),

	"bold":      colorTemplate(color.Bold),
	"italic":    colorTemplate(color.OpItalic),
	"crossout":  colorTemplate(color.OpStrikethrough),
	"underline": colorTemplate(color.OpUnderscore),
	"overline":  overline,
}
