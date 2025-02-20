package main

import (
	"fmt"
	"go-crash-course/data_types"
	"go-crash-course/greetings"
)

var Global = 5

func ModifyGlobal() {
	fmt.Println("\nBefore modify: ", Global)

	original := Global

	defer func() {
		Global = original
		fmt.Println("After defer restore initial value: ", Global)
	}()

	Global = 10
	fmt.Println("After modify: ", Global)
}

func main() {
	greetings.SayHello()
	data_types.WorkWithPrimitives()

	var myAge data_types.Age = 30

	fmt.Println("Age category: ", data_types.Category(myAge))

	fmt.Println(greetings.GreetByYear(2029))

	ModifyGlobal()
	println("in main after call ModifyGlobal: ", Global)
}
