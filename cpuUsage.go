package main

import (
	"fmt"
	"syscall"
)

func main() {
	rusage := &syscall.Rusage{}
	err := syscall.Getrusage(syscall.RUSAGE_SELF, rusage)

	if err != nil {
		fmt.Printf("rusage err: %v\n", err)
	}
	fmt.Printf("rusage detail: %v\n", rusage.Utime)

	return
}
