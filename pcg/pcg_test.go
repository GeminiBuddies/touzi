package pcg

import (
	"testing"
)

func TestPCG_Next(t *testing.T) {

}

func BenchmarkPCG_Next(b *testing.B) {
	// p := Factory(uint128.From64(42), uint128.From64(54))
	p := New()

	for i := 0; i < b.N; i++ {
		p.Next()
	}
}
