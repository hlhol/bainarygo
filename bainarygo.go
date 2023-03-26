package main

import (
	"fmt"
)

//create a method called decimalToBinary

func decimalToBinary(decimal int) string {
	// the decimal 0 the bainary of it is 0
	if decimal == 0 {
		return "0"
	}

	binary := ""
	//start for loop to convert
	for decimal > 0 {
		//we can get reminder by dvide by 2
		remainder := decimal % 2
		binary = fmt.Sprintf("%d%s", remainder, binary)
		decimal = decimal / 2
	}
	return binary
}

func main() {
	decimal := 44
	binary := decimalToBinary(decimal)
	fmt.Printf("Decimal %d = Binary %s\n", decimal, binary)
}
