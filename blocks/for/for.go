package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("Current number:", i)

		// Break the loop when i equals 5
		if i == 5 {
			fmt.Println("Breaking the loop")
			break
		}
	}

	fmt.Println("Loop exited")
}
