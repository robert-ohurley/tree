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
var dirOnly = flag.Bool("dir-only", false, "Only print directories")
var maxDepth = flag.Int("D", 10, "Depth to traverse")
var filePointer = flag.String("file-pointer", "|−-", "String used to point to files")
var lineSeparator = flag.String("separator", "|", "String used to separate lines")
var dirColor = flag.String("directory-color", "blue", "Color to print directories")
var fileColor = flag.String("file-color", "white", "Color to print files")
var showFullPath = flag.Bool("fullpath", false, "Show full path name for directories")

type Printer struct {
	strings.Builder
}

var sb = Printer{}

func (sb *Printer) FormatPath(node *Node) string {
	if node.depth == 0 || *showFullPath == true {
		return node.FullPath()
	}
	return fmt.Sprint(node.baseName)
}

func (sb *Printer) GetIndentation(depth int, isDir bool) string {
	sb.Reset()

	for i := 0; i < depth; i++ {
		if isDir == true && i+1 == depth {
			sb.WriteString(*lineSeparator)
			sb.WriteString("\t")
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

func (sb *Printer) ColorString(text string, color string) string {
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

func (n *Node) Print() {
	indent := sb.GetIndentation(n.depth, n.isDir)

	if n.isDir == true {
		fmt.Println(indent, sb.ColorString(sb.FormatPath(n), *dirColor))
	} else {
		fmt.Println(indent, "", sb.ColorString(sb.FormatPath(n), *fileColor))
	}
}

func dfs(node *Node) {
	node.Print()

	for _, item := range node.subDirs {
		dfs(item)
	}

	for _, item := range node.files {
		item.Print()
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
}

func createTree(items []fs.DirEntry, parent *Node) {
	for _, item := range items {

		//don't recurse greater that the max depth
		if parent.depth+1 == *maxDepth {
			break
		}

		if isDir := item.IsDir(); isDir == true {
			depth := parent.depth + 1
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
		} else if *dirOnly == false {
			//depth remains same for files
			depth := parent.depth

			if *showHiddenFiles == false && item.Name()[0] == byte('.') {
				continue
			} else {
				parent.files = append(parent.files, NewNode("/", item.Name(), depth, false))
			}
		}
	}
}
