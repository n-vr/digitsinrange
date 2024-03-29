package digitsinrange

func countDInN(d, n int) int {
	var count int
	for n > 0 {
		if n%10 == d {
			count++
		}
		n /= 10
	}
	return count
}

// CountDigitOccurancesInRangeNaive counts the number of occurances of a digit d
// in the range [0, n) by iterating over the range and counting the number of
// occurances of the digit in each number.
func CountDigitOccurancesInRangeNaive(d, n int) int {
	var count int
	for i := 0; i <= n; i++ {
		count += countDInN(d, i)
	}
	return count
}

// CountDigitOccurancesInRange counts the number of occurances of a digit d
// in the range [0, n) by relying on some math.
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
	prevLSD := 0

	for n > 0 {
		lsd := n % 10
		n /= 10

		count += n * factor

		if lsd >= d {
			count += prevLSD + 1
		}

		// fmt.Println("lsd:", lsd, "n:", n, "factor:", factor, "count:", count)

		factor *= 10
		prevLSD = lsd
	}

	return count
}
