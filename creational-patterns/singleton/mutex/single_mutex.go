package main

import (
	"fmt"
	"sync"
	"time"
)

type single struct {
}

var (
	singleInstance *single
	lock           = &sync.Mutex{}
)

func GetInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Mutex - Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Printf("addr = %p : Instance already created.\n", &singleInstance)
		}
	} else {
		fmt.Printf("addr = %p : Instance already created.\n", &singleInstance)
	}

	return singleInstance
}

func main() {
	for i := 0; i < 10; i++ {
		go GetInstance()
	}

	time.Sleep(2 * time.Second)

	var input string
	fmt.Println("Press Enter to exit.")
	fmt.Scanln(&input)
}
