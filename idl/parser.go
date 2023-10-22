package idl

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/cupen/yobuffer/idl/ast"
	"github.com/cupen/yobuffer/idl/errorhandler"
	"github.com/cupen/yobuffer/idl/parser"
	"golang.org/x/exp/slog"
)

type Parser struct {
	*parser.BaseYobufferListener
	c            *ast.Context
	structDef    *ast.Struct
	meta         *ast.Meta
	service      *ast.Service
	serviceField *ast.RPC
	err          antlr.RecognitionException
}

func New() *Parser {
	return &Parser{
		c: &ast.Context{},
	}
}

func (p *Parser) Parse(input []byte) (c *ast.Context, err error) {
	defer func() {
		if _err := recover(); _err != nil {
			var ok bool
			err, ok = _err.(error)
			if !ok {
				err = fmt.Errorf("%v", _err)
			}
		}
	}()
	c, err = p.parse(input)
	return
}

func (p *Parser) parse(input []byte) (*ast.Context, error) {
	stream := antlr.NewInputStream(string(input))
	lexer := parser.NewYobufferLexer(stream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	ps := parser.NewYobufferParser(tokens)

	errorListener := errorhandler.NewReturnErrorListener()

	lexerMetas := &parser.YobufferLexerLexerStaticData
	errorStrategy := errorhandler.NewReturnErrorStrategy(lexerMetas.LiteralNames, lexerMetas.SymbolicNames)
	_ = errorStrategy

	ps.RemoveErrorListeners()
	ps.AddErrorListener(errorListener)
	antlr.ParseTreeWalkerDefault.Walk(p, ps.Schema())
	if ps.HasError() {
		e := ps.GetError()
		return nil, fmt.Errorf("%v", e)
	}
	return p.c, nil
}

func (p *Parser) EnterSchema(ctx *parser.SchemaContext) {
	slog.Info("EnterSchema")
	if ctx.GetParser().HasError() {
		e := ctx.GetParser().GetError()
		slog.Error(fmt.Sprintf("%v", e))
	}
}

func (p *Parser) EnterPackageStmt(ctx *parser.PackageStmtContext) {
	slog.Info("EnterPackage")
	if p.c.Package != nil {
		// panic("package already defined")
	}
	end := ctx.END()
	if end == nil || end.GetText() != ";" {
		// panic("package statement should end with ';'")
	}

	name := ctx.Name().GetText()
	p.c.Package = &ast.Package{
		Name: name,
	}
}

func (p *Parser) EnterService(ctx *parser.ServiceContext) {
	slog.Info("EnterService")
	name := ctx.Name().GetText()
	length := ctx.SERVICE().GetChildCount()

	rpcs := []*ast.RPC{}
	for i := 0; i < length; i++ {
		elem := ctx.ServiceElement(i)
		rpc := elem.Rpc()

		rpcs = append(rpcs, &ast.RPC{
			Name:   rpc.Name().GetText(),
			Fields: nil,
		})
	}
	p.service = &ast.Service{
		Name: name,
		RPCs: rpcs,
	}
}

func (p *Parser) parseRpc(ctx *parser.RpcContext) *ast.RPC {
	name := ctx.Name().GetText()
	return &ast.RPC{
		Name: name,
	}
}

func (p *Parser) parseFields(ctx *parser.FieldContext) *ast.Field {
	name := ctx.Name().GetText()
	_type := ctx.FieldType().GetText()
	_tag := ctx.FieldTag().GetText()
	tag, err := strconv.Atoi(_tag)
	if err != nil {
		panic(err)
	}
	return &ast.Field{
		Name: name,
		Type: _type,
		Tag:  tag,
	}
}

func (p *Parser) ExitService(ctx *parser.ServiceContext) {
}
