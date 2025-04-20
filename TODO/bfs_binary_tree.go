package main

import "fmt"

// To implement here
func bfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

// func levelOrderTraverse1(root *TreeNode) {
//     if root == nil {
//         return
//     }
//     q := []*TreeNode{}
//     q = append(q, root)
//     for len(q) > 0 {
//         cur := q[0]
//         q = q[1:]
//         // 访问 cur 节点
//         fmt.Println(cur.Val)

//         // 把 cur 的左右子节点加入队列
//         if cur.Left != nil {
//             q = append(q, cur.Left)
//         }
//         if cur.Right != nil {
//             q = append(q, cur.Right)
//         }
//     }
// }

// func levelOrderTraverse2(root *TreeNode) {
//     if root == nil {
//         return
//     }
//     q := []*TreeNode{root}
//     // 记录当前遍历到的层数（根节点视为第 1 层）
//     depth := 1

//     for len(q) > 0 {
//         // 获取当前队列长度
//         sz := len(q)
//         for i := 0; i < sz; i++ {
//             // 弹出队列头
//             cur := q[0]
//             q = q[1:]
//             // 访问 cur 节点，同时知道它所在的层数
//             fmt.Printf("depth = %d, val = %d\n", depth, cur.Val)

//             // 把 cur 的左右子节点加入队列
//             if cur.Left != nil {
//                 q = append(q, cur.Left)
//             }
//             if cur.Right != nil {
//                 q = append(q, cur.Right)
//             }
//         }
//         depth++
//     }
// }

// type State struct {
//     node  *TreeNode
//     depth int
// }

// func levelOrderTraverse3(root *TreeNode) {
//     if root == nil {
//         return
//     }
//     q := []State{{root, 1}}

//     for len(q) > 0 {
//         cur := q[0]
//         q = q[1:]
//         // 访问 cur 节点，同时知道它的路径权重和
//         fmt.Printf("depth = %d, val = %d\n", cur.depth, cur.node.Val)

//         // 把 cur 的左右子节点加入队列
//         if cur.node.Left != nil {
//             q = append(q, State{cur.node.Left, cur.depth + 1})
//         }
//         if cur.node.Right != nil {
//             q = append(q, State{cur.node.Right, cur.depth + 1})
//         }
//     }
// }

func print_bfs() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	result := bfs(root)
	fmt.Println(result) // Output: [1 2 3 4 5 6 7]
}
