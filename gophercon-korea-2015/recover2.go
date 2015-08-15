package main

import (
	"fmt"
	"log"
)

func divide(a, b int) int {
	return a / b
}

// START OMIT
func protect(g func()) {
	defer func() {
		log.Println("done")

		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()
	log.Println("start")
	g()
}
func main() {
	protect(func() {
		fmt.Println(divide(4, 0))
	})
}

// END OMIT