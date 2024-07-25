package main

import (
	treePkg "myPackage/practice-tree/treePkg"
)

func main() {

	root := treePkg.BuildTree()
	//treePkg.PreOrder(root)
	// treePkg.InOrder(root)
	// fmt.Println(">>>>>>>>>>>>>>----------- ")
	// treePkg.PostOrder(root)
	treePkg.LevelOrder2(root)
	// fmt.Println(root)
}
