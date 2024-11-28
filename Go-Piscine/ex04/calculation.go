package main

import (
	"fmt"
	"os"
	"strconv"
)

const ERROR_MSG string = "Arguments is invalid."

func main() {
	s, ok := calculationStr(os.Args)
	if !ok {
		fmt.Println(ERROR_MSG)
		os.Exit(1)
	}
	fmt.Print(s)
}

func calculationStr(args []string) (string, bool) {
	if len(args) != 3 {
		return "", false
	}
	num1, err1 := strconv.Atoi(args[1])
	num2, err2 := strconv.Atoi(args[2])
	if err1 != nil || err2 != nil {
		return "", false
	}
	sum := num1 + num2
	difference := abs(num1 - num2)
	product := num1 * num2
	quotient := num1 / num2
	result := fmt.Sprintf(
		"sum: %v\n"+
			"difference: %v\n"+
			"product: %v\n"+
			"quotient: %v\n",
		sum, difference, product, quotient,
	)
	return result, true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
