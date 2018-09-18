#### File Compressor and Decompressor written in [Go](https://golang.org)

As part of my explorings with the Go language I decided to write a simple File Compressor, for it is the kind of project in which [buffered reading, writing and seeking](https://www.devdungeon.com/content/working-files-go#write_buffered) are staples (otherwise the compressed data may not preserve it's original state). Plus, I had to work with Go's very abstract way of dealing with pointers, errors and loops.

I also played around with [this neat progress bar package](https://gopkg.in/cheggaaa/pb.v1).

##### A little about Huffman Coding

This [entropy encoding](https://en.wikipedia.org/wiki/Entropy_encoding) technique consists in creating a binary tree representation of the input data that can be stored and rebuilt later on. 

1. A list containing the symbols frequency (in case of files, symbols are bytes) is built.
2. The list is then sorted. This can be done through a [Heap](https://en.wikipedia.org/wiki/Heap_(data_structure)).
3. The following steps are to be repeated until there's no symbol left:
   1. Get the two symbols of smaller frequency from the list.
   2. Create a tree containing the two elements as children nodes. 
   3. Create a parent node storing the sum of two children elements frequency. 
   4. Add the parent element to the list, that must, after the addition, still have its order preserved.
   5. Delete the children nodes.
4. A code word is then assigned to each element based on its path out of the root.

##### Usage

###### Compiling

```shell
$ go build main
```

###### Running

```shell
$ ./main -c uncompressed compressed
$ ./main -x compressed uncompressed
```

