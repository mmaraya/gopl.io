package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%v, %v", i, os.Args[i])
		fmt.Println()
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Elapsed time: %v\n", elapsed)

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}
