package main

import (
	"fmt"
)

func createAndCatchPanic() error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Caught panic: %v", r)
		}
	}()
	panic("Sending panic")
	return err
}

func createAndCatchPanicPtr(errPtr *error) {
	defer func() {
		if r := recover(); r != nil {
			*errPtr = fmt.Errorf("Caught panic: %v", r)
		}
	}()
	panic("Temporary panic pointer")
}

func main() {
	err := createAndCatchPanic()
	fmt.Printf("Main error no pointer func returned: %v\n", err)
	err = nil
	createAndCatchPanicPtr(&err)
	fmt.Printf("Main error pointer func returned: %v\n", err)
}
