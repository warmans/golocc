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

//Result - container for analysis results
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
	Test             int
	Assertion        int
	IfStatement      int
	SwitchStatement  int
	GoStatement      int
}

//Parser - Code parser struct
type Parser struct{}

//ParseDir - Parse all files within directory
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
		loc, cloc, assertions := p.CountLOC(file.Name())
		res.LOC += loc
		res.CLOC += cloc
		res.NCLOC += (loc - cloc)
		res.Assertion += assertions
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
						if strings.HasPrefix(x.Name.String(), "Test") {
							if len(x.Type.Params.List) != 0 {
								xt := x.Type.Params.List[0].Type.(*ast.StarExpr)
								xtx := xt.X.(*ast.SelectorExpr)
								for _, validArgType := range []string{"testing.T", "testing.M", "testing.B"} {
									if fmt.Sprintf("%s.%s", xtx.X, xtx.Sel) == validArgType {
										res.Test++
									}
								}
							}
						}
					}
				} else {
					res.Method++
					if x.Name.IsExported() {
						res.ExportedMethod++
					}
				}
			case *ast.ImportSpec:
				res.Import++
			case *ast.IfStmt:
				res.IfStatement++
			case *ast.SwitchStmt:
				res.SwitchStatement++
			case *ast.GoStmt:
				res.GoStatement++
			}
			return true
		})
	}

	return res
}

//CountLOC - count lines of code, pull LOC, Comments, assertions
func (p *Parser) CountLOC(filePath string) (int, int, int) {

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	var loc int
	var cloc int
	var assertions int
	var inBlockComment bool

	assertionPrefixes := []string{
		"So(",
		"convey.So(",
		"assert.",
	}

	for {
		line, isPrefix, err := r.ReadLine()
		if err == io.EOF {
			return loc, cloc, assertions
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
		for _, prefix := range assertionPrefixes {
			if strings.HasPrefix(strings.TrimSpace(string(line)), prefix) {
				assertions++
			}
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
}

func main() {

	//default to current working dir
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current working dir: ", err.Error())
	}

	targetDir := flag.String("d", pwd, "target directory")
	outputFmt := flag.String("o", "text", "output format")
	flag.Parse()

	var report ReportInterface
	switch *outputFmt {
	case "text":
		report = &TextReport{}
	case "json":
		report = &JSONReport{}
	}

	parser := Parser{}
	report.Print(parser.ParseDir(*targetDir))
}
