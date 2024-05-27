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

func (sb *Formatter) getPreceedingText(d int) string {
	sb.Reset()

	for i := 0; i < d; i++ {
		sb.WriteString("|\t")
	}
	sb.WriteString("|__>")

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
		if item.isDir == true && string(item.baseName[0]) != string(".") {
			tabs := sb.getPreceedingText(item.depth)
			fmt.Println(tabs, item.getFullPath())
		} else {
			tabs := sb.getPreceedingText(item.depth)
			fmt.Println(tabs, "", item.getFullPath())
		}
	}
}

type Node struct {
	baseName   string
	parentName string
	files      []*Node
	subDirs    []*Node
	depth      int
	isDir      bool
}

func (n *Node) getFullPath() string {
	return n.parentName + n.baseName
}

func NewNode(basename, parentName string, depth int, isDir bool) *Node {
	files := []*Node{}
	subDirs := []*Node{}
	return &Node{
		baseName:   basename,
		parentName: parentName,
		files:      files,
		subDirs:    subDirs,
		depth:      depth,
		isDir:      isDir,
	}
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

	root := NewNode(os.Args[1], "", 0, true)
	items, err := os.ReadDir(root.baseName)

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
			parentDir := parent.baseName + "/"
			subDirNode := NewNode(item.Name(), parentDir, depth, true)
			subDirFiles, _ := os.ReadDir(parentDir + item.Name())
			parent.subDirs = append(parent.subDirs, subDirNode)
			createTree(subDirFiles, subDirNode)
		} else {
			parent.files = append(parent.files, NewNode(item.Name(), parent.baseName+"/", depth, false))
		}
	}
}
