package main

import (
	"fmt"
)

// TreeNode 定义二叉树的节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrderTraverse 层序遍历函数
func levelOrderTraverse(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var currentLevel []int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentLevel)
	}

	return result
}

// main 函数用于测试
func main() {
	// 构建一个简单的二叉树
	//       1
	//      / \
	//     2   3
	//    / \   \
	//   4   5   6
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	// 调用层序遍历函数
	result := levelOrderTraverse(root)

	// 打印结果
	fmt.Println(result) // 输出: [[1], [2, 3], [4, 5, 6]]
}
