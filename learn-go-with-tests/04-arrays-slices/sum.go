package arrays_slices

func Sum(numbers []int) (sum int) {
	for _, v := range numbers {
		sum += v
	}

	return
}
