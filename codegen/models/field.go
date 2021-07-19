package models

import "github.com/cupen/yobuffer/encodings/types"

type Field struct {
	Name   string
	Type   string
	Offset int // base-0
}

func (f *Field) GetSizeFixed() int {
	return types.SizeOf(f.Type)
}
