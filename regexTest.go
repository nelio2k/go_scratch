package main

import (
	"fmt"
	"regexp"
)

func main() {
	key := "`[$%XDCRInternalKey*%$]` = \"something\""
	keyEscape := "``[$%XDCRInternalKey*%$]`` = \"something\""

	rg := regexp.MustCompile("``[$%XDCRInternalKey*%$]``")
	if rg.MatchString(key) {
		fmt.Printf("Match1\n")
	}
	if rg.MatchString(keyEscape) {
		fmt.Printf("Match2\n")
	}

	return
}
