package main

import (
	"testing"
)

func TestCLOC(t *testing.T) {

	parser := Parser{}
	result := parser.ParseDir("./fixture")

	if result.CLOC != 8 {
		t.Error("expected 8 comment lines")
	}
}

func TestStructCount(t *testing.T) {

	parser := Parser{}
	result := parser.ParseDir("./fixture")

	if result.Struct != 2 {
		t.Error("expected 2 structs")
	}
}

func TestInterfaceCount(t *testing.T) {

	parser := Parser{}
	result := parser.ParseDir("./fixture")

	if result.Interface != 1 {
		t.Error("expected 1 interface")
	}
}

func TestMethodCount(t *testing.T) {

	parser := Parser{}
	result := parser.ParseDir("./fixture")

	if result.Method != 2 {
		t.Error("expected 2 methods")
	}

	if result.ExportedMethod != 1 {
		t.Error("expected 1 exported method")
	}
}

func TestFunctionCount(t *testing.T) {

	parser := Parser{}
	result := parser.ParseDir("./fixture")

	if result.Function != 5 {
		t.Error("expected 5 functions got", result.Function)
	}

	if result.ExportedFunction != 4 {
		t.Error("expected 4 exported functions")
	}
}

func TestImportCount(t *testing.T) {

	parser := Parser{}
	result := parser.ParseDir("./fixture")

	if result.Import != 5 {
		t.Error("expected 5 imports got", result.Import)
	}
}

func TestTestCount(t *testing.T) {
	parser := Parser{}
	result := parser.ParseDir("./fixture")
	if result.Test != 3 {
		t.Error("expected 3 tests got", result.Test)
	}
}

func TestAssertCount(t *testing.T) {
	parser := Parser{}
	result := parser.ParseDir("./fixture")
	if result.Assertion != 2 {
		t.Error("expected 2 tests got", result.Test)
	}
}

func TestIfStatementCount(t *testing.T) {
	parser := Parser{}
	result := parser.ParseDir("./fixture")
	if result.IfStatement != 2 {
		t.Error("expected 2 if statements got", result.IfStatement)
	}
}

func TestSwitchStatementCount(t *testing.T) {
	parser := Parser{}
	result := parser.ParseDir("./fixture")
	if result.SwitchStatement != 1 {
		t.Error("expected 1 switch statement got", result.SwitchStatement)
	}
}

func TestGoStatementCount(t *testing.T) {
	parser := Parser{}
	result := parser.ParseDir("./fixture")
	if result.GoStatement != 2 {
		t.Error("expected 2 switch statement got", result.GoStatement)
	}
}
