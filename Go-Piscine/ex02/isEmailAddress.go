package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Invalid argument")
		return
	}
	emailRegex := `^[\w\-\._]+@[\w\-\._]+\.[A-Za-z]+$`
	re := regexp.MustCompile(emailRegex)
	for _, arg := range args {
		if re.MatchString(arg) {
			fmt.Printf("%v is a valid email address.\n", arg)
		} else {
			fmt.Printf("%v is NOT a valid email address.\n", arg)
		}
	}
}
