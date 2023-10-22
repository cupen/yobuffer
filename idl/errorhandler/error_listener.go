package errorhandler

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

var ErrSytax = fmt.Errorf("syntax error")

type ReturnErrorListener struct {
	*antlr.DefaultErrorListener
}

func NewReturnErrorListener() *ReturnErrorListener {
	return &ReturnErrorListener{}
}

func (rel *ReturnErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	p, ok := recognizer.(*antlr.BaseParser)
	if !ok {
		panic(fmt.Errorf("%w: line %d:%d: %v", ErrSytax, line, column, msg))
	}

	currTok := p.GetCurrentToken()
	tokLine := currTok.GetLine()
	tokCol := currTok.GetColumn()
	curTok := currTok.GetText()

	err := fmt.Errorf("%w: line %d:%d: unexpected '%v': %s", ErrSytax, tokLine, tokCol, curTok, msg)

	expected := p.GetExpectedTokens().StringVerbose(p.LiteralNames, p.SymbolicNames, false)
	if expected != "" {
		err = fmt.Errorf("%w missing '%s'", err, expected)
	}
	panic(err)
}
