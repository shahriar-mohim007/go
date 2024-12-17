package main

import "fmt"

type Foo struct {
	Field1 string
	Field2 int
}

func MakeFoo() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20,
	}
	return f, nil
}
func MakeFooBad(f *Foo) error {
	f.Field1 = "val1"
	f.Field2 = 21
	return nil
}
func main() {
	f, err := MakeFoo()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(f) // Prints: {val 20}

	g := Foo{}
	err = MakeFooBad(&g)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(g)
}
