// Code generated from Yobuffer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Yobuffer
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by YobufferParser.
type YobufferVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by YobufferParser#schema.
	VisitSchema(ctx *SchemaContext) interface{}

	// Visit a parse tree produced by YobufferParser#packageStmt.
	VisitPackageStmt(ctx *PackageStmtContext) interface{}

	// Visit a parse tree produced by YobufferParser#importStmt.
	VisitImportStmt(ctx *ImportStmtContext) interface{}

	// Visit a parse tree produced by YobufferParser#meta.
	VisitMeta(ctx *MetaContext) interface{}

	// Visit a parse tree produced by YobufferParser#metaElement.
	VisitMetaElement(ctx *MetaElementContext) interface{}

	// Visit a parse tree produced by YobufferParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by YobufferParser#fieldTag.
	VisitFieldTag(ctx *FieldTagContext) interface{}

	// Visit a parse tree produced by YobufferParser#fieldType.
	VisitFieldType(ctx *FieldTypeContext) interface{}

	// Visit a parse tree produced by YobufferParser#message.
	VisitMessage(ctx *MessageContext) interface{}

	// Visit a parse tree produced by YobufferParser#messageElement.
	VisitMessageElement(ctx *MessageElementContext) interface{}

	// Visit a parse tree produced by YobufferParser#messageType.
	VisitMessageType(ctx *MessageTypeContext) interface{}

	// Visit a parse tree produced by YobufferParser#struct.
	VisitStruct(ctx *StructContext) interface{}

	// Visit a parse tree produced by YobufferParser#structElement.
	VisitStructElement(ctx *StructElementContext) interface{}

	// Visit a parse tree produced by YobufferParser#service.
	VisitService(ctx *ServiceContext) interface{}

	// Visit a parse tree produced by YobufferParser#serviceElement.
	VisitServiceElement(ctx *ServiceElementContext) interface{}

	// Visit a parse tree produced by YobufferParser#rpc.
	VisitRpc(ctx *RpcContext) interface{}

	// Visit a parse tree produced by YobufferParser#rpcField.
	VisitRpcField(ctx *RpcFieldContext) interface{}

	// Visit a parse tree produced by YobufferParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by YobufferParser#emptyStmt.
	VisitEmptyStmt(ctx *EmptyStmtContext) interface{}

	// Visit a parse tree produced by YobufferParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by YobufferParser#ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by YobufferParser#fullIdent.
	VisitFullIdent(ctx *FullIdentContext) interface{}

	// Visit a parse tree produced by YobufferParser#keywords.
	VisitKeywords(ctx *KeywordsContext) interface{}
}
