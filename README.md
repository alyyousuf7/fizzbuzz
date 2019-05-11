# fizzbuzz
FizzBuzz implementation in Go.

Usage:
```go
package main

import (
	"fmt"
	
	"github.com/alyyousuf7/fizzbuzz"
)

func main() {
	fb, err := fizzbuzz.New(
		[]int{3, 5, 7},
		[]string{"fizz", "buzz", "bazz"},
	)
	if err != nil {
		panic(err)
	}

	for i := 1; i <= 15; i++ {
		fmt.Println(fb.Process(i))
	}
}
```

Outputs:
```bash
1
2
fizz
4
buzz
fizz
bazz
8
fizz
buzz
11
fizz
13
bazz
fizzbuzz
```
