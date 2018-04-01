package ohai

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func Ohai(a ...interface{}) (int, error) {
	return Ohaif("%s", a...)
}

func Ohaif(format string, a ...interface{}) (int, error) {
	return fmt.Printf(fmt.Sprintf("==> %s", format), a...)
}

func Ohailn(a ...interface{}) (int, error) {
	return Ohaif("%s\n", a...)
}

func Success(a ...interface{}) (int, error) {
	return Successf("%s", a...)
}

func Successf(format string, a ...interface{}) (int, error) {
	return fmt.Printf(emoji.Sprintf(":tropical_fish: %s", format), a...)
}

func Successln(a ...interface{}) (int, error) {
	return Successf("%s\n", a...)
}
