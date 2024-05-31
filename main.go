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
var maxDepth = flag.Int("D", 10, "Depth to traverse")
var filePointer = flag.String("file-pointer", "|−-", "String used to point to files")
var lineSeparator = flag.String("separator", "|", "String used to separate lines")
var dirColor = flag.String("directory-color", "blue", "Color to print directories")

type Formatter struct {
	strings.Builder
}

var sb = Formatter{}

func (sb *Formatter) GetIndentation(depth int, isDir bool) string {
	sb.Reset()

	for i := 0; i <= depth; i++ {
		if isDir == true && i == depth {
			sb.WriteString("∟")
		} else {
			sb.WriteString(*lineSeparator)
			sb.WriteString("\t")
		}
	}

	if isDir == false {
		sb.WriteString(*filePointer)
	}

	return sb.String()
}

func (sb *Formatter) ColorString(text string, color string) string {
	end := "\033[0m"

	switch color {
	case "black":
		return fmt.Sprint("\033[30m", text, end)
	case "red":
		return fmt.Sprint("\033[31m", text, end)
	case "green":
		return fmt.Sprint("\033[32m", text, end)
	case "yellow":
		return fmt.Sprint("\033[33m", text, end)
	case "blue":
		return fmt.Sprint("\033[34m", text, end)
	case "magenta":
		return fmt.Sprint("\033[35m", text, end)
	case "cyan":
		return fmt.Sprint("\033[36m", text, end)
	case "white":
		return fmt.Sprint("\033[37m", text, end)
	case "bright-black":
		return fmt.Sprint("\033[90m", text, end)
	case "bright-red":
		return fmt.Sprint("\033[91m", text, end)
	case "bright-green":
		return fmt.Sprint("\033[92m", text, end)
	case "bright-yellow":
		return fmt.Sprint("\033[93m", text, end)
	case "bright-blue":
		return fmt.Sprint("\033[94m", text, end)
	case "bright-magenta":
		return fmt.Sprint("\033[95m", text, end)
	case "bright-cyan":
		return fmt.Sprint("\033[96m", text, end)
	case "bright-white":
		return fmt.Sprint("\033[97m", text, end)
	default:
		return text
	}
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
		indent := sb.GetIndentation(item.depth, item.isDir)

		if item.isDir == true {
			fmt.Println(indent, sb.ColorString(item.FullPath(), *dirColor))
		} else {
			fmt.Println(indent, "", item.baseName)
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

		if depth > *maxDepth {
			break
		}

		if isDir := item.IsDir(); isDir == true {
			//create node for subdirectory
			if *showHiddenFiles == false && item.Name()[0] == byte('.') {
				continue
			}

			subDirNode := NewNode(parent.FullPath(), item.Name(), depth, true)

			//append node to parents subdirectories
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
