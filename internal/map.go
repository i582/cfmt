package internal

type styleCallback func(func(string) string) func(string) string

// Map is the styles map
var Map = map[string]styleCallback{
	"red":          Red,
	"cyan":         Cyan,
	"gray":         Gray,
	"blue":         Blue,
	"black":        Black,
	"green":        Green,
	"white":        White,
	"yellow":       Yellow,
	"magenta":      Magenta,
	"lightRed":     LightRed,
	"lightCyan":    LightCyan,
	"lightBlue":    LightBlue,
	"lightGreen":   LightGreen,
	"lightWhite":   LightWhite,
	"lightYellow":  LightYellow,
	"lightMagenta": LightMagenta,

	"bgRed":          BgRed,
	"bgGray":         BgGray,
	"bgCyan":         BgCyan,
	"bgBlue":         BgBlue,
	"bgWhite":        BgWhite,
	"bgBlack":        BgBlack,
	"bgGreen":        BgGreen,
	"bgYellow":       BgYellow,
	"bgMagenta":      BgMagenta,
	"bgDefault":      BgDefault,
	"bgDarkGray":     BgDarkGray,
	"bgLightRed":     BgLightRed,
	"bgLightCyan":    BgLightCyan,
	"bgLightBlue":    BgLightBlue,
	"bgLightWhite":   BgLightWhite,
	"bgLightGreen":   BgLightGreen,
	"bgLightYellow":  BgLightYellow,
	"bgLightMagenta": BgLightMagenta,

	"bold":      Bold,
	"italic":    Italic,
	"crossout":  CrossOut,
	"underline": Underline,
	"overline":  Overline,
}
