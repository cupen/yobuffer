package parser

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/cupen/yobuffer/codegen/models"
	"github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
)

var (
	ErrNoStructDefined = fmt.Errorf("no struct defined")
)

var (
	regexStructField = regexp.MustCompile("^[ \t]*(\\w+)[ \t]+(\\w+)[ \t]*=[ \t]*(\\d+)[ \t]*;[ \t\r\n]*$")
	regexRPC         = regexp.MustCompile("^[ \t]*rpc[ \t]+(\\w+)[ \t]*(\\([^\\)]*\\))(.*)$")
	// regexWord        = regexp.MustCompile("[a-zA-Z_]+")
)

type idlParser struct {
	lexer  *lexmachine.Lexer
	ctx    *models.Context
	states struct {
		actions int
		_struct *models.Struct
		_fields map[string]struct{}
		meta    *models.Meta
		rpc     *models.RPC
	}
}

func New() *idlParser {
	obj := &idlParser{
		lexer: lexmachine.NewLexer(),
	}
	lexer := obj.lexer
	lexer.Add([]byte("\\s*//[^\n]*\n*"), obj.actionSkip)
	lexer.Add([]byte("package \\w+\n+"), obj.tokenBeginPackage)
	lexer.Add([]byte("\\s*@meta\\([^\n]*\\)\n+"), obj.tokenMeta)
	lexer.Add([]byte("\\s*global_\\w+\\([^\n]+\\)\n+"), obj.tokenGlobal)
	lexer.Add([]byte("\\s*struct *\\w+\\ *{\\s*\n+"), obj.tokenStructBegin)
	lexer.Add([]byte("\\s*\\w+[ \t]+\\w+[ \t]+= [1-9]+;\\s*\n+"), obj.tokenStructFields)
	lexer.Add([]byte("\\s*\\}\n+"), obj.tokenStructEnd)
	lexer.Add([]byte("\\s*rpc +\\w+ *\\([^\n]*\\)\\s+\\([^\n]+\\)\n+"), obj.tokenRPC)
	lexer.Add([]byte("\\s*\\w+:\\w+\\s*\n+"), obj.tokenArgs)
	lexer.Add([]byte("\\s*external\\s*+struct\\s+(\\w+)\n+"), obj.tokenExternalStruct)
	if err := lexer.Compile(); err != nil {
		panic(err)
	}
	return obj
}

func (this *idlParser) receiveTokens(token interface{}) error {
	this.states.actions++
	switch tk := token.(type) {
	case *BeginPackage:
		if this.ctx.Package != "" {
			return fmt.Errorf("already declared package %s", this.ctx.Package)
		}
		this.ctx.Package = strings.Trim(tk.Name, " \t")
		return nil
	case *StructBegin:
		return this.receiveStruct(tk, nil, nil)
	case *StructField:
		return this.receiveStruct(nil, tk, nil)
	case *StructEnd:
		return this.receiveStruct(nil, nil, tk)
	case *Meta:
		return this.receiveMeta(tk)
	case *RPC:
		return this.receiveRPC(tk)
	case *ExternalStruct:
		return this.receiveExternalStruct(tk)

	default:
		return fmt.Errorf("invalid token: %v", reflect.TypeOf(token))
	}
}

func (this *idlParser) Parse(data []byte) (*models.Context, error) {
	c := models.NewContext()
	this.ctx = c

	s, err := this.lexer.Scanner(data)
	if err != nil {
		return nil, err
	}
	for token, err, eos := s.Next(); !eos; token, err, eos = s.Next() {
		if err != nil {
			log.Printf("fail: token:%+v err:%v", token, err)
			return nil, err
		}
		if token == nil {
			continue
		}
		if err := this.receiveTokens(token); err != nil {
			log.Printf("fail: receive token:%+v failed. err:%v", reflect.TypeOf(token), err)
			return nil, err
		}
	}
	return c, nil
}

func show(caseName string, s *lexmachine.Scanner, m *machines.Match) {
	// log.Printf("%s | scanner:%v match:%s", caseName, string(s.Text), string(m.Bytes))
	log.Printf("%24s | %s", caseName, string(trim(m.Bytes)))
}

func (this *idlParser) actionSkip(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
	show("comment", s, m)
	return nil, nil
}

func (this *idlParser) tokenBeginPackage(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
	show("package", s, m)
	trim(m.Bytes)
	name := string(trim(m.Bytes[len("package"):]))
	return &BeginPackage{Name: name}, nil
}

