package strgen

import (
	"regexp"
	"testing"
)

// Test for RandString
func TestHexRandString(t *testing.T) {
    result := HexRandString(10)
    if len(result) != 10 {
        t.Errorf("Expected string of length 10, got %d", len(result))
    }
    if !regexp.MustCompile("^[0-9a-f]+$").MatchString(result) {
        t.Errorf("String contains non-hex characters: %s", result)
    }
}

// Benchmark for RandString
func BenchmarkHexRandString(b *testing.B) {
    for i := 0; i < b.N; i++ {
        HexRandString(10)
    }
}