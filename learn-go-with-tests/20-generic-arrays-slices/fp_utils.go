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

type findResult[T any] struct {
	item  T
	found bool
}

func Find[T any](xs []T, pred func(T) bool) (item T, found bool) {
	result := findResult[T]{item, found}

	result = Reduce(xs, result, func(acc findResult[T], curr T) findResult[T] {
		if pred(curr) && !acc.found {
			return findResult[T]{curr, true}
		} else {
			return acc
		}
	})

	return result.item, result.found
}

type Account struct {
	Name    string
	Balance float64
}

type Transaction struct {
	From, To string
	Sum      float64
}

func NewTransaction(a, b Account, amount float64) Transaction {
	return Transaction{a.Name, b.Name, amount}
}

func NewBalanceFor(acc Account, trx []Transaction) Account {
	return Reduce(trx, acc, func(acc Account, curr Transaction) Account {
		if curr.From == acc.Name {
			acc.Balance -= curr.Sum
		}

		if curr.To == acc.Name {
			acc.Balance += curr.Sum
		}

		return acc
	})
}
