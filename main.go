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
	Tests            int
	Assertions       int
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
		res.Assertions += assertions
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
								nodePos := fset.Position(x.Type.Params.List[0].Type.Pos())
								nodeEnd := fset.Position(x.Type.Params.List[0].Type.End())
								nodeFile, _ := os.Open(nodePos.Filename)
								defer nodeFile.Close()
								node := make([]byte, (nodeEnd.Offset - nodePos.Offset))
								nodeFile.ReadAt(node, int64(nodePos.Offset))
								paramTypes := []string{
									"*testing.T",
									"*testing.M",
									"*testing.B",
								}
								for _, paramType := range paramTypes {
									if string(node) == paramType {
										res.Tests++
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

//ReportInterface - reports that parse results and print out a report
type ReportInterface interface {
	Print(*Result)
}

//JSONReport json structure for LOC report
type JSONReport struct {
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

//Print - print out parsed report in json format
func (j *JSONReport) Print(res *Result) {
	j.LOC.CLOC = res.CLOC
	j.LOC.NCLOC = res.NCLOC
	j.Imports = res.Import
	j.Structs = res.Struct
	j.Interfaces = res.Interface
	j.Methods.Total = res.Method
	j.Methods.Exported = res.ExportedMethod
	j.Functions.Total = res.Function
	j.Functions.Exported = res.ExportedFunction
	j.Testing.Cases = res.Tests
	j.Testing.Assertions = res.Assertions
	jsonOutput, _ := json.MarshalIndent(j, "", "  ")
	fmt.Print(string(jsonOutput))
}

//TextReport - plaintext report output
type TextReport struct {
}

//Print - print out plaintext report
func (t *TextReport) Print(res *Result) {
	fmt.Printf("LOC:        %v (%v CLOC, %v NCLOC)\n", res.LOC, res.CLOC, res.NCLOC)
	fmt.Printf("Imports:    %v\n", res.Import)
	fmt.Printf("Structs:    %v\n", res.Struct)
	fmt.Printf("Interfaces: %v\n", res.Interface)
	fmt.Printf("Methods:    %v (%v Exported)\n", res.Method, res.ExportedMethod)
	fmt.Printf("Functions:  %v (%v Exported)\n", res.Function, res.ExportedFunction)
	fmt.Printf("Tests:      %v \n", res.Tests)
	fmt.Printf("Assertions: %v \n", res.Assertions)
}

func main() {

	//default to current working dir
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current working dir: ", err.Error())
	}

	targetDir := flag.String("d", pwd, "target directory")
	outputFmt := flag.String("o", "text", "output format")
	flag.Parse()

	fmt.Println("Parsing dir: ", *targetDir)

	parser := Parser{}
	result := parser.ParseDir(*targetDir)
	var report ReportInterface
	switch *outputFmt {
	case "text":
		report = &TextReport{}
	case "json":
		report = &JSONReport{}
	}
	report.Print(result)
}
