package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	maxAstarisk, err := strconv.Atoi(os.Args[1])
	if err != nil || maxAstarisk < 0 || maxAstarisk > 10000 {
		return
	}
	totalAstarisk := 0
	for i := 1; totalAstarisk+i <= maxAstarisk; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
		totalAstarisk += i
	}
}
