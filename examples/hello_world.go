package main

import (
	"os"
	. "github.com/garslo/gogen"
)

func main() {
	pkg := Package{Name: "main"}
	pkg.Declare(Import{"fmt"})
	pkg.Declare(Function{
		Name: "main",
		Body: []Statement{
			CallFunction{
				Func:   Dotted{Var{"fmt"}, "Println"},
				Params: []Expression{Var{`"Hello World!"`}},
			},
		},
	})
	pkg.WriteTo(os.Stdout)
}
