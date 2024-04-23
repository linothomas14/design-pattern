package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type log struct {
}

var logInstance *log

func getLog() *log {
	if logInstance == nil {
		lock.Lock() // Untuk Memastikan hanya ada 1 thread (go routine) yang mengakses method ini
		defer lock.Unlock()
		if logInstance == nil {
			fmt.Println("Creating log instance now.")
			logInstance = &log{}
		} else {
			fmt.Println("log instance already created.")
		}
	} else {
		fmt.Println("log instance already created.")
	}

	return logInstance
}
