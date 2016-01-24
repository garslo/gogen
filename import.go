package gogen

type Import struct {
	Name string
}

type Imports []Import

func (me *Imports) Add(imp Import) {
	*me = append(*me, imp)
}
