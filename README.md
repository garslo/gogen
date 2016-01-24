# gogen

A simplification of Go's `go/ast` package that allows for some
interesting code generation. Currently very rough.

# Example

Just for fun, and the resulting code won't compile (type issues), but
shows the state of things:

```go
package main

import (
	"go/format"
	"go/token"
	"os"
	. "github.com/garslo/gogen"
)

func main() {
	pkg := Package{
		Name: "foo",
		Imports: Imports{
			Import{"os"},
		},
		Functions: Functions{
			Function{
				Name: "main",
				Body: []Statement{
					Range{
						Key:        Var{"i"},
						Value:      Var{"x"},
						RangeValue: Var{"xs"},
						Body: []Statement{
							CallFunction{
								Dotted{Var{"os"}, "Exit"},
								[]Expression{Var{"1"}},
							},
						},
					},
					For{
						Init:      DeclareAndAssign{Var{"i"}, Var{"0"}},
						Condition: Var{"i"},
						Post:      Assign{Var{"i"}, Var{"0"}},
						Body: []Statement{
							DeclareAndAssign{Var{"j"}, Var{"i"}},
						},
					},
				},
			},
		},
	}

	fset := token.NewFileSet()
	format.Node(os.Stdout, fset, pkg.Ast())
}
```

Running that gives output of

```go
package foo

import "os"

func main() {
	for i, x := range xs {
		os.Exit(1)
	}
	for i := 0; i; i = 0 {
		j := i
	}
}
```
