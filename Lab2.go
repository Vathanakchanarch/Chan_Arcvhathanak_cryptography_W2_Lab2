package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter first integers: ")
	fmt.Scan(&a)
	
	fmt.Print("Enter second integers: ")
	fmt.Scan(&b)
	
	fmt.Println("Both numbers are positive:", a>0 && b>0)

	fmt.Println("One number is greater than other:", a>b || a<b)

	fmt.Println("numbers are not equal:", !(a==b))

}
