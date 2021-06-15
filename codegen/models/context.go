package models

import "fmt"

type Context struct {
	structs map[int]*Struct
	rpcs    map[int]*RPC
}

func NewContext() *Context {
	return &Context{
		structs: map[int]*Struct{},
		rpcs:    map[int]*RPC{},
	}
}

// DefineStruct ...
func (this *Context) DefineStruct(s *Struct) error {
	if s == nil {
		return fmt.Errorf("nil struct")
	}

	old, ok := this.structs[s.ID]
	if ok {
		return fmt.Errorf("duplicated Struct(id:%d name:%s) with Struct(id:%d name:%s) ", s.ID, s.Name, old.ID, old.Name)
	}
	this.structs[s.ID] = s
	return nil
}

// DefineRPC ...
func (this *Context) DefineRPC(r *RPC) error {
	if r == nil {
		return fmt.Errorf("nil struct")
	}

	old, ok := this.rpcs[r.ID]
	if ok {
		return fmt.Errorf("duplicated RPC(id:%d name:%s) with RPC(id:%d name:%s) ", r.ID, r.Name, old.ID, old.Name)
	}
	this.rpcs[r.ID] = r
	return nil
}
