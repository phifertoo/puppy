// Package puppy is small
package puppy

import "fmt"

// Bark says woof
func Bark() string {
	return "woof!"
}

func Barks() string {
	return "woof woof woof" // test
}

func Tag2() {
	fmt.Println("from tag 1.0.1")
}

func Tag3() {
	fmt.Println("from tag 1.0.2")
}
