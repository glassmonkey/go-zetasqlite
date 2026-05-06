// Package zsqlcompat re-exports go-zetasql-style names for symbols that
// are spread across zetasql-wasm packages but were originally available
// under github.com/goccy/go-zetasql{,/types,/resolved_ast}.
//
// This package only exists to keep the go-zetasqlite migration tractable.
// New code should use the zetasql-wasm packages directly.
package zsqlcompat

import (
	"github.com/glassmonkey/zetasql-wasm/resolved_ast"
	"github.com/glassmonkey/zetasql-wasm/types"
	"github.com/glassmonkey/zetasql-wasm/wasm/generated"
)

// BaseFunctionCall is the common surface that go-zetasql exposed via its
// ResolvedBaseFunctionCall pseudo-base. zetasql-wasm flattens that out: each
// of FunctionCallNode / AggregateFunctionCallNode / AnalyticFunctionCallNode
// declares the same four methods directly. This interface lets fork code
// continue to share helpers across all three.
type BaseFunctionCall interface {
	ArgumentList() []resolved_ast.ExprNode
	Function() *generated.FunctionRefProto
	Signature() *generated.FunctionSignatureProto
	ErrorMode() generated.ResolvedFunctionCallBaseEnums_ErrorMode
}

// ExprType pulls the *generated.TypeProto off any resolved expression node
// that carries one. Most concrete ExprNode implementations expose a
// `Type() *generated.TypeProto` method that the resolved_ast.ExprNode
// interface itself doesn't surface; this helper does the type assertion
// once and returns nil for the rare cases that don't.
func ExprType(e resolved_ast.ExprNode) *generated.TypeProto {
	if t, ok := e.(interface{ Type() *generated.TypeProto }); ok {
		return t.Type()
	}
	return nil
}

// LanguageFeature is the language-feature flag type.
type LanguageFeature = generated.LanguageFeature

