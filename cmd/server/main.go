// Package main is the entry into the application
package main

import "fmt"

// Run - is going to be reponsible for
// the instantiation and startup of our
// go application
func Run() error {
	fmt.Println("Starting up our application")
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
