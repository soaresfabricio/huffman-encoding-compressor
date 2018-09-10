package main

import (
	"fmt"
	"huffman/huffman"
	"os"
)

func printUsage() {
	fmt.Println("Usage: [this tool] [-c | -x] [input file] [output file]")
	fmt.Println("\t-c \tCompresses input to output.")
	fmt.Println("\t-x \tDecompresses input to output.")
	fmt.Println("File extensions should be added mannually to prevent overwriting.")
}

func main() {
	if len(os.Args) != 4 {
		printUsage()
	} else if os.Args[1] == "-c" {
		compressor.CompressFile(os.Args[2], os.Args[3])
	} else if os.Args[1] == "-x" {
		compressor.DecompressFile(os.Args[2], os.Args[3])
	} else {
		printUsage()
	}

}
