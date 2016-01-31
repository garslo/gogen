package fn

import "github.com/garslo/gogen"

func (me Functions) Print(xs ...gogen.Expression) gogen.Function {
	fn := gogen.Function{
		Name: "Print",
		Parameters: gogen.Types{
			gogen.Type{
				Name:     "x",
				TypeName: "interface{}",
			},
		},
		Body: []gogen.Statement{
			gogen.CallFunction{
				Func:   gogen.Dotted{gogen.Pkg("fmt"), "Println"},
				Params: []gogen.Expression{gogen.Var{"x"}},
			},
		},
	}
	me.Package.Declare(fn)
	return fn
}

func FmtPrintln(xs ...gogen.Expression) gogen.CallFunction {
	return gogen.CallFunction{
		Func:   gogen.Dotted{gogen.Pkg("fmt"), "Println"},
		Params: xs,
	}
}
