package word

import (
	"fmt"
	"testing"
)

const quote = `couple word for test benchmark`

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote)
	}
}
func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote)
	}
}

func ExampleCount() {
	fmt.Println(Count("Greetings, everyone"))
	// Output:
	// 2
}

// bad practice example for map
func ExampleUseCount() {
	fmt.Println(UseCount("Hello Hello"))
	// Output:
	// map[Hello:2]
}

func TestCount(t *testing.T) {
	x := Count("Hello, Hello everyone")
	if x != 3 {
		t.Error("Expected", 3, "got", x)
	}
}

func TestUseCount(t *testing.T) {
	x := UseCount("Greet, greet greet, greet, hello hello hello")
	a := map[string]int{
		"hello":  3,
		"Greet,": 1,
		"greet,": 2,
		"greet":  1,
	}

	for k := range x {
		aVal := a[k]
		xVal := x[k]

		if aVal != xVal {
			t.Error("Expected", a, "got", x, "with key", k)
		}
	}
}