func (this *idlParser) tokenMeta(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
	show("meta", s, m)
	line := trim(m.Bytes[len("@meta("):])
	line = line[0 : len(line)-1]

	// TODO: 待优化
	fields := map[string]string{}
	for _, l := range strings.Split(string(line), ",") {
		tmpArr := strings.Split(l, "=")
		if len(tmpArr) != 2 {
			return nil, Errorf(m, "invalid meta:%s", string(m.Bytes))
		}
		k, v := trims(tmpArr[0]), trims(tmpArr[1])
		if _, ok := fields[k]; ok {
			return nil, Errorf(m, "duplicated field:%s", k)
		}
		fields[k] = v
	}
	return &Meta{Fields: fields}, nil
}

func (this *idlParser) tokenGlobal(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
	show("global", s, m)
	return nil, nil
}

func (this *idlParser) tokenStructBegin(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
	line := string(m.Bytes)
	line = strings.TrimPrefix(line, "struct")
	line = strings.Trim(line, " \t{\n\r")
	show("begin-struct", s, m)
	return &StructBegin{Name: line}, nil
}

func (this *idlParser) tokenStructEnd(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
	_case := "end-struct"
	show(_case, s, m)
	return &StructEnd{}, nil
}

func (this *idlParser) tokenRPC(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
	show("rpc", s, m)
	p := regexRPC
	matchesAll := p.FindStringSubmatch(string(trim(m.Bytes)))
	if len(matchesAll) != 4 {
		return nil, Errorf(m, "syntax error %#v", matchesAll)
	}
	matches := matchesAll[1:]

	fields := []*StructField{}
	line := strings.ReplaceAll(matches[1], " ", "")
	line = strings.ReplaceAll(line, "\t", "")
	line = strings.Trim(line, "() \t")
	if line != "" {
		for _, l := range strings.Split(line, ",") {
			kv := strings.Split(l, ":")
			fname := kv[0]
			if fname == "" {
				return nil, Errorf(m, "empty field name")
			}
			ftype := ""
			if len(kv) == 2 {
				ftype = kv[1]
			}
			fields = append(fields, &StructField{
				Name: fname,
				Type: ftype,
			})
		}
	}

	// if args := matches[1]; args != "" {
	// 	sot := 0 // start of token
	// 	eot := 0 // end of token
	// 	lastToken := ""
	// 	fname := ""
	// 	ftype := ""
	// 	for i := 0; i < len(args); i++ {
	// 		switch args[i] {
	// 		case '(':
	// 			sot = i
	// 			continue
	// 		case ' ', '\t':
	// 			if sot != 0 {
	// 				fname = args[sot:i]
	// 				sot = 0
	// 			}
	// 			continue
	// 		case ':':
	// 			if sot == 0 {
	// 				return nil, Errorf(m, "invalid ':'")
	// 			}
	// 			continue
	// 		case ',': // next field
	// 		default:
	// 			sot = i
	// 		}

	// 		fields = append(fields, &StructField{
	// 			Name: fname,
	// 			Type: ftype,
	// 		})
	// 	}
	// }
	// offset, err := strconv.Atoi(matches[1])
	// if err != nil {
	// 	return nil, Errorf(m, err.Error())
	// }
	returnType := strings.Trim(matches[2], " \t")
	if returnType == "" {
		return nil, Errorf(m, "invalid return '%s'", returnType)
	}
	if returnType[0] != '(' || returnType[len(returnType)-1] != ')' {
		return nil, Errorf(m, "missing return of rpc")
	}
	returnType = strings.Trim(returnType, "() \t")
	if returnType == "" {
		return nil, Errorf(m, "invalid return '%s'", returnType)
	}
	return &RPC{
		Name:   string(matches[0]),
		Fields: fields,
		Return: returnType,
	}, nil
}

func (this *idlParser) tokenStructFields(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
	show("field", s, m)
	p := regexStructField
	matchesAll := p.FindStringSubmatch(string(m.Bytes))
	if len(matchesAll) != 4 {
		return nil, Errorf(m, "syntax error %#v", matchesAll)
	}
	matches := matchesAll[1:]
	offset, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, Errorf(m, err.Error())
	}
	return &StructField{
		Type:   strings.ToLower(matches[0]),
		Name:   strings.Title(matches[1]),
		Offset: offset,
	}, nil
}

