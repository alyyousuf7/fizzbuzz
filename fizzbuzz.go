package fizzbuzz

import (
	"fmt"
	"strconv"
	"strings"
)

type fizzbuzz struct {
	nums  []int
	words []string
}

func New(nums []int, words []string) (*fizzbuzz, error) {
	if len(nums) != len(words) {
		return nil, fmt.Errorf("Numbers and words should be equal")
	}

	return &fizzbuzz{nums, words}, nil
}

func (fb fizzbuzz) Process(n int) string {
	out := []string{}
	for i := 0; i < len(fb.nums); i++ {
		if n%fb.nums[i] == 0 {
			out = append(out, fb.words[i])
		}
	}

	if len(out) == 0 {
		return strconv.FormatInt(int64(n), 10)
	}

	return strings.Join(out, "")
}
