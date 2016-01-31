package main

import (
	"os"
	. "github.com/garslo/gogen"
)

func main() {
	FmtPrintln := Functor{Dotted{Var{"fmt"}, "Println"}}
	fooParamName := "x"
	Foo := Function{
		Name: "Foo",
		Parameters: Types{
			Type{
				Name:     fooParamName,
				TypeName: StringT,
			},
		},
		Body: []Statement{
			FmtPrintln.Call(Var{fooParamName}),
		},
	}
	toPrint := Var{"y"}
	pkg := Package{Name: "main"}
	pkg.Declare(Import{Fmt})
	pkg.Declare(Foo)
	pkg.Declare(Function{
		Name: "main",
		Body: []Statement{
			DeclareAndAssign{
				toPrint,
				String("i'm being printed from a code-genned function"),
			},
			Foo.Call(toPrint),
		},
	})
	pkg.WriteTo(os.Stdout)
}
