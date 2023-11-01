package pkgo

import (
	"fmt"
	"testing"
)

func TestApproach8(t *testing.T) {
	fmt.Println(RandStr(100))
}

func BenchmarkApproach8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandStr(10)
	}
}
