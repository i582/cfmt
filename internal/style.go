package internal

import (
	"fmt"
	"strings"
)

// StyleBuilder is a function that returns a styled string based on the supplied style.
func StyleBuilder(styles []string, text string) (string, error) {
	if len(styles) == 0 {
		return "", fmt.Errorf("style string is empty")
	}

	var err error
	for _, style := range styles {
		text, err = ApplyStyle(text, style)
		if err != nil {
			return "", err
		}
	}

	return text, nil
}

// StyleBuilder is a function that apply a style for string.
func ApplyStyle(text, style string) (string, error) {
	if isHex(style) {
		err := checkHex(style)
		if err != nil {
			return "", err
		}
		text = hexForegroundColorFunc(style, text)
		return text, nil
	}

	if isBackgroundHex(style) {
		rawHex := strings.TrimPrefix(style, "bg")
		err := checkHex(rawHex)
		if err != nil {
			return "", err
		}
		text = hexBackgroundColorFunc(style, text)
		return text, nil
	}

	clr, ok := Map[style]
	if !ok {
		customStyleFunc, ok := CustomMap[style]
		if !ok {
			return "", fmt.Errorf("unknown style %s", style)
		}
		text = customStyleFunc(text)
		return text, nil
	}

	return clr.Sprint(text), nil
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
