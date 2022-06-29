package main

import "fmt"

type errorPer struct {
	info string
}

func (ep errorPer) Error() string {
	return fmt.Sprintf("aqui esta el error: %v", ep.info)
}

func main() {
	e1 := errorPer{
		info: "Necesito dormir mas",
	}

	foo(e1)
}

func foo(e error) {
	fmt.Println("foo corrio -", e, "\n", e.(errorPer).info)
}
