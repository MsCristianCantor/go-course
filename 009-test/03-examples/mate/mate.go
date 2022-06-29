// Package mate help us to add
package mate

// Sum add all of number delivered
func Sum(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}
