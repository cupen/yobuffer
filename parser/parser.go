package parser

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/cupen/yobuffer/parser/ast"
)

type Parser struct {
	*BaseYobufferListener
	c            *ast.Context
	structDef    *ast.Struct
	meta         *ast.Meta
	service      *ast.Service
	serviceField *ast.RPC
}

func New() *Parser {
	return &Parser{
		c: &ast.Context{},
	}
}

func (p *Parser) Parse(input []byte) (*ast.Context, error) {
	stream := antlr.NewInputStream(string(input))
	lexer := NewYobufferLexer(stream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	ps := NewYobufferParser(tokens)
	antlr.ParseTreeWalkerDefault.Walk(p, ps.PackageStmt())
	return p.c, nil
}

func (p *Parser) EnterPackageStmt(ctx *PackageStmtContext) {
	if p.c.Package != nil {
		panic("package already defined")
	}
	if ctx.END() == nil {
		panic("package statement not end with ';'")
	}

	// name:= ctx.PACKAGE().GetText()
	name := ctx.FullIdent().GetText()
	p.c.Package = &ast.Package{
		Name: name,
	}
}

func (p *Parser) EnterService(ctx *ServiceContext) {
	ctx.GetTokens(YobufferParserRULE_serviceName)
	ctx.SERVICE()
	if p.service != nil {
		panic("service already defined")
	}
	p.service = &ast.Service{
		Name: ctx.ServiceName().GetText(),
	}
}

// ExitService is called when production service is exited.
func (p *Parser) ExitService(ctx *ServiceContext) {
	p.c.Services = append(p.c.Services, p.service)
	p.service = nil
}
