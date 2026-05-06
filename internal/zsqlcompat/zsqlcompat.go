// Package zsqlcompat holds fork-private aliases for resolved-AST enum
// constants and statement-kind lists that aren't yet exposed by name from
// zetasql-wasm. Most of what used to live here has been moved upstream
// (LanguageFeature, ParameterMode, ProductMode, NameResolutionMode,
// ParseLocationRecordType, Cardinality, SignatureArgumentKind constants,
// TypeFromKind / TypeFromProto helpers, BaseFunctionCall interface,
// ExprType / ScanColumnList helpers, and TypeKind.String). The remaining
// content is generator-side enums on resolved_ast nodes that still expose
// the proto type directly via accessor return values; until that
// generator update lands, the named constants stay here so go-zetasqlite
// can reference them by short name in switch statements.
package zsqlcompat

import (
	"github.com/glassmonkey/zetasql-wasm/wasm/generated"
)

// === Resolved AST enum aliases ===
//
// Each block aliases a generator-output proto enum so call sites can use
// short names like `JoinTypeFull` instead of
// `generated.ResolvedJoinScanEnums_FULL`. When the resolved_astgen tool is
// updated to emit named-type enums upstream, these blocks delete and the
// imports point to resolved_ast directly.

// JoinType.
const (
	JoinTypeInner = generated.ResolvedJoinScanEnums_INNER
	JoinTypeLeft  = generated.ResolvedJoinScanEnums_LEFT
	JoinTypeRight = generated.ResolvedJoinScanEnums_RIGHT
	JoinTypeFull  = generated.ResolvedJoinScanEnums_FULL
)

// Window-frame unit.
type FrameUnit = generated.ResolvedWindowFrameEnums_FrameUnit

const (
	FrameUnitRows  = generated.ResolvedWindowFrameEnums_ROWS
	FrameUnitRange = generated.ResolvedWindowFrameEnums_RANGE
)

// Create-statement mode and scope.
type CreateMode = generated.ResolvedCreateStatementEnums_CreateMode

const (
	CreateDefaultMode     = generated.ResolvedCreateStatementEnums_CREATE_DEFAULT
	CreateOrReplaceMode   = generated.ResolvedCreateStatementEnums_CREATE_OR_REPLACE
	CreateIfNotExistsMode = generated.ResolvedCreateStatementEnums_CREATE_IF_NOT_EXISTS
	CreateScopeTemp       = generated.ResolvedCreateStatementEnums_CREATE_TEMP
)

// Merge-when match type and action type.
const (
	MatchTypeMatched            = generated.ResolvedMergeWhenEnums_MATCHED
	MatchTypeNotMatchedBySource = generated.ResolvedMergeWhenEnums_NOT_MATCHED_BY_SOURCE
	MatchTypeNotMatchedByTarget = generated.ResolvedMergeWhenEnums_NOT_MATCHED_BY_TARGET

	ActionTypeInsert = generated.ResolvedMergeWhenEnums_INSERT
	ActionTypeUpdate = generated.ResolvedMergeWhenEnums_UPDATE
	ActionTypeDelete = generated.ResolvedMergeWhenEnums_DELETE
)

// Subquery type.
const (
	SubqueryTypeScalar  = generated.ResolvedSubqueryExprEnums_SCALAR
	SubqueryTypeArray   = generated.ResolvedSubqueryExprEnums_ARRAY
	SubqueryTypeExists  = generated.ResolvedSubqueryExprEnums_EXISTS
	SubqueryTypeIn      = generated.ResolvedSubqueryExprEnums_IN
	SubqueryTypeLikeAny = generated.ResolvedSubqueryExprEnums_LIKE_ANY
	SubqueryTypeLikeAll = generated.ResolvedSubqueryExprEnums_LIKE_ALL
)

// NULL-order mode for ORDER BY.
const (
	NullOrderModeNullsFirst = generated.ResolvedOrderByItemEnums_NULLS_FIRST
	NullOrderModeNullsLast  = generated.ResolvedOrderByItemEnums_NULLS_LAST
)

// Set-operation type.
const (
	SetOperationTypeUnionAll          = generated.ResolvedSetOperationScanEnums_UNION_ALL
	SetOperationTypeUnionDistinct     = generated.ResolvedSetOperationScanEnums_UNION_DISTINCT
	SetOperationTypeIntersectAll      = generated.ResolvedSetOperationScanEnums_INTERSECT_ALL
	SetOperationTypeIntersectDistinct = generated.ResolvedSetOperationScanEnums_INTERSECT_DISTINCT
	SetOperationTypeExceptAll         = generated.ResolvedSetOperationScanEnums_EXCEPT_ALL
	SetOperationTypeExceptDistinct    = generated.ResolvedSetOperationScanEnums_EXCEPT_DISTINCT
)

// Window-frame boundary.
type BoundaryType = generated.ResolvedWindowFrameExprEnums_BoundaryType

const (
	UnboundedPrecedingType = generated.ResolvedWindowFrameExprEnums_UNBOUNDED_PRECEDING
	OffsetPrecedingType    = generated.ResolvedWindowFrameExprEnums_OFFSET_PRECEDING
	CurrentRowType         = generated.ResolvedWindowFrameExprEnums_CURRENT_ROW
	OffsetFollowingType    = generated.ResolvedWindowFrameExprEnums_OFFSET_FOLLOWING
	UnboundedFollowingType = generated.ResolvedWindowFrameExprEnums_UNBOUNDED_FOLLOWING
)

// Null-handling modifier (resolved AST variant — used on aggregate/analytic calls).
const (
	IgnoreNulls  = generated.ResolvedNonScalarFunctionCallBaseEnums_IGNORE_NULLS
	RespectNulls = generated.ResolvedNonScalarFunctionCallBaseEnums_RESPECT_NULLS
)

// Function-call error mode.
const SafeErrorMode = generated.ResolvedFunctionCallBaseEnums_SAFE_ERROR_MODE

// SupportedStatementKinds returns the resolved-node statement kinds that
// go-zetasqlite asks the analyzer to accept. The list is fork-specific so
// it stays here rather than moving upstream.
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
