package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 5}

		got := Sum(numbers)
		want := 8

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("sum of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{5, 9})
		want := []int{3, 14}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("sum of all sets of slices", func(t *testing.T) {
		got := SumAllSets([]int{1, 2}, []int{3, 6}, []int{9, 10})
		want := 31
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 5, 10}, []int{}, []int{7, 10, 9})
		want := []int{17, 0, 19}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
