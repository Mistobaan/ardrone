package navdata

import (
	"testing"
)

func BenchmarkReader_ReadNavdata(b *testing.B) {
	b.StopTimer()
	reader := NewReader(fixture())
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		reader.ReadNavdata()
	}
}

func BenchmarkParse(b *testing.B) {
	b.StopTimer()
	buf := fixtureBytes()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Parse(buf)
	}
}
