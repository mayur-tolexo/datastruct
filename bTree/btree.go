package bTree

//Tree model
type Tree struct {
	Left  *Tree
	Right *Tree
	Val   int
}


//NewNode will create new btree node
func NewNode(v int) *Tree{
	return &Tree{
		Val:v,
	}
}


//Exists will check node exists or not
func Exists(t *Tree)bool {
	if t==nil{
		return false
	}
	return true
}

//LeftExists will check left node exists or not
func (t *Tree)LeftExists()bool{
	return Exists(t.Left)
}

//RightExists will check right node exists or not
func (t *Tree)RightExists()bool{
	return Exists(t.Right)
}
