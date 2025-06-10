package script

import (
	"bytes"
	"testing"
	"unicode/utf8"
)

func TestTrimToMaxBytes(t *testing.T) {
	t.Run("returns original slice when under limit", func(t *testing.T) {
		input := []byte("hello")
		res := trimToMaxBytes(input, 10)
		if !bytes.Equal(res, input) {
			t.Fatalf("expected %q got %q", input, res)
		}
		if len(input) > 0 && &res[0] != &input[0] {
			t.Fatalf("function should return the original slice reference")
		}
	})

	t.Run("trims and keeps utf8", func(t *testing.T) {
		input := []byte("안녕하세요")
		res := trimToMaxBytes(input, 8)
		if !utf8.Valid(res) {
			t.Fatalf("result is not valid utf8: %v", res)
		}
		if string(res) != "안녕" {
			t.Fatalf("expected %q got %q", "안녕", string(res))
		}
		if len(res) > 8 {
			t.Fatalf("result length %d exceeds limit", len(res))
		}
	})
}