// LanguageFeature constants used by go-zetasqlite.
const (
	FeatureAnalyticFunctions             LanguageFeature = generated.LanguageFeature_FEATURE_ANALYTIC_FUNCTIONS
	FeatureBignumericType                LanguageFeature = generated.LanguageFeature_FEATURE_BIGNUMERIC_TYPE
	FeatureCreateTableAsSelectColumnList LanguageFeature = generated.LanguageFeature_FEATURE_CREATE_TABLE_AS_SELECT_COLUMN_LIST
	FeatureCreateTableNotNull            LanguageFeature = generated.LanguageFeature_FEATURE_CREATE_TABLE_NOT_NULL
	FeatureGeography                     LanguageFeature = generated.LanguageFeature_FEATURE_GEOGRAPHY
	FeatureGroupByRollup                 LanguageFeature = generated.LanguageFeature_FEATURE_GROUP_BY_ROLLUP
	FeatureIntervalType                  LanguageFeature = generated.LanguageFeature_FEATURE_INTERVAL_TYPE
	FeatureJsonArrayFunctions            LanguageFeature = generated.LanguageFeature_FEATURE_JSON_ARRAY_FUNCTIONS
	FeatureJsonStrictNumberParsing       LanguageFeature = generated.LanguageFeature_FEATURE_JSON_STRICT_NUMBER_PARSING
	FeatureJsonType                      LanguageFeature = generated.LanguageFeature_FEATURE_JSON_TYPE
	FeatureNamedArguments                LanguageFeature = generated.LanguageFeature_FEATURE_NAMED_ARGUMENTS
	FeatureNumericType                   LanguageFeature = generated.LanguageFeature_FEATURE_NUMERIC_TYPE
	FeatureParameterizedTypes            LanguageFeature = generated.LanguageFeature_FEATURE_PARAMETERIZED_TYPES
	FeatureTablesample                   LanguageFeature = generated.LanguageFeature_FEATURE_TABLESAMPLE
	FeatureTemplateFunctions             LanguageFeature = generated.LanguageFeature_FEATURE_TEMPLATE_FUNCTIONS
	FeatureTimestampNanos                LanguageFeature = generated.LanguageFeature_FEATURE_TIMESTAMP_NANOS

	FeatureV11HavingInAggregate               LanguageFeature = generated.LanguageFeature_FEATURE_V_1_1_HAVING_IN_AGGREGATE
	FeatureV11LimitInAggregate                LanguageFeature = generated.LanguageFeature_FEATURE_V_1_1_LIMIT_IN_AGGREGATE
	FeatureV11NullHandlingModifierInAggregate LanguageFeature = generated.LanguageFeature_FEATURE_V_1_1_NULL_HANDLING_MODIFIER_IN_AGGREGATE
	FeatureV11NullHandlingModifierInAnalytic  LanguageFeature = generated.LanguageFeature_FEATURE_V_1_1_NULL_HANDLING_MODIFIER_IN_ANALYTIC
	FeatureV11OrderByCollate                  LanguageFeature = generated.LanguageFeature_FEATURE_V_1_1_ORDER_BY_COLLATE
	FeatureV11OrderByInAggregate              LanguageFeature = generated.LanguageFeature_FEATURE_V_1_1_ORDER_BY_IN_AGGREGATE
	FeatureV11SelectStarExceptReplace         LanguageFeature = generated.LanguageFeature_FEATURE_V_1_1_SELECT_STAR_EXCEPT_REPLACE
	FeatureV11WithOnSubquery                  LanguageFeature = generated.LanguageFeature_FEATURE_V_1_1_WITH_ON_SUBQUERY
	FeatureV12CivilTime                       LanguageFeature = generated.LanguageFeature_FEATURE_V_1_2_CIVIL_TIME
	FeatureV12SafeFunctionCall                LanguageFeature = generated.LanguageFeature_FEATURE_V_1_2_SAFE_FUNCTION_CALL
	FeatureV12WeekWithWeekday                 LanguageFeature = generated.LanguageFeature_FEATURE_V_1_2_WEEK_WITH_WEEKDAY
	FeatureV13AllowDashesInTableName          LanguageFeature = generated.LanguageFeature_FEATURE_V_1_3_ALLOW_DASHES_IN_TABLE_NAME
	FeatureV13DateArithmetics                 LanguageFeature = generated.LanguageFeature_FEATURE_V_1_3_DATE_ARITHMETICS
	FeatureV13DateTimeConstructors            LanguageFeature = generated.LanguageFeature_FEATURE_V_1_3_DATE_TIME_CONSTRUCTORS
	FeatureV13DecimalAlias                    LanguageFeature = generated.LanguageFeature_FEATURE_V_1_3_DECIMAL_ALIAS
	FeatureV13ExtendedDateTimeSignatures      LanguageFeature = generated.LanguageFeature_FEATURE_V_1_3_EXTENDED_DATE_TIME_SIGNATURES
	FeatureV13ExtendedGeographyParsers        LanguageFeature = generated.LanguageFeature_FEATURE_V_1_3_EXTENDED_GEOGRAPHY_PARSERS
	FeatureV13FormatInCast                    LanguageFeature = generated.LanguageFeature_FEATURE_V_1_3_FORMAT_IN_CAST
	FeatureV13IsDistinct                      LanguageFeature = generated.LanguageFeature_FEATURE_V_1_3_IS_DISTINCT
	FeatureV13NullsFirstLastInOrderBy         LanguageFeature = generated.LanguageFeature_FEATURE_V_1_3_NULLS_FIRST_LAST_IN_ORDER_BY
	FeatureV13Pivot                           LanguageFeature = generated.LanguageFeature_FEATURE_V_1_3_PIVOT
	FeatureV13Qualify                         LanguageFeature = generated.LanguageFeature_FEATURE_V_1_3_QUALIFY
	FeatureV13Unpivot                         LanguageFeature = generated.LanguageFeature_FEATURE_V_1_3_UNPIVOT
)

// Parameter mode (analyzer option).
type ParameterMode = generated.ParameterMode

const (
	ParameterNamed      ParameterMode = generated.ParameterMode_PARAMETER_NAMED
	ParameterPositional ParameterMode = generated.ParameterMode_PARAMETER_POSITIONAL
	ParameterNone       ParameterMode = generated.ParameterMode_PARAMETER_NONE
)

// NameResolutionDefault is the default name-resolution mode.
const NameResolutionDefault = generated.NameResolutionMode_NAME_RESOLUTION_DEFAULT

// ParseLocationRecordFullNodeScope makes the analyzer attach a parse-location
// range to every resolved AST node.
const ParseLocationRecordFullNodeScope = generated.ParseLocationRecordType_PARSE_LOCATION_RECORD_FULL_NODE_SCOPE

