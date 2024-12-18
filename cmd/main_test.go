package main

import "testing"

func BenchmarkTwoSum(b *testing.B) {
	// Benchmarking the reverseWordsSimple function
	b.Run("reverseWords", func(b *testing.B) {
		addWaitgroup()
		for i := 0; i < b.N; i++ {
			synchronousMain()
		}
	})

	// Benchmarking the reverseWordsComplex function
	b.Run("reverseWordsSlow", func(b *testing.B) {
		addWaitgroup()
		for i := 0; i < b.N; i++ {
			asynchronousMain()
		}
	})
}