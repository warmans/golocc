package main

import (
	"bufio"
	"encoding/json"
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
	Tests            int
	TLOC             int
	Assertions       int
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
		loc, cloc, assertions, tests := p.CountLOC(file.Name())
		res.LOC += loc
		res.CLOC += cloc
		res.NCLOC += (loc - cloc)
		res.Assertions += assertions
		res.Tests += tests
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
							nodePos := fset.Position(x.Type.Params.List[0].Type.Pos())
							nodeEnd := fset.Position(x.Type.Params.List[0].Type.End())
							nodeFile, _ := os.Open(nodePos.Filename)
							defer nodeFile.Close()
							node := make([]byte, (nodeEnd.Offset - nodePos.Offset))
							nodeFile.ReadAt(node, int64(nodePos.Offset))
							paramTypes := []string{
								"*testing.T",
								"*testing.M",
							}
							for _, paramType := range paramTypes {
								if string(node) == paramType {
									res.Tests++
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
			}
			return true
		})
	}

	return res
}

func (p *Parser) CountLOC(filePath string) (int, int, int, int) {

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	var loc int
	var cloc int
	var assertions int
	var tests int
	var inBlockComment bool

	assertionPrefixes := []string{
		"So(",
		"convey.So(",
		"assert.",
	}

	for {
		line, isPrefix, err := r.ReadLine()
		if err == io.EOF {
			return loc, cloc, assertions, tests
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

	return loc, cloc, assertions, tests
}

type TextReport struct {
}

//JsonReport json structure for LOC report
type JsonReport struct {
	LOC struct {
		CLOC  int
		NCLOC int
	}
	Imports    int
	Structs    int
	Interfaces int
	Methods    struct {
		Total    int
		Exported int
	}
	Functions struct {
		Total    int
		Exported int
	}
	Testing struct {
		Cases      int
		Assertions int
	}
}

func (t *TextReport) Print(res *Result, outputFmt string) {
	switch outputFmt {
	case "plain":
		fmt.Printf("LOC:        %v (%v CLOC, %v NCLOC)\n", res.LOC, res.CLOC, res.NCLOC)
		fmt.Printf("Imports:    %v\n", res.Import)
		fmt.Printf("Structs:    %v\n", res.Struct)
		fmt.Printf("Interfaces: %v\n", res.Interface)
		fmt.Printf("Methods:    %v (%v Exported)\n", res.Method, res.ExportedMethod)
		fmt.Printf("Functions:  %v (%v Exported)\n", res.Function, res.ExportedFunction)
		fmt.Printf("Tests:      %v \n", res.Tests)
		fmt.Printf("Assertions: %v \n", res.Assertions)
	case "json":
		report := JsonReport{}
		report.LOC.CLOC = res.CLOC
		report.LOC.NCLOC = res.NCLOC
		report.Imports = res.Import
		report.Structs = res.Struct
		report.Interfaces = res.Interface
		report.Methods.Total = res.Method
		report.Methods.Exported = res.ExportedMethod
		report.Functions.Total = res.Function
		report.Functions.Exported = res.ExportedFunction
		report.Testing.Cases = res.Tests
		report.Testing.Assertions = res.Assertions
		jsonOutput, _ := json.MarshalIndent(report, "", "  ")
		fmt.Print(string(jsonOutput))
	}
}

func main() {

	//default to current working dir
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current working dir: ", err.Error())
	}

	targetDir := flag.String("d", pwd, "target directory")
	outputFmt := flag.String("o", "plain", "output format")
	flag.Parse()

	fmt.Println("Parsing dir: ", *targetDir)

	parser := Parser{}
	result := parser.ParseDir(*targetDir)

	report := &TextReport{}
	report.Print(result, *outputFmt)
}
