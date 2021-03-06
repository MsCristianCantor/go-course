package perrito

import (
	"fmt"
	"testing"
)

func TestEdad(t *testing.T) {
	n := Edad(10)
	if n != 70 {
		t.Error("Got", n, "Expected", 70)
	}
}

func TestEdadDos(t *testing.T) {
	n := EdadDos(10)
	if n != 70 {
		t.Error("Got", n, "Expected", 70)
	}
}

func ExampleEdad() {
	fmt.Println(Edad(10))
	//Output:
	//70
}

func ExampleEdadDos() {
	fmt.Println(EdadDos(10))
	//Output:
	//70
}

func BenchmarkEdad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Edad(10)
	}
}

func BenchmarkEdadDos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EdadDos(10)
	}
}
