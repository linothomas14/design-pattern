package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		go getLog()
	}

	fmt.Scanln() // Untuk menahan program agar tetap berjalan
}
