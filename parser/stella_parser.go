// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // StellaParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type StellaParser struct {
	*antlr.BaseParser
}

var stellaparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func stellaparserParserInit() {
	staticData := &stellaparserParserStaticData
	staticData.literalNames = []string{
		"", "','", "';'", "'('", "')'", "'{'", "'}'", "'='", "':'", "'->'",
		"'=>'", "'|'", "'<|'", "'|>'", "'['", "']'", "'<'", "'<='", "'>'", "'>='",
		"'=='", "'!='", "'+'", "'-'", "'*'", "'/'", "'.'", "'List::head'", "'List::isempty'",
		"'List::tail'", "'Nat::pred'", "'Nat::iszero'", "'Nat::rec'", "'Bool'",
		"'Nat'", "'Unit'", "'and'", "'as'", "'cons'", "'core'", "'else'", "'extend'",
		"'false'", "'fix'", "'fn'", "'fold'", "'if'", "'in'", "'inl'", "'inline'",
		"'inr'", "'language'", "'let'", "'letrec'", "'match'", "'not'", "'or'",
		"'return'", "'succ'", "'then'", "'throws'", "'true'", "'type'", "'unfold'",
		"'unit'", "'with'", "'\\u00B5'",
	}
	staticData.symbolicNames = []string{
		"", "Surrogate_id_SYMB_0", "Surrogate_id_SYMB_1", "Surrogate_id_SYMB_2",
		"Surrogate_id_SYMB_3", "Surrogate_id_SYMB_4", "Surrogate_id_SYMB_5",
		"Surrogate_id_SYMB_6", "Surrogate_id_SYMB_7", "Surrogate_id_SYMB_8",
		"Surrogate_id_SYMB_9", "Surrogate_id_SYMB_10", "Surrogate_id_SYMB_11",
		"Surrogate_id_SYMB_12", "Surrogate_id_SYMB_13", "Surrogate_id_SYMB_14",
		"Surrogate_id_SYMB_15", "Surrogate_id_SYMB_16", "Surrogate_id_SYMB_17",
		"Surrogate_id_SYMB_18", "Surrogate_id_SYMB_19", "Surrogate_id_SYMB_20",
		"Surrogate_id_SYMB_21", "Surrogate_id_SYMB_22", "Surrogate_id_SYMB_23",
		"Surrogate_id_SYMB_24", "Surrogate_id_SYMB_25", "Surrogate_id_SYMB_26",
		"Surrogate_id_SYMB_27", "Surrogate_id_SYMB_28", "Surrogate_id_SYMB_29",
		"Surrogate_id_SYMB_30", "Surrogate_id_SYMB_31", "Surrogate_id_SYMB_32",
		"Surrogate_id_SYMB_33", "Surrogate_id_SYMB_34", "Surrogate_id_SYMB_35",
		"Surrogate_id_SYMB_36", "Surrogate_id_SYMB_37", "Surrogate_id_SYMB_38",
		"Surrogate_id_SYMB_39", "Surrogate_id_SYMB_40", "Surrogate_id_SYMB_41",
		"Surrogate_id_SYMB_42", "Surrogate_id_SYMB_43", "Surrogate_id_SYMB_44",
		"Surrogate_id_SYMB_45", "Surrogate_id_SYMB_46", "Surrogate_id_SYMB_47",
		"Surrogate_id_SYMB_48", "Surrogate_id_SYMB_49", "Surrogate_id_SYMB_50",
		"Surrogate_id_SYMB_51", "Surrogate_id_SYMB_52", "Surrogate_id_SYMB_53",
		"Surrogate_id_SYMB_54", "Surrogate_id_SYMB_55", "Surrogate_id_SYMB_56",
		"Surrogate_id_SYMB_57", "Surrogate_id_SYMB_58", "Surrogate_id_SYMB_59",
		"Surrogate_id_SYMB_60", "Surrogate_id_SYMB_61", "Surrogate_id_SYMB_62",
		"Surrogate_id_SYMB_63", "Surrogate_id_SYMB_64", "Surrogate_id_SYMB_65",
		"COMMENT_antlr_builtin", "MULTICOMMENT_antlr_builtin", "StellaIdent",
		"ExtensionName", "INTEGER", "WS", "ErrorToken",
	}
	staticData.ruleNames = []string{
		"start_Program", "start_Expr", "start_Type", "program", "languageDecl",
		"extension", "decl", "annotation", "paramDecl", "expr", "patternBinding",
		"binding", "match_case", "pattern", "labelledPattern", "stellatype",
		"recordFieldType", "variantFieldType",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 73, 577, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 5, 3, 48, 8, 3, 10, 3, 12, 3, 51, 9, 3, 1, 3, 5, 3,
		54, 8, 3, 10, 3, 12, 3, 57, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 5, 5, 68, 8, 5, 10, 5, 12, 5, 71, 9, 5, 1, 5, 1, 5, 1, 6,
		5, 6, 76, 8, 6, 10, 6, 12, 6, 79, 9, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 5, 6, 87, 8, 6, 10, 6, 12, 6, 90, 9, 6, 3, 6, 92, 8, 6, 1, 6, 1, 6,
		1, 6, 3, 6, 97, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 103, 8, 6, 10, 6, 12,
		6, 106, 9, 6, 3, 6, 108, 8, 6, 1, 6, 1, 6, 5, 6, 112, 8, 6, 10, 6, 12,
		6, 115, 9, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 125,
		8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		5, 9, 222, 8, 9, 10, 9, 12, 9, 225, 9, 9, 3, 9, 227, 8, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 239, 8, 9, 10, 9,
		12, 9, 242, 9, 9, 3, 9, 244, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9,
		251, 8, 9, 10, 9, 12, 9, 254, 9, 9, 3, 9, 256, 8, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 3, 9, 263, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		5, 9, 272, 8, 9, 10, 9, 12, 9, 275, 9, 9, 3, 9, 277, 8, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 286, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 300, 8, 9, 10, 9,
		12, 9, 303, 9, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 312,
		8, 9, 10, 9, 12, 9, 315, 9, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		3, 9, 324, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 376, 8, 9, 10, 9, 12, 9, 379, 9, 9,
		3, 9, 381, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 389, 8, 9, 10,
		9, 12, 9, 392, 9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 410,
		8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 427, 8, 13, 10, 13, 12, 13,
		430, 9, 13, 3, 13, 432, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13,
		439, 8, 13, 10, 13, 12, 13, 442, 9, 13, 3, 13, 444, 8, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 5, 13, 451, 8, 13, 10, 13, 12, 13, 454, 9, 13, 3,
		13, 456, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 3, 13, 480, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 494, 8,
		15, 10, 15, 12, 15, 497, 9, 15, 3, 15, 499, 8, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 513,
		8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 520, 8, 15, 10, 15, 12,
		15, 523, 9, 15, 3, 15, 525, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5,
		15, 532, 8, 15, 10, 15, 12, 15, 535, 9, 15, 3, 15, 537, 8, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 5, 15, 544, 8, 15, 10, 15, 12, 15, 547, 9, 15,
		3, 15, 549, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3,
		15, 558, 8, 15, 1, 15, 1, 15, 1, 15, 5, 15, 563, 8, 15, 10, 15, 12, 15,
		566, 9, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 3, 17, 575,
		8, 17, 1, 17, 0, 2, 18, 30, 18, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 0, 0, 670, 0, 36, 1, 0, 0, 0, 2, 39, 1, 0,
		0, 0, 4, 42, 1, 0, 0, 0, 6, 45, 1, 0, 0, 0, 8, 58, 1, 0, 0, 0, 10, 62,
		1, 0, 0, 0, 12, 124, 1, 0, 0, 0, 14, 126, 1, 0, 0, 0, 16, 128, 1, 0, 0,
		0, 18, 323, 1, 0, 0, 0, 20, 393, 1, 0, 0, 0, 22, 397, 1, 0, 0, 0, 24, 401,
		1, 0, 0, 0, 26, 479, 1, 0, 0, 0, 28, 481, 1, 0, 0, 0, 30, 557, 1, 0, 0,
		0, 32, 567, 1, 0, 0, 0, 34, 571, 1, 0, 0, 0, 36, 37, 3, 6, 3, 0, 37, 38,
		5, 0, 0, 1, 38, 1, 1, 0, 0, 0, 39, 40, 3, 18, 9, 0, 40, 41, 5, 0, 0, 1,
		41, 3, 1, 0, 0, 0, 42, 43, 3, 30, 15, 0, 43, 44, 5, 0, 0, 1, 44, 5, 1,
		0, 0, 0, 45, 49, 3, 8, 4, 0, 46, 48, 3, 10, 5, 0, 47, 46, 1, 0, 0, 0, 48,
		51, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 55, 1, 0, 0,
		0, 51, 49, 1, 0, 0, 0, 52, 54, 3, 12, 6, 0, 53, 52, 1, 0, 0, 0, 54, 57,
		1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 7, 1, 0, 0, 0,
		57, 55, 1, 0, 0, 0, 58, 59, 5, 51, 0, 0, 59, 60, 5, 39, 0, 0, 60, 61, 5,
		2, 0, 0, 61, 9, 1, 0, 0, 0, 62, 63, 5, 41, 0, 0, 63, 64, 5, 65, 0, 0, 64,
		69, 5, 70, 0, 0, 65, 66, 5, 1, 0, 0, 66, 68, 5, 70, 0, 0, 67, 65, 1, 0,
		0, 0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 72,
		1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 73, 5, 2, 0, 0, 73, 11, 1, 0, 0, 0,
		74, 76, 3, 14, 7, 0, 75, 74, 1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1,
		0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 80, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80,
		81, 5, 44, 0, 0, 81, 82, 5, 69, 0, 0, 82, 91, 5, 3, 0, 0, 83, 88, 3, 16,
		8, 0, 84, 85, 5, 1, 0, 0, 85, 87, 3, 16, 8, 0, 86, 84, 1, 0, 0, 0, 87,
		90, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 92, 1, 0, 0,
		0, 90, 88, 1, 0, 0, 0, 91, 83, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93,
		1, 0, 0, 0, 93, 96, 5, 4, 0, 0, 94, 95, 5, 9, 0, 0, 95, 97, 3, 30, 15,
		0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 107, 1, 0, 0, 0, 98, 99,
		5, 60, 0, 0, 99, 104, 3, 30, 15, 0, 100, 101, 5, 1, 0, 0, 101, 103, 3,
		30, 15, 0, 102, 100, 1, 0, 0, 0, 103, 106, 1, 0, 0, 0, 104, 102, 1, 0,
		0, 0, 104, 105, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0,
		107, 98, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 113,
		5, 5, 0, 0, 110, 112, 3, 12, 6, 0, 111, 110, 1, 0, 0, 0, 112, 115, 1, 0,
		0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 116, 1, 0, 0, 0,
		115, 113, 1, 0, 0, 0, 116, 117, 5, 57, 0, 0, 117, 118, 3, 18, 9, 0, 118,
		119, 5, 6, 0, 0, 119, 125, 1, 0, 0, 0, 120, 121, 5, 62, 0, 0, 121, 122,
		5, 69, 0, 0, 122, 123, 5, 7, 0, 0, 123, 125, 3, 30, 15, 0, 124, 77, 1,
		0, 0, 0, 124, 120, 1, 0, 0, 0, 125, 13, 1, 0, 0, 0, 126, 127, 5, 49, 0,
		0, 127, 15, 1, 0, 0, 0, 128, 129, 5, 69, 0, 0, 129, 130, 5, 8, 0, 0, 130,
		131, 3, 30, 15, 0, 131, 17, 1, 0, 0, 0, 132, 133, 6, 9, -1, 0, 133, 324,
		5, 61, 0, 0, 134, 324, 5, 42, 0, 0, 135, 324, 5, 64, 0, 0, 136, 324, 5,
		71, 0, 0, 137, 324, 5, 69, 0, 0, 138, 139, 5, 48, 0, 0, 139, 140, 5, 3,
		0, 0, 140, 141, 3, 18, 9, 0, 141, 142, 5, 4, 0, 0, 142, 324, 1, 0, 0, 0,
		143, 144, 5, 50, 0, 0, 144, 145, 5, 3, 0, 0, 145, 146, 3, 18, 9, 0, 146,
		147, 5, 4, 0, 0, 147, 324, 1, 0, 0, 0, 148, 149, 5, 38, 0, 0, 149, 150,
		5, 3, 0, 0, 150, 151, 3, 18, 9, 0, 151, 152, 5, 1, 0, 0, 152, 153, 3, 18,
		9, 0, 153, 154, 5, 4, 0, 0, 154, 324, 1, 0, 0, 0, 155, 156, 5, 27, 0, 0,
		156, 157, 5, 3, 0, 0, 157, 158, 3, 18, 9, 0, 158, 159, 5, 4, 0, 0, 159,
		324, 1, 0, 0, 0, 160, 161, 5, 28, 0, 0, 161, 162, 5, 3, 0, 0, 162, 163,
		3, 18, 9, 0, 163, 164, 5, 4, 0, 0, 164, 324, 1, 0, 0, 0, 165, 166, 5, 29,
		0, 0, 166, 167, 5, 3, 0, 0, 167, 168, 3, 18, 9, 0, 168, 169, 5, 4, 0, 0,
		169, 324, 1, 0, 0, 0, 170, 171, 5, 58, 0, 0, 171, 172, 5, 3, 0, 0, 172,
		173, 3, 18, 9, 0, 173, 174, 5, 4, 0, 0, 174, 324, 1, 0, 0, 0, 175, 176,
		5, 55, 0, 0, 176, 177, 5, 3, 0, 0, 177, 178, 3, 18, 9, 0, 178, 179, 5,
		4, 0, 0, 179, 324, 1, 0, 0, 0, 180, 181, 5, 30, 0, 0, 181, 182, 5, 3, 0,
		0, 182, 183, 3, 18, 9, 0, 183, 184, 5, 4, 0, 0, 184, 324, 1, 0, 0, 0, 185,
		186, 5, 31, 0, 0, 186, 187, 5, 3, 0, 0, 187, 188, 3, 18, 9, 0, 188, 189,
		5, 4, 0, 0, 189, 324, 1, 0, 0, 0, 190, 191, 5, 43, 0, 0, 191, 192, 5, 3,
		0, 0, 192, 193, 3, 18, 9, 0, 193, 194, 5, 4, 0, 0, 194, 324, 1, 0, 0, 0,
		195, 196, 5, 32, 0, 0, 196, 197, 5, 3, 0, 0, 197, 198, 3, 18, 9, 0, 198,
		199, 5, 1, 0, 0, 199, 200, 3, 18, 9, 0, 200, 201, 5, 1, 0, 0, 201, 202,
		3, 18, 9, 0, 202, 203, 5, 4, 0, 0, 203, 324, 1, 0, 0, 0, 204, 205, 5, 45,
		0, 0, 205, 206, 5, 14, 0, 0, 206, 207, 3, 30, 15, 0, 207, 208, 5, 15, 0,
		0, 208, 209, 3, 18, 9, 28, 209, 324, 1, 0, 0, 0, 210, 211, 5, 63, 0, 0,
		211, 212, 5, 14, 0, 0, 212, 213, 3, 30, 15, 0, 213, 214, 5, 15, 0, 0, 214,
		215, 3, 18, 9, 27, 215, 324, 1, 0, 0, 0, 216, 217, 5, 44, 0, 0, 217, 226,
		5, 3, 0, 0, 218, 223, 3, 16, 8, 0, 219, 220, 5, 1, 0, 0, 220, 222, 3, 16,
		8, 0, 221, 219, 1, 0, 0, 0, 222, 225, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0,
		223, 224, 1, 0, 0, 0, 224, 227, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 226,
		218, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 229,
		5, 4, 0, 0, 229, 230, 5, 5, 0, 0, 230, 231, 5, 57, 0, 0, 231, 232, 3, 18,
		9, 0, 232, 233, 5, 6, 0, 0, 233, 324, 1, 0, 0, 0, 234, 243, 5, 5, 0, 0,
		235, 240, 3, 18, 9, 0, 236, 237, 5, 1, 0, 0, 237, 239, 3, 18, 9, 0, 238,
		236, 1, 0, 0, 0, 239, 242, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 240, 241,
		1, 0, 0, 0, 241, 244, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 243, 235, 1, 0,
		0, 0, 243, 244, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 324, 5, 6, 0, 0,
		246, 255, 5, 5, 0, 0, 247, 252, 3, 22, 11, 0, 248, 249, 5, 1, 0, 0, 249,
		251, 3, 22, 11, 0, 250, 248, 1, 0, 0, 0, 251, 254, 1, 0, 0, 0, 252, 250,
		1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 256, 1, 0, 0, 0, 254, 252, 1, 0,
		0, 0, 255, 247, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0,
		257, 324, 5, 6, 0, 0, 258, 259, 5, 12, 0, 0, 259, 262, 5, 69, 0, 0, 260,
		261, 5, 7, 0, 0, 261, 263, 3, 18, 9, 0, 262, 260, 1, 0, 0, 0, 262, 263,
		1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 324, 5, 13, 0, 0, 265, 266, 5, 54,
		0, 0, 266, 267, 3, 18, 9, 0, 267, 276, 5, 5, 0, 0, 268, 273, 3, 24, 12,
		0, 269, 270, 5, 11, 0, 0, 270, 272, 3, 24, 12, 0, 271, 269, 1, 0, 0, 0,
		272, 275, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274,
		277, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 276, 268, 1, 0, 0, 0, 276, 277,
		1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278, 279, 5, 6, 0, 0, 279, 324, 1, 0,
		0, 0, 280, 285, 5, 14, 0, 0, 281, 282, 3, 18, 9, 0, 282, 283, 5, 1, 0,
		0, 283, 284, 3, 18, 9, 0, 284, 286, 1, 0, 0, 0, 285, 281, 1, 0, 0, 0, 285,
		286, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 324, 5, 15, 0, 0, 288, 289,
		5, 46, 0, 0, 289, 290, 3, 18, 9, 0, 290, 291, 5, 59, 0, 0, 291, 292, 3,
		18, 9, 0, 292, 293, 5, 40, 0, 0, 293, 294, 3, 18, 9, 6, 294, 324, 1, 0,
		0, 0, 295, 296, 5, 52, 0, 0, 296, 301, 3, 20, 10, 0, 297, 298, 5, 1, 0,
		0, 298, 300, 3, 20, 10, 0, 299, 297, 1, 0, 0, 0, 300, 303, 1, 0, 0, 0,
		301, 299, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 304, 1, 0, 0, 0, 303,
		301, 1, 0, 0, 0, 304, 305, 5, 47, 0, 0, 305, 306, 3, 18, 9, 5, 306, 324,
		1, 0, 0, 0, 307, 308, 5, 53, 0, 0, 308, 313, 3, 20, 10, 0, 309, 310, 5,
		1, 0, 0, 310, 312, 3, 20, 10, 0, 311, 309, 1, 0, 0, 0, 312, 315, 1, 0,
		0, 0, 313, 311, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 316, 1, 0, 0, 0,
		315, 313, 1, 0, 0, 0, 316, 317, 5, 47, 0, 0, 317, 318, 3, 18, 9, 4, 318,
		324, 1, 0, 0, 0, 319, 320, 5, 3, 0, 0, 320, 321, 3, 18, 9, 0, 321, 322,
		5, 4, 0, 0, 322, 324, 1, 0, 0, 0, 323, 132, 1, 0, 0, 0, 323, 134, 1, 0,
		0, 0, 323, 135, 1, 0, 0, 0, 323, 136, 1, 0, 0, 0, 323, 137, 1, 0, 0, 0,
		323, 138, 1, 0, 0, 0, 323, 143, 1, 0, 0, 0, 323, 148, 1, 0, 0, 0, 323,
		155, 1, 0, 0, 0, 323, 160, 1, 0, 0, 0, 323, 165, 1, 0, 0, 0, 323, 170,
		1, 0, 0, 0, 323, 175, 1, 0, 0, 0, 323, 180, 1, 0, 0, 0, 323, 185, 1, 0,
		0, 0, 323, 190, 1, 0, 0, 0, 323, 195, 1, 0, 0, 0, 323, 204, 1, 0, 0, 0,
		323, 210, 1, 0, 0, 0, 323, 216, 1, 0, 0, 0, 323, 234, 1, 0, 0, 0, 323,
		246, 1, 0, 0, 0, 323, 258, 1, 0, 0, 0, 323, 265, 1, 0, 0, 0, 323, 280,
		1, 0, 0, 0, 323, 288, 1, 0, 0, 0, 323, 295, 1, 0, 0, 0, 323, 307, 1, 0,
		0, 0, 323, 319, 1, 0, 0, 0, 324, 390, 1, 0, 0, 0, 325, 326, 10, 25, 0,
		0, 326, 327, 5, 24, 0, 0, 327, 389, 3, 18, 9, 26, 328, 329, 10, 24, 0,
		0, 329, 330, 5, 25, 0, 0, 330, 389, 3, 18, 9, 25, 331, 332, 10, 23, 0,
		0, 332, 333, 5, 36, 0, 0, 333, 389, 3, 18, 9, 24, 334, 335, 10, 22, 0,
		0, 335, 336, 5, 22, 0, 0, 336, 389, 3, 18, 9, 23, 337, 338, 10, 21, 0,
		0, 338, 339, 5, 23, 0, 0, 339, 389, 3, 18, 9, 22, 340, 341, 10, 20, 0,
		0, 341, 342, 5, 56, 0, 0, 342, 389, 3, 18, 9, 21, 343, 344, 10, 12, 0,
		0, 344, 345, 5, 16, 0, 0, 345, 389, 3, 18, 9, 13, 346, 347, 10, 11, 0,
		0, 347, 348, 5, 17, 0, 0, 348, 389, 3, 18, 9, 12, 349, 350, 10, 10, 0,
		0, 350, 351, 5, 18, 0, 0, 351, 389, 3, 18, 9, 11, 352, 353, 10, 9, 0, 0,
		353, 354, 5, 19, 0, 0, 354, 389, 3, 18, 9, 10, 355, 356, 10, 8, 0, 0, 356,
		357, 5, 20, 0, 0, 357, 389, 3, 18, 9, 9, 358, 359, 10, 7, 0, 0, 359, 360,
		5, 21, 0, 0, 360, 389, 3, 18, 9, 8, 361, 362, 10, 2, 0, 0, 362, 363, 5,
		2, 0, 0, 363, 389, 3, 18, 9, 3, 364, 365, 10, 47, 0, 0, 365, 366, 5, 26,
		0, 0, 366, 389, 5, 69, 0, 0, 367, 368, 10, 46, 0, 0, 368, 369, 5, 26, 0,
		0, 369, 389, 5, 71, 0, 0, 370, 371, 10, 26, 0, 0, 371, 380, 5, 3, 0, 0,
		372, 377, 3, 18, 9, 0, 373, 374, 5, 1, 0, 0, 374, 376, 3, 18, 9, 0, 375,
		373, 1, 0, 0, 0, 376, 379, 1, 0, 0, 0, 377, 375, 1, 0, 0, 0, 377, 378,
		1, 0, 0, 0, 378, 381, 1, 0, 0, 0, 379, 377, 1, 0, 0, 0, 380, 372, 1, 0,
		0, 0, 380, 381, 1, 0, 0, 0, 381, 382, 1, 0, 0, 0, 382, 389, 5, 4, 0, 0,
		383, 384, 10, 19, 0, 0, 384, 385, 5, 37, 0, 0, 385, 389, 3, 30, 15, 0,
		386, 387, 10, 1, 0, 0, 387, 389, 5, 2, 0, 0, 388, 325, 1, 0, 0, 0, 388,
		328, 1, 0, 0, 0, 388, 331, 1, 0, 0, 0, 388, 334, 1, 0, 0, 0, 388, 337,
		1, 0, 0, 0, 388, 340, 1, 0, 0, 0, 388, 343, 1, 0, 0, 0, 388, 346, 1, 0,
		0, 0, 388, 349, 1, 0, 0, 0, 388, 352, 1, 0, 0, 0, 388, 355, 1, 0, 0, 0,
		388, 358, 1, 0, 0, 0, 388, 361, 1, 0, 0, 0, 388, 364, 1, 0, 0, 0, 388,
		367, 1, 0, 0, 0, 388, 370, 1, 0, 0, 0, 388, 383, 1, 0, 0, 0, 388, 386,
		1, 0, 0, 0, 389, 392, 1, 0, 0, 0, 390, 388, 1, 0, 0, 0, 390, 391, 1, 0,
		0, 0, 391, 19, 1, 0, 0, 0, 392, 390, 1, 0, 0, 0, 393, 394, 3, 26, 13, 0,
		394, 395, 5, 7, 0, 0, 395, 396, 3, 18, 9, 0, 396, 21, 1, 0, 0, 0, 397,
		398, 5, 69, 0, 0, 398, 399, 5, 7, 0, 0, 399, 400, 3, 18, 9, 0, 400, 23,
		1, 0, 0, 0, 401, 402, 3, 26, 13, 0, 402, 403, 5, 10, 0, 0, 403, 404, 3,
		18, 9, 0, 404, 25, 1, 0, 0, 0, 405, 406, 5, 12, 0, 0, 406, 409, 5, 69,
		0, 0, 407, 408, 5, 7, 0, 0, 408, 410, 3, 26, 13, 0, 409, 407, 1, 0, 0,
		0, 409, 410, 1, 0, 0, 0, 410, 411, 1, 0, 0, 0, 411, 480, 5, 13, 0, 0, 412,
		413, 5, 48, 0, 0, 413, 414, 5, 3, 0, 0, 414, 415, 3, 26, 13, 0, 415, 416,
		5, 4, 0, 0, 416, 480, 1, 0, 0, 0, 417, 418, 5, 50, 0, 0, 418, 419, 5, 3,
		0, 0, 419, 420, 3, 26, 13, 0, 420, 421, 5, 4, 0, 0, 421, 480, 1, 0, 0,
		0, 422, 431, 5, 5, 0, 0, 423, 428, 3, 26, 13, 0, 424, 425, 5, 1, 0, 0,
		425, 427, 3, 26, 13, 0, 426, 424, 1, 0, 0, 0, 427, 430, 1, 0, 0, 0, 428,
		426, 1, 0, 0, 0, 428, 429, 1, 0, 0, 0, 429, 432, 1, 0, 0, 0, 430, 428,
		1, 0, 0, 0, 431, 423, 1, 0, 0, 0, 431, 432, 1, 0, 0, 0, 432, 433, 1, 0,
		0, 0, 433, 480, 5, 6, 0, 0, 434, 443, 5, 5, 0, 0, 435, 440, 3, 28, 14,
		0, 436, 437, 5, 1, 0, 0, 437, 439, 3, 28, 14, 0, 438, 436, 1, 0, 0, 0,
		439, 442, 1, 0, 0, 0, 440, 438, 1, 0, 0, 0, 440, 441, 1, 0, 0, 0, 441,
		444, 1, 0, 0, 0, 442, 440, 1, 0, 0, 0, 443, 435, 1, 0, 0, 0, 443, 444,
		1, 0, 0, 0, 444, 445, 1, 0, 0, 0, 445, 480, 5, 6, 0, 0, 446, 455, 5, 14,
		0, 0, 447, 452, 3, 26, 13, 0, 448, 449, 5, 1, 0, 0, 449, 451, 3, 26, 13,
		0, 450, 448, 1, 0, 0, 0, 451, 454, 1, 0, 0, 0, 452, 450, 1, 0, 0, 0, 452,
		453, 1, 0, 0, 0, 453, 456, 1, 0, 0, 0, 454, 452, 1, 0, 0, 0, 455, 447,
		1, 0, 0, 0, 455, 456, 1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457, 480, 5, 15,
		0, 0, 458, 459, 5, 38, 0, 0, 459, 460, 5, 3, 0, 0, 460, 461, 3, 26, 13,
		0, 461, 462, 5, 1, 0, 0, 462, 463, 3, 26, 13, 0, 463, 464, 5, 4, 0, 0,
		464, 480, 1, 0, 0, 0, 465, 480, 5, 42, 0, 0, 466, 480, 5, 61, 0, 0, 467,
		480, 5, 64, 0, 0, 468, 480, 5, 71, 0, 0, 469, 470, 5, 58, 0, 0, 470, 471,
		5, 3, 0, 0, 471, 472, 3, 26, 13, 0, 472, 473, 5, 4, 0, 0, 473, 480, 1,
		0, 0, 0, 474, 480, 5, 69, 0, 0, 475, 476, 5, 3, 0, 0, 476, 477, 3, 26,
		13, 0, 477, 478, 5, 4, 0, 0, 478, 480, 1, 0, 0, 0, 479, 405, 1, 0, 0, 0,
		479, 412, 1, 0, 0, 0, 479, 417, 1, 0, 0, 0, 479, 422, 1, 0, 0, 0, 479,
		434, 1, 0, 0, 0, 479, 446, 1, 0, 0, 0, 479, 458, 1, 0, 0, 0, 479, 465,
		1, 0, 0, 0, 479, 466, 1, 0, 0, 0, 479, 467, 1, 0, 0, 0, 479, 468, 1, 0,
		0, 0, 479, 469, 1, 0, 0, 0, 479, 474, 1, 0, 0, 0, 479, 475, 1, 0, 0, 0,
		480, 27, 1, 0, 0, 0, 481, 482, 5, 69, 0, 0, 482, 483, 5, 7, 0, 0, 483,
		484, 3, 26, 13, 0, 484, 29, 1, 0, 0, 0, 485, 486, 6, 15, -1, 0, 486, 558,
		5, 33, 0, 0, 487, 558, 5, 34, 0, 0, 488, 489, 5, 44, 0, 0, 489, 498, 5,
		3, 0, 0, 490, 495, 3, 30, 15, 0, 491, 492, 5, 1, 0, 0, 492, 494, 3, 30,
		15, 0, 493, 491, 1, 0, 0, 0, 494, 497, 1, 0, 0, 0, 495, 493, 1, 0, 0, 0,
		495, 496, 1, 0, 0, 0, 496, 499, 1, 0, 0, 0, 497, 495, 1, 0, 0, 0, 498,
		490, 1, 0, 0, 0, 498, 499, 1, 0, 0, 0, 499, 500, 1, 0, 0, 0, 500, 501,
		5, 4, 0, 0, 501, 502, 5, 9, 0, 0, 502, 558, 3, 30, 15, 10, 503, 504, 5,
		66, 0, 0, 504, 505, 5, 69, 0, 0, 505, 506, 5, 26, 0, 0, 506, 558, 3, 30,
		15, 9, 507, 512, 5, 5, 0, 0, 508, 509, 3, 30, 15, 0, 509, 510, 5, 1, 0,
		0, 510, 511, 3, 30, 15, 0, 511, 513, 1, 0, 0, 0, 512, 508, 1, 0, 0, 0,
		512, 513, 1, 0, 0, 0, 513, 514, 1, 0, 0, 0, 514, 558, 5, 6, 0, 0, 515,
		524, 5, 5, 0, 0, 516, 521, 3, 32, 16, 0, 517, 518, 5, 1, 0, 0, 518, 520,
		3, 32, 16, 0, 519, 517, 1, 0, 0, 0, 520, 523, 1, 0, 0, 0, 521, 519, 1,
		0, 0, 0, 521, 522, 1, 0, 0, 0, 522, 525, 1, 0, 0, 0, 523, 521, 1, 0, 0,
		0, 524, 516, 1, 0, 0, 0, 524, 525, 1, 0, 0, 0, 525, 526, 1, 0, 0, 0, 526,
		558, 5, 6, 0, 0, 527, 536, 5, 5, 0, 0, 528, 533, 3, 34, 17, 0, 529, 530,
		5, 1, 0, 0, 530, 532, 3, 34, 17, 0, 531, 529, 1, 0, 0, 0, 532, 535, 1,
		0, 0, 0, 533, 531, 1, 0, 0, 0, 533, 534, 1, 0, 0, 0, 534, 537, 1, 0, 0,
		0, 535, 533, 1, 0, 0, 0, 536, 528, 1, 0, 0, 0, 536, 537, 1, 0, 0, 0, 537,
		538, 1, 0, 0, 0, 538, 558, 5, 6, 0, 0, 539, 548, 5, 14, 0, 0, 540, 545,
		3, 30, 15, 0, 541, 542, 5, 1, 0, 0, 542, 544, 3, 30, 15, 0, 543, 541, 1,
		0, 0, 0, 544, 547, 1, 0, 0, 0, 545, 543, 1, 0, 0, 0, 545, 546, 1, 0, 0,
		0, 546, 549, 1, 0, 0, 0, 547, 545, 1, 0, 0, 0, 548, 540, 1, 0, 0, 0, 548,
		549, 1, 0, 0, 0, 549, 550, 1, 0, 0, 0, 550, 558, 5, 15, 0, 0, 551, 558,
		5, 35, 0, 0, 552, 558, 5, 69, 0, 0, 553, 554, 5, 3, 0, 0, 554, 555, 3,
		30, 15, 0, 555, 556, 5, 4, 0, 0, 556, 558, 1, 0, 0, 0, 557, 485, 1, 0,
		0, 0, 557, 487, 1, 0, 0, 0, 557, 488, 1, 0, 0, 0, 557, 503, 1, 0, 0, 0,
		557, 507, 1, 0, 0, 0, 557, 515, 1, 0, 0, 0, 557, 527, 1, 0, 0, 0, 557,
		539, 1, 0, 0, 0, 557, 551, 1, 0, 0, 0, 557, 552, 1, 0, 0, 0, 557, 553,
		1, 0, 0, 0, 558, 564, 1, 0, 0, 0, 559, 560, 10, 8, 0, 0, 560, 561, 5, 22,
		0, 0, 561, 563, 3, 30, 15, 9, 562, 559, 1, 0, 0, 0, 563, 566, 1, 0, 0,
		0, 564, 562, 1, 0, 0, 0, 564, 565, 1, 0, 0, 0, 565, 31, 1, 0, 0, 0, 566,
		564, 1, 0, 0, 0, 567, 568, 5, 69, 0, 0, 568, 569, 5, 8, 0, 0, 569, 570,
		3, 30, 15, 0, 570, 33, 1, 0, 0, 0, 571, 574, 5, 69, 0, 0, 572, 573, 5,
		8, 0, 0, 573, 575, 3, 30, 15, 0, 574, 572, 1, 0, 0, 0, 574, 575, 1, 0,
		0, 0, 575, 35, 1, 0, 0, 0, 48, 49, 55, 69, 77, 88, 91, 96, 104, 107, 113,
		124, 223, 226, 240, 243, 252, 255, 262, 273, 276, 285, 301, 313, 323, 377,
		380, 388, 390, 409, 428, 431, 440, 443, 452, 455, 479, 495, 498, 512, 521,
		524, 533, 536, 545, 548, 557, 564, 574,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// StellaParserInit initializes any static state used to implement StellaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewStellaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func StellaParserInit() {
	staticData := &stellaparserParserStaticData
	staticData.once.Do(stellaparserParserInit)
}

// NewStellaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewStellaParser(input antlr.TokenStream) *StellaParser {
	StellaParserInit()
	this := new(StellaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &stellaparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// StellaParser tokens.
const (
	StellaParserEOF                        = antlr.TokenEOF
	StellaParserSurrogate_id_SYMB_0        = 1
	StellaParserSurrogate_id_SYMB_1        = 2
	StellaParserSurrogate_id_SYMB_2        = 3
	StellaParserSurrogate_id_SYMB_3        = 4
	StellaParserSurrogate_id_SYMB_4        = 5
	StellaParserSurrogate_id_SYMB_5        = 6
	StellaParserSurrogate_id_SYMB_6        = 7
	StellaParserSurrogate_id_SYMB_7        = 8
	StellaParserSurrogate_id_SYMB_8        = 9
	StellaParserSurrogate_id_SYMB_9        = 10
	StellaParserSurrogate_id_SYMB_10       = 11
	StellaParserSurrogate_id_SYMB_11       = 12
	StellaParserSurrogate_id_SYMB_12       = 13
	StellaParserSurrogate_id_SYMB_13       = 14
	StellaParserSurrogate_id_SYMB_14       = 15
	StellaParserSurrogate_id_SYMB_15       = 16
	StellaParserSurrogate_id_SYMB_16       = 17
	StellaParserSurrogate_id_SYMB_17       = 18
	StellaParserSurrogate_id_SYMB_18       = 19
	StellaParserSurrogate_id_SYMB_19       = 20
	StellaParserSurrogate_id_SYMB_20       = 21
	StellaParserSurrogate_id_SYMB_21       = 22
	StellaParserSurrogate_id_SYMB_22       = 23
	StellaParserSurrogate_id_SYMB_23       = 24
	StellaParserSurrogate_id_SYMB_24       = 25
	StellaParserSurrogate_id_SYMB_25       = 26
	StellaParserSurrogate_id_SYMB_26       = 27
	StellaParserSurrogate_id_SYMB_27       = 28
	StellaParserSurrogate_id_SYMB_28       = 29
	StellaParserSurrogate_id_SYMB_29       = 30
	StellaParserSurrogate_id_SYMB_30       = 31
	StellaParserSurrogate_id_SYMB_31       = 32
	StellaParserSurrogate_id_SYMB_32       = 33
	StellaParserSurrogate_id_SYMB_33       = 34
	StellaParserSurrogate_id_SYMB_34       = 35
	StellaParserSurrogate_id_SYMB_35       = 36
	StellaParserSurrogate_id_SYMB_36       = 37
	StellaParserSurrogate_id_SYMB_37       = 38
	StellaParserSurrogate_id_SYMB_38       = 39
	StellaParserSurrogate_id_SYMB_39       = 40
	StellaParserSurrogate_id_SYMB_40       = 41
	StellaParserSurrogate_id_SYMB_41       = 42
	StellaParserSurrogate_id_SYMB_42       = 43
	StellaParserSurrogate_id_SYMB_43       = 44
	StellaParserSurrogate_id_SYMB_44       = 45
	StellaParserSurrogate_id_SYMB_45       = 46
	StellaParserSurrogate_id_SYMB_46       = 47
	StellaParserSurrogate_id_SYMB_47       = 48
	StellaParserSurrogate_id_SYMB_48       = 49
	StellaParserSurrogate_id_SYMB_49       = 50
	StellaParserSurrogate_id_SYMB_50       = 51
	StellaParserSurrogate_id_SYMB_51       = 52
	StellaParserSurrogate_id_SYMB_52       = 53
	StellaParserSurrogate_id_SYMB_53       = 54
	StellaParserSurrogate_id_SYMB_54       = 55
	StellaParserSurrogate_id_SYMB_55       = 56
	StellaParserSurrogate_id_SYMB_56       = 57
	StellaParserSurrogate_id_SYMB_57       = 58
	StellaParserSurrogate_id_SYMB_58       = 59
	StellaParserSurrogate_id_SYMB_59       = 60
	StellaParserSurrogate_id_SYMB_60       = 61
	StellaParserSurrogate_id_SYMB_61       = 62
	StellaParserSurrogate_id_SYMB_62       = 63
	StellaParserSurrogate_id_SYMB_63       = 64
	StellaParserSurrogate_id_SYMB_64       = 65
	StellaParserSurrogate_id_SYMB_65       = 66
	StellaParserCOMMENT_antlr_builtin      = 67
	StellaParserMULTICOMMENT_antlr_builtin = 68
	StellaParserStellaIdent                = 69
	StellaParserExtensionName              = 70
	StellaParserINTEGER                    = 71
	StellaParserWS                         = 72
	StellaParserErrorToken                 = 73
)

// StellaParser rules.
const (
	StellaParserRULE_start_Program    = 0
	StellaParserRULE_start_Expr       = 1
	StellaParserRULE_start_Type       = 2
	StellaParserRULE_program          = 3
	StellaParserRULE_languageDecl     = 4
	StellaParserRULE_extension        = 5
	StellaParserRULE_decl             = 6
	StellaParserRULE_annotation       = 7
	StellaParserRULE_paramDecl        = 8
	StellaParserRULE_expr             = 9
	StellaParserRULE_patternBinding   = 10
	StellaParserRULE_binding          = 11
	StellaParserRULE_match_case       = 12
	StellaParserRULE_pattern          = 13
	StellaParserRULE_labelledPattern  = 14
	StellaParserRULE_stellatype       = 15
	StellaParserRULE_recordFieldType  = 16
	StellaParserRULE_variantFieldType = 17
)

// IStart_ProgramContext is an interface to support dynamic dispatch.
type IStart_ProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetX returns the x rule contexts.
	GetX() IProgramContext

	// SetX sets the x rule contexts.
	SetX(IProgramContext)

	// IsStart_ProgramContext differentiates from other interfaces.
	IsStart_ProgramContext()
}

type Start_ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	x      IProgramContext
}

