package so

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance.")
				singleInstance = &single{}
			})
	} else {
		fmt.Printf("addr = %p\n", &singleInstance)
		fmt.Println("Instance already created.")
	}
	return singleInstance
}

// Try Something
func GetInstanceRC() *single {
	instance := &single{}
	fmt.Println("instance addr ==>", &instance)
	return instance
}
