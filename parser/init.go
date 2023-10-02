package parser

//go:generate antlr4 -Dlanguage=Go -visitor -listener -package parser Yobuffer.g4

var (
	defaultParser = New()
)

// func Parse(file string) (*models.Context, error) {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		return nil, err
// 	}
// 	data, err := io.ReadAll(f)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return defaultParser.Parse(data)
// }