func NewEmptyStart_ProgramContext() *Start_ProgramContext {
	var p = new(Start_ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StellaParserRULE_start_Program
	return p
}

func (*Start_ProgramContext) IsStart_ProgramContext() {}

func NewStart_ProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Start_ProgramContext {
	var p = new(Start_ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StellaParserRULE_start_Program

	return p
}

func (s *Start_ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *Start_ProgramContext) GetX() IProgramContext { return s.x }

func (s *Start_ProgramContext) SetX(v IProgramContext) { s.x = v }

func (s *Start_ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(StellaParserEOF, 0)
}

func (s *Start_ProgramContext) Program() IProgramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramContext)
}

func (s *Start_ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Start_ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Start_ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterStart_Program(s)
	}
}

func (s *Start_ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitStart_Program(s)
	}
}

func (p *StellaParser) Start_Program() (localctx IStart_ProgramContext) {
	this := p
	_ = this

	localctx = NewStart_ProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, StellaParserRULE_start_Program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)

		var _x = p.Program()

		localctx.(*Start_ProgramContext).x = _x
	}
	{
		p.SetState(37)
		p.Match(StellaParserEOF)
	}

	return localctx
}

// IStart_ExprContext is an interface to support dynamic dispatch.
type IStart_ExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetX returns the x rule contexts.
	GetX() IExprContext

	// SetX sets the x rule contexts.
	SetX(IExprContext)

	// IsStart_ExprContext differentiates from other interfaces.
	IsStart_ExprContext()
}

type Start_ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	x      IExprContext
}

func NewEmptyStart_ExprContext() *Start_ExprContext {
	var p = new(Start_ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StellaParserRULE_start_Expr
	return p
}

func (*Start_ExprContext) IsStart_ExprContext() {}

func NewStart_ExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Start_ExprContext {
	var p = new(Start_ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StellaParserRULE_start_Expr

	return p
}

func (s *Start_ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *Start_ExprContext) GetX() IExprContext { return s.x }

func (s *Start_ExprContext) SetX(v IExprContext) { s.x = v }

func (s *Start_ExprContext) EOF() antlr.TerminalNode {
	return s.GetToken(StellaParserEOF, 0)
}

func (s *Start_ExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Start_ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Start_ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Start_ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterStart_Expr(s)
	}
}

func (s *Start_ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitStart_Expr(s)
	}
}

func (p *StellaParser) Start_Expr() (localctx IStart_ExprContext) {
	this := p
	_ = this

	localctx = NewStart_ExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, StellaParserRULE_start_Expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)

		var _x = p.expr(0)

		localctx.(*Start_ExprContext).x = _x
	}
	{
		p.SetState(40)
		p.Match(StellaParserEOF)
	}

	return localctx
}

