// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type StellaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var stellalexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func stellalexerLexerInit() {
	staticData := &stellalexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "','", "';'", "'('", "')'", "'{'", "'}'", "'='", "':'", "'->'",
		"'=>'", "'<|'", "'|>'", "'['", "']'", "'<'", "'<='", "'>'", "'>='",
		"'=='", "'!='", "'+'", "'*'", "'List::head'", "'List::isempty'", "'List::tail'",
		"'Nat::pred'", "'Nat::iszero'", "'Nat::rec'", "'.'", "'Bool'", "'Nat'",
		"'Unit'", "'and'", "'as'", "'cons'", "'core'", "'else'", "'extend'",
		"'false'", "'fix'", "'fn'", "'fold'", "'if'", "'in'", "'inline'", "'language'",
		"'let'", "'match'", "'not'", "'or'", "'record'", "'return'", "'succ'",
		"'then'", "'throws'", "'true'", "'type'", "'unfold'", "'variant'", "'with'",
		"'\\u00B5'",
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
		"Surrogate_id_SYMB_60", "COMMENT_antlr_builtin", "StellaIdent", "ExtensionName",
		"INTEGER", "WS", "ErrorToken",
	}
	staticData.ruleNames = []string{
		"LETTER", "CAPITAL", "SMALL", "DIGIT", "Surrogate_id_SYMB_0", "Surrogate_id_SYMB_1",
		"Surrogate_id_SYMB_2", "Surrogate_id_SYMB_3", "Surrogate_id_SYMB_4",
		"Surrogate_id_SYMB_5", "Surrogate_id_SYMB_6", "Surrogate_id_SYMB_7",
		"Surrogate_id_SYMB_8", "Surrogate_id_SYMB_9", "Surrogate_id_SYMB_10",
		"Surrogate_id_SYMB_11", "Surrogate_id_SYMB_12", "Surrogate_id_SYMB_13",
		"Surrogate_id_SYMB_14", "Surrogate_id_SYMB_15", "Surrogate_id_SYMB_16",
		"Surrogate_id_SYMB_17", "Surrogate_id_SYMB_18", "Surrogate_id_SYMB_19",
		"Surrogate_id_SYMB_20", "Surrogate_id_SYMB_21", "Surrogate_id_SYMB_22",
		"Surrogate_id_SYMB_23", "Surrogate_id_SYMB_24", "Surrogate_id_SYMB_25",
		"Surrogate_id_SYMB_26", "Surrogate_id_SYMB_27", "Surrogate_id_SYMB_28",
		"Surrogate_id_SYMB_29", "Surrogate_id_SYMB_30", "Surrogate_id_SYMB_31",
		"Surrogate_id_SYMB_32", "Surrogate_id_SYMB_33", "Surrogate_id_SYMB_34",
		"Surrogate_id_SYMB_35", "Surrogate_id_SYMB_36", "Surrogate_id_SYMB_37",
		"Surrogate_id_SYMB_38", "Surrogate_id_SYMB_39", "Surrogate_id_SYMB_40",
		"Surrogate_id_SYMB_41", "Surrogate_id_SYMB_42", "Surrogate_id_SYMB_43",
		"Surrogate_id_SYMB_44", "Surrogate_id_SYMB_45", "Surrogate_id_SYMB_46",
		"Surrogate_id_SYMB_47", "Surrogate_id_SYMB_48", "Surrogate_id_SYMB_49",
		"Surrogate_id_SYMB_50", "Surrogate_id_SYMB_51", "Surrogate_id_SYMB_52",
		"Surrogate_id_SYMB_53", "Surrogate_id_SYMB_54", "Surrogate_id_SYMB_55",
		"Surrogate_id_SYMB_56", "Surrogate_id_SYMB_57", "Surrogate_id_SYMB_58",
		"Surrogate_id_SYMB_59", "Surrogate_id_SYMB_60", "COMMENT_antlr_builtin",
		"StellaIdent", "ExtensionName", "INTEGER", "WS", "Escapable", "ErrorToken",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 67, 497, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 1, 0, 1, 0, 3,
		0, 148, 8, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32,
		1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37,
		1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1,
		43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46,
		1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1,
		48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49,
		1, 50, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1,
		52, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 54,
		1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1,
		56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 58,
		1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 59, 1, 59, 1, 59, 1, 59, 1,
		59, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 61, 1, 61, 1, 61, 1, 61, 1, 61,
		1, 61, 1, 61, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1,
		63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 64, 1, 64, 1, 65, 1, 65, 1, 65, 1, 65,
		5, 65, 444, 8, 65, 10, 65, 12, 65, 447, 9, 65, 1, 65, 3, 65, 450, 8, 65,
		1, 65, 1, 65, 3, 65, 454, 8, 65, 1, 65, 1, 65, 1, 66, 1, 66, 3, 66, 460,
		8, 66, 1, 66, 1, 66, 1, 66, 3, 66, 465, 8, 66, 5, 66, 467, 8, 66, 10, 66,
		12, 66, 470, 9, 66, 1, 67, 1, 67, 1, 67, 1, 67, 3, 67, 476, 8, 67, 4, 67,
		478, 8, 67, 11, 67, 12, 67, 479, 1, 68, 4, 68, 483, 8, 68, 11, 68, 12,
		68, 484, 1, 69, 4, 69, 488, 8, 69, 11, 69, 12, 69, 489, 1, 69, 1, 69, 1,
		70, 1, 70, 1, 71, 1, 71, 0, 0, 72, 1, 0, 3, 0, 5, 0, 7, 0, 9, 1, 11, 2,
		13, 3, 15, 4, 17, 5, 19, 6, 21, 7, 23, 8, 25, 9, 27, 10, 29, 11, 31, 12,
		33, 13, 35, 14, 37, 15, 39, 16, 41, 17, 43, 18, 45, 19, 47, 20, 49, 21,
		51, 22, 53, 23, 55, 24, 57, 25, 59, 26, 61, 27, 63, 28, 65, 29, 67, 30,
		69, 31, 71, 32, 73, 33, 75, 34, 77, 35, 79, 36, 81, 37, 83, 38, 85, 39,
		87, 40, 89, 41, 91, 42, 93, 43, 95, 44, 97, 45, 99, 46, 101, 47, 103, 48,
		105, 49, 107, 50, 109, 51, 111, 52, 113, 53, 115, 54, 117, 55, 119, 56,
		121, 57, 123, 58, 125, 59, 127, 60, 129, 61, 131, 62, 133, 63, 135, 64,
		137, 65, 139, 66, 141, 0, 143, 67, 1, 0, 8, 3, 0, 65, 90, 192, 214, 216,
		222, 3, 0, 97, 122, 223, 246, 248, 255, 1, 0, 48, 57, 2, 0, 10, 10, 13,
		13, 5, 0, 33, 33, 45, 45, 58, 58, 63, 63, 95, 95, 2, 0, 45, 45, 95, 95,
		3, 0, 9, 10, 12, 13, 32, 32, 6, 0, 34, 34, 92, 92, 102, 102, 110, 110,
		114, 114, 116, 116, 504, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0,
		0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0,
		0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1,
		0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67,
		1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0,
		75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0,
		0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0,
		0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0,
		0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105,
		1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0,
		0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1,
		0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0,
		127, 1, 0, 0, 0, 0, 129, 1, 0, 0, 0, 0, 131, 1, 0, 0, 0, 0, 133, 1, 0,
		0, 0, 0, 135, 1, 0, 0, 0, 0, 137, 1, 0, 0, 0, 0, 139, 1, 0, 0, 0, 0, 143,
		1, 0, 0, 0, 1, 147, 1, 0, 0, 0, 3, 149, 1, 0, 0, 0, 5, 151, 1, 0, 0, 0,
		7, 153, 1, 0, 0, 0, 9, 155, 1, 0, 0, 0, 11, 157, 1, 0, 0, 0, 13, 159, 1,
		0, 0, 0, 15, 161, 1, 0, 0, 0, 17, 163, 1, 0, 0, 0, 19, 165, 1, 0, 0, 0,
		21, 167, 1, 0, 0, 0, 23, 169, 1, 0, 0, 0, 25, 171, 1, 0, 0, 0, 27, 174,
		1, 0, 0, 0, 29, 177, 1, 0, 0, 0, 31, 180, 1, 0, 0, 0, 33, 183, 1, 0, 0,
		0, 35, 185, 1, 0, 0, 0, 37, 187, 1, 0, 0, 0, 39, 189, 1, 0, 0, 0, 41, 192,
		1, 0, 0, 0, 43, 194, 1, 0, 0, 0, 45, 197, 1, 0, 0, 0, 47, 200, 1, 0, 0,
		0, 49, 203, 1, 0, 0, 0, 51, 205, 1, 0, 0, 0, 53, 207, 1, 0, 0, 0, 55, 218,
		1, 0, 0, 0, 57, 232, 1, 0, 0, 0, 59, 243, 1, 0, 0, 0, 61, 253, 1, 0, 0,
		0, 63, 265, 1, 0, 0, 0, 65, 274, 1, 0, 0, 0, 67, 276, 1, 0, 0, 0, 69, 281,
		1, 0, 0, 0, 71, 285, 1, 0, 0, 0, 73, 290, 1, 0, 0, 0, 75, 294, 1, 0, 0,
		0, 77, 297, 1, 0, 0, 0, 79, 302, 1, 0, 0, 0, 81, 307, 1, 0, 0, 0, 83, 312,
		1, 0, 0, 0, 85, 319, 1, 0, 0, 0, 87, 325, 1, 0, 0, 0, 89, 329, 1, 0, 0,
		0, 91, 332, 1, 0, 0, 0, 93, 337, 1, 0, 0, 0, 95, 340, 1, 0, 0, 0, 97, 343,
		1, 0, 0, 0, 99, 350, 1, 0, 0, 0, 101, 359, 1, 0, 0, 0, 103, 363, 1, 0,
		0, 0, 105, 369, 1, 0, 0, 0, 107, 373, 1, 0, 0, 0, 109, 376, 1, 0, 0, 0,
		111, 383, 1, 0, 0, 0, 113, 390, 1, 0, 0, 0, 115, 395, 1, 0, 0, 0, 117,
		400, 1, 0, 0, 0, 119, 407, 1, 0, 0, 0, 121, 412, 1, 0, 0, 0, 123, 417,
		1, 0, 0, 0, 125, 424, 1, 0, 0, 0, 127, 432, 1, 0, 0, 0, 129, 437, 1, 0,
		0, 0, 131, 439, 1, 0, 0, 0, 133, 459, 1, 0, 0, 0, 135, 471, 1, 0, 0, 0,
		137, 482, 1, 0, 0, 0, 139, 487, 1, 0, 0, 0, 141, 493, 1, 0, 0, 0, 143,
		495, 1, 0, 0, 0, 145, 148, 3, 3, 1, 0, 146, 148, 3, 5, 2, 0, 147, 145,
		1, 0, 0, 0, 147, 146, 1, 0, 0, 0, 148, 2, 1, 0, 0, 0, 149, 150, 7, 0, 0,
		0, 150, 4, 1, 0, 0, 0, 151, 152, 7, 1, 0, 0, 152, 6, 1, 0, 0, 0, 153, 154,
		7, 2, 0, 0, 154, 8, 1, 0, 0, 0, 155, 156, 5, 44, 0, 0, 156, 10, 1, 0, 0,
		0, 157, 158, 5, 59, 0, 0, 158, 12, 1, 0, 0, 0, 159, 160, 5, 40, 0, 0, 160,
		14, 1, 0, 0, 0, 161, 162, 5, 41, 0, 0, 162, 16, 1, 0, 0, 0, 163, 164, 5,
		123, 0, 0, 164, 18, 1, 0, 0, 0, 165, 166, 5, 125, 0, 0, 166, 20, 1, 0,
		0, 0, 167, 168, 5, 61, 0, 0, 168, 22, 1, 0, 0, 0, 169, 170, 5, 58, 0, 0,
		170, 24, 1, 0, 0, 0, 171, 172, 5, 45, 0, 0, 172, 173, 5, 62, 0, 0, 173,
		26, 1, 0, 0, 0, 174, 175, 5, 61, 0, 0, 175, 176, 5, 62, 0, 0, 176, 28,
		1, 0, 0, 0, 177, 178, 5, 60, 0, 0, 178, 179, 5, 124, 0, 0, 179, 30, 1,
		0, 0, 0, 180, 181, 5, 124, 0, 0, 181, 182, 5, 62, 0, 0, 182, 32, 1, 0,
		0, 0, 183, 184, 5, 91, 0, 0, 184, 34, 1, 0, 0, 0, 185, 186, 5, 93, 0, 0,
		186, 36, 1, 0, 0, 0, 187, 188, 5, 60, 0, 0, 188, 38, 1, 0, 0, 0, 189, 190,
		5, 60, 0, 0, 190, 191, 5, 61, 0, 0, 191, 40, 1, 0, 0, 0, 192, 193, 5, 62,
		0, 0, 193, 42, 1, 0, 0, 0, 194, 195, 5, 62, 0, 0, 195, 196, 5, 61, 0, 0,
		196, 44, 1, 0, 0, 0, 197, 198, 5, 61, 0, 0, 198, 199, 5, 61, 0, 0, 199,
		46, 1, 0, 0, 0, 200, 201, 5, 33, 0, 0, 201, 202, 5, 61, 0, 0, 202, 48,
		1, 0, 0, 0, 203, 204, 5, 43, 0, 0, 204, 50, 1, 0, 0, 0, 205, 206, 5, 42,
		0, 0, 206, 52, 1, 0, 0, 0, 207, 208, 5, 76, 0, 0, 208, 209, 5, 105, 0,
		0, 209, 210, 5, 115, 0, 0, 210, 211, 5, 116, 0, 0, 211, 212, 5, 58, 0,
		0, 212, 213, 5, 58, 0, 0, 213, 214, 5, 104, 0, 0, 214, 215, 5, 101, 0,
		0, 215, 216, 5, 97, 0, 0, 216, 217, 5, 100, 0, 0, 217, 54, 1, 0, 0, 0,
		218, 219, 5, 76, 0, 0, 219, 220, 5, 105, 0, 0, 220, 221, 5, 115, 0, 0,
		221, 222, 5, 116, 0, 0, 222, 223, 5, 58, 0, 0, 223, 224, 5, 58, 0, 0, 224,
		225, 5, 105, 0, 0, 225, 226, 5, 115, 0, 0, 226, 227, 5, 101, 0, 0, 227,
		228, 5, 109, 0, 0, 228, 229, 5, 112, 0, 0, 229, 230, 5, 116, 0, 0, 230,
		231, 5, 121, 0, 0, 231, 56, 1, 0, 0, 0, 232, 233, 5, 76, 0, 0, 233, 234,
		5, 105, 0, 0, 234, 235, 5, 115, 0, 0, 235, 236, 5, 116, 0, 0, 236, 237,
		5, 58, 0, 0, 237, 238, 5, 58, 0, 0, 238, 239, 5, 116, 0, 0, 239, 240, 5,
		97, 0, 0, 240, 241, 5, 105, 0, 0, 241, 242, 5, 108, 0, 0, 242, 58, 1, 0,
		0, 0, 243, 244, 5, 78, 0, 0, 244, 245, 5, 97, 0, 0, 245, 246, 5, 116, 0,
		0, 246, 247, 5, 58, 0, 0, 247, 248, 5, 58, 0, 0, 248, 249, 5, 112, 0, 0,
		249, 250, 5, 114, 0, 0, 250, 251, 5, 101, 0, 0, 251, 252, 5, 100, 0, 0,
		252, 60, 1, 0, 0, 0, 253, 254, 5, 78, 0, 0, 254, 255, 5, 97, 0, 0, 255,
		256, 5, 116, 0, 0, 256, 257, 5, 58, 0, 0, 257, 258, 5, 58, 0, 0, 258, 259,
		5, 105, 0, 0, 259, 260, 5, 115, 0, 0, 260, 261, 5, 122, 0, 0, 261, 262,
		5, 101, 0, 0, 262, 263, 5, 114, 0, 0, 263, 264, 5, 111, 0, 0, 264, 62,
		1, 0, 0, 0, 265, 266, 5, 78, 0, 0, 266, 267, 5, 97, 0, 0, 267, 268, 5,
		116, 0, 0, 268, 269, 5, 58, 0, 0, 269, 270, 5, 58, 0, 0, 270, 271, 5, 114,
		0, 0, 271, 272, 5, 101, 0, 0, 272, 273, 5, 99, 0, 0, 273, 64, 1, 0, 0,
		0, 274, 275, 5, 46, 0, 0, 275, 66, 1, 0, 0, 0, 276, 277, 5, 66, 0, 0, 277,
		278, 5, 111, 0, 0, 278, 279, 5, 111, 0, 0, 279, 280, 5, 108, 0, 0, 280,
		68, 1, 0, 0, 0, 281, 282, 5, 78, 0, 0, 282, 283, 5, 97, 0, 0, 283, 284,
		5, 116, 0, 0, 284, 70, 1, 0, 0, 0, 285, 286, 5, 85, 0, 0, 286, 287, 5,
		110, 0, 0, 287, 288, 5, 105, 0, 0, 288, 289, 5, 116, 0, 0, 289, 72, 1,
		0, 0, 0, 290, 291, 5, 97, 0, 0, 291, 292, 5, 110, 0, 0, 292, 293, 5, 100,
		0, 0, 293, 74, 1, 0, 0, 0, 294, 295, 5, 97, 0, 0, 295, 296, 5, 115, 0,
		0, 296, 76, 1, 0, 0, 0, 297, 298, 5, 99, 0, 0, 298, 299, 5, 111, 0, 0,
		299, 300, 5, 110, 0, 0, 300, 301, 5, 115, 0, 0, 301, 78, 1, 0, 0, 0, 302,
		303, 5, 99, 0, 0, 303, 304, 5, 111, 0, 0, 304, 305, 5, 114, 0, 0, 305,
		306, 5, 101, 0, 0, 306, 80, 1, 0, 0, 0, 307, 308, 5, 101, 0, 0, 308, 309,
		5, 108, 0, 0, 309, 310, 5, 115, 0, 0, 310, 311, 5, 101, 0, 0, 311, 82,
		1, 0, 0, 0, 312, 313, 5, 101, 0, 0, 313, 314, 5, 120, 0, 0, 314, 315, 5,
		116, 0, 0, 315, 316, 5, 101, 0, 0, 316, 317, 5, 110, 0, 0, 317, 318, 5,
		100, 0, 0, 318, 84, 1, 0, 0, 0, 319, 320, 5, 102, 0, 0, 320, 321, 5, 97,
		0, 0, 321, 322, 5, 108, 0, 0, 322, 323, 5, 115, 0, 0, 323, 324, 5, 101,
		0, 0, 324, 86, 1, 0, 0, 0, 325, 326, 5, 102, 0, 0, 326, 327, 5, 105, 0,
		0, 327, 328, 5, 120, 0, 0, 328, 88, 1, 0, 0, 0, 329, 330, 5, 102, 0, 0,
		330, 331, 5, 110, 0, 0, 331, 90, 1, 0, 0, 0, 332, 333, 5, 102, 0, 0, 333,
		334, 5, 111, 0, 0, 334, 335, 5, 108, 0, 0, 335, 336, 5, 100, 0, 0, 336,
		92, 1, 0, 0, 0, 337, 338, 5, 105, 0, 0, 338, 339, 5, 102, 0, 0, 339, 94,
		1, 0, 0, 0, 340, 341, 5, 105, 0, 0, 341, 342, 5, 110, 0, 0, 342, 96, 1,
		0, 0, 0, 343, 344, 5, 105, 0, 0, 344, 345, 5, 110, 0, 0, 345, 346, 5, 108,
		0, 0, 346, 347, 5, 105, 0, 0, 347, 348, 5, 110, 0, 0, 348, 349, 5, 101,
		0, 0, 349, 98, 1, 0, 0, 0, 350, 351, 5, 108, 0, 0, 351, 352, 5, 97, 0,
		0, 352, 353, 5, 110, 0, 0, 353, 354, 5, 103, 0, 0, 354, 355, 5, 117, 0,
		0, 355, 356, 5, 97, 0, 0, 356, 357, 5, 103, 0, 0, 357, 358, 5, 101, 0,
		0, 358, 100, 1, 0, 0, 0, 359, 360, 5, 108, 0, 0, 360, 361, 5, 101, 0, 0,
		361, 362, 5, 116, 0, 0, 362, 102, 1, 0, 0, 0, 363, 364, 5, 109, 0, 0, 364,
		365, 5, 97, 0, 0, 365, 366, 5, 116, 0, 0, 366, 367, 5, 99, 0, 0, 367, 368,
		5, 104, 0, 0, 368, 104, 1, 0, 0, 0, 369, 370, 5, 110, 0, 0, 370, 371, 5,
		111, 0, 0, 371, 372, 5, 116, 0, 0, 372, 106, 1, 0, 0, 0, 373, 374, 5, 111,
		0, 0, 374, 375, 5, 114, 0, 0, 375, 108, 1, 0, 0, 0, 376, 377, 5, 114, 0,
		0, 377, 378, 5, 101, 0, 0, 378, 379, 5, 99, 0, 0, 379, 380, 5, 111, 0,
		0, 380, 381, 5, 114, 0, 0, 381, 382, 5, 100, 0, 0, 382, 110, 1, 0, 0, 0,
		383, 384, 5, 114, 0, 0, 384, 385, 5, 101, 0, 0, 385, 386, 5, 116, 0, 0,
		386, 387, 5, 117, 0, 0, 387, 388, 5, 114, 0, 0, 388, 389, 5, 110, 0, 0,
		389, 112, 1, 0, 0, 0, 390, 391, 5, 115, 0, 0, 391, 392, 5, 117, 0, 0, 392,
		393, 5, 99, 0, 0, 393, 394, 5, 99, 0, 0, 394, 114, 1, 0, 0, 0, 395, 396,
		5, 116, 0, 0, 396, 397, 5, 104, 0, 0, 397, 398, 5, 101, 0, 0, 398, 399,
		5, 110, 0, 0, 399, 116, 1, 0, 0, 0, 400, 401, 5, 116, 0, 0, 401, 402, 5,
		104, 0, 0, 402, 403, 5, 114, 0, 0, 403, 404, 5, 111, 0, 0, 404, 405, 5,
		119, 0, 0, 405, 406, 5, 115, 0, 0, 406, 118, 1, 0, 0, 0, 407, 408, 5, 116,
		0, 0, 408, 409, 5, 114, 0, 0, 409, 410, 5, 117, 0, 0, 410, 411, 5, 101,
		0, 0, 411, 120, 1, 0, 0, 0, 412, 413, 5, 116, 0, 0, 413, 414, 5, 121, 0,
		0, 414, 415, 5, 112, 0, 0, 415, 416, 5, 101, 0, 0, 416, 122, 1, 0, 0, 0,
		417, 418, 5, 117, 0, 0, 418, 419, 5, 110, 0, 0, 419, 420, 5, 102, 0, 0,
		420, 421, 5, 111, 0, 0, 421, 422, 5, 108, 0, 0, 422, 423, 5, 100, 0, 0,
		423, 124, 1, 0, 0, 0, 424, 425, 5, 118, 0, 0, 425, 426, 5, 97, 0, 0, 426,
		427, 5, 114, 0, 0, 427, 428, 5, 105, 0, 0, 428, 429, 5, 97, 0, 0, 429,
		430, 5, 110, 0, 0, 430, 431, 5, 116, 0, 0, 431, 126, 1, 0, 0, 0, 432, 433,
		5, 119, 0, 0, 433, 434, 5, 105, 0, 0, 434, 435, 5, 116, 0, 0, 435, 436,
		5, 104, 0, 0, 436, 128, 1, 0, 0, 0, 437, 438, 5, 181, 0, 0, 438, 130, 1,
		0, 0, 0, 439, 440, 5, 47, 0, 0, 440, 441, 5, 47, 0, 0, 441, 445, 1, 0,
		0, 0, 442, 444, 8, 3, 0, 0, 443, 442, 1, 0, 0, 0, 444, 447, 1, 0, 0, 0,
		445, 443, 1, 0, 0, 0, 445, 446, 1, 0, 0, 0, 446, 453, 1, 0, 0, 0, 447,
		445, 1, 0, 0, 0, 448, 450, 5, 13, 0, 0, 449, 448, 1, 0, 0, 0, 449, 450,
		1, 0, 0, 0, 450, 451, 1, 0, 0, 0, 451, 454, 5, 10, 0, 0, 452, 454, 5, 0,
		0, 1, 453, 449, 1, 0, 0, 0, 453, 452, 1, 0, 0, 0, 454, 455, 1, 0, 0, 0,
		455, 456, 6, 65, 0, 0, 456, 132, 1, 0, 0, 0, 457, 460, 5, 95, 0, 0, 458,
		460, 3, 1, 0, 0, 459, 457, 1, 0, 0, 0, 459, 458, 1, 0, 0, 0, 460, 468,
		1, 0, 0, 0, 461, 467, 7, 4, 0, 0, 462, 465, 3, 7, 3, 0, 463, 465, 3, 1,
		0, 0, 464, 462, 1, 0, 0, 0, 464, 463, 1, 0, 0, 0, 465, 467, 1, 0, 0, 0,
		466, 461, 1, 0, 0, 0, 466, 464, 1, 0, 0, 0, 467, 470, 1, 0, 0, 0, 468,
		466, 1, 0, 0, 0, 468, 469, 1, 0, 0, 0, 469, 134, 1, 0, 0, 0, 470, 468,
		1, 0, 0, 0, 471, 477, 5, 35, 0, 0, 472, 478, 7, 5, 0, 0, 473, 476, 3, 7,
		3, 0, 474, 476, 3, 1, 0, 0, 475, 473, 1, 0, 0, 0, 475, 474, 1, 0, 0, 0,
		476, 478, 1, 0, 0, 0, 477, 472, 1, 0, 0, 0, 477, 475, 1, 0, 0, 0, 478,
		479, 1, 0, 0, 0, 479, 477, 1, 0, 0, 0, 479, 480, 1, 0, 0, 0, 480, 136,
		1, 0, 0, 0, 481, 483, 3, 7, 3, 0, 482, 481, 1, 0, 0, 0, 483, 484, 1, 0,
		0, 0, 484, 482, 1, 0, 0, 0, 484, 485, 1, 0, 0, 0, 485, 138, 1, 0, 0, 0,
		486, 488, 7, 6, 0, 0, 487, 486, 1, 0, 0, 0, 488, 489, 1, 0, 0, 0, 489,
		487, 1, 0, 0, 0, 489, 490, 1, 0, 0, 0, 490, 491, 1, 0, 0, 0, 491, 492,
		6, 69, 0, 0, 492, 140, 1, 0, 0, 0, 493, 494, 7, 7, 0, 0, 494, 142, 1, 0,
		0, 0, 495, 496, 9, 0, 0, 0, 496, 144, 1, 0, 0, 0, 14, 0, 147, 445, 449,
		453, 459, 464, 466, 468, 475, 477, 479, 484, 489, 1, 6, 0, 0,
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

// StellaLexerInit initializes any static state used to implement StellaLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewStellaLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func StellaLexerInit() {
	staticData := &stellalexerLexerStaticData
	staticData.once.Do(stellalexerLexerInit)
}

// NewStellaLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewStellaLexer(input antlr.CharStream) *StellaLexer {
	StellaLexerInit()
	l := new(StellaLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &stellalexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "StellaLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// StellaLexer tokens.
const (
	StellaLexerSurrogate_id_SYMB_0   = 1
	StellaLexerSurrogate_id_SYMB_1   = 2
	StellaLexerSurrogate_id_SYMB_2   = 3
	StellaLexerSurrogate_id_SYMB_3   = 4
	StellaLexerSurrogate_id_SYMB_4   = 5
	StellaLexerSurrogate_id_SYMB_5   = 6
	StellaLexerSurrogate_id_SYMB_6   = 7
	StellaLexerSurrogate_id_SYMB_7   = 8
	StellaLexerSurrogate_id_SYMB_8   = 9
	StellaLexerSurrogate_id_SYMB_9   = 10
	StellaLexerSurrogate_id_SYMB_10  = 11
	StellaLexerSurrogate_id_SYMB_11  = 12
	StellaLexerSurrogate_id_SYMB_12  = 13
	StellaLexerSurrogate_id_SYMB_13  = 14
	StellaLexerSurrogate_id_SYMB_14  = 15
	StellaLexerSurrogate_id_SYMB_15  = 16
	StellaLexerSurrogate_id_SYMB_16  = 17
	StellaLexerSurrogate_id_SYMB_17  = 18
	StellaLexerSurrogate_id_SYMB_18  = 19
	StellaLexerSurrogate_id_SYMB_19  = 20
	StellaLexerSurrogate_id_SYMB_20  = 21
	StellaLexerSurrogate_id_SYMB_21  = 22
	StellaLexerSurrogate_id_SYMB_22  = 23
	StellaLexerSurrogate_id_SYMB_23  = 24
	StellaLexerSurrogate_id_SYMB_24  = 25
	StellaLexerSurrogate_id_SYMB_25  = 26
	StellaLexerSurrogate_id_SYMB_26  = 27
	StellaLexerSurrogate_id_SYMB_27  = 28
	StellaLexerSurrogate_id_SYMB_28  = 29
	StellaLexerSurrogate_id_SYMB_29  = 30
	StellaLexerSurrogate_id_SYMB_30  = 31
	StellaLexerSurrogate_id_SYMB_31  = 32
	StellaLexerSurrogate_id_SYMB_32  = 33
	StellaLexerSurrogate_id_SYMB_33  = 34
	StellaLexerSurrogate_id_SYMB_34  = 35
	StellaLexerSurrogate_id_SYMB_35  = 36
	StellaLexerSurrogate_id_SYMB_36  = 37
	StellaLexerSurrogate_id_SYMB_37  = 38
	StellaLexerSurrogate_id_SYMB_38  = 39
	StellaLexerSurrogate_id_SYMB_39  = 40
	StellaLexerSurrogate_id_SYMB_40  = 41
	StellaLexerSurrogate_id_SYMB_41  = 42
	StellaLexerSurrogate_id_SYMB_42  = 43
	StellaLexerSurrogate_id_SYMB_43  = 44
	StellaLexerSurrogate_id_SYMB_44  = 45
	StellaLexerSurrogate_id_SYMB_45  = 46
	StellaLexerSurrogate_id_SYMB_46  = 47
	StellaLexerSurrogate_id_SYMB_47  = 48
	StellaLexerSurrogate_id_SYMB_48  = 49
	StellaLexerSurrogate_id_SYMB_49  = 50
	StellaLexerSurrogate_id_SYMB_50  = 51
	StellaLexerSurrogate_id_SYMB_51  = 52
	StellaLexerSurrogate_id_SYMB_52  = 53
	StellaLexerSurrogate_id_SYMB_53  = 54
	StellaLexerSurrogate_id_SYMB_54  = 55
	StellaLexerSurrogate_id_SYMB_55  = 56
	StellaLexerSurrogate_id_SYMB_56  = 57
	StellaLexerSurrogate_id_SYMB_57  = 58
	StellaLexerSurrogate_id_SYMB_58  = 59
	StellaLexerSurrogate_id_SYMB_59  = 60
	StellaLexerSurrogate_id_SYMB_60  = 61
	StellaLexerCOMMENT_antlr_builtin = 62
	StellaLexerStellaIdent           = 63
	StellaLexerExtensionName         = 64
	StellaLexerINTEGER               = 65
	StellaLexerWS                    = 66
	StellaLexerErrorToken            = 67
)
