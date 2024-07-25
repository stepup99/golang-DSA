package treePkg

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

func NewNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func Insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return NewNode(val)
	} else {
		if val <= root.Val {
			root.Left = Insert(root.Left, val)
		} else {
			root.Right = Insert(root.Right, val)
		}
		return root
	}
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
	_, err := json.Marshal(root)
	if err != nil {
		log.Fatal(err)
	}
	return root
}

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func InOrder(root *TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Println(root.Val)
	InOrder(root.Right)

}

func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Println(root.Val)
}

func LevelOrder(root *TreeNode) {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Printf("%d..", cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
}

func LevelOrder2(root *TreeNode) {
	queue := []*TreeNode{root}
	var ans [][]int
	for len(queue) > 0 {

		sz := len(queue)

		var temp []int

		for i := 0; i < sz; i++ {

			cur := queue[0]

			queue = queue[1:]

			temp = append(temp, cur.Val)

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}

		}
		ans = append(ans, temp)
	}

	fmt.Println(ans)
}
