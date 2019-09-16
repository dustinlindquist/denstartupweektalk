package benchmark

import (
	"testing"
)

func BenchmarkConcat100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var start = ""
		for m := 0; m < 100; m++ {
			start = Concat(start, "123456789")
		}
	}
}

func BenchmarkSprintfConcat100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var start = ""
		for m := 0; m < 100; m++ {
			start = SprintfConcat(start, "123456789")
		}
	}
}

func BenchmarkBufferConcat100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var start = ""
		for m := 0; m < 100; m++ {
			start = BufferConcat(start, "123456789")
		}
	}
}

func BenchmarkStringsConcat100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var start = ""
		for m := 0; m < 100; m++ {
			start = StringsConcat(start, "123456789")
		}
	}
}