// IStart_TypeContext is an interface to support dynamic dispatch.
type IStart_TypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetX returns the x rule contexts.
	GetX() IStellatypeContext

	// SetX sets the x rule contexts.
	SetX(IStellatypeContext)

	// IsStart_TypeContext differentiates from other interfaces.
	IsStart_TypeContext()
}

type Start_TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	x      IStellatypeContext
}

func NewEmptyStart_TypeContext() *Start_TypeContext {
	var p = new(Start_TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StellaParserRULE_start_Type
	return p
}

func (*Start_TypeContext) IsStart_TypeContext() {}

func NewStart_TypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Start_TypeContext {
	var p = new(Start_TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StellaParserRULE_start_Type

	return p
}

func (s *Start_TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *Start_TypeContext) GetX() IStellatypeContext { return s.x }

func (s *Start_TypeContext) SetX(v IStellatypeContext) { s.x = v }

func (s *Start_TypeContext) EOF() antlr.TerminalNode {
	return s.GetToken(StellaParserEOF, 0)
}

func (s *Start_TypeContext) Stellatype() IStellatypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStellatypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStellatypeContext)
}

func (s *Start_TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Start_TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Start_TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterStart_Type(s)
	}
}

func (s *Start_TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitStart_Type(s)
	}
}

func (p *StellaParser) Start_Type() (localctx IStart_TypeContext) {
	this := p
	_ = this

	localctx = NewStart_TypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, StellaParserRULE_start_Type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)

		var _x = p.stellatype(0)

		localctx.(*Start_TypeContext).x = _x
	}
	{
		p.SetState(43)
		p.Match(StellaParserEOF)
	}

	return localctx
}

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_extension returns the _extension rule contexts.
	Get_extension() IExtensionContext

	// Get_decl returns the _decl rule contexts.
	Get_decl() IDeclContext

	// Set_extension sets the _extension rule contexts.
	Set_extension(IExtensionContext)

	// Set_decl sets the _decl rule contexts.
	Set_decl(IDeclContext)

	// GetExtensions returns the extensions rule context list.
	GetExtensions() []IExtensionContext

	// GetDecls returns the decls rule context list.
	GetDecls() []IDeclContext

	// SetExtensions sets the extensions rule context list.
	SetExtensions([]IExtensionContext)

	// SetDecls sets the decls rule context list.
	SetDecls([]IDeclContext)

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	_extension IExtensionContext
	extensions []IExtensionContext
	_decl      IDeclContext
	decls      []IDeclContext
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StellaParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StellaParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Get_extension() IExtensionContext { return s._extension }

func (s *ProgramContext) Get_decl() IDeclContext { return s._decl }

func (s *ProgramContext) Set_extension(v IExtensionContext) { s._extension = v }

func (s *ProgramContext) Set_decl(v IDeclContext) { s._decl = v }

func (s *ProgramContext) GetExtensions() []IExtensionContext { return s.extensions }

func (s *ProgramContext) GetDecls() []IDeclContext { return s.decls }

func (s *ProgramContext) SetExtensions(v []IExtensionContext) { s.extensions = v }

func (s *ProgramContext) SetDecls(v []IDeclContext) { s.decls = v }

func (s *ProgramContext) LanguageDecl() ILanguageDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILanguageDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILanguageDeclContext)
}

func (s *ProgramContext) AllExtension() []IExtensionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExtensionContext); ok {
			len++
		}
	}

	tst := make([]IExtensionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExtensionContext); ok {
			tst[i] = t.(IExtensionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Extension(i int) IExtensionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtensionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtensionContext)
}

func (s *ProgramContext) AllDecl() []IDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclContext); ok {
			len++
		}
	}

	tst := make([]IDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclContext); ok {
			tst[i] = t.(IDeclContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Decl(i int) IDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *StellaParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, StellaParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.LanguageDecl()
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == StellaParserSurrogate_id_SYMB_40 {
		{
			p.SetState(46)

			var _x = p.Extension()

			localctx.(*ProgramContext)._extension = _x
		}
		localctx.(*ProgramContext).extensions = append(localctx.(*ProgramContext).extensions, localctx.(*ProgramContext)._extension)

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4612266560566853632) != 0 {
		{
			p.SetState(52)

			var _x = p.Decl()

			localctx.(*ProgramContext)._decl = _x
		}
		localctx.(*ProgramContext).decls = append(localctx.(*ProgramContext).decls, localctx.(*ProgramContext)._decl)

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILanguageDeclContext is an interface to support dynamic dispatch.
type ILanguageDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLanguageDeclContext differentiates from other interfaces.
	IsLanguageDeclContext()
}

type LanguageDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLanguageDeclContext() *LanguageDeclContext {
	var p = new(LanguageDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StellaParserRULE_languageDecl
	return p
}

func (*LanguageDeclContext) IsLanguageDeclContext() {}

func NewLanguageDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LanguageDeclContext {
	var p = new(LanguageDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StellaParserRULE_languageDecl

	return p
}

func (s *LanguageDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *LanguageDeclContext) CopyFrom(ctx *LanguageDeclContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LanguageDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LanguageDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LanguageCoreContext struct {
	*LanguageDeclContext
}

func NewLanguageCoreContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LanguageCoreContext {
	var p = new(LanguageCoreContext)

	p.LanguageDeclContext = NewEmptyLanguageDeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LanguageDeclContext))

	return p
}

func (s *LanguageCoreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LanguageCoreContext) Surrogate_id_SYMB_50() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_50, 0)
}

func (s *LanguageCoreContext) Surrogate_id_SYMB_38() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_38, 0)
}

func (s *LanguageCoreContext) Surrogate_id_SYMB_1() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_1, 0)
}

func (s *LanguageCoreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterLanguageCore(s)
	}
}

func (s *LanguageCoreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitLanguageCore(s)
	}
}

func (p *StellaParser) LanguageDecl() (localctx ILanguageDeclContext) {
	this := p
	_ = this

	localctx = NewLanguageDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, StellaParserRULE_languageDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewLanguageCoreContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(StellaParserSurrogate_id_SYMB_50)
	}
	{
		p.SetState(59)
		p.Match(StellaParserSurrogate_id_SYMB_38)
	}
	{
		p.SetState(60)
		p.Match(StellaParserSurrogate_id_SYMB_1)
	}

	return localctx
}

// IExtensionContext is an interface to support dynamic dispatch.
type IExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtensionContext differentiates from other interfaces.
	IsExtensionContext()
}

type ExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtensionContext() *ExtensionContext {
	var p = new(ExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StellaParserRULE_extension
	return p
}

func (*ExtensionContext) IsExtensionContext() {}

func NewExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensionContext {
	var p = new(ExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StellaParserRULE_extension

	return p
}

func (s *ExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensionContext) CopyFrom(ctx *ExtensionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AnExtensionContext struct {
	*ExtensionContext
	_ExtensionName antlr.Token
	extensionNames []antlr.Token
}

func NewAnExtensionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AnExtensionContext {
	var p = new(AnExtensionContext)

	p.ExtensionContext = NewEmptyExtensionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExtensionContext))

	return p
}

func (s *AnExtensionContext) Get_ExtensionName() antlr.Token { return s._ExtensionName }

func (s *AnExtensionContext) Set_ExtensionName(v antlr.Token) { s._ExtensionName = v }

func (s *AnExtensionContext) GetExtensionNames() []antlr.Token { return s.extensionNames }

func (s *AnExtensionContext) SetExtensionNames(v []antlr.Token) { s.extensionNames = v }

func (s *AnExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnExtensionContext) Surrogate_id_SYMB_40() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_40, 0)
}

func (s *AnExtensionContext) Surrogate_id_SYMB_64() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_64, 0)
}

func (s *AnExtensionContext) Surrogate_id_SYMB_1() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_1, 0)
}

func (s *AnExtensionContext) AllExtensionName() []antlr.TerminalNode {
	return s.GetTokens(StellaParserExtensionName)
}

func (s *AnExtensionContext) ExtensionName(i int) antlr.TerminalNode {
	return s.GetToken(StellaParserExtensionName, i)
}

func (s *AnExtensionContext) AllSurrogate_id_SYMB_0() []antlr.TerminalNode {
	return s.GetTokens(StellaParserSurrogate_id_SYMB_0)
}

func (s *AnExtensionContext) Surrogate_id_SYMB_0(i int) antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, i)
}

func (s *AnExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterAnExtension(s)
	}
}

func (s *AnExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitAnExtension(s)
	}
}

func (p *StellaParser) Extension() (localctx IExtensionContext) {
	this := p
	_ = this

	localctx = NewExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, StellaParserRULE_extension)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewAnExtensionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(StellaParserSurrogate_id_SYMB_40)
	}
	{
		p.SetState(63)
		p.Match(StellaParserSurrogate_id_SYMB_64)
	}
	{
		p.SetState(64)

		var _m = p.Match(StellaParserExtensionName)

		localctx.(*AnExtensionContext)._ExtensionName = _m
	}
	localctx.(*AnExtensionContext).extensionNames = append(localctx.(*AnExtensionContext).extensionNames, localctx.(*AnExtensionContext)._ExtensionName)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == StellaParserSurrogate_id_SYMB_0 {
		{
			p.SetState(65)
			p.Match(StellaParserSurrogate_id_SYMB_0)
		}
		{
			p.SetState(66)

			var _m = p.Match(StellaParserExtensionName)

			localctx.(*AnExtensionContext)._ExtensionName = _m
		}
		localctx.(*AnExtensionContext).extensionNames = append(localctx.(*AnExtensionContext).extensionNames, localctx.(*AnExtensionContext)._ExtensionName)

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(72)
		p.Match(StellaParserSurrogate_id_SYMB_1)
	}

	return localctx
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StellaParserRULE_decl
	return p
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StellaParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) CopyFrom(ctx *DeclContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclTypeAliasContext struct {
	*DeclContext
	name  antlr.Token
	atype IStellatypeContext
}

func NewDeclTypeAliasContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclTypeAliasContext {
	var p = new(DeclTypeAliasContext)

	p.DeclContext = NewEmptyDeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclContext))

	return p
}

func (s *DeclTypeAliasContext) GetName() antlr.Token { return s.name }

func (s *DeclTypeAliasContext) SetName(v antlr.Token) { s.name = v }

func (s *DeclTypeAliasContext) GetAtype() IStellatypeContext { return s.atype }

func (s *DeclTypeAliasContext) SetAtype(v IStellatypeContext) { s.atype = v }

func (s *DeclTypeAliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclTypeAliasContext) Surrogate_id_SYMB_61() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_61, 0)
}

func (s *DeclTypeAliasContext) Surrogate_id_SYMB_6() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_6, 0)
}

func (s *DeclTypeAliasContext) StellaIdent() antlr.TerminalNode {
	return s.GetToken(StellaParserStellaIdent, 0)
}

func (s *DeclTypeAliasContext) Stellatype() IStellatypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStellatypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStellatypeContext)
}

func (s *DeclTypeAliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterDeclTypeAlias(s)
	}
}

func (s *DeclTypeAliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitDeclTypeAlias(s)
	}
}

type DeclFunContext struct {
	*DeclContext
	_annotation IAnnotationContext
	annotations []IAnnotationContext
	name        antlr.Token
	_paramDecl  IParamDeclContext
	paramDecls  []IParamDeclContext
	returnType  IStellatypeContext
	_stellatype IStellatypeContext
	throwTypes  []IStellatypeContext
	_decl       IDeclContext
	localDecls  []IDeclContext
	returnExpr  IExprContext
}

func NewDeclFunContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclFunContext {
	var p = new(DeclFunContext)

	p.DeclContext = NewEmptyDeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclContext))

	return p
}

func (s *DeclFunContext) GetName() antlr.Token { return s.name }

func (s *DeclFunContext) SetName(v antlr.Token) { s.name = v }

func (s *DeclFunContext) Get_annotation() IAnnotationContext { return s._annotation }

func (s *DeclFunContext) Get_paramDecl() IParamDeclContext { return s._paramDecl }

func (s *DeclFunContext) GetReturnType() IStellatypeContext { return s.returnType }

func (s *DeclFunContext) Get_stellatype() IStellatypeContext { return s._stellatype }

func (s *DeclFunContext) Get_decl() IDeclContext { return s._decl }

func (s *DeclFunContext) GetReturnExpr() IExprContext { return s.returnExpr }

func (s *DeclFunContext) Set_annotation(v IAnnotationContext) { s._annotation = v }

func (s *DeclFunContext) Set_paramDecl(v IParamDeclContext) { s._paramDecl = v }

func (s *DeclFunContext) SetReturnType(v IStellatypeContext) { s.returnType = v }

func (s *DeclFunContext) Set_stellatype(v IStellatypeContext) { s._stellatype = v }

func (s *DeclFunContext) Set_decl(v IDeclContext) { s._decl = v }

func (s *DeclFunContext) SetReturnExpr(v IExprContext) { s.returnExpr = v }

func (s *DeclFunContext) GetAnnotations() []IAnnotationContext { return s.annotations }

func (s *DeclFunContext) GetParamDecls() []IParamDeclContext { return s.paramDecls }

func (s *DeclFunContext) GetThrowTypes() []IStellatypeContext { return s.throwTypes }

func (s *DeclFunContext) GetLocalDecls() []IDeclContext { return s.localDecls }

func (s *DeclFunContext) SetAnnotations(v []IAnnotationContext) { s.annotations = v }

func (s *DeclFunContext) SetParamDecls(v []IParamDeclContext) { s.paramDecls = v }

func (s *DeclFunContext) SetThrowTypes(v []IStellatypeContext) { s.throwTypes = v }

func (s *DeclFunContext) SetLocalDecls(v []IDeclContext) { s.localDecls = v }

func (s *DeclFunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclFunContext) Surrogate_id_SYMB_43() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_43, 0)
}

func (s *DeclFunContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *DeclFunContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *DeclFunContext) Surrogate_id_SYMB_4() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_4, 0)
}

func (s *DeclFunContext) Surrogate_id_SYMB_56() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_56, 0)
}

func (s *DeclFunContext) Surrogate_id_SYMB_5() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_5, 0)
}

func (s *DeclFunContext) StellaIdent() antlr.TerminalNode {
	return s.GetToken(StellaParserStellaIdent, 0)
}

func (s *DeclFunContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclFunContext) Surrogate_id_SYMB_8() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_8, 0)
}

func (s *DeclFunContext) Surrogate_id_SYMB_59() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_59, 0)
}

func (s *DeclFunContext) AllAnnotation() []IAnnotationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAnnotationContext); ok {
			len++
		}
	}

	tst := make([]IAnnotationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAnnotationContext); ok {
			tst[i] = t.(IAnnotationContext)
			i++
		}
	}

	return tst
}

func (s *DeclFunContext) Annotation(i int) IAnnotationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *DeclFunContext) AllParamDecl() []IParamDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamDeclContext); ok {
			len++
		}
	}

	tst := make([]IParamDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamDeclContext); ok {
			tst[i] = t.(IParamDeclContext)
			i++
		}
	}

	return tst
}

func (s *DeclFunContext) ParamDecl(i int) IParamDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamDeclContext)
}

func (s *DeclFunContext) AllStellatype() []IStellatypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStellatypeContext); ok {
			len++
		}
	}

	tst := make([]IStellatypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStellatypeContext); ok {
			tst[i] = t.(IStellatypeContext)
			i++
		}
	}

	return tst
}

func (s *DeclFunContext) Stellatype(i int) IStellatypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStellatypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStellatypeContext)
}

func (s *DeclFunContext) AllDecl() []IDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclContext); ok {
			len++
		}
	}

	tst := make([]IDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclContext); ok {
			tst[i] = t.(IDeclContext)
			i++
		}
	}

	return tst
}

func (s *DeclFunContext) Decl(i int) IDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *DeclFunContext) AllSurrogate_id_SYMB_0() []antlr.TerminalNode {
	return s.GetTokens(StellaParserSurrogate_id_SYMB_0)
}

func (s *DeclFunContext) Surrogate_id_SYMB_0(i int) antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, i)
}

func (s *DeclFunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterDeclFun(s)
	}
}

func (s *DeclFunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitDeclFun(s)
	}
}

func (p *StellaParser) Decl() (localctx IDeclContext) {
	this := p
	_ = this

	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, StellaParserRULE_decl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(124)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StellaParserSurrogate_id_SYMB_43, StellaParserSurrogate_id_SYMB_48:
		localctx = NewDeclFunContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StellaParserSurrogate_id_SYMB_48 {
			{
				p.SetState(74)

				var _x = p.Annotation()

				localctx.(*DeclFunContext)._annotation = _x
			}
			localctx.(*DeclFunContext).annotations = append(localctx.(*DeclFunContext).annotations, localctx.(*DeclFunContext)._annotation)

			p.SetState(79)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(80)
			p.Match(StellaParserSurrogate_id_SYMB_43)
		}
		{
			p.SetState(81)

			var _m = p.Match(StellaParserStellaIdent)

			localctx.(*DeclFunContext).name = _m
		}
		{
			p.SetState(82)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == StellaParserStellaIdent {
			{
				p.SetState(83)

				var _x = p.ParamDecl()

				localctx.(*DeclFunContext)._paramDecl = _x
			}
			localctx.(*DeclFunContext).paramDecls = append(localctx.(*DeclFunContext).paramDecls, localctx.(*DeclFunContext)._paramDecl)
			p.SetState(88)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == StellaParserSurrogate_id_SYMB_0 {
				{
					p.SetState(84)
					p.Match(StellaParserSurrogate_id_SYMB_0)
				}
				{
					p.SetState(85)

					var _x = p.ParamDecl()

					localctx.(*DeclFunContext)._paramDecl = _x
				}
				localctx.(*DeclFunContext).paramDecls = append(localctx.(*DeclFunContext).paramDecls, localctx.(*DeclFunContext)._paramDecl)

				p.SetState(90)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(93)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == StellaParserSurrogate_id_SYMB_8 {
			{
				p.SetState(94)
				p.Match(StellaParserSurrogate_id_SYMB_8)
			}
			{
				p.SetState(95)

				var _x = p.stellatype(0)

				localctx.(*DeclFunContext).returnType = _x
			}

		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == StellaParserSurrogate_id_SYMB_59 {
			{
				p.SetState(98)
				p.Match(StellaParserSurrogate_id_SYMB_59)
			}
			{
				p.SetState(99)

				var _x = p.stellatype(0)

				localctx.(*DeclFunContext)._stellatype = _x
			}
			localctx.(*DeclFunContext).throwTypes = append(localctx.(*DeclFunContext).throwTypes, localctx.(*DeclFunContext)._stellatype)
			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == StellaParserSurrogate_id_SYMB_0 {
				{
					p.SetState(100)
					p.Match(StellaParserSurrogate_id_SYMB_0)
				}
				{
					p.SetState(101)

					var _x = p.stellatype(0)

					localctx.(*DeclFunContext)._stellatype = _x
				}
				localctx.(*DeclFunContext).throwTypes = append(localctx.(*DeclFunContext).throwTypes, localctx.(*DeclFunContext)._stellatype)

				p.SetState(106)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(109)
			p.Match(StellaParserSurrogate_id_SYMB_4)
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4612266560566853632) != 0 {
			{
				p.SetState(110)

				var _x = p.Decl()

				localctx.(*DeclFunContext)._decl = _x
			}
			localctx.(*DeclFunContext).localDecls = append(localctx.(*DeclFunContext).localDecls, localctx.(*DeclFunContext)._decl)

			p.SetState(115)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(116)
			p.Match(StellaParserSurrogate_id_SYMB_56)
		}
		{
			p.SetState(117)

			var _x = p.expr(0)

			localctx.(*DeclFunContext).returnExpr = _x
		}
		{
			p.SetState(118)
			p.Match(StellaParserSurrogate_id_SYMB_5)
		}

	case StellaParserSurrogate_id_SYMB_61:
		localctx = NewDeclTypeAliasContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Match(StellaParserSurrogate_id_SYMB_61)
		}
		{
			p.SetState(121)

			var _m = p.Match(StellaParserStellaIdent)

			localctx.(*DeclTypeAliasContext).name = _m
		}
		{
			p.SetState(122)
			p.Match(StellaParserSurrogate_id_SYMB_6)
		}
		{
			p.SetState(123)

			var _x = p.stellatype(0)

			localctx.(*DeclTypeAliasContext).atype = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAnnotationContext is an interface to support dynamic dispatch.
type IAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotationContext differentiates from other interfaces.
	IsAnnotationContext()
}

type AnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StellaParserRULE_annotation
	return p
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StellaParserRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationContext) CopyFrom(ctx *AnnotationContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InlineAnnotationContext struct {
	*AnnotationContext
}

func NewInlineAnnotationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InlineAnnotationContext {
	var p = new(InlineAnnotationContext)

	p.AnnotationContext = NewEmptyAnnotationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AnnotationContext))

	return p
}

func (s *InlineAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineAnnotationContext) Surrogate_id_SYMB_48() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_48, 0)
}

func (s *InlineAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterInlineAnnotation(s)
	}
}

func (s *InlineAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitInlineAnnotation(s)
	}
}

func (p *StellaParser) Annotation() (localctx IAnnotationContext) {
	this := p
	_ = this

	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, StellaParserRULE_annotation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewInlineAnnotationContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(StellaParserSurrogate_id_SYMB_48)
	}

	return localctx
}

// IParamDeclContext is an interface to support dynamic dispatch.
type IParamDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetParamType returns the paramType rule contexts.
	GetParamType() IStellatypeContext

	// SetParamType sets the paramType rule contexts.
	SetParamType(IStellatypeContext)

	// IsParamDeclContext differentiates from other interfaces.
	IsParamDeclContext()
}

type ParamDeclContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	name      antlr.Token
	paramType IStellatypeContext
}

func NewEmptyParamDeclContext() *ParamDeclContext {
	var p = new(ParamDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StellaParserRULE_paramDecl
	return p
}

func (*ParamDeclContext) IsParamDeclContext() {}

func NewParamDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamDeclContext {
	var p = new(ParamDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StellaParserRULE_paramDecl

	return p
}

func (s *ParamDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamDeclContext) GetName() antlr.Token { return s.name }

func (s *ParamDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *ParamDeclContext) GetParamType() IStellatypeContext { return s.paramType }

func (s *ParamDeclContext) SetParamType(v IStellatypeContext) { s.paramType = v }

func (s *ParamDeclContext) Surrogate_id_SYMB_7() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_7, 0)
}

func (s *ParamDeclContext) StellaIdent() antlr.TerminalNode {
	return s.GetToken(StellaParserStellaIdent, 0)
}

func (s *ParamDeclContext) Stellatype() IStellatypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStellatypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStellatypeContext)
}

func (s *ParamDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterParamDecl(s)
	}
}

func (s *ParamDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitParamDecl(s)
	}
}

func (p *StellaParser) ParamDecl() (localctx IParamDeclContext) {
	this := p
	_ = this

	localctx = NewParamDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, StellaParserRULE_paramDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)

		var _m = p.Match(StellaParserStellaIdent)

		localctx.(*ParamDeclContext).name = _m
	}
	{
		p.SetState(129)
		p.Match(StellaParserSurrogate_id_SYMB_7)
	}
	{
		p.SetState(130)

		var _x = p.stellatype(0)

		localctx.(*ParamDeclContext).paramType = _x
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StellaParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StellaParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IsEmptyContext struct {
	*ExprContext
	list IExprContext
}

func NewIsEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsEmptyContext {
	var p = new(IsEmptyContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IsEmptyContext) GetList() IExprContext { return s.list }

func (s *IsEmptyContext) SetList(v IExprContext) { s.list = v }

func (s *IsEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsEmptyContext) Surrogate_id_SYMB_27() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_27, 0)
}

func (s *IsEmptyContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *IsEmptyContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *IsEmptyContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IsEmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterIsEmpty(s)
	}
}

func (s *IsEmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitIsEmpty(s)
	}
}

type FoldContext struct {
	*ExprContext
	type_ IStellatypeContext
	expr_ IExprContext
}

func NewFoldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FoldContext {
	var p = new(FoldContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FoldContext) GetType_() IStellatypeContext { return s.type_ }

func (s *FoldContext) GetExpr_() IExprContext { return s.expr_ }

func (s *FoldContext) SetType_(v IStellatypeContext) { s.type_ = v }

func (s *FoldContext) SetExpr_(v IExprContext) { s.expr_ = v }

func (s *FoldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FoldContext) Surrogate_id_SYMB_44() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_44, 0)
}

func (s *FoldContext) Surrogate_id_SYMB_13() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_13, 0)
}

func (s *FoldContext) Surrogate_id_SYMB_14() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_14, 0)
}

func (s *FoldContext) Stellatype() IStellatypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStellatypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStellatypeContext)
}

func (s *FoldContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FoldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterFold(s)
	}
}

func (s *FoldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitFold(s)
	}
}

type AddContext struct {
	*ExprContext
}

func NewAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddContext {
	var p = new(AddContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AddContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddContext) Surrogate_id_SYMB_21() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_21, 0)
}

func (s *AddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterAdd(s)
	}
}

func (s *AddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitAdd(s)
	}
}

type IsZeroContext struct {
	*ExprContext
	n IExprContext
}

func NewIsZeroContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsZeroContext {
	var p = new(IsZeroContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IsZeroContext) GetN() IExprContext { return s.n }

func (s *IsZeroContext) SetN(v IExprContext) { s.n = v }

func (s *IsZeroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsZeroContext) Surrogate_id_SYMB_30() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_30, 0)
}

func (s *IsZeroContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *IsZeroContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *IsZeroContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IsZeroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterIsZero(s)
	}
}

func (s *IsZeroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitIsZero(s)
	}
}

type LessThanOrEqualContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewLessThanOrEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessThanOrEqualContext {
	var p = new(LessThanOrEqualContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LessThanOrEqualContext) GetLeft() IExprContext { return s.left }

func (s *LessThanOrEqualContext) GetRight() IExprContext { return s.right }

func (s *LessThanOrEqualContext) SetLeft(v IExprContext) { s.left = v }

func (s *LessThanOrEqualContext) SetRight(v IExprContext) { s.right = v }

func (s *LessThanOrEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanOrEqualContext) Surrogate_id_SYMB_16() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_16, 0)
}

func (s *LessThanOrEqualContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LessThanOrEqualContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LessThanOrEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterLessThanOrEqual(s)
	}
}

func (s *LessThanOrEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitLessThanOrEqual(s)
	}
}

type SuccContext struct {
	*ExprContext
	n IExprContext
}

func NewSuccContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SuccContext {
	var p = new(SuccContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SuccContext) GetN() IExprContext { return s.n }

func (s *SuccContext) SetN(v IExprContext) { s.n = v }

func (s *SuccContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuccContext) Surrogate_id_SYMB_57() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_57, 0)
}

func (s *SuccContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *SuccContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *SuccContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SuccContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterSucc(s)
	}
}

func (s *SuccContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitSucc(s)
	}
}

type VarContext struct {
	*ExprContext
	name antlr.Token
}

func NewVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarContext {
	var p = new(VarContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *VarContext) GetName() antlr.Token { return s.name }

func (s *VarContext) SetName(v antlr.Token) { s.name = v }

func (s *VarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarContext) StellaIdent() antlr.TerminalNode {
	return s.GetToken(StellaParserStellaIdent, 0)
}

func (s *VarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterVar(s)
	}
}

func (s *VarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitVar(s)
	}
}

type InlContext struct {
	*ExprContext
	expr_ IExprContext
}

func NewInlContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InlContext {
	var p = new(InlContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *InlContext) GetExpr_() IExprContext { return s.expr_ }

func (s *InlContext) SetExpr_(v IExprContext) { s.expr_ = v }

func (s *InlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlContext) Surrogate_id_SYMB_47() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_47, 0)
}

func (s *InlContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *InlContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *InlContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *InlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterInl(s)
	}
}

func (s *InlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitInl(s)
	}
}

type GreaterThanOrEqualContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewGreaterThanOrEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterThanOrEqualContext {
	var p = new(GreaterThanOrEqualContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GreaterThanOrEqualContext) GetLeft() IExprContext { return s.left }

func (s *GreaterThanOrEqualContext) GetRight() IExprContext { return s.right }

func (s *GreaterThanOrEqualContext) SetLeft(v IExprContext) { s.left = v }

func (s *GreaterThanOrEqualContext) SetRight(v IExprContext) { s.right = v }

func (s *GreaterThanOrEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanOrEqualContext) Surrogate_id_SYMB_18() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_18, 0)
}

func (s *GreaterThanOrEqualContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *GreaterThanOrEqualContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GreaterThanOrEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterGreaterThanOrEqual(s)
	}
}

func (s *GreaterThanOrEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitGreaterThanOrEqual(s)
	}
}

type InrContext struct {
	*ExprContext
	expr_ IExprContext
}

func NewInrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InrContext {
	var p = new(InrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *InrContext) GetExpr_() IExprContext { return s.expr_ }

func (s *InrContext) SetExpr_(v IExprContext) { s.expr_ = v }

func (s *InrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InrContext) Surrogate_id_SYMB_49() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_49, 0)
}

func (s *InrContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *InrContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *InrContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *InrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterInr(s)
	}
}

func (s *InrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitInr(s)
	}
}

type DivideContext struct {
	*ExprContext
}

func NewDivideContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivideContext {
	var p = new(DivideContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DivideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivideContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *DivideContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DivideContext) Surrogate_id_SYMB_24() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_24, 0)
}

func (s *DivideContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterDivide(s)
	}
}

func (s *DivideContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitDivide(s)
	}
}

type LessThanContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewLessThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessThanContext {
	var p = new(LessThanContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LessThanContext) GetLeft() IExprContext { return s.left }

func (s *LessThanContext) GetRight() IExprContext { return s.right }

func (s *LessThanContext) SetLeft(v IExprContext) { s.left = v }

func (s *LessThanContext) SetRight(v IExprContext) { s.right = v }

func (s *LessThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanContext) Surrogate_id_SYMB_15() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_15, 0)
}

func (s *LessThanContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LessThanContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LessThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterLessThan(s)
	}
}

func (s *LessThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitLessThan(s)
	}
}

type LogicNotContext struct {
	*ExprContext
	expr_ IExprContext
}

func NewLogicNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicNotContext {
	var p = new(LogicNotContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LogicNotContext) GetExpr_() IExprContext { return s.expr_ }

func (s *LogicNotContext) SetExpr_(v IExprContext) { s.expr_ = v }

func (s *LogicNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicNotContext) Surrogate_id_SYMB_54() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_54, 0)
}

func (s *LogicNotContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *LogicNotContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *LogicNotContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LogicNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterLogicNot(s)
	}
}

func (s *LogicNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitLogicNot(s)
	}
}

type DotRecordContext struct {
	*ExprContext
	expr_ IExprContext
	label antlr.Token
}

func NewDotRecordContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DotRecordContext {
	var p = new(DotRecordContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DotRecordContext) GetLabel() antlr.Token { return s.label }

func (s *DotRecordContext) SetLabel(v antlr.Token) { s.label = v }

func (s *DotRecordContext) GetExpr_() IExprContext { return s.expr_ }

func (s *DotRecordContext) SetExpr_(v IExprContext) { s.expr_ = v }

func (s *DotRecordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotRecordContext) Surrogate_id_SYMB_25() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_25, 0)
}

func (s *DotRecordContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DotRecordContext) StellaIdent() antlr.TerminalNode {
	return s.GetToken(StellaParserStellaIdent, 0)
}

func (s *DotRecordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterDotRecord(s)
	}
}

func (s *DotRecordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitDotRecord(s)
	}
}

type ParenthesisedExprContext struct {
	*ExprContext
}

func NewParenthesisedExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesisedExprContext {
	var p = new(ParenthesisedExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParenthesisedExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisedExprContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *ParenthesisedExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenthesisedExprContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *ParenthesisedExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterParenthesisedExpr(s)
	}
}

func (s *ParenthesisedExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitParenthesisedExpr(s)
	}
}

type GreaterThanContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewGreaterThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterThanContext {
	var p = new(GreaterThanContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GreaterThanContext) GetLeft() IExprContext { return s.left }

func (s *GreaterThanContext) GetRight() IExprContext { return s.right }

func (s *GreaterThanContext) SetLeft(v IExprContext) { s.left = v }

func (s *GreaterThanContext) SetRight(v IExprContext) { s.right = v }

func (s *GreaterThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanContext) Surrogate_id_SYMB_17() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_17, 0)
}

func (s *GreaterThanContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *GreaterThanContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GreaterThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterGreaterThan(s)
	}
}

func (s *GreaterThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitGreaterThan(s)
	}
}

type EqualContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualContext {
	var p = new(EqualContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EqualContext) GetLeft() IExprContext { return s.left }

func (s *EqualContext) GetRight() IExprContext { return s.right }

func (s *EqualContext) SetLeft(v IExprContext) { s.left = v }

func (s *EqualContext) SetRight(v IExprContext) { s.right = v }

func (s *EqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualContext) Surrogate_id_SYMB_19() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_19, 0)
}

func (s *EqualContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EqualContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterEqual(s)
	}
}

func (s *EqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitEqual(s)
	}
}

type TailContext struct {
	*ExprContext
	list IExprContext
}

func NewTailContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TailContext {
	var p = new(TailContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TailContext) GetList() IExprContext { return s.list }

func (s *TailContext) SetList(v IExprContext) { s.list = v }

func (s *TailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TailContext) Surrogate_id_SYMB_28() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_28, 0)
}

func (s *TailContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *TailContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *TailContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterTail(s)
	}
}

func (s *TailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitTail(s)
	}
}

type MultiplyContext struct {
	*ExprContext
}

func NewMultiplyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplyContext {
	var p = new(MultiplyContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MultiplyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MultiplyContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MultiplyContext) Surrogate_id_SYMB_23() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_23, 0)
}

func (s *MultiplyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterMultiply(s)
	}
}

func (s *MultiplyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitMultiply(s)
	}
}

type RecordContext struct {
	*ExprContext
	_binding IBindingContext
	bindings []IBindingContext
}

func NewRecordContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RecordContext {
	var p = new(RecordContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RecordContext) Get_binding() IBindingContext { return s._binding }

func (s *RecordContext) Set_binding(v IBindingContext) { s._binding = v }

func (s *RecordContext) GetBindings() []IBindingContext { return s.bindings }

func (s *RecordContext) SetBindings(v []IBindingContext) { s.bindings = v }

func (s *RecordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordContext) Surrogate_id_SYMB_4() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_4, 0)
}

func (s *RecordContext) Surrogate_id_SYMB_5() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_5, 0)
}

func (s *RecordContext) AllBinding() []IBindingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBindingContext); ok {
			len++
		}
	}

	tst := make([]IBindingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBindingContext); ok {
			tst[i] = t.(IBindingContext)
			i++
		}
	}

	return tst
}

func (s *RecordContext) Binding(i int) IBindingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindingContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindingContext)
}

func (s *RecordContext) AllSurrogate_id_SYMB_0() []antlr.TerminalNode {
	return s.GetTokens(StellaParserSurrogate_id_SYMB_0)
}

func (s *RecordContext) Surrogate_id_SYMB_0(i int) antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, i)
}

func (s *RecordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterRecord(s)
	}
}

func (s *RecordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitRecord(s)
	}
}

type ListContext struct {
	*ExprContext
	_expr IExprContext
	exprs []IExprContext
}

func NewListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListContext {
	var p = new(ListContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ListContext) Get_expr() IExprContext { return s._expr }

func (s *ListContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListContext) GetExprs() []IExprContext { return s.exprs }

func (s *ListContext) SetExprs(v []IExprContext) { s.exprs = v }

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) Surrogate_id_SYMB_13() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_13, 0)
}

func (s *ListContext) Surrogate_id_SYMB_14() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_14, 0)
}

func (s *ListContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListContext) Surrogate_id_SYMB_0() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, 0)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitList(s)
	}
}

type LogicAndContext struct {
	*ExprContext
}

func NewLogicAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicAndContext {
	var p = new(LogicAndContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LogicAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicAndContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LogicAndContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LogicAndContext) Surrogate_id_SYMB_35() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_35, 0)
}

func (s *LogicAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterLogicAnd(s)
	}
}

func (s *LogicAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitLogicAnd(s)
	}
}

type LogicOrContext struct {
	*ExprContext
}

func NewLogicOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicOrContext {
	var p = new(LogicOrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LogicOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicOrContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LogicOrContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LogicOrContext) Surrogate_id_SYMB_55() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_55, 0)
}

func (s *LogicOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterLogicOr(s)
	}
}

func (s *LogicOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitLogicOr(s)
	}
}

type HeadContext struct {
	*ExprContext
	list IExprContext
}

func NewHeadContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HeadContext {
	var p = new(HeadContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *HeadContext) GetList() IExprContext { return s.list }

func (s *HeadContext) SetList(v IExprContext) { s.list = v }

func (s *HeadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeadContext) Surrogate_id_SYMB_26() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_26, 0)
}

func (s *HeadContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *HeadContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *HeadContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *HeadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterHead(s)
	}
}

func (s *HeadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitHead(s)
	}
}

type TerminatingSemicolonContext struct {
	*ExprContext
}

func NewTerminatingSemicolonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TerminatingSemicolonContext {
	var p = new(TerminatingSemicolonContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TerminatingSemicolonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TerminatingSemicolonContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TerminatingSemicolonContext) Surrogate_id_SYMB_1() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_1, 0)
}

func (s *TerminatingSemicolonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterTerminatingSemicolon(s)
	}
}

func (s *TerminatingSemicolonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitTerminatingSemicolon(s)
	}
}

type NotEqualContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewNotEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotEqualContext {
	var p = new(NotEqualContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotEqualContext) GetLeft() IExprContext { return s.left }

func (s *NotEqualContext) GetRight() IExprContext { return s.right }

func (s *NotEqualContext) SetLeft(v IExprContext) { s.left = v }

func (s *NotEqualContext) SetRight(v IExprContext) { s.right = v }

func (s *NotEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotEqualContext) Surrogate_id_SYMB_20() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_20, 0)
}

func (s *NotEqualContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *NotEqualContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterNotEqual(s)
	}
}

func (s *NotEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitNotEqual(s)
	}
}

type ConstUnitContext struct {
	*ExprContext
}

func NewConstUnitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstUnitContext {
	var p = new(ConstUnitContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ConstUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstUnitContext) Surrogate_id_SYMB_63() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_63, 0)
}

func (s *ConstUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterConstUnit(s)
	}
}

func (s *ConstUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitConstUnit(s)
	}
}

type PredContext struct {
	*ExprContext
	n IExprContext
}

func NewPredContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PredContext {
	var p = new(PredContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PredContext) GetN() IExprContext { return s.n }

func (s *PredContext) SetN(v IExprContext) { s.n = v }

func (s *PredContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredContext) Surrogate_id_SYMB_29() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_29, 0)
}

func (s *PredContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *PredContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *PredContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PredContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterPred(s)
	}
}

func (s *PredContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitPred(s)
	}
}

type MatchContext struct {
	*ExprContext
	_match_case IMatch_caseContext
	cases       []IMatch_caseContext
}

func NewMatchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatchContext {
	var p = new(MatchContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MatchContext) Get_match_case() IMatch_caseContext { return s._match_case }

func (s *MatchContext) Set_match_case(v IMatch_caseContext) { s._match_case = v }

func (s *MatchContext) GetCases() []IMatch_caseContext { return s.cases }

func (s *MatchContext) SetCases(v []IMatch_caseContext) { s.cases = v }

func (s *MatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchContext) Surrogate_id_SYMB_53() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_53, 0)
}

func (s *MatchContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MatchContext) Surrogate_id_SYMB_4() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_4, 0)
}

func (s *MatchContext) Surrogate_id_SYMB_5() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_5, 0)
}

func (s *MatchContext) AllMatch_case() []IMatch_caseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMatch_caseContext); ok {
			len++
		}
	}

	tst := make([]IMatch_caseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMatch_caseContext); ok {
			tst[i] = t.(IMatch_caseContext)
			i++
		}
	}

	return tst
}

func (s *MatchContext) Match_case(i int) IMatch_caseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatch_caseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatch_caseContext)
}

func (s *MatchContext) AllSurrogate_id_SYMB_10() []antlr.TerminalNode {
	return s.GetTokens(StellaParserSurrogate_id_SYMB_10)
}

func (s *MatchContext) Surrogate_id_SYMB_10(i int) antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_10, i)
}

func (s *MatchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterMatch(s)
	}
}

func (s *MatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitMatch(s)
	}
}

type TypeAscContext struct {
	*ExprContext
	expr_ IExprContext
	type_ IStellatypeContext
}

func NewTypeAscContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeAscContext {
	var p = new(TypeAscContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TypeAscContext) GetExpr_() IExprContext { return s.expr_ }

func (s *TypeAscContext) GetType_() IStellatypeContext { return s.type_ }

func (s *TypeAscContext) SetExpr_(v IExprContext) { s.expr_ = v }

func (s *TypeAscContext) SetType_(v IStellatypeContext) { s.type_ = v }

func (s *TypeAscContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAscContext) Surrogate_id_SYMB_36() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_36, 0)
}

func (s *TypeAscContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TypeAscContext) Stellatype() IStellatypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStellatypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStellatypeContext)
}

func (s *TypeAscContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterTypeAsc(s)
	}
}

func (s *TypeAscContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitTypeAsc(s)
	}
}

type NatRecContext struct {
	*ExprContext
	n       IExprContext
	initial IExprContext
	step    IExprContext
}

func NewNatRecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NatRecContext {
	var p = new(NatRecContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NatRecContext) GetN() IExprContext { return s.n }

func (s *NatRecContext) GetInitial() IExprContext { return s.initial }

func (s *NatRecContext) GetStep() IExprContext { return s.step }

func (s *NatRecContext) SetN(v IExprContext) { s.n = v }

func (s *NatRecContext) SetInitial(v IExprContext) { s.initial = v }

func (s *NatRecContext) SetStep(v IExprContext) { s.step = v }

func (s *NatRecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NatRecContext) Surrogate_id_SYMB_31() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_31, 0)
}

func (s *NatRecContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *NatRecContext) AllSurrogate_id_SYMB_0() []antlr.TerminalNode {
	return s.GetTokens(StellaParserSurrogate_id_SYMB_0)
}

func (s *NatRecContext) Surrogate_id_SYMB_0(i int) antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, i)
}

func (s *NatRecContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *NatRecContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *NatRecContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NatRecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterNatRec(s)
	}
}

func (s *NatRecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitNatRec(s)
	}
}

type SequenceContext struct {
	*ExprContext
}

func NewSequenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SequenceContext {
	var p = new(SequenceContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequenceContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *SequenceContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SequenceContext) Surrogate_id_SYMB_1() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_1, 0)
}