// Product mode.
type ProductMode = generated.ProductMode

const (
	ProductInternal ProductMode = generated.ProductMode_PRODUCT_INTERNAL
	ProductExternal ProductMode = generated.ProductMode_PRODUCT_EXTERNAL
)

// Function-argument cardinality and signature kinds.
const (
	RequiredArgumentCardinality = generated.FunctionEnums_REQUIRED
	OptionalArgumentCardinality = generated.FunctionEnums_OPTIONAL
	RepeatedArgumentCardinality = generated.FunctionEnums_REPEATED

	ArgTypeFixed     = generated.SignatureArgumentKind_ARG_TYPE_FIXED
	ArgTypeAny1      = generated.SignatureArgumentKind_ARG_TYPE_ANY_1
	ArgTypeAny2      = generated.SignatureArgumentKind_ARG_TYPE_ANY_2
	ArgArrayTypeAny1 = generated.SignatureArgumentKind_ARG_ARRAY_TYPE_ANY_1
	ArgArrayTypeAny2 = generated.SignatureArgumentKind_ARG_ARRAY_TYPE_ANY_2
)

// NewFunctionArgumentTypeOptions mirrors the go-zetasql constructor
// signature, which takes only a cardinality.
func NewFunctionArgumentTypeOptions(c generated.FunctionEnums_ArgumentCardinality) *types.FunctionArgumentTypeOptions {
	return &types.FunctionArgumentTypeOptions{Cardinality: c}
}

// TypeFromKind returns the singleton zetasql-wasm Type for a scalar kind.
// Composite kinds (Array, Struct) return nil since they require element /
// field information that this kind alone cannot supply.
func TypeFromKind(k types.TypeKind) types.Type {
	switch k {
	case types.Int32:
		return types.Int32Type()
	case types.Int64:
		return types.Int64Type()
	case types.Uint32:
		return types.Uint32Type()
	case types.Uint64:
		return types.Uint64Type()
	case types.Bool:
		return types.BoolType()
	case types.Float:
		return types.FloatType()
	case types.Double:
		return types.DoubleType()
	case types.String:
		return types.StringType()
	case types.Bytes:
		return types.BytesType()
	case types.Date:
		return types.DateType()
	case types.Timestamp:
		return types.TimestampType()
	case types.Time:
		return types.TimeType()
	case types.Datetime:
		return types.DatetimeType()
	case types.Geography:
		return types.GeographyType()
	case types.Numeric:
		return types.NumericType()
	case types.BigNumeric:
		return types.BigNumericType()
	case types.Json:
		return types.JsonType()
	case types.Interval:
		return types.IntervalType()
	}
	return nil
}

// KindString returns the canonical proto name for a TypeKind (e.g. "TYPE_INT64").
// zetasql-wasm's TypeKind itself doesn't have a String method, but the
// underlying generated.TypeKind does.
func KindString(k types.TypeKind) string {
	return generated.TypeKind(k).String()
}

// TypeFromProto reconstructs a zetasql-wasm types.Type from its proto
// representation, walking ArrayType / StructType recursively. Returns nil
// for an unrecognised TypeKind. Used at the seam where a resolved AST
// node hands us a *generated.TypeProto (e.g. an argument or column type)
// but downstream code wants the typed Go interface.
func TypeFromProto(p *generated.TypeProto) (types.Type, error) {
	if p == nil {
		return nil, nil
	}
	k := types.TypeKind(p.GetTypeKind())
	switch k {
	case types.Array:
		elem, err := TypeFromProto(p.GetArrayType().GetElementType())
		if err != nil {
			return nil, err
		}
		return types.NewArrayType(elem)
	case types.Struct:
		fields := make([]*types.StructField, 0, len(p.GetStructType().GetField()))
		for _, f := range p.GetStructType().GetField() {
			ft, err := TypeFromProto(f.GetFieldType())
			if err != nil {
				return nil, err
			}
			fields = append(fields, &types.StructField{Name: f.GetFieldName(), Type: ft})
		}
		return types.NewStructType(fields)
	}
	return TypeFromKind(k), nil
}

// SupportedStatementKinds returns the resolved-node statement kinds that
// go-zetasqlite asks the analyzer to accept.
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

// === Resolved AST aliases ===

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
