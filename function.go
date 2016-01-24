package gogen

type Function struct {
	Name        string
	ReturnTypes Types
	Parameters  Types
	Body        []Statement
}

type Functions []Function

func (me *Functions) Add(fn Function) {
	*me = append(*me, fn)
}
