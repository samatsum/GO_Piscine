package main

import "fmt"

func main() {
	var stringValue string = "42"
	var uintValue uint = 42
	var intValue int = 42
	var uint8Value uint8 = 42
	var int16Value int16 = 42
	var uint32Value uint32 = 42
	var int64Value int64 = 42
	var boolValue bool = false
	var float32Value float32 = 42
	var float64Value float64 = 42
	var complex64Value complex64 = 42 + 0i
	var complex128Value complex128 = 42 + 21i
	var structValue = struct{}{}
	var arrayValue = [1]int{42}
	var mapValue = map[string]int{"42": 42}
	var pointerValue *int
	var sliceValue = []int{}
	var chanValue chan int

	fmt.Printf("%v is a string.\n", stringValue)
	fmt.Printf("%v is a uint.\n", uintValue)
	fmt.Printf("%v is an int.\n", intValue)
	fmt.Printf("%v is a uint8.\n", uint8Value)
	fmt.Printf("%v is an int16.\n", int16Value)
	fmt.Printf("%v is a uint32.\n", uint32Value)
	fmt.Printf("%v is an int64.\n", int64Value)
	fmt.Printf("%v is a bool.\n", boolValue)
	fmt.Printf("%v is a float32.\n", float32Value)
	fmt.Printf("%v is a float64.\n", float64Value)
	fmt.Printf("%v is a complex64.\n", complex64Value)
	fmt.Printf("%v is a complex128.\n", complex128Value)
	fmt.Printf("%v is a main.FortyTwo.\n", structValue)
	fmt.Printf("%v is a [1]int.\n", arrayValue)
	fmt.Printf("%v is a map[string]int.\n", mapValue)
	fmt.Printf("%v is an *int.\n", pointerValue)
	fmt.Printf("%v is a []int.\n", sliceValue)
	fmt.Printf("%v is a chan int.\n", chanValue)
}
