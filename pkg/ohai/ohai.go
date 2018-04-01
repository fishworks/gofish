package ohai

import (
	"fmt"
)

func Ohai(a ...interface{}) (int, error) {
	a = append([]interface{}{"==> "}, a...)
	return fmt.Println(a...)
}

func Ohaif(format string, a ...interface{}) (int, error) {
	return fmt.Printf(fmt.Sprintf("==> %s", format), a...)
}

func Ohailn(a ...interface{}) (int, error) {
	a = append([]interface{}{"==> "}, a...)
	return fmt.Println(a...)
}

func Success(a ...interface{}) (int, error) {
	a = append([]interface{}{"🐟  "}, a...)
	return fmt.Print(a...)
}

func Successf(format string, a ...interface{}) (int, error) {
	return fmt.Printf(fmt.Sprintf("🐟  %s", format), a...)
}

func Successln(a ...interface{}) (int, error) {
	a = append([]interface{}{"🐟  "}, a...)
	return fmt.Println(a...)
}
