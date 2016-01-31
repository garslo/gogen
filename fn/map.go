package fn

import (
	"fmt"

	"github.com/garslo/gogen"
)

func validateMapParams(fun gogen.Function) {
	if len(fun.Parameters) != 1 {
		panic(fmt.Sprintf("Map() fun needs to have 1 param, has %d", len(fun.Parameters)))
	}

	if len(fun.ReturnTypes) != 1 {
		panic(fmt.Sprintf("Map() fun needs to have 1 return type, has %d", len(fun.ReturnTypes)))
	}
}

func (me Functions) Map(fun gogen.Function) gogen.Function {
	validateMapParams(fun)
	paramType := fun.Parameters[0].TypeName
	returnType := fun.ReturnTypes[0].TypeName
	fn := gogen.Function{
		Name: "Map_" + paramType + "_" + returnType,
		ReturnTypes: gogen.Types{
			gogen.Type{
				Name:     "ret",
				TypeName: "[]" + returnType,
			},
		},
		Parameters: gogen.Types{
			gogen.Type{
				Name:     "fn",
				TypeName: fmt.Sprintf("func(%s) %s", paramType, returnType),
			},
			gogen.Type{
				Name:     "xs",
				TypeName: "[]" + paramType,
			},
		},
		Body: []gogen.Statement{
			gogen.Range{
				Value:      gogen.Var{"x"},
				RangeValue: gogen.Var{"xs"},
				Body: []gogen.Statement{
					gogen.Assign{
						Lhs: gogen.Var{"ret"},
						Rhs: gogen.CallFunction{
							Func: gogen.Var{"append"},
							Params: []gogen.Expression{
								gogen.Var{"ret"},
								gogen.CallFunction{
									Func:   gogen.Var{"fn"},
									Params: []gogen.Expression{gogen.Var{"x"}},
								},
							},
						},
					},
				},
			},
		},
	}
	me.Package.Declare(fn)
	return fn
}
