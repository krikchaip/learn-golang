package arrays_slices

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
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	// ** calling reflect.DeepEqual is now typesafe 👍🏻
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		// ** you can't compare two slices with the `==` operator
		// if got != want {
		// 	t.Errorf("got %v want %v", got, want)
		// }

		// ** use this instead
		// ** NOTE: using this method requires manual typecheck (due to the `any` params)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("V1", func(t *testing.T) {
		got := SumAllV1([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		checkSums(t, got, want)
	})

	t.Run("latest", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
