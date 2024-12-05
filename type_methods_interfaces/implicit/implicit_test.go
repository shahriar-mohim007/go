package main

import (
	"fmt"
	"testing"
)

// Notifier interface defines the behavior for sending notifications.
type Notifier interface {
	Notify(message string) error
}

// UserService depends on a Notifier to send welcome emails.
type UserService struct {
	notifier Notifier
}

// SendWelcomeEmail uses the notifier to send a welcome message to the user.
func (u *UserService) SendWelcomeEmail(user string) error {
	message := "Welcome, " + user + "!"
	return u.notifier.Notify(message)
}

// EmailNotifier is a concrete implementation of Notifier for sending emails.
type EmailNotifier struct{}

// Notify sends an email notification.
func (e EmailNotifier) Notify(message string) error {
	fmt.Println("Sending email:", message)
	return nil // Simulate successful email sending
}

// MockNotifier is a mock implementation of Notifier for testing purposes.
type MockNotifier struct{}

// Notify simulates sending a notification without any real action.
func (m MockNotifier) Notify(message string) error {
	fmt.Println("Mock notify called with message:", message)
	return nil // Simulate success
}

// Main function demonstrating dependency injection in production.
func main() {
	// Create a real email notifier
	emailNotifier := EmailNotifier{}

	// Inject the real notifier into the UserService
	service := UserService{notifier: emailNotifier}

	// Call the service method
	err := service.SendWelcomeEmail("Alice")
	if err != nil {
		fmt.Println("Error:", err)
	}

}

//Dependency Injection in main:
//
//The UserService is injected with the EmailNotifier to send real notifications.

// Test function demonstrating dependency injection with a mock implementation.
func TestSendWelcomeEmail(t *testing.T) {
	// Create a mock notifier
	mockNotifier := MockNotifier{}

	// Inject the mock notifier into the UserService
	service := UserService{notifier: mockNotifier}

	// Call the service method
	err := service.SendWelcomeEmail("Bob")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

//In Go, implicit interfaces simplify dependency injection by allowing you to define dependencies
//in terms of the behavior (methods) required, rather than the concrete implementation.
//	This eliminates the need for tight coupling between components
//and their dependencies, making applications more flexible, testable, and maintainable.

//How Implicit Interfaces Aid Dependency Injection
//
//Define Dependencies by Behavior:
//Instead of requiring a specific implementation, you specify the behavior (methods) your component depends on.
//
//Inject Any Implementation:
//You can provide any implementation (mock, stub, or real) that satisfies the interface, making it easy to swap dependencies.
//
//Avoid Coupling:
//Since the component depends only on an interface, it avoids being tied to specific concrete implementations.

//Benefits of Implicit Interfaces for Dependency Injection
//
//Loose Coupling:
//Code depends on abstractions (interfaces) rather than concrete implementations.
//
//Flexibility:
//Any type that satisfies the interface can be injected.
//
//Improved Testability:
//Mock implementations can be injected for unit testing without modifying the actual code.
//
//Ease of Maintenance:
//Swapping out dependencies becomes trivial since they only need to satisfy the interface.

//Why This Is Powerful
//
//Go’s implicit interfaces align well with the principles of dependency inversion and interface
//segregation from SOLID design principles. They enable you to:
//
//Focus on defining behaviors instead of implementations.
//Keep your code modular and extensible.
//
//By leveraging implicit interfaces, dependency injection in Go becomes both intuitive and natural,
//leading to cleaner and more maintainable applications.
