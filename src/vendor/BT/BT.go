package BT

type BTNode struct {
	Val   int
	Right *BTNode
	Left  *BTNode
}

type BT struct {
}

func (r *BT) PreOrderTraversal(root *BTNode) (v []int) {

	if root == nil {
		return v
	}

	if root.Left != nil {
		v = append(v, r.PreOrderTraversal(root.Left)...)
	}
	if root.Right != nil {
		v = append(v, r.PreOrderTraversal(root.Right)...)
	}
	return append(v, root.Val)

}
