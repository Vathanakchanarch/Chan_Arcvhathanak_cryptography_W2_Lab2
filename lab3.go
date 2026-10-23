package main

import "fmt"

func myXor(a, b int) int        { return a ^ b }
func myNOT(a, b int) (int, int) { 
	return ^a, ^b 
}
func myOR(a, b int) int         { return a | b }
func myAND(a, b int) int        { return a & b }

func myLeftShift(a, b int) int  { return a << b }
func myRightShift(a, b int) int { return a >> b }

func main() {
	var a, b int
	fmt.Print("Enter first integers: ")
	fmt.Scan(&a)
	
	fmt.Print("Enter second integers: ")
	fmt.Scan(&b)

	fmt.Println("\nBitwise Operations on two integers: ")
	fmt.Println("XOR: ", myXor(a, b))
	notA, notB := myNOT(a, b)
	fmt.Println("Not A: ", notA)
	fmt.Println("Not B: ", notB)

	fmt.Println("OR: ", myOR(a, b))
	fmt.Println("AND: ", myAND(a, b))

	fmt.Println("Left Shift: ", myLeftShift(a, b))
	fmt.Println("Right Shift: ", myRightShift(a, b))	

	
}
