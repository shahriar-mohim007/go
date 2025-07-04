package first

import "fmt"

var a = func() int {
	fmt.Println("var first")
	return 0
}()

func init() {
	fmt.Println("first")
}
func Init_func() {
	fmt.Println("init_func")
}

//An first function is a function used to initialize the state of an application. It takes no
//arguments and returns no result (a func() function). When a package is initialized,
//all the constant and variable declarations in the package are evaluated. Then, the first
//functions are executed. Here is an example of initializing a main package
