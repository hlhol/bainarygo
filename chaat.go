package main

import (
	"fmt"
)

var pr = fmt.Println
var num1 = 1
var num2 = 55

func main() {

	for num1 <= 15 {
		pr(num1)
		num1 = num1 + 1

	}

	if num2 >= 15 {
		pr("ooh right, we are fainally here")

	} else {
		pr("oooh nooooo")
	}

}
