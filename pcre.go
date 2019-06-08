package main

import (
	"flag"
	"fmt"
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
	"os"
)

type PcreWrapper struct {
	pcreRegex *pcre.Regexp
}

func MakePcreRegex(expression string) (*PcreWrapper, error) {
	pcreRegex := pcre.MustCompile(expression, 0 /*flags*/)
	wrapper := &PcreWrapper{pcreRegex: &pcreRegex}
	return wrapper, nil
}

func (p *PcreWrapper) ReplaceAll(orig, strToSub []byte, flags int) []byte {
	return p.pcreRegex.ReplaceAll(orig, strToSub, flags)
}

func (p *PcreWrapper) Match(str string) bool {
	matcher := pcre.Matcher{}
	matcher.Reset(*p.pcreRegex, []byte(str), 0)
	return matcher.Matches()
}

var options struct {
	expression string
	testString string
}

func argParse() {
	flag.StringVar(&options.expression, "pcreExpression", "", "PCRE expression")
	flag.StringVar(&options.testString, "testString", "", "String to validate against")
	flag.Parse()
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	argParse()

	if options.expression == "" {
		fmt.Printf("Expression required\n")
		usage()
		os.Exit(-1)
	} else if options.testString == "" {
		fmt.Printf("TestString required\n")
		usage()
		os.Exit(-1)
	}

	pcreRegex, err := MakePcreRegex(options.expression)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error compiling PCRE: %v\n", err.Error())
		os.Exit(-1)
	}

	res := pcreRegex.Match(options.testString)
	if res {
		fmt.Println("True")
		os.Exit(0)
	} else {
		fmt.Println("False")
		os.Exit(1)
	}
}
