// Package puppy is small
package puppy

import "fmt"

// Bark returns the sound "woof"
func Bark() string {
	return "woof!"
}

// Barks returns the sound "woof woof woof"
func Barks() string {
	return "woof woof woof" // test
}

// Tag2 prints a message indicating it's from tag 1.0.1
func Tag2() {
	fmt.Println("from tag 1.0.1")
}

// Tag3 prints a message indicating it's from tag 1.0.2
func Tag3() {
	fmt.Println("from tag 1.0.2")
}
