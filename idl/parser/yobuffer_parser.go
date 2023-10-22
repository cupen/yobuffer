// Code generated from Yobuffer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Yobuffer
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type YobufferParser struct {
	*antlr.BaseParser
}

var YobufferParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func yobufferParserInit() {
	staticData := &YobufferParserStaticData
	staticData.LiteralNames = []string{
		"", "'->'", "'meta'", "'import'", "'package'", "'int8'", "'int16'",
		"'int32'", "'int64'", "'uint8'", "'uint16'", "'uint32'", "'uint64'",
		"'bool'", "'string'", "'float32'", "'float64'", "'bytes'", "'message'",
		"'service'", "'rpc'", "'stream'", "'struct'", "'void'", "'map'", "'list'",
		"'@'", "';'", "'='", "'('", "')'", "'['", "']'", "'{'", "'}'", "'<'",
		"'>'", "'.'", "','", "':'", "'+'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "", "META", "IMPORT", "PACKAGE", "INT8", "INT16", "INT32", "INT64",
		"UINT8", "UINT16", "UINT32", "UINT64", "BOOL", "STRING", "FLOAT32",
		"FLOAT64", "BYTES", "MESSAGE", "SERVICE", "RPC", "STREAM", "STRUCT",
		"VOID", "MAP", "LIST", "AT", "END", "EQ", "LP", "RP", "LB", "RB", "LC",
		"RC", "LT", "GT", "DOT", "COMMA", "COLON", "PLUS", "MINUS", "NEWLINE",
		"STRING_VALUE", "BOOL_VALUE", "FLOAT_VALUE", "INT_VALUE", "IDENTIFIER",
		"WHITESPACE", "COMMENT",
	}
	staticData.RuleNames = []string{
		"schema", "packageStmt", "importStmt", "meta", "metaElement", "field",
		"fieldTag", "fieldType", "message", "messageElement", "messageType",
		"struct", "structElement", "service", "serviceElement", "rpc", "rpcField",
		"constant", "emptyStmt", "name", "ident", "fullIdent", "keywords",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 49, 211, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 1, 0, 3, 0, 48, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		5, 0, 54, 8, 0, 10, 0, 12, 0, 57, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 71, 8, 3, 10, 3, 12, 3, 74,
		9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 103, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8,
		109, 8, 8, 10, 8, 12, 8, 112, 9, 8, 1, 8, 1, 8, 1, 9, 1, 9, 3, 9, 118,
		8, 9, 1, 10, 3, 10, 121, 8, 10, 1, 10, 1, 10, 1, 10, 5, 10, 126, 8, 10,
		10, 10, 12, 10, 129, 9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 5,
		11, 137, 8, 11, 10, 11, 12, 11, 140, 9, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		3, 12, 146, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 152, 8, 13, 10, 13,
		12, 13, 155, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 3, 14, 161, 8, 14, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 171, 8, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 3, 17, 182,
		8, 17, 1, 17, 1, 17, 3, 17, 186, 8, 17, 1, 17, 1, 17, 1, 17, 3, 17, 191,
		8, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 3, 20, 199, 8, 20, 1,
		21, 1, 21, 1, 21, 5, 21, 204, 8, 21, 10, 21, 12, 21, 207, 9, 21, 1, 22,
		1, 22, 1, 22, 0, 0, 23, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 0, 2, 1, 0, 40, 41, 2, 0, 3, 20,
		23, 23, 221, 0, 47, 1, 0, 0, 0, 2, 58, 1, 0, 0, 0, 4, 62, 1, 0, 0, 0, 6,
		66, 1, 0, 0, 0, 8, 77, 1, 0, 0, 0, 10, 82, 1, 0, 0, 0, 12, 88, 1, 0, 0,
		0, 14, 102, 1, 0, 0, 0, 16, 104, 1, 0, 0, 0, 18, 117, 1, 0, 0, 0, 20, 120,
		1, 0, 0, 0, 22, 132, 1, 0, 0, 0, 24, 145, 1, 0, 0, 0, 26, 147, 1, 0, 0,
		0, 28, 160, 1, 0, 0, 0, 30, 162, 1, 0, 0, 0, 32, 174, 1, 0, 0, 0, 34, 190,
		1, 0, 0, 0, 36, 192, 1, 0, 0, 0, 38, 194, 1, 0, 0, 0, 40, 198, 1, 0, 0,
		0, 42, 200, 1, 0, 0, 0, 44, 208, 1, 0, 0, 0, 46, 48, 3, 2, 1, 0, 47, 46,
		1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 55, 1, 0, 0, 0, 49, 54, 3, 4, 2, 0,
		50, 54, 3, 6, 3, 0, 51, 54, 3, 16, 8, 0, 52, 54, 3, 26, 13, 0, 53, 49,
		1, 0, 0, 0, 53, 50, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 52, 1, 0, 0, 0,
		54, 57, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 1, 1, 0,
		0, 0, 57, 55, 1, 0, 0, 0, 58, 59, 5, 4, 0, 0, 59, 60, 3, 38, 19, 0, 60,
		61, 5, 27, 0, 0, 61, 3, 1, 0, 0, 0, 62, 63, 5, 3, 0, 0, 63, 64, 5, 43,
		0, 0, 64, 65, 5, 27, 0, 0, 65, 5, 1, 0, 0, 0, 66, 67, 5, 26, 0, 0, 67,
		68, 5, 2, 0, 0, 68, 72, 5, 33, 0, 0, 69, 71, 3, 8, 4, 0, 70, 69, 1, 0,
		0, 0, 71, 74, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 75,
		1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 75, 76, 5, 34, 0, 0, 76, 7, 1, 0, 0, 0,
		77, 78, 3, 38, 19, 0, 78, 79, 5, 28, 0, 0, 79, 80, 3, 34, 17, 0, 80, 81,
		5, 27, 0, 0, 81, 9, 1, 0, 0, 0, 82, 83, 3, 38, 19, 0, 83, 84, 5, 39, 0,
		0, 84, 85, 3, 14, 7, 0, 85, 86, 5, 28, 0, 0, 86, 87, 3, 12, 6, 0, 87, 11,
		1, 0, 0, 0, 88, 89, 5, 46, 0, 0, 89, 13, 1, 0, 0, 0, 90, 103, 5, 5, 0,
		0, 91, 103, 5, 6, 0, 0, 92, 103, 5, 7, 0, 0, 93, 103, 5, 8, 0, 0, 94, 103,
		5, 9, 0, 0, 95, 103, 5, 10, 0, 0, 96, 103, 5, 11, 0, 0, 97, 103, 5, 12,
		0, 0, 98, 103, 5, 13, 0, 0, 99, 103, 5, 14, 0, 0, 100, 103, 5, 17, 0, 0,
		101, 103, 3, 20, 10, 0, 102, 90, 1, 0, 0, 0, 102, 91, 1, 0, 0, 0, 102,
		92, 1, 0, 0, 0, 102, 93, 1, 0, 0, 0, 102, 94, 1, 0, 0, 0, 102, 95, 1, 0,
		0, 0, 102, 96, 1, 0, 0, 0, 102, 97, 1, 0, 0, 0, 102, 98, 1, 0, 0, 0, 102,
		99, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 101, 1, 0, 0, 0, 103, 15, 1,
		0, 0, 0, 104, 105, 5, 18, 0, 0, 105, 106, 3, 38, 19, 0, 106, 110, 5, 33,
		0, 0, 107, 109, 3, 18, 9, 0, 108, 107, 1, 0, 0, 0, 109, 112, 1, 0, 0, 0,
		110, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 113, 1, 0, 0, 0, 112,
		110, 1, 0, 0, 0, 113, 114, 5, 34, 0, 0, 114, 17, 1, 0, 0, 0, 115, 118,
		3, 10, 5, 0, 116, 118, 3, 16, 8, 0, 117, 115, 1, 0, 0, 0, 117, 116, 1,
		0, 0, 0, 118, 19, 1, 0, 0, 0, 119, 121, 5, 37, 0, 0, 120, 119, 1, 0, 0,
		0, 120, 121, 1, 0, 0, 0, 121, 127, 1, 0, 0, 0, 122, 123, 3, 40, 20, 0,
		123, 124, 5, 37, 0, 0, 124, 126, 1, 0, 0, 0, 125, 122, 1, 0, 0, 0, 126,
		129, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 130,
		1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 130, 131, 3, 38, 19, 0, 131, 21, 1, 0,
		0, 0, 132, 133, 5, 22, 0, 0, 133, 134, 3, 38, 19, 0, 134, 138, 5, 33, 0,
		0, 135, 137, 3, 24, 12, 0, 136, 135, 1, 0, 0, 0, 137, 140, 1, 0, 0, 0,
		138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 141, 1, 0, 0, 0, 140,
		138, 1, 0, 0, 0, 141, 142, 5, 34, 0, 0, 142, 23, 1, 0, 0, 0, 143, 146,
		3, 10, 5, 0, 144, 146, 3, 22, 11, 0, 145, 143, 1, 0, 0, 0, 145, 144, 1,
		0, 0, 0, 146, 25, 1, 0, 0, 0, 147, 148, 5, 19, 0, 0, 148, 149, 3, 38, 19,
		0, 149, 153, 5, 33, 0, 0, 150, 152, 3, 28, 14, 0, 151, 150, 1, 0, 0, 0,
		152, 155, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154,
		156, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 157, 5, 34, 0, 0, 157, 27,
		1, 0, 0, 0, 158, 161, 3, 30, 15, 0, 159, 161, 3, 36, 18, 0, 160, 158, 1,
		0, 0, 0, 160, 159, 1, 0, 0, 0, 161, 29, 1, 0, 0, 0, 162, 163, 5, 20, 0,
		0, 163, 164, 3, 38, 19, 0, 164, 165, 5, 29, 0, 0, 165, 166, 3, 20, 10,
		0, 166, 167, 5, 30, 0, 0, 167, 170, 5, 1, 0, 0, 168, 171, 3, 20, 10, 0,
		169, 171, 5, 23, 0, 0, 170, 168, 1, 0, 0, 0, 170, 169, 1, 0, 0, 0, 171,
		172, 1, 0, 0, 0, 172, 173, 5, 27, 0, 0, 173, 31, 1, 0, 0, 0, 174, 175,
		3, 38, 19, 0, 175, 176, 5, 28, 0, 0, 176, 177, 3, 12, 6, 0, 177, 178, 5,
		27, 0, 0, 178, 33, 1, 0, 0, 0, 179, 191, 3, 42, 21, 0, 180, 182, 7, 0,
		0, 0, 181, 180, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0,
		183, 191, 5, 46, 0, 0, 184, 186, 7, 0, 0, 0, 185, 184, 1, 0, 0, 0, 185,
		186, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 191, 5, 45, 0, 0, 188, 191,
		5, 43, 0, 0, 189, 191, 5, 44, 0, 0, 190, 179, 1, 0, 0, 0, 190, 181, 1,
		0, 0, 0, 190, 185, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190, 189, 1, 0, 0,
		0, 191, 35, 1, 0, 0, 0, 192, 193, 5, 27, 0, 0, 193, 37, 1, 0, 0, 0, 194,
		195, 5, 47, 0, 0, 195, 39, 1, 0, 0, 0, 196, 199, 5, 47, 0, 0, 197, 199,
		3, 44, 22, 0, 198, 196, 1, 0, 0, 0, 198, 197, 1, 0, 0, 0, 199, 41, 1, 0,
		0, 0, 200, 205, 3, 40, 20, 0, 201, 202, 5, 37, 0, 0, 202, 204, 3, 40, 20,
		0, 203, 201, 1, 0, 0, 0, 204, 207, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 205,
		206, 1, 0, 0, 0, 206, 43, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 208, 209, 7,
		1, 0, 0, 209, 45, 1, 0, 0, 0, 19, 47, 53, 55, 72, 102, 110, 117, 120, 127,
		138, 145, 153, 160, 170, 181, 185, 190, 198, 205,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// YobufferParserInit initializes any static state used to implement YobufferParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewYobufferParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func YobufferParserInit() {
	staticData := &YobufferParserStaticData
	staticData.once.Do(yobufferParserInit)
}

// NewYobufferParser produces a new parser instance for the optional input antlr.TokenStream.
func NewYobufferParser(input antlr.TokenStream) *YobufferParser {
	YobufferParserInit()
	this := new(YobufferParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &YobufferParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Yobuffer.g4"

	return this
}

// YobufferParser tokens.
const (
	YobufferParserEOF          = antlr.TokenEOF
	YobufferParserT__0         = 1
	YobufferParserMETA         = 2
	YobufferParserIMPORT       = 3
	YobufferParserPACKAGE      = 4
	YobufferParserINT8         = 5
	YobufferParserINT16        = 6
	YobufferParserINT32        = 7
	YobufferParserINT64        = 8
	YobufferParserUINT8        = 9
	YobufferParserUINT16       = 10
	YobufferParserUINT32       = 11
	YobufferParserUINT64       = 12
	YobufferParserBOOL         = 13
	YobufferParserSTRING       = 14
	YobufferParserFLOAT32      = 15
	YobufferParserFLOAT64      = 16
	YobufferParserBYTES        = 17
	YobufferParserMESSAGE      = 18
	YobufferParserSERVICE      = 19
	YobufferParserRPC          = 20
	YobufferParserSTREAM       = 21
	YobufferParserSTRUCT       = 22
	YobufferParserVOID         = 23
	YobufferParserMAP          = 24
	YobufferParserLIST         = 25
	YobufferParserAT           = 26
	YobufferParserEND          = 27
	YobufferParserEQ           = 28
	YobufferParserLP           = 29
	YobufferParserRP           = 30
	YobufferParserLB           = 31
	YobufferParserRB           = 32
	YobufferParserLC           = 33
	YobufferParserRC           = 34
	YobufferParserLT           = 35
	YobufferParserGT           = 36
	YobufferParserDOT          = 37
	YobufferParserCOMMA        = 38
	YobufferParserCOLON        = 39
	YobufferParserPLUS         = 40
	YobufferParserMINUS        = 41
	YobufferParserNEWLINE      = 42
	YobufferParserSTRING_VALUE = 43
	YobufferParserBOOL_VALUE   = 44
	YobufferParserFLOAT_VALUE  = 45
	YobufferParserINT_VALUE    = 46
	YobufferParserIDENTIFIER   = 47
	YobufferParserWHITESPACE   = 48
	YobufferParserCOMMENT      = 49
)

// YobufferParser rules.
const (
	YobufferParserRULE_schema         = 0
	YobufferParserRULE_packageStmt    = 1
	YobufferParserRULE_importStmt     = 2
	YobufferParserRULE_meta           = 3
	YobufferParserRULE_metaElement    = 4
	YobufferParserRULE_field          = 5
	YobufferParserRULE_fieldTag       = 6
	YobufferParserRULE_fieldType      = 7
	YobufferParserRULE_message        = 8
	YobufferParserRULE_messageElement = 9
	YobufferParserRULE_messageType    = 10
	YobufferParserRULE_struct         = 11
	YobufferParserRULE_structElement  = 12
	YobufferParserRULE_service        = 13
	YobufferParserRULE_serviceElement = 14
	YobufferParserRULE_rpc            = 15
	YobufferParserRULE_rpcField       = 16
	YobufferParserRULE_constant       = 17
	YobufferParserRULE_emptyStmt      = 18
	YobufferParserRULE_name           = 19
	YobufferParserRULE_ident          = 20
	YobufferParserRULE_fullIdent      = 21
	YobufferParserRULE_keywords       = 22
)

// ISchemaContext is an interface to support dynamic dispatch.
type ISchemaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PackageStmt() IPackageStmtContext
	AllImportStmt() []IImportStmtContext
	ImportStmt(i int) IImportStmtContext
	AllMeta() []IMetaContext
	Meta(i int) IMetaContext
	AllMessage() []IMessageContext
	Message(i int) IMessageContext
	AllService() []IServiceContext
	Service(i int) IServiceContext

	// IsSchemaContext differentiates from other interfaces.
	IsSchemaContext()
}

type SchemaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaContext() *SchemaContext {
	var p = new(SchemaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_schema
	return p
}

func InitEmptySchemaContext(p *SchemaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_schema
}

func (*SchemaContext) IsSchemaContext() {}

func NewSchemaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaContext {
	var p = new(SchemaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_schema

	return p
}

func (s *SchemaContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaContext) PackageStmt() IPackageStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPackageStmtContext)
}

func (s *SchemaContext) AllImportStmt() []IImportStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportStmtContext); ok {
			len++
		}
	}

	tst := make([]IImportStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportStmtContext); ok {
			tst[i] = t.(IImportStmtContext)
			i++
		}
	}

	return tst
}

