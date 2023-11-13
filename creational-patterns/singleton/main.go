package main

import (
	"fmt"

	so "github.com/lexyu-golang-project-collection/creational-patterns/singleton/sync-once"
)

func main() {

	/*
		for i := 0; i < 10; i++ {
			so.GetInstance()
		}
	*/

	// /*
	for i := 0; i < 100; i++ {
		go func() {
			so.GetInstanceRC()
		}()

	}
	// */

	fmt.Scanln()
}
