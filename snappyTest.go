package main

// A quick test program that compresses a specified file and outputs compression ratio and other stats

import (
	"flag"
	"fmt"
	"github.com/golang/snappy"
	"io/ioutil"
)

func main() {
	fmt.Println("Starting snappy test program")
	fileNamePtr := flag.String("datafile", "none", "Data file for input")

	flag.Parse()

	fileName := *fileNamePtr
	if fileName == "none" {
		fmt.Println("No filename entered.")
		fmt.Println("prog -datafile [filename]")
		return
	} else {
		fmt.Println("Data file: ", fileName)
	}

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file ", fileName)
		return
	}
	compressedContent := snappy.Encode(nil, content)

	bytesOfContent := len(content)
	bytesOfContentFromCompressed, err := snappy.DecodedLen(compressedContent)
	if err != nil {
		fmt.Printf("Error decoding length from compressed data: %v\n", err.Error())
		return
	}
	bytesOfCompressedContent := len(compressedContent)
	compressionRatio := float64(bytesOfContent) / float64(bytesOfCompressedContent)

	fmt.Println("Size of uncompressed file (bytes): ", bytesOfContent)
	fmt.Println("Size of uncompressed file from header (bytes): ", bytesOfContentFromCompressed)
	fmt.Println("Size of compressed file (bytes): ", bytesOfCompressedContent)
	fmt.Println("Compression Ratio: ", compressionRatio)

	return
}
