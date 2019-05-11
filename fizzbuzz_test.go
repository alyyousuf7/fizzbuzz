package fizzbuzz_test

import (
	"testing"

	"github.com/alyyousuf7/fizzbuzz"
	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	_, err := fizzbuzz.New(
		[]int{1, 2, 3},
		[]string{"one", "two", "three"},
	)
	assert.NoError(t, err)

	_, err = fizzbuzz.New(
		[]int{1, 2, 3},
		[]string{"one", "two"},
	)
	assert.Error(t, err)

	_, err = fizzbuzz.New(
		[]int{1, 2},
		[]string{"one", "two", "three"},
	)
	assert.Error(t, err)
}

func TestProcess(t *testing.T) {
	fb, err := fizzbuzz.New(
		[]int{3},
		[]string{"fizz"},
	)
	assert.NoError(t, err)
	assert.NotNil(t, fb)

	assert.Equal(t, "fizz", fb.Process(0))
	assert.Equal(t, "1", fb.Process(1))
	assert.Equal(t, "2", fb.Process(2))
	assert.Equal(t, "fizz", fb.Process(3))
	assert.Equal(t, "4", fb.Process(4))
	assert.Equal(t, "5", fb.Process(5))
	assert.Equal(t, "fizz", fb.Process(6))
	assert.Equal(t, "7", fb.Process(7))
	assert.Equal(t, "8", fb.Process(8))
	assert.Equal(t, "fizz", fb.Process(9))
	assert.Equal(t, "10", fb.Process(10))

	assert.Equal(t, "14", fb.Process(14))
	assert.Equal(t, "fizz", fb.Process(15))

	assert.Equal(t, "fizz", fb.Process(105))

	fb, err = fizzbuzz.New(
		[]int{3, 5},
		[]string{"fizz", "buzz"},
	)
	assert.NoError(t, err)
	assert.NotNil(t, fb)

	assert.Equal(t, "fizzbuzz", fb.Process(0))
	assert.Equal(t, "1", fb.Process(1))
	assert.Equal(t, "2", fb.Process(2))
	assert.Equal(t, "fizz", fb.Process(3))
	assert.Equal(t, "4", fb.Process(4))
	assert.Equal(t, "buzz", fb.Process(5))
	assert.Equal(t, "fizz", fb.Process(6))
	assert.Equal(t, "7", fb.Process(7))
	assert.Equal(t, "8", fb.Process(8))
	assert.Equal(t, "fizz", fb.Process(9))
	assert.Equal(t, "buzz", fb.Process(10))

	assert.Equal(t, "14", fb.Process(14))
	assert.Equal(t, "fizzbuzz", fb.Process(15))

	assert.Equal(t, "fizzbuzz", fb.Process(105))

	fb, err = fizzbuzz.New(
		[]int{3, 5, 7},
		[]string{"fizz", "buzz", "bazz"},
	)
	assert.NoError(t, err)
	assert.NotNil(t, fb)

	assert.Equal(t, "fizzbuzzbazz", fb.Process(0))
	assert.Equal(t, "1", fb.Process(1))
	assert.Equal(t, "2", fb.Process(2))
	assert.Equal(t, "fizz", fb.Process(3))
	assert.Equal(t, "4", fb.Process(4))
	assert.Equal(t, "buzz", fb.Process(5))
	assert.Equal(t, "fizz", fb.Process(6))
	assert.Equal(t, "bazz", fb.Process(7))
	assert.Equal(t, "8", fb.Process(8))
	assert.Equal(t, "fizz", fb.Process(9))
	assert.Equal(t, "buzz", fb.Process(10))

	assert.Equal(t, "bazz", fb.Process(14))
	assert.Equal(t, "fizzbuzz", fb.Process(15))

	assert.Equal(t, "fizzbuzzbazz", fb.Process(105))
}