func (s *SequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterSequence(s)
	}
}

func (s *SequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitSequence(s)
	}
}

type ConstFalseContext struct {
	*ExprContext
}

func NewConstFalseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstFalseContext {
	var p = new(ConstFalseContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ConstFalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstFalseContext) Surrogate_id_SYMB_41() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_41, 0)
}

func (s *ConstFalseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterConstFalse(s)
	}
}

func (s *ConstFalseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitConstFalse(s)
	}
}

type AbstractionContext struct {
	*ExprContext
	_paramDecl IParamDeclContext
	paramDecls []IParamDeclContext
	returnExpr IExprContext
}

func NewAbstractionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AbstractionContext {
	var p = new(AbstractionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AbstractionContext) Get_paramDecl() IParamDeclContext { return s._paramDecl }

func (s *AbstractionContext) GetReturnExpr() IExprContext { return s.returnExpr }

func (s *AbstractionContext) Set_paramDecl(v IParamDeclContext) { s._paramDecl = v }

func (s *AbstractionContext) SetReturnExpr(v IExprContext) { s.returnExpr = v }

func (s *AbstractionContext) GetParamDecls() []IParamDeclContext { return s.paramDecls }

func (s *AbstractionContext) SetParamDecls(v []IParamDeclContext) { s.paramDecls = v }

func (s *AbstractionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbstractionContext) Surrogate_id_SYMB_43() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_43, 0)
}

func (s *AbstractionContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *AbstractionContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *AbstractionContext) Surrogate_id_SYMB_4() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_4, 0)
}

func (s *AbstractionContext) Surrogate_id_SYMB_56() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_56, 0)
}

func (s *AbstractionContext) Surrogate_id_SYMB_5() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_5, 0)
}

func (s *AbstractionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AbstractionContext) AllParamDecl() []IParamDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamDeclContext); ok {
			len++
		}
	}

	tst := make([]IParamDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamDeclContext); ok {
			tst[i] = t.(IParamDeclContext)
			i++
		}
	}

	return tst
}

func (s *AbstractionContext) ParamDecl(i int) IParamDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamDeclContext)
}

func (s *AbstractionContext) AllSurrogate_id_SYMB_0() []antlr.TerminalNode {
	return s.GetTokens(StellaParserSurrogate_id_SYMB_0)
}

func (s *AbstractionContext) Surrogate_id_SYMB_0(i int) antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, i)
}

func (s *AbstractionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterAbstraction(s)
	}
}

func (s *AbstractionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitAbstraction(s)
	}
}

type ConstIntContext struct {
	*ExprContext
	n antlr.Token
}

func NewConstIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstIntContext {
	var p = new(ConstIntContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ConstIntContext) GetN() antlr.Token { return s.n }

func (s *ConstIntContext) SetN(v antlr.Token) { s.n = v }

func (s *ConstIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstIntContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(StellaParserINTEGER, 0)
}

func (s *ConstIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterConstInt(s)
	}
}

func (s *ConstIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitConstInt(s)
	}
}

type UnfoldContext struct {
	*ExprContext
	type_ IStellatypeContext
	expr_ IExprContext
}

func NewUnfoldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnfoldContext {
	var p = new(UnfoldContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *UnfoldContext) GetType_() IStellatypeContext { return s.type_ }

func (s *UnfoldContext) GetExpr_() IExprContext { return s.expr_ }

func (s *UnfoldContext) SetType_(v IStellatypeContext) { s.type_ = v }

func (s *UnfoldContext) SetExpr_(v IExprContext) { s.expr_ = v }

func (s *UnfoldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnfoldContext) Surrogate_id_SYMB_62() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_62, 0)
}

func (s *UnfoldContext) Surrogate_id_SYMB_13() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_13, 0)
}

func (s *UnfoldContext) Surrogate_id_SYMB_14() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_14, 0)
}

func (s *UnfoldContext) Stellatype() IStellatypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStellatypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStellatypeContext)
}

func (s *UnfoldContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnfoldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterUnfold(s)
	}
}

func (s *UnfoldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitUnfold(s)
	}
}

type VariantContext struct {
	*ExprContext
	label antlr.Token
	rhs   IExprContext
}

func NewVariantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariantContext {
	var p = new(VariantContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *VariantContext) GetLabel() antlr.Token { return s.label }

func (s *VariantContext) SetLabel(v antlr.Token) { s.label = v }

func (s *VariantContext) GetRhs() IExprContext { return s.rhs }

func (s *VariantContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *VariantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariantContext) Surrogate_id_SYMB_11() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_11, 0)
}

func (s *VariantContext) Surrogate_id_SYMB_12() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_12, 0)
}

func (s *VariantContext) StellaIdent() antlr.TerminalNode {
	return s.GetToken(StellaParserStellaIdent, 0)
}

func (s *VariantContext) Surrogate_id_SYMB_6() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_6, 0)
}

func (s *VariantContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VariantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterVariant(s)
	}
}

func (s *VariantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitVariant(s)
	}
}

type ConstTrueContext struct {
	*ExprContext
}

func NewConstTrueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstTrueContext {
	var p = new(ConstTrueContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ConstTrueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstTrueContext) Surrogate_id_SYMB_60() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_60, 0)
}

func (s *ConstTrueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterConstTrue(s)
	}
}

func (s *ConstTrueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitConstTrue(s)
	}
}

type DotTupleContext struct {
	*ExprContext
	expr_ IExprContext
	index antlr.Token
}

func NewDotTupleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DotTupleContext {
	var p = new(DotTupleContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DotTupleContext) GetIndex() antlr.Token { return s.index }

func (s *DotTupleContext) SetIndex(v antlr.Token) { s.index = v }

func (s *DotTupleContext) GetExpr_() IExprContext { return s.expr_ }

func (s *DotTupleContext) SetExpr_(v IExprContext) { s.expr_ = v }

func (s *DotTupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotTupleContext) Surrogate_id_SYMB_25() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_25, 0)
}

func (s *DotTupleContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DotTupleContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(StellaParserINTEGER, 0)
}

func (s *DotTupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterDotTuple(s)
	}
}

func (s *DotTupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitDotTuple(s)
	}
}

type FixContext struct {
	*ExprContext
	expr_ IExprContext
}

func NewFixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FixContext {
	var p = new(FixContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FixContext) GetExpr_() IExprContext { return s.expr_ }

func (s *FixContext) SetExpr_(v IExprContext) { s.expr_ = v }

func (s *FixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FixContext) Surrogate_id_SYMB_42() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_42, 0)
}

func (s *FixContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *FixContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *FixContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterFix(s)
	}
}

func (s *FixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitFix(s)
	}
}

type SubtractContext struct {
	*ExprContext
}

func NewSubtractContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubtractContext {
	var p = new(SubtractContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SubtractContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubtractContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *SubtractContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SubtractContext) Surrogate_id_SYMB_22() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_22, 0)
}

func (s *SubtractContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterSubtract(s)
	}
}

func (s *SubtractContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitSubtract(s)
	}
}

type LetContext struct {
	*ExprContext
	_patternBinding IPatternBindingContext
	patternBindings []IPatternBindingContext
	body            IExprContext
}

func NewLetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetContext {
	var p = new(LetContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LetContext) Get_patternBinding() IPatternBindingContext { return s._patternBinding }

func (s *LetContext) GetBody() IExprContext { return s.body }

func (s *LetContext) Set_patternBinding(v IPatternBindingContext) { s._patternBinding = v }

func (s *LetContext) SetBody(v IExprContext) { s.body = v }

func (s *LetContext) GetPatternBindings() []IPatternBindingContext { return s.patternBindings }

func (s *LetContext) SetPatternBindings(v []IPatternBindingContext) { s.patternBindings = v }

func (s *LetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetContext) Surrogate_id_SYMB_51() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_51, 0)
}

func (s *LetContext) Surrogate_id_SYMB_46() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_46, 0)
}

func (s *LetContext) AllPatternBinding() []IPatternBindingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPatternBindingContext); ok {
			len++
		}
	}

	tst := make([]IPatternBindingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPatternBindingContext); ok {
			tst[i] = t.(IPatternBindingContext)
			i++
		}
	}

	return tst
}

func (s *LetContext) PatternBinding(i int) IPatternBindingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternBindingContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternBindingContext)
}

func (s *LetContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LetContext) AllSurrogate_id_SYMB_0() []antlr.TerminalNode {
	return s.GetTokens(StellaParserSurrogate_id_SYMB_0)
}

func (s *LetContext) Surrogate_id_SYMB_0(i int) antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, i)
}

func (s *LetContext) Surrogate_id_SYMB_52() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_52, 0)
}

func (s *LetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterLet(s)
	}
}

func (s *LetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitLet(s)
	}
}

type IfContext struct {
	*ExprContext
	condition IExprContext
	thenExpr  IExprContext
	elseExpr  IExprContext
}

func NewIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfContext {
	var p = new(IfContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IfContext) GetCondition() IExprContext { return s.condition }

func (s *IfContext) GetThenExpr() IExprContext { return s.thenExpr }

func (s *IfContext) GetElseExpr() IExprContext { return s.elseExpr }

func (s *IfContext) SetCondition(v IExprContext) { s.condition = v }

func (s *IfContext) SetThenExpr(v IExprContext) { s.thenExpr = v }

func (s *IfContext) SetElseExpr(v IExprContext) { s.elseExpr = v }

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) Surrogate_id_SYMB_45() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_45, 0)
}

func (s *IfContext) Surrogate_id_SYMB_58() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_58, 0)
}

func (s *IfContext) Surrogate_id_SYMB_39() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_39, 0)
}

func (s *IfContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *IfContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitIf(s)
	}
}

type ApplicationContext struct {
	*ExprContext
	fun   IExprContext
	_expr IExprContext
	args  []IExprContext
}

func NewApplicationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ApplicationContext {
	var p = new(ApplicationContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ApplicationContext) GetFun() IExprContext { return s.fun }

func (s *ApplicationContext) Get_expr() IExprContext { return s._expr }

func (s *ApplicationContext) SetFun(v IExprContext) { s.fun = v }

func (s *ApplicationContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ApplicationContext) GetArgs() []IExprContext { return s.args }

func (s *ApplicationContext) SetArgs(v []IExprContext) { s.args = v }

func (s *ApplicationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApplicationContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *ApplicationContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *ApplicationContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ApplicationContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ApplicationContext) AllSurrogate_id_SYMB_0() []antlr.TerminalNode {
	return s.GetTokens(StellaParserSurrogate_id_SYMB_0)
}

func (s *ApplicationContext) Surrogate_id_SYMB_0(i int) antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, i)
}

func (s *ApplicationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterApplication(s)
	}
}

func (s *ApplicationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitApplication(s)
	}
}

type TupleContext struct {
	*ExprContext
	_expr IExprContext
	exprs []IExprContext
}

func NewTupleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleContext {
	var p = new(TupleContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TupleContext) Get_expr() IExprContext { return s._expr }

func (s *TupleContext) Set_expr(v IExprContext) { s._expr = v }

func (s *TupleContext) GetExprs() []IExprContext { return s.exprs }

func (s *TupleContext) SetExprs(v []IExprContext) { s.exprs = v }

func (s *TupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleContext) Surrogate_id_SYMB_4() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_4, 0)
}

func (s *TupleContext) Surrogate_id_SYMB_5() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_5, 0)
}

func (s *TupleContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *TupleContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TupleContext) AllSurrogate_id_SYMB_0() []antlr.TerminalNode {
	return s.GetTokens(StellaParserSurrogate_id_SYMB_0)
}

func (s *TupleContext) Surrogate_id_SYMB_0(i int) antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, i)
}

func (s *TupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterTuple(s)
	}
}

func (s *TupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitTuple(s)
	}
}

type ConsListContext struct {
	*ExprContext
	head IExprContext
	tail IExprContext
}

func NewConsListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConsListContext {
	var p = new(ConsListContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ConsListContext) GetHead() IExprContext { return s.head }

func (s *ConsListContext) GetTail() IExprContext { return s.tail }

func (s *ConsListContext) SetHead(v IExprContext) { s.head = v }

func (s *ConsListContext) SetTail(v IExprContext) { s.tail = v }

func (s *ConsListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsListContext) Surrogate_id_SYMB_37() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_37, 0)
}

func (s *ConsListContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *ConsListContext) Surrogate_id_SYMB_0() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, 0)
}

func (s *ConsListContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *ConsListContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ConsListContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConsListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterConsList(s)
	}
}

func (s *ConsListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitConsList(s)
	}
}

