package main

import "fmt"

// Printer interface defines the behavior of printing.
type Printer interface {
	Print(content string)
}

// Scanner interface defines the behavior of scanning.
type Scanner interface {
	Scan() string
}

// PrintScanner interface embeds Printer and Scanner, combining their behaviors.
type PrintScanner interface {
	Printer
	Scanner
}

// Device struct implements both Printer and Scanner interfaces.
type Device struct{}

// Implement the Print method for Printer interface.
func (d Device) Print(content string) {
	fmt.Println("Printing:", content)
}

// Implement the Scan method for Scanner interface.
func (d Device) Scan() string {
	return "Scanned content"
}

func main() {
	// Create an instance of Device
	var ps PrintScanner = Device{}

	// Use the Print method from Printer interface
	ps.Print("Hello, World!")

	// Use the Scan method from Scanner interface
	scanned := ps.Scan()
	fmt.Println(scanned)
}
