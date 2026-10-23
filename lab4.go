package main

import "fmt"



func main() {
	
	var a, b,choices int

	fmt.Print("Enter first integers: ")
	fmt.Scan(&a)

	fmt.Print("Enter second integers: ")
	fmt.Scan(&b)
	for {
		fmt.Println("Choose the operation you want to perform: ")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Modulus")
		fmt.Println("6. Exit")
		fmt.Scan(&choices)
		if choices == 6 {
			break
		}
		switch choices {
		case 1:
			fmt.Println("Addition: ", a+b)
		case 2:
			fmt.Println("Subtraction: ", a-b)
		case 3:
			fmt.Println("Multiplication: ", a*b)
		case 4:
			if b == 0 {
				fmt.Println("Division by zero is not allowed.")
			} else {
				fmt.Println("Division: ", a/b)
			}
		case 5:
			if b == 0 {
				fmt.Println("Modulus by zero is not allowed.")
			} else {
				fmt.Println("Modulus: ", a%b)
			}
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}


}