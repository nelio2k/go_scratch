package main 

import (
	"fmt"
)

func main() {
	f := fibonacci()
	for i :=0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	var prev1 int = 0
	var prev int = 1
	var count int = 0
	return func() int {
		if count == 0 {
			count++
			return 0;
		} else if count == 1 {
			count++
			return 1;
		} else {
			var retVal int = prev + prev1
			prev1 = prev
			prev = retVal
			return retVal
		}
	}
}

