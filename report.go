package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

//ReportInterface - reports that parse results and print out a report
type ReportInterface interface {
	Print(*Result)
}

//JSONReport json structure for LOC report
type JSONReport struct {
}

//Print - print out parsed report in json format
func (j *JSONReport) Print(res *Result) {
	jsonOutput, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(jsonOutput))
}

//TextReport - plaintext report output
type TextReport struct {
}

//Print - print out plaintext report
func (t *TextReport) Print(res *Result) {

	fmt.Printf("\n")
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("Lines of Code: %v (%v CLOC, %v NCLOC)\n", res.LOC, res.CLOC, res.NCLOC)
	fmt.Printf("Imports:       %v\n", res.Import)
	fmt.Printf("Structs:       %v\n", res.Struct)
	fmt.Printf("Interfaces:    %v\n", res.Interface)
	fmt.Printf("Methods:       %v (%v Exported, %v LOC, %v LOC Avg.)\n", res.Method, res.ExportedMethod, res.MethodLOC, t.avg(res.MethodLOC, res.Method))
	fmt.Printf("Functions:     %v (%v Exported, %v LOC, %v LOC Avg.)\n", res.Function, res.ExportedFunction, res.FunctionLOC, t.avg(res.FunctionLOC, res.Function))
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("Ifs:           %v \n", res.IfStatement)
	fmt.Printf("Switches:      %v \n", res.IfStatement)
	fmt.Printf("Go Routines:   %v \n", res.GoStatement)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("Tests:         %v \n", res.Test)
	fmt.Printf("Assertions:    %v \n", res.Assertion)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("\n")
}

func (t *TextReport) avg(num, dividedBy int) int {
	if dividedBy == 0 {
		return 0
	}
	return num / dividedBy
}
