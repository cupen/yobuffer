package errorhandler

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type ReturnErrorStrategy struct {
	*antlr.DefaultErrorStrategy
	Err        error
	Hint       string
	ErrContext antlr.ParserRuleContext
	RE         antlr.RecognitionException

	litNames, symbNames []string
}

var _ antlr.ErrorStrategy = &ReturnErrorStrategy{}

func NewReturnErrorStrategy(litNames, symbNames []string) *ReturnErrorStrategy {
	res := &ReturnErrorStrategy{
		litNames:  litNames,
		symbNames: symbNames,
	}
	res.DefaultErrorStrategy = antlr.NewDefaultErrorStrategy()
	return res
}

func (res *ReturnErrorStrategy) Recover(recognizer antlr.Parser, e antlr.RecognitionException) {
	if res.Err == nil {
		res.RE = e
		res.Err = fmt.Errorf("invalid syntax")
		res.ErrContext = recognizer.GetParserRuleContext()
		expected := recognizer.GetExpectedTokens().StringVerbose(res.litNames, res.symbNames, false)
		res.Hint = fmt.Sprintf("missing '%s'", expected)
		switch expected {
		case "EQUALS":
			res.Hint += " shit2"
		}
	}
	context := recognizer.GetParserRuleContext()
	for context != nil {
		context.SetException(e)
		var ok bool
		context, ok = context.GetParent().(antlr.ParserRuleContext)
		if !ok {
			break
		}
	}
}

func (res *ReturnErrorStrategy) RecoverInline(recognizer antlr.Parser) antlr.Token {
	res.Recover(recognizer, antlr.NewInputMisMatchException(recognizer))
	return recognizer.GetCurrentToken()
}

func (res *ReturnErrorStrategy) Sync(recognizer antlr.Parser) {
}
