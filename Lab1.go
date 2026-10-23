package main

import "fmt"


// func main() {
// 	fmt.Println("Hello World!")
// }
func main() {
	var a, b int
	var result int 

	fmt.Print("Enter first integers: ")

	fmt.Scan(&a)

	fmt.Print("Enter second integers: ")
	fmt.Scan(&b)

	fmt.Println("\nOperations on two integers: ")
	result = b

	fmt.Println("= ", result)

	result = a + b

	fmt.Println("+= ", result)

	result = a - b

	fmt.Println("-= ", result)

	result = a * b
	
	fmt.Println("*= ", result)

	result = a / b
	
	fmt.Println("/= ", result)

	result = a % b
	
	fmt.Println("%= ", result)


}