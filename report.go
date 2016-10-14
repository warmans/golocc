package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

//ReportInterface - reports that parse results and print out a report
type ReportInterface interface {
	Print(*Result)
}

//JSONReport json structure for LOC report
type JSONReport struct {
	writer io.Writer
}

//Print - print out parsed report in json format
func (j *JSONReport) Print(res *Result) {
	jsonOutput, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(j.writer, string(jsonOutput))
}

//TextReport - plaintext report output
type TextReport struct {
	writer io.Writer
}

//Print - print out plaintext report
func (t *TextReport) Print(res *Result) {
	fmt.Fprintf(t.writer, "\n")
	fmt.Fprintln(t.writer, strings.Repeat("-", 80))
	fmt.Fprintf(t.writer, "Lines of Code: %v (%v CLOC, %v NCLOC)\n", res.LOC, res.CLOC, res.NCLOC)
	fmt.Fprintf(t.writer, "Packages:      %v\n", res.Package)
	fmt.Fprintf(t.writer, "Imports:       %v\n", res.Import)
	fmt.Fprintf(t.writer, "Structs:       %v\n", res.Struct)
	fmt.Fprintf(t.writer, "Interfaces:    %v\n", res.Interface)
	fmt.Fprintf(t.writer, "Methods:       %v (%v Exported, %v LOC, %v LOC Avg.)\n", res.Method, res.ExportedMethod, res.MethodLOC, t.divideBy(res.MethodLOC, res.Method))
	fmt.Fprintf(t.writer, "Functions:     %v (%v Exported, %v LOC, %v LOC Avg.)\n", res.Function, res.ExportedFunction, res.FunctionLOC, t.divideBy(res.FunctionLOC, res.Function))
	fmt.Fprintln(t.writer, strings.Repeat("-", 80))
	fmt.Fprintf(t.writer, "Ifs:           %v \n", res.IfStatement)
	fmt.Fprintf(t.writer, "Switches:      %v \n", res.SwitchStatement)
	fmt.Fprintf(t.writer, "Go Routines:   %v \n", res.GoStatement)
	fmt.Fprintln(t.writer, strings.Repeat("-", 80))
	fmt.Fprintf(t.writer, "Tests:         %v \n", res.Test)
	fmt.Fprintf(t.writer, "Benchmarks:    %v \n", res.Benchmark)
	fmt.Fprintf(t.writer, "Assertions:    %v \n", res.Assertion)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Fprintf(t.writer, "\n")
}

//Safely divide where vivisor might be zero
func (t *TextReport) divideBy(num, dividedBy int) int {
	if dividedBy == 0 {
		return 0
	}
	return num / dividedBy
}
