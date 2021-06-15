package parser

import "github.com/cupen/yobuffer/codegen/models"

var defaultParser = &treeParser{}

func Parse(file string) {

}

type treeParser struct {
}

func (this *treeParser) Parse(data []byte) *models.Context {
	c := models.NewContext()
	c.DefineStruct(&models.Struct{
		ID:       10001,
		Name:     "GatePlayerInfo",
		Desc:     "...",
		Encoding: "tiny",
		Fields: []*models.Field{
			{},
		},
	})

}
