package trees

import "fmt"

// https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/
// https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/solutions/3750938/my-java-solution-99-faster/

func distanceK_BFS(root *TreeNode, target *TreeNode, k int) []int {

	g := map[int][]int{}
	var tree2GraphMaker func(node *TreeNode)
	tree2GraphMaker = func(node *TreeNode) {
		if node == nil {
			return
		}

		if _, ok := g[node.Val]; !ok {
			g[node.Val] = []int{}
		}

		if node.Left != nil {
			g[node.Val] = append(g[node.Val], node.Left.Val)
			if _, ok := g[node.Left.Val]; !ok {
				g[node.Left.Val] = []int{}
			}
			g[node.Left.Val] = append(g[node.Left.Val], node.Val)
		}

		if node.Right != nil {
			g[node.Val] = append(g[node.Val], node.Right.Val)
			if _, ok := g[node.Right.Val]; !ok {
				g[node.Right.Val] = []int{}
			}
			g[node.Right.Val] = append(g[node.Right.Val], node.Val)
		}

		tree2GraphMaker(node.Left)
		tree2GraphMaker(node.Right)
	}
	tree2GraphMaker(root)

	fmt.Println(g)

	q := [][2]int{}
	// [ [5,  0] ]

	visited := map[int]bool{}

	q = append(q, [2]int{target.Val, 0})
	visited[target.Val] = true
	res := []int{}

	for len(q) > 0 {
		curr := q[0] // [3, 1], [6, 1], [2, 1]
		q = q[1:]
		if curr[1] == k {
			res = append(res, curr[0])
		}

		for _, v := range g[curr[0]] {
			_, ok := visited[v]
			if curr[1]+1 <= k && !ok {
				q = append(q, [2]int{v, curr[1] + 1})
			}
		}
	}

	return res
}

// https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/solutions/3750834/java-solution-beats-94-full-intuition-building-and-easy-solution/
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	directions := []string{}

	res := []int{}
	directionFinder(root, target, &directions)
	fmt.Println(directions)
	upperDis(root, directions, len(directions), 0, k, &res)
	lowerDis(target, 0, k, &res)
	return res
}

func directionFinder(node *TreeNode, target *TreeNode, dir *[]string) bool {
	if node == nil {
		return false
	}

	if node == target {
		return true
	}

	if directionFinder(node.Left, target, dir) {
		*dir = append(*dir, "l")
		return true
	}
	if directionFinder(node.Right, target, dir) {
		*dir = append(*dir, "r")
		return true
	}

	return false
}

func upperDis(node *TreeNode, dirs []string, dis, flag, k int, res *[]int) {
	if node == nil {
		return
	}
	if flag == 0 {
		if dis == k {
			*res = append(*res, node.Val)
			return
		}
		if dis == k {
			return
		}

		d := dirs[dis-1]
		if d == "l" {
			upperDis(node.Left, dirs, dis-1, 0, k, res)
			upperDis(node.Right, dirs, dis+1, 1, k, res)
		} else {
			upperDis(node.Left, dirs, dis+1, 1, k, res)
			upperDis(node.Right, dirs, dis-1, 0, k, res)
		}
	} else {
		// going away from target.. so dis is incremented
		if dis == k {
			*res = append(*res, node.Val)
			return
		}
		if dis > k {
			return
		}
		upperDis(node.Left, dirs, dis+1, 1, k, res)
		upperDis(node.Right, dirs, dis+1, 1, k, res)
	}
}

func lowerDis(node *TreeNode, curr, k int, res *[]int) {
	if node == nil {
		return
	}
	if curr > k {
		return
	}

	if curr == k {
		*res = append(*res, node.Val)
		return
	}

	lowerDis(node.Left, curr+1, k, res)
	lowerDis(node.Right, curr+1, k, res)
}
