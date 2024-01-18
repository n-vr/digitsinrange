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

func TestCountDigitOccurancesInRangeNaive(t *testing.T) {
	testCases := []struct {
		d, n int
		want int
	}{
		{5, 0, 0},
		{5, 10, 1},
		{5, 15, 1},
		{5, 16, 2},
		{5, 20, 2},
		{5, 100, 20},
		{5, 110, 21},
		{5, 350, 65},
	}

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
	want := digitsinrange.CountDigitOccurancesInRangeNaive(9, 1250000)
	got := digitsinrange.CountDigitOccurancesInRange(9, 1250000)

	if got != want {
		t.Errorf("CountDigitOccurancesInRange(9, 1250000) = %d; want %d", got, want)
	}
}
