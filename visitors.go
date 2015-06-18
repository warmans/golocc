package main

import (
	"go/ast"
	"strings"
	"fmt"
	"go/token"
)

type AstVisitor interface {
	Visit(ast.Node)
}

//TypeVisitor visits types to count structs, interfaces
type TypeVisitor struct {
	res *Result
}

func (v *TypeVisitor) Visit(node ast.Node) {
	switch node.(type) {
	case *ast.StructType:
		v.res.Struct++
	case *ast.InterfaceType:
		v.res.Interface++
	}
}

//FuncVisitor visits functions to count functions, methods tests etc.
type FuncVisitor struct {
	res *Result
	fset *token.FileSet
}

func (v *FuncVisitor) Visit(node ast.Node) {

	switch x:= node.(type) {
		case *ast.FuncDecl:

		nodeLOC := (v.fset.Position(x.End()).Line) - (v.fset.Position(x.Pos()).Line + 1)

		if x.Recv == nil {
			//count function
			v.res.Function++

			//count function lines
			v.res.FunctionLOC += nodeLOC

			//check if the function is a test
			if x.Name.IsExported() {
				v.res.ExportedFunction++
				if strings.HasPrefix(x.Name.String(), "Test") {
					if len(x.Type.Params.List) != 0 {
						xt := x.Type.Params.List[0].Type.(*ast.StarExpr)
						xtx := xt.X.(*ast.SelectorExpr)
						for _, validArgType := range []string{"testing.T", "testing.M", "testing.B"} {
							if fmt.Sprintf("%s.%s", xtx.X, xtx.Sel) == validArgType {
								v.res.Test++
							}
						}
					}
				}
			}
		} else {

			v.res.Method++

			//count function lines
			v.res.MethodLOC += nodeLOC

			if x.Name.IsExported() {
				v.res.ExportedMethod++
			}

		}
	}
}

//ImportVisitor visits import statements
type ImportVisitor struct {
	res *Result
}

func (v *ImportVisitor) Visit(node ast.Node) {
	switch node.(type) {
		case *ast.ImportSpec:
		v.res.Import++
	}
}

//FlowControlVisitor visits ifs, cases etc.
type FlowControlVisitor struct {
	res *Result
}

func (v *FlowControlVisitor) Visit(node ast.Node) {
	switch node.(type) {
		case *ast.IfStmt:
		v.res.IfStatement++
		case *ast.SwitchStmt:
		v.res.SwitchStatement++
		case *ast.GoStmt:
		v.res.GoStatement++
	}
}
