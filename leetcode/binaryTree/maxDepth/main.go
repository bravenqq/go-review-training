// Package main provides ...
package main

import "fmt"

/*给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。*/

/*
* 思路一：使用广度优先遍历遍历每一层，记录遍历的层数即可
* 时间复杂度：O(n)
* 空间复杂度：线性表最差O(n),二叉树完全平衡最好O(log n)
 */
func main() {
	root := &TreeNode{Val: 3}
	fmt.Println(maxDepth(root))

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var depth int

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	currentLists := []*TreeNode{root}
	getMaxDepth(currentLists)
	return depth
}
func getMaxDepth(current []*TreeNode) {
	if len(current) == 0 {
		return
	}
	depth++
	var nextLists []*TreeNode
	for _, node := range current {
		if node.Left != nil {
			nextLists = append(nextLists, node.Left)
		}
		if node.Right != nil {
			nextLists = append(nextLists, node.Right)
		}
	}
	fmt.Println("nextLists:", len(nextLists))
	getMaxDepth(nextLists)
}
