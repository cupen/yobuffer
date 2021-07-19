package models

import (
	"fmt"
	"io"
	"text/template"

	"github.com/cupen/yobuffer/codegen/languages"
	"github.com/cupen/yobuffer/encodings/types"
)

type Context struct {
	Package string
	Structs map[int]*Struct
	RPCs    map[int]*RPC
}

func NewContext() *Context {
	return &Context{
		Structs: map[int]*Struct{},
		RPCs:    map[int]*RPC{},
	}
}

// DefineStruct ...
func (this *Context) DefineStruct(s *Struct) error {
	if s == nil {
		return fmt.Errorf("nil struct")
	}

	old, ok := this.Structs[s.ID]
	if ok {
		return fmt.Errorf("duplicated Struct(id:%d name:%s) with Struct(id:%d name:%s) ", s.ID, s.Name, old.ID, old.Name)
	}
	this.Structs[s.ID] = s
	return nil
}

// DefineRPC ...
func (this *Context) DefineRPC(r *RPC) error {
	if r == nil {
		return fmt.Errorf("nil struct")
	}

	old, ok := this.RPCs[r.ID]
	if ok {
		return fmt.Errorf("duplicated RPC(id:%d name:%s) with RPC(id:%d name:%s) ", r.ID, r.Name, old.ID, old.Name)
	}
	this.RPCs[r.ID] = r
	return nil
}

func (this *Context) WriteTo(w io.Writer) error {
	t := template.New("go-coder").Funcs(template.FuncMap{
		"fieldSizeOf": types.SizeOf,
		"panic": func(err string) int {
			panic(fmt.Errorf(err))
		},
	})

	t, err := t.Parse(string(languages.Go))
	if err != nil {
		return err
	}
	return t.Execute(w, this)
}
