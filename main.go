package main

import (
	"calcolatrice/math"
	"flag"
	"fmt"
	"os"
)

func main() {
	add := flag.Bool("add", false, "Add two numbers")
	subs := flag.Bool("subtract", false, "Subtract two numbers")
	mult := flag.Bool("multiply", false, "Multiply two numbers")
	div := flag.Bool("divide", false, "Divide two numbers")

	flag.Parse()

	var a, b int
	fmt.Println("Enter 1st Number: ")
	fmt.Scan(&a)
	fmt.Println("Enter 2nd Number: ")
	fmt.Scan(&b)

	switch {
	case *add:
		fmt.Printf("Additon: %d \n", math.Addition(a, b))
	case *subs:
		fmt.Printf("Difference: %d \n", math.Subtraction(a, b))
	case *mult:
		fmt.Printf("Product: %d \n", math.Multiplication(a, b))
	case *div:
		fmt.Printf("Division: %f \n", math.Division(float64(a), float64(b)))
	default:
		fmt.Fprintln(os.Stderr, "Wrong option try with add, subtract, div and multply")
		os.Exit(1)
	}
}