func (p *StellaParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *StellaParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, StellaParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		localctx = NewConstTrueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(133)
			p.Match(StellaParserSurrogate_id_SYMB_60)
		}

	case 2:
		localctx = NewConstFalseContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(134)
			p.Match(StellaParserSurrogate_id_SYMB_41)
		}

	case 3:
		localctx = NewConstUnitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(135)
			p.Match(StellaParserSurrogate_id_SYMB_63)
		}

	case 4:
		localctx = NewConstIntContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(136)

			var _m = p.Match(StellaParserINTEGER)

			localctx.(*ConstIntContext).n = _m
		}

	case 5:
		localctx = NewVarContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(137)

			var _m = p.Match(StellaParserStellaIdent)

			localctx.(*VarContext).name = _m
		}

	case 6:
		localctx = NewInlContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(138)
			p.Match(StellaParserSurrogate_id_SYMB_47)
		}
		{
			p.SetState(139)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(140)

			var _x = p.expr(0)

			localctx.(*InlContext).expr_ = _x
		}
		{
			p.SetState(141)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	case 7:
		localctx = NewInrContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(143)
			p.Match(StellaParserSurrogate_id_SYMB_49)
		}
		{
			p.SetState(144)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(145)

			var _x = p.expr(0)

			localctx.(*InrContext).expr_ = _x
		}
		{
			p.SetState(146)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	case 8:
		localctx = NewConsListContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(148)
			p.Match(StellaParserSurrogate_id_SYMB_37)
		}
		{
			p.SetState(149)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(150)

			var _x = p.expr(0)

			localctx.(*ConsListContext).head = _x
		}
		{
			p.SetState(151)
			p.Match(StellaParserSurrogate_id_SYMB_0)
		}
		{
			p.SetState(152)

			var _x = p.expr(0)

			localctx.(*ConsListContext).tail = _x
		}
		{
			p.SetState(153)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	case 9:
		localctx = NewHeadContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(155)
			p.Match(StellaParserSurrogate_id_SYMB_26)
		}
		{
			p.SetState(156)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(157)

			var _x = p.expr(0)

			localctx.(*HeadContext).list = _x
		}
		{
			p.SetState(158)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	case 10:
		localctx = NewIsEmptyContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(160)
			p.Match(StellaParserSurrogate_id_SYMB_27)
		}
		{
			p.SetState(161)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(162)

			var _x = p.expr(0)

			localctx.(*IsEmptyContext).list = _x
		}
		{
			p.SetState(163)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	case 11:
		localctx = NewTailContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(165)
			p.Match(StellaParserSurrogate_id_SYMB_28)
		}
		{
			p.SetState(166)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(167)

			var _x = p.expr(0)

			localctx.(*TailContext).list = _x
		}
		{
			p.SetState(168)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	case 12:
		localctx = NewSuccContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(170)
			p.Match(StellaParserSurrogate_id_SYMB_57)
		}
		{
			p.SetState(171)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(172)

			var _x = p.expr(0)

			localctx.(*SuccContext).n = _x
		}
		{
			p.SetState(173)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	case 13:
		localctx = NewLogicNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(175)
			p.Match(StellaParserSurrogate_id_SYMB_54)
		}
		{
			p.SetState(176)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(177)

			var _x = p.expr(0)

			localctx.(*LogicNotContext).expr_ = _x
		}
		{
			p.SetState(178)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	case 14:
		localctx = NewPredContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(180)
			p.Match(StellaParserSurrogate_id_SYMB_29)
		}
		{
			p.SetState(181)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(182)

			var _x = p.expr(0)

			localctx.(*PredContext).n = _x
		}
		{
			p.SetState(183)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	case 15:
		localctx = NewIsZeroContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(185)
			p.Match(StellaParserSurrogate_id_SYMB_30)
		}
		{
			p.SetState(186)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(187)

			var _x = p.expr(0)

			localctx.(*IsZeroContext).n = _x
		}
		{
			p.SetState(188)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	case 16:
		localctx = NewFixContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(190)
			p.Match(StellaParserSurrogate_id_SYMB_42)
		}
		{
			p.SetState(191)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(192)

			var _x = p.expr(0)

			localctx.(*FixContext).expr_ = _x
		}
		{
			p.SetState(193)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	case 17:
		localctx = NewNatRecContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(195)
			p.Match(StellaParserSurrogate_id_SYMB_31)
		}
		{
			p.SetState(196)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(197)

			var _x = p.expr(0)

			localctx.(*NatRecContext).n = _x
		}
		{
			p.SetState(198)
			p.Match(StellaParserSurrogate_id_SYMB_0)
		}
		{
			p.SetState(199)

			var _x = p.expr(0)

			localctx.(*NatRecContext).initial = _x
		}
		{
			p.SetState(200)
			p.Match(StellaParserSurrogate_id_SYMB_0)
		}
		{
			p.SetState(201)

			var _x = p.expr(0)

			localctx.(*NatRecContext).step = _x
		}
		{
			p.SetState(202)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	case 18:
		localctx = NewFoldContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(204)
			p.Match(StellaParserSurrogate_id_SYMB_44)
		}
		{
			p.SetState(205)
			p.Match(StellaParserSurrogate_id_SYMB_13)
		}
		{
			p.SetState(206)

			var _x = p.stellatype(0)

			localctx.(*FoldContext).type_ = _x
		}
		{
			p.SetState(207)
			p.Match(StellaParserSurrogate_id_SYMB_14)
		}
		{
			p.SetState(208)

			var _x = p.expr(28)

			localctx.(*FoldContext).expr_ = _x
		}

	case 19:
		localctx = NewUnfoldContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(210)
			p.Match(StellaParserSurrogate_id_SYMB_62)
		}
		{
			p.SetState(211)
			p.Match(StellaParserSurrogate_id_SYMB_13)
		}
		{
			p.SetState(212)

			var _x = p.stellatype(0)

			localctx.(*UnfoldContext).type_ = _x
		}
		{
			p.SetState(213)
			p.Match(StellaParserSurrogate_id_SYMB_14)
		}
		{
			p.SetState(214)

			var _x = p.expr(27)

			localctx.(*UnfoldContext).expr_ = _x
		}

	case 20:
		localctx = NewAbstractionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(216)
			p.Match(StellaParserSurrogate_id_SYMB_43)
		}
		{
			p.SetState(217)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == StellaParserStellaIdent {
			{
				p.SetState(218)

				var _x = p.ParamDecl()

				localctx.(*AbstractionContext)._paramDecl = _x
			}
			localctx.(*AbstractionContext).paramDecls = append(localctx.(*AbstractionContext).paramDecls, localctx.(*AbstractionContext)._paramDecl)
			p.SetState(223)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == StellaParserSurrogate_id_SYMB_0 {
				{
					p.SetState(219)
					p.Match(StellaParserSurrogate_id_SYMB_0)
				}
				{
					p.SetState(220)

					var _x = p.ParamDecl()

					localctx.(*AbstractionContext)._paramDecl = _x
				}
				localctx.(*AbstractionContext).paramDecls = append(localctx.(*AbstractionContext).paramDecls, localctx.(*AbstractionContext)._paramDecl)

				p.SetState(225)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(228)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}
		{
			p.SetState(229)
			p.Match(StellaParserSurrogate_id_SYMB_4)
		}
		{
			p.SetState(230)
			p.Match(StellaParserSurrogate_id_SYMB_56)
		}
		{
			p.SetState(231)

			var _x = p.expr(0)

			localctx.(*AbstractionContext).returnExpr = _x
		}
		{
			p.SetState(232)
			p.Match(StellaParserSurrogate_id_SYMB_5)
		}

	case 21:
		localctx = NewTupleContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(234)
			p.Match(StellaParserSurrogate_id_SYMB_4)
		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-6560200659419770840) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&161) != 0 {
			{
				p.SetState(235)

				var _x = p.expr(0)

				localctx.(*TupleContext)._expr = _x
			}
			localctx.(*TupleContext).exprs = append(localctx.(*TupleContext).exprs, localctx.(*TupleContext)._expr)
			p.SetState(240)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == StellaParserSurrogate_id_SYMB_0 {
				{
					p.SetState(236)
					p.Match(StellaParserSurrogate_id_SYMB_0)
				}
				{
					p.SetState(237)

					var _x = p.expr(0)

					localctx.(*TupleContext)._expr = _x
				}
				localctx.(*TupleContext).exprs = append(localctx.(*TupleContext).exprs, localctx.(*TupleContext)._expr)

				p.SetState(242)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(245)
			p.Match(StellaParserSurrogate_id_SYMB_5)
		}

	case 22:
		localctx = NewRecordContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(246)
			p.Match(StellaParserSurrogate_id_SYMB_4)
		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == StellaParserStellaIdent {
			{
				p.SetState(247)

				var _x = p.Binding()

				localctx.(*RecordContext)._binding = _x
			}
			localctx.(*RecordContext).bindings = append(localctx.(*RecordContext).bindings, localctx.(*RecordContext)._binding)
			p.SetState(252)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == StellaParserSurrogate_id_SYMB_0 {
				{
					p.SetState(248)
					p.Match(StellaParserSurrogate_id_SYMB_0)
				}
				{
					p.SetState(249)

					var _x = p.Binding()

					localctx.(*RecordContext)._binding = _x
				}
				localctx.(*RecordContext).bindings = append(localctx.(*RecordContext).bindings, localctx.(*RecordContext)._binding)

				p.SetState(254)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(257)
			p.Match(StellaParserSurrogate_id_SYMB_5)
		}

	case 23:
		localctx = NewVariantContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(258)
			p.Match(StellaParserSurrogate_id_SYMB_11)
		}
		{
			p.SetState(259)

			var _m = p.Match(StellaParserStellaIdent)

			localctx.(*VariantContext).label = _m
		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == StellaParserSurrogate_id_SYMB_6 {
			{
				p.SetState(260)
				p.Match(StellaParserSurrogate_id_SYMB_6)
			}
			{
				p.SetState(261)

				var _x = p.expr(0)

				localctx.(*VariantContext).rhs = _x
			}

		}
		{
			p.SetState(264)
			p.Match(StellaParserSurrogate_id_SYMB_12)
		}

	case 24:
		localctx = NewMatchContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(265)
			p.Match(StellaParserSurrogate_id_SYMB_53)
		}
		{
			p.SetState(266)
			p.expr(0)
		}
		{
			p.SetState(267)
			p.Match(StellaParserSurrogate_id_SYMB_4)
		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2595485433173397544) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&161) != 0 {
			{
				p.SetState(268)

				var _x = p.Match_case()

				localctx.(*MatchContext)._match_case = _x
			}
			localctx.(*MatchContext).cases = append(localctx.(*MatchContext).cases, localctx.(*MatchContext)._match_case)
			p.SetState(273)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == StellaParserSurrogate_id_SYMB_10 {
				{
					p.SetState(269)
					p.Match(StellaParserSurrogate_id_SYMB_10)
				}
				{
					p.SetState(270)

					var _x = p.Match_case()

					localctx.(*MatchContext)._match_case = _x
				}
				localctx.(*MatchContext).cases = append(localctx.(*MatchContext).cases, localctx.(*MatchContext)._match_case)

				p.SetState(275)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(278)
			p.Match(StellaParserSurrogate_id_SYMB_5)
		}

	case 25:
		localctx = NewListContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(280)
			p.Match(StellaParserSurrogate_id_SYMB_13)
		}
		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-6560200659419770840) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&161) != 0 {
			{
				p.SetState(281)

				var _x = p.expr(0)

				localctx.(*ListContext)._expr = _x
			}
			localctx.(*ListContext).exprs = append(localctx.(*ListContext).exprs, localctx.(*ListContext)._expr)

			{
				p.SetState(282)
				p.Match(StellaParserSurrogate_id_SYMB_0)
			}
			{
				p.SetState(283)

				var _x = p.expr(0)

				localctx.(*ListContext)._expr = _x
			}
			localctx.(*ListContext).exprs = append(localctx.(*ListContext).exprs, localctx.(*ListContext)._expr)

		}
		{
			p.SetState(287)
			p.Match(StellaParserSurrogate_id_SYMB_14)
		}

	case 26:
		localctx = NewIfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(288)
			p.Match(StellaParserSurrogate_id_SYMB_45)
		}
		{
			p.SetState(289)

			var _x = p.expr(0)

			localctx.(*IfContext).condition = _x
		}
		{
			p.SetState(290)
			p.Match(StellaParserSurrogate_id_SYMB_58)
		}
		{
			p.SetState(291)

			var _x = p.expr(0)

			localctx.(*IfContext).thenExpr = _x
		}
		{
			p.SetState(292)
			p.Match(StellaParserSurrogate_id_SYMB_39)
		}
		{
			p.SetState(293)

			var _x = p.expr(6)

			localctx.(*IfContext).elseExpr = _x
		}

	case 27:
		localctx = NewLetContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(295)
			p.Match(StellaParserSurrogate_id_SYMB_51)
		}
		{
			p.SetState(296)

			var _x = p.PatternBinding()

			localctx.(*LetContext)._patternBinding = _x
		}
		localctx.(*LetContext).patternBindings = append(localctx.(*LetContext).patternBindings, localctx.(*LetContext)._patternBinding)
		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StellaParserSurrogate_id_SYMB_0 {
			{
				p.SetState(297)
				p.Match(StellaParserSurrogate_id_SYMB_0)
			}
			{
				p.SetState(298)

				var _x = p.PatternBinding()

				localctx.(*LetContext)._patternBinding = _x
			}
			localctx.(*LetContext).patternBindings = append(localctx.(*LetContext).patternBindings, localctx.(*LetContext)._patternBinding)

			p.SetState(303)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(304)
			p.Match(StellaParserSurrogate_id_SYMB_46)
		}
		{
			p.SetState(305)

			var _x = p.expr(5)

			localctx.(*LetContext).body = _x
		}

	case 28:
		localctx = NewLetContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(307)
			p.Match(StellaParserSurrogate_id_SYMB_52)
		}
		{
			p.SetState(308)

			var _x = p.PatternBinding()

			localctx.(*LetContext)._patternBinding = _x
		}
		localctx.(*LetContext).patternBindings = append(localctx.(*LetContext).patternBindings, localctx.(*LetContext)._patternBinding)
		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StellaParserSurrogate_id_SYMB_0 {
			{
				p.SetState(309)
				p.Match(StellaParserSurrogate_id_SYMB_0)
			}
			{
				p.SetState(310)

				var _x = p.PatternBinding()

				localctx.(*LetContext)._patternBinding = _x
			}
			localctx.(*LetContext).patternBindings = append(localctx.(*LetContext).patternBindings, localctx.(*LetContext)._patternBinding)

			p.SetState(315)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(316)
			p.Match(StellaParserSurrogate_id_SYMB_46)
		}
		{
			p.SetState(317)

			var _x = p.expr(4)

			localctx.(*LetContext).body = _x
		}

	case 29:
		localctx = NewParenthesisedExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(319)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(320)
			p.expr(0)
		}
		{
			p.SetState(321)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(388)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplyContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_expr)
				p.SetState(325)

				if !(p.Precpred(p.GetParserRuleContext(), 25)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 25)", ""))
				}
				{
					p.SetState(326)
					p.Match(StellaParserSurrogate_id_SYMB_23)
				}
				{
					p.SetState(327)
					p.expr(26)
				}

			case 2:
				localctx = NewDivideContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_expr)
				p.SetState(328)

				if !(p.Precpred(p.GetParserRuleContext(), 24)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 24)", ""))
				}
				{
					p.SetState(329)
					p.Match(StellaParserSurrogate_id_SYMB_24)
				}
				{
					p.SetState(330)
					p.expr(25)
				}

			case 3:
				localctx = NewLogicAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_expr)
				p.SetState(331)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
				}
				{
					p.SetState(332)
					p.Match(StellaParserSurrogate_id_SYMB_35)
				}
				{
					p.SetState(333)
					p.expr(24)
				}

			case 4:
				localctx = NewAddContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_expr)
				p.SetState(334)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
				}
				{
					p.SetState(335)
					p.Match(StellaParserSurrogate_id_SYMB_21)
				}
				{
					p.SetState(336)
					p.expr(23)
				}

			case 5:
				localctx = NewSubtractContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_expr)
				p.SetState(337)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
				}
				{
					p.SetState(338)
					p.Match(StellaParserSurrogate_id_SYMB_22)
				}
				{
					p.SetState(339)
					p.expr(22)
				}

			case 6:
				localctx = NewLogicOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_expr)
				p.SetState(340)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}
				{
					p.SetState(341)
					p.Match(StellaParserSurrogate_id_SYMB_55)
				}
				{
					p.SetState(342)
					p.expr(21)
				}

			case 7:
				localctx = NewLessThanContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LessThanContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_expr)
				p.SetState(343)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(344)
					p.Match(StellaParserSurrogate_id_SYMB_15)
				}
				{
					p.SetState(345)

					var _x = p.expr(13)

					localctx.(*LessThanContext).right = _x
				}

			case 8:
				localctx = NewLessThanOrEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LessThanOrEqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_expr)
				p.SetState(346)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(347)
					p.Match(StellaParserSurrogate_id_SYMB_16)
				}
				{
					p.SetState(348)

					var _x = p.expr(12)

					localctx.(*LessThanOrEqualContext).right = _x
				}

			case 9:
				localctx = NewGreaterThanContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*GreaterThanContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_expr)
				p.SetState(349)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(350)
					p.Match(StellaParserSurrogate_id_SYMB_17)
				}
				{
					p.SetState(351)

					var _x = p.expr(11)

					localctx.(*GreaterThanContext).right = _x
				}

			case 10:
				localctx = NewGreaterThanOrEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*GreaterThanOrEqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_expr)
				p.SetState(352)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(353)
					p.Match(StellaParserSurrogate_id_SYMB_18)
				}
				{
					p.SetState(354)

					var _x = p.expr(10)

					localctx.(*GreaterThanOrEqualContext).right = _x
				}

			case 11:
				localctx = NewEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*EqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_expr)
				p.SetState(355)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(356)
					p.Match(StellaParserSurrogate_id_SYMB_19)
				}
				{
					p.SetState(357)

					var _x = p.expr(9)

					localctx.(*EqualContext).right = _x
				}

			case 12:
				localctx = NewNotEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*NotEqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_expr)
				p.SetState(358)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(359)
					p.Match(StellaParserSurrogate_id_SYMB_20)
				}
				{
					p.SetState(360)

					var _x = p.expr(8)

					localctx.(*NotEqualContext).right = _x
				}

			case 13:
				localctx = NewSequenceContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_expr)
				p.SetState(361)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(362)
					p.Match(StellaParserSurrogate_id_SYMB_1)
				}
				{
					p.SetState(363)
					p.expr(3)
				}

			case 14:
				localctx = NewDotRecordContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*DotRecordContext).expr_ = _prevctx

				p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_expr)
				p.SetState(364)

				if !(p.Precpred(p.GetParserRuleContext(), 47)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 47)", ""))
				}
				{
					p.SetState(365)
					p.Match(StellaParserSurrogate_id_SYMB_25)
				}
				{
					p.SetState(366)

					var _m = p.Match(StellaParserStellaIdent)

					localctx.(*DotRecordContext).label = _m
				}

			case 15:
				localctx = NewDotTupleContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*DotTupleContext).expr_ = _prevctx

				p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_expr)
				p.SetState(367)

				if !(p.Precpred(p.GetParserRuleContext(), 46)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 46)", ""))
				}
				{
					p.SetState(368)
					p.Match(StellaParserSurrogate_id_SYMB_25)
				}
				{
					p.SetState(369)

					var _m = p.Match(StellaParserINTEGER)

					localctx.(*DotTupleContext).index = _m
				}

			case 16:
				localctx = NewApplicationContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ApplicationContext).fun = _prevctx

				p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_expr)
				p.SetState(370)

				if !(p.Precpred(p.GetParserRuleContext(), 26)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 26)", ""))
				}
				{
					p.SetState(371)
					p.Match(StellaParserSurrogate_id_SYMB_2)
				}
				p.SetState(380)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-6560200659419770840) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&161) != 0 {
					{
						p.SetState(372)

						var _x = p.expr(0)

						localctx.(*ApplicationContext)._expr = _x
					}
					localctx.(*ApplicationContext).args = append(localctx.(*ApplicationContext).args, localctx.(*ApplicationContext)._expr)
					p.SetState(377)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == StellaParserSurrogate_id_SYMB_0 {
						{
							p.SetState(373)
							p.Match(StellaParserSurrogate_id_SYMB_0)
						}
						{
							p.SetState(374)

							var _x = p.expr(0)

							localctx.(*ApplicationContext)._expr = _x
						}
						localctx.(*ApplicationContext).args = append(localctx.(*ApplicationContext).args, localctx.(*ApplicationContext)._expr)

						p.SetState(379)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(382)
					p.Match(StellaParserSurrogate_id_SYMB_3)
				}

			case 17:
				localctx = NewTypeAscContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*TypeAscContext).expr_ = _prevctx

				p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_expr)
				p.SetState(383)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(384)
					p.Match(StellaParserSurrogate_id_SYMB_36)
				}
				{
					p.SetState(385)

					var _x = p.stellatype(0)

					localctx.(*TypeAscContext).type_ = _x
				}

			case 18:
				localctx = NewTerminatingSemicolonContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_expr)
				p.SetState(386)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(387)
					p.Match(StellaParserSurrogate_id_SYMB_1)
				}

			}

		}
		p.SetState(392)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}

	return localctx
}

// IPatternBindingContext is an interface to support dynamic dispatch.
type IPatternBindingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPat returns the pat rule contexts.
	GetPat() IPatternContext

	// GetRhs returns the rhs rule contexts.
	GetRhs() IExprContext

	// SetPat sets the pat rule contexts.
	SetPat(IPatternContext)

	// SetRhs sets the rhs rule contexts.
	SetRhs(IExprContext)

	// IsPatternBindingContext differentiates from other interfaces.
	IsPatternBindingContext()
}

type PatternBindingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	pat    IPatternContext
	rhs    IExprContext
}

func NewEmptyPatternBindingContext() *PatternBindingContext {
	var p = new(PatternBindingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StellaParserRULE_patternBinding
	return p
}

func (*PatternBindingContext) IsPatternBindingContext() {}

func NewPatternBindingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternBindingContext {
	var p = new(PatternBindingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StellaParserRULE_patternBinding

	return p
}

func (s *PatternBindingContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternBindingContext) GetPat() IPatternContext { return s.pat }

func (s *PatternBindingContext) GetRhs() IExprContext { return s.rhs }

func (s *PatternBindingContext) SetPat(v IPatternContext) { s.pat = v }

func (s *PatternBindingContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *PatternBindingContext) Surrogate_id_SYMB_6() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_6, 0)
}

func (s *PatternBindingContext) Pattern() IPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *PatternBindingContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PatternBindingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternBindingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatternBindingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterPatternBinding(s)
	}
}

func (s *PatternBindingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitPatternBinding(s)
	}
}

func (p *StellaParser) PatternBinding() (localctx IPatternBindingContext) {
	this := p
	_ = this

	localctx = NewPatternBindingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, StellaParserRULE_patternBinding)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(393)

		var _x = p.Pattern()

		localctx.(*PatternBindingContext).pat = _x
	}
	{
		p.SetState(394)
		p.Match(StellaParserSurrogate_id_SYMB_6)
	}
	{
		p.SetState(395)

		var _x = p.expr(0)

		localctx.(*PatternBindingContext).rhs = _x
	}

	return localctx
}

// IBindingContext is an interface to support dynamic dispatch.
type IBindingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetRhs returns the rhs rule contexts.
	GetRhs() IExprContext

	// SetRhs sets the rhs rule contexts.
	SetRhs(IExprContext)

	// IsBindingContext differentiates from other interfaces.
	IsBindingContext()
}

type BindingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	rhs    IExprContext
}

func NewEmptyBindingContext() *BindingContext {
	var p = new(BindingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StellaParserRULE_binding
	return p
}

func (*BindingContext) IsBindingContext() {}

func NewBindingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindingContext {
	var p = new(BindingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StellaParserRULE_binding

	return p
}

func (s *BindingContext) GetParser() antlr.Parser { return s.parser }

func (s *BindingContext) GetName() antlr.Token { return s.name }

func (s *BindingContext) SetName(v antlr.Token) { s.name = v }

func (s *BindingContext) GetRhs() IExprContext { return s.rhs }

func (s *BindingContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *BindingContext) Surrogate_id_SYMB_6() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_6, 0)
}

func (s *BindingContext) StellaIdent() antlr.TerminalNode {
	return s.GetToken(StellaParserStellaIdent, 0)
}

func (s *BindingContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BindingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterBinding(s)
	}
}

func (s *BindingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitBinding(s)
	}
}

func (p *StellaParser) Binding() (localctx IBindingContext) {
	this := p
	_ = this

	localctx = NewBindingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, StellaParserRULE_binding)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)

		var _m = p.Match(StellaParserStellaIdent)

		localctx.(*BindingContext).name = _m
	}
	{
		p.SetState(398)
		p.Match(StellaParserSurrogate_id_SYMB_6)
	}
	{
		p.SetState(399)

		var _x = p.expr(0)

		localctx.(*BindingContext).rhs = _x
	}

	return localctx
}

// IMatch_caseContext is an interface to support dynamic dispatch.
type IMatch_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMatch_caseContext differentiates from other interfaces.
	IsMatch_caseContext()
}

type Match_caseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatch_caseContext() *Match_caseContext {
	var p = new(Match_caseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StellaParserRULE_match_case
	return p
}

func (*Match_caseContext) IsMatch_caseContext() {}

func NewMatch_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Match_caseContext {
	var p = new(Match_caseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StellaParserRULE_match_case

	return p
}

func (s *Match_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Match_caseContext) Pattern() IPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *Match_caseContext) Surrogate_id_SYMB_9() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_9, 0)
}

func (s *Match_caseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Match_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Match_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Match_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterMatch_case(s)
	}
}

func (s *Match_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitMatch_case(s)
	}
}

func (p *StellaParser) Match_case() (localctx IMatch_caseContext) {
	this := p
	_ = this

	localctx = NewMatch_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, StellaParserRULE_match_case)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(401)
		p.Pattern()
	}
	{
		p.SetState(402)
		p.Match(StellaParserSurrogate_id_SYMB_9)
	}
	{
		p.SetState(403)
		p.expr(0)
	}

	return localctx
}

// IPatternContext is an interface to support dynamic dispatch.
type IPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPatternContext differentiates from other interfaces.
	IsPatternContext()
}

type PatternContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternContext() *PatternContext {
	var p = new(PatternContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StellaParserRULE_pattern
	return p
}

func (*PatternContext) IsPatternContext() {}

func NewPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternContext {
	var p = new(PatternContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StellaParserRULE_pattern

	return p
}

func (s *PatternContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternContext) CopyFrom(ctx *PatternContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PatternConsContext struct {
	*PatternContext
	head IPatternContext
	tail IPatternContext
}

func NewPatternConsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatternConsContext {
	var p = new(PatternConsContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *PatternConsContext) GetHead() IPatternContext { return s.head }

func (s *PatternConsContext) GetTail() IPatternContext { return s.tail }

func (s *PatternConsContext) SetHead(v IPatternContext) { s.head = v }

func (s *PatternConsContext) SetTail(v IPatternContext) { s.tail = v }

func (s *PatternConsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternConsContext) Surrogate_id_SYMB_37() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_37, 0)
}

func (s *PatternConsContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *PatternConsContext) Surrogate_id_SYMB_0() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, 0)
}

func (s *PatternConsContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *PatternConsContext) AllPattern() []IPatternContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPatternContext); ok {
			len++
		}
	}

	tst := make([]IPatternContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPatternContext); ok {
			tst[i] = t.(IPatternContext)
			i++
		}
	}

	return tst
}

func (s *PatternConsContext) Pattern(i int) IPatternContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *PatternConsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterPatternCons(s)
	}
}

func (s *PatternConsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitPatternCons(s)
	}
}

type PatternTupleContext struct {
	*PatternContext
	_pattern IPatternContext
	patterns []IPatternContext
}

func NewPatternTupleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatternTupleContext {
	var p = new(PatternTupleContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *PatternTupleContext) Get_pattern() IPatternContext { return s._pattern }

func (s *PatternTupleContext) Set_pattern(v IPatternContext) { s._pattern = v }

func (s *PatternTupleContext) GetPatterns() []IPatternContext { return s.patterns }

func (s *PatternTupleContext) SetPatterns(v []IPatternContext) { s.patterns = v }

func (s *PatternTupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternTupleContext) Surrogate_id_SYMB_4() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_4, 0)
}

func (s *PatternTupleContext) Surrogate_id_SYMB_5() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_5, 0)
}

func (s *PatternTupleContext) AllPattern() []IPatternContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPatternContext); ok {
			len++
		}
	}

	tst := make([]IPatternContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPatternContext); ok {
			tst[i] = t.(IPatternContext)
			i++
		}
	}

	return tst
}

func (s *PatternTupleContext) Pattern(i int) IPatternContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *PatternTupleContext) AllSurrogate_id_SYMB_0() []antlr.TerminalNode {
	return s.GetTokens(StellaParserSurrogate_id_SYMB_0)
}

func (s *PatternTupleContext) Surrogate_id_SYMB_0(i int) antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, i)
}

func (s *PatternTupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterPatternTuple(s)
	}
}

func (s *PatternTupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitPatternTuple(s)
	}
}

type PatternListContext struct {
	*PatternContext
	_pattern IPatternContext
	patterns []IPatternContext
}

func NewPatternListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatternListContext {
	var p = new(PatternListContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *PatternListContext) Get_pattern() IPatternContext { return s._pattern }

func (s *PatternListContext) Set_pattern(v IPatternContext) { s._pattern = v }

func (s *PatternListContext) GetPatterns() []IPatternContext { return s.patterns }

func (s *PatternListContext) SetPatterns(v []IPatternContext) { s.patterns = v }

func (s *PatternListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternListContext) Surrogate_id_SYMB_13() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_13, 0)
}

func (s *PatternListContext) Surrogate_id_SYMB_14() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_14, 0)
}

func (s *PatternListContext) AllPattern() []IPatternContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPatternContext); ok {
			len++
		}
	}

	tst := make([]IPatternContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPatternContext); ok {
			tst[i] = t.(IPatternContext)
			i++
		}
	}

	return tst
}

func (s *PatternListContext) Pattern(i int) IPatternContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *PatternListContext) AllSurrogate_id_SYMB_0() []antlr.TerminalNode {
	return s.GetTokens(StellaParserSurrogate_id_SYMB_0)
}

func (s *PatternListContext) Surrogate_id_SYMB_0(i int) antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, i)
}

