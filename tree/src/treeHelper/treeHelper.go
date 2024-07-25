package treehelper

import (
	"encoding/json"
	"fmt"
	"log"
)

type TreeNode struct {
	Val   int       `json:"val"`
	Left  *TreeNode `json:"left"`
	Right *TreeNode `json:"right"`
}

// create a new tree node (instatiating a tree node object)
func NewNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func Insert(treeNode *TreeNode, val int) *TreeNode {
	if treeNode == nil {
		return NewNode(val)
	} else {
		if val <= treeNode.Val {
			treeNode.Left = Insert(treeNode.Left, val)
		} else {
			treeNode.Right = Insert(treeNode.Right, val)
		}
	}
	return treeNode
}

func BuildTree() *TreeNode {
	var root *TreeNode
	root = Insert(root, 21)
	root = Insert(root, 2)
	root = Insert(root, 1)
	root = Insert(root, 6)
	root = Insert(root, 4)
	root = Insert(root, 11)
	root = Insert(root, 15)
	root = Insert(root, 23)
	root = Insert(root, 43)
	root = Insert(root, 3)

	data, err := json.Marshal(root)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data))

	return root

}

/*

      21
      /  \
     2    23
    / \     \
   1   6    43
      / \
     4  11
    /     \
   3       15

*/

// preorder

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("..%d..", root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func InOrder(root *TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Printf("..%d..", root.Val)
	InOrder(root.Right)
}

func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Printf("..%d..", root.Val)
}

func LevelOrder(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {

		current := queue[0]

		queue = queue[1:]

		fmt.Printf("%d...", current.Val)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

}

// [
// 	[21],
// 	[2, 23],
// 	[1, 6, 43]
// 	[4, 11]
// 	[3 ,15]
// ]

func LevelOrder2(root *TreeNode) {
	if root == nil {
		fmt.Println([][]int{})
		return
	}
	queue := []*TreeNode{root}
	var ans [][]int
	for len(queue) > 0 {

		var temp []int

		sz := len(queue)

		for i := 0; i < sz; i++ {
			current := queue[0]

			queue = queue[1:]

			temp = append(temp, current.Val)

			if current.Left != nil {
				queue = append(queue, current.Left)
			}

			if current.Right != nil {
				queue = append(queue, current.Right)
			}

		}

		ans = append(ans, temp)
	}
	fmt.Println(ans)
}
