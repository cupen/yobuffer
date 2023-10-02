grammar Yobuffer;


// Schema
schema: packageStmt? (importStmt | meta | message | service)*;


// Package
packageStmt: PACKAGE fullIdent END;


// Import
importStmt: IMPORT STRING_VALUE END;


// Meta
meta : AT META LC metaElement* RC END;
metaName: ident;
metaElement : metaName EQ constant END;


// Field
field: (meta)? fieldType fieldName EQ fieldTag END;
fieldName: ident;
fieldTag: INT_VALUE;

// field types
fieldType
  : INT8
  | INT16
  | INT32
  | INT64
  | UINT8
  | UINT16
  | UINT32
  | UINT64
  | BOOL
  | STRING
  | BYTES
  | messageType
  ;


// Message
message: MESSAGE messageName LC messageElement* RC;
messageName: ident;
messageElement: field | message;
messageType: ( DOT )? ( ident DOT )* messageName;


// Struct
struct: STRUCT structName LC structElement* RC;
structName: ident;
structElement: field | struct;
// structField: (meta)? fieldType fieldName EQ fieldTag END;
// structFieldName: ident;
// structFieldTag: INT_VALUE;


// Service
service: SERVICE serviceName LC serviceElement* RC;
serviceName: ident;
serviceElement : rpc | emptyStmt;
rpc: RPC rpcName LP messageType RP COLON (messageType | VOID) END;
rpcName: ident;


// lexical
constant
  : fullIdent
  | (MINUS | PLUS)? INT_VALUE 
  | (MINUS | PLUS)? FLOAT_VALUE 
  | STRING_VALUE
  | BOOL_VALUE
  ;

emptyStmt: END;


// Lexical elements
ident: IDENTIFIER | keywords;
fullIdent: ident ( DOT ident )*;


// Keywords
keywords
  : IMPORT
  | PACKAGE
  | INT8
  | INT16
  | INT32
  | INT64
  | UINT8
  | UINT16
  | UINT32
  | UINT64
  | BOOL
  | STRING
  | FLOAT32
  | FLOAT64
  | BYTES
  | MESSAGE
  | SERVICE
  | RPC
  // | STREAM
  | VOID
  ;

META: 'meta';
IMPORT: 'import';
PACKAGE: 'package';
INT8: 'int8';
INT16: 'int16';
INT32: 'int32';
INT64: 'int64';
UINT8: 'uint8';
UINT16: 'uint16';
UINT32: 'uint32';
UINT64: 'uint64';
BOOL: 'bool';
STRING: 'string';
FLOAT32: 'float32';
FLOAT64: 'float64';
BYTES: 'bytes';
MESSAGE: 'message';
SERVICE: 'service';
RPC: 'rpc';
STREAM: 'stream';
STRUCT: 'struct';
VOID: 'void';


// Symbols
AT: '@';
END: ';';
EQ: '=';
LP: '(';
RP: ')';
LB: '[';
RB: ']';
LC: '{';
RC: '}';
LT: '<';
GT: '>';
DOT: '.';
COMMA: ',';
COLON: ':';
PLUS: '+';
MINUS: '-';


STRING_VALUE: ( '\'' ( CHAR_VALUE )* '\'' ) |  ( '"' ( CHAR_VALUE )* '"' );
fragment CHAR_VALUE: HEX_ESCAPE | OCT_ESCAPE | CHAR_ESCAPE | ~[\u0000\n\\];
fragment HEX_ESCAPE: '\\' ( 'x' | 'X' ) HEX_DIGIT HEX_DIGIT;
fragment OCT_ESCAPE: '\\' OCTAL_DIGIT OCTAL_DIGIT OCTAL_DIGIT;
fragment CHAR_ESCAPE: '\\' ( 'a' | 'b' | 'f' | 'n' | 'r' | 't' | 'v' | '\\' | '\'' | '"' );

BOOL_VALUE: 'true' | 'false';

FLOAT_VALUE: ( DECIMALS DOT DECIMALS? EXPONENT? | DECIMALS EXPONENT | DOT DECIMALS EXPONENT? ) | 'inf' | 'nan';
fragment EXPONENT  : ( 'e' | 'E' ) (PLUS | MINUS)? DECIMALS;
fragment DECIMALS  : DECIMAL_DIGIT+;

INT_VALUE: DECIMAL_LIT | OCTAL_LIT | HEX_LIT;
fragment DECIMAL_LIT : ( [1-9] ) DECIMAL_DIGIT*;
fragment OCTAL_LIT   : '0' OCTAL_DIGIT*;
fragment HEX_LIT     : '0' ( 'x' | 'X' ) HEX_DIGIT+ ;

IDENTIFIER: LETTER ( LETTER | DECIMAL_DIGIT )*;

fragment LETTER: [A-Za-z_];
fragment DECIMAL_DIGIT: [0-9];
fragment OCTAL_DIGIT: [0-7];
fragment HEX_DIGIT: [0-9A-Fa-f];

// comments
WHITESPACE:   [ \t\r\n\u000C]+ -> skip;
COMMENT: '//' ~[\r\n]* -> channel(HIDDEN);

