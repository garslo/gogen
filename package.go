package gogen

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"io"
)

type Package struct {
	Name      string
	Imports   Imports
	Functions Functions
}

func (me *Package) Import(pkg string) *Package {
	me.Imports.Add(Import{pkg})
	return me
}

func (me *Package) Function(fn Function) *Package {
	me.Functions.Add(fn)
	return me
}

func (me *Package) Ast() ast.Node {
	importDecls := make([]ast.Decl, len(me.Imports))
	for i, imp := range me.Imports {
		importDecls[i] = &ast.GenDecl{
			Tok: token.IMPORT,
			Specs: []ast.Spec{
				&ast.ImportSpec{
					Path: &ast.BasicLit{
						Kind:  token.STRING,
						Value: fmt.Sprintf(`"%s"`, imp.Name),
					},
				},
			},
		}
	}
	funcDecls := make([]ast.Decl, len(me.Functions))
	for i, fn := range me.Functions {
		paramFields := make([]*ast.Field, len(fn.Parameters))
		for j, param := range fn.Parameters {
			var names []*ast.Ident
			if param.Name != "" {
				names = []*ast.Ident{
					&ast.Ident{
						Name: param.Name,
						Obj: &ast.Object{
							Kind: ast.Var,
							Name: param.Name,
						},
					},
				}
			}
			paramFields[j] = &ast.Field{
				Names: names,
				Type: &ast.Ident{
					Name: param.TypeName,
				},
			}
		}
		returnFields := make([]*ast.Field, len(fn.ReturnTypes))
		for j, ret := range fn.ReturnTypes {
			var names []*ast.Ident
			if ret.Name != "" {
				names = []*ast.Ident{
					&ast.Ident{
						Name: ret.Name,
						Obj: &ast.Object{
							Kind: ast.Var,
							Name: ret.Name,
						},
					},
				}
			}
			returnFields[j] = &ast.Field{
				Names: names,
				Type: &ast.Ident{
					Name: ret.TypeName,
				},
			}
		}
		stmts := make([]ast.Stmt, len(fn.Body))
		for j, stmt := range fn.Body {
			stmts[j] = stmt.Ast()
		}

		funcDecls[i] = &ast.FuncDecl{
			Name: &ast.Ident{
				Name: fn.Name,
				Obj: &ast.Object{
					Kind: ast.Fun,
					Name: fn.Name,
				},
			},
			Type: &ast.FuncType{
				Params: &ast.FieldList{
					List: paramFields,
				},
				Results: &ast.FieldList{
					List: returnFields,
				},
			},
			Body: &ast.BlockStmt{
				List: stmts,
			},
		}
	}
	return &ast.File{
		Name: &ast.Ident{
			Name: me.Name,
		},
		Decls: append(importDecls, funcDecls...),
	}
}

func (me *Package) WriteTo(w io.Writer) error {
	fset := token.NewFileSet()
	return format.Node(w, fset, me.Ast())
}
