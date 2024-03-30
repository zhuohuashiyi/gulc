package gulc

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


type Point struct {
	X, Y int
}


type Rectangle struct {
	Min, Max Point
}