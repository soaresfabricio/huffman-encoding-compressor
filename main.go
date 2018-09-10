package main

import (
	"fmt"
	"huffman/huffman"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Incorrect usage.")
	} else if os.Args[1] == "-c" {
		compressor.CompressFile(os.Args[2], os.Args[3])
	} else if os.Args[1] == "-x" {
		compressor.DecompressFile(os.Args[2], os.Args[3])
	} else {
		fmt.Println("Incorrect usage.")
	}

}
