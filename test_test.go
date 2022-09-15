package main_test

import (
	"go-bench/arr"
	"testing"
)

func TestGoRoutinesArr(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		go arr.TestArr()
	}
}
