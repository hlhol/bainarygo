package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	get := bufio.NewScanner(os.Stdin)
	fmt.Println("input text:")
	get.Scan()
	input := get.Text()

	fmt.Printf("This is my first scanner in go language %v ", input)

}
