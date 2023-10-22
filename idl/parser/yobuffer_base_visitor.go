// Code generated from Yobuffer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Yobuffer
import "github.com/antlr4-go/antlr/v4"

type BaseYobufferVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseYobufferVisitor) VisitSchema(ctx *SchemaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitPackageStmt(ctx *PackageStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitImportStmt(ctx *ImportStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitMeta(ctx *MetaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitMetaElement(ctx *MetaElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitFieldTag(ctx *FieldTagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitFieldType(ctx *FieldTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitMessage(ctx *MessageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitMessageElement(ctx *MessageElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitMessageType(ctx *MessageTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitStruct(ctx *StructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitStructElement(ctx *StructElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitService(ctx *ServiceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitServiceElement(ctx *ServiceElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitRpc(ctx *RpcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitRpcField(ctx *RpcFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitEmptyStmt(ctx *EmptyStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitFullIdent(ctx *FullIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseYobufferVisitor) VisitKeywords(ctx *KeywordsContext) interface{} {
	return v.VisitChildren(ctx)
}
