package styles

import (
	"fmt"
	"strings"
)

func StyleBuilder(styleString string) (func(string) string, error) {
	if styleString == "" {
		return nil, fmt.Errorf("style string is empty")
	}

	styleParts := strings.Split(styleString, "|")

	var funcs []StyleCallback
	var hexColorFuncs []func(string, func(string) string) func(string) string
	var hexColors []string

	for _, style := range styleParts {
		if isHex(style) {
			err := checkHex(style)
			if err != nil {
				return nil, err
			}

			hexColorFuncs = append(hexColorFuncs, HexForegroundColor)
			hexColors = append(hexColors, style)
			continue
		}

		if isBackgroundHex(style) {
			err := checkHex(strings.TrimPrefix(style, "bg"))
			if err != nil {
				return nil, err
			}

			hexColorFuncs = append(hexColorFuncs, HexBackgroundColor)
			hexColors = append(hexColors, strings.TrimPrefix(style, "bg"))
			continue
		}

		fn, ok := Map[style]
		if !ok {
			return nil, fmt.Errorf("unknown style %s", style)
		}

		funcs = append(funcs, fn)
	}

	outFun := func(text string) string { return text }

	for _, fun := range funcs {
		outFun = fun(outFun)
	}

	for index, fun := range hexColorFuncs {
		var color string
		if index >= 0 && index < len(hexColors) {
			color = hexColors[index]
		}

		outFun = fun(color, outFun)
	}

	return outFun, nil
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
