package main

import (
	"fmt"
	"time"

	mx "github.com/lexyu-golang-project-collection/creational-patterns/singleton/mutex"
	// so "github.com/lexyu-golang-project-collection/creational-patterns/singleton/sync-once"
)

func main() {

	/*
		for i := 0; i < 10; i++ {
			go so.GetInstance()
		}
	*/

	/*
		for i := 0; i < 100; i++ {
			go func() {
				so.GetInstanceRC()
			}()

		}
	*/

	// /*
	for i := 0; i < 10; i++ {
		go mx.GetInstance()
	}
	// */

	time.Sleep(2 * time.Second)

	var input string
	fmt.Println("Press Enter to exit.")
	fmt.Scanln(&input)
}