func (s *SchemaContext) ImportStmt(i int) IImportStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportStmtContext)
}

func (s *SchemaContext) AllMeta() []IMetaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMetaContext); ok {
			len++
		}
	}

	tst := make([]IMetaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMetaContext); ok {
			tst[i] = t.(IMetaContext)
			i++
		}
	}

	return tst
}

func (s *SchemaContext) Meta(i int) IMetaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaContext)
}

func (s *SchemaContext) AllMessage() []IMessageContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMessageContext); ok {
			len++
		}
	}

	tst := make([]IMessageContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMessageContext); ok {
			tst[i] = t.(IMessageContext)
			i++
		}
	}

	return tst
}

func (s *SchemaContext) Message(i int) IMessageContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *SchemaContext) AllService() []IServiceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IServiceContext); ok {
			len++
		}
	}

	tst := make([]IServiceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IServiceContext); ok {
			tst[i] = t.(IServiceContext)
			i++
		}
	}

	return tst
}

func (s *SchemaContext) Service(i int) IServiceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceContext)
}

func (s *SchemaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterSchema(s)
	}
}

func (s *SchemaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitSchema(s)
	}
}

func (s *SchemaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitSchema(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) Schema() (localctx ISchemaContext) {
	localctx = NewSchemaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, YobufferParserRULE_schema)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YobufferParserPACKAGE {
		{
			p.SetState(46)
			p.PackageStmt()
		}

	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67895304) != 0 {
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case YobufferParserIMPORT:
			{
				p.SetState(49)
				p.ImportStmt()
			}

		case YobufferParserAT:
			{
				p.SetState(50)
				p.Meta()
			}

		case YobufferParserMESSAGE:
			{
				p.SetState(51)
				p.Message()
			}

		case YobufferParserSERVICE:
			{
				p.SetState(52)
				p.Service()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPackageStmtContext is an interface to support dynamic dispatch.
type IPackageStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PACKAGE() antlr.TerminalNode
	Name() INameContext
	END() antlr.TerminalNode

	// IsPackageStmtContext differentiates from other interfaces.
	IsPackageStmtContext()
}

type PackageStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageStmtContext() *PackageStmtContext {
	var p = new(PackageStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_packageStmt
	return p
}

func InitEmptyPackageStmtContext(p *PackageStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_packageStmt
}

func (*PackageStmtContext) IsPackageStmtContext() {}

func NewPackageStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageStmtContext {
	var p = new(PackageStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_packageStmt

	return p
}

func (s *PackageStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageStmtContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(YobufferParserPACKAGE, 0)
}

func (s *PackageStmtContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *PackageStmtContext) END() antlr.TerminalNode {
	return s.GetToken(YobufferParserEND, 0)
}

func (s *PackageStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterPackageStmt(s)
	}
}

func (s *PackageStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitPackageStmt(s)
	}
}

func (s *PackageStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitPackageStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) PackageStmt() (localctx IPackageStmtContext) {
	localctx = NewPackageStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, YobufferParserRULE_packageStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(YobufferParserPACKAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.Name()
	}
	{
		p.SetState(60)
		p.Match(YobufferParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportStmtContext is an interface to support dynamic dispatch.
type IImportStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	STRING_VALUE() antlr.TerminalNode
	END() antlr.TerminalNode

	// IsImportStmtContext differentiates from other interfaces.
	IsImportStmtContext()
}

type ImportStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStmtContext() *ImportStmtContext {
	var p = new(ImportStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_importStmt
	return p
}

func InitEmptyImportStmtContext(p *ImportStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_importStmt
}

func (*ImportStmtContext) IsImportStmtContext() {}

func NewImportStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStmtContext {
	var p = new(ImportStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_importStmt

	return p
}

func (s *ImportStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStmtContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(YobufferParserIMPORT, 0)
}

func (s *ImportStmtContext) STRING_VALUE() antlr.TerminalNode {
	return s.GetToken(YobufferParserSTRING_VALUE, 0)
}

func (s *ImportStmtContext) END() antlr.TerminalNode {
	return s.GetToken(YobufferParserEND, 0)
}

func (s *ImportStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterImportStmt(s)
	}
}

func (s *ImportStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitImportStmt(s)
	}
}

func (s *ImportStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitImportStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) ImportStmt() (localctx IImportStmtContext) {
	localctx = NewImportStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, YobufferParserRULE_importStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(YobufferParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
		p.Match(YobufferParserSTRING_VALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Match(YobufferParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMetaContext is an interface to support dynamic dispatch.
type IMetaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	META() antlr.TerminalNode
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllMetaElement() []IMetaElementContext
	MetaElement(i int) IMetaElementContext

	// IsMetaContext differentiates from other interfaces.
	IsMetaContext()
}

type MetaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetaContext() *MetaContext {
	var p = new(MetaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_meta
	return p
}

func InitEmptyMetaContext(p *MetaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_meta
}

func (*MetaContext) IsMetaContext() {}

func NewMetaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetaContext {
	var p = new(MetaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_meta

	return p
}

func (s *MetaContext) GetParser() antlr.Parser { return s.parser }

func (s *MetaContext) AT() antlr.TerminalNode {
	return s.GetToken(YobufferParserAT, 0)
}

func (s *MetaContext) META() antlr.TerminalNode {
	return s.GetToken(YobufferParserMETA, 0)
}

func (s *MetaContext) LC() antlr.TerminalNode {
	return s.GetToken(YobufferParserLC, 0)
}

func (s *MetaContext) RC() antlr.TerminalNode {
	return s.GetToken(YobufferParserRC, 0)
}

func (s *MetaContext) AllMetaElement() []IMetaElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMetaElementContext); ok {
			len++
		}
	}

	tst := make([]IMetaElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMetaElementContext); ok {
			tst[i] = t.(IMetaElementContext)
			i++
		}
	}

	return tst
}

func (s *MetaContext) MetaElement(i int) IMetaElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaElementContext)
}

func (s *MetaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterMeta(s)
	}
}

func (s *MetaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitMeta(s)
	}
}

func (s *MetaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitMeta(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) Meta() (localctx IMetaContext) {
	localctx = NewMetaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, YobufferParserRULE_meta)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(YobufferParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Match(YobufferParserMETA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Match(YobufferParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == YobufferParserIDENTIFIER {
		{
			p.SetState(69)
			p.MetaElement()
		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(75)
		p.Match(YobufferParserRC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMetaElementContext is an interface to support dynamic dispatch.
type IMetaElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	EQ() antlr.TerminalNode
	Constant() IConstantContext
	END() antlr.TerminalNode

	// IsMetaElementContext differentiates from other interfaces.
	IsMetaElementContext()
}

type MetaElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetaElementContext() *MetaElementContext {
	var p = new(MetaElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_metaElement
	return p
}

func InitEmptyMetaElementContext(p *MetaElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_metaElement
}

func (*MetaElementContext) IsMetaElementContext() {}

func NewMetaElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetaElementContext {
	var p = new(MetaElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_metaElement

	return p
}

func (s *MetaElementContext) GetParser() antlr.Parser { return s.parser }

func (s *MetaElementContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *MetaElementContext) EQ() antlr.TerminalNode {
	return s.GetToken(YobufferParserEQ, 0)
}

func (s *MetaElementContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *MetaElementContext) END() antlr.TerminalNode {
	return s.GetToken(YobufferParserEND, 0)
}

func (s *MetaElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetaElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetaElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterMetaElement(s)
	}
}

func (s *MetaElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitMetaElement(s)
	}
}

func (s *MetaElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitMetaElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) MetaElement() (localctx IMetaElementContext) {
	localctx = NewMetaElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, YobufferParserRULE_metaElement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Name()
	}
	{
		p.SetState(78)
		p.Match(YobufferParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Constant()
	}
	{
		p.SetState(80)
		p.Match(YobufferParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	COLON() antlr.TerminalNode
	FieldType() IFieldTypeContext
	EQ() antlr.TerminalNode
	FieldTag() IFieldTagContext

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FieldContext) COLON() antlr.TerminalNode {
	return s.GetToken(YobufferParserCOLON, 0)
}

func (s *FieldContext) FieldType() IFieldTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *FieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(YobufferParserEQ, 0)
}

func (s *FieldContext) FieldTag() IFieldTagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldTagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldTagContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, YobufferParserRULE_field)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Name()
	}
	{
		p.SetState(83)
		p.Match(YobufferParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.FieldType()
	}
	{
		p.SetState(85)
		p.Match(YobufferParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.FieldTag()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldTagContext is an interface to support dynamic dispatch.
type IFieldTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_VALUE() antlr.TerminalNode

	// IsFieldTagContext differentiates from other interfaces.
	IsFieldTagContext()
}

type FieldTagContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldTagContext() *FieldTagContext {
	var p = new(FieldTagContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_fieldTag
	return p
}

func InitEmptyFieldTagContext(p *FieldTagContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_fieldTag
}

func (*FieldTagContext) IsFieldTagContext() {}

func NewFieldTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldTagContext {
	var p = new(FieldTagContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_fieldTag

	return p
}

func (s *FieldTagContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldTagContext) INT_VALUE() antlr.TerminalNode {
	return s.GetToken(YobufferParserINT_VALUE, 0)
}

func (s *FieldTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterFieldTag(s)
	}
}

func (s *FieldTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitFieldTag(s)
	}
}

func (s *FieldTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitFieldTag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) FieldTag() (localctx IFieldTagContext) {
	localctx = NewFieldTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, YobufferParserRULE_fieldTag)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(YobufferParserINT_VALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldTypeContext is an interface to support dynamic dispatch.
type IFieldTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT8() antlr.TerminalNode
	INT16() antlr.TerminalNode
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	UINT8() antlr.TerminalNode
	UINT16() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	BYTES() antlr.TerminalNode
	MessageType() IMessageTypeContext

	// IsFieldTypeContext differentiates from other interfaces.
	IsFieldTypeContext()
}

type FieldTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldTypeContext() *FieldTypeContext {
	var p = new(FieldTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_fieldType
	return p
}

func InitEmptyFieldTypeContext(p *FieldTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_fieldType
}

func (*FieldTypeContext) IsFieldTypeContext() {}

func NewFieldTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldTypeContext {
	var p = new(FieldTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_fieldType

	return p
}

func (s *FieldTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldTypeContext) INT8() antlr.TerminalNode {
	return s.GetToken(YobufferParserINT8, 0)
}

func (s *FieldTypeContext) INT16() antlr.TerminalNode {
	return s.GetToken(YobufferParserINT16, 0)
}

func (s *FieldTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(YobufferParserINT32, 0)
}

func (s *FieldTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(YobufferParserINT64, 0)
}

func (s *FieldTypeContext) UINT8() antlr.TerminalNode {
	return s.GetToken(YobufferParserUINT8, 0)
}

func (s *FieldTypeContext) UINT16() antlr.TerminalNode {
	return s.GetToken(YobufferParserUINT16, 0)
}

func (s *FieldTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(YobufferParserUINT32, 0)
}

func (s *FieldTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(YobufferParserUINT64, 0)
}

func (s *FieldTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(YobufferParserBOOL, 0)
}

func (s *FieldTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(YobufferParserSTRING, 0)
}

func (s *FieldTypeContext) BYTES() antlr.TerminalNode {
	return s.GetToken(YobufferParserBYTES, 0)
}

func (s *FieldTypeContext) MessageType() IMessageTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageTypeContext)
}

func (s *FieldTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterFieldType(s)
	}
}

func (s *FieldTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitFieldType(s)
	}
}

func (s *FieldTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitFieldType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) FieldType() (localctx IFieldTypeContext) {
	localctx = NewFieldTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, YobufferParserRULE_fieldType)
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.Match(YobufferParserINT8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.Match(YobufferParserINT16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(92)
			p.Match(YobufferParserINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(93)
			p.Match(YobufferParserINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(94)
			p.Match(YobufferParserUINT8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(95)
			p.Match(YobufferParserUINT16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(96)
			p.Match(YobufferParserUINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(97)
			p.Match(YobufferParserUINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(98)
			p.Match(YobufferParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(99)
			p.Match(YobufferParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(100)
			p.Match(YobufferParserBYTES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(101)
			p.MessageType()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMessageContext is an interface to support dynamic dispatch.
type IMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MESSAGE() antlr.TerminalNode
	Name() INameContext
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllMessageElement() []IMessageElementContext
	MessageElement(i int) IMessageElementContext

	// IsMessageContext differentiates from other interfaces.
	IsMessageContext()
}

type MessageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageContext() *MessageContext {
	var p = new(MessageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_message
	return p
}

func InitEmptyMessageContext(p *MessageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_message
}

func (*MessageContext) IsMessageContext() {}

func NewMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageContext {
	var p = new(MessageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_message

	return p
}

func (s *MessageContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(YobufferParserMESSAGE, 0)
}

func (s *MessageContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *MessageContext) LC() antlr.TerminalNode {
	return s.GetToken(YobufferParserLC, 0)
}

func (s *MessageContext) RC() antlr.TerminalNode {
	return s.GetToken(YobufferParserRC, 0)
}

func (s *MessageContext) AllMessageElement() []IMessageElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMessageElementContext); ok {
			len++
		}
	}

	tst := make([]IMessageElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMessageElementContext); ok {
			tst[i] = t.(IMessageElementContext)
			i++
		}
	}

	return tst
}

func (s *MessageContext) MessageElement(i int) IMessageElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageElementContext)
}

func (s *MessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterMessage(s)
	}
}

func (s *MessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitMessage(s)
	}
}

func (s *MessageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitMessage(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) Message() (localctx IMessageContext) {
	localctx = NewMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, YobufferParserRULE_message)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(YobufferParserMESSAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Name()
	}
	{
		p.SetState(106)
		p.Match(YobufferParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == YobufferParserMESSAGE || _la == YobufferParserIDENTIFIER {
		{
			p.SetState(107)
			p.MessageElement()
		}

		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(113)
		p.Match(YobufferParserRC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMessageElementContext is an interface to support dynamic dispatch.
type IMessageElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field() IFieldContext
	Message() IMessageContext

	// IsMessageElementContext differentiates from other interfaces.
	IsMessageElementContext()
}

type MessageElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageElementContext() *MessageElementContext {
	var p = new(MessageElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_messageElement
	return p
}

func InitEmptyMessageElementContext(p *MessageElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_messageElement
}

func (*MessageElementContext) IsMessageElementContext() {}

func NewMessageElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageElementContext {
	var p = new(MessageElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_messageElement

	return p
}

func (s *MessageElementContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageElementContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *MessageElementContext) Message() IMessageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *MessageElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterMessageElement(s)
	}
}

func (s *MessageElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitMessageElement(s)
	}
}

func (s *MessageElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitMessageElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) MessageElement() (localctx IMessageElementContext) {
	localctx = NewMessageElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, YobufferParserRULE_messageElement)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YobufferParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.Field()
		}

	case YobufferParserMESSAGE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.Message()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMessageTypeContext is an interface to support dynamic dispatch.
type IMessageTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext

	// IsMessageTypeContext differentiates from other interfaces.
	IsMessageTypeContext()
}

type MessageTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageTypeContext() *MessageTypeContext {
	var p = new(MessageTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_messageType
	return p
}

func InitEmptyMessageTypeContext(p *MessageTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_messageType
}

func (*MessageTypeContext) IsMessageTypeContext() {}

func NewMessageTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageTypeContext {
	var p = new(MessageTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_messageType

	return p
}

func (s *MessageTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageTypeContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *MessageTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(YobufferParserDOT)
}

func (s *MessageTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(YobufferParserDOT, i)
}

func (s *MessageTypeContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *MessageTypeContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *MessageTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterMessageType(s)
	}
}

func (s *MessageTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitMessageType(s)
	}
}

func (s *MessageTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitMessageType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) MessageType() (localctx IMessageTypeContext) {
	localctx = NewMessageTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, YobufferParserRULE_messageType)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YobufferParserDOT {
		{
			p.SetState(119)
			p.Match(YobufferParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(122)
				p.Ident()
			}
			{
				p.SetState(123)
				p.Match(YobufferParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.Name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructContext is an interface to support dynamic dispatch.
type IStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	Name() INameContext
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllStructElement() []IStructElementContext
	StructElement(i int) IStructElementContext

	// IsStructContext differentiates from other interfaces.
	IsStructContext()
}

type StructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructContext() *StructContext {
	var p = new(StructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_struct
	return p
}

func InitEmptyStructContext(p *StructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_struct
}

func (*StructContext) IsStructContext() {}

func NewStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructContext {
	var p = new(StructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_struct

	return p
}

func (s *StructContext) GetParser() antlr.Parser { return s.parser }

func (s *StructContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(YobufferParserSTRUCT, 0)
}

func (s *StructContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *StructContext) LC() antlr.TerminalNode {
	return s.GetToken(YobufferParserLC, 0)
}

func (s *StructContext) RC() antlr.TerminalNode {
	return s.GetToken(YobufferParserRC, 0)
}

func (s *StructContext) AllStructElement() []IStructElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructElementContext); ok {
			len++
		}
	}

	tst := make([]IStructElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructElementContext); ok {
			tst[i] = t.(IStructElementContext)
			i++
		}
	}

	return tst
}

func (s *StructContext) StructElement(i int) IStructElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructElementContext)
}

func (s *StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterStruct(s)
	}
}

func (s *StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitStruct(s)
	}
}

func (s *StructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) Struct_() (localctx IStructContext) {
	localctx = NewStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, YobufferParserRULE_struct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(YobufferParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Name()
	}
	{
		p.SetState(134)
		p.Match(YobufferParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == YobufferParserSTRUCT || _la == YobufferParserIDENTIFIER {
		{
			p.SetState(135)
			p.StructElement()
		}

		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(141)
		p.Match(YobufferParserRC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructElementContext is an interface to support dynamic dispatch.
type IStructElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field() IFieldContext
	Struct_() IStructContext

	// IsStructElementContext differentiates from other interfaces.
	IsStructElementContext()
}

type StructElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructElementContext() *StructElementContext {
	var p = new(StructElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_structElement
	return p
}

func InitEmptyStructElementContext(p *StructElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_structElement
}

func (*StructElementContext) IsStructElementContext() {}

func NewStructElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructElementContext {
	var p = new(StructElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_structElement

	return p
}

func (s *StructElementContext) GetParser() antlr.Parser { return s.parser }

func (s *StructElementContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *StructElementContext) Struct_() IStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructContext)
}

func (s *StructElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterStructElement(s)
	}
}

func (s *StructElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitStructElement(s)
	}
}

func (s *StructElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitStructElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) StructElement() (localctx IStructElementContext) {
	localctx = NewStructElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, YobufferParserRULE_structElement)
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YobufferParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(143)
			p.Field()
		}

	case YobufferParserSTRUCT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(144)
			p.Struct_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IServiceContext is an interface to support dynamic dispatch.
type IServiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SERVICE() antlr.TerminalNode
	Name() INameContext
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllServiceElement() []IServiceElementContext
	ServiceElement(i int) IServiceElementContext

	// IsServiceContext differentiates from other interfaces.
	IsServiceContext()
}

type ServiceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceContext() *ServiceContext {
	var p = new(ServiceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_service
	return p
}

func InitEmptyServiceContext(p *ServiceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_service
}

func (*ServiceContext) IsServiceContext() {}

func NewServiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceContext {
	var p = new(ServiceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_service

	return p
}

func (s *ServiceContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(YobufferParserSERVICE, 0)
}

func (s *ServiceContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ServiceContext) LC() antlr.TerminalNode {
	return s.GetToken(YobufferParserLC, 0)
}

func (s *ServiceContext) RC() antlr.TerminalNode {
	return s.GetToken(YobufferParserRC, 0)
}

func (s *ServiceContext) AllServiceElement() []IServiceElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IServiceElementContext); ok {
			len++
		}
	}

	tst := make([]IServiceElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IServiceElementContext); ok {
			tst[i] = t.(IServiceElementContext)
			i++
		}
	}

	return tst
}

func (s *ServiceContext) ServiceElement(i int) IServiceElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceElementContext)
}

func (s *ServiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterService(s)
	}
}

func (s *ServiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitService(s)
	}
}

func (s *ServiceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitService(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) Service() (localctx IServiceContext) {
	localctx = NewServiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, YobufferParserRULE_service)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(YobufferParserSERVICE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Name()
	}
	{
		p.SetState(149)
		p.Match(YobufferParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == YobufferParserRPC || _la == YobufferParserEND {
		{
			p.SetState(150)
			p.ServiceElement()
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(156)
		p.Match(YobufferParserRC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IServiceElementContext is an interface to support dynamic dispatch.
type IServiceElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Rpc() IRpcContext
	EmptyStmt() IEmptyStmtContext

	// IsServiceElementContext differentiates from other interfaces.
	IsServiceElementContext()
}

type ServiceElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceElementContext() *ServiceElementContext {
	var p = new(ServiceElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_serviceElement
	return p
}

func InitEmptyServiceElementContext(p *ServiceElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_serviceElement
}

func (*ServiceElementContext) IsServiceElementContext() {}

func NewServiceElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceElementContext {
	var p = new(ServiceElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_serviceElement

	return p
}

func (s *ServiceElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceElementContext) Rpc() IRpcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcContext)
}

func (s *ServiceElementContext) EmptyStmt() IEmptyStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStmtContext)
}

func (s *ServiceElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterServiceElement(s)
	}
}

func (s *ServiceElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitServiceElement(s)
	}
}

func (s *ServiceElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitServiceElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) ServiceElement() (localctx IServiceElementContext) {
	localctx = NewServiceElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, YobufferParserRULE_serviceElement)
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YobufferParserRPC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Rpc()
		}

	case YobufferParserEND:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			p.EmptyStmt()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRpcContext is an interface to support dynamic dispatch.
type IRpcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RPC() antlr.TerminalNode
	Name() INameContext
	LP() antlr.TerminalNode
	AllMessageType() []IMessageTypeContext
	MessageType(i int) IMessageTypeContext
	RP() antlr.TerminalNode
	END() antlr.TerminalNode
	VOID() antlr.TerminalNode

	// IsRpcContext differentiates from other interfaces.
	IsRpcContext()
}

type RpcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcContext() *RpcContext {
	var p = new(RpcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_rpc
	return p
}

func InitEmptyRpcContext(p *RpcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_rpc
}

func (*RpcContext) IsRpcContext() {}

func NewRpcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcContext {
	var p = new(RpcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_rpc

	return p
}

func (s *RpcContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcContext) RPC() antlr.TerminalNode {
	return s.GetToken(YobufferParserRPC, 0)
}

func (s *RpcContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *RpcContext) LP() antlr.TerminalNode {
	return s.GetToken(YobufferParserLP, 0)
}

func (s *RpcContext) AllMessageType() []IMessageTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMessageTypeContext); ok {
			len++
		}
	}

	tst := make([]IMessageTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMessageTypeContext); ok {
			tst[i] = t.(IMessageTypeContext)
			i++
		}
	}

	return tst
}

