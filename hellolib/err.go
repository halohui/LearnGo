package hellolib

import (
	"errors"
	"log"
	"fmt"
)

var errDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errDivByZero
	}
	return x / y, nil
}

func TestDiv() {
	z, err := div(5, 0)
	if err == errDivByZero {
		log.Fatalln(err)
	}
	fmt.Println(z)
}

type DivError struct {
	x, y int
}

func (DivError) Error() string {
	return "division by zero"
}

func div1(x, y int) (int, error) {
	if y == 0 {
		return 0, DivError{x, y}
	}

	return x / y, nil
}

func TestPanic() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()

	panic("I am dead")
	println("exit") //不会被执行
}
