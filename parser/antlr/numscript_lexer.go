// Code generated from Numscript.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type NumscriptLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var NumscriptLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func numscriptlexerLexerInit() {
	staticData := &NumscriptLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "", "'vars'", "'max'", "'source'", "'destination'",
		"'send'", "'from'", "'to'", "'remaining'", "'allowing'", "'unbounded'",
		"'overdraft'", "'('", "')'", "'['", "']'", "'{'", "'}'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "NEWLINE", "MULTILINE_COMMENT", "LINE_COMMENT", "VARS", "MAX",
		"SOURCE", "DESTINATION", "SEND", "FROM", "TO", "REMAINING", "ALLOWING",
		"UNBOUNDED", "OVERDRAFT", "LPARENS", "RPARENS", "LBRACKET", "RBRACKET",
		"LBRACE", "RBRACE", "EQ", "RATIO_PORTION_LITERAL", "PERCENTAGE_PORTION_LITERAL",
		"TYPE_IDENT", "NUMBER", "VARIABLE_NAME", "ACCOUNT", "ASSET",
	}
	staticData.RuleNames = []string{
		"WS", "NEWLINE", "MULTILINE_COMMENT", "LINE_COMMENT", "VARS", "MAX",
		"SOURCE", "DESTINATION", "SEND", "FROM", "TO", "REMAINING", "ALLOWING",
		"UNBOUNDED", "OVERDRAFT", "LPARENS", "RPARENS", "LBRACKET", "RBRACKET",
		"LBRACE", "RBRACE", "EQ", "RATIO_PORTION_LITERAL", "PERCENTAGE_PORTION_LITERAL",
		"TYPE_IDENT", "NUMBER", "VARIABLE_NAME", "ACCOUNT", "ASSET",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 29, 269, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 4, 0, 61, 8, 0, 11, 0,
		12, 0, 62, 1, 0, 1, 0, 1, 1, 4, 1, 68, 8, 1, 11, 1, 12, 1, 69, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 5, 2, 77, 8, 2, 10, 2, 12, 2, 80, 9, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 91, 8, 3, 10, 3, 12, 3,
		94, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		22, 4, 22, 195, 8, 22, 11, 22, 12, 22, 196, 1, 22, 3, 22, 200, 8, 22, 1,
		22, 1, 22, 3, 22, 204, 8, 22, 1, 22, 4, 22, 207, 8, 22, 11, 22, 12, 22,
		208, 1, 23, 4, 23, 212, 8, 23, 11, 23, 12, 23, 213, 1, 23, 1, 23, 4, 23,
		218, 8, 23, 11, 23, 12, 23, 219, 3, 23, 222, 8, 23, 1, 23, 1, 23, 1, 24,
		4, 24, 227, 8, 24, 11, 24, 12, 24, 228, 1, 25, 4, 25, 232, 8, 25, 11, 25,
		12, 25, 233, 1, 26, 1, 26, 4, 26, 238, 8, 26, 11, 26, 12, 26, 239, 1, 26,
		5, 26, 243, 8, 26, 10, 26, 12, 26, 246, 9, 26, 1, 27, 1, 27, 4, 27, 250,
		8, 27, 11, 27, 12, 27, 251, 1, 27, 1, 27, 4, 27, 256, 8, 27, 11, 27, 12,
		27, 257, 5, 27, 260, 8, 27, 10, 27, 12, 27, 263, 9, 27, 1, 28, 4, 28, 266,
		8, 28, 11, 28, 12, 28, 267, 2, 78, 92, 0, 29, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 1, 0, 9, 3, 0, 9, 10, 13,
		13, 32, 32, 2, 0, 10, 10, 13, 13, 1, 0, 48, 57, 1, 0, 32, 32, 1, 0, 97,
		122, 2, 0, 95, 95, 97, 122, 3, 0, 48, 57, 95, 95, 97, 122, 5, 0, 45, 45,
		48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 47, 57, 65, 90, 288, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0,
		0, 0, 0, 57, 1, 0, 0, 0, 1, 60, 1, 0, 0, 0, 3, 67, 1, 0, 0, 0, 5, 71, 1,
		0, 0, 0, 7, 86, 1, 0, 0, 0, 9, 99, 1, 0, 0, 0, 11, 104, 1, 0, 0, 0, 13,
		108, 1, 0, 0, 0, 15, 115, 1, 0, 0, 0, 17, 127, 1, 0, 0, 0, 19, 132, 1,
		0, 0, 0, 21, 137, 1, 0, 0, 0, 23, 140, 1, 0, 0, 0, 25, 150, 1, 0, 0, 0,
		27, 159, 1, 0, 0, 0, 29, 169, 1, 0, 0, 0, 31, 179, 1, 0, 0, 0, 33, 181,
		1, 0, 0, 0, 35, 183, 1, 0, 0, 0, 37, 185, 1, 0, 0, 0, 39, 187, 1, 0, 0,
		0, 41, 189, 1, 0, 0, 0, 43, 191, 1, 0, 0, 0, 45, 194, 1, 0, 0, 0, 47, 211,
		1, 0, 0, 0, 49, 226, 1, 0, 0, 0, 51, 231, 1, 0, 0, 0, 53, 235, 1, 0, 0,
		0, 55, 247, 1, 0, 0, 0, 57, 265, 1, 0, 0, 0, 59, 61, 7, 0, 0, 0, 60, 59,
		1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0,
		63, 64, 1, 0, 0, 0, 64, 65, 6, 0, 0, 0, 65, 2, 1, 0, 0, 0, 66, 68, 7, 1,
		0, 0, 67, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70,
		1, 0, 0, 0, 70, 4, 1, 0, 0, 0, 71, 72, 5, 47, 0, 0, 72, 73, 5, 42, 0, 0,
		73, 78, 1, 0, 0, 0, 74, 77, 3, 5, 2, 0, 75, 77, 9, 0, 0, 0, 76, 74, 1,
		0, 0, 0, 76, 75, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 78,
		76, 1, 0, 0, 0, 79, 81, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 82, 5, 42,
		0, 0, 82, 83, 5, 47, 0, 0, 83, 84, 1, 0, 0, 0, 84, 85, 6, 2, 0, 0, 85,
		6, 1, 0, 0, 0, 86, 87, 5, 47, 0, 0, 87, 88, 5, 47, 0, 0, 88, 92, 1, 0,
		0, 0, 89, 91, 9, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 93,
		1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 93, 95, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0,
		95, 96, 3, 3, 1, 0, 96, 97, 1, 0, 0, 0, 97, 98, 6, 3, 0, 0, 98, 8, 1, 0,
		0, 0, 99, 100, 5, 118, 0, 0, 100, 101, 5, 97, 0, 0, 101, 102, 5, 114, 0,
		0, 102, 103, 5, 115, 0, 0, 103, 10, 1, 0, 0, 0, 104, 105, 5, 109, 0, 0,
		105, 106, 5, 97, 0, 0, 106, 107, 5, 120, 0, 0, 107, 12, 1, 0, 0, 0, 108,
		109, 5, 115, 0, 0, 109, 110, 5, 111, 0, 0, 110, 111, 5, 117, 0, 0, 111,
		112, 5, 114, 0, 0, 112, 113, 5, 99, 0, 0, 113, 114, 5, 101, 0, 0, 114,
		14, 1, 0, 0, 0, 115, 116, 5, 100, 0, 0, 116, 117, 5, 101, 0, 0, 117, 118,
		5, 115, 0, 0, 118, 119, 5, 116, 0, 0, 119, 120, 5, 105, 0, 0, 120, 121,
		5, 110, 0, 0, 121, 122, 5, 97, 0, 0, 122, 123, 5, 116, 0, 0, 123, 124,
		5, 105, 0, 0, 124, 125, 5, 111, 0, 0, 125, 126, 5, 110, 0, 0, 126, 16,
		1, 0, 0, 0, 127, 128, 5, 115, 0, 0, 128, 129, 5, 101, 0, 0, 129, 130, 5,
		110, 0, 0, 130, 131, 5, 100, 0, 0, 131, 18, 1, 0, 0, 0, 132, 133, 5, 102,
		0, 0, 133, 134, 5, 114, 0, 0, 134, 135, 5, 111, 0, 0, 135, 136, 5, 109,
		0, 0, 136, 20, 1, 0, 0, 0, 137, 138, 5, 116, 0, 0, 138, 139, 5, 111, 0,
		0, 139, 22, 1, 0, 0, 0, 140, 141, 5, 114, 0, 0, 141, 142, 5, 101, 0, 0,
		142, 143, 5, 109, 0, 0, 143, 144, 5, 97, 0, 0, 144, 145, 5, 105, 0, 0,
		145, 146, 5, 110, 0, 0, 146, 147, 5, 105, 0, 0, 147, 148, 5, 110, 0, 0,
		148, 149, 5, 103, 0, 0, 149, 24, 1, 0, 0, 0, 150, 151, 5, 97, 0, 0, 151,
		152, 5, 108, 0, 0, 152, 153, 5, 108, 0, 0, 153, 154, 5, 111, 0, 0, 154,
		155, 5, 119, 0, 0, 155, 156, 5, 105, 0, 0, 156, 157, 5, 110, 0, 0, 157,
		158, 5, 103, 0, 0, 158, 26, 1, 0, 0, 0, 159, 160, 5, 117, 0, 0, 160, 161,
		5, 110, 0, 0, 161, 162, 5, 98, 0, 0, 162, 163, 5, 111, 0, 0, 163, 164,
		5, 117, 0, 0, 164, 165, 5, 110, 0, 0, 165, 166, 5, 100, 0, 0, 166, 167,
		5, 101, 0, 0, 167, 168, 5, 100, 0, 0, 168, 28, 1, 0, 0, 0, 169, 170, 5,
		111, 0, 0, 170, 171, 5, 118, 0, 0, 171, 172, 5, 101, 0, 0, 172, 173, 5,
		114, 0, 0, 173, 174, 5, 100, 0, 0, 174, 175, 5, 114, 0, 0, 175, 176, 5,
		97, 0, 0, 176, 177, 5, 102, 0, 0, 177, 178, 5, 116, 0, 0, 178, 30, 1, 0,
		0, 0, 179, 180, 5, 40, 0, 0, 180, 32, 1, 0, 0, 0, 181, 182, 5, 41, 0, 0,
		182, 34, 1, 0, 0, 0, 183, 184, 5, 91, 0, 0, 184, 36, 1, 0, 0, 0, 185, 186,
		5, 93, 0, 0, 186, 38, 1, 0, 0, 0, 187, 188, 5, 123, 0, 0, 188, 40, 1, 0,
		0, 0, 189, 190, 5, 125, 0, 0, 190, 42, 1, 0, 0, 0, 191, 192, 5, 61, 0,
		0, 192, 44, 1, 0, 0, 0, 193, 195, 7, 2, 0, 0, 194, 193, 1, 0, 0, 0, 195,
		196, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 199,
		1, 0, 0, 0, 198, 200, 7, 3, 0, 0, 199, 198, 1, 0, 0, 0, 199, 200, 1, 0,
		0, 0, 200, 201, 1, 0, 0, 0, 201, 203, 5, 47, 0, 0, 202, 204, 7, 3, 0, 0,
		203, 202, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 206, 1, 0, 0, 0, 205,
		207, 7, 2, 0, 0, 206, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 206,
		1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 46, 1, 0, 0, 0, 210, 212, 7, 2,
		0, 0, 211, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0,
		213, 214, 1, 0, 0, 0, 214, 221, 1, 0, 0, 0, 215, 217, 5, 46, 0, 0, 216,
		218, 7, 2, 0, 0, 217, 216, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 217,
		1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 222, 1, 0, 0, 0, 221, 215, 1, 0,
		0, 0, 221, 222, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 224, 5, 37, 0, 0,
		224, 48, 1, 0, 0, 0, 225, 227, 7, 4, 0, 0, 226, 225, 1, 0, 0, 0, 227, 228,
		1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 50, 1, 0,
		0, 0, 230, 232, 7, 2, 0, 0, 231, 230, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0,
		233, 231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 52, 1, 0, 0, 0, 235, 237,
		5, 36, 0, 0, 236, 238, 7, 5, 0, 0, 237, 236, 1, 0, 0, 0, 238, 239, 1, 0,
		0, 0, 239, 237, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 244, 1, 0, 0, 0,
		241, 243, 7, 6, 0, 0, 242, 241, 1, 0, 0, 0, 243, 246, 1, 0, 0, 0, 244,
		242, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 54, 1, 0, 0, 0, 246, 244, 1,
		0, 0, 0, 247, 249, 5, 64, 0, 0, 248, 250, 7, 7, 0, 0, 249, 248, 1, 0, 0,
		0, 250, 251, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252,
		261, 1, 0, 0, 0, 253, 255, 5, 58, 0, 0, 254, 256, 7, 7, 0, 0, 255, 254,
		1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 257, 258, 1, 0,
		0, 0, 258, 260, 1, 0, 0, 0, 259, 253, 1, 0, 0, 0, 260, 263, 1, 0, 0, 0,
		261, 259, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 56, 1, 0, 0, 0, 263, 261,
		1, 0, 0, 0, 264, 266, 7, 8, 0, 0, 265, 264, 1, 0, 0, 0, 266, 267, 1, 0,
		0, 0, 267, 265, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 58, 1, 0, 0, 0,
		21, 0, 62, 69, 76, 78, 92, 196, 199, 203, 208, 213, 219, 221, 228, 233,
		239, 244, 251, 257, 261, 267, 1, 6, 0, 0,
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

// NumscriptLexerInit initializes any static state used to implement NumscriptLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewNumscriptLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func NumscriptLexerInit() {
	staticData := &NumscriptLexerLexerStaticData
	staticData.once.Do(numscriptlexerLexerInit)
}

// NewNumscriptLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewNumscriptLexer(input antlr.CharStream) *NumscriptLexer {
	NumscriptLexerInit()
	l := new(NumscriptLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &NumscriptLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Numscript.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NumscriptLexer tokens.
const (
	NumscriptLexerWS                         = 1
	NumscriptLexerNEWLINE                    = 2
	NumscriptLexerMULTILINE_COMMENT          = 3
	NumscriptLexerLINE_COMMENT               = 4
	NumscriptLexerVARS                       = 5
	NumscriptLexerMAX                        = 6
	NumscriptLexerSOURCE                     = 7
	NumscriptLexerDESTINATION                = 8
	NumscriptLexerSEND                       = 9
	NumscriptLexerFROM                       = 10
	NumscriptLexerTO                         = 11
	NumscriptLexerREMAINING                  = 12
	NumscriptLexerALLOWING                   = 13
	NumscriptLexerUNBOUNDED                  = 14
	NumscriptLexerOVERDRAFT                  = 15
	NumscriptLexerLPARENS                    = 16
	NumscriptLexerRPARENS                    = 17
	NumscriptLexerLBRACKET                   = 18
	NumscriptLexerRBRACKET                   = 19
	NumscriptLexerLBRACE                     = 20
	NumscriptLexerRBRACE                     = 21
	NumscriptLexerEQ                         = 22
	NumscriptLexerRATIO_PORTION_LITERAL      = 23
	NumscriptLexerPERCENTAGE_PORTION_LITERAL = 24
	NumscriptLexerTYPE_IDENT                 = 25
	NumscriptLexerNUMBER                     = 26
	NumscriptLexerVARIABLE_NAME              = 27
	NumscriptLexerACCOUNT                    = 28
	NumscriptLexerASSET                      = 29
)
