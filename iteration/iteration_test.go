package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"

	if expected != repeated {
		t.Errorf("Expected %q but got %q", repeated, expected)
	}
}

func ExampleRepeat() {
	repeated := Repeat("v", 4)
	fmt.Println(repeated)
	// Output: vvvv
}

func TestIndex(t *testing.T) {
	index := checkIndex("escape", "s")
	expected := 1

	if expected != index {
		t.Errorf("Expected %d but got %d", expected, index)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}
