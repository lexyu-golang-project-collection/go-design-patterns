package mx

import (
	"fmt"
	"sync"
)

type single struct {
}

var singleInstance *single
var lock = &sync.Mutex{}

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
