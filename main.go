package main

import (
	"fmt"
	"tykProject/MinHeap"
)

func main() {
	fmt.Println("hello world")
	heap := MinHeap.Constructor(2, []int{1, 2, 3, 4})
	heap.Add(2)
	fmt.Println("testGitCommit")
}
