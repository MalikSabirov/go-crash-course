package data_types

import (
	"fmt"
	"strings"
)

func WorkWithPrimitives() {
	var str string = "Hello, Go!"

	fmt.Println("Length of str", len(str))
	fmt.Println("Upper case", strings.ToUpper(str))
	fmt.Println("'Go' in str", strings.Contains(str, "Go"))

	var age int = 25
	fmt.Printf("age: %d \n", age)

	var pi float64 = 3.1415
	fmt.Printf("number Pi: %.2f\n", pi)
}
