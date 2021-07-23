package parser

type BeginPackage struct {
	Name string
}

type StructBegin struct {
	Name string
}

type StructEnd struct{}

type StructField struct {
	Name   string
	Type   string
	Offset int
}

type Meta struct {
	Fields map[string]string
}

type RPC struct {
	Name   string
	Fields []*StructField
	Return string
}
type ExternalStruct struct {
	Name string
}
