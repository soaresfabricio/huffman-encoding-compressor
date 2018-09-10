package compressor

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/cheggaaa/pb.v1"
)

// CompressFile will compress input to ouput using huffman
func CompressFile(input string, output string) {

	inputFile, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	inputData, err := ioutil.ReadAll(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	byteCount := make([]uint32, 256)

	for _, x := range inputData {
		byteCount[x]++
	}

	outputFile, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	bufferedWriter := bufio.NewWriter(outputFile)

	binary.Write(bufferedWriter, binary.LittleEndian, byteCount)

	n, err := outputFile.Seek(4, 1)
	if err != nil {
		log.Fatal(n, err)
	}

	fmt.Println("Compressing...")

	root := BuildHuffmanTree(byteCount)

	aux := byte(0)
	size := uint32(0)

	bar := pb.StartNew(len(inputData))

	buffer := make([]byte, 1024)

	for _, c := range inputData {

		bar.Increment()

		getCode(root, &c, buffer, 0)

		for _, i := range buffer {

			if i == 2 {
				break
			}

			if i == 1 {
				aux = aux | (1 << (size % 8))
			}

			size++

			if (size % 8) == 0 {
				bufferedWriter.WriteByte(aux)
				aux = 0
			}
		}
	}

	bufferedWriter.WriteByte(aux)
	bufferedWriter.Flush()

	outputFile.Seek(1024, 0)

	binary.Write(outputFile, binary.LittleEndian, size)

}

func getCode(n *treeNode, c *byte, buffer []byte, size int) bool {
	if !(n.left != nil || n.right != nil) && (n.c == *c) {
		buffer[size] = 2
		return true
	} else {
		found := false
		if n.left != nil {
			buffer[size] = 0
			found = getCode(n.left, c, buffer, size+1)
		}
		if found == false && n.right != nil {
			buffer[size] = 1
			found = getCode(n.right, c, buffer, size+1)
		}
		if found != true {
			buffer[size] = 2
		}
		return found
	}
}

func DecompressFile(input string, output string) {

	inputFile, err := os.Open(input)
	if err != nil {
		log.Panicf("There was a problem reading file: %s", err)
	}

	outputFile, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	var byteCount = make([]uint32, 256)
	binary.Read(inputFile, binary.LittleEndian, &byteCount)

	var size uint32
	binary.Read(inputFile, binary.LittleEndian, &size)

	position := uint32(0)
	aux := byte(0)

	fmt.Println("Decompressing...")

	root := BuildHuffmanTree(byteCount)

	bar := pb.StartNew(int(size) + 1)

	bufferedWriter := bufio.NewWriter(outputFile)
	bufferedReader := bufio.NewReader(inputFile)

	for position < size {

		currentNode := root
		for (currentNode.left != nil) || (currentNode.right != nil) {
			if generateBit(bufferedReader, position, &aux) {
				currentNode = currentNode.right
				position = position + 1
				bar.Increment()

			} else {
				currentNode = currentNode.left
				position = position + 1
				bar.Increment()

			}

		}
		bufferedWriter.WriteByte(currentNode.c)
	}
	bufferedWriter.Flush()

	outputFile.Close()
	inputFile.Close()
}

func generateBit(inReader *bufio.Reader, position uint32, aux *byte) bool {

	if (position % 8) == 0 {
		binary.Read(inReader, binary.LittleEndian, aux)
	}

	if (*aux)&(1<<(position%8)) != 0 {
		return true
	}
	return false

}
