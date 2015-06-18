#GOLOCC

Utility for counting lines of code (LOC, CLOC, NCLOC) and structures/statements in a go package.

### Usage
`golocc -d $GOPATH/src/whatever`

or

`cd $GOPATH/src/whatever; golocc`


### Options

| flag | function                                           |
|------|----------------------------------------------------|
| d    | Target directory. Defaults to current working dir  |
| o    | Output format. Supports `text`, `json`             |


### Sample `text` output:

```
--------------------------------------------------------------------------------
Lines of Code: 66 (8 CLOC, 58 NCLOC)
Imports:       5
Structs:       2
Interfaces:    1
Methods:       2 (1 Exported, 15 LOC, 7 LOC Avg.)
Functions:     5 (4 Exported, 10 LOC, 2 LOC Avg.)
--------------------------------------------------------------------------------
Ifs:           2
Switches:      2
Go Routines:   2
--------------------------------------------------------------------------------
Tests:         3
Assertions:    2
--------------------------------------------------------------------------------

```
