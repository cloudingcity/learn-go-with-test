package integer

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("2 plus 2 equala 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

	t.Run("3 plus 3 equala 6", func(t *testing.T) {
		sum := Add(3, 3)
		expected := 6

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