func (this *idlParser) tokenArgs(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
	show("args", s, m)
	return nil, nil
}

func (this *idlParser) tokenExternalStruct(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
	show("external struct", s, m)
	line := m.Bytes
	legalSyntax := "external struct"
	if len(line) < len(legalSyntax) {
		return nil, Errorf(m, "is this '%s'?", legalSyntax)
	}
	for i := 0; i > len(legalSyntax); i++ {
		if line[i] != legalSyntax[i] {
			return nil, Errorf(m, "is this '%s'? too many spaces or tabs?", legalSyntax)
		}
	}
	line = trim(line[len(legalSyntax)+1:])
	return &ExternalStruct{
		Name: string(line),
	}, nil
}

func trim(s []byte) []byte {
	return bytes.Trim(s, "\t\n")
}

func trims(s string) string {
	return strings.Trim(s, " \t\n")
}

func (this *idlParser) getFieldsMap() map[string]struct{} {
	if this.states._fields == nil {
		this.states._fields = map[string]struct{}{}
	}
	return this.states._fields
}

func (this *idlParser) receiveMeta(meta *Meta) error {
	if this.states.meta != nil {
		return fmt.Errorf("invalid meta")
	}
	this.states.meta = &models.Meta{
		Fields: meta.Fields,
	}
	return nil
}

func (this *idlParser) receiveRPC(rpc *RPC) error {
	if this.states.meta == nil {
		return fmt.Errorf("missing meta of rpc '%s'", rpc.Name)
	}
	meta := this.states.meta
	this.states.meta = nil

	id, err := meta.GetInt("id")
	if err != nil {
		return err
	}
	desc, err := meta.GetString("desc")
	if err != nil {
		return err
	}
	args := []*models.Field{}
	for _, f := range rpc.Fields {
		if f.Type == "" {
			f.Type = "string"
		}
		args = append(args, &models.Field{
			Name:   f.Name,
			Type:   f.Type,
			Offset: 0,
		})
	}
	var arg *models.Field
	if len(args) > 0 {
		arg = args[0]
	}

	r := &models.RPC{
		ID:     id,
		Name:   rpc.Name,
		Desc:   desc,
		Args:   args,
		Arg:    arg,
		Return: rpc.Return,
	}
	return this.ctx.DefineRPC(r)
}

func (this *idlParser) receiveStruct(begin *StructBegin, field *StructField, end *StructEnd) error {
	state := &this.states
	if begin != nil {
		if state.meta == nil {
			return fmt.Errorf("missing meta for struct '%s'", begin.Name)
		}
		if state._struct != nil {
			return fmt.Errorf("'%s' is not ended", state._struct)
		}
		id, err := state.meta.GetInt("id")
		if err != nil {
			return fmt.Errorf("meta.id %w", err)
		}
		desc, err := state.meta.GetString("desc")
		if err != nil {
			return fmt.Errorf("meta.desc %w", err)
		}
		encodingName, err := state.meta.GetString("encoding")
		if err != nil {
			return fmt.Errorf("meta.encoding %w", err)
		}
		obj := models.Struct{
			ID:       id,
			Name:     begin.Name,
			Desc:     desc,
			Encoding: encodingName,
		}
		state._struct = &obj
		// return this.ctx.DefineStruct(&obj)
	}

	if field != nil {
		if state._struct == nil {
			return ErrNoStructDefined
		}

		fields := this.getFieldsMap()
		if _, ok := fields[field.Name]; ok {
			return fmt.Errorf("duplicated field")
		}

		f := models.Field{
			Name:   field.Name,
			Type:   field.Type,
			Offset: field.Offset,
		}
		state._struct.Fields = append(state._struct.Fields, &f)
		fields[field.Name] = struct{}{}
	}

	if end != nil {
		if state._struct == nil {
			return ErrNoStructDefined
		}
		s := state._struct
		state._struct = nil
		state._fields = nil
		state.meta = nil
		return this.ctx.DefineStruct(s)
	}
	return nil
}

func (this *idlParser) receiveExternalStruct(e *ExternalStruct) error {
	state := &this.states
	if state.meta == nil {
		return fmt.Errorf("missing meta for ExternalStruct %s", e.Name)
	}
	id, err := state.meta.GetInt("id")
	if err != nil {
		return err
	}
	obj := models.ExternalStruct{
		ID:   id,
		Name: e.Name,
	}
	state.meta = nil
	return this.ctx.DefineExternalStruct(&obj)
}
