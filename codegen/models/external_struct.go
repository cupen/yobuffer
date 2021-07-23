package models

type ExternalStruct struct {
	ID   int
	Name string
}

func (es *ExternalStruct) IsValid() bool {
	return es.ID > 0 && es.Name != ""
}
