package bTree_test

import (
	"testing"

	"github.com/mayur-tolexo/datastruct/bTree"
	"github.com/stretchr/testify/assert"
)

//TestGetDepth will test Btree GetDepth func
func TestGetDepth(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, bTree.GetDepth(nil))

	//only root
	root := bTree.NewNode(1)
	assert.Equal(1, bTree.GetDepth(root))

	root = bTree.NewNode(1)
	root.Left = bTree.NewNode(2)
	root.Right = bTree.NewNode(3)
	root.Left.Left = bTree.NewNode(4)
	assert.Equal(3, bTree.GetDepth(root))
}
