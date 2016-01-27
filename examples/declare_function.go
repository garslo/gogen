package main

import (
	"os"
	. "github.com/garslo/gogen"
)

func main() {
	fooFuncParamName := "x"
	fooFunc := Function{
		Name: "foo",
		Parameters: Types{
			Type{
				Name:     fooFuncParamName,
				TypeName: StringT,
			},
		},
		Body: []Statement{
			CallFunction{
				Func: Dotted{Pkg(Fmt), "Println"},
				Params: []Expression{
					Var{fooFuncParamName},
				},
			},
		},
	}
	toPrint := Var{"y"}
	pkg := Package{Name: "main"}
	pkg.Declare(Import{Fmt})
	pkg.Declare(fooFunc)
	pkg.Declare(Function{
		Name: "main",
		Body: []Statement{
			DeclareAndAssign{
				toPrint,
				String("i'm being printed from a code-genned function"),
			},
			CallFunction{
				Func:   Var{fooFunc.Name},
				Params: []Expression{toPrint},
			},
		},
	})
	pkg.WriteTo(os.Stdout)
}
