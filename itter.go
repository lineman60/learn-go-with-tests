package integers

import "testing"

const repeateCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeateCount; i++ {
		repeated += character
	}
	return repeated
}

func BenchmarkRepeate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
