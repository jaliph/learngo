// https://leetcode.com/problems/construct-quad-tree/

package trees

/**
 * Definition for a QuadTree node.
 */
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func Construct(grid [][]int) *Node {
	n := len(grid)

	var flag bool
	var creatNode func(r, c, n int) *Node
	creatNode = func(r, c, n int) *Node {
		flag = true
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if grid[r][c] != grid[i+r][j+c] {
					flag = false
				}
			}
			if !flag {
				break
			}
		}

		if flag {
			return &Node{Val: grid[r][c] == 1, IsLeaf: true}
		} else {
			n = n >> 1
			return &Node{
				Val:         false,
				IsLeaf:      false,
				TopLeft:     creatNode(r, c, n),
				TopRight:    creatNode(r, c+n, n),
				BottomLeft:  creatNode(r+n, c, n),
				BottomRight: creatNode(r+n, c+n, n),
			}
		}
	}

	return creatNode(0, 0, n)
}
