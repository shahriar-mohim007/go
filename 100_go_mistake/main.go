package main

import (
	_ "again_go/100_go_mistake/first"
	"again_go/100_go_mistake/redis"
	"fmt"
)

func init() {
	// ...
	fmt.Println("main first")
}
func main() {

	err := redis.Store("foo", "bar")
	if err != nil {
		fmt.Println(err)
	}
	// ...
}
