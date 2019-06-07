package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("NEIL Time comparison")
	time1, err := time.Parse(time.RFC3339, "2018-11-21T00:01:02.03Z")
	if err != nil {
		fmt.Printf("error1: %v\n", err)
		return
	}
	time2, err := time.Parse(time.RFC3339, "2018-11-21T00:01:02Z")
	if err != nil {
		fmt.Printf("error2: %v\n", err)
		return
	}

	isEqual := time1.Equal(time2)

	fmt.Printf("time1 is equal to time2? %v\n", isEqual)
	return
}
