package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%v, %v", i, os.Args[i])
		fmt.Println()
	}
}
