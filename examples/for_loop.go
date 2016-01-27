package main

import (
	"os"
	. "github.com/garslo/gogen"
)

func main() {
	pkg := Package{
		Name: "main",
	}
	i := Name("i")
	pkg.Declare(Import{"os"}).
		Declare(Import{"fmt"}).
		Declare(Function{
		Name: "main",
		Body: []Statement{
			Declare{i.Name, "int"},
			For{
				Init:      Assign{i, Int(0)},
				Condition: LessThanOrEqual{i, Int(10)},
				Post:      Increment{i},
				Body: []Statement{
					CallFunction{
						Dotted{Pkg("fmt"), "Println"},
						[]Expression{i},
					},
				},
			},
			CallFunction{
				Dotted{Pkg("os"), "Exit"},
				[]Expression{i},
			},
		},
	})
	pkg.WriteTo(os.Stdout)
}
