package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		Usage()

	}
	for _, file := range os.Args[1:] {
		DieIfError(PrintAst(file))

	}

}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s file [file...]\n", os.Args[0])
	os.Exit(0)

}

func DieIfError(err error) {
	if err != nil {
		if errList, ok := err.(scanner.ErrorList); ok {
			HandleScannerError(errList)

		}
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)

	}

}

func HandleScannerError(errList scanner.ErrorList) {
	fmt.Fprintf(os.Stderr, "%d errors\n", errList.Len())
	maxErrs := 10
	errCount := 0
	for _, scannerErr := range errList {
		if errCount == maxErrs {
			fmt.Fprint(os.Stderr, "too many errors")
			break

		}
		errCount++
		fmt.Fprintf(os.Stderr, "Error: %v\n", scannerErr)

	}
	os.Exit(1)

}

func PrintAst(file string) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return err

	}
	ast.Print(fset, f)
	return nil
}
