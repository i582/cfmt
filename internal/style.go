package internal

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

// StyleBuilder is a function that returns a callback function that will style
// the supplied string according to the supplied format string.
func StyleBuilder(styleString string) (func(string) string, error) {
	if styleString == "" {
		return nil, fmt.Errorf("style string is empty")
	}

	var styleParts = strings.Split(styleString, "|")

	var hexColorFuncs []func(string, func(string) string) func(string) string
	var hexColors []string

	var customStyles []func(func(string) string) func(string) string

	var colors = make([]color.Color, 0, len(styleParts))
	color.New()

	for _, style := range styleParts {
		if isHex(style) {
			err := checkHex(style)
			if err != nil {
				return nil, err
			}

			hexColorFuncs = append(hexColorFuncs, hexForegroundColor)
			hexColors = append(hexColors, style)
			continue
		}

		if isBackgroundHex(style) {
			rawHex := strings.TrimPrefix(style, "bg")
			err := checkHex(rawHex)
			if err != nil {
				return nil, err
			}

			hexColorFuncs = append(hexColorFuncs, hexBackgroundColor)
			hexColors = append(hexColors, rawHex)
			continue
		}

		clr, ok := Map[style]
		if !ok {
			customStyle, ok := CustomMap[style]
			if !ok {
				return nil, fmt.Errorf("unknown style %s", style)
			}

			customStyles = append(customStyles, customStyle)
			continue
		}
		colors = append(colors, clr)
	}

	outFun := func(text string) string { return text }

	for index, fun := range hexColorFuncs {
		clr := hexColors[index]
		outFun = fun(clr, outFun)
	}

	for _, style := range customStyles {
		outFun = style(outFun)
	}

	colorStyle := color.New(colors...)

	return func(s string) string {
		return outFun(colorStyle.Sprint(s))
	}, nil
}

func isHex(val string) bool {
	return strings.HasPrefix(val, "#")
}

func checkHex(val string) error {
	if len(val) != 7 {
		return fmt.Errorf("invalid hex: length of hex color must be 6")
	}

	return nil
}

func isBackgroundHex(val string) bool {
	return strings.HasPrefix(val, "bg#")
}
