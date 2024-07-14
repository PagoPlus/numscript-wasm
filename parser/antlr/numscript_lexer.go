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
		"'send'", "'from'", "'to'", "'remaining'", "'('", "')'", "'['", "']'",
		"'{'", "'}'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "NEWLINE", "MULTILINE_COMMENT", "LINE_COMMENT", "VARS", "MAX",
		"SOURCE", "DESTINATION", "SEND", "FROM", "TO", "REMAINING", "LPARENS",
		"RPARENS", "LBRACKET", "RBRACKET", "LBRACE", "RBRACE", "EQ", "RATIO_PORTION_LITERAL",
		"PERCENTAGE_PORTION_LITERAL", "TYPE_IDENT", "NUMBER", "VARIABLE_NAME",
		"ACCOUNT", "ASSET",
	}
	staticData.RuleNames = []string{
		"WS", "NEWLINE", "MULTILINE_COMMENT", "LINE_COMMENT", "VARS", "MAX",
		"SOURCE", "DESTINATION", "SEND", "FROM", "TO", "REMAINING", "LPARENS",
		"RPARENS", "LBRACKET", "RBRACKET", "LBRACE", "RBRACE", "EQ", "RATIO_PORTION_LITERAL",
		"PERCENTAGE_PORTION_LITERAL", "TYPE_IDENT", "NUMBER", "VARIABLE_NAME",
		"ACCOUNT", "ASSET",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 26, 234, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		1, 0, 4, 0, 55, 8, 0, 11, 0, 12, 0, 56, 1, 0, 1, 0, 1, 1, 4, 1, 62, 8,
		1, 11, 1, 12, 1, 63, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 71, 8, 2, 10,
		2, 12, 2, 74, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 5, 3, 85, 8, 3, 10, 3, 12, 3, 88, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 4, 19, 160, 8,
		19, 11, 19, 12, 19, 161, 1, 19, 3, 19, 165, 8, 19, 1, 19, 1, 19, 3, 19,
		169, 8, 19, 1, 19, 4, 19, 172, 8, 19, 11, 19, 12, 19, 173, 1, 20, 4, 20,
		177, 8, 20, 11, 20, 12, 20, 178, 1, 20, 1, 20, 4, 20, 183, 8, 20, 11, 20,
		12, 20, 184, 3, 20, 187, 8, 20, 1, 20, 1, 20, 1, 21, 4, 21, 192, 8, 21,
		11, 21, 12, 21, 193, 1, 22, 4, 22, 197, 8, 22, 11, 22, 12, 22, 198, 1,
		23, 1, 23, 4, 23, 203, 8, 23, 11, 23, 12, 23, 204, 1, 23, 5, 23, 208, 8,
		23, 10, 23, 12, 23, 211, 9, 23, 1, 24, 1, 24, 4, 24, 215, 8, 24, 11, 24,
		12, 24, 216, 1, 24, 1, 24, 4, 24, 221, 8, 24, 11, 24, 12, 24, 222, 5, 24,
		225, 8, 24, 10, 24, 12, 24, 228, 9, 24, 1, 25, 4, 25, 231, 8, 25, 11, 25,
		12, 25, 232, 2, 72, 86, 0, 26, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25,
		51, 26, 1, 0, 9, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 1,
		0, 48, 57, 1, 0, 32, 32, 1, 0, 97, 122, 2, 0, 95, 95, 97, 122, 3, 0, 48,
		57, 95, 95, 97, 122, 5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 2,
		0, 47, 57, 65, 90, 253, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0,
		0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0,
		0, 0, 1, 54, 1, 0, 0, 0, 3, 61, 1, 0, 0, 0, 5, 65, 1, 0, 0, 0, 7, 80, 1,
		0, 0, 0, 9, 93, 1, 0, 0, 0, 11, 98, 1, 0, 0, 0, 13, 102, 1, 0, 0, 0, 15,
		109, 1, 0, 0, 0, 17, 121, 1, 0, 0, 0, 19, 126, 1, 0, 0, 0, 21, 131, 1,
		0, 0, 0, 23, 134, 1, 0, 0, 0, 25, 144, 1, 0, 0, 0, 27, 146, 1, 0, 0, 0,
		29, 148, 1, 0, 0, 0, 31, 150, 1, 0, 0, 0, 33, 152, 1, 0, 0, 0, 35, 154,
		1, 0, 0, 0, 37, 156, 1, 0, 0, 0, 39, 159, 1, 0, 0, 0, 41, 176, 1, 0, 0,
		0, 43, 191, 1, 0, 0, 0, 45, 196, 1, 0, 0, 0, 47, 200, 1, 0, 0, 0, 49, 212,
		1, 0, 0, 0, 51, 230, 1, 0, 0, 0, 53, 55, 7, 0, 0, 0, 54, 53, 1, 0, 0, 0,
		55, 56, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 58, 1,
		0, 0, 0, 58, 59, 6, 0, 0, 0, 59, 2, 1, 0, 0, 0, 60, 62, 7, 1, 0, 0, 61,
		60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0,
		0, 64, 4, 1, 0, 0, 0, 65, 66, 5, 47, 0, 0, 66, 67, 5, 42, 0, 0, 67, 72,
		1, 0, 0, 0, 68, 71, 3, 5, 2, 0, 69, 71, 9, 0, 0, 0, 70, 68, 1, 0, 0, 0,
		70, 69, 1, 0, 0, 0, 71, 74, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 72, 70, 1,
		0, 0, 0, 73, 75, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 75, 76, 5, 42, 0, 0, 76,
		77, 5, 47, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79, 6, 2, 0, 0, 79, 6, 1, 0, 0,
		0, 80, 81, 5, 47, 0, 0, 81, 82, 5, 47, 0, 0, 82, 86, 1, 0, 0, 0, 83, 85,
		9, 0, 0, 0, 84, 83, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0,
		86, 84, 1, 0, 0, 0, 87, 89, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 90, 3,
		3, 1, 0, 90, 91, 1, 0, 0, 0, 91, 92, 6, 3, 0, 0, 92, 8, 1, 0, 0, 0, 93,
		94, 5, 118, 0, 0, 94, 95, 5, 97, 0, 0, 95, 96, 5, 114, 0, 0, 96, 97, 5,
		115, 0, 0, 97, 10, 1, 0, 0, 0, 98, 99, 5, 109, 0, 0, 99, 100, 5, 97, 0,
		0, 100, 101, 5, 120, 0, 0, 101, 12, 1, 0, 0, 0, 102, 103, 5, 115, 0, 0,
		103, 104, 5, 111, 0, 0, 104, 105, 5, 117, 0, 0, 105, 106, 5, 114, 0, 0,
		106, 107, 5, 99, 0, 0, 107, 108, 5, 101, 0, 0, 108, 14, 1, 0, 0, 0, 109,
		110, 5, 100, 0, 0, 110, 111, 5, 101, 0, 0, 111, 112, 5, 115, 0, 0, 112,
		113, 5, 116, 0, 0, 113, 114, 5, 105, 0, 0, 114, 115, 5, 110, 0, 0, 115,
		116, 5, 97, 0, 0, 116, 117, 5, 116, 0, 0, 117, 118, 5, 105, 0, 0, 118,
		119, 5, 111, 0, 0, 119, 120, 5, 110, 0, 0, 120, 16, 1, 0, 0, 0, 121, 122,
		5, 115, 0, 0, 122, 123, 5, 101, 0, 0, 123, 124, 5, 110, 0, 0, 124, 125,
		5, 100, 0, 0, 125, 18, 1, 0, 0, 0, 126, 127, 5, 102, 0, 0, 127, 128, 5,
		114, 0, 0, 128, 129, 5, 111, 0, 0, 129, 130, 5, 109, 0, 0, 130, 20, 1,
		0, 0, 0, 131, 132, 5, 116, 0, 0, 132, 133, 5, 111, 0, 0, 133, 22, 1, 0,
		0, 0, 134, 135, 5, 114, 0, 0, 135, 136, 5, 101, 0, 0, 136, 137, 5, 109,
		0, 0, 137, 138, 5, 97, 0, 0, 138, 139, 5, 105, 0, 0, 139, 140, 5, 110,
		0, 0, 140, 141, 5, 105, 0, 0, 141, 142, 5, 110, 0, 0, 142, 143, 5, 103,
		0, 0, 143, 24, 1, 0, 0, 0, 144, 145, 5, 40, 0, 0, 145, 26, 1, 0, 0, 0,
		146, 147, 5, 41, 0, 0, 147, 28, 1, 0, 0, 0, 148, 149, 5, 91, 0, 0, 149,
		30, 1, 0, 0, 0, 150, 151, 5, 93, 0, 0, 151, 32, 1, 0, 0, 0, 152, 153, 5,
		123, 0, 0, 153, 34, 1, 0, 0, 0, 154, 155, 5, 125, 0, 0, 155, 36, 1, 0,
		0, 0, 156, 157, 5, 61, 0, 0, 157, 38, 1, 0, 0, 0, 158, 160, 7, 2, 0, 0,
		159, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161,
		162, 1, 0, 0, 0, 162, 164, 1, 0, 0, 0, 163, 165, 7, 3, 0, 0, 164, 163,
		1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 168, 5, 47,
		0, 0, 167, 169, 7, 3, 0, 0, 168, 167, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0,
		169, 171, 1, 0, 0, 0, 170, 172, 7, 2, 0, 0, 171, 170, 1, 0, 0, 0, 172,
		173, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 40, 1,
		0, 0, 0, 175, 177, 7, 2, 0, 0, 176, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0,
		0, 178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 186, 1, 0, 0, 0, 180,
		182, 5, 46, 0, 0, 181, 183, 7, 2, 0, 0, 182, 181, 1, 0, 0, 0, 183, 184,
		1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 187, 1, 0,
		0, 0, 186, 180, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0,
		188, 189, 5, 37, 0, 0, 189, 42, 1, 0, 0, 0, 190, 192, 7, 4, 0, 0, 191,
		190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193, 194,
		1, 0, 0, 0, 194, 44, 1, 0, 0, 0, 195, 197, 7, 2, 0, 0, 196, 195, 1, 0,
		0, 0, 197, 198, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0,
		199, 46, 1, 0, 0, 0, 200, 202, 5, 36, 0, 0, 201, 203, 7, 5, 0, 0, 202,
		201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 204, 205,
		1, 0, 0, 0, 205, 209, 1, 0, 0, 0, 206, 208, 7, 6, 0, 0, 207, 206, 1, 0,
		0, 0, 208, 211, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0,
		210, 48, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 212, 214, 5, 64, 0, 0, 213,
		215, 7, 7, 0, 0, 214, 213, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 214,
		1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 226, 1, 0, 0, 0, 218, 220, 5, 58,
		0, 0, 219, 221, 7, 7, 0, 0, 220, 219, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0,
		222, 220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 225, 1, 0, 0, 0, 224,
		218, 1, 0, 0, 0, 225, 228, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 226, 227,
		1, 0, 0, 0, 227, 50, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 229, 231, 7, 8,
		0, 0, 230, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0,
		232, 233, 1, 0, 0, 0, 233, 52, 1, 0, 0, 0, 21, 0, 56, 63, 70, 72, 86, 161,
		164, 168, 173, 178, 184, 186, 193, 198, 204, 209, 216, 222, 226, 232, 1,
		6, 0, 0,
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
	NumscriptLexerLPARENS                    = 13
	NumscriptLexerRPARENS                    = 14
	NumscriptLexerLBRACKET                   = 15
	NumscriptLexerRBRACKET                   = 16
	NumscriptLexerLBRACE                     = 17
	NumscriptLexerRBRACE                     = 18
	NumscriptLexerEQ                         = 19
	NumscriptLexerRATIO_PORTION_LITERAL      = 20
	NumscriptLexerPERCENTAGE_PORTION_LITERAL = 21
	NumscriptLexerTYPE_IDENT                 = 22
	NumscriptLexerNUMBER                     = 23
	NumscriptLexerVARIABLE_NAME              = 24
	NumscriptLexerACCOUNT                    = 25
	NumscriptLexerASSET                      = 26
)
