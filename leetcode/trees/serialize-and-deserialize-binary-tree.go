package trees

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	res := []string{}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			res = append(res, "N")
			return
		}
		res = append(res, fmt.Sprintf("%d", node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	res := strings.Split(data, ",")
	i := 0
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if res[i] == "N" {
			i++
			return nil
		}
		num, _ := strconv.Atoi(res[i])
		curr := &TreeNode{Val: num}
		i++
		curr.Left = dfs()
		curr.Right = dfs()
		return curr
	}
	return dfs()
}

/**
* Your Codec object will be instantiated and called as such:
* ser := Constructor();
* deser := Constructor();
* data := ser.serialize(root);
* ans := deser.deserialize(data);
 */
