package matematicas

import (
	"fmt"
	"testing"
)

func TestPromedioCentrado(t *testing.T) {
	type prueba struct {
		datos     []int
		respuesta float64
	}

	prubas := []prueba{
		{
			datos:     []int{1, 2, 3, 4, 5},
			respuesta: 3,
		},
		{
			datos:     []int{3, 5, 7, 8},
			respuesta: 6,
		},
		{
			datos:     []int{10, 20, 30, 40, 50},
			respuesta: 30,
		},
		{
			datos:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			respuesta: 5,
		},
	}

	for _, v := range prubas {
		x := PromedioCentrado(v.datos)
		if x != v.respuesta {
			t.Error("Expected", v.respuesta, "Got", x)
		}
	}
}

func ExamplePromedioCentrado() {
	fmt.Println(PromedioCentrado([]int{1, 2, 3, 4, 5}))
	//Output
	//3
}

func BenchmarkPromedioCentrado(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PromedioCentrado([]int{1, 2, 3, 4, 5})
	}
}
