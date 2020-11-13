package cfmt

import (
	"github.com/i582/cfmt/internal/styles"
)

func RegisterStyle(name string, fn func(string) string) {
	styles.Map[name] = func(f func(text string) string) func(text string) string {
		return func(text string) string {
			t := f(text)
			t = fn(t)
			return t
		}
	}
}
