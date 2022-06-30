package palabra

import (
	"fmt"
	"testing"

	"github.com/MsCristianCantor/go-course/010-Ninja-nivel-13/02/cita"
)

func TestConteoUso(t *testing.T) {
	m := ConteoUso("uno dos tres tres tres")
	for k, v := range m {
		switch k {
		case "uno":
			if v != 1 {
				t.Error("Expected", 1, "Got", v)
			}
		case "dos":
			if v != 1 {
				t.Error("Expected", 1, "Got", v)
			}
		case "tres":
			if v != 3 {
				t.Error("Expected", 3, "Got", v)
			}
		}
	}
}

func TestConteo(t *testing.T) {
	n := Conteo(cita.Cuando)
	if n != 355 {
		t.Error("Got", n, "Expected", 355)
	}
}

func BenchmarkConteo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Conteo(cita.Cuando)
	}
}

func BenchmarkConteoUso(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConteoUso(cita.Cuando)
	}
}

func ExampleConteo() {
	fmt.Println(Conteo("uno dos tres"))
	//Output:
	//3
}
