package internal

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

// StyleBuilder is a function that returns a styled string based on the supplied style.
func StyleBuilder(styleString string, text string) (string, error) {
	if styleString == "" {
		return "", fmt.Errorf("style string is empty")
	}

	var styleParts = strings.Split(styleString, "|")
	var colors = make([]color.Color, 0, len(styleParts))
	color.New()

	for _, style := range styleParts {
		if isHex(style) {
			err := checkHex(style)
			if err != nil {
				return "", err
			}
			text = hexForegroundColorFunc(style, text)
			continue
		}

		if isBackgroundHex(style) {
			rawHex := strings.TrimPrefix(style, "bg")
			err := checkHex(rawHex)
			if err != nil {
				return "", err
			}
			text = hexBackgroundColorFunc(style, text)
			continue
		}

		clr, ok := Map[style]
		if !ok {
			customStyleFunc, ok := CustomMap[style]
			if !ok {
				return "", fmt.Errorf("unknown style %s", style)
			}
			text = customStyleFunc(text)
			continue
		}
		colors = append(colors, clr)
	}

	colorStyle := color.New(colors...)
	return colorStyle.Sprint(text), nil
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
