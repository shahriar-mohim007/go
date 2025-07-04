package redis

import "fmt"

// imports
var a = func() int {
	fmt.Println("var redis")
	return 0
}()

func init() {
	// ...
	fmt.Println("redis")
}
func Store(key, value string) error {
	// ...
	fmt.Println("store", key, value)
	return fmt.Errorf("")
}
