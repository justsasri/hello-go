package utils

import (
	"fmt"
)

func Greet(s string) {
	res := fmt.Sprintf("%s said hello!", s)
	fmt.Println(res)
}
