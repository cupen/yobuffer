package ast

import (
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/cupen/yobuffer/codegen/languages"
	"github.com/cupen/yobuffer/encodings/types"
)

type Context struct {
	Package  *Package
	Imports  []*Import
	Structs  []*Struct
	Services []*Service
}

type Package struct {
	Name string
}

type Import struct {
}

func (ctx *Context) WriteTo(w io.Writer) error {
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
	return t.Execute(w, ctx)
}
