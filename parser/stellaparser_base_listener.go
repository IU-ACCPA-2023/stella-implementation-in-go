// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // StellaParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseStellaParserListener is a complete listener for a parse tree produced by StellaParser.
type BaseStellaParserListener struct{}

var _ StellaParserListener = &BaseStellaParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStellaParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStellaParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStellaParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStellaParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart_Program is called when production start_Program is entered.
func (s *BaseStellaParserListener) EnterStart_Program(ctx *Start_ProgramContext) {}

// ExitStart_Program is called when production start_Program is exited.
func (s *BaseStellaParserListener) ExitStart_Program(ctx *Start_ProgramContext) {}

// EnterStart_Expr is called when production start_Expr is entered.
func (s *BaseStellaParserListener) EnterStart_Expr(ctx *Start_ExprContext) {}

// ExitStart_Expr is called when production start_Expr is exited.
func (s *BaseStellaParserListener) ExitStart_Expr(ctx *Start_ExprContext) {}

// EnterStart_Type is called when production start_Type is entered.
func (s *BaseStellaParserListener) EnterStart_Type(ctx *Start_TypeContext) {}

// ExitStart_Type is called when production start_Type is exited.
func (s *BaseStellaParserListener) ExitStart_Type(ctx *Start_TypeContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseStellaParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseStellaParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterLanguageCore is called when production LanguageCore is entered.
func (s *BaseStellaParserListener) EnterLanguageCore(ctx *LanguageCoreContext) {}

// ExitLanguageCore is called when production LanguageCore is exited.
func (s *BaseStellaParserListener) ExitLanguageCore(ctx *LanguageCoreContext) {}

// EnterAnExtension is called when production AnExtension is entered.
func (s *BaseStellaParserListener) EnterAnExtension(ctx *AnExtensionContext) {}

// ExitAnExtension is called when production AnExtension is exited.
func (s *BaseStellaParserListener) ExitAnExtension(ctx *AnExtensionContext) {}

// EnterDeclFun is called when production DeclFun is entered.
func (s *BaseStellaParserListener) EnterDeclFun(ctx *DeclFunContext) {}

// ExitDeclFun is called when production DeclFun is exited.
func (s *BaseStellaParserListener) ExitDeclFun(ctx *DeclFunContext) {}

// EnterDeclTypeAlias is called when production DeclTypeAlias is entered.
func (s *BaseStellaParserListener) EnterDeclTypeAlias(ctx *DeclTypeAliasContext) {}

// ExitDeclTypeAlias is called when production DeclTypeAlias is exited.
func (s *BaseStellaParserListener) ExitDeclTypeAlias(ctx *DeclTypeAliasContext) {}

// EnterInlineAnnotation is called when production InlineAnnotation is entered.
func (s *BaseStellaParserListener) EnterInlineAnnotation(ctx *InlineAnnotationContext) {}

// ExitInlineAnnotation is called when production InlineAnnotation is exited.
func (s *BaseStellaParserListener) ExitInlineAnnotation(ctx *InlineAnnotationContext) {}

// EnterParamDecl is called when production paramDecl is entered.
func (s *BaseStellaParserListener) EnterParamDecl(ctx *ParamDeclContext) {}

// ExitParamDecl is called when production paramDecl is exited.
func (s *BaseStellaParserListener) ExitParamDecl(ctx *ParamDeclContext) {}

// EnterIsEmpty is called when production IsEmpty is entered.
func (s *BaseStellaParserListener) EnterIsEmpty(ctx *IsEmptyContext) {}

// ExitIsEmpty is called when production IsEmpty is exited.
func (s *BaseStellaParserListener) ExitIsEmpty(ctx *IsEmptyContext) {}

// EnterFold is called when production Fold is entered.
func (s *BaseStellaParserListener) EnterFold(ctx *FoldContext) {}

// ExitFold is called when production Fold is exited.
func (s *BaseStellaParserListener) ExitFold(ctx *FoldContext) {}

// EnterAdd is called when production Add is entered.
func (s *BaseStellaParserListener) EnterAdd(ctx *AddContext) {}

// ExitAdd is called when production Add is exited.
func (s *BaseStellaParserListener) ExitAdd(ctx *AddContext) {}

// EnterIsZero is called when production IsZero is entered.
func (s *BaseStellaParserListener) EnterIsZero(ctx *IsZeroContext) {}

// ExitIsZero is called when production IsZero is exited.
func (s *BaseStellaParserListener) ExitIsZero(ctx *IsZeroContext) {}

// EnterLessThanOrEqual is called when production LessThanOrEqual is entered.
func (s *BaseStellaParserListener) EnterLessThanOrEqual(ctx *LessThanOrEqualContext) {}

// ExitLessThanOrEqual is called when production LessThanOrEqual is exited.
func (s *BaseStellaParserListener) ExitLessThanOrEqual(ctx *LessThanOrEqualContext) {}

// EnterSucc is called when production Succ is entered.
func (s *BaseStellaParserListener) EnterSucc(ctx *SuccContext) {}

// ExitSucc is called when production Succ is exited.
func (s *BaseStellaParserListener) ExitSucc(ctx *SuccContext) {}

// EnterVar is called when production Var is entered.
func (s *BaseStellaParserListener) EnterVar(ctx *VarContext) {}

// ExitVar is called when production Var is exited.
func (s *BaseStellaParserListener) ExitVar(ctx *VarContext) {}

// EnterInl is called when production Inl is entered.
func (s *BaseStellaParserListener) EnterInl(ctx *InlContext) {}

// ExitInl is called when production Inl is exited.
func (s *BaseStellaParserListener) ExitInl(ctx *InlContext) {}

// EnterGreaterThanOrEqual is called when production GreaterThanOrEqual is entered.
func (s *BaseStellaParserListener) EnterGreaterThanOrEqual(ctx *GreaterThanOrEqualContext) {}

// ExitGreaterThanOrEqual is called when production GreaterThanOrEqual is exited.
func (s *BaseStellaParserListener) ExitGreaterThanOrEqual(ctx *GreaterThanOrEqualContext) {}

// EnterInr is called when production Inr is entered.
func (s *BaseStellaParserListener) EnterInr(ctx *InrContext) {}

// ExitInr is called when production Inr is exited.
func (s *BaseStellaParserListener) ExitInr(ctx *InrContext) {}

// EnterDivide is called when production Divide is entered.
func (s *BaseStellaParserListener) EnterDivide(ctx *DivideContext) {}

// ExitDivide is called when production Divide is exited.
func (s *BaseStellaParserListener) ExitDivide(ctx *DivideContext) {}

// EnterLessThan is called when production LessThan is entered.
func (s *BaseStellaParserListener) EnterLessThan(ctx *LessThanContext) {}

// ExitLessThan is called when production LessThan is exited.
func (s *BaseStellaParserListener) ExitLessThan(ctx *LessThanContext) {}

// EnterLogicNot is called when production LogicNot is entered.
func (s *BaseStellaParserListener) EnterLogicNot(ctx *LogicNotContext) {}

// ExitLogicNot is called when production LogicNot is exited.
func (s *BaseStellaParserListener) ExitLogicNot(ctx *LogicNotContext) {}

// EnterDotRecord is called when production DotRecord is entered.
func (s *BaseStellaParserListener) EnterDotRecord(ctx *DotRecordContext) {}

// ExitDotRecord is called when production DotRecord is exited.
func (s *BaseStellaParserListener) ExitDotRecord(ctx *DotRecordContext) {}

// EnterParenthesisedExpr is called when production ParenthesisedExpr is entered.
func (s *BaseStellaParserListener) EnterParenthesisedExpr(ctx *ParenthesisedExprContext) {}

// ExitParenthesisedExpr is called when production ParenthesisedExpr is exited.
func (s *BaseStellaParserListener) ExitParenthesisedExpr(ctx *ParenthesisedExprContext) {}

// EnterGreaterThan is called when production GreaterThan is entered.
func (s *BaseStellaParserListener) EnterGreaterThan(ctx *GreaterThanContext) {}

// ExitGreaterThan is called when production GreaterThan is exited.
func (s *BaseStellaParserListener) ExitGreaterThan(ctx *GreaterThanContext) {}

// EnterEqual is called when production Equal is entered.
func (s *BaseStellaParserListener) EnterEqual(ctx *EqualContext) {}

// ExitEqual is called when production Equal is exited.
func (s *BaseStellaParserListener) ExitEqual(ctx *EqualContext) {}

// EnterTail is called when production Tail is entered.
func (s *BaseStellaParserListener) EnterTail(ctx *TailContext) {}

// ExitTail is called when production Tail is exited.
func (s *BaseStellaParserListener) ExitTail(ctx *TailContext) {}

// EnterMultiply is called when production Multiply is entered.
func (s *BaseStellaParserListener) EnterMultiply(ctx *MultiplyContext) {}

// ExitMultiply is called when production Multiply is exited.
func (s *BaseStellaParserListener) ExitMultiply(ctx *MultiplyContext) {}

// EnterRecord is called when production Record is entered.
func (s *BaseStellaParserListener) EnterRecord(ctx *RecordContext) {}

// ExitRecord is called when production Record is exited.
func (s *BaseStellaParserListener) ExitRecord(ctx *RecordContext) {}

// EnterList is called when production List is entered.
func (s *BaseStellaParserListener) EnterList(ctx *ListContext) {}

// ExitList is called when production List is exited.
func (s *BaseStellaParserListener) ExitList(ctx *ListContext) {}

// EnterLogicAnd is called when production LogicAnd is entered.
func (s *BaseStellaParserListener) EnterLogicAnd(ctx *LogicAndContext) {}

// ExitLogicAnd is called when production LogicAnd is exited.
func (s *BaseStellaParserListener) ExitLogicAnd(ctx *LogicAndContext) {}

// EnterLogicOr is called when production LogicOr is entered.
func (s *BaseStellaParserListener) EnterLogicOr(ctx *LogicOrContext) {}

// ExitLogicOr is called when production LogicOr is exited.
func (s *BaseStellaParserListener) ExitLogicOr(ctx *LogicOrContext) {}

// EnterHead is called when production Head is entered.
func (s *BaseStellaParserListener) EnterHead(ctx *HeadContext) {}

// ExitHead is called when production Head is exited.
func (s *BaseStellaParserListener) ExitHead(ctx *HeadContext) {}

// EnterTerminatingSemicolon is called when production TerminatingSemicolon is entered.
func (s *BaseStellaParserListener) EnterTerminatingSemicolon(ctx *TerminatingSemicolonContext) {}

// ExitTerminatingSemicolon is called when production TerminatingSemicolon is exited.
func (s *BaseStellaParserListener) ExitTerminatingSemicolon(ctx *TerminatingSemicolonContext) {}

// EnterNotEqual is called when production NotEqual is entered.
func (s *BaseStellaParserListener) EnterNotEqual(ctx *NotEqualContext) {}

// ExitNotEqual is called when production NotEqual is exited.
func (s *BaseStellaParserListener) ExitNotEqual(ctx *NotEqualContext) {}

// EnterConstUnit is called when production ConstUnit is entered.
func (s *BaseStellaParserListener) EnterConstUnit(ctx *ConstUnitContext) {}

// ExitConstUnit is called when production ConstUnit is exited.
func (s *BaseStellaParserListener) ExitConstUnit(ctx *ConstUnitContext) {}

// EnterPred is called when production Pred is entered.
func (s *BaseStellaParserListener) EnterPred(ctx *PredContext) {}

// ExitPred is called when production Pred is exited.
func (s *BaseStellaParserListener) ExitPred(ctx *PredContext) {}

// EnterMatch is called when production match is entered.
func (s *BaseStellaParserListener) EnterMatch(ctx *MatchContext) {}

// ExitMatch is called when production match is exited.
func (s *BaseStellaParserListener) ExitMatch(ctx *MatchContext) {}

// EnterTypeAsc is called when production TypeAsc is entered.
func (s *BaseStellaParserListener) EnterTypeAsc(ctx *TypeAscContext) {}

// ExitTypeAsc is called when production TypeAsc is exited.
func (s *BaseStellaParserListener) ExitTypeAsc(ctx *TypeAscContext) {}

// EnterNatRec is called when production NatRec is entered.
func (s *BaseStellaParserListener) EnterNatRec(ctx *NatRecContext) {}

// ExitNatRec is called when production NatRec is exited.
func (s *BaseStellaParserListener) ExitNatRec(ctx *NatRecContext) {}

// EnterSequence is called when production Sequence is entered.
func (s *BaseStellaParserListener) EnterSequence(ctx *SequenceContext) {}

// ExitSequence is called when production Sequence is exited.
func (s *BaseStellaParserListener) ExitSequence(ctx *SequenceContext) {}

// EnterConstFalse is called when production ConstFalse is entered.
func (s *BaseStellaParserListener) EnterConstFalse(ctx *ConstFalseContext) {}

// ExitConstFalse is called when production ConstFalse is exited.
func (s *BaseStellaParserListener) ExitConstFalse(ctx *ConstFalseContext) {}

// EnterAbstraction is called when production Abstraction is entered.
func (s *BaseStellaParserListener) EnterAbstraction(ctx *AbstractionContext) {}

// ExitAbstraction is called when production Abstraction is exited.
func (s *BaseStellaParserListener) ExitAbstraction(ctx *AbstractionContext) {}

// EnterConstInt is called when production ConstInt is entered.
func (s *BaseStellaParserListener) EnterConstInt(ctx *ConstIntContext) {}

// ExitConstInt is called when production ConstInt is exited.
func (s *BaseStellaParserListener) ExitConstInt(ctx *ConstIntContext) {}

// EnterUnfold is called when production Unfold is entered.
func (s *BaseStellaParserListener) EnterUnfold(ctx *UnfoldContext) {}

// ExitUnfold is called when production Unfold is exited.
func (s *BaseStellaParserListener) ExitUnfold(ctx *UnfoldContext) {}

// EnterVariant is called when production Variant is entered.
func (s *BaseStellaParserListener) EnterVariant(ctx *VariantContext) {}

// ExitVariant is called when production Variant is exited.
func (s *BaseStellaParserListener) ExitVariant(ctx *VariantContext) {}

// EnterConstTrue is called when production ConstTrue is entered.
func (s *BaseStellaParserListener) EnterConstTrue(ctx *ConstTrueContext) {}

// ExitConstTrue is called when production ConstTrue is exited.
func (s *BaseStellaParserListener) ExitConstTrue(ctx *ConstTrueContext) {}

// EnterDotTuple is called when production DotTuple is entered.
func (s *BaseStellaParserListener) EnterDotTuple(ctx *DotTupleContext) {}

// ExitDotTuple is called when production DotTuple is exited.
func (s *BaseStellaParserListener) ExitDotTuple(ctx *DotTupleContext) {}

// EnterFix is called when production Fix is entered.
func (s *BaseStellaParserListener) EnterFix(ctx *FixContext) {}

// ExitFix is called when production Fix is exited.
func (s *BaseStellaParserListener) ExitFix(ctx *FixContext) {}

// EnterSubtract is called when production Subtract is entered.
func (s *BaseStellaParserListener) EnterSubtract(ctx *SubtractContext) {}

// ExitSubtract is called when production Subtract is exited.
func (s *BaseStellaParserListener) ExitSubtract(ctx *SubtractContext) {}

// EnterLet is called when production Let is entered.
func (s *BaseStellaParserListener) EnterLet(ctx *LetContext) {}

// ExitLet is called when production Let is exited.
func (s *BaseStellaParserListener) ExitLet(ctx *LetContext) {}

// EnterIf is called when production If is entered.
func (s *BaseStellaParserListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production If is exited.
func (s *BaseStellaParserListener) ExitIf(ctx *IfContext) {}

// EnterApplication is called when production Application is entered.
func (s *BaseStellaParserListener) EnterApplication(ctx *ApplicationContext) {}

// ExitApplication is called when production Application is exited.
func (s *BaseStellaParserListener) ExitApplication(ctx *ApplicationContext) {}

// EnterTuple is called when production Tuple is entered.
func (s *BaseStellaParserListener) EnterTuple(ctx *TupleContext) {}

// ExitTuple is called when production Tuple is exited.
func (s *BaseStellaParserListener) ExitTuple(ctx *TupleContext) {}

// EnterConsList is called when production ConsList is entered.
func (s *BaseStellaParserListener) EnterConsList(ctx *ConsListContext) {}

// ExitConsList is called when production ConsList is exited.
func (s *BaseStellaParserListener) ExitConsList(ctx *ConsListContext) {}

// EnterPatternBinding is called when production patternBinding is entered.
func (s *BaseStellaParserListener) EnterPatternBinding(ctx *PatternBindingContext) {}

// ExitPatternBinding is called when production patternBinding is exited.
func (s *BaseStellaParserListener) ExitPatternBinding(ctx *PatternBindingContext) {}

// EnterBinding is called when production binding is entered.
func (s *BaseStellaParserListener) EnterBinding(ctx *BindingContext) {}

// ExitBinding is called when production binding is exited.
func (s *BaseStellaParserListener) ExitBinding(ctx *BindingContext) {}

// EnterMatch_case is called when production match_case is entered.
func (s *BaseStellaParserListener) EnterMatch_case(ctx *Match_caseContext) {}

// ExitMatch_case is called when production match_case is exited.
func (s *BaseStellaParserListener) ExitMatch_case(ctx *Match_caseContext) {}

// EnterPatternVariant is called when production PatternVariant is entered.
func (s *BaseStellaParserListener) EnterPatternVariant(ctx *PatternVariantContext) {}

// ExitPatternVariant is called when production PatternVariant is exited.
func (s *BaseStellaParserListener) ExitPatternVariant(ctx *PatternVariantContext) {}

// EnterPatternInl is called when production PatternInl is entered.
func (s *BaseStellaParserListener) EnterPatternInl(ctx *PatternInlContext) {}

// ExitPatternInl is called when production PatternInl is exited.
func (s *BaseStellaParserListener) ExitPatternInl(ctx *PatternInlContext) {}

// EnterPatternInr is called when production PatternInr is entered.
func (s *BaseStellaParserListener) EnterPatternInr(ctx *PatternInrContext) {}

// ExitPatternInr is called when production PatternInr is exited.
func (s *BaseStellaParserListener) ExitPatternInr(ctx *PatternInrContext) {}

// EnterPatternTuple is called when production PatternTuple is entered.
func (s *BaseStellaParserListener) EnterPatternTuple(ctx *PatternTupleContext) {}

// ExitPatternTuple is called when production PatternTuple is exited.
func (s *BaseStellaParserListener) ExitPatternTuple(ctx *PatternTupleContext) {}

// EnterPatternRecord is called when production PatternRecord is entered.
func (s *BaseStellaParserListener) EnterPatternRecord(ctx *PatternRecordContext) {}

// ExitPatternRecord is called when production PatternRecord is exited.
func (s *BaseStellaParserListener) ExitPatternRecord(ctx *PatternRecordContext) {}

// EnterPatternList is called when production PatternList is entered.
func (s *BaseStellaParserListener) EnterPatternList(ctx *PatternListContext) {}

// ExitPatternList is called when production PatternList is exited.
func (s *BaseStellaParserListener) ExitPatternList(ctx *PatternListContext) {}

// EnterPatternCons is called when production PatternCons is entered.
func (s *BaseStellaParserListener) EnterPatternCons(ctx *PatternConsContext) {}

// ExitPatternCons is called when production PatternCons is exited.
func (s *BaseStellaParserListener) ExitPatternCons(ctx *PatternConsContext) {}

// EnterPatternFalse is called when production PatternFalse is entered.
func (s *BaseStellaParserListener) EnterPatternFalse(ctx *PatternFalseContext) {}

// ExitPatternFalse is called when production PatternFalse is exited.
func (s *BaseStellaParserListener) ExitPatternFalse(ctx *PatternFalseContext) {}

// EnterPatternTrue is called when production PatternTrue is entered.
func (s *BaseStellaParserListener) EnterPatternTrue(ctx *PatternTrueContext) {}

// ExitPatternTrue is called when production PatternTrue is exited.
func (s *BaseStellaParserListener) ExitPatternTrue(ctx *PatternTrueContext) {}

// EnterPatternUnit is called when production PatternUnit is entered.
func (s *BaseStellaParserListener) EnterPatternUnit(ctx *PatternUnitContext) {}

// ExitPatternUnit is called when production PatternUnit is exited.
func (s *BaseStellaParserListener) ExitPatternUnit(ctx *PatternUnitContext) {}

// EnterPatternInt is called when production PatternInt is entered.
func (s *BaseStellaParserListener) EnterPatternInt(ctx *PatternIntContext) {}

// ExitPatternInt is called when production PatternInt is exited.
func (s *BaseStellaParserListener) ExitPatternInt(ctx *PatternIntContext) {}

// EnterPatternSucc is called when production PatternSucc is entered.
func (s *BaseStellaParserListener) EnterPatternSucc(ctx *PatternSuccContext) {}

// ExitPatternSucc is called when production PatternSucc is exited.
func (s *BaseStellaParserListener) ExitPatternSucc(ctx *PatternSuccContext) {}

// EnterPatternVar is called when production PatternVar is entered.
func (s *BaseStellaParserListener) EnterPatternVar(ctx *PatternVarContext) {}

// ExitPatternVar is called when production PatternVar is exited.
func (s *BaseStellaParserListener) ExitPatternVar(ctx *PatternVarContext) {}

// EnterParenthesisedPattern is called when production ParenthesisedPattern is entered.
func (s *BaseStellaParserListener) EnterParenthesisedPattern(ctx *ParenthesisedPatternContext) {}

// ExitParenthesisedPattern is called when production ParenthesisedPattern is exited.
func (s *BaseStellaParserListener) ExitParenthesisedPattern(ctx *ParenthesisedPatternContext) {}

// EnterLabelledPattern is called when production labelledPattern is entered.
func (s *BaseStellaParserListener) EnterLabelledPattern(ctx *LabelledPatternContext) {}

// ExitLabelledPattern is called when production labelledPattern is exited.
func (s *BaseStellaParserListener) ExitLabelledPattern(ctx *LabelledPatternContext) {}

// EnterTypeTuple is called when production TypeTuple is entered.
func (s *BaseStellaParserListener) EnterTypeTuple(ctx *TypeTupleContext) {}

// ExitTypeTuple is called when production TypeTuple is exited.
func (s *BaseStellaParserListener) ExitTypeTuple(ctx *TypeTupleContext) {}

// EnterTypeVar is called when production TypeVar is entered.
func (s *BaseStellaParserListener) EnterTypeVar(ctx *TypeVarContext) {}

// ExitTypeVar is called when production TypeVar is exited.
func (s *BaseStellaParserListener) ExitTypeVar(ctx *TypeVarContext) {}

// EnterTypeVariant is called when production TypeVariant is entered.
func (s *BaseStellaParserListener) EnterTypeVariant(ctx *TypeVariantContext) {}

// ExitTypeVariant is called when production TypeVariant is exited.
func (s *BaseStellaParserListener) ExitTypeVariant(ctx *TypeVariantContext) {}

// EnterTypeUnit is called when production TypeUnit is entered.
func (s *BaseStellaParserListener) EnterTypeUnit(ctx *TypeUnitContext) {}

// ExitTypeUnit is called when production TypeUnit is exited.
func (s *BaseStellaParserListener) ExitTypeUnit(ctx *TypeUnitContext) {}

// EnterTypeBool is called when production TypeBool is entered.
func (s *BaseStellaParserListener) EnterTypeBool(ctx *TypeBoolContext) {}

// ExitTypeBool is called when production TypeBool is exited.
func (s *BaseStellaParserListener) ExitTypeBool(ctx *TypeBoolContext) {}

// EnterTypeNat is called when production TypeNat is entered.
func (s *BaseStellaParserListener) EnterTypeNat(ctx *TypeNatContext) {}

// ExitTypeNat is called when production TypeNat is exited.
func (s *BaseStellaParserListener) ExitTypeNat(ctx *TypeNatContext) {}

// EnterTypeRec is called when production TypeRec is entered.
func (s *BaseStellaParserListener) EnterTypeRec(ctx *TypeRecContext) {}

// ExitTypeRec is called when production TypeRec is exited.
func (s *BaseStellaParserListener) ExitTypeRec(ctx *TypeRecContext) {}

// EnterTypeParens is called when production TypeParens is entered.
func (s *BaseStellaParserListener) EnterTypeParens(ctx *TypeParensContext) {}

// ExitTypeParens is called when production TypeParens is exited.
func (s *BaseStellaParserListener) ExitTypeParens(ctx *TypeParensContext) {}

// EnterTypeFun is called when production TypeFun is entered.
func (s *BaseStellaParserListener) EnterTypeFun(ctx *TypeFunContext) {}

// ExitTypeFun is called when production TypeFun is exited.
func (s *BaseStellaParserListener) ExitTypeFun(ctx *TypeFunContext) {}

// EnterTypeRecord is called when production TypeRecord is entered.
func (s *BaseStellaParserListener) EnterTypeRecord(ctx *TypeRecordContext) {}

// ExitTypeRecord is called when production TypeRecord is exited.
func (s *BaseStellaParserListener) ExitTypeRecord(ctx *TypeRecordContext) {}

// EnterTypeList is called when production TypeList is entered.
func (s *BaseStellaParserListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production TypeList is exited.
func (s *BaseStellaParserListener) ExitTypeList(ctx *TypeListContext) {}

// EnterTypeSum is called when production TypeSum is entered.
func (s *BaseStellaParserListener) EnterTypeSum(ctx *TypeSumContext) {}

// ExitTypeSum is called when production TypeSum is exited.
func (s *BaseStellaParserListener) ExitTypeSum(ctx *TypeSumContext) {}

// EnterRecordFieldType is called when production recordFieldType is entered.
func (s *BaseStellaParserListener) EnterRecordFieldType(ctx *RecordFieldTypeContext) {}

// ExitRecordFieldType is called when production recordFieldType is exited.
func (s *BaseStellaParserListener) ExitRecordFieldType(ctx *RecordFieldTypeContext) {}

// EnterVariantFieldType is called when production variantFieldType is entered.
func (s *BaseStellaParserListener) EnterVariantFieldType(ctx *VariantFieldTypeContext) {}

// ExitVariantFieldType is called when production variantFieldType is exited.
func (s *BaseStellaParserListener) ExitVariantFieldType(ctx *VariantFieldTypeContext) {}
