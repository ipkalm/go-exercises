package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	const multiplier int = 7
	type tableTest struct {
		hy int
		dy int
	}
	tableTests := []tableTest{}

	humanYears := []int{7, 14, 700, 21}

	for _, humYear := range humanYears {
		tableTests = append(tableTests, tableTest{humYear, (humYear / multiplier)})
	}

	for _, v := range tableTests {
		x := Years(v.hy)
		if x != v.dy {
			t.Error("Expected", v.dy, "got", x)
		}
	}
}

func TestYearsTwo(t *testing.T) {
	const multiplier int = 7
	type tableTest struct {
		hy int
		dy int
	}
	tableTests := []tableTest{}

	humanYears := []int{7, 14, 700, 21}

	for _, humYear := range humanYears {
		tableTests = append(tableTests, tableTest{humYear, (humYear / multiplier)})
	}

	for _, v := range tableTests {
		x := YearsTwo(v.hy)
		if x != v.dy {
			t.Error("Expected", v.dy, "got", x)
		}
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}

func ExampleYears() {
	fmt.Println(Years(70))
	// Output:
	// 10
}

func ExampleYearsTwo() {
	fmt.Println(Years(21))
	// Output:
	// 3
}
