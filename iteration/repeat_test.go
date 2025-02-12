package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("passed repeatCount", func(t *testing.T) {
		repeatCount := 5
		character := "a"
		repeated := Repeat(character, repeatCount)

		assertRepeated(t, repeatCount, character, repeated)
	})
	t.Run("not passed repeatCount", func(t *testing.T) {
		repeatCount := 0
		character := "b"
		repeated := Repeat(character, repeatCount)

		assertRepeated(t, repeatCount, character, repeated)
	})
}

func assertRepeated(t testing.TB, repeatCount int, character, repeated string) {
	t.Helper()
	if repeatCount <= 0 {
		repeatCount = 5
	}
	counter := strings.Count(repeated, character)
	if counter != repeatCount {
		t.Errorf(
			"expected %q repeated %d times, got repeated %d times",
			character,
			repeatCount,
			counter)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("x", 3)
	fmt.Println(repeated)
	// Output: xxx
}
