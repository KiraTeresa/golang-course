package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Jenny")
	if s != "Welcome my dear Jenny" {
		t.Error("got", s, "expected", "Welcome my dear Jenny")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Jenny"))
	// Output:
	// Welcome my dear Jenny
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Jenny")
	}
}
