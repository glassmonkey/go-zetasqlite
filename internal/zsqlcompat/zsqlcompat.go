// Package zsqlcompat collects the few residual fork-specific aliases that
// don't naturally belong upstream. Most of what once lived here has been
// adopted by zetasql-wasm proper (named-type enums in zetasql/types/
// resolved_ast, BaseFunctionCall + ExprType + ScanColumnList helpers,
// TypeFromKind / TypeFromProto, the LanguageFeature / ParameterMode /
// ProductMode / NameResolutionMode / ParseLocationRecordType named types,
// the Cardinality / SignatureArgumentKind constants, etc.).
//
// What remains is fork-specific configuration; if more arrives it should
// land here only when it's clearly not a candidate for upstream.
package zsqlcompat

import "github.com/glassmonkey/zetasql-wasm/wasm/generated"

// SupportedStatementKinds returns the resolved-node statement kinds that
// go-zetasqlite asks the analyzer to accept. The exact list is a
// fork-side product decision (which DDL/DML/script-control statements
// the SQLite-backed engine handles), so it lives here rather than in
// zetasql-wasm.
func SupportedStatementKinds() []generated.ResolvedNodeKind {
	return []generated.ResolvedNodeKind{
		generated.ResolvedNodeKind_RESOLVED_BEGIN_STMT,
		generated.ResolvedNodeKind_RESOLVED_COMMIT_STMT,
		generated.ResolvedNodeKind_RESOLVED_MERGE_STMT,
		generated.ResolvedNodeKind_RESOLVED_QUERY_STMT,
		generated.ResolvedNodeKind_RESOLVED_INSERT_STMT,
		generated.ResolvedNodeKind_RESOLVED_UPDATE_STMT,
		generated.ResolvedNodeKind_RESOLVED_DELETE_STMT,
		generated.ResolvedNodeKind_RESOLVED_DROP_STMT,
		generated.ResolvedNodeKind_RESOLVED_TRUNCATE_STMT,
		generated.ResolvedNodeKind_RESOLVED_CREATE_TABLE_STMT,
		generated.ResolvedNodeKind_RESOLVED_CREATE_TABLE_AS_SELECT_STMT,
		generated.ResolvedNodeKind_RESOLVED_CREATE_PROCEDURE_STMT,
		generated.ResolvedNodeKind_RESOLVED_CREATE_FUNCTION_STMT,
		generated.ResolvedNodeKind_RESOLVED_CREATE_TABLE_FUNCTION_STMT,
		generated.ResolvedNodeKind_RESOLVED_CREATE_VIEW_STMT,
		generated.ResolvedNodeKind_RESOLVED_DROP_FUNCTION_STMT,
	}
}
