package render

import "testing"

// Some benchmarks that handle large slices and reusing memory

func BenchmarkAppendingReuse(b *testing.B) {
	count := 10000
	vals := make([]int, 0, count)

	for i := 0; i < b.N; i++ {
		vals = vals[:0]

		for j := 0; j < count; j++ {
			vals = append(vals, j)
		}
	}
}

func BenchmarkAppendingMakeScratchWithCapacity(b *testing.B) {
	count := 10000

	for i := 0; i < b.N; i++ {
		vals := make([]int, 0, count)
		for j := 0; j < count; j++ {
			vals = append(vals, j)
		}
	}
}

func BenchmarkAppendingMakeScratchNoCapacity(b *testing.B) {
	count := 10000

	for i := 0; i < b.N; i++ {
		vals := make([]int, 0)
		for j := 0; j < count; j++ {
			vals = append(vals, j)
		}
	}
}

func BenchmarkInsertToSlice(b *testing.B) {
	count := 10000
	half := count / 2

	vals := make([]int, count, count+1)

	for i := 0; i < b.N; i++ {
		vals = vals[:count]
		copy(vals[half+1:], vals[half:])
		vals[half-1] = i

		if len(vals) != count+1 {
			b.Fatalf("Length wrong: %d", len(vals))
		}
	}
}
