package main

type 二叉树 struct {
	LeftTree  *二叉树
	Value     int
	RightTree *二叉树
}

func main() {
	root := 二叉树{Value: 1}
	print(root.Value)
}
