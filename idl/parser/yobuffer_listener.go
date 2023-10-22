// Code generated from Yobuffer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Yobuffer
import "github.com/antlr4-go/antlr/v4"

// YobufferListener is a complete listener for a parse tree produced by YobufferParser.
type YobufferListener interface {
	antlr.ParseTreeListener

	// EnterSchema is called when entering the schema production.
	EnterSchema(c *SchemaContext)

	// EnterPackageStmt is called when entering the packageStmt production.
	EnterPackageStmt(c *PackageStmtContext)

	// EnterImportStmt is called when entering the importStmt production.
	EnterImportStmt(c *ImportStmtContext)

	// EnterMeta is called when entering the meta production.
	EnterMeta(c *MetaContext)

	// EnterMetaElement is called when entering the metaElement production.
	EnterMetaElement(c *MetaElementContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldTag is called when entering the fieldTag production.
	EnterFieldTag(c *FieldTagContext)

	// EnterFieldType is called when entering the fieldType production.
	EnterFieldType(c *FieldTypeContext)

	// EnterMessage is called when entering the message production.
	EnterMessage(c *MessageContext)

	// EnterMessageElement is called when entering the messageElement production.
	EnterMessageElement(c *MessageElementContext)

	// EnterMessageType is called when entering the messageType production.
	EnterMessageType(c *MessageTypeContext)

	// EnterStruct is called when entering the struct production.
	EnterStruct(c *StructContext)

	// EnterStructElement is called when entering the structElement production.
	EnterStructElement(c *StructElementContext)

	// EnterService is called when entering the service production.
	EnterService(c *ServiceContext)

	// EnterServiceElement is called when entering the serviceElement production.
	EnterServiceElement(c *ServiceElementContext)

	// EnterRpc is called when entering the rpc production.
	EnterRpc(c *RpcContext)

	// EnterRpcField is called when entering the rpcField production.
	EnterRpcField(c *RpcFieldContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterEmptyStmt is called when entering the emptyStmt production.
	EnterEmptyStmt(c *EmptyStmtContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterFullIdent is called when entering the fullIdent production.
	EnterFullIdent(c *FullIdentContext)

	// EnterKeywords is called when entering the keywords production.
	EnterKeywords(c *KeywordsContext)

	// ExitSchema is called when exiting the schema production.
	ExitSchema(c *SchemaContext)

	// ExitPackageStmt is called when exiting the packageStmt production.
	ExitPackageStmt(c *PackageStmtContext)

	// ExitImportStmt is called when exiting the importStmt production.
	ExitImportStmt(c *ImportStmtContext)

	// ExitMeta is called when exiting the meta production.
	ExitMeta(c *MetaContext)

	// ExitMetaElement is called when exiting the metaElement production.
	ExitMetaElement(c *MetaElementContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldTag is called when exiting the fieldTag production.
	ExitFieldTag(c *FieldTagContext)

	// ExitFieldType is called when exiting the fieldType production.
	ExitFieldType(c *FieldTypeContext)

	// ExitMessage is called when exiting the message production.
	ExitMessage(c *MessageContext)

	// ExitMessageElement is called when exiting the messageElement production.
	ExitMessageElement(c *MessageElementContext)

	// ExitMessageType is called when exiting the messageType production.
	ExitMessageType(c *MessageTypeContext)

	// ExitStruct is called when exiting the struct production.
	ExitStruct(c *StructContext)

	// ExitStructElement is called when exiting the structElement production.
	ExitStructElement(c *StructElementContext)

	// ExitService is called when exiting the service production.
	ExitService(c *ServiceContext)

	// ExitServiceElement is called when exiting the serviceElement production.
	ExitServiceElement(c *ServiceElementContext)

	// ExitRpc is called when exiting the rpc production.
	ExitRpc(c *RpcContext)

	// ExitRpcField is called when exiting the rpcField production.
	ExitRpcField(c *RpcFieldContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitEmptyStmt is called when exiting the emptyStmt production.
	ExitEmptyStmt(c *EmptyStmtContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitFullIdent is called when exiting the fullIdent production.
	ExitFullIdent(c *FullIdentContext)

	// ExitKeywords is called when exiting the keywords production.
	ExitKeywords(c *KeywordsContext)
}
