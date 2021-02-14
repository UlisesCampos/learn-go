package main

func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}
func SumAll(slices ...[]int) []int {
	// lengthOfNumbers := len(slices)
	// sums := make([]int, lengthOfNumbers)

	var sums []int
	for _, v := range slices {
		//sums[i] = Sum(v)
		sums = append(sums, Sum(v))
	}
	return sums
}
func SumAllSets(slices ...[]int) int {
	var sums []int
	var a int
	for _, v := range slices {
		sums = append(sums, Sum(v))
	}
	for _, v := range sums {
		a += v
	}
	return a
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
