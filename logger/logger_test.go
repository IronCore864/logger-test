package logger

import (
	"io"
	"os"
	"testing"
)

// BenchmarkLoggerNoticef tests Noticef with string formatting.
func BenchmarkLoggerNoticef(b *testing.B) {
	logger := New(io.Discard, "Benchmark: ")
	SetLogger(logger)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Noticef("Formatted message with number: %d and string: %s", 42, "test")
	}
}

// BenchmarkLoggerNoticefConcurrent tests concurrent logging performance.
func BenchmarkLoggerNoticefConcurrent(b *testing.B) {
	logger := New(io.Discard, "Benchmark: ")
	SetLogger(logger)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Noticef("Concurrent message from goroutiner: %d and string: %s", 42, "test")
		}
	})
}

// BenchmarkLoggerNoticefStderr tests Noticef with real os.Stderr output.
func BenchmarkLoggerNoticefStderr(b *testing.B) {
	logger := New(os.Stderr, "Benchmark: ")
	SetLogger(logger)

	b.ResetTimer()
	for b.Loop() {
		Noticef("Formatted message with number: %d and string: %s", 42, "test")
	}
}

// BenchmarkLoggerNoticefStderrConcurrent tests concurrent logging to os.Stderr.
func BenchmarkLoggerNoticefStderrConcurrent(b *testing.B) {
	logger := New(os.Stderr, "Benchmark: ")
	SetLogger(logger)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Noticef("Concurrent message from goroutiner: %d and string: %s", 42, "test")
		}
	})
}
