package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"strings"
)

type Result struct {
	LOC              int
	CLOC             int
	NCLOC            int
	Struct           int
	Interface        int
	Method           int
	ExportedMethod   int
	Function         int
	ExportedFunction int
	Import           int
}

type Parser struct{}

func (p *Parser) ParseDir(targetDir string) *Result {

	res := &Result{}

	//create the file set
	fset := token.NewFileSet()
	d, err := parser.ParseDir(fset, targetDir, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	//count up lines
	fset.Iterate(func(file *token.File) bool {
		loc, cloc := p.CountLOC(file.Name())
		res.LOC += loc
		res.CLOC += cloc
		res.NCLOC += (loc - cloc)
		return true
	})

	//count entities
	for _, pkg := range d {

		ast.Inspect(pkg, func(n ast.Node) bool {

			switch x := n.(type) {
			case *ast.StructType:
				res.Struct++
			case *ast.InterfaceType:
				res.Interface++
			case *ast.FuncDecl:
				if x.Recv == nil {
					res.Function++
					if x.Name.IsExported() {
						res.ExportedFunction++
					}
				} else {
					res.Method++
					if x.Name.IsExported() {
						res.ExportedMethod++
					}
				}
			case *ast.ImportSpec:
				res.Import++
			}
			return true
		})
	}

	return res
}

func (p *Parser) CountLOC(filePath string) (int, int) {

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	var loc int
	var cloc int
	var inBlockComment bool

	for {
		line, isPrefix, err := r.ReadLine()
		if err == io.EOF {
			return loc, cloc
		}
		if isPrefix == true {
			continue //incomplete line
		}
		if len(line) == 0 {
			continue //empty line
		}
		if strings.Index(strings.TrimSpace(string(line)), "//") == 0 {
			cloc++ //slash comment at start of line
			continue
		}

		blockCommentStartPos := strings.Index(strings.TrimSpace(string(line)), "/*")
		blockCommentEndPos := strings.LastIndex(strings.TrimSpace(string(line)), "*/")

		if blockCommentStartPos > -1 {
			//block was started and not terminated
			if blockCommentEndPos == -1 || blockCommentStartPos > blockCommentEndPos {
				inBlockComment = true
			}
		}
		if blockCommentEndPos > -1 {
			//end of block is found and no new block was started
			if blockCommentStartPos == -1 || blockCommentEndPos > blockCommentStartPos {
				inBlockComment = false
				cloc++ //end of block counts as a comment line but we're already out of the block
			}
		}

		loc++
		if inBlockComment {
			cloc++
		}
	}

	return loc, cloc
}

type TextReport struct {

}

func (t *TextReport) Print(res *Result) {
	fmt.Printf("LOC:        %v (%v CLOC, %v NCLOC)\n", res.LOC, res.CLOC, res.NCLOC)
	fmt.Printf("Imports:    %v\n", res.Import)
	fmt.Printf("Structs:    %v\n", res.Struct)
	fmt.Printf("Interfaces: %v\n", res.Interface)
	fmt.Printf("Methods:    %v (%v Exported)\n", res.Method, res.ExportedMethod)
	fmt.Printf("Functions:  %v (%v Exported)\n", res.Function, res.ExportedFunction)
}

func main() {

	//default to current working dir
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current working dir: ", err.Error())
	}

	targetDir := flag.String("d", pwd, "target directory")
	flag.Parse()

	fmt.Println("Parsing dir: ", *targetDir)

	parser := Parser{}
	result := parser.ParseDir(*targetDir)

	report := &TextReport{}
	report.Print(result)
}
