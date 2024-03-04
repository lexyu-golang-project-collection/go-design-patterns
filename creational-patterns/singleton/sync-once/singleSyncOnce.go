package so

import (
	"fmt"
	"sync"
)

type single struct {
}

var once sync.Once
var singleInstance *single

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

// Try Something
func GetInstanceRC() *single {
	instance := &single{}
	fmt.Println("instance addr ==>", &instance)
	return instance
}
