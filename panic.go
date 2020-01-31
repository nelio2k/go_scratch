package main

import (
	"fmt"
)

func createAndCatchPanicNoReturn() error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Caught panic: %v", r)
		}
	}()
	panic("Sending panic")
	return err
}

func createAndCatchPanic() (err error) {
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
	err := createAndCatchPanicNoReturn()
	fmt.Printf("Main error no pointer func returned: %v\n", err)
	err = nil
	createAndCatchPanicPtr(&err)
	fmt.Printf("Main error pointer func returned: %v\n", err)
	err = createAndCatchPanic()
	fmt.Printf("Main error no pointer func2 returned: %v\n", err)
}
