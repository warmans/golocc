package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestTextOutput(t *testing.T) {

	buff := &bytes.Buffer{}
	report := &TextReport{writer: buff}
	res := &Result{
		LOC:              10,
		CLOC:             11,
		NCLOC:            12,
		Struct:           13,
		Interface:        14,
		Method:           15,
		ExportedMethod:   16,
		MethodLOC:        17,
		Function:         18,
		ExportedFunction: 19,
		FunctionLOC:      20,
		Import:           21,
		IfStatement:      22,
		SwitchStatement:  23,
		GoStatement:      24,
		Test:             25,
		Assertion:        26}

	report.Print(res)

	if !strings.Contains(buff.String(), "Lines of Code: 10 (11 CLOC, 12 NCLOC)") {
		t.Error("Unexpected output in text report (LOC)")
	}

	if !strings.Contains(buff.String(), "Imports:       21") {
		t.Error("Unexpected output in text report (Imports)")
	}

	if !strings.Contains(buff.String(), "Structs:       13") {
		t.Error("Unexpected output in text report (Structs)")
	}

	if !strings.Contains(buff.String(), "Interfaces:    14") {
		t.Error("Unexpected output in text report (Interfaces)")
	}

	if !strings.Contains(buff.String(), "Methods:       15 (16 Exported, 17 LOC, 1 LOC Avg.)") {
		t.Error("Unexpected output in text report (Methods)")
	}

	if !strings.Contains(buff.String(), "Functions:     18 (19 Exported, 20 LOC, 1 LOC Avg.)") {
		t.Error("Unexpected output in text report (Functions)")
	}

	if !strings.Contains(buff.String(), "Ifs:           22") {
		t.Error("Unexpected output in text report (Ifs)")
	}

	if !strings.Contains(buff.String(), "Switches:      23") {
		t.Error("Unexpected output in text report (Switches)")
	}

	if !strings.Contains(buff.String(), "Go Routines:   24") {
		t.Error("Unexpected output in text report (Go Routines)")
	}

	if !strings.Contains(buff.String(), "Tests:         25") {
		t.Error("Unexpected output in text report (Tests)")
	}

	if !strings.Contains(buff.String(), "Assertions:    26") {
		t.Error("Unexpected output in text report (Assertions)")
	}
}
