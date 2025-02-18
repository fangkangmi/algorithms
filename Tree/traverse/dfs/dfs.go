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

func levelOrderTraverse(root *TreeNode) [][]int {
	// base case
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	dfs(root, 0, &result)
	return result
}

func dfs(node *TreeNode, level int, result *[][]int) {
	if node == nil {
		return
	}

	// Ensure the result slice has enough levels
	if len(*result) == level {
		*result = append(*result, []int{})
	}

	// Add the current node's value to the corresponding level
	(*result)[level] = append((*result)[level], node.Val)

	// Recursively traverse the left and right subtrees
	dfs(node.Left, level+1, result)
	dfs(node.Right, level+1, result)
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
