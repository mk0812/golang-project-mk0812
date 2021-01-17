package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

//import "fmt"
//import "math" の様に個別でもimportが可能

var lang string

var lang2 string = "Go"

var slice1 = []string{"Golang", "Ruby"}

func main() {
	const GREET_MEET = "Hello World"
	sum := 0
	defer fmt.Println(hello2("Hello", GREET_MEET))
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println(condition("Ruby"))

	var mike Person
	mike.name = "Mike"
	mike.age = 23

	fmt.Println(mike)

	for index, value := range slice1 {
		fmt.Println(index, value)
	}
}

func hello2(arg1, arg2 string) string {
	return arg1 + " " + arg2
}

func condition(arg string) string {
	switch arg {
	case "Ruby":
		return "This is Ruby"
	case "Go": //これ以降のcaseやdefaultは実行されない。
		return "This is Go"
	case "JS":
		return "This is JS"
	default:
		return "I don't know what this is"
	}
}
