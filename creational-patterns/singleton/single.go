package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

var instance *single

func getInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance.")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Instance already created.")
	}
	return singleInstance
}

func getInstanceRC() *single {
	instance := &single{}
	fmt.Println("instance addr ==>", &instance)
	return instance
}
