// Code generated from Yobuffer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Yobuffer
import "github.com/antlr4-go/antlr/v4"

// BaseYobufferListener is a complete listener for a parse tree produced by YobufferParser.
type BaseYobufferListener struct{}

var _ YobufferListener = &BaseYobufferListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseYobufferListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseYobufferListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseYobufferListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseYobufferListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSchema is called when production schema is entered.
func (s *BaseYobufferListener) EnterSchema(ctx *SchemaContext) {}

// ExitSchema is called when production schema is exited.
func (s *BaseYobufferListener) ExitSchema(ctx *SchemaContext) {}

// EnterPackageStmt is called when production packageStmt is entered.
func (s *BaseYobufferListener) EnterPackageStmt(ctx *PackageStmtContext) {}

// ExitPackageStmt is called when production packageStmt is exited.
func (s *BaseYobufferListener) ExitPackageStmt(ctx *PackageStmtContext) {}

// EnterImportStmt is called when production importStmt is entered.
func (s *BaseYobufferListener) EnterImportStmt(ctx *ImportStmtContext) {}

// ExitImportStmt is called when production importStmt is exited.
func (s *BaseYobufferListener) ExitImportStmt(ctx *ImportStmtContext) {}

// EnterMeta is called when production meta is entered.
func (s *BaseYobufferListener) EnterMeta(ctx *MetaContext) {}

// ExitMeta is called when production meta is exited.
func (s *BaseYobufferListener) ExitMeta(ctx *MetaContext) {}

// EnterMetaName is called when production metaName is entered.
func (s *BaseYobufferListener) EnterMetaName(ctx *MetaNameContext) {}

// ExitMetaName is called when production metaName is exited.
func (s *BaseYobufferListener) ExitMetaName(ctx *MetaNameContext) {}

// EnterMetaElement is called when production metaElement is entered.
func (s *BaseYobufferListener) EnterMetaElement(ctx *MetaElementContext) {}

// ExitMetaElement is called when production metaElement is exited.
func (s *BaseYobufferListener) ExitMetaElement(ctx *MetaElementContext) {}

// EnterField is called when production field is entered.
func (s *BaseYobufferListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseYobufferListener) ExitField(ctx *FieldContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseYobufferListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseYobufferListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterFieldTag is called when production fieldTag is entered.
func (s *BaseYobufferListener) EnterFieldTag(ctx *FieldTagContext) {}

// ExitFieldTag is called when production fieldTag is exited.
func (s *BaseYobufferListener) ExitFieldTag(ctx *FieldTagContext) {}

// EnterFieldType is called when production fieldType is entered.
func (s *BaseYobufferListener) EnterFieldType(ctx *FieldTypeContext) {}

// ExitFieldType is called when production fieldType is exited.
func (s *BaseYobufferListener) ExitFieldType(ctx *FieldTypeContext) {}

// EnterMessage is called when production message is entered.
func (s *BaseYobufferListener) EnterMessage(ctx *MessageContext) {}

// ExitMessage is called when production message is exited.
func (s *BaseYobufferListener) ExitMessage(ctx *MessageContext) {}

// EnterMessageName is called when production messageName is entered.
func (s *BaseYobufferListener) EnterMessageName(ctx *MessageNameContext) {}

// ExitMessageName is called when production messageName is exited.
func (s *BaseYobufferListener) ExitMessageName(ctx *MessageNameContext) {}

// EnterMessageElement is called when production messageElement is entered.
func (s *BaseYobufferListener) EnterMessageElement(ctx *MessageElementContext) {}

// ExitMessageElement is called when production messageElement is exited.
func (s *BaseYobufferListener) ExitMessageElement(ctx *MessageElementContext) {}

// EnterMessageType is called when production messageType is entered.
func (s *BaseYobufferListener) EnterMessageType(ctx *MessageTypeContext) {}

// ExitMessageType is called when production messageType is exited.
func (s *BaseYobufferListener) ExitMessageType(ctx *MessageTypeContext) {}

// EnterStruct is called when production struct is entered.
func (s *BaseYobufferListener) EnterStruct(ctx *StructContext) {}

// ExitStruct is called when production struct is exited.
func (s *BaseYobufferListener) ExitStruct(ctx *StructContext) {}

// EnterStructName is called when production structName is entered.
func (s *BaseYobufferListener) EnterStructName(ctx *StructNameContext) {}

// ExitStructName is called when production structName is exited.
func (s *BaseYobufferListener) ExitStructName(ctx *StructNameContext) {}

// EnterStructElement is called when production structElement is entered.
func (s *BaseYobufferListener) EnterStructElement(ctx *StructElementContext) {}

// ExitStructElement is called when production structElement is exited.
func (s *BaseYobufferListener) ExitStructElement(ctx *StructElementContext) {}

// EnterService is called when production service is entered.
func (s *BaseYobufferListener) EnterService(ctx *ServiceContext) {}

// ExitService is called when production service is exited.
func (s *BaseYobufferListener) ExitService(ctx *ServiceContext) {}

// EnterServiceName is called when production serviceName is entered.
func (s *BaseYobufferListener) EnterServiceName(ctx *ServiceNameContext) {}

// ExitServiceName is called when production serviceName is exited.
func (s *BaseYobufferListener) ExitServiceName(ctx *ServiceNameContext) {}

// EnterServiceElement is called when production serviceElement is entered.
func (s *BaseYobufferListener) EnterServiceElement(ctx *ServiceElementContext) {}

// ExitServiceElement is called when production serviceElement is exited.
func (s *BaseYobufferListener) ExitServiceElement(ctx *ServiceElementContext) {}

// EnterRpc is called when production rpc is entered.
func (s *BaseYobufferListener) EnterRpc(ctx *RpcContext) {}

// ExitRpc is called when production rpc is exited.
func (s *BaseYobufferListener) ExitRpc(ctx *RpcContext) {}

// EnterRpcName is called when production rpcName is entered.
func (s *BaseYobufferListener) EnterRpcName(ctx *RpcNameContext) {}

// ExitRpcName is called when production rpcName is exited.
func (s *BaseYobufferListener) ExitRpcName(ctx *RpcNameContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseYobufferListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseYobufferListener) ExitConstant(ctx *ConstantContext) {}

// EnterEmptyStmt is called when production emptyStmt is entered.
func (s *BaseYobufferListener) EnterEmptyStmt(ctx *EmptyStmtContext) {}

// ExitEmptyStmt is called when production emptyStmt is exited.
func (s *BaseYobufferListener) ExitEmptyStmt(ctx *EmptyStmtContext) {}

// EnterIdent is called when production ident is entered.
func (s *BaseYobufferListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseYobufferListener) ExitIdent(ctx *IdentContext) {}

// EnterFullIdent is called when production fullIdent is entered.
func (s *BaseYobufferListener) EnterFullIdent(ctx *FullIdentContext) {}

// ExitFullIdent is called when production fullIdent is exited.
func (s *BaseYobufferListener) ExitFullIdent(ctx *FullIdentContext) {}

// EnterKeywords is called when production keywords is entered.
func (s *BaseYobufferListener) EnterKeywords(ctx *KeywordsContext) {}

// ExitKeywords is called when production keywords is exited.
func (s *BaseYobufferListener) ExitKeywords(ctx *KeywordsContext) {}
