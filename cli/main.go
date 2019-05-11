package main

import (
	"fmt"

	"github.com/alyyousuf7/fizzbuzz"
)

func main() {
	fb, err := fizzbuzz.New(
		[]uint{3, 5, 7},
		[]string{"fizz", "buzz", "bazz"},
	)
	if err != nil {
		panic(err)
	}

	for i := 1; i <= 15; i++ {
		fmt.Println(fb.Process(i))
	}
}
