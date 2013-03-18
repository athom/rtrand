package main

import (
	"fmt"
	"github.com/athom/rtrand"
	"math/rand"
)

func main() {
	i := 0
	for i < 10 {
		fmt.Println(rand.Float32())
		i += 1
	}

	fmt.Println("-----")

	i = 0
	for i < 10 {
		fmt.Println(rtrand.Float32())
		i += 1
	}
}
