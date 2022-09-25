package pow

import "testing"

var (
	ps = NewPOWServer(30, 6)
)

func BenchmarkGenerate512Hash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generate512Hash(ps.GenerateUniqKey())
	}
}
func BenchmarkGenerate3Hash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generate3Hash(ps.GenerateUniqKey())
	}
}
