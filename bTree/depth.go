/*
Create a function to get tree depth
*/

package bTree


//GetDepth will return tree depth
func GetDepth(root *Tree) (depth int) {
	if root == nil {
		return
	}
	return 1 + max(GetDepth(root.Left), GetDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
