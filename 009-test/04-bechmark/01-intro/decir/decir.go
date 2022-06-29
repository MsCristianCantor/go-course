package decir

import "fmt"

// Saludar return salute
func Saludar(s string) string {
	return fmt.Sprint("Bienvenido querido ", s)
}
