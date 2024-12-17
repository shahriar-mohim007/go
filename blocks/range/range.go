package main

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
	for _, v := range evenVals {
		fmt.Println(v)
	}
	uniqueNames := map[string]bool{
		"Fred":  true,
		"Raul":  true,
		"Wilma": true,
	}
	for k := range uniqueNames {
		fmt.Println(k)
	}

	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals)
	for i := range evenVals {
		evenVals[i] *= 2
	}
	fmt.Println(evenVals)
}

//The for-range value is a copy
//You should be aware that each time the for-range loop iterates over your
//compound type, it copies the value from the compound type to the value
//variable. Modifying the value variable will not modify the value in the
//compound type.
