package main

import (
	"errors"
	"fmt"
	"log"
)

type raizError struct {
	lat string
	lon string
	err error
}

func (re raizError) Error() string {
	return fmt.Sprintf("error matematico: %v %v %v", re.lat, re.lon, re.err)
}

func main() {
	_, err := raiz(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func raiz(f float64) (float64, error) {
	if f < 0 {
		e := errors.New("Murio")
		return 0, raizError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}
