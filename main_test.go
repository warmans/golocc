package main

import (
	"testing"
)

func TestIgnoreVendor(t *testing.T) {
	parser1 := Parser{}
	result1 := parser1.ParsePackages([]string{"./fixture/..."}, true)

	parser2 := Parser{}
	result2 := parser2.ParsePackages([]string{"./fixture/..."}, false)

	if result1.CLOC == result2.CLOC {
		t.Error("got same CLOC result with and without vendor dir")
	}
}

func TestCLOC(t *testing.T) {

	parser := Parser{}
	result := parser.ParsePackages([]string{"./fixture/..."}, true)

	if result.CLOC != 8 {
		t.Error("expected 8 comment lines got", result.CLOC)
	}
}

func TestStructCount(t *testing.T) {

	parser := Parser{}
	result := parser.ParsePackages([]string{"./fixture/..."}, true)

	if result.Struct != 3 {
		t.Error("expected 3 structs got", result.Struct)
	}
}

func TestInterfaceCount(t *testing.T) {

	parser := Parser{}
	result := parser.ParsePackages([]string{"./fixture/..."}, true)

	if result.Interface != 1 {
		t.Error("expected 1 interface got", result.Interface)
	}
}

func TestMethodCount(t *testing.T) {

	parser := Parser{}
	result := parser.ParsePackages([]string{"./fixture/..."}, true)

	if result.Method != 2 {
		t.Error("expected 2 methods")
	}

	if result.ExportedMethod != 1 {
		t.Error("expected 1 exported method got", result.ExportedMethod)
	}
}

func TestMethodLineCount(t *testing.T) {

	parser := Parser{}
	result := parser.ParsePackages([]string{"./fixture/..."}, true)

	if result.MethodLOC != 15 {
		t.Error("expected 15 method lines got", result.MethodLOC)
	}
}

func TestFunctionCount(t *testing.T) {

	parser := Parser{}
	result := parser.ParsePackages([]string{"./fixture/..."}, true)

	if result.Function != 5 {
		t.Error("expected 5 functions got", result.Function)
	}

	if result.ExportedFunction != 4 {
		t.Error("expected 4 exported functions")
	}
}

func TestFunctionLineCount(t *testing.T) {

	parser := Parser{}
	result := parser.ParsePackages([]string{"./fixture/..."}, true)

	if result.FunctionLOC != 10 {
		t.Error("expected 10 function lines got", result.FunctionLOC)
	}
}

func TestImportCount(t *testing.T) {

	parser := Parser{}
	result := parser.ParsePackages([]string{"./fixture/..."}, true)

	if result.Import != 5 {
		t.Error("expected 5 imports got", result.Import)
	}
}

func TestTestCount(t *testing.T) {
	parser := Parser{}
	result := parser.ParsePackages([]string{"./fixture/..."}, true)
	if result.Test != 2 {
		t.Error("expected 2 tests got", result.Test)
	}
}

func TestBenchmarkCount(t *testing.T) {
	parser := Parser{}
	result := parser.ParsePackages([]string{"./fixture/..."}, true)
	if result.Benchmark != 1 {
		t.Error("expected 1 benchmark got", result.Benchmark)
	}
}

func TestAssertCount(t *testing.T) {
	parser := Parser{}
	result := parser.ParsePackages([]string{"./fixture/..."}, true)
	if result.Assertion != 2 {
		t.Error("expected 2 assertions got", result.Assertion)
	}
}

func TestIfStatementCount(t *testing.T) {
	parser := Parser{}
	result := parser.ParsePackages([]string{"./fixture/..."}, true)
	if result.IfStatement != 2 {
		t.Error("expected 2 if statements got", result.IfStatement)
	}
}

func TestSwitchStatementCount(t *testing.T) {
	parser := Parser{}
	result := parser.ParsePackages([]string{"./fixture/..."}, true)
	if result.SwitchStatement != 1 {
		t.Error("expected 1 switch statement got", result.SwitchStatement)
	}
}

func TestGoStatementCount(t *testing.T) {
	parser := Parser{}
	result := parser.ParsePackages([]string{"./fixture/..."}, true)
	if result.GoStatement != 2 {
		t.Error("expected 2 switch statement got", result.GoStatement)
	}
}

func TestPackageCount(t *testing.T) {
	parser := Parser{}
	result := parser.ParsePackages([]string{"./fixture/..."}, true)
	if result.GoStatement != 2 {
		t.Error("expected 2 package got", result.Package)
	}
}