func (s *RpcContext) MessageType(i int) IMessageTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageTypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageTypeContext)
}

func (s *RpcContext) RP() antlr.TerminalNode {
	return s.GetToken(YobufferParserRP, 0)
}

func (s *RpcContext) END() antlr.TerminalNode {
	return s.GetToken(YobufferParserEND, 0)
}

func (s *RpcContext) VOID() antlr.TerminalNode {
	return s.GetToken(YobufferParserVOID, 0)
}

func (s *RpcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterRpc(s)
	}
}

func (s *RpcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitRpc(s)
	}
}

func (s *RpcContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitRpc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) Rpc() (localctx IRpcContext) {
	localctx = NewRpcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, YobufferParserRULE_rpc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(YobufferParserRPC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Name()
	}
	{
		p.SetState(164)
		p.Match(YobufferParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.MessageType()
	}
	{
		p.SetState(166)
		p.Match(YobufferParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Match(YobufferParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(168)
			p.MessageType()
		}

	case 2:
		{
			p.SetState(169)
			p.Match(YobufferParserVOID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(172)
		p.Match(YobufferParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRpcFieldContext is an interface to support dynamic dispatch.
type IRpcFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	EQ() antlr.TerminalNode
	FieldTag() IFieldTagContext
	END() antlr.TerminalNode

	// IsRpcFieldContext differentiates from other interfaces.
	IsRpcFieldContext()
}

type RpcFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcFieldContext() *RpcFieldContext {
	var p = new(RpcFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_rpcField
	return p
}

func InitEmptyRpcFieldContext(p *RpcFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_rpcField
}

func (*RpcFieldContext) IsRpcFieldContext() {}

func NewRpcFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcFieldContext {
	var p = new(RpcFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_rpcField

	return p
}

func (s *RpcFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcFieldContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *RpcFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(YobufferParserEQ, 0)
}

func (s *RpcFieldContext) FieldTag() IFieldTagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldTagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldTagContext)
}

func (s *RpcFieldContext) END() antlr.TerminalNode {
	return s.GetToken(YobufferParserEND, 0)
}

func (s *RpcFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterRpcField(s)
	}
}

func (s *RpcFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitRpcField(s)
	}
}

func (s *RpcFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitRpcField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) RpcField() (localctx IRpcFieldContext) {
	localctx = NewRpcFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, YobufferParserRULE_rpcField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Name()
	}
	{
		p.SetState(175)
		p.Match(YobufferParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.FieldTag()
	}
	{
		p.SetState(177)
		p.Match(YobufferParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FullIdent() IFullIdentContext
	INT_VALUE() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	FLOAT_VALUE() antlr.TerminalNode
	STRING_VALUE() antlr.TerminalNode
	BOOL_VALUE() antlr.TerminalNode

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) FullIdent() IFullIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *ConstantContext) INT_VALUE() antlr.TerminalNode {
	return s.GetToken(YobufferParserINT_VALUE, 0)
}

func (s *ConstantContext) MINUS() antlr.TerminalNode {
	return s.GetToken(YobufferParserMINUS, 0)
}

func (s *ConstantContext) PLUS() antlr.TerminalNode {
	return s.GetToken(YobufferParserPLUS, 0)
}

func (s *ConstantContext) FLOAT_VALUE() antlr.TerminalNode {
	return s.GetToken(YobufferParserFLOAT_VALUE, 0)
}

func (s *ConstantContext) STRING_VALUE() antlr.TerminalNode {
	return s.GetToken(YobufferParserSTRING_VALUE, 0)
}

func (s *ConstantContext) BOOL_VALUE() antlr.TerminalNode {
	return s.GetToken(YobufferParserBOOL_VALUE, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, YobufferParserRULE_constant)
	var _la int

	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(179)
			p.FullIdent()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == YobufferParserPLUS || _la == YobufferParserMINUS {
			{
				p.SetState(180)
				_la = p.GetTokenStream().LA(1)

				if !(_la == YobufferParserPLUS || _la == YobufferParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(183)
			p.Match(YobufferParserINT_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == YobufferParserPLUS || _la == YobufferParserMINUS {
			{
				p.SetState(184)
				_la = p.GetTokenStream().LA(1)

				if !(_la == YobufferParserPLUS || _la == YobufferParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(187)
			p.Match(YobufferParserFLOAT_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(188)
			p.Match(YobufferParserSTRING_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(189)
			p.Match(YobufferParserBOOL_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEmptyStmtContext is an interface to support dynamic dispatch.
type IEmptyStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	END() antlr.TerminalNode

	// IsEmptyStmtContext differentiates from other interfaces.
	IsEmptyStmtContext()
}

type EmptyStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyStmtContext() *EmptyStmtContext {
	var p = new(EmptyStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_emptyStmt
	return p
}

func InitEmptyEmptyStmtContext(p *EmptyStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_emptyStmt
}

func (*EmptyStmtContext) IsEmptyStmtContext() {}

func NewEmptyStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStmtContext {
	var p = new(EmptyStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_emptyStmt

	return p
}

func (s *EmptyStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *EmptyStmtContext) END() antlr.TerminalNode {
	return s.GetToken(YobufferParserEND, 0)
}

func (s *EmptyStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterEmptyStmt(s)
	}
}

func (s *EmptyStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitEmptyStmt(s)
	}
}

func (s *EmptyStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitEmptyStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) EmptyStmt() (localctx IEmptyStmtContext) {
	localctx = NewEmptyStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, YobufferParserRULE_emptyStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(YobufferParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(YobufferParserIDENTIFIER, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitName(s)
	}
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, YobufferParserRULE_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.Match(YobufferParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Keywords() IKeywordsContext

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_ident
	return p
}

func InitEmptyIdentContext(p *IdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_ident
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(YobufferParserIDENTIFIER, 0)
}

func (s *IdentContext) Keywords() IKeywordsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordsContext)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitIdent(s)
	}
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) Ident() (localctx IIdentContext) {
	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, YobufferParserRULE_ident)
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YobufferParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(196)
			p.Match(YobufferParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case YobufferParserIMPORT, YobufferParserPACKAGE, YobufferParserINT8, YobufferParserINT16, YobufferParserINT32, YobufferParserINT64, YobufferParserUINT8, YobufferParserUINT16, YobufferParserUINT32, YobufferParserUINT64, YobufferParserBOOL, YobufferParserSTRING, YobufferParserFLOAT32, YobufferParserFLOAT64, YobufferParserBYTES, YobufferParserMESSAGE, YobufferParserSERVICE, YobufferParserRPC, YobufferParserVOID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(197)
			p.Keywords()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFullIdentContext is an interface to support dynamic dispatch.
type IFullIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsFullIdentContext differentiates from other interfaces.
	IsFullIdentContext()
}

type FullIdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullIdentContext() *FullIdentContext {
	var p = new(FullIdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_fullIdent
	return p
}

func InitEmptyFullIdentContext(p *FullIdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_fullIdent
}

func (*FullIdentContext) IsFullIdentContext() {}

func NewFullIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullIdentContext {
	var p = new(FullIdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_fullIdent

	return p
}

func (s *FullIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *FullIdentContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *FullIdentContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FullIdentContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(YobufferParserDOT)
}

func (s *FullIdentContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(YobufferParserDOT, i)
}

func (s *FullIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterFullIdent(s)
	}
}

func (s *FullIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitFullIdent(s)
	}
}

func (s *FullIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitFullIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) FullIdent() (localctx IFullIdentContext) {
	localctx = NewFullIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, YobufferParserRULE_fullIdent)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Ident()
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == YobufferParserDOT {
		{
			p.SetState(201)
			p.Match(YobufferParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)
			p.Ident()
		}

		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeywordsContext is an interface to support dynamic dispatch.
type IKeywordsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	PACKAGE() antlr.TerminalNode
	INT8() antlr.TerminalNode
	INT16() antlr.TerminalNode
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	UINT8() antlr.TerminalNode
	UINT16() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	FLOAT32() antlr.TerminalNode
	FLOAT64() antlr.TerminalNode
	BYTES() antlr.TerminalNode
	MESSAGE() antlr.TerminalNode
	SERVICE() antlr.TerminalNode
	RPC() antlr.TerminalNode
	VOID() antlr.TerminalNode

	// IsKeywordsContext differentiates from other interfaces.
	IsKeywordsContext()
}

type KeywordsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordsContext() *KeywordsContext {
	var p = new(KeywordsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_keywords
	return p
}

func InitEmptyKeywordsContext(p *KeywordsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YobufferParserRULE_keywords
}

func (*KeywordsContext) IsKeywordsContext() {}

func NewKeywordsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordsContext {
	var p = new(KeywordsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YobufferParserRULE_keywords

	return p
}

func (s *KeywordsContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordsContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(YobufferParserIMPORT, 0)
}

func (s *KeywordsContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(YobufferParserPACKAGE, 0)
}

func (s *KeywordsContext) INT8() antlr.TerminalNode {
	return s.GetToken(YobufferParserINT8, 0)
}

func (s *KeywordsContext) INT16() antlr.TerminalNode {
	return s.GetToken(YobufferParserINT16, 0)
}

func (s *KeywordsContext) INT32() antlr.TerminalNode {
	return s.GetToken(YobufferParserINT32, 0)
}

func (s *KeywordsContext) INT64() antlr.TerminalNode {
	return s.GetToken(YobufferParserINT64, 0)
}

func (s *KeywordsContext) UINT8() antlr.TerminalNode {
	return s.GetToken(YobufferParserUINT8, 0)
}

func (s *KeywordsContext) UINT16() antlr.TerminalNode {
	return s.GetToken(YobufferParserUINT16, 0)
}

func (s *KeywordsContext) UINT32() antlr.TerminalNode {
	return s.GetToken(YobufferParserUINT32, 0)
}

func (s *KeywordsContext) UINT64() antlr.TerminalNode {
	return s.GetToken(YobufferParserUINT64, 0)
}

func (s *KeywordsContext) BOOL() antlr.TerminalNode {
	return s.GetToken(YobufferParserBOOL, 0)
}

func (s *KeywordsContext) STRING() antlr.TerminalNode {
	return s.GetToken(YobufferParserSTRING, 0)
}

func (s *KeywordsContext) FLOAT32() antlr.TerminalNode {
	return s.GetToken(YobufferParserFLOAT32, 0)
}

func (s *KeywordsContext) FLOAT64() antlr.TerminalNode {
	return s.GetToken(YobufferParserFLOAT64, 0)
}

func (s *KeywordsContext) BYTES() antlr.TerminalNode {
	return s.GetToken(YobufferParserBYTES, 0)
}

func (s *KeywordsContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(YobufferParserMESSAGE, 0)
}

func (s *KeywordsContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(YobufferParserSERVICE, 0)
}

func (s *KeywordsContext) RPC() antlr.TerminalNode {
	return s.GetToken(YobufferParserRPC, 0)
}

func (s *KeywordsContext) VOID() antlr.TerminalNode {
	return s.GetToken(YobufferParserVOID, 0)
}

func (s *KeywordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.EnterKeywords(s)
	}
}

func (s *KeywordsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YobufferListener); ok {
		listenerT.ExitKeywords(s)
	}
}

func (s *KeywordsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case YobufferVisitor:
		return t.VisitKeywords(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *YobufferParser) Keywords() (localctx IKeywordsContext) {
	localctx = NewKeywordsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, YobufferParserRULE_keywords)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&10485752) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
