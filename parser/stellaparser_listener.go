// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // StellaParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// StellaParserListener is a complete listener for a parse tree produced by StellaParser.
type StellaParserListener interface {
	antlr.ParseTreeListener

	// EnterStart_Program is called when entering the start_Program production.
	EnterStart_Program(c *Start_ProgramContext)

	// EnterStart_Expr is called when entering the start_Expr production.
	EnterStart_Expr(c *Start_ExprContext)

	// EnterStart_Type is called when entering the start_Type production.
	EnterStart_Type(c *Start_TypeContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterLanguageCore is called when entering the LanguageCore production.
	EnterLanguageCore(c *LanguageCoreContext)

	// EnterAnExtension is called when entering the AnExtension production.
	EnterAnExtension(c *AnExtensionContext)

	// EnterDeclFun is called when entering the DeclFun production.
	EnterDeclFun(c *DeclFunContext)

	// EnterDeclTypeAlias is called when entering the DeclTypeAlias production.
	EnterDeclTypeAlias(c *DeclTypeAliasContext)

	// EnterInlineAnnotation is called when entering the InlineAnnotation production.
	EnterInlineAnnotation(c *InlineAnnotationContext)

	// EnterParamDecl is called when entering the paramDecl production.
	EnterParamDecl(c *ParamDeclContext)

	// EnterIsEmpty is called when entering the IsEmpty production.
	EnterIsEmpty(c *IsEmptyContext)

	// EnterFold is called when entering the Fold production.
	EnterFold(c *FoldContext)

	// EnterAdd is called when entering the Add production.
	EnterAdd(c *AddContext)

	// EnterIsZero is called when entering the IsZero production.
	EnterIsZero(c *IsZeroContext)

	// EnterLessThanOrEqual is called when entering the LessThanOrEqual production.
	EnterLessThanOrEqual(c *LessThanOrEqualContext)

	// EnterSucc is called when entering the Succ production.
	EnterSucc(c *SuccContext)

	// EnterVar is called when entering the Var production.
	EnterVar(c *VarContext)

	// EnterGreaterThanOrEqual is called when entering the GreaterThanOrEqual production.
	EnterGreaterThanOrEqual(c *GreaterThanOrEqualContext)

	// EnterLessThan is called when entering the LessThan production.
	EnterLessThan(c *LessThanContext)

	// EnterLogicNot is called when entering the LogicNot production.
	EnterLogicNot(c *LogicNotContext)

	// EnterDotRecord is called when entering the DotRecord production.
	EnterDotRecord(c *DotRecordContext)

	// EnterParenthesisedExpr is called when entering the ParenthesisedExpr production.
	EnterParenthesisedExpr(c *ParenthesisedExprContext)

	// EnterGreaterThan is called when entering the GreaterThan production.
	EnterGreaterThan(c *GreaterThanContext)

	// EnterEqual is called when entering the Equal production.
	EnterEqual(c *EqualContext)

	// EnterTail is called when entering the Tail production.
	EnterTail(c *TailContext)

	// EnterMultiply is called when entering the Multiply production.
	EnterMultiply(c *MultiplyContext)

	// EnterRecord is called when entering the Record production.
	EnterRecord(c *RecordContext)

	// EnterList is called when entering the List production.
	EnterList(c *ListContext)

	// EnterLogicAnd is called when entering the LogicAnd production.
	EnterLogicAnd(c *LogicAndContext)

	// EnterLogicOr is called when entering the LogicOr production.
	EnterLogicOr(c *LogicOrContext)

	// EnterHead is called when entering the Head production.
	EnterHead(c *HeadContext)

	// EnterNotEqual is called when entering the NotEqual production.
	EnterNotEqual(c *NotEqualContext)

	// EnterPred is called when entering the Pred production.
	EnterPred(c *PredContext)

	// EnterMatch is called when entering the match production.
	EnterMatch(c *MatchContext)

	// EnterTypeAsc is called when entering the TypeAsc production.
	EnterTypeAsc(c *TypeAscContext)

	// EnterNatRec is called when entering the NatRec production.
	EnterNatRec(c *NatRecContext)

	// EnterConstFalse is called when entering the ConstFalse production.
	EnterConstFalse(c *ConstFalseContext)

	// EnterAbstraction is called when entering the Abstraction production.
	EnterAbstraction(c *AbstractionContext)

	// EnterConstInt is called when entering the ConstInt production.
	EnterConstInt(c *ConstIntContext)

	// EnterUnfold is called when entering the Unfold production.
	EnterUnfold(c *UnfoldContext)

	// EnterVariant is called when entering the Variant production.
	EnterVariant(c *VariantContext)

	// EnterConstTrue is called when entering the ConstTrue production.
	EnterConstTrue(c *ConstTrueContext)

	// EnterDotTuple is called when entering the DotTuple production.
	EnterDotTuple(c *DotTupleContext)

	// EnterFix is called when entering the Fix production.
	EnterFix(c *FixContext)

	// EnterLet is called when entering the Let production.
	EnterLet(c *LetContext)

	// EnterIf is called when entering the If production.
	EnterIf(c *IfContext)

	// EnterApplication is called when entering the Application production.
	EnterApplication(c *ApplicationContext)

	// EnterTuple is called when entering the Tuple production.
	EnterTuple(c *TupleContext)

	// EnterConsList is called when entering the ConsList production.
	EnterConsList(c *ConsListContext)

	// EnterBinding is called when entering the binding production.
	EnterBinding(c *BindingContext)

	// EnterMatch_case is called when entering the match_case production.
	EnterMatch_case(c *Match_caseContext)

	// EnterPatternVariant is called when entering the PatternVariant production.
	EnterPatternVariant(c *PatternVariantContext)

	// EnterPatternTuple is called when entering the PatternTuple production.
	EnterPatternTuple(c *PatternTupleContext)

	// EnterPatternRecord is called when entering the PatternRecord production.
	EnterPatternRecord(c *PatternRecordContext)

	// EnterPatternList is called when entering the PatternList production.
	EnterPatternList(c *PatternListContext)

	// EnterPatternCons is called when entering the PatternCons production.
	EnterPatternCons(c *PatternConsContext)

	// EnterPatternFalse is called when entering the PatternFalse production.
	EnterPatternFalse(c *PatternFalseContext)

	// EnterPatternTrue is called when entering the PatternTrue production.
	EnterPatternTrue(c *PatternTrueContext)

	// EnterPatternInt is called when entering the PatternInt production.
	EnterPatternInt(c *PatternIntContext)

	// EnterPatternSucc is called when entering the PatternSucc production.
	EnterPatternSucc(c *PatternSuccContext)

	// EnterPatternVar is called when entering the PatternVar production.
	EnterPatternVar(c *PatternVarContext)

	// EnterParenthesisedPattern is called when entering the ParenthesisedPattern production.
	EnterParenthesisedPattern(c *ParenthesisedPatternContext)

	// EnterLabelledPattern is called when entering the labelledPattern production.
	EnterLabelledPattern(c *LabelledPatternContext)

	// EnterTypeTuple is called when entering the TypeTuple production.
	EnterTypeTuple(c *TypeTupleContext)

	// EnterTypeVar is called when entering the TypeVar production.
	EnterTypeVar(c *TypeVarContext)

	// EnterTypeVariant is called when entering the TypeVariant production.
	EnterTypeVariant(c *TypeVariantContext)

	// EnterTypeUnit is called when entering the TypeUnit production.
	EnterTypeUnit(c *TypeUnitContext)

	// EnterTypeBool is called when entering the TypeBool production.
	EnterTypeBool(c *TypeBoolContext)

	// EnterTypeNat is called when entering the TypeNat production.
	EnterTypeNat(c *TypeNatContext)

	// EnterTypeRec is called when entering the TypeRec production.
	EnterTypeRec(c *TypeRecContext)

	// EnterTypeParens is called when entering the TypeParens production.
	EnterTypeParens(c *TypeParensContext)

	// EnterTypeFun is called when entering the TypeFun production.
	EnterTypeFun(c *TypeFunContext)

	// EnterTypeRecord is called when entering the TypeRecord production.
	EnterTypeRecord(c *TypeRecordContext)

	// EnterTypeList is called when entering the TypeList production.
	EnterTypeList(c *TypeListContext)

	// EnterTypeSum is called when entering the TypeSum production.
	EnterTypeSum(c *TypeSumContext)

	// EnterRecordFieldType is called when entering the recordFieldType production.
	EnterRecordFieldType(c *RecordFieldTypeContext)

	// EnterVariantFieldType is called when entering the variantFieldType production.
	EnterVariantFieldType(c *VariantFieldTypeContext)

	// ExitStart_Program is called when exiting the start_Program production.
	ExitStart_Program(c *Start_ProgramContext)

	// ExitStart_Expr is called when exiting the start_Expr production.
	ExitStart_Expr(c *Start_ExprContext)

	// ExitStart_Type is called when exiting the start_Type production.
	ExitStart_Type(c *Start_TypeContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitLanguageCore is called when exiting the LanguageCore production.
	ExitLanguageCore(c *LanguageCoreContext)

	// ExitAnExtension is called when exiting the AnExtension production.
	ExitAnExtension(c *AnExtensionContext)

	// ExitDeclFun is called when exiting the DeclFun production.
	ExitDeclFun(c *DeclFunContext)

	// ExitDeclTypeAlias is called when exiting the DeclTypeAlias production.
	ExitDeclTypeAlias(c *DeclTypeAliasContext)

	// ExitInlineAnnotation is called when exiting the InlineAnnotation production.
	ExitInlineAnnotation(c *InlineAnnotationContext)

	// ExitParamDecl is called when exiting the paramDecl production.
	ExitParamDecl(c *ParamDeclContext)

	// ExitIsEmpty is called when exiting the IsEmpty production.
	ExitIsEmpty(c *IsEmptyContext)

	// ExitFold is called when exiting the Fold production.
	ExitFold(c *FoldContext)

	// ExitAdd is called when exiting the Add production.
	ExitAdd(c *AddContext)

	// ExitIsZero is called when exiting the IsZero production.
	ExitIsZero(c *IsZeroContext)

	// ExitLessThanOrEqual is called when exiting the LessThanOrEqual production.
	ExitLessThanOrEqual(c *LessThanOrEqualContext)

	// ExitSucc is called when exiting the Succ production.
	ExitSucc(c *SuccContext)

	// ExitVar is called when exiting the Var production.
	ExitVar(c *VarContext)

	// ExitGreaterThanOrEqual is called when exiting the GreaterThanOrEqual production.
	ExitGreaterThanOrEqual(c *GreaterThanOrEqualContext)

	// ExitLessThan is called when exiting the LessThan production.
	ExitLessThan(c *LessThanContext)

	// ExitLogicNot is called when exiting the LogicNot production.
	ExitLogicNot(c *LogicNotContext)

	// ExitDotRecord is called when exiting the DotRecord production.
	ExitDotRecord(c *DotRecordContext)

	// ExitParenthesisedExpr is called when exiting the ParenthesisedExpr production.
	ExitParenthesisedExpr(c *ParenthesisedExprContext)

	// ExitGreaterThan is called when exiting the GreaterThan production.
	ExitGreaterThan(c *GreaterThanContext)

	// ExitEqual is called when exiting the Equal production.
	ExitEqual(c *EqualContext)

	// ExitTail is called when exiting the Tail production.
	ExitTail(c *TailContext)

	// ExitMultiply is called when exiting the Multiply production.
	ExitMultiply(c *MultiplyContext)

	// ExitRecord is called when exiting the Record production.
	ExitRecord(c *RecordContext)

	// ExitList is called when exiting the List production.
	ExitList(c *ListContext)

	// ExitLogicAnd is called when exiting the LogicAnd production.
	ExitLogicAnd(c *LogicAndContext)

	// ExitLogicOr is called when exiting the LogicOr production.
	ExitLogicOr(c *LogicOrContext)

	// ExitHead is called when exiting the Head production.
	ExitHead(c *HeadContext)

	// ExitNotEqual is called when exiting the NotEqual production.
	ExitNotEqual(c *NotEqualContext)

	// ExitPred is called when exiting the Pred production.
	ExitPred(c *PredContext)

	// ExitMatch is called when exiting the match production.
	ExitMatch(c *MatchContext)

	// ExitTypeAsc is called when exiting the TypeAsc production.
	ExitTypeAsc(c *TypeAscContext)

	// ExitNatRec is called when exiting the NatRec production.
	ExitNatRec(c *NatRecContext)

	// ExitConstFalse is called when exiting the ConstFalse production.
	ExitConstFalse(c *ConstFalseContext)

	// ExitAbstraction is called when exiting the Abstraction production.
	ExitAbstraction(c *AbstractionContext)

	// ExitConstInt is called when exiting the ConstInt production.
	ExitConstInt(c *ConstIntContext)

	// ExitUnfold is called when exiting the Unfold production.
	ExitUnfold(c *UnfoldContext)

	// ExitVariant is called when exiting the Variant production.
	ExitVariant(c *VariantContext)

	// ExitConstTrue is called when exiting the ConstTrue production.
	ExitConstTrue(c *ConstTrueContext)

	// ExitDotTuple is called when exiting the DotTuple production.
	ExitDotTuple(c *DotTupleContext)

	// ExitFix is called when exiting the Fix production.
	ExitFix(c *FixContext)

	// ExitLet is called when exiting the Let production.
	ExitLet(c *LetContext)

	// ExitIf is called when exiting the If production.
	ExitIf(c *IfContext)

	// ExitApplication is called when exiting the Application production.
	ExitApplication(c *ApplicationContext)

	// ExitTuple is called when exiting the Tuple production.
	ExitTuple(c *TupleContext)

	// ExitConsList is called when exiting the ConsList production.
	ExitConsList(c *ConsListContext)

	// ExitBinding is called when exiting the binding production.
	ExitBinding(c *BindingContext)

	// ExitMatch_case is called when exiting the match_case production.
	ExitMatch_case(c *Match_caseContext)

	// ExitPatternVariant is called when exiting the PatternVariant production.
	ExitPatternVariant(c *PatternVariantContext)

	// ExitPatternTuple is called when exiting the PatternTuple production.
	ExitPatternTuple(c *PatternTupleContext)

	// ExitPatternRecord is called when exiting the PatternRecord production.
	ExitPatternRecord(c *PatternRecordContext)

	// ExitPatternList is called when exiting the PatternList production.
	ExitPatternList(c *PatternListContext)

	// ExitPatternCons is called when exiting the PatternCons production.
	ExitPatternCons(c *PatternConsContext)

	// ExitPatternFalse is called when exiting the PatternFalse production.
	ExitPatternFalse(c *PatternFalseContext)

	// ExitPatternTrue is called when exiting the PatternTrue production.
	ExitPatternTrue(c *PatternTrueContext)

	// ExitPatternInt is called when exiting the PatternInt production.
	ExitPatternInt(c *PatternIntContext)

	// ExitPatternSucc is called when exiting the PatternSucc production.
	ExitPatternSucc(c *PatternSuccContext)

	// ExitPatternVar is called when exiting the PatternVar production.
	ExitPatternVar(c *PatternVarContext)

	// ExitParenthesisedPattern is called when exiting the ParenthesisedPattern production.
	ExitParenthesisedPattern(c *ParenthesisedPatternContext)

	// ExitLabelledPattern is called when exiting the labelledPattern production.
	ExitLabelledPattern(c *LabelledPatternContext)

	// ExitTypeTuple is called when exiting the TypeTuple production.
	ExitTypeTuple(c *TypeTupleContext)

	// ExitTypeVar is called when exiting the TypeVar production.
	ExitTypeVar(c *TypeVarContext)

	// ExitTypeVariant is called when exiting the TypeVariant production.
	ExitTypeVariant(c *TypeVariantContext)

	// ExitTypeUnit is called when exiting the TypeUnit production.
	ExitTypeUnit(c *TypeUnitContext)

	// ExitTypeBool is called when exiting the TypeBool production.
	ExitTypeBool(c *TypeBoolContext)

	// ExitTypeNat is called when exiting the TypeNat production.
	ExitTypeNat(c *TypeNatContext)

	// ExitTypeRec is called when exiting the TypeRec production.
	ExitTypeRec(c *TypeRecContext)

	// ExitTypeParens is called when exiting the TypeParens production.
	ExitTypeParens(c *TypeParensContext)

	// ExitTypeFun is called when exiting the TypeFun production.
	ExitTypeFun(c *TypeFunContext)

	// ExitTypeRecord is called when exiting the TypeRecord production.
	ExitTypeRecord(c *TypeRecordContext)

	// ExitTypeList is called when exiting the TypeList production.
	ExitTypeList(c *TypeListContext)

	// ExitTypeSum is called when exiting the TypeSum production.
	ExitTypeSum(c *TypeSumContext)

	// ExitRecordFieldType is called when exiting the recordFieldType production.
	ExitRecordFieldType(c *RecordFieldTypeContext)

	// ExitVariantFieldType is called when exiting the variantFieldType production.
	ExitVariantFieldType(c *VariantFieldTypeContext)
}
