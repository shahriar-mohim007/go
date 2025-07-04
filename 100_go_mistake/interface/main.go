package main

import "fmt"

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Logger interface {
	Read(p []byte) (n int, err error)
}

type File struct{}

func (f File) Read(p []byte) (int, error) {
	fmt.Println("Reading data...")
	return len(p), nil
}

func main() {

	var r Reader = File{}
	var l Logger = File{}

	r.Read([]byte("data")) // Treated as Reader
	l.Read([]byte("data")) // Treated as Logger
	fmt.Println(r == l)
	fmt.Println("Reader and Logger both satisfied.")
}
