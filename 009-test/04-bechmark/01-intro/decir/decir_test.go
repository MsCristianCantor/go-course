package decir

import (
	"fmt"
	"testing"
)

func TestSaludar(t *testing.T) {
	s := Saludar("Camilo")
	if s != "Bienvenido querido Camilo" {
		t.Error("Expected", "Bienvenido querido Camilo", "Got", s)
	}
}

func ExampleSaludar() {
	fmt.Println(Saludar("Camilo"))
	//Output:
	//Bienvenido querido Camilo
}

func BenchmarkSaludar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Saludar("Camilo")
	}
}
