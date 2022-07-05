package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// 从上至下，从左至右打印二叉树
func PrintBinaryTree1(root *Node) {
	result := make([]int, 0)
	queue := make([]*Node, 0)
	if root != nil {
		queue = append(queue, root)
		for {
			if len(queue) <= 0 {
				break
			}
			curr := queue[0]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			queue = queue[1:]
			result = append(result, curr.Val)
		}
	}
	for _, item := range result {
		fmt.Printf("%d ", item)
	}
}

// 之字形打印二叉树
func PrintBinaryTree2(root *Node) {

}
