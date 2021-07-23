package models

type RPC struct {
	ID   int
	Name string
	Desc string
	// v1
	Args []*Field

	// v2
	Arg    *Field
	Return string
}

func (r *RPC) HasArg() bool {
	return r.Arg != nil
}

func (r *RPC) HasReturn() bool {
	return r.Return != ""
}