func (s *PatternListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterPatternList(s)
	}
}

func (s *PatternListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitPatternList(s)
	}
}

type PatternRecordContext struct {
	*PatternContext
	_labelledPattern ILabelledPatternContext
	patterns         []ILabelledPatternContext
}

func NewPatternRecordContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatternRecordContext {
	var p = new(PatternRecordContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *PatternRecordContext) Get_labelledPattern() ILabelledPatternContext {
	return s._labelledPattern
}

func (s *PatternRecordContext) Set_labelledPattern(v ILabelledPatternContext) { s._labelledPattern = v }

func (s *PatternRecordContext) GetPatterns() []ILabelledPatternContext { return s.patterns }

func (s *PatternRecordContext) SetPatterns(v []ILabelledPatternContext) { s.patterns = v }

func (s *PatternRecordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternRecordContext) Surrogate_id_SYMB_4() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_4, 0)
}

func (s *PatternRecordContext) Surrogate_id_SYMB_5() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_5, 0)
}

func (s *PatternRecordContext) AllLabelledPattern() []ILabelledPatternContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILabelledPatternContext); ok {
			len++
		}
	}

	tst := make([]ILabelledPatternContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILabelledPatternContext); ok {
			tst[i] = t.(ILabelledPatternContext)
			i++
		}
	}

	return tst
}

func (s *PatternRecordContext) LabelledPattern(i int) ILabelledPatternContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelledPatternContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelledPatternContext)
}

func (s *PatternRecordContext) AllSurrogate_id_SYMB_0() []antlr.TerminalNode {
	return s.GetTokens(StellaParserSurrogate_id_SYMB_0)
}

func (s *PatternRecordContext) Surrogate_id_SYMB_0(i int) antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, i)
}

func (s *PatternRecordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterPatternRecord(s)
	}
}

func (s *PatternRecordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitPatternRecord(s)
	}
}

type PatternVariantContext struct {
	*PatternContext
	label antlr.Token
}

func NewPatternVariantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatternVariantContext {
	var p = new(PatternVariantContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *PatternVariantContext) GetLabel() antlr.Token { return s.label }

func (s *PatternVariantContext) SetLabel(v antlr.Token) { s.label = v }

func (s *PatternVariantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternVariantContext) Surrogate_id_SYMB_11() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_11, 0)
}

func (s *PatternVariantContext) Surrogate_id_SYMB_12() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_12, 0)
}

func (s *PatternVariantContext) StellaIdent() antlr.TerminalNode {
	return s.GetToken(StellaParserStellaIdent, 0)
}

func (s *PatternVariantContext) Surrogate_id_SYMB_6() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_6, 0)
}

func (s *PatternVariantContext) Pattern() IPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *PatternVariantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterPatternVariant(s)
	}
}

func (s *PatternVariantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitPatternVariant(s)
	}
}

type PatternIntContext struct {
	*PatternContext
	n antlr.Token
}

func NewPatternIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatternIntContext {
	var p = new(PatternIntContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *PatternIntContext) GetN() antlr.Token { return s.n }

func (s *PatternIntContext) SetN(v antlr.Token) { s.n = v }

func (s *PatternIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternIntContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(StellaParserINTEGER, 0)
}

func (s *PatternIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterPatternInt(s)
	}
}

func (s *PatternIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitPatternInt(s)
	}
}

type PatternInrContext struct {
	*PatternContext
	pat IPatternContext
}

func NewPatternInrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatternInrContext {
	var p = new(PatternInrContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *PatternInrContext) GetPat() IPatternContext { return s.pat }

func (s *PatternInrContext) SetPat(v IPatternContext) { s.pat = v }

func (s *PatternInrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternInrContext) Surrogate_id_SYMB_49() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_49, 0)
}

func (s *PatternInrContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *PatternInrContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *PatternInrContext) Pattern() IPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *PatternInrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterPatternInr(s)
	}
}

func (s *PatternInrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitPatternInr(s)
	}
}

type PatternTrueContext struct {
	*PatternContext
}

func NewPatternTrueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatternTrueContext {
	var p = new(PatternTrueContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *PatternTrueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternTrueContext) Surrogate_id_SYMB_60() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_60, 0)
}

func (s *PatternTrueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterPatternTrue(s)
	}
}

func (s *PatternTrueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitPatternTrue(s)
	}
}

type PatternInlContext struct {
	*PatternContext
	pat IPatternContext
}

func NewPatternInlContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatternInlContext {
	var p = new(PatternInlContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *PatternInlContext) GetPat() IPatternContext { return s.pat }

func (s *PatternInlContext) SetPat(v IPatternContext) { s.pat = v }

func (s *PatternInlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternInlContext) Surrogate_id_SYMB_47() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_47, 0)
}

func (s *PatternInlContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *PatternInlContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *PatternInlContext) Pattern() IPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *PatternInlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterPatternInl(s)
	}
}

func (s *PatternInlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitPatternInl(s)
	}
}

type PatternVarContext struct {
	*PatternContext
	name antlr.Token
}

func NewPatternVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatternVarContext {
	var p = new(PatternVarContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *PatternVarContext) GetName() antlr.Token { return s.name }

func (s *PatternVarContext) SetName(v antlr.Token) { s.name = v }

func (s *PatternVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternVarContext) StellaIdent() antlr.TerminalNode {
	return s.GetToken(StellaParserStellaIdent, 0)
}

func (s *PatternVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterPatternVar(s)
	}
}

func (s *PatternVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitPatternVar(s)
	}
}

type ParenthesisedPatternContext struct {
	*PatternContext
}

func NewParenthesisedPatternContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesisedPatternContext {
	var p = new(ParenthesisedPatternContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *ParenthesisedPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisedPatternContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *ParenthesisedPatternContext) Pattern() IPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *ParenthesisedPatternContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *ParenthesisedPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterParenthesisedPattern(s)
	}
}

func (s *ParenthesisedPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitParenthesisedPattern(s)
	}
}

type PatternSuccContext struct {
	*PatternContext
	n IPatternContext
}

func NewPatternSuccContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatternSuccContext {
	var p = new(PatternSuccContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *PatternSuccContext) GetN() IPatternContext { return s.n }

func (s *PatternSuccContext) SetN(v IPatternContext) { s.n = v }

func (s *PatternSuccContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternSuccContext) Surrogate_id_SYMB_57() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_57, 0)
}

func (s *PatternSuccContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *PatternSuccContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *PatternSuccContext) Pattern() IPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *PatternSuccContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterPatternSucc(s)
	}
}

func (s *PatternSuccContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitPatternSucc(s)
	}
}

type PatternFalseContext struct {
	*PatternContext
}

func NewPatternFalseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatternFalseContext {
	var p = new(PatternFalseContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *PatternFalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternFalseContext) Surrogate_id_SYMB_41() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_41, 0)
}

func (s *PatternFalseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterPatternFalse(s)
	}
}

func (s *PatternFalseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitPatternFalse(s)
	}
}

type PatternUnitContext struct {
	*PatternContext
}

func NewPatternUnitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatternUnitContext {
	var p = new(PatternUnitContext)

	p.PatternContext = NewEmptyPatternContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PatternContext))

	return p
}

func (s *PatternUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternUnitContext) Surrogate_id_SYMB_63() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_63, 0)
}

func (s *PatternUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterPatternUnit(s)
	}
}

func (s *PatternUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitPatternUnit(s)
	}
}

func (p *StellaParser) Pattern() (localctx IPatternContext) {
	this := p
	_ = this

	localctx = NewPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, StellaParserRULE_pattern)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPatternVariantContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(405)
			p.Match(StellaParserSurrogate_id_SYMB_11)
		}
		{
			p.SetState(406)

			var _m = p.Match(StellaParserStellaIdent)

			localctx.(*PatternVariantContext).label = _m
		}
		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == StellaParserSurrogate_id_SYMB_6 {
			{
				p.SetState(407)
				p.Match(StellaParserSurrogate_id_SYMB_6)
			}
			{
				p.SetState(408)
				p.Pattern()
			}

		}
		{
			p.SetState(411)
			p.Match(StellaParserSurrogate_id_SYMB_12)
		}

	case 2:
		localctx = NewPatternInlContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(412)
			p.Match(StellaParserSurrogate_id_SYMB_47)
		}
		{
			p.SetState(413)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(414)

			var _x = p.Pattern()

			localctx.(*PatternInlContext).pat = _x
		}
		{
			p.SetState(415)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	case 3:
		localctx = NewPatternInrContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(417)
			p.Match(StellaParserSurrogate_id_SYMB_49)
		}
		{
			p.SetState(418)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(419)

			var _x = p.Pattern()

			localctx.(*PatternInrContext).pat = _x
		}
		{
			p.SetState(420)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	case 4:
		localctx = NewPatternTupleContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(422)
			p.Match(StellaParserSurrogate_id_SYMB_4)
		}
		p.SetState(431)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2595485433173397544) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&161) != 0 {
			{
				p.SetState(423)

				var _x = p.Pattern()

				localctx.(*PatternTupleContext)._pattern = _x
			}
			localctx.(*PatternTupleContext).patterns = append(localctx.(*PatternTupleContext).patterns, localctx.(*PatternTupleContext)._pattern)
			p.SetState(428)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == StellaParserSurrogate_id_SYMB_0 {
				{
					p.SetState(424)
					p.Match(StellaParserSurrogate_id_SYMB_0)
				}
				{
					p.SetState(425)

					var _x = p.Pattern()

					localctx.(*PatternTupleContext)._pattern = _x
				}
				localctx.(*PatternTupleContext).patterns = append(localctx.(*PatternTupleContext).patterns, localctx.(*PatternTupleContext)._pattern)

				p.SetState(430)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(433)
			p.Match(StellaParserSurrogate_id_SYMB_5)
		}

	case 5:
		localctx = NewPatternRecordContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(434)
			p.Match(StellaParserSurrogate_id_SYMB_4)
		}
		p.SetState(443)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == StellaParserStellaIdent {
			{
				p.SetState(435)

				var _x = p.LabelledPattern()

				localctx.(*PatternRecordContext)._labelledPattern = _x
			}
			localctx.(*PatternRecordContext).patterns = append(localctx.(*PatternRecordContext).patterns, localctx.(*PatternRecordContext)._labelledPattern)
			p.SetState(440)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == StellaParserSurrogate_id_SYMB_0 {
				{
					p.SetState(436)
					p.Match(StellaParserSurrogate_id_SYMB_0)
				}
				{
					p.SetState(437)

					var _x = p.LabelledPattern()

					localctx.(*PatternRecordContext)._labelledPattern = _x
				}
				localctx.(*PatternRecordContext).patterns = append(localctx.(*PatternRecordContext).patterns, localctx.(*PatternRecordContext)._labelledPattern)

				p.SetState(442)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(445)
			p.Match(StellaParserSurrogate_id_SYMB_5)
		}

	case 6:
		localctx = NewPatternListContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(446)
			p.Match(StellaParserSurrogate_id_SYMB_13)
		}
		p.SetState(455)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2595485433173397544) != 0 || (int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&161) != 0 {
			{
				p.SetState(447)

				var _x = p.Pattern()

				localctx.(*PatternListContext)._pattern = _x
			}
			localctx.(*PatternListContext).patterns = append(localctx.(*PatternListContext).patterns, localctx.(*PatternListContext)._pattern)
			p.SetState(452)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == StellaParserSurrogate_id_SYMB_0 {
				{
					p.SetState(448)
					p.Match(StellaParserSurrogate_id_SYMB_0)
				}
				{
					p.SetState(449)

					var _x = p.Pattern()

					localctx.(*PatternListContext)._pattern = _x
				}
				localctx.(*PatternListContext).patterns = append(localctx.(*PatternListContext).patterns, localctx.(*PatternListContext)._pattern)

				p.SetState(454)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(457)
			p.Match(StellaParserSurrogate_id_SYMB_14)
		}

	case 7:
		localctx = NewPatternConsContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(458)
			p.Match(StellaParserSurrogate_id_SYMB_37)
		}
		{
			p.SetState(459)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(460)

			var _x = p.Pattern()

			localctx.(*PatternConsContext).head = _x
		}
		{
			p.SetState(461)
			p.Match(StellaParserSurrogate_id_SYMB_0)
		}
		{
			p.SetState(462)

			var _x = p.Pattern()

			localctx.(*PatternConsContext).tail = _x
		}
		{
			p.SetState(463)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	case 8:
		localctx = NewPatternFalseContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(465)
			p.Match(StellaParserSurrogate_id_SYMB_41)
		}

	case 9:
		localctx = NewPatternTrueContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(466)
			p.Match(StellaParserSurrogate_id_SYMB_60)
		}

	case 10:
		localctx = NewPatternUnitContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(467)
			p.Match(StellaParserSurrogate_id_SYMB_63)
		}

	case 11:
		localctx = NewPatternIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(468)

			var _m = p.Match(StellaParserINTEGER)

			localctx.(*PatternIntContext).n = _m
		}

	case 12:
		localctx = NewPatternSuccContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(469)
			p.Match(StellaParserSurrogate_id_SYMB_57)
		}
		{
			p.SetState(470)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(471)

			var _x = p.Pattern()

			localctx.(*PatternSuccContext).n = _x
		}
		{
			p.SetState(472)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	case 13:
		localctx = NewPatternVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(474)

			var _m = p.Match(StellaParserStellaIdent)

			localctx.(*PatternVarContext).name = _m
		}

	case 14:
		localctx = NewParenthesisedPatternContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(475)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(476)
			p.Pattern()
		}
		{
			p.SetState(477)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	}

	return localctx
}

// ILabelledPatternContext is an interface to support dynamic dispatch.
type ILabelledPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLabel returns the label token.
	GetLabel() antlr.Token

	// SetLabel sets the label token.
	SetLabel(antlr.Token)

	// IsLabelledPatternContext differentiates from other interfaces.
	IsLabelledPatternContext()
}

type LabelledPatternContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	label  antlr.Token
}

func NewEmptyLabelledPatternContext() *LabelledPatternContext {
	var p = new(LabelledPatternContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StellaParserRULE_labelledPattern
	return p
}

func (*LabelledPatternContext) IsLabelledPatternContext() {}

func NewLabelledPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelledPatternContext {
	var p = new(LabelledPatternContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StellaParserRULE_labelledPattern

	return p
}

func (s *LabelledPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelledPatternContext) GetLabel() antlr.Token { return s.label }

func (s *LabelledPatternContext) SetLabel(v antlr.Token) { s.label = v }

func (s *LabelledPatternContext) Surrogate_id_SYMB_6() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_6, 0)
}

func (s *LabelledPatternContext) Pattern() IPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *LabelledPatternContext) StellaIdent() antlr.TerminalNode {
	return s.GetToken(StellaParserStellaIdent, 0)
}

func (s *LabelledPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelledPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelledPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterLabelledPattern(s)
	}
}

func (s *LabelledPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitLabelledPattern(s)
	}
}

func (p *StellaParser) LabelledPattern() (localctx ILabelledPatternContext) {
	this := p
	_ = this

	localctx = NewLabelledPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, StellaParserRULE_labelledPattern)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(481)

		var _m = p.Match(StellaParserStellaIdent)

		localctx.(*LabelledPatternContext).label = _m
	}
	{
		p.SetState(482)
		p.Match(StellaParserSurrogate_id_SYMB_6)
	}
	{
		p.SetState(483)
		p.Pattern()
	}

	return localctx
}

// IStellatypeContext is an interface to support dynamic dispatch.
type IStellatypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStellatypeContext differentiates from other interfaces.
	IsStellatypeContext()
}

type StellatypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStellatypeContext() *StellatypeContext {
	var p = new(StellatypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StellaParserRULE_stellatype
	return p
}

func (*StellatypeContext) IsStellatypeContext() {}

func NewStellatypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StellatypeContext {
	var p = new(StellatypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StellaParserRULE_stellatype

	return p
}

func (s *StellatypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StellatypeContext) CopyFrom(ctx *StellatypeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StellatypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StellatypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeTupleContext struct {
	*StellatypeContext
	_stellatype IStellatypeContext
	types       []IStellatypeContext
}

func NewTypeTupleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeTupleContext {
	var p = new(TypeTupleContext)

	p.StellatypeContext = NewEmptyStellatypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StellatypeContext))

	return p
}

func (s *TypeTupleContext) Get_stellatype() IStellatypeContext { return s._stellatype }

func (s *TypeTupleContext) Set_stellatype(v IStellatypeContext) { s._stellatype = v }

func (s *TypeTupleContext) GetTypes() []IStellatypeContext { return s.types }

func (s *TypeTupleContext) SetTypes(v []IStellatypeContext) { s.types = v }

func (s *TypeTupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeTupleContext) Surrogate_id_SYMB_4() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_4, 0)
}

func (s *TypeTupleContext) Surrogate_id_SYMB_5() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_5, 0)
}

func (s *TypeTupleContext) AllStellatype() []IStellatypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStellatypeContext); ok {
			len++
		}
	}

	tst := make([]IStellatypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStellatypeContext); ok {
			tst[i] = t.(IStellatypeContext)
			i++
		}
	}

	return tst
}

func (s *TypeTupleContext) Stellatype(i int) IStellatypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStellatypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStellatypeContext)
}

func (s *TypeTupleContext) Surrogate_id_SYMB_0() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, 0)
}

func (s *TypeTupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterTypeTuple(s)
	}
}

func (s *TypeTupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitTypeTuple(s)
	}
}

type TypeVarContext struct {
	*StellatypeContext
	name antlr.Token
}

func NewTypeVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeVarContext {
	var p = new(TypeVarContext)

	p.StellatypeContext = NewEmptyStellatypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StellatypeContext))

	return p
}

func (s *TypeVarContext) GetName() antlr.Token { return s.name }

func (s *TypeVarContext) SetName(v antlr.Token) { s.name = v }

func (s *TypeVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeVarContext) StellaIdent() antlr.TerminalNode {
	return s.GetToken(StellaParserStellaIdent, 0)
}

func (s *TypeVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterTypeVar(s)
	}
}

func (s *TypeVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitTypeVar(s)
	}
}

type TypeVariantContext struct {
	*StellatypeContext
	_variantFieldType IVariantFieldTypeContext
	fieldTypes        []IVariantFieldTypeContext
}

func NewTypeVariantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeVariantContext {
	var p = new(TypeVariantContext)

	p.StellatypeContext = NewEmptyStellatypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StellatypeContext))

	return p
}

func (s *TypeVariantContext) Get_variantFieldType() IVariantFieldTypeContext {
	return s._variantFieldType
}

func (s *TypeVariantContext) Set_variantFieldType(v IVariantFieldTypeContext) {
	s._variantFieldType = v
}

func (s *TypeVariantContext) GetFieldTypes() []IVariantFieldTypeContext { return s.fieldTypes }

func (s *TypeVariantContext) SetFieldTypes(v []IVariantFieldTypeContext) { s.fieldTypes = v }

func (s *TypeVariantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeVariantContext) Surrogate_id_SYMB_4() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_4, 0)
}

func (s *TypeVariantContext) Surrogate_id_SYMB_5() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_5, 0)
}

func (s *TypeVariantContext) AllVariantFieldType() []IVariantFieldTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariantFieldTypeContext); ok {
			len++
		}
	}

	tst := make([]IVariantFieldTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariantFieldTypeContext); ok {
			tst[i] = t.(IVariantFieldTypeContext)
			i++
		}
	}

	return tst
}

func (s *TypeVariantContext) VariantFieldType(i int) IVariantFieldTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariantFieldTypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariantFieldTypeContext)
}

func (s *TypeVariantContext) AllSurrogate_id_SYMB_0() []antlr.TerminalNode {
	return s.GetTokens(StellaParserSurrogate_id_SYMB_0)
}

func (s *TypeVariantContext) Surrogate_id_SYMB_0(i int) antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, i)
}

func (s *TypeVariantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterTypeVariant(s)
	}
}

func (s *TypeVariantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitTypeVariant(s)
	}
}

type TypeUnitContext struct {
	*StellatypeContext
}

func NewTypeUnitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeUnitContext {
	var p = new(TypeUnitContext)

	p.StellatypeContext = NewEmptyStellatypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StellatypeContext))

	return p
}

func (s *TypeUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeUnitContext) Surrogate_id_SYMB_34() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_34, 0)
}

func (s *TypeUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterTypeUnit(s)
	}
}

func (s *TypeUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitTypeUnit(s)
	}
}

type TypeBoolContext struct {
	*StellatypeContext
}

func NewTypeBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeBoolContext {
	var p = new(TypeBoolContext)

	p.StellatypeContext = NewEmptyStellatypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StellatypeContext))

	return p
}

func (s *TypeBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeBoolContext) Surrogate_id_SYMB_32() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_32, 0)
}

func (s *TypeBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterTypeBool(s)
	}
}

func (s *TypeBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitTypeBool(s)
	}
}

type TypeNatContext struct {
	*StellatypeContext
}

func NewTypeNatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeNatContext {
	var p = new(TypeNatContext)

	p.StellatypeContext = NewEmptyStellatypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StellatypeContext))

	return p
}

func (s *TypeNatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNatContext) Surrogate_id_SYMB_33() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_33, 0)
}

func (s *TypeNatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterTypeNat(s)
	}
}

func (s *TypeNatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitTypeNat(s)
	}
}

type TypeRecContext struct {
	*StellatypeContext
	var_  antlr.Token
	type_ IStellatypeContext
}

func NewTypeRecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeRecContext {
	var p = new(TypeRecContext)

	p.StellatypeContext = NewEmptyStellatypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StellatypeContext))

	return p
}

