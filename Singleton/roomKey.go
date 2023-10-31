package main

import (
	"fmt"
	"sync"
)

var mu = &sync.Mutex{}

type key struct {
	name string
}

var singleInstance *key

func getInstance() *key {
	if singleInstance == nil {
		mu.Lock()
		defer mu.Unlock()

		if singleInstance == nil {
			fmt.Println("Creating single instance of the room key.")
			singleInstance = &key{"111 room"}
		} else {
			fmt.Println(singleInstance.name, " key is already created.")
		}
	} else {
		fmt.Println(singleInstance.name, " key is already created.")
	}

	return singleInstance
}
