package images

import (
	"fmt"
	"math"
	"testing"
)

func Test_deltaBlockSize(t *testing.T) {
	tests := []struct {
		x    int64
		want uint32
	}{
		{0, 256},
		{1, 256},
		{100, 256},
		{1_024, 256},
		{33_333, 256},
		{65_536, 256},
		{88_887, 512},
		{262_144, 512},
		{262_145, 512},
		{777_111, 1024},
		{22_654_123, 8192},
		{1_333_555_888, 65536},
		{35_000_000_000, 262144},
		{123_456_678_901, 524288},
		{4_611_686_018_427_387_904, 2147483648},
		{5_000_000_000_000_000_000, 2147483648},
		{math.MaxInt64, 2147483648},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("deltaBlockSize(%v)", tt.x), func(t *testing.T) {
			if got := deltaBlockSize(tt.x); got != tt.want {
				t.Errorf("got deltaBlockSize(%v) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}
