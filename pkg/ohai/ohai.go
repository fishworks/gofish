package ohai

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

// Ohai displays an informative message.
func Ohai(a ...interface{}) (int, error) {
	return Ohaif("%s", a...)
}

// Ohaif displays an informative message.
func Ohaif(format string, a ...interface{}) (int, error) {
	return fmt.Printf(fmt.Sprintf("==> %s", format), a...)
}

// Ohailn displays an informative message.
func Ohailn(a ...interface{}) (int, error) {
	return Ohaif("%s\n", a...)
}

// Success displays a success message.
func Success(a ...interface{}) (int, error) {
	return Successf("%s", a...)
}

// Successf displays a success message.
func Successf(format string, a ...interface{}) (int, error) {
	return fmt.Printf(emoji.Sprintf(":tropical_fish: %s", format), a...)
}

// Successln displays a success message.
func Successln(a ...interface{}) (int, error) {
	return Successf("%s\n", a...)
}

// Warning displays a warning message.
func Warning(a ...interface{}) (int, error) {
	return Warningf("%s", a...)
}

// Warningf displays a warning message.
func Warningf(format string, a ...interface{}) (int, error) {
	return fmt.Printf(emoji.Sprintf("!!! %s", format), a...)
}

// Warningln displays a warning message.
func Warningln(a ...interface{}) (int, error) {
	return Warningf("%s\n", a...)
}
