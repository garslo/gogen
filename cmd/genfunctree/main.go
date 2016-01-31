package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"

	"github.com/garslo/gogen"
)

type Function struct {
	Name   string
	Params *ast.FieldList
	Return *ast.FieldList
}

type FunctionInfo map[string][]Function

func (me FunctionInfo) Add(pkg string, fn Function) {
	// <|> TODO May not need this
	if _, ok := me[pkg]; !ok {
		me[pkg] = []Function{fn}
	} else {
		me[pkg] = append(me[pkg], fn)
	}
}

type Walker struct {
	Functions FunctionInfo
	Fset      *token.FileSet
	ThisFile  *ast.File
}

func (me *Walker) Walk(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	if filepath.Ext(info.Name()) == ".go" {
		me.ThisFile, err = parser.ParseFile(me.Fset, path, nil, parser.ParseComments)
		if err != nil {
			log.Printf(err.Error())
		}
		ast.Walk(me, me.ThisFile)
	}
	return nil
}

func (me *Walker) Visit(node ast.Node) ast.Visitor {
	switch fn := node.(type) {
	case *ast.FuncDecl:
		if !fn.Name.IsExported() {
			return nil
		}
		if fn.Recv != nil {
			return nil
		}

		toAdd := Function{
			Name:   fn.Name.Name,
			Params: fn.Type.Params,
			Return: fn.Type.Results,
		}
		me.Functions.Add(me.ThisFile.Name.Name, toAdd)
		return nil
	default:
		return me
	}
}

func main() {
	var (
		root string
	)
	flag.StringVar(&root, "r", ".", "package root")
	flag.Parse()

	pwd, _ := os.Getwd()
	os.Chdir(root)
	defer os.Chdir(pwd)
	walker := Walker{
		Functions: make(FunctionInfo),
		Fset:      token.NewFileSet(),
	}
	filepath.Walk(root, walker.Walk)

	pkg := &gogen.Package{Name: "mymath"}
	pkg.Declare(gogen.Import{"math"})
	pkg.Declare(gogen.Import{"github.com/garslo/gogen"})

	for _, f := range walker.Functions["math"] {
		if f.Name == "Abs" {
			ret := gogen.Var{"ret"}
			pkg.Declare(
				gogen.Function{
					Name: "Math" + f.Name,
					Parameters: gogen.Types{
						gogen.Type{
							Name:     "expr",
							TypeName: "gogen.Expression",
						},
					},
					ReturnTypes: gogen.Types{
						gogen.Type{
							TypeName: "gogen.Function",
						},
					},
					Body: []gogen.Statement{
						gogen.DeclareAndAssign{
							Lhs: ret,
							Rhs: gogen.Thunk{
								Expr: &ast.CompositeLit{
									Type: &ast.SelectorExpr{
										X: &ast.Ident{
											Name: "gogen",
										},
										Sel: &ast.Ident{
											Name: "Function",
										},
									},
									Elts: []ast.Expr{
										&ast.KeyValueExpr{
											Key: &ast.Ident{
												Name: "Name",
											},
											Value: &ast.BasicLit{
												Kind:  token.STRING,
												Value: "Math" + f.Name,
											},
										},
									},
								},
							},
						},
						gogen.Return{
							[]gogen.Expression{ret},
						},
						//gogen.Return{
						//	Values: []Expression{
						//		Thunk{
						//			Expr: &ast.CompositeLit{
						//				Type: &ast.SelectorExpr{
						//					X: &ast.Ident{
						//						Name: "gogen",
						//					},
						//					Sel: &ast.Ident{
						//						Name: "Function",
						//					},
						//				},
						//			},
						//		},
						//	},
						//},
					},
				},
			)
			break
		}
	}
	pkg.WriteTo(os.Stdout)
}
