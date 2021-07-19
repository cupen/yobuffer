package models

import "fmt"

type Struct struct {
	ID       int
	Name     string
	Desc     string
	Encoding string
	Fields   []*Field
	Source   string
}

func (this *Struct) String() string {
	return fmt.Sprintf("struct(%d)%s", this.ID, this.Name)
}
