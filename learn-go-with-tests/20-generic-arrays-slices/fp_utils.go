package fp_utils

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	add := func(acc, curr int) int {
		return acc + curr
	}

	return Reduce(numbers, 0, add)
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
func SumAllTails(numbersToSum ...[]int) []int {
	sumTail := func(acc, curr []int) []int {
		if len(curr) == 0 {
			return append(acc, 0)
		} else {
			tail := curr[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(numbersToSum, []int{}, sumTail)
}

func Reduce[T, U any](
	collection []T,
	initialValue U,
	accumulator func(U, T) U) U {
	result := initialValue

	for _, item := range collection {
		result = accumulator(result, item)
	}

	return result
}

type Transaction struct {
	From, To string
	Sum      float64
}

func BalanceFor(trx []Transaction, name string) float64 {
	type BalanceMap = map[string]float64
	var bMap = make(BalanceMap)

	bMap = Reduce(trx, bMap, func(acc BalanceMap, curr Transaction) BalanceMap {
		bMap[curr.From] -= curr.Sum
		bMap[curr.To] += curr.Sum
		return acc
	})

	return bMap[name]
}
