package gulc

import (
	"math/rand"
	"time"
)

const (
	MAX_HEIGHT=32
	Rand_LEVEL_P = 0.25
)

type SkiplistNodeLevel struct {
	level int
	next *SkiplistNode   // 指向下一个节点
}


type SkiplistNode struct {
	val int
	prev *SkiplistNode   // 指向前一个节点
	level int
	levels []*SkiplistNodeLevel
}


func NewSkiplistNode(val int, level int) *SkiplistNode {
	node := &SkiplistNode{val: val, levels: make([]*SkiplistNodeLevel, level), level: level}
	for i := 0; i < level; i++ {
		node.levels[i] = &SkiplistNodeLevel{level: i}
	}
	return node
}


func NewSkiplistNodeNoVal(level int) *SkiplistNode {
	node := &SkiplistNode{levels: make([]*SkiplistNodeLevel, level), level: level}
	for i := 0; i < level; i++ {
		node.levels[i] = &SkiplistNodeLevel{level: i}
	}
	return node
}


type Skiplist struct {
	level int   // 除伪头部外最高层次
	length int // 元素数量
	head, tail *SkiplistNode
}


func NewSkiplist() *Skiplist {
	sl := &Skiplist{level: 1, head: NewSkiplistNodeNoVal(MAX_HEIGHT)}
	return sl
}

func (s *Skiplist) Search(target int) bool {
	var cur *SkiplistNode = s.head
	for i := s.level - 1; i >= 0; i-- {
		for cur.levels[i].next != nil && cur.levels[i].next.val < target {
			cur = cur.levels[i].next
		}
	}
	if cur.levels[0].next != nil && cur.levels[0].next.val == target {
		return true
	}
	return false
}

func (s *Skiplist) randomLevel() int {
	rand.Seed(time.Now().UnixNano())
	level := 1
	for (level < MAX_HEIGHT && 0x7fff * rand.Float64() < 0x7fff * Rand_LEVEL_P) {
		level++
	}
	return min(level, MAX_HEIGHT)
}

func (s *Skiplist) Insert(val int) {
	var cur *SkiplistNode = s.head
	prevNodes := make([]*SkiplistNode, MAX_HEIGHT)
	for i := s.level - 1; i >= 0; i-- {
		for cur.levels[i].next != nil && cur.levels[i].next.val < val {
			cur = cur.levels[i].next
		}
		prevNodes[i] = cur
	}

	level := s.randomLevel()
	node := NewSkiplistNode(val, level)

	for i := min(level - 1, s.level - 1); i >= 0; i-- {
		node.levels[i].next = prevNodes[i].levels[i].next
		prevNodes[i].levels[i].next = node
	}
	if level > s.level {
		for i := s.level; i < level; i++ {
			s.head.levels[i].next = node
		}
		s.level = level
	}
	node.prev = prevNodes[0]
	if node.levels[0].next != nil {
		node.levels[0].next.prev = node
	} else {
		s.tail = node
	}
}


func (s *Skiplist) Delete(val int) bool {
	var cur *SkiplistNode = s.head
	prevNodes := make([]*SkiplistNode, MAX_HEIGHT)
	for i := s.level - 1; i >= 0; i-- {
		for cur.levels[i].next != nil && cur.levels[i].next.val < val {
			cur = cur.levels[i].next
		}
		prevNodes[i] = cur
	}
	if cur.levels[0].next == nil || cur.levels[0].next.val != val {
		return false
	}

	deletingNode := cur.levels[0].next

	for i := deletingNode.level - 1; i >= 0; i-- {
		prevNodes[i].levels[i].next = deletingNode.levels[i].next
	}

	if deletingNode.levels[0].next != nil {
		deletingNode.levels[0].next.prev = prevNodes[0]
	} else {
		s.tail = prevNodes[0]
	}
	return true
}