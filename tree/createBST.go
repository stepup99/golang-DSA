package main

import (
	"fmt"
	treehelper "myPackage/src/treeHelper"
)

func main() {
	tree := treehelper.BuildTree()
	// treehelper.PreOrder(tree)
	// fmt.Println()
	// treehelper.InOrder(tree)
	// fmt.Println()
	// treehelper.PostOrder(tree)
	fmt.Println("hi >>>>>>> ")
	treehelper.LevelOrder2(tree)
}
