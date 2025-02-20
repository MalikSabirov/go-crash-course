package main

import (
	"fmt"
	"go-crash-course/data_types"
	"go-crash-course/greetings"
)

func main() {
	greetings.SayHello()
	data_types.WorkWithPrimitives()

	var myAge data_types.Age = 30

	fmt.Println("Age category: ", data_types.Category(myAge))

	fmt.Println(greetings.GreetByYear(2029))
}
