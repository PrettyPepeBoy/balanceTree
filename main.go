package main

import (
	"dima/tree"
)

func main() {
	t := tree.CreateTree(8)
	t.Put(6)
	t.Put(10)
	t.Put(14)
	t.Put(22)
	t.Print()

}
