package main

import (
	"fmt"
	"sync"
	"time"
)

type single struct {
}

var (
	once           sync.Once
	singleInstance *single
)

func GetInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Printf("Once - Creating single instance addr = %p\n", &singleInstance)
				singleInstance = &single{}
			})
	} else {
		fmt.Printf("addr = %p : Instance already created.\n", &singleInstance)
	}
	return singleInstance
}

// Failed, Test
func GetInstanceRC() *single {
	if singleInstance == nil {
		singleInstance := &single{}
		fmt.Println("instance addr ==>", &singleInstance)
	} else {
		fmt.Printf("addr = %p : Instance already created.\n", &singleInstance)
	}
	return singleInstance
}

func main() {
	/*
		for i := 0; i < 10; i++ {
			go GetInstance()
		}
	*/

	// /*
	for i := 0; i < 100; i++ {
		go GetInstanceRC()
	}
	// */

	time.Sleep(2 * time.Second)

	var input string
	fmt.Println("Press Enter to exit.")
	fmt.Scanln(&input)
}
