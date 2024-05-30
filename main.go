package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

var dir, _ = os.Getwd()
var showHiddenFiles = flag.Bool("h", false, "Show hidden files")
var selectedDir = flag.String("d", dir, "Directory to print")

type Formatter struct {
	strings.Builder
}

var sb = Formatter{}

func (sb *Formatter) GetIndentation(d int, isDir bool) string {
	sb.Reset()

	for i := 0; i < d; i++ {
		sb.WriteString("|\t")
	}

	if isDir == true {
		sb.WriteString("|\t")
	} else {
		sb.WriteString("|__>")
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
		//

		if item.isDir == true {
			tabs := sb.GetIndentation(item.depth, item.isDir)
			fmt.Println(tabs, item.FullPath())
		} else {
			tabs := sb.GetIndentation(item.depth, item.isDir)
			fmt.Println(tabs, "", item.baseName)
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

func (n *Node) FullPath() string {
	if n.isDir == true {
		//TODO: may need a / inserted in the middle here.
		return n.parentName + n.baseName + "/"
	} else {
		return n.parentName + n.baseName
	}

}

func NewNode(parentName, baseName string, depth int, isDir bool) *Node {
	files := []*Node{}
	subDirs := []*Node{}
	return &Node{
		parentName: parentName,
		baseName:   baseName,
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
	var root *Node

	flag.Parse()

	fmt.Println(*selectedDir)
	root = NewNode(*selectedDir, "", 0, true)

	items, err := os.ReadDir(root.FullPath())

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
			//create node for subdirectory
			if *showHiddenFiles == false && item.Name()[0] == byte('.') {
				continue
			}

			subDirNode := NewNode(parent.FullPath(), item.Name(), depth, true)

			//append node to parents subsdirectories
			parent.subDirs = append(parent.subDirs, subDirNode)

			//get all files within the subdirectory.
			subDirFiles, _ := os.ReadDir(subDirNode.FullPath())

			//recurse into subdirectory
			createTree(subDirFiles, subDirNode)
		} else {
			if *showHiddenFiles == false && item.Name()[0] != byte('.') {
				parent.files = append(parent.files, NewNode(parent.baseName+"/", item.Name(), depth, false))
			}
		}
	}
}
