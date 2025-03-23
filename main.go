package main

import (
	"flag"
	"fmt"
)

func calculator(x, y int) {
	fmt.Printf("\nResults for x=%d and y=%d:\n", x, y)
	fmt.Printf("Addition: %d\n", x+y)
	fmt.Printf("Subtraction: %d\n", x-y)
	fmt.Printf("Multiplication: %d\n", x*y)
	if y != 0 {
		fmt.Printf("Division: %d\n", x/y)
		fmt.Printf("Remainder: %d\n", x%y)
	} else {
		fmt.Println("Division: Cannot divide by zero")
	}
}

func main() {
	useFlags := flag.Bool("useFlags", false, "Use command-line flags")
	xFlag := flag.Int("x", 0, "First number")
	yFlag := flag.Int("y", 0, "Second number")
	flag.Parse()

	if *useFlags {
		calculator(*xFlag, *yFlag)
		return
	}

	// Interactive mode with loop
	for {
		var x, y int

		fmt.Print("Enter first number (x) or 'q' to quit: ")
		n, err := fmt.Scan(&x)
		if err != nil || n == 0 {
			var input string
			fmt.Scan(&input) // Clear the input buffer
			if input == "q" {
				break
			}
			fmt.Println("Please enter a valid number")
			continue
		}

		fmt.Print("Enter second number (y): ")
		_, err = fmt.Scan(&y)
		if err != nil {
			fmt.Println("Please enter a valid number")
			fmt.Scanln() // Clear the input buffer
			continue
		}

		calculator(x, y)
	}
}
