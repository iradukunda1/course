package main

import (
	"log"
)

// The entry point of the program
func main() {

	log.Println("Starting the application...")

	defer log.Println("... Stopping the application")

	a := []string{"a", "b", "c", "d"}
	b := []string{"a", "b", "e", "f"}

	add, remove := Diff(a, b)

	log.Printf("Add: %v", add)
	log.Printf("Remove: %v", remove)
}

// Diff returns the difference between two arrays
func Diff(a, b []string) (add []string, remove []string) {

	for _, v := range a {
		if !Contains(b, v) {
			add = append(add, v)
		}
	}

	for _, v := range b {
		if !Contains(a, v) {
			remove = append(remove, v)
		}
	}

	return
}

// Contains returns true if the array contains provided value
func Contains(a []string, v string) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}