func (s *TypeRecContext) GetVar_() antlr.Token { return s.var_ }

func (s *TypeRecContext) SetVar_(v antlr.Token) { s.var_ = v }

func (s *TypeRecContext) GetType_() IStellatypeContext { return s.type_ }

func (s *TypeRecContext) SetType_(v IStellatypeContext) { s.type_ = v }

func (s *TypeRecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRecContext) Surrogate_id_SYMB_65() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_65, 0)
}

func (s *TypeRecContext) Surrogate_id_SYMB_25() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_25, 0)
}

func (s *TypeRecContext) StellaIdent() antlr.TerminalNode {
	return s.GetToken(StellaParserStellaIdent, 0)
}

func (s *TypeRecContext) Stellatype() IStellatypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStellatypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStellatypeContext)
}

func (s *TypeRecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterTypeRec(s)
	}
}

func (s *TypeRecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitTypeRec(s)
	}
}

type TypeParensContext struct {
	*StellatypeContext
}

func NewTypeParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeParensContext {
	var p = new(TypeParensContext)

	p.StellatypeContext = NewEmptyStellatypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StellatypeContext))

	return p
}

func (s *TypeParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParensContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *TypeParensContext) Stellatype() IStellatypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStellatypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStellatypeContext)
}

func (s *TypeParensContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *TypeParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterTypeParens(s)
	}
}

func (s *TypeParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitTypeParens(s)
	}
}

type TypeFunContext struct {
	*StellatypeContext
	_stellatype IStellatypeContext
	paramTypes  []IStellatypeContext
	returnType  IStellatypeContext
}

func NewTypeFunContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeFunContext {
	var p = new(TypeFunContext)

	p.StellatypeContext = NewEmptyStellatypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StellatypeContext))

	return p
}

func (s *TypeFunContext) Get_stellatype() IStellatypeContext { return s._stellatype }

func (s *TypeFunContext) GetReturnType() IStellatypeContext { return s.returnType }

func (s *TypeFunContext) Set_stellatype(v IStellatypeContext) { s._stellatype = v }

func (s *TypeFunContext) SetReturnType(v IStellatypeContext) { s.returnType = v }

func (s *TypeFunContext) GetParamTypes() []IStellatypeContext { return s.paramTypes }

func (s *TypeFunContext) SetParamTypes(v []IStellatypeContext) { s.paramTypes = v }

func (s *TypeFunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeFunContext) Surrogate_id_SYMB_43() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_43, 0)
}

func (s *TypeFunContext) Surrogate_id_SYMB_2() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_2, 0)
}

func (s *TypeFunContext) Surrogate_id_SYMB_3() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_3, 0)
}

func (s *TypeFunContext) Surrogate_id_SYMB_8() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_8, 0)
}

func (s *TypeFunContext) AllStellatype() []IStellatypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStellatypeContext); ok {
			len++
		}
	}

	tst := make([]IStellatypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStellatypeContext); ok {
			tst[i] = t.(IStellatypeContext)
			i++
		}
	}

	return tst
}

func (s *TypeFunContext) Stellatype(i int) IStellatypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStellatypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStellatypeContext)
}

func (s *TypeFunContext) AllSurrogate_id_SYMB_0() []antlr.TerminalNode {
	return s.GetTokens(StellaParserSurrogate_id_SYMB_0)
}

func (s *TypeFunContext) Surrogate_id_SYMB_0(i int) antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, i)
}

func (s *TypeFunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterTypeFun(s)
	}
}

func (s *TypeFunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitTypeFun(s)
	}
}

type TypeRecordContext struct {
	*StellatypeContext
	_recordFieldType IRecordFieldTypeContext
	fieldTypes       []IRecordFieldTypeContext
}

func NewTypeRecordContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeRecordContext {
	var p = new(TypeRecordContext)

	p.StellatypeContext = NewEmptyStellatypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StellatypeContext))

	return p
}

func (s *TypeRecordContext) Get_recordFieldType() IRecordFieldTypeContext { return s._recordFieldType }

func (s *TypeRecordContext) Set_recordFieldType(v IRecordFieldTypeContext) { s._recordFieldType = v }

func (s *TypeRecordContext) GetFieldTypes() []IRecordFieldTypeContext { return s.fieldTypes }

func (s *TypeRecordContext) SetFieldTypes(v []IRecordFieldTypeContext) { s.fieldTypes = v }

func (s *TypeRecordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRecordContext) Surrogate_id_SYMB_4() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_4, 0)
}

func (s *TypeRecordContext) Surrogate_id_SYMB_5() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_5, 0)
}

func (s *TypeRecordContext) AllRecordFieldType() []IRecordFieldTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRecordFieldTypeContext); ok {
			len++
		}
	}

	tst := make([]IRecordFieldTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRecordFieldTypeContext); ok {
			tst[i] = t.(IRecordFieldTypeContext)
			i++
		}
	}

	return tst
}

func (s *TypeRecordContext) RecordFieldType(i int) IRecordFieldTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecordFieldTypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRecordFieldTypeContext)
}

func (s *TypeRecordContext) AllSurrogate_id_SYMB_0() []antlr.TerminalNode {
	return s.GetTokens(StellaParserSurrogate_id_SYMB_0)
}

func (s *TypeRecordContext) Surrogate_id_SYMB_0(i int) antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, i)
}

func (s *TypeRecordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterTypeRecord(s)
	}
}

func (s *TypeRecordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitTypeRecord(s)
	}
}

type TypeListContext struct {
	*StellatypeContext
	_stellatype IStellatypeContext
	types       []IStellatypeContext
}

func NewTypeListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeListContext {
	var p = new(TypeListContext)

	p.StellatypeContext = NewEmptyStellatypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StellatypeContext))

	return p
}

func (s *TypeListContext) Get_stellatype() IStellatypeContext { return s._stellatype }

func (s *TypeListContext) Set_stellatype(v IStellatypeContext) { s._stellatype = v }

func (s *TypeListContext) GetTypes() []IStellatypeContext { return s.types }

func (s *TypeListContext) SetTypes(v []IStellatypeContext) { s.types = v }

func (s *TypeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeListContext) Surrogate_id_SYMB_13() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_13, 0)
}

func (s *TypeListContext) Surrogate_id_SYMB_14() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_14, 0)
}

func (s *TypeListContext) AllStellatype() []IStellatypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStellatypeContext); ok {
			len++
		}
	}

	tst := make([]IStellatypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStellatypeContext); ok {
			tst[i] = t.(IStellatypeContext)
			i++
		}
	}

	return tst
}

func (s *TypeListContext) Stellatype(i int) IStellatypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStellatypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStellatypeContext)
}

func (s *TypeListContext) AllSurrogate_id_SYMB_0() []antlr.TerminalNode {
	return s.GetTokens(StellaParserSurrogate_id_SYMB_0)
}

func (s *TypeListContext) Surrogate_id_SYMB_0(i int) antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_0, i)
}

func (s *TypeListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterTypeList(s)
	}
}

func (s *TypeListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitTypeList(s)
	}
}

type TypeSumContext struct {
	*StellatypeContext
	left  IStellatypeContext
	right IStellatypeContext
}

func NewTypeSumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeSumContext {
	var p = new(TypeSumContext)

	p.StellatypeContext = NewEmptyStellatypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StellatypeContext))

	return p
}

func (s *TypeSumContext) GetLeft() IStellatypeContext { return s.left }

func (s *TypeSumContext) GetRight() IStellatypeContext { return s.right }

func (s *TypeSumContext) SetLeft(v IStellatypeContext) { s.left = v }

func (s *TypeSumContext) SetRight(v IStellatypeContext) { s.right = v }

func (s *TypeSumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSumContext) Surrogate_id_SYMB_21() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_21, 0)
}

func (s *TypeSumContext) AllStellatype() []IStellatypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStellatypeContext); ok {
			len++
		}
	}

	tst := make([]IStellatypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStellatypeContext); ok {
			tst[i] = t.(IStellatypeContext)
			i++
		}
	}

	return tst
}

func (s *TypeSumContext) Stellatype(i int) IStellatypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStellatypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStellatypeContext)
}

func (s *TypeSumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterTypeSum(s)
	}
}

func (s *TypeSumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitTypeSum(s)
	}
}

func (p *StellaParser) Stellatype() (localctx IStellatypeContext) {
	return p.stellatype(0)
}

func (p *StellaParser) stellatype(_p int) (localctx IStellatypeContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewStellatypeContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IStellatypeContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, StellaParserRULE_stellatype, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(557)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeBoolContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(486)
			p.Match(StellaParserSurrogate_id_SYMB_32)
		}

	case 2:
		localctx = NewTypeNatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(487)
			p.Match(StellaParserSurrogate_id_SYMB_33)
		}

	case 3:
		localctx = NewTypeFunContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(488)
			p.Match(StellaParserSurrogate_id_SYMB_43)
		}
		{
			p.SetState(489)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		p.SetState(498)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17652315602984) != 0 || _la == StellaParserSurrogate_id_SYMB_65 || _la == StellaParserStellaIdent {
			{
				p.SetState(490)

				var _x = p.stellatype(0)

				localctx.(*TypeFunContext)._stellatype = _x
			}
			localctx.(*TypeFunContext).paramTypes = append(localctx.(*TypeFunContext).paramTypes, localctx.(*TypeFunContext)._stellatype)
			p.SetState(495)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == StellaParserSurrogate_id_SYMB_0 {
				{
					p.SetState(491)
					p.Match(StellaParserSurrogate_id_SYMB_0)
				}
				{
					p.SetState(492)

					var _x = p.stellatype(0)

					localctx.(*TypeFunContext)._stellatype = _x
				}
				localctx.(*TypeFunContext).paramTypes = append(localctx.(*TypeFunContext).paramTypes, localctx.(*TypeFunContext)._stellatype)

				p.SetState(497)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(500)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}
		{
			p.SetState(501)
			p.Match(StellaParserSurrogate_id_SYMB_8)
		}
		{
			p.SetState(502)

			var _x = p.stellatype(10)

			localctx.(*TypeFunContext).returnType = _x
		}

	case 4:
		localctx = NewTypeRecContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(503)
			p.Match(StellaParserSurrogate_id_SYMB_65)
		}
		{
			p.SetState(504)

			var _m = p.Match(StellaParserStellaIdent)

			localctx.(*TypeRecContext).var_ = _m
		}
		{
			p.SetState(505)
			p.Match(StellaParserSurrogate_id_SYMB_25)
		}
		{
			p.SetState(506)

			var _x = p.stellatype(9)

			localctx.(*TypeRecContext).type_ = _x
		}

	case 5:
		localctx = NewTypeTupleContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(507)
			p.Match(StellaParserSurrogate_id_SYMB_4)
		}
		p.SetState(512)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17652315602984) != 0 || _la == StellaParserSurrogate_id_SYMB_65 || _la == StellaParserStellaIdent {
			{
				p.SetState(508)

				var _x = p.stellatype(0)

				localctx.(*TypeTupleContext)._stellatype = _x
			}
			localctx.(*TypeTupleContext).types = append(localctx.(*TypeTupleContext).types, localctx.(*TypeTupleContext)._stellatype)

			{
				p.SetState(509)
				p.Match(StellaParserSurrogate_id_SYMB_0)
			}
			{
				p.SetState(510)

				var _x = p.stellatype(0)

				localctx.(*TypeTupleContext)._stellatype = _x
			}
			localctx.(*TypeTupleContext).types = append(localctx.(*TypeTupleContext).types, localctx.(*TypeTupleContext)._stellatype)

		}
		{
			p.SetState(514)
			p.Match(StellaParserSurrogate_id_SYMB_5)
		}

	case 6:
		localctx = NewTypeRecordContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(515)
			p.Match(StellaParserSurrogate_id_SYMB_4)
		}
		p.SetState(524)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == StellaParserStellaIdent {
			{
				p.SetState(516)

				var _x = p.RecordFieldType()

				localctx.(*TypeRecordContext)._recordFieldType = _x
			}
			localctx.(*TypeRecordContext).fieldTypes = append(localctx.(*TypeRecordContext).fieldTypes, localctx.(*TypeRecordContext)._recordFieldType)
			p.SetState(521)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == StellaParserSurrogate_id_SYMB_0 {
				{
					p.SetState(517)
					p.Match(StellaParserSurrogate_id_SYMB_0)
				}
				{
					p.SetState(518)

					var _x = p.RecordFieldType()

					localctx.(*TypeRecordContext)._recordFieldType = _x
				}
				localctx.(*TypeRecordContext).fieldTypes = append(localctx.(*TypeRecordContext).fieldTypes, localctx.(*TypeRecordContext)._recordFieldType)

				p.SetState(523)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(526)
			p.Match(StellaParserSurrogate_id_SYMB_5)
		}

	case 7:
		localctx = NewTypeVariantContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(527)
			p.Match(StellaParserSurrogate_id_SYMB_4)
		}
		p.SetState(536)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == StellaParserStellaIdent {
			{
				p.SetState(528)

				var _x = p.VariantFieldType()

				localctx.(*TypeVariantContext)._variantFieldType = _x
			}
			localctx.(*TypeVariantContext).fieldTypes = append(localctx.(*TypeVariantContext).fieldTypes, localctx.(*TypeVariantContext)._variantFieldType)
			p.SetState(533)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == StellaParserSurrogate_id_SYMB_0 {
				{
					p.SetState(529)
					p.Match(StellaParserSurrogate_id_SYMB_0)
				}
				{
					p.SetState(530)

					var _x = p.VariantFieldType()

					localctx.(*TypeVariantContext)._variantFieldType = _x
				}
				localctx.(*TypeVariantContext).fieldTypes = append(localctx.(*TypeVariantContext).fieldTypes, localctx.(*TypeVariantContext)._variantFieldType)

				p.SetState(535)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(538)
			p.Match(StellaParserSurrogate_id_SYMB_5)
		}

	case 8:
		localctx = NewTypeListContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(539)
			p.Match(StellaParserSurrogate_id_SYMB_13)
		}
		p.SetState(548)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17652315602984) != 0 || _la == StellaParserSurrogate_id_SYMB_65 || _la == StellaParserStellaIdent {
			{
				p.SetState(540)

				var _x = p.stellatype(0)

				localctx.(*TypeListContext)._stellatype = _x
			}
			localctx.(*TypeListContext).types = append(localctx.(*TypeListContext).types, localctx.(*TypeListContext)._stellatype)
			p.SetState(545)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == StellaParserSurrogate_id_SYMB_0 {
				{
					p.SetState(541)
					p.Match(StellaParserSurrogate_id_SYMB_0)
				}
				{
					p.SetState(542)

					var _x = p.stellatype(0)

					localctx.(*TypeListContext)._stellatype = _x
				}
				localctx.(*TypeListContext).types = append(localctx.(*TypeListContext).types, localctx.(*TypeListContext)._stellatype)

				p.SetState(547)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(550)
			p.Match(StellaParserSurrogate_id_SYMB_14)
		}

	case 9:
		localctx = NewTypeUnitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(551)
			p.Match(StellaParserSurrogate_id_SYMB_34)
		}

	case 10:
		localctx = NewTypeVarContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(552)

			var _m = p.Match(StellaParserStellaIdent)

			localctx.(*TypeVarContext).name = _m
		}

	case 11:
		localctx = NewTypeParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(553)
			p.Match(StellaParserSurrogate_id_SYMB_2)
		}
		{
			p.SetState(554)
			p.stellatype(0)
		}
		{
			p.SetState(555)
			p.Match(StellaParserSurrogate_id_SYMB_3)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(564)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTypeSumContext(p, NewStellatypeContext(p, _parentctx, _parentState))
			localctx.(*TypeSumContext).left = _prevctx

			p.PushNewRecursionContext(localctx, _startState, StellaParserRULE_stellatype)
			p.SetState(559)

			if !(p.Precpred(p.GetParserRuleContext(), 8)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
			}
			{
				p.SetState(560)
				p.Match(StellaParserSurrogate_id_SYMB_21)
			}
			{
				p.SetState(561)

				var _x = p.stellatype(9)

				localctx.(*TypeSumContext).right = _x
			}

		}
		p.SetState(566)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())
	}

	return localctx
}

// IRecordFieldTypeContext is an interface to support dynamic dispatch.
type IRecordFieldTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLabel returns the label token.
	GetLabel() antlr.Token

	// SetLabel sets the label token.
	SetLabel(antlr.Token)

	// GetType_ returns the type_ rule contexts.
	GetType_() IStellatypeContext

	// SetType_ sets the type_ rule contexts.
	SetType_(IStellatypeContext)

	// IsRecordFieldTypeContext differentiates from other interfaces.
	IsRecordFieldTypeContext()
}

type RecordFieldTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	label  antlr.Token
	type_  IStellatypeContext
}

func NewEmptyRecordFieldTypeContext() *RecordFieldTypeContext {
	var p = new(RecordFieldTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StellaParserRULE_recordFieldType
	return p
}

func (*RecordFieldTypeContext) IsRecordFieldTypeContext() {}

func NewRecordFieldTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecordFieldTypeContext {
	var p = new(RecordFieldTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StellaParserRULE_recordFieldType

	return p
}

func (s *RecordFieldTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *RecordFieldTypeContext) GetLabel() antlr.Token { return s.label }

func (s *RecordFieldTypeContext) SetLabel(v antlr.Token) { s.label = v }

func (s *RecordFieldTypeContext) GetType_() IStellatypeContext { return s.type_ }

func (s *RecordFieldTypeContext) SetType_(v IStellatypeContext) { s.type_ = v }

func (s *RecordFieldTypeContext) Surrogate_id_SYMB_7() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_7, 0)
}

func (s *RecordFieldTypeContext) StellaIdent() antlr.TerminalNode {
	return s.GetToken(StellaParserStellaIdent, 0)
}

func (s *RecordFieldTypeContext) Stellatype() IStellatypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStellatypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStellatypeContext)
}

func (s *RecordFieldTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordFieldTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecordFieldTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterRecordFieldType(s)
	}
}

func (s *RecordFieldTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitRecordFieldType(s)
	}
}

func (p *StellaParser) RecordFieldType() (localctx IRecordFieldTypeContext) {
	this := p
	_ = this

	localctx = NewRecordFieldTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, StellaParserRULE_recordFieldType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(567)

		var _m = p.Match(StellaParserStellaIdent)

		localctx.(*RecordFieldTypeContext).label = _m
	}
	{
		p.SetState(568)
		p.Match(StellaParserSurrogate_id_SYMB_7)
	}
	{
		p.SetState(569)

		var _x = p.stellatype(0)

		localctx.(*RecordFieldTypeContext).type_ = _x
	}

	return localctx
}

// IVariantFieldTypeContext is an interface to support dynamic dispatch.
type IVariantFieldTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLabel returns the label token.
	GetLabel() antlr.Token

	// SetLabel sets the label token.
	SetLabel(antlr.Token)

	// GetType_ returns the type_ rule contexts.
	GetType_() IStellatypeContext

	// SetType_ sets the type_ rule contexts.
	SetType_(IStellatypeContext)

	// IsVariantFieldTypeContext differentiates from other interfaces.
	IsVariantFieldTypeContext()
}

type VariantFieldTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	label  antlr.Token
	type_  IStellatypeContext
}

func NewEmptyVariantFieldTypeContext() *VariantFieldTypeContext {
	var p = new(VariantFieldTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StellaParserRULE_variantFieldType
	return p
}

func (*VariantFieldTypeContext) IsVariantFieldTypeContext() {}

func NewVariantFieldTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariantFieldTypeContext {
	var p = new(VariantFieldTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StellaParserRULE_variantFieldType

	return p
}

func (s *VariantFieldTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *VariantFieldTypeContext) GetLabel() antlr.Token { return s.label }

func (s *VariantFieldTypeContext) SetLabel(v antlr.Token) { s.label = v }

func (s *VariantFieldTypeContext) GetType_() IStellatypeContext { return s.type_ }

func (s *VariantFieldTypeContext) SetType_(v IStellatypeContext) { s.type_ = v }

func (s *VariantFieldTypeContext) StellaIdent() antlr.TerminalNode {
	return s.GetToken(StellaParserStellaIdent, 0)
}

func (s *VariantFieldTypeContext) Surrogate_id_SYMB_7() antlr.TerminalNode {
	return s.GetToken(StellaParserSurrogate_id_SYMB_7, 0)
}

func (s *VariantFieldTypeContext) Stellatype() IStellatypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStellatypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStellatypeContext)
}

func (s *VariantFieldTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariantFieldTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariantFieldTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.EnterVariantFieldType(s)
	}
}

func (s *VariantFieldTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StellaParserListener); ok {
		listenerT.ExitVariantFieldType(s)
	}
}

func (p *StellaParser) VariantFieldType() (localctx IVariantFieldTypeContext) {
	this := p
	_ = this

	localctx = NewVariantFieldTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, StellaParserRULE_variantFieldType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(571)

		var _m = p.Match(StellaParserStellaIdent)

		localctx.(*VariantFieldTypeContext).label = _m
	}
	p.SetState(574)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StellaParserSurrogate_id_SYMB_7 {
		{
			p.SetState(572)
			p.Match(StellaParserSurrogate_id_SYMB_7)
		}
		{
			p.SetState(573)

			var _x = p.stellatype(0)

			localctx.(*VariantFieldTypeContext).type_ = _x
		}

	}

	return localctx
}

func (p *StellaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 15:
		var t *StellatypeContext = nil
		if localctx != nil {
			t = localctx.(*StellatypeContext)
		}
		return p.Stellatype_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *StellaParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 25)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 24)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 47)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 46)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 26)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StellaParser) Stellatype_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 18:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
