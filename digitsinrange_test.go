package digitsinrange_test

import (
	"fmt"
	"testing"

	"github.com/n-vr/digitsinrange"
)

func BenchmarkCountDigitOccurancesInRangeNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		digitsinrange.CountDigitOccurancesInRangeNaive(9, 1250000)
	}
}

func BenchmarkCountDigitOccurancesInRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		digitsinrange.CountDigitOccurancesInRange(9, 1250000)
	}
}

var testCases = []struct {
	d, n int
	want int
}{
	{5, 0, 0},
	{5, 0, 0},
	{5, 10, 1},
	{5, 15, 1},
	{5, 16, 2},
	{5, 20, 2},
	{5, 100, 20},
	{5, 110, 21},
	{5, 350, 65},
	{1, 13, 5},
}

func TestCountDigitOccurancesInRangeNaive(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("d%dn%d", tc.d, tc.n), func(t *testing.T) {
			got := digitsinrange.CountDigitOccurancesInRangeNaive(tc.d, tc.n)
			if got != tc.want {
				t.Errorf("CountDigitOccurancesInRangeNaive(%d, %d) = %d; want %d", tc.d, tc.n, got, tc.want)
			}
		})
	}
}

func TestCountDigitOccurancesInRange(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("d%dn%d", tc.d, tc.n), func(t *testing.T) {
			match, naive, fast := matchApproaches(tc.d, tc.n)
			if !match {
				t.Errorf("CountDigitOccurancesInRange(%d, %d) = %d; want %d", tc.d, tc.n, fast, naive)
			}
		})
	}
}

func matchApproaches(d, n int) (bool, int, int) {
	naive := digitsinrange.CountDigitOccurancesInRangeNaive(d, n)
	fast := digitsinrange.CountDigitOccurancesInRange(d, n)
	return naive == fast, naive, fast
}
