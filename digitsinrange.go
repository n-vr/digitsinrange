package digitsinrange

func countXInN(x, n int) int {
	var count int
	for n > 0 {
		if n%10 == x {
			count++
		}
		n /= 10
	}
	return count
}

// CountDigitOccurancesInRangeNaive counts the number of occurances of a digit x
// in the range [0, max) by iterating over the range and counting the number of
// occurances of the digit in each number.
func CountDigitOccurancesInRangeNaive(x, max int) int {
	var count int
	for i := 0; i < max; i++ {
		count += countXInN(x, i)
	}
	return count
}

// CountDigitOccurancesInRange counts the number of occurances of a digit d
// in the range [0, n) by relying on some math. d cannot be 0.
//
//   - 1 occurrance of x from 0 to 10
//   - 10 occurances of x from 0 to 100
//   - 100 occurances of x from 0 to 1000
//   - etc.
//
// In cases where the least significant digit (lsd) is greater than digit d, we need
// to add the number of occurances of x in the range [0, lsd) to the count.
func CountDigitOccurancesInRange(d, n int) int {
	count := 0
	factor := 1

	for n > 0 {
		lsd := n % 10
		n /= 10

		count += n * factor

		if lsd >= d {
			count += (n + 1) * factor
		}

		factor *= 10
	}

	return count
}
