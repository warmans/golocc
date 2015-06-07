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

	if result.Function != 3 {
		t.Error("expected 3 function got", result.Function)
	}

	if result.ExportedFunction != 2 {
		t.Error("expected 2 exported function")
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
	if result.Tests != 1 {
		t.Error("expected 1 rest got", result.Tests)
	}
}

func TestAssertCount(t *testing.T) {
	parser := Parser{}
	result := parser.ParseDir("./fixture")
	if result.Assertions != 2 {
		t.Error("expected 2 rest got", result.Tests)
	}
}
