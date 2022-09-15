package main

import (
	"errors"
	"fmt"
)

func main2() {
	e := errors.New("first error")
	e2 := wrap(e, "second error")
	e3 := wrap(e2, "test")
	fmt.Println(e3.Error())
	fmt.Println(Unwrap(e3).Error())
	fmt.Println(Unwrap(e2).Error())
	fmt.Println(Unwrap(e))
}

func printError(err error) {
	fmt.Println(err.Error())
}

type ali struct {
	b   string
	err error
}

func wrap(err error, msg string) error {
	return ali{
		b:   msg,
		err: err,
	}
}
func (a ali) Error() string {
	return a.b
}

func Unwrap(e error) error {
	switch t := e.(type) {
	case ali:
		return t.err
	default:
		return nil
	}
}
