#GOLOCC

[![Build Status](https://travis-ci.org/warmans/golocc.svg)](https://travis-ci.org/warmans/golocc) [![code-coverage](http://gocover.io/_badge/github.com/warmans/golocc)](http://gocover.io/github.com/warmans/golocc)

Utility for counting lines of code (LOC, CLOC, NCLOC) and structures/statements in a go package.

### Usage

Within a package: 

`golocc ./...`

or 

`golocc --no-vendor ./...`

or from outside the package:

`goloc $GOPATH/src/[some package]`


### Options

| flag | function                                           |
|------|----------------------------------------------------|
| d    | Target directory. Defaults to current working dir (deprecated)  |
| o    | Output format. Supports `text`, `json`             |


### Sample `text` output:

```
--------------------------------------------------------------------------------
Lines of Code: 547 (36 CLOC, 511 NCLOC)
Packages:      4
Imports:       28
Structs:       12
Interfaces:    3
Methods:       11 (9 Exported, 221 LOC, 20 LOC Avg.)
Functions:     24 (21 Exported, 215 LOC, 8 LOC Avg.)
--------------------------------------------------------------------------------
Ifs:           57 
Switches:      2 
Go Routines:   2 
--------------------------------------------------------------------------------
Tests:         19 
Benchmarks:    1 
Assertions:    2 
--------------------------------------------------------------------------------
```
