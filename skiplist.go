package main

import (
	"fmt"
	"math/rand"
)

const MaxLevel = 20
const MaxCountOfSub = 32

type Node struct {
	Forward []Node
	Value   int
}

func NewNode(v, level int) *Node {
	return &Node{Value: v, Forward: make([]Node, level)}
}

type SkipList struct {
	Header *Node
	Level  int
}

func main() {
	slt := NewSkipList()
	for i := 100; i > 0; i-- {
		slt.Insert(i)
	}
	slt.PrintSkipList()
	slt.Search(15)
	slt.Search(93)
	slt.Remove(93)
	slt.PrintSkipList()
}
func NewSkipList() *SkipList {
	return &SkipList{Level: 0, Header: NewNode(0, MaxCountOfSub)}
}

func (skipList *SkipList) Insert(key int) {

	update := make(map[int]*Node)
	node := skipList.Header

	for i := skipList.Level; i >= 0; i-- {
		for {
			if node.Forward[i].Value != 0 && node.Forward[i].Value < key {
				node = &node.Forward[i]
			} else {
				break
			}
		}
		update[i] = node
	}

	level := skipList.RandomLevel()
	if level > skipList.Level {
		for i := skipList.Level; i < level; i++ {
			update[i] = skipList.Header
		}
		skipList.Level = level
	}

	newNode := NewNode(key, MaxCountOfSub)

	for i := 0; i < level; i++ {
		newNode.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = *newNode
	}

}

func (skipList *SkipList) RandomLevel() int {
	return rand.Intn(MaxLevel + 1)
}

func (skipList *SkipList) PrintSkipList() {

	fmt.Println("SkipList-------------------------------------------")
	for i := MaxLevel - 1; i >= 0; i-- {

		fmt.Println("level:", i)
		node := skipList.Header.Forward[i]
		for {
			if node.Value != 0 {
				fmt.Printf("%d->", node.Value)
				node = node.Forward[i]
			} else {
				break
			}
		}
		fmt.Println("--------------------------------------------------------")
	} //end for

	fmt.Println("Current MaxLevel:", skipList.Level)
}

func (skipList *SkipList) Search(key int) *Node {

	node := skipList.Header
	for i := skipList.Level - 1; i >= 0; i-- {

		fmt.Printf("Search Level=%d ...\n", i)
		for {
			if node.Forward[i].Value == 0 {
				break
			}

			if node.Forward[i].Value == key {
				fmt.Printf("\nkey=%d found in level=%d \n", key, i)
				return &node.Forward[i]
			}

			if node.Forward[i].Value < key {
				node = &node.Forward[i]
				continue
			} else { // > key
				break
			}
		} //end for find

	} //end level
	return nil
}

func (skipList *SkipList) Remove(key int) {

	update := make(map[int]*Node)
	node := skipList.Header
	for i := skipList.Level - 1; i >= 0; i-- {

		for {

			if node.Forward[i].Value == 0 {
				break
			}

			if node.Forward[i].Value == key {
				update[i] = node
				break
			}

			if node.Forward[i].Value < key {
				node = &node.Forward[i]
				continue
			} else { // > key
				break
			}

		} //end for find

	} //end level

	for i, v := range update {
		if v == skipList.Header {
			skipList.Level--
		}
		v.Forward[i] = v.Forward[i].Forward[i]
	}
}
