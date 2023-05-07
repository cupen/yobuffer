package models

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"text/template"

	"github.com/cupen/yobuffer/codegen/languages"
	"github.com/cupen/yobuffer/encodings/types"
)

type Context struct {
	Package          string
	Structs          map[int]*Struct
	Structss         map[string]*Struct
	RPCs             map[int]*RPC
	RPCss            map[string]*RPC
	ExternalStructs  map[int]*ExternalStruct
	ExternalStructss map[string]*ExternalStruct
}

func NewContext() *Context {
	obj := Context{
		Structs:          map[int]*Struct{},
		Structss:         map[string]*Struct{},
		RPCs:             map[int]*RPC{},
		RPCss:            map[string]*RPC{},
		ExternalStructs:  map[int]*ExternalStruct{},
		ExternalStructss: map[string]*ExternalStruct{},
	}
	obj.DefineStruct(&Struct{
		ID:   0,
		Name: "void",
	})
	return &obj
}

// DefineStruct ...
func (this *Context) DefineStruct(s *Struct) error {
	if s == nil {
		return fmt.Errorf("nil struct")
	}
	if old, ok := this.Structs[s.ID]; ok {
		return fmt.Errorf("duplicated Struct(id:%d name:%s) with Struct(id:%d name:%s) ", s.ID, s.Name, old.ID, old.Name)
	}
	if old, ok := this.Structss[s.Name]; ok {
		return fmt.Errorf("duplicated Struct(id:%d name:%s) with Struct(id:%d name:%s) ", s.ID, s.Name, old.ID, old.Name)
	}
	this.Structs[s.ID] = s
	this.Structss[s.Name] = s
	return nil
}

// DefineRPC ...
func (this *Context) DefineRPC(r *RPC) error {
	if r == nil {
		return fmt.Errorf("nil struct")
	}
	if r.ID <= 0 || r.Name == "" {
		return fmt.Errorf("invalid RPC(id:%d name:%s)", r.ID, r.Name)
	}

	old, ok := this.RPCs[r.ID]
	if ok {
		return fmt.Errorf("duplicated RPC(id:%d name:%s) with RPC(id:%d name:%s) ", r.ID, r.Name, old.ID, old.Name)
	}
	old, ok = this.RPCss[r.Name]
	if ok {
		return fmt.Errorf("duplicated RPC(id:%d name:%s) with RPC(id:%d name:%s) ", r.ID, r.Name, old.ID, old.Name)
	}
	this.RPCs[r.ID] = r
	this.RPCss[r.Name] = r
	return nil
}

func (this *Context) DefineExternalStruct(r *ExternalStruct) error {
	if !r.IsValid() {
		return fmt.Errorf("invalid external struct %+v", r)
	}
	if old, ok := this.ExternalStructs[r.ID]; ok {
		return fmt.Errorf("duplicated ExternalStruct(id:%d name:%s) with RPC(id:%d name:%s) ", r.ID, r.Name, old.ID, old.Name)
	}
	if old, ok := this.ExternalStructss[r.Name]; ok {
		return fmt.Errorf("duplicated ExternalStruct(id:%d name:%s) with RPC(id:%d name:%s) ", r.ID, r.Name, old.ID, old.Name)
	}
	this.ExternalStructs[r.ID] = r
	this.ExternalStructss[r.Name] = r
	return nil
}

func (this *Context) GetRPCList() []*RPC {
	list := []*RPC{}
	for _, r := range this.RPCs {
		list = append(list, r)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})
	return list
}

func (this *Context) GetExternalStructs() []*ExternalStruct {
	list := make([]*ExternalStruct, len(this.ExternalStructs))
	i := 0
	for _, obj := range this.ExternalStructs {
		list[i] = obj
		i++
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})
	return list
}

func (this *Context) WriteTo(w io.Writer) error {
	t := template.New("go-coder").Funcs(template.FuncMap{
		"fieldSizeOf": types.SizeOf,
		"title":       strings.Title,
		"upper":       strings.ToUpper,
		"lower":       strings.ToLower,
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
