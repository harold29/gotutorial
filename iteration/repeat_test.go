package iteration

import "testing"
import "fmt"

func TestRepeat(t *testing.T) {
	// t.Fatal("not implemented")
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}

	t.Run("with a desired number of repeated char", func(t *testing.T) {
		repeated := Repeat("a", 8)
		expected := "aaaaaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 8)
	fmt.Println(repeated)
	// Output: bbbbbbbb
}
