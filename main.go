package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

type Formatter struct {
	strings.Builder
}

var sb = Formatter{}

func (sb *Formatter) getTabs(n *Node) string {
	sb.Reset()

	sb.WriteString("|")
	for i := 0; i <= n.depth; i++ {
		sb.WriteString(" ")
	}

	return sb.String()
}

type Stack struct {
	stack []*Node
}

var stack = Stack{}

func (s *Stack) push(n *Node) {
	s.stack = append(s.stack, n)
}

func (s *Stack) pop() *Node {
	n := s.stack[len(s.stack)-1]
	s.stack = s.stack[0 : len(s.stack)-1]
	return n
}

func (s *Stack) print() {
	for _, item := range s.stack {
		if item.isDir == true {
			fmt.Println(item.name)
			fmt.Print("\n")
		} else {
			tabs := sb.getTabs(item)
			fmt.Println(tabs, item.name)
			fmt.Print("\n")
		}
	}
}

type Node struct {
	name    string
	files   []*Node
	subDirs []*Node
	depth   int
	isDir   bool
}

func NewNode(name string, depth int, isDir bool) *Node {
	files := []*Node{}
	subDirs := []*Node{}
	return &Node{name, files, subDirs, depth, isDir}
}

func dfs(node *Node) {
	stack.push(node)

	for _, item := range node.subDirs {
		dfs(item)
	}

	for _, item := range node.files {
		stack.push(item)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please provide a directory")
	}
	root := NewNode(os.Args[1], 0, true)

	items, err := os.ReadDir(root.name)

	if err != nil {
		fmt.Println(err)
	}

	createTree(items, root)
	dfs(root)
	stack.print()
}

func createTree(items []fs.DirEntry, parent *Node) {
	for _, item := range items {
		depth := parent.depth + 1

		if isDir := item.IsDir(); isDir == true {
			subDirName := parent.name + "/" + item.Name()
			subDirNode := NewNode(subDirName, depth, true)
			subDirFiles, _ := os.ReadDir(subDirName)
			parent.subDirs = append(parent.subDirs, subDirNode)
			createTree(subDirFiles, subDirNode)
		} else {
			parent.files = append(parent.files, NewNode(item.Name(), depth, false))
		}
	}
}
