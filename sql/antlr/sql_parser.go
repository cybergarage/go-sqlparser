// Code generated from SQL.g4 by ANTLR 4.12.0. DO NOT EDIT.

package antlr // SQL
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

type SQLParser struct {
	*antlr.BaseParser
}

var sqlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sqlParserInit() {
	staticData := &sqlParserStaticData
	staticData.literalNames = []string{
		"", "'('", "')'", "'{'", "'}'", "'['", "']'", "':'", "'*'", "'='", "'=='",
		"'<'", "'<='", "'>'", "'>='", "'!='", "", "", "','", "';'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "ASTERISK", "SINGLE_EQ", "DOUBLE_EQ",
		"OP_LT", "LE", "GT", "GE", "NOTEQ", "AND", "OR", "COMMA", "SEMICOLON",
		"ALL", "ANCESTOR", "AS", "ASC", "ASYNC", "BY", "CREATE", "COLLECTION",
		"CURRENT_DATE", "CURRENT_TIME", "CURRENT_TIMESTAMP", "DATABASE", "DESC",
		"DELETE", "DISTINCT", "DROP", "EACH", "EXCEPT", "FLATTEN", "FROM", "GROUP",
		"HAVING", "IN", "COLLECTION_INDEX", "INSERT", "INTERSECT", "INTO", "IS",
		"LIMIT", "NIL", "OFFSET", "OPTIONS", "ORDER", "SELECT", "SET", "SHOW",
		"SYNC", "USE", "UNION", "UPDATE", "WHERE", "VALUE", "VALUES", "TRUE",
		"FALSE", "WS", "IDENTIFIER", "NUMBER", "FLOAT", "STRING",
	}
	staticData.ruleNames = []string{
		"queries", "statement", "showStmt", "useStmt", "createStmt", "createDatabase",
		"createCollection", "drop_collectionStmt", "create_indexStmt", "drop_indexStmt",
		"selectStmt", "columns", "from_section", "table_name", "data_source",
		"grouping_section", "having_section", "sorting_section", "sorting_item",
		"sorting_specification", "limit_section", "offset_section", "insertStmt",
		"insert_columns_section", "insert_values_section", "updateStmt", "property_section",
		"deleteStmt", "expression", "expression_list", "expression_literal",
		"expression_literal_value", "expression_dictionary", "dictionary_literal",
		"expression_array", "array_literal", "expression_logic_operator", "expression_binary_operator",
		"expression_function", "function_name", "function_value", "expression_operator",
		"binary_operator", "logical_operator", "property_literal", "integer_literal",
		"real_literal", "string_literal", "true_literal", "false_literal", "sync_operator",
		"compound_operator", "condition_operator", "property", "value", "name",
		"collection_section", "collectionName", "databaseName", "column", "index_section",
		"index_name", "where_section", "unreserved_keyword",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 69, 496, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 1, 0, 1, 0, 1, 0, 5, 0, 132, 8, 0, 10, 0, 12, 0, 135, 9, 0,
		1, 0, 5, 0, 138, 8, 0, 10, 0, 12, 0, 141, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 153, 8, 1, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 163, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 174, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 3, 10, 190,
		8, 10, 1, 10, 1, 10, 1, 10, 3, 10, 195, 8, 10, 1, 10, 3, 10, 198, 8, 10,
		1, 10, 3, 10, 201, 8, 10, 1, 10, 3, 10, 204, 8, 10, 1, 10, 3, 10, 207,
		8, 10, 1, 10, 3, 10, 210, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 216,
		8, 11, 10, 11, 12, 11, 219, 9, 11, 3, 11, 221, 8, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 5, 12, 228, 8, 12, 10, 12, 12, 12, 231, 9, 12, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 242, 8,
		15, 10, 15, 12, 15, 245, 9, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 5, 17, 255, 8, 17, 10, 17, 12, 17, 258, 9, 17, 1, 18, 1,
		18, 3, 18, 262, 8, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20,
		270, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 3, 22, 277, 8, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 3, 22, 283, 8, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 5, 23, 291, 8, 23, 10, 23, 12, 23, 294, 9, 23, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 302, 8, 24, 1, 25, 3, 25, 305, 8, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 313, 8, 25, 10, 25, 12,
		25, 316, 9, 25, 1, 25, 3, 25, 319, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		27, 3, 27, 326, 8, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 332, 8, 27, 1,
		28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 342, 8, 29,
		10, 29, 12, 29, 345, 9, 29, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 351, 8,
		29, 10, 29, 12, 29, 354, 9, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		5, 29, 362, 8, 29, 10, 29, 12, 29, 365, 9, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 5, 29, 373, 8, 29, 10, 29, 12, 29, 376, 9, 29, 1, 29,
		1, 29, 1, 29, 3, 29, 381, 8, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 395, 8, 31, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1,
		35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 3, 38, 416, 8, 38,
		1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 5, 40, 425, 8, 40, 10,
		40, 12, 40, 428, 9, 40, 1, 40, 3, 40, 431, 8, 40, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1,
		46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51,
		3, 51, 457, 8, 51, 1, 51, 1, 51, 3, 51, 461, 8, 51, 1, 52, 1, 52, 1, 53,
		1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 3, 55, 472, 8, 55, 1, 56, 1,
		56, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 59, 3, 59, 483, 8, 59,
		1, 60, 1, 60, 1, 61, 1, 61, 3, 61, 489, 8, 61, 1, 62, 1, 62, 1, 62, 1,
		63, 1, 63, 1, 63, 0, 0, 64, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
		60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94,
		96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124,
		126, 0, 5, 2, 0, 20, 20, 34, 34, 2, 0, 23, 23, 32, 32, 1, 0, 9, 15, 1,
		0, 16, 17, 2, 0, 24, 24, 56, 56, 496, 0, 128, 1, 0, 0, 0, 2, 152, 1, 0,
		0, 0, 4, 154, 1, 0, 0, 0, 6, 157, 1, 0, 0, 0, 8, 162, 1, 0, 0, 0, 10, 164,
		1, 0, 0, 0, 12, 168, 1, 0, 0, 0, 14, 175, 1, 0, 0, 0, 16, 179, 1, 0, 0,
		0, 18, 183, 1, 0, 0, 0, 20, 187, 1, 0, 0, 0, 22, 220, 1, 0, 0, 0, 24, 222,
		1, 0, 0, 0, 26, 232, 1, 0, 0, 0, 28, 234, 1, 0, 0, 0, 30, 236, 1, 0, 0,
		0, 32, 246, 1, 0, 0, 0, 34, 249, 1, 0, 0, 0, 36, 259, 1, 0, 0, 0, 38, 263,
		1, 0, 0, 0, 40, 265, 1, 0, 0, 0, 42, 273, 1, 0, 0, 0, 44, 276, 1, 0, 0,
		0, 46, 286, 1, 0, 0, 0, 48, 301, 1, 0, 0, 0, 50, 304, 1, 0, 0, 0, 52, 320,
		1, 0, 0, 0, 54, 325, 1, 0, 0, 0, 56, 333, 1, 0, 0, 0, 58, 380, 1, 0, 0,
		0, 60, 382, 1, 0, 0, 0, 62, 394, 1, 0, 0, 0, 64, 396, 1, 0, 0, 0, 66, 400,
		1, 0, 0, 0, 68, 404, 1, 0, 0, 0, 70, 406, 1, 0, 0, 0, 72, 408, 1, 0, 0,
		0, 74, 410, 1, 0, 0, 0, 76, 412, 1, 0, 0, 0, 78, 419, 1, 0, 0, 0, 80, 430,
		1, 0, 0, 0, 82, 432, 1, 0, 0, 0, 84, 436, 1, 0, 0, 0, 86, 438, 1, 0, 0,
		0, 88, 440, 1, 0, 0, 0, 90, 442, 1, 0, 0, 0, 92, 444, 1, 0, 0, 0, 94, 446,
		1, 0, 0, 0, 96, 448, 1, 0, 0, 0, 98, 450, 1, 0, 0, 0, 100, 452, 1, 0, 0,
		0, 102, 460, 1, 0, 0, 0, 104, 462, 1, 0, 0, 0, 106, 464, 1, 0, 0, 0, 108,
		466, 1, 0, 0, 0, 110, 471, 1, 0, 0, 0, 112, 473, 1, 0, 0, 0, 114, 475,
		1, 0, 0, 0, 116, 477, 1, 0, 0, 0, 118, 479, 1, 0, 0, 0, 120, 484, 1, 0,
		0, 0, 122, 488, 1, 0, 0, 0, 124, 490, 1, 0, 0, 0, 126, 493, 1, 0, 0, 0,
		128, 133, 3, 2, 1, 0, 129, 130, 5, 19, 0, 0, 130, 132, 3, 2, 1, 0, 131,
		129, 1, 0, 0, 0, 132, 135, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134,
		1, 0, 0, 0, 134, 139, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 136, 138, 5, 19,
		0, 0, 137, 136, 1, 0, 0, 0, 138, 141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0,
		139, 140, 1, 0, 0, 0, 140, 1, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 153,
		3, 4, 2, 0, 143, 153, 3, 6, 3, 0, 144, 153, 3, 8, 4, 0, 145, 153, 3, 16,
		8, 0, 146, 153, 3, 14, 7, 0, 147, 153, 3, 18, 9, 0, 148, 153, 3, 20, 10,
		0, 149, 153, 3, 44, 22, 0, 150, 153, 3, 50, 25, 0, 151, 153, 3, 54, 27,
		0, 152, 142, 1, 0, 0, 0, 152, 143, 1, 0, 0, 0, 152, 144, 1, 0, 0, 0, 152,
		145, 1, 0, 0, 0, 152, 146, 1, 0, 0, 0, 152, 147, 1, 0, 0, 0, 152, 148,
		1, 0, 0, 0, 152, 149, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152, 151, 1, 0,
		0, 0, 153, 3, 1, 0, 0, 0, 154, 155, 5, 55, 0, 0, 155, 156, 3, 112, 56,
		0, 156, 5, 1, 0, 0, 0, 157, 158, 5, 57, 0, 0, 158, 159, 3, 112, 56, 0,
		159, 7, 1, 0, 0, 0, 160, 163, 3, 10, 5, 0, 161, 163, 3, 12, 6, 0, 162,
		160, 1, 0, 0, 0, 162, 161, 1, 0, 0, 0, 163, 9, 1, 0, 0, 0, 164, 165, 5,
		26, 0, 0, 165, 166, 5, 31, 0, 0, 166, 167, 3, 116, 58, 0, 167, 11, 1, 0,
		0, 0, 168, 169, 5, 26, 0, 0, 169, 170, 5, 27, 0, 0, 170, 173, 3, 112, 56,
		0, 171, 172, 5, 51, 0, 0, 172, 174, 3, 56, 28, 0, 173, 171, 1, 0, 0, 0,
		173, 174, 1, 0, 0, 0, 174, 13, 1, 0, 0, 0, 175, 176, 5, 35, 0, 0, 176,
		177, 5, 27, 0, 0, 177, 178, 3, 112, 56, 0, 178, 15, 1, 0, 0, 0, 179, 180,
		5, 26, 0, 0, 180, 181, 5, 43, 0, 0, 181, 182, 3, 120, 60, 0, 182, 17, 1,
		0, 0, 0, 183, 184, 5, 35, 0, 0, 184, 185, 5, 43, 0, 0, 185, 186, 3, 120,
		60, 0, 186, 19, 1, 0, 0, 0, 187, 189, 5, 53, 0, 0, 188, 190, 7, 0, 0, 0,
		189, 188, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191,
		192, 3, 22, 11, 0, 192, 194, 3, 24, 12, 0, 193, 195, 3, 124, 62, 0, 194,
		193, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 197, 1, 0, 0, 0, 196, 198,
		3, 30, 15, 0, 197, 196, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 200, 1,
		0, 0, 0, 199, 201, 3, 32, 16, 0, 200, 199, 1, 0, 0, 0, 200, 201, 1, 0,
		0, 0, 201, 203, 1, 0, 0, 0, 202, 204, 3, 34, 17, 0, 203, 202, 1, 0, 0,
		0, 203, 204, 1, 0, 0, 0, 204, 206, 1, 0, 0, 0, 205, 207, 3, 40, 20, 0,
		206, 205, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 209, 1, 0, 0, 0, 208,
		210, 3, 42, 21, 0, 209, 208, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 21,
		1, 0, 0, 0, 211, 221, 5, 8, 0, 0, 212, 217, 3, 118, 59, 0, 213, 214, 5,
		18, 0, 0, 214, 216, 3, 118, 59, 0, 215, 213, 1, 0, 0, 0, 216, 219, 1, 0,
		0, 0, 217, 215, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 221, 1, 0, 0, 0,
		219, 217, 1, 0, 0, 0, 220, 211, 1, 0, 0, 0, 220, 212, 1, 0, 0, 0, 221,
		23, 1, 0, 0, 0, 222, 223, 5, 39, 0, 0, 223, 224, 3, 26, 13, 0, 224, 229,
		1, 0, 0, 0, 225, 226, 5, 18, 0, 0, 226, 228, 3, 26, 13, 0, 227, 225, 1,
		0, 0, 0, 228, 231, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 229, 230, 1, 0, 0,
		0, 230, 25, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 232, 233, 3, 28, 14, 0, 233,
		27, 1, 0, 0, 0, 234, 235, 3, 114, 57, 0, 235, 29, 1, 0, 0, 0, 236, 237,
		5, 40, 0, 0, 237, 238, 5, 25, 0, 0, 238, 243, 3, 56, 28, 0, 239, 240, 5,
		18, 0, 0, 240, 242, 3, 56, 28, 0, 241, 239, 1, 0, 0, 0, 242, 245, 1, 0,
		0, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 31, 1, 0, 0, 0,
		245, 243, 1, 0, 0, 0, 246, 247, 5, 41, 0, 0, 247, 248, 3, 56, 28, 0, 248,
		33, 1, 0, 0, 0, 249, 250, 5, 52, 0, 0, 250, 251, 5, 25, 0, 0, 251, 256,
		3, 36, 18, 0, 252, 253, 5, 18, 0, 0, 253, 255, 3, 36, 18, 0, 254, 252,
		1, 0, 0, 0, 255, 258, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0, 256, 257, 1, 0,
		0, 0, 257, 35, 1, 0, 0, 0, 258, 256, 1, 0, 0, 0, 259, 261, 3, 106, 53,
		0, 260, 262, 3, 38, 19, 0, 261, 260, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0,
		262, 37, 1, 0, 0, 0, 263, 264, 7, 1, 0, 0, 264, 39, 1, 0, 0, 0, 265, 269,
		5, 48, 0, 0, 266, 267, 3, 60, 30, 0, 267, 268, 5, 18, 0, 0, 268, 270, 1,
		0, 0, 0, 269, 266, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 271, 1, 0, 0,
		0, 271, 272, 3, 60, 30, 0, 272, 41, 1, 0, 0, 0, 273, 274, 5, 50, 0, 0,
		274, 43, 1, 0, 0, 0, 275, 277, 3, 100, 50, 0, 276, 275, 1, 0, 0, 0, 276,
		277, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278, 279, 5, 44, 0, 0, 279, 280,
		5, 46, 0, 0, 280, 282, 3, 112, 56, 0, 281, 283, 3, 46, 23, 0, 282, 281,
		1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284, 285, 3, 48,
		24, 0, 285, 45, 1, 0, 0, 0, 286, 287, 5, 1, 0, 0, 287, 292, 3, 118, 59,
		0, 288, 289, 5, 18, 0, 0, 289, 291, 3, 118, 59, 0, 290, 288, 1, 0, 0, 0,
		291, 294, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293,
		295, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 295, 296, 5, 2, 0, 0, 296, 47, 1,
		0, 0, 0, 297, 298, 5, 61, 0, 0, 298, 302, 3, 56, 28, 0, 299, 300, 5, 62,
		0, 0, 300, 302, 3, 56, 28, 0, 301, 297, 1, 0, 0, 0, 301, 299, 1, 0, 0,
		0, 302, 49, 1, 0, 0, 0, 303, 305, 3, 100, 50, 0, 304, 303, 1, 0, 0, 0,
		304, 305, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306, 307, 5, 59, 0, 0, 307,
		308, 3, 112, 56, 0, 308, 309, 5, 54, 0, 0, 309, 314, 3, 52, 26, 0, 310,
		311, 5, 18, 0, 0, 311, 313, 3, 52, 26, 0, 312, 310, 1, 0, 0, 0, 313, 316,
		1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 318, 1, 0,
		0, 0, 316, 314, 1, 0, 0, 0, 317, 319, 3, 124, 62, 0, 318, 317, 1, 0, 0,
		0, 318, 319, 1, 0, 0, 0, 319, 51, 1, 0, 0, 0, 320, 321, 3, 106, 53, 0,
		321, 322, 5, 9, 0, 0, 322, 323, 3, 60, 30, 0, 323, 53, 1, 0, 0, 0, 324,
		326, 3, 100, 50, 0, 325, 324, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 327,
		1, 0, 0, 0, 327, 328, 5, 33, 0, 0, 328, 329, 5, 39, 0, 0, 329, 331, 3,
		112, 56, 0, 330, 332, 3, 124, 62, 0, 331, 330, 1, 0, 0, 0, 331, 332, 1,
		0, 0, 0, 332, 55, 1, 0, 0, 0, 333, 334, 3, 58, 29, 0, 334, 57, 1, 0, 0,
		0, 335, 381, 3, 60, 30, 0, 336, 381, 3, 76, 38, 0, 337, 343, 3, 74, 37,
		0, 338, 339, 3, 72, 36, 0, 339, 340, 3, 74, 37, 0, 340, 342, 1, 0, 0, 0,
		341, 338, 1, 0, 0, 0, 342, 345, 1, 0, 0, 0, 343, 341, 1, 0, 0, 0, 343,
		344, 1, 0, 0, 0, 344, 381, 1, 0, 0, 0, 345, 343, 1, 0, 0, 0, 346, 347,
		5, 1, 0, 0, 347, 352, 3, 68, 34, 0, 348, 349, 5, 18, 0, 0, 349, 351, 3,
		68, 34, 0, 350, 348, 1, 0, 0, 0, 351, 354, 1, 0, 0, 0, 352, 350, 1, 0,
		0, 0, 352, 353, 1, 0, 0, 0, 353, 355, 1, 0, 0, 0, 354, 352, 1, 0, 0, 0,
		355, 356, 5, 2, 0, 0, 356, 381, 1, 0, 0, 0, 357, 358, 5, 3, 0, 0, 358,
		363, 3, 64, 32, 0, 359, 360, 5, 18, 0, 0, 360, 362, 3, 64, 32, 0, 361,
		359, 1, 0, 0, 0, 362, 365, 1, 0, 0, 0, 363, 361, 1, 0, 0, 0, 363, 364,
		1, 0, 0, 0, 364, 366, 1, 0, 0, 0, 365, 363, 1, 0, 0, 0, 366, 367, 5, 4,
		0, 0, 367, 381, 1, 0, 0, 0, 368, 369, 5, 5, 0, 0, 369, 374, 3, 68, 34,
		0, 370, 371, 5, 18, 0, 0, 371, 373, 3, 68, 34, 0, 372, 370, 1, 0, 0, 0,
		373, 376, 1, 0, 0, 0, 374, 372, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375,
		377, 1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 377, 378, 5, 6, 0, 0, 378, 381,
		1, 0, 0, 0, 379, 381, 3, 126, 63, 0, 380, 335, 1, 0, 0, 0, 380, 336, 1,
		0, 0, 0, 380, 337, 1, 0, 0, 0, 380, 346, 1, 0, 0, 0, 380, 357, 1, 0, 0,
		0, 380, 368, 1, 0, 0, 0, 380, 379, 1, 0, 0, 0, 381, 59, 1, 0, 0, 0, 382,
		383, 3, 62, 31, 0, 383, 61, 1, 0, 0, 0, 384, 395, 3, 88, 44, 0, 385, 395,
		3, 90, 45, 0, 386, 395, 3, 92, 46, 0, 387, 395, 3, 94, 47, 0, 388, 395,
		3, 96, 48, 0, 389, 395, 3, 98, 49, 0, 390, 395, 5, 49, 0, 0, 391, 395,
		5, 29, 0, 0, 392, 395, 5, 28, 0, 0, 393, 395, 5, 30, 0, 0, 394, 384, 1,
		0, 0, 0, 394, 385, 1, 0, 0, 0, 394, 386, 1, 0, 0, 0, 394, 387, 1, 0, 0,
		0, 394, 388, 1, 0, 0, 0, 394, 389, 1, 0, 0, 0, 394, 390, 1, 0, 0, 0, 394,
		391, 1, 0, 0, 0, 394, 392, 1, 0, 0, 0, 394, 393, 1, 0, 0, 0, 395, 63, 1,
		0, 0, 0, 396, 397, 3, 110, 55, 0, 397, 398, 5, 7, 0, 0, 398, 399, 3, 60,
		30, 0, 399, 65, 1, 0, 0, 0, 400, 401, 3, 110, 55, 0, 401, 402, 5, 7, 0,
		0, 402, 403, 3, 60, 30, 0, 403, 67, 1, 0, 0, 0, 404, 405, 3, 60, 30, 0,
		405, 69, 1, 0, 0, 0, 406, 407, 3, 60, 30, 0, 407, 71, 1, 0, 0, 0, 408,
		409, 3, 86, 43, 0, 409, 73, 1, 0, 0, 0, 410, 411, 3, 82, 41, 0, 411, 75,
		1, 0, 0, 0, 412, 413, 5, 66, 0, 0, 413, 415, 5, 1, 0, 0, 414, 416, 3, 80,
		40, 0, 415, 414, 1, 0, 0, 0, 415, 416, 1, 0, 0, 0, 416, 417, 1, 0, 0, 0,
		417, 418, 5, 2, 0, 0, 418, 77, 1, 0, 0, 0, 419, 420, 5, 66, 0, 0, 420,
		79, 1, 0, 0, 0, 421, 426, 3, 56, 28, 0, 422, 423, 5, 18, 0, 0, 423, 425,
		3, 56, 28, 0, 424, 422, 1, 0, 0, 0, 425, 428, 1, 0, 0, 0, 426, 424, 1,
		0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 431, 1, 0, 0, 0, 428, 426, 1, 0, 0,
		0, 429, 431, 5, 8, 0, 0, 430, 421, 1, 0, 0, 0, 430, 429, 1, 0, 0, 0, 431,
		81, 1, 0, 0, 0, 432, 433, 3, 60, 30, 0, 433, 434, 3, 84, 42, 0, 434, 435,
		3, 60, 30, 0, 435, 83, 1, 0, 0, 0, 436, 437, 7, 2, 0, 0, 437, 85, 1, 0,
		0, 0, 438, 439, 7, 3, 0, 0, 439, 87, 1, 0, 0, 0, 440, 441, 5, 66, 0, 0,
		441, 89, 1, 0, 0, 0, 442, 443, 5, 67, 0, 0, 443, 91, 1, 0, 0, 0, 444, 445,
		5, 68, 0, 0, 445, 93, 1, 0, 0, 0, 446, 447, 5, 69, 0, 0, 447, 95, 1, 0,
		0, 0, 448, 449, 5, 63, 0, 0, 449, 97, 1, 0, 0, 0, 450, 451, 5, 64, 0, 0,
		451, 99, 1, 0, 0, 0, 452, 453, 7, 4, 0, 0, 453, 101, 1, 0, 0, 0, 454, 456,
		5, 58, 0, 0, 455, 457, 5, 20, 0, 0, 456, 455, 1, 0, 0, 0, 456, 457, 1,
		0, 0, 0, 457, 461, 1, 0, 0, 0, 458, 461, 5, 45, 0, 0, 459, 461, 5, 37,
		0, 0, 460, 454, 1, 0, 0, 0, 460, 458, 1, 0, 0, 0, 460, 459, 1, 0, 0, 0,
		461, 103, 1, 0, 0, 0, 462, 463, 7, 2, 0, 0, 463, 105, 1, 0, 0, 0, 464,
		465, 5, 66, 0, 0, 465, 107, 1, 0, 0, 0, 466, 467, 5, 66, 0, 0, 467, 109,
		1, 0, 0, 0, 468, 472, 5, 66, 0, 0, 469, 472, 3, 126, 63, 0, 470, 472, 1,
		0, 0, 0, 471, 468, 1, 0, 0, 0, 471, 469, 1, 0, 0, 0, 471, 470, 1, 0, 0,
		0, 472, 111, 1, 0, 0, 0, 473, 474, 3, 114, 57, 0, 474, 113, 1, 0, 0, 0,
		475, 476, 3, 110, 55, 0, 476, 115, 1, 0, 0, 0, 477, 478, 3, 110, 55, 0,
		478, 117, 1, 0, 0, 0, 479, 482, 3, 56, 28, 0, 480, 481, 5, 22, 0, 0, 481,
		483, 3, 110, 55, 0, 482, 480, 1, 0, 0, 0, 482, 483, 1, 0, 0, 0, 483, 119,
		1, 0, 0, 0, 484, 485, 3, 122, 61, 0, 485, 121, 1, 0, 0, 0, 486, 489, 5,
		66, 0, 0, 487, 489, 3, 94, 47, 0, 488, 486, 1, 0, 0, 0, 488, 487, 1, 0,
		0, 0, 489, 123, 1, 0, 0, 0, 490, 491, 5, 60, 0, 0, 491, 492, 3, 56, 28,
		0, 492, 125, 1, 0, 0, 0, 493, 494, 5, 50, 0, 0, 494, 127, 1, 0, 0, 0, 42,
		133, 139, 152, 162, 173, 189, 194, 197, 200, 203, 206, 209, 217, 220, 229,
		243, 256, 261, 269, 276, 282, 292, 301, 304, 314, 318, 325, 331, 343, 352,
		363, 374, 380, 394, 415, 426, 430, 456, 460, 471, 482, 488,
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

// SQLParserInit initializes any static state used to implement SQLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSQLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SQLParserInit() {
	staticData := &sqlParserStaticData
	staticData.once.Do(sqlParserInit)
}

// NewSQLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSQLParser(input antlr.TokenStream) *SQLParser {
	SQLParserInit()
	this := new(SQLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sqlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "SQL.g4"

	return this
}

// SQLParser tokens.
const (
	SQLParserEOF               = antlr.TokenEOF
	SQLParserT__0              = 1
	SQLParserT__1              = 2
	SQLParserT__2              = 3
	SQLParserT__3              = 4
	SQLParserT__4              = 5
	SQLParserT__5              = 6
	SQLParserT__6              = 7
	SQLParserASTERISK          = 8
	SQLParserSINGLE_EQ         = 9
	SQLParserDOUBLE_EQ         = 10
	SQLParserOP_LT             = 11
	SQLParserLE                = 12
	SQLParserGT                = 13
	SQLParserGE                = 14
	SQLParserNOTEQ             = 15
	SQLParserAND               = 16
	SQLParserOR                = 17
	SQLParserCOMMA             = 18
	SQLParserSEMICOLON         = 19
	SQLParserALL               = 20
	SQLParserANCESTOR          = 21
	SQLParserAS                = 22
	SQLParserASC               = 23
	SQLParserASYNC             = 24
	SQLParserBY                = 25
	SQLParserCREATE            = 26
	SQLParserCOLLECTION        = 27
	SQLParserCURRENT_DATE      = 28
	SQLParserCURRENT_TIME      = 29
	SQLParserCURRENT_TIMESTAMP = 30
	SQLParserDATABASE          = 31
	SQLParserDESC              = 32
	SQLParserDELETE            = 33
	SQLParserDISTINCT          = 34
	SQLParserDROP              = 35
	SQLParserEACH              = 36
	SQLParserEXCEPT            = 37
	SQLParserFLATTEN           = 38
	SQLParserFROM              = 39
	SQLParserGROUP             = 40
	SQLParserHAVING            = 41
	SQLParserIN                = 42
	SQLParserCOLLECTION_INDEX  = 43
	SQLParserINSERT            = 44
	SQLParserINTERSECT         = 45
	SQLParserINTO              = 46
	SQLParserIS                = 47
	SQLParserLIMIT             = 48
	SQLParserNIL               = 49
	SQLParserOFFSET            = 50
	SQLParserOPTIONS           = 51
	SQLParserORDER             = 52
	SQLParserSELECT            = 53
	SQLParserSET               = 54
	SQLParserSHOW              = 55
	SQLParserSYNC              = 56
	SQLParserUSE               = 57
	SQLParserUNION             = 58
	SQLParserUPDATE            = 59
	SQLParserWHERE             = 60
	SQLParserVALUE             = 61
	SQLParserVALUES            = 62
	SQLParserTRUE              = 63
	SQLParserFALSE             = 64
	SQLParserWS                = 65
	SQLParserIDENTIFIER        = 66
	SQLParserNUMBER            = 67
	SQLParserFLOAT             = 68
	SQLParserSTRING            = 69
)

// SQLParser rules.
const (
	SQLParserRULE_queries                    = 0
	SQLParserRULE_statement                  = 1
	SQLParserRULE_showStmt                   = 2
	SQLParserRULE_useStmt                    = 3
	SQLParserRULE_createStmt                 = 4
	SQLParserRULE_createDatabase             = 5
	SQLParserRULE_createCollection           = 6
	SQLParserRULE_drop_collectionStmt        = 7
	SQLParserRULE_create_indexStmt           = 8
	SQLParserRULE_drop_indexStmt             = 9
	SQLParserRULE_selectStmt                 = 10
	SQLParserRULE_columns                    = 11
	SQLParserRULE_from_section               = 12
	SQLParserRULE_table_name                 = 13
	SQLParserRULE_data_source                = 14
	SQLParserRULE_grouping_section           = 15
	SQLParserRULE_having_section             = 16
	SQLParserRULE_sorting_section            = 17
	SQLParserRULE_sorting_item               = 18
	SQLParserRULE_sorting_specification      = 19
	SQLParserRULE_limit_section              = 20
	SQLParserRULE_offset_section             = 21
	SQLParserRULE_insertStmt                 = 22
	SQLParserRULE_insert_columns_section     = 23
	SQLParserRULE_insert_values_section      = 24
	SQLParserRULE_updateStmt                 = 25
	SQLParserRULE_property_section           = 26
	SQLParserRULE_deleteStmt                 = 27
	SQLParserRULE_expression                 = 28
	SQLParserRULE_expression_list            = 29
	SQLParserRULE_expression_literal         = 30
	SQLParserRULE_expression_literal_value   = 31
	SQLParserRULE_expression_dictionary      = 32
	SQLParserRULE_dictionary_literal         = 33
	SQLParserRULE_expression_array           = 34
	SQLParserRULE_array_literal              = 35
	SQLParserRULE_expression_logic_operator  = 36
	SQLParserRULE_expression_binary_operator = 37
	SQLParserRULE_expression_function        = 38
	SQLParserRULE_function_name              = 39
	SQLParserRULE_function_value             = 40
	SQLParserRULE_expression_operator        = 41
	SQLParserRULE_binary_operator            = 42
	SQLParserRULE_logical_operator           = 43
	SQLParserRULE_property_literal           = 44
	SQLParserRULE_integer_literal            = 45
	SQLParserRULE_real_literal               = 46
	SQLParserRULE_string_literal             = 47
	SQLParserRULE_true_literal               = 48
	SQLParserRULE_false_literal              = 49
	SQLParserRULE_sync_operator              = 50
	SQLParserRULE_compound_operator          = 51
	SQLParserRULE_condition_operator         = 52
	SQLParserRULE_property                   = 53
	SQLParserRULE_value                      = 54
	SQLParserRULE_name                       = 55
	SQLParserRULE_collection_section         = 56
	SQLParserRULE_collectionName             = 57
	SQLParserRULE_databaseName               = 58
	SQLParserRULE_column                     = 59
	SQLParserRULE_index_section              = 60
	SQLParserRULE_index_name                 = 61
	SQLParserRULE_where_section              = 62
	SQLParserRULE_unreserved_keyword         = 63
)

// IQueriesContext is an interface to support dynamic dispatch.
type IQueriesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

	// IsQueriesContext differentiates from other interfaces.
	IsQueriesContext()
}

type QueriesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueriesContext() *QueriesContext {
	var p = new(QueriesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_queries
	return p
}

func (*QueriesContext) IsQueriesContext() {}

func NewQueriesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueriesContext {
	var p = new(QueriesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_queries

	return p
}

func (s *QueriesContext) GetParser() antlr.Parser { return s.parser }

func (s *QueriesContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *QueriesContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *QueriesContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(SQLParserSEMICOLON)
}

func (s *QueriesContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserSEMICOLON, i)
}

func (s *QueriesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueriesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueriesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterQueries(s)
	}
}

func (s *QueriesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitQueries(s)
	}
}

func (p *SQLParser) Queries() (localctx IQueriesContext) {
	this := p
	_ = this

	localctx = NewQueriesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SQLParserRULE_queries)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Statement()
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(129)
				p.Match(SQLParserSEMICOLON)
			}
			{
				p.SetState(130)
				p.Statement()
			}

		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserSEMICOLON {
		{
			p.SetState(136)
			p.Match(SQLParserSEMICOLON)
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ShowStmt() IShowStmtContext
	UseStmt() IUseStmtContext
	CreateStmt() ICreateStmtContext
	Create_indexStmt() ICreate_indexStmtContext
	Drop_collectionStmt() IDrop_collectionStmtContext
	Drop_indexStmt() IDrop_indexStmtContext
	SelectStmt() ISelectStmtContext
	InsertStmt() IInsertStmtContext
	UpdateStmt() IUpdateStmtContext
	DeleteStmt() IDeleteStmtContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ShowStmt() IShowStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShowStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShowStmtContext)
}

func (s *StatementContext) UseStmt() IUseStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUseStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUseStmtContext)
}

func (s *StatementContext) CreateStmt() ICreateStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateStmtContext)
}

func (s *StatementContext) Create_indexStmt() ICreate_indexStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreate_indexStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreate_indexStmtContext)
}

func (s *StatementContext) Drop_collectionStmt() IDrop_collectionStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDrop_collectionStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDrop_collectionStmtContext)
}

func (s *StatementContext) Drop_indexStmt() IDrop_indexStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDrop_indexStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDrop_indexStmtContext)
}

func (s *StatementContext) SelectStmt() ISelectStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectStmtContext)
}

func (s *StatementContext) InsertStmt() IInsertStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInsertStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInsertStmtContext)
}

func (s *StatementContext) UpdateStmt() IUpdateStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdateStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdateStmtContext)
}

func (s *StatementContext) DeleteStmt() IDeleteStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeleteStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeleteStmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *SQLParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SQLParserRULE_statement)

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

	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.ShowStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)
			p.UseStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(144)
			p.CreateStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(145)
			p.Create_indexStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(146)
			p.Drop_collectionStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(147)
			p.Drop_indexStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(148)
			p.SelectStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(149)
			p.InsertStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(150)
			p.UpdateStmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(151)
			p.DeleteStmt()
		}

	}

	return localctx
}

// IShowStmtContext is an interface to support dynamic dispatch.
type IShowStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SHOW() antlr.TerminalNode
	Collection_section() ICollection_sectionContext

	// IsShowStmtContext differentiates from other interfaces.
	IsShowStmtContext()
}

type ShowStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShowStmtContext() *ShowStmtContext {
	var p = new(ShowStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_showStmt
	return p
}

func (*ShowStmtContext) IsShowStmtContext() {}

func NewShowStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShowStmtContext {
	var p = new(ShowStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_showStmt

	return p
}

func (s *ShowStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ShowStmtContext) SHOW() antlr.TerminalNode {
	return s.GetToken(SQLParserSHOW, 0)
}

func (s *ShowStmtContext) Collection_section() ICollection_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollection_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollection_sectionContext)
}

func (s *ShowStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShowStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShowStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterShowStmt(s)
	}
}

func (s *ShowStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitShowStmt(s)
	}
}

func (p *SQLParser) ShowStmt() (localctx IShowStmtContext) {
	this := p
	_ = this

	localctx = NewShowStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SQLParserRULE_showStmt)

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
		p.SetState(154)
		p.Match(SQLParserSHOW)
	}
	{
		p.SetState(155)
		p.Collection_section()
	}

	return localctx
}

// IUseStmtContext is an interface to support dynamic dispatch.
type IUseStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	USE() antlr.TerminalNode
	Collection_section() ICollection_sectionContext

	// IsUseStmtContext differentiates from other interfaces.
	IsUseStmtContext()
}

type UseStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUseStmtContext() *UseStmtContext {
	var p = new(UseStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_useStmt
	return p
}

func (*UseStmtContext) IsUseStmtContext() {}

func NewUseStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UseStmtContext {
	var p = new(UseStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_useStmt

	return p
}

func (s *UseStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *UseStmtContext) USE() antlr.TerminalNode {
	return s.GetToken(SQLParserUSE, 0)
}

func (s *UseStmtContext) Collection_section() ICollection_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollection_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollection_sectionContext)
}

func (s *UseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UseStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterUseStmt(s)
	}
}

func (s *UseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitUseStmt(s)
	}
}

func (p *SQLParser) UseStmt() (localctx IUseStmtContext) {
	this := p
	_ = this

	localctx = NewUseStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SQLParserRULE_useStmt)

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
		p.SetState(157)
		p.Match(SQLParserUSE)
	}
	{
		p.SetState(158)
		p.Collection_section()
	}

	return localctx
}

// ICreateStmtContext is an interface to support dynamic dispatch.
type ICreateStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CreateDatabase() ICreateDatabaseContext
	CreateCollection() ICreateCollectionContext

	// IsCreateStmtContext differentiates from other interfaces.
	IsCreateStmtContext()
}

type CreateStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateStmtContext() *CreateStmtContext {
	var p = new(CreateStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_createStmt
	return p
}

func (*CreateStmtContext) IsCreateStmtContext() {}

func NewCreateStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateStmtContext {
	var p = new(CreateStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_createStmt

	return p
}

func (s *CreateStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateStmtContext) CreateDatabase() ICreateDatabaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateDatabaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateDatabaseContext)
}

func (s *CreateStmtContext) CreateCollection() ICreateCollectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateCollectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateCollectionContext)
}

func (s *CreateStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterCreateStmt(s)
	}
}

func (s *CreateStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitCreateStmt(s)
	}
}

func (p *SQLParser) CreateStmt() (localctx ICreateStmtContext) {
	this := p
	_ = this

	localctx = NewCreateStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SQLParserRULE_createStmt)

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

	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.CreateDatabase()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)
			p.CreateCollection()
		}

	}

	return localctx
}

// ICreateDatabaseContext is an interface to support dynamic dispatch.
type ICreateDatabaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CREATE() antlr.TerminalNode
	DATABASE() antlr.TerminalNode
	DatabaseName() IDatabaseNameContext

	// IsCreateDatabaseContext differentiates from other interfaces.
	IsCreateDatabaseContext()
}

type CreateDatabaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateDatabaseContext() *CreateDatabaseContext {
	var p = new(CreateDatabaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_createDatabase
	return p
}

func (*CreateDatabaseContext) IsCreateDatabaseContext() {}

func NewCreateDatabaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateDatabaseContext {
	var p = new(CreateDatabaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_createDatabase

	return p
}

func (s *CreateDatabaseContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateDatabaseContext) CREATE() antlr.TerminalNode {
	return s.GetToken(SQLParserCREATE, 0)
}

func (s *CreateDatabaseContext) DATABASE() antlr.TerminalNode {
	return s.GetToken(SQLParserDATABASE, 0)
}

func (s *CreateDatabaseContext) DatabaseName() IDatabaseNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatabaseNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatabaseNameContext)
}

func (s *CreateDatabaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateDatabaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateDatabaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterCreateDatabase(s)
	}
}

func (s *CreateDatabaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitCreateDatabase(s)
	}
}

func (p *SQLParser) CreateDatabase() (localctx ICreateDatabaseContext) {
	this := p
	_ = this

	localctx = NewCreateDatabaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SQLParserRULE_createDatabase)

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
		p.SetState(164)
		p.Match(SQLParserCREATE)
	}
	{
		p.SetState(165)
		p.Match(SQLParserDATABASE)
	}
	{
		p.SetState(166)
		p.DatabaseName()
	}

	return localctx
}

// ICreateCollectionContext is an interface to support dynamic dispatch.
type ICreateCollectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CREATE() antlr.TerminalNode
	COLLECTION() antlr.TerminalNode
	Collection_section() ICollection_sectionContext
	OPTIONS() antlr.TerminalNode
	Expression() IExpressionContext

	// IsCreateCollectionContext differentiates from other interfaces.
	IsCreateCollectionContext()
}

type CreateCollectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateCollectionContext() *CreateCollectionContext {
	var p = new(CreateCollectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_createCollection
	return p
}

func (*CreateCollectionContext) IsCreateCollectionContext() {}

func NewCreateCollectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateCollectionContext {
	var p = new(CreateCollectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_createCollection

	return p
}

func (s *CreateCollectionContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateCollectionContext) CREATE() antlr.TerminalNode {
	return s.GetToken(SQLParserCREATE, 0)
}

func (s *CreateCollectionContext) COLLECTION() antlr.TerminalNode {
	return s.GetToken(SQLParserCOLLECTION, 0)
}

func (s *CreateCollectionContext) Collection_section() ICollection_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollection_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollection_sectionContext)
}

func (s *CreateCollectionContext) OPTIONS() antlr.TerminalNode {
	return s.GetToken(SQLParserOPTIONS, 0)
}

func (s *CreateCollectionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CreateCollectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateCollectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateCollectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterCreateCollection(s)
	}
}

func (s *CreateCollectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitCreateCollection(s)
	}
}

func (p *SQLParser) CreateCollection() (localctx ICreateCollectionContext) {
	this := p
	_ = this

	localctx = NewCreateCollectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SQLParserRULE_createCollection)
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
		p.SetState(168)
		p.Match(SQLParserCREATE)
	}
	{
		p.SetState(169)
		p.Match(SQLParserCOLLECTION)
	}
	{
		p.SetState(170)
		p.Collection_section()
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserOPTIONS {
		{
			p.SetState(171)
			p.Match(SQLParserOPTIONS)
		}
		{
			p.SetState(172)
			p.Expression()
		}

	}

	return localctx
}

// IDrop_collectionStmtContext is an interface to support dynamic dispatch.
type IDrop_collectionStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DROP() antlr.TerminalNode
	COLLECTION() antlr.TerminalNode
	Collection_section() ICollection_sectionContext

	// IsDrop_collectionStmtContext differentiates from other interfaces.
	IsDrop_collectionStmtContext()
}

type Drop_collectionStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDrop_collectionStmtContext() *Drop_collectionStmtContext {
	var p = new(Drop_collectionStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_drop_collectionStmt
	return p
}

func (*Drop_collectionStmtContext) IsDrop_collectionStmtContext() {}

func NewDrop_collectionStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Drop_collectionStmtContext {
	var p = new(Drop_collectionStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_drop_collectionStmt

	return p
}

func (s *Drop_collectionStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Drop_collectionStmtContext) DROP() antlr.TerminalNode {
	return s.GetToken(SQLParserDROP, 0)
}

func (s *Drop_collectionStmtContext) COLLECTION() antlr.TerminalNode {
	return s.GetToken(SQLParserCOLLECTION, 0)
}

func (s *Drop_collectionStmtContext) Collection_section() ICollection_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollection_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollection_sectionContext)
}

func (s *Drop_collectionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Drop_collectionStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Drop_collectionStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterDrop_collectionStmt(s)
	}
}

func (s *Drop_collectionStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitDrop_collectionStmt(s)
	}
}

func (p *SQLParser) Drop_collectionStmt() (localctx IDrop_collectionStmtContext) {
	this := p
	_ = this

	localctx = NewDrop_collectionStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SQLParserRULE_drop_collectionStmt)

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
		p.SetState(175)
		p.Match(SQLParserDROP)
	}
	{
		p.SetState(176)
		p.Match(SQLParserCOLLECTION)
	}
	{
		p.SetState(177)
		p.Collection_section()
	}

	return localctx
}

// ICreate_indexStmtContext is an interface to support dynamic dispatch.
type ICreate_indexStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CREATE() antlr.TerminalNode
	COLLECTION_INDEX() antlr.TerminalNode
	Index_section() IIndex_sectionContext

	// IsCreate_indexStmtContext differentiates from other interfaces.
	IsCreate_indexStmtContext()
}

type Create_indexStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreate_indexStmtContext() *Create_indexStmtContext {
	var p = new(Create_indexStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_create_indexStmt
	return p
}

func (*Create_indexStmtContext) IsCreate_indexStmtContext() {}

func NewCreate_indexStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Create_indexStmtContext {
	var p = new(Create_indexStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_create_indexStmt

	return p
}

func (s *Create_indexStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Create_indexStmtContext) CREATE() antlr.TerminalNode {
	return s.GetToken(SQLParserCREATE, 0)
}

func (s *Create_indexStmtContext) COLLECTION_INDEX() antlr.TerminalNode {
	return s.GetToken(SQLParserCOLLECTION_INDEX, 0)
}

func (s *Create_indexStmtContext) Index_section() IIndex_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndex_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndex_sectionContext)
}

func (s *Create_indexStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Create_indexStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Create_indexStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterCreate_indexStmt(s)
	}
}

func (s *Create_indexStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitCreate_indexStmt(s)
	}
}

func (p *SQLParser) Create_indexStmt() (localctx ICreate_indexStmtContext) {
	this := p
	_ = this

	localctx = NewCreate_indexStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SQLParserRULE_create_indexStmt)

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
		p.SetState(179)
		p.Match(SQLParserCREATE)
	}
	{
		p.SetState(180)
		p.Match(SQLParserCOLLECTION_INDEX)
	}
	{
		p.SetState(181)
		p.Index_section()
	}

	return localctx
}

// IDrop_indexStmtContext is an interface to support dynamic dispatch.
type IDrop_indexStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DROP() antlr.TerminalNode
	COLLECTION_INDEX() antlr.TerminalNode
	Index_section() IIndex_sectionContext

	// IsDrop_indexStmtContext differentiates from other interfaces.
	IsDrop_indexStmtContext()
}

type Drop_indexStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDrop_indexStmtContext() *Drop_indexStmtContext {
	var p = new(Drop_indexStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_drop_indexStmt
	return p
}

func (*Drop_indexStmtContext) IsDrop_indexStmtContext() {}

func NewDrop_indexStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Drop_indexStmtContext {
	var p = new(Drop_indexStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_drop_indexStmt

	return p
}

func (s *Drop_indexStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Drop_indexStmtContext) DROP() antlr.TerminalNode {
	return s.GetToken(SQLParserDROP, 0)
}

func (s *Drop_indexStmtContext) COLLECTION_INDEX() antlr.TerminalNode {
	return s.GetToken(SQLParserCOLLECTION_INDEX, 0)
}

func (s *Drop_indexStmtContext) Index_section() IIndex_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndex_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndex_sectionContext)
}

func (s *Drop_indexStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Drop_indexStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Drop_indexStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterDrop_indexStmt(s)
	}
}

func (s *Drop_indexStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitDrop_indexStmt(s)
	}
}

func (p *SQLParser) Drop_indexStmt() (localctx IDrop_indexStmtContext) {
	this := p
	_ = this

	localctx = NewDrop_indexStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SQLParserRULE_drop_indexStmt)

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
		p.SetState(183)
		p.Match(SQLParserDROP)
	}
	{
		p.SetState(184)
		p.Match(SQLParserCOLLECTION_INDEX)
	}
	{
		p.SetState(185)
		p.Index_section()
	}

	return localctx
}

// ISelectStmtContext is an interface to support dynamic dispatch.
type ISelectStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SELECT() antlr.TerminalNode
	Columns() IColumnsContext
	From_section() IFrom_sectionContext
	Where_section() IWhere_sectionContext
	Grouping_section() IGrouping_sectionContext
	Having_section() IHaving_sectionContext
	Sorting_section() ISorting_sectionContext
	Limit_section() ILimit_sectionContext
	Offset_section() IOffset_sectionContext
	DISTINCT() antlr.TerminalNode
	ALL() antlr.TerminalNode

	// IsSelectStmtContext differentiates from other interfaces.
	IsSelectStmtContext()
}

type SelectStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectStmtContext() *SelectStmtContext {
	var p = new(SelectStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_selectStmt
	return p
}

func (*SelectStmtContext) IsSelectStmtContext() {}

func NewSelectStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStmtContext {
	var p = new(SelectStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_selectStmt

	return p
}

func (s *SelectStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectStmtContext) SELECT() antlr.TerminalNode {
	return s.GetToken(SQLParserSELECT, 0)
}

func (s *SelectStmtContext) Columns() IColumnsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnsContext)
}

func (s *SelectStmtContext) From_section() IFrom_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrom_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrom_sectionContext)
}

func (s *SelectStmtContext) Where_section() IWhere_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhere_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhere_sectionContext)
}

func (s *SelectStmtContext) Grouping_section() IGrouping_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGrouping_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGrouping_sectionContext)
}

func (s *SelectStmtContext) Having_section() IHaving_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHaving_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHaving_sectionContext)
}

func (s *SelectStmtContext) Sorting_section() ISorting_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISorting_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISorting_sectionContext)
}

func (s *SelectStmtContext) Limit_section() ILimit_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimit_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimit_sectionContext)
}

func (s *SelectStmtContext) Offset_section() IOffset_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOffset_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOffset_sectionContext)
}

func (s *SelectStmtContext) DISTINCT() antlr.TerminalNode {
	return s.GetToken(SQLParserDISTINCT, 0)
}

func (s *SelectStmtContext) ALL() antlr.TerminalNode {
	return s.GetToken(SQLParserALL, 0)
}

func (s *SelectStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterSelectStmt(s)
	}
}

func (s *SelectStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitSelectStmt(s)
	}
}

func (p *SQLParser) SelectStmt() (localctx ISelectStmtContext) {
	this := p
	_ = this

	localctx = NewSelectStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SQLParserRULE_selectStmt)
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
		p.SetState(187)
		p.Match(SQLParserSELECT)
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserALL || _la == SQLParserDISTINCT {
		{
			p.SetState(188)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SQLParserALL || _la == SQLParserDISTINCT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(191)
		p.Columns()
	}
	{
		p.SetState(192)
		p.From_section()
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserWHERE {
		{
			p.SetState(193)
			p.Where_section()
		}

	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserGROUP {
		{
			p.SetState(196)
			p.Grouping_section()
		}

	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserHAVING {
		{
			p.SetState(199)
			p.Having_section()
		}

	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserORDER {
		{
			p.SetState(202)
			p.Sorting_section()
		}

	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserLIMIT {
		{
			p.SetState(205)
			p.Limit_section()
		}

	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserOFFSET {
		{
			p.SetState(208)
			p.Offset_section()
		}

	}

	return localctx
}

// IColumnsContext is an interface to support dynamic dispatch.
type IColumnsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASTERISK() antlr.TerminalNode
	AllColumn() []IColumnContext
	Column(i int) IColumnContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsColumnsContext differentiates from other interfaces.
	IsColumnsContext()
}

type ColumnsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnsContext() *ColumnsContext {
	var p = new(ColumnsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_columns
	return p
}

func (*ColumnsContext) IsColumnsContext() {}

func NewColumnsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnsContext {
	var p = new(ColumnsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_columns

	return p
}

func (s *ColumnsContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnsContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(SQLParserASTERISK, 0)
}

func (s *ColumnsContext) AllColumn() []IColumnContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumnContext); ok {
			len++
		}
	}

	tst := make([]IColumnContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumnContext); ok {
			tst[i] = t.(IColumnContext)
			i++
		}
	}

	return tst
}

func (s *ColumnsContext) Column(i int) IColumnContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnContext); ok {
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

	return t.(IColumnContext)
}

func (s *ColumnsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLParserCOMMA)
}

func (s *ColumnsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *ColumnsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterColumns(s)
	}
}

func (s *ColumnsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitColumns(s)
	}
}

func (p *SQLParser) Columns() (localctx IColumnsContext) {
	this := p
	_ = this

	localctx = NewColumnsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SQLParserRULE_columns)
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

	p.SetState(220)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLParserASTERISK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.Match(SQLParserASTERISK)
		}

	case SQLParserT__0, SQLParserT__2, SQLParserT__4, SQLParserCURRENT_DATE, SQLParserCURRENT_TIME, SQLParserCURRENT_TIMESTAMP, SQLParserNIL, SQLParserOFFSET, SQLParserTRUE, SQLParserFALSE, SQLParserIDENTIFIER, SQLParserNUMBER, SQLParserFLOAT, SQLParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.Column()
		}
		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLParserCOMMA {
			{
				p.SetState(213)
				p.Match(SQLParserCOMMA)
			}
			{
				p.SetState(214)
				p.Column()
			}

			p.SetState(219)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFrom_sectionContext is an interface to support dynamic dispatch.
type IFrom_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FROM() antlr.TerminalNode
	AllTable_name() []ITable_nameContext
	Table_name(i int) ITable_nameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFrom_sectionContext differentiates from other interfaces.
	IsFrom_sectionContext()
}

type From_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrom_sectionContext() *From_sectionContext {
	var p = new(From_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_from_section
	return p
}

func (*From_sectionContext) IsFrom_sectionContext() {}

func NewFrom_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *From_sectionContext {
	var p = new(From_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_from_section

	return p
}

func (s *From_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *From_sectionContext) FROM() antlr.TerminalNode {
	return s.GetToken(SQLParserFROM, 0)
}

func (s *From_sectionContext) AllTable_name() []ITable_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITable_nameContext); ok {
			len++
		}
	}

	tst := make([]ITable_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITable_nameContext); ok {
			tst[i] = t.(ITable_nameContext)
			i++
		}
	}

	return tst
}

func (s *From_sectionContext) Table_name(i int) ITable_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_nameContext); ok {
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

	return t.(ITable_nameContext)
}

func (s *From_sectionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLParserCOMMA)
}

func (s *From_sectionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *From_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *From_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *From_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterFrom_section(s)
	}
}

func (s *From_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitFrom_section(s)
	}
}

func (p *SQLParser) From_section() (localctx IFrom_sectionContext) {
	this := p
	_ = this

	localctx = NewFrom_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SQLParserRULE_from_section)
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
		p.SetState(222)
		p.Match(SQLParserFROM)
	}
	{
		p.SetState(223)
		p.Table_name()
	}

	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserCOMMA {
		{
			p.SetState(225)
			p.Match(SQLParserCOMMA)
		}
		{
			p.SetState(226)
			p.Table_name()
		}

		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITable_nameContext is an interface to support dynamic dispatch.
type ITable_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Data_source() IData_sourceContext

	// IsTable_nameContext differentiates from other interfaces.
	IsTable_nameContext()
}

type Table_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_nameContext() *Table_nameContext {
	var p = new(Table_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_table_name
	return p
}

func (*Table_nameContext) IsTable_nameContext() {}

func NewTable_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_nameContext {
	var p = new(Table_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_table_name

	return p
}

func (s *Table_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_nameContext) Data_source() IData_sourceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IData_sourceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IData_sourceContext)
}

func (s *Table_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterTable_name(s)
	}
}

func (s *Table_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitTable_name(s)
	}
}

func (p *SQLParser) Table_name() (localctx ITable_nameContext) {
	this := p
	_ = this

	localctx = NewTable_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SQLParserRULE_table_name)

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
		p.SetState(232)
		p.Data_source()
	}

	return localctx
}

// IData_sourceContext is an interface to support dynamic dispatch.
type IData_sourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CollectionName() ICollectionNameContext

	// IsData_sourceContext differentiates from other interfaces.
	IsData_sourceContext()
}

type Data_sourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyData_sourceContext() *Data_sourceContext {
	var p = new(Data_sourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_data_source
	return p
}

func (*Data_sourceContext) IsData_sourceContext() {}

func NewData_sourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_sourceContext {
	var p = new(Data_sourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_data_source

	return p
}

func (s *Data_sourceContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_sourceContext) CollectionName() ICollectionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollectionNameContext)
}

func (s *Data_sourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_sourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_sourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterData_source(s)
	}
}

func (s *Data_sourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitData_source(s)
	}
}

func (p *SQLParser) Data_source() (localctx IData_sourceContext) {
	this := p
	_ = this

	localctx = NewData_sourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SQLParserRULE_data_source)

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
		p.SetState(234)
		p.CollectionName()
	}

	return localctx
}

// IGrouping_sectionContext is an interface to support dynamic dispatch.
type IGrouping_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GROUP() antlr.TerminalNode
	BY() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsGrouping_sectionContext differentiates from other interfaces.
	IsGrouping_sectionContext()
}

type Grouping_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGrouping_sectionContext() *Grouping_sectionContext {
	var p = new(Grouping_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_grouping_section
	return p
}

func (*Grouping_sectionContext) IsGrouping_sectionContext() {}

func NewGrouping_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Grouping_sectionContext {
	var p = new(Grouping_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_grouping_section

	return p
}

func (s *Grouping_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Grouping_sectionContext) GROUP() antlr.TerminalNode {
	return s.GetToken(SQLParserGROUP, 0)
}

func (s *Grouping_sectionContext) BY() antlr.TerminalNode {
	return s.GetToken(SQLParserBY, 0)
}

func (s *Grouping_sectionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Grouping_sectionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *Grouping_sectionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLParserCOMMA)
}

func (s *Grouping_sectionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *Grouping_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Grouping_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Grouping_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterGrouping_section(s)
	}
}

func (s *Grouping_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitGrouping_section(s)
	}
}

func (p *SQLParser) Grouping_section() (localctx IGrouping_sectionContext) {
	this := p
	_ = this

	localctx = NewGrouping_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SQLParserRULE_grouping_section)
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
		p.SetState(236)
		p.Match(SQLParserGROUP)
	}
	{
		p.SetState(237)
		p.Match(SQLParserBY)
	}
	{
		p.SetState(238)
		p.Expression()
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserCOMMA {
		{
			p.SetState(239)
			p.Match(SQLParserCOMMA)
		}
		{
			p.SetState(240)
			p.Expression()
		}

		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHaving_sectionContext is an interface to support dynamic dispatch.
type IHaving_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HAVING() antlr.TerminalNode
	Expression() IExpressionContext

	// IsHaving_sectionContext differentiates from other interfaces.
	IsHaving_sectionContext()
}

type Having_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHaving_sectionContext() *Having_sectionContext {
	var p = new(Having_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_having_section
	return p
}

func (*Having_sectionContext) IsHaving_sectionContext() {}

func NewHaving_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Having_sectionContext {
	var p = new(Having_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_having_section

	return p
}

func (s *Having_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Having_sectionContext) HAVING() antlr.TerminalNode {
	return s.GetToken(SQLParserHAVING, 0)
}

func (s *Having_sectionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Having_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Having_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Having_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterHaving_section(s)
	}
}

func (s *Having_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitHaving_section(s)
	}
}

func (p *SQLParser) Having_section() (localctx IHaving_sectionContext) {
	this := p
	_ = this

	localctx = NewHaving_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SQLParserRULE_having_section)

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
		p.SetState(246)
		p.Match(SQLParserHAVING)
	}
	{
		p.SetState(247)
		p.Expression()
	}

	return localctx
}

// ISorting_sectionContext is an interface to support dynamic dispatch.
type ISorting_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ORDER() antlr.TerminalNode
	BY() antlr.TerminalNode
	AllSorting_item() []ISorting_itemContext
	Sorting_item(i int) ISorting_itemContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsSorting_sectionContext differentiates from other interfaces.
	IsSorting_sectionContext()
}

type Sorting_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySorting_sectionContext() *Sorting_sectionContext {
	var p = new(Sorting_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_sorting_section
	return p
}

func (*Sorting_sectionContext) IsSorting_sectionContext() {}

func NewSorting_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sorting_sectionContext {
	var p = new(Sorting_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_sorting_section

	return p
}

func (s *Sorting_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Sorting_sectionContext) ORDER() antlr.TerminalNode {
	return s.GetToken(SQLParserORDER, 0)
}

func (s *Sorting_sectionContext) BY() antlr.TerminalNode {
	return s.GetToken(SQLParserBY, 0)
}

func (s *Sorting_sectionContext) AllSorting_item() []ISorting_itemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISorting_itemContext); ok {
			len++
		}
	}

	tst := make([]ISorting_itemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISorting_itemContext); ok {
			tst[i] = t.(ISorting_itemContext)
			i++
		}
	}

	return tst
}

func (s *Sorting_sectionContext) Sorting_item(i int) ISorting_itemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISorting_itemContext); ok {
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

	return t.(ISorting_itemContext)
}

func (s *Sorting_sectionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLParserCOMMA)
}

func (s *Sorting_sectionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *Sorting_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sorting_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sorting_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterSorting_section(s)
	}
}

func (s *Sorting_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitSorting_section(s)
	}
}

func (p *SQLParser) Sorting_section() (localctx ISorting_sectionContext) {
	this := p
	_ = this

	localctx = NewSorting_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SQLParserRULE_sorting_section)
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
		p.SetState(249)
		p.Match(SQLParserORDER)
	}
	{
		p.SetState(250)
		p.Match(SQLParserBY)
	}
	{
		p.SetState(251)
		p.Sorting_item()
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserCOMMA {
		{
			p.SetState(252)
			p.Match(SQLParserCOMMA)
		}
		{
			p.SetState(253)
			p.Sorting_item()
		}

		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISorting_itemContext is an interface to support dynamic dispatch.
type ISorting_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Property() IPropertyContext
	Sorting_specification() ISorting_specificationContext

	// IsSorting_itemContext differentiates from other interfaces.
	IsSorting_itemContext()
}

type Sorting_itemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySorting_itemContext() *Sorting_itemContext {
	var p = new(Sorting_itemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_sorting_item
	return p
}

func (*Sorting_itemContext) IsSorting_itemContext() {}

func NewSorting_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sorting_itemContext {
	var p = new(Sorting_itemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_sorting_item

	return p
}

func (s *Sorting_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Sorting_itemContext) Property() IPropertyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *Sorting_itemContext) Sorting_specification() ISorting_specificationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISorting_specificationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISorting_specificationContext)
}

func (s *Sorting_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sorting_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sorting_itemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterSorting_item(s)
	}
}

func (s *Sorting_itemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitSorting_item(s)
	}
}

func (p *SQLParser) Sorting_item() (localctx ISorting_itemContext) {
	this := p
	_ = this

	localctx = NewSorting_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SQLParserRULE_sorting_item)
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
		p.SetState(259)
		p.Property()
	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserASC || _la == SQLParserDESC {
		{
			p.SetState(260)
			p.Sorting_specification()
		}

	}

	return localctx
}

// ISorting_specificationContext is an interface to support dynamic dispatch.
type ISorting_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASC() antlr.TerminalNode
	DESC() antlr.TerminalNode

	// IsSorting_specificationContext differentiates from other interfaces.
	IsSorting_specificationContext()
}

type Sorting_specificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySorting_specificationContext() *Sorting_specificationContext {
	var p = new(Sorting_specificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_sorting_specification
	return p
}

func (*Sorting_specificationContext) IsSorting_specificationContext() {}

func NewSorting_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sorting_specificationContext {
	var p = new(Sorting_specificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_sorting_specification

	return p
}

func (s *Sorting_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *Sorting_specificationContext) ASC() antlr.TerminalNode {
	return s.GetToken(SQLParserASC, 0)
}

func (s *Sorting_specificationContext) DESC() antlr.TerminalNode {
	return s.GetToken(SQLParserDESC, 0)
}

func (s *Sorting_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sorting_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sorting_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterSorting_specification(s)
	}
}

func (s *Sorting_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitSorting_specification(s)
	}
}

func (p *SQLParser) Sorting_specification() (localctx ISorting_specificationContext) {
	this := p
	_ = this

	localctx = NewSorting_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SQLParserRULE_sorting_specification)
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
		p.SetState(263)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SQLParserASC || _la == SQLParserDESC) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILimit_sectionContext is an interface to support dynamic dispatch.
type ILimit_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LIMIT() antlr.TerminalNode
	AllExpression_literal() []IExpression_literalContext
	Expression_literal(i int) IExpression_literalContext
	COMMA() antlr.TerminalNode

	// IsLimit_sectionContext differentiates from other interfaces.
	IsLimit_sectionContext()
}

type Limit_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimit_sectionContext() *Limit_sectionContext {
	var p = new(Limit_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_limit_section
	return p
}

func (*Limit_sectionContext) IsLimit_sectionContext() {}

func NewLimit_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Limit_sectionContext {
	var p = new(Limit_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_limit_section

	return p
}

func (s *Limit_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Limit_sectionContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(SQLParserLIMIT, 0)
}

func (s *Limit_sectionContext) AllExpression_literal() []IExpression_literalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpression_literalContext); ok {
			len++
		}
	}

	tst := make([]IExpression_literalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpression_literalContext); ok {
			tst[i] = t.(IExpression_literalContext)
			i++
		}
	}

	return tst
}

func (s *Limit_sectionContext) Expression_literal(i int) IExpression_literalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_literalContext); ok {
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

	return t.(IExpression_literalContext)
}

func (s *Limit_sectionContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, 0)
}

func (s *Limit_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Limit_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Limit_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterLimit_section(s)
	}
}

func (s *Limit_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitLimit_section(s)
	}
}

func (p *SQLParser) Limit_section() (localctx ILimit_sectionContext) {
	this := p
	_ = this

	localctx = NewLimit_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SQLParserRULE_limit_section)

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
		p.SetState(265)
		p.Match(SQLParserLIMIT)
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(266)
			p.Expression_literal()
		}
		{
			p.SetState(267)
			p.Match(SQLParserCOMMA)
		}

	}
	{
		p.SetState(271)
		p.Expression_literal()
	}

	return localctx
}

// IOffset_sectionContext is an interface to support dynamic dispatch.
type IOffset_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OFFSET() antlr.TerminalNode

	// IsOffset_sectionContext differentiates from other interfaces.
	IsOffset_sectionContext()
}

type Offset_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOffset_sectionContext() *Offset_sectionContext {
	var p = new(Offset_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_offset_section
	return p
}

func (*Offset_sectionContext) IsOffset_sectionContext() {}

func NewOffset_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Offset_sectionContext {
	var p = new(Offset_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_offset_section

	return p
}

func (s *Offset_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Offset_sectionContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(SQLParserOFFSET, 0)
}

func (s *Offset_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Offset_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Offset_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterOffset_section(s)
	}
}

func (s *Offset_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitOffset_section(s)
	}
}

func (p *SQLParser) Offset_section() (localctx IOffset_sectionContext) {
	this := p
	_ = this

	localctx = NewOffset_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SQLParserRULE_offset_section)

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
		p.SetState(273)
		p.Match(SQLParserOFFSET)
	}

	return localctx
}

// IInsertStmtContext is an interface to support dynamic dispatch.
type IInsertStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INSERT() antlr.TerminalNode
	INTO() antlr.TerminalNode
	Collection_section() ICollection_sectionContext
	Insert_values_section() IInsert_values_sectionContext
	Sync_operator() ISync_operatorContext
	Insert_columns_section() IInsert_columns_sectionContext

	// IsInsertStmtContext differentiates from other interfaces.
	IsInsertStmtContext()
}

type InsertStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsertStmtContext() *InsertStmtContext {
	var p = new(InsertStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_insertStmt
	return p
}

func (*InsertStmtContext) IsInsertStmtContext() {}

func NewInsertStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertStmtContext {
	var p = new(InsertStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_insertStmt

	return p
}

func (s *InsertStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *InsertStmtContext) INSERT() antlr.TerminalNode {
	return s.GetToken(SQLParserINSERT, 0)
}

func (s *InsertStmtContext) INTO() antlr.TerminalNode {
	return s.GetToken(SQLParserINTO, 0)
}

func (s *InsertStmtContext) Collection_section() ICollection_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollection_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollection_sectionContext)
}

func (s *InsertStmtContext) Insert_values_section() IInsert_values_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInsert_values_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInsert_values_sectionContext)
}

func (s *InsertStmtContext) Sync_operator() ISync_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISync_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISync_operatorContext)
}

func (s *InsertStmtContext) Insert_columns_section() IInsert_columns_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInsert_columns_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInsert_columns_sectionContext)
}

func (s *InsertStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InsertStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterInsertStmt(s)
	}
}

func (s *InsertStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitInsertStmt(s)
	}
}

func (p *SQLParser) InsertStmt() (localctx IInsertStmtContext) {
	this := p
	_ = this

	localctx = NewInsertStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SQLParserRULE_insertStmt)
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
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserASYNC || _la == SQLParserSYNC {
		{
			p.SetState(275)
			p.Sync_operator()
		}

	}
	{
		p.SetState(278)
		p.Match(SQLParserINSERT)
	}
	{
		p.SetState(279)
		p.Match(SQLParserINTO)
	}
	{
		p.SetState(280)
		p.Collection_section()
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserT__0 {
		{
			p.SetState(281)
			p.Insert_columns_section()
		}

	}
	{
		p.SetState(284)
		p.Insert_values_section()
	}

	return localctx
}

// IInsert_columns_sectionContext is an interface to support dynamic dispatch.
type IInsert_columns_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllColumn() []IColumnContext
	Column(i int) IColumnContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsInsert_columns_sectionContext differentiates from other interfaces.
	IsInsert_columns_sectionContext()
}

type Insert_columns_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsert_columns_sectionContext() *Insert_columns_sectionContext {
	var p = new(Insert_columns_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_insert_columns_section
	return p
}

func (*Insert_columns_sectionContext) IsInsert_columns_sectionContext() {}

func NewInsert_columns_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Insert_columns_sectionContext {
	var p = new(Insert_columns_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_insert_columns_section

	return p
}

func (s *Insert_columns_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Insert_columns_sectionContext) AllColumn() []IColumnContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumnContext); ok {
			len++
		}
	}

	tst := make([]IColumnContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumnContext); ok {
			tst[i] = t.(IColumnContext)
			i++
		}
	}

	return tst
}

func (s *Insert_columns_sectionContext) Column(i int) IColumnContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnContext); ok {
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

	return t.(IColumnContext)
}

func (s *Insert_columns_sectionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLParserCOMMA)
}

func (s *Insert_columns_sectionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *Insert_columns_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Insert_columns_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Insert_columns_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterInsert_columns_section(s)
	}
}

func (s *Insert_columns_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitInsert_columns_section(s)
	}
}

func (p *SQLParser) Insert_columns_section() (localctx IInsert_columns_sectionContext) {
	this := p
	_ = this

	localctx = NewInsert_columns_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SQLParserRULE_insert_columns_section)
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
		p.SetState(286)
		p.Match(SQLParserT__0)
	}
	{
		p.SetState(287)
		p.Column()
	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserCOMMA {
		{
			p.SetState(288)
			p.Match(SQLParserCOMMA)
		}
		{
			p.SetState(289)
			p.Column()
		}

		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(295)
		p.Match(SQLParserT__1)
	}

	return localctx
}

// IInsert_values_sectionContext is an interface to support dynamic dispatch.
type IInsert_values_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VALUE() antlr.TerminalNode
	Expression() IExpressionContext
	VALUES() antlr.TerminalNode

	// IsInsert_values_sectionContext differentiates from other interfaces.
	IsInsert_values_sectionContext()
}

type Insert_values_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsert_values_sectionContext() *Insert_values_sectionContext {
	var p = new(Insert_values_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_insert_values_section
	return p
}

func (*Insert_values_sectionContext) IsInsert_values_sectionContext() {}

func NewInsert_values_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Insert_values_sectionContext {
	var p = new(Insert_values_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_insert_values_section

	return p
}

func (s *Insert_values_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Insert_values_sectionContext) VALUE() antlr.TerminalNode {
	return s.GetToken(SQLParserVALUE, 0)
}

func (s *Insert_values_sectionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Insert_values_sectionContext) VALUES() antlr.TerminalNode {
	return s.GetToken(SQLParserVALUES, 0)
}

func (s *Insert_values_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Insert_values_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Insert_values_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterInsert_values_section(s)
	}
}

func (s *Insert_values_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitInsert_values_section(s)
	}
}

func (p *SQLParser) Insert_values_section() (localctx IInsert_values_sectionContext) {
	this := p
	_ = this

	localctx = NewInsert_values_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SQLParserRULE_insert_values_section)

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

	p.SetState(301)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLParserVALUE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(297)
			p.Match(SQLParserVALUE)
		}
		{
			p.SetState(298)
			p.Expression()
		}

	case SQLParserVALUES:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(299)
			p.Match(SQLParserVALUES)
		}
		{
			p.SetState(300)
			p.Expression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUpdateStmtContext is an interface to support dynamic dispatch.
type IUpdateStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UPDATE() antlr.TerminalNode
	Collection_section() ICollection_sectionContext
	SET() antlr.TerminalNode
	AllProperty_section() []IProperty_sectionContext
	Property_section(i int) IProperty_sectionContext
	Sync_operator() ISync_operatorContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	Where_section() IWhere_sectionContext

	// IsUpdateStmtContext differentiates from other interfaces.
	IsUpdateStmtContext()
}

type UpdateStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateStmtContext() *UpdateStmtContext {
	var p = new(UpdateStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_updateStmt
	return p
}

func (*UpdateStmtContext) IsUpdateStmtContext() {}

func NewUpdateStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateStmtContext {
	var p = new(UpdateStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_updateStmt

	return p
}

func (s *UpdateStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateStmtContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(SQLParserUPDATE, 0)
}

func (s *UpdateStmtContext) Collection_section() ICollection_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollection_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollection_sectionContext)
}

func (s *UpdateStmtContext) SET() antlr.TerminalNode {
	return s.GetToken(SQLParserSET, 0)
}

func (s *UpdateStmtContext) AllProperty_section() []IProperty_sectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProperty_sectionContext); ok {
			len++
		}
	}

	tst := make([]IProperty_sectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProperty_sectionContext); ok {
			tst[i] = t.(IProperty_sectionContext)
			i++
		}
	}

	return tst
}

func (s *UpdateStmtContext) Property_section(i int) IProperty_sectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProperty_sectionContext); ok {
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

	return t.(IProperty_sectionContext)
}

func (s *UpdateStmtContext) Sync_operator() ISync_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISync_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISync_operatorContext)
}

func (s *UpdateStmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLParserCOMMA)
}

func (s *UpdateStmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *UpdateStmtContext) Where_section() IWhere_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhere_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhere_sectionContext)
}

func (s *UpdateStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpdateStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterUpdateStmt(s)
	}
}

func (s *UpdateStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitUpdateStmt(s)
	}
}

func (p *SQLParser) UpdateStmt() (localctx IUpdateStmtContext) {
	this := p
	_ = this

	localctx = NewUpdateStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SQLParserRULE_updateStmt)
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
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserASYNC || _la == SQLParserSYNC {
		{
			p.SetState(303)
			p.Sync_operator()
		}

	}
	{
		p.SetState(306)
		p.Match(SQLParserUPDATE)
	}
	{
		p.SetState(307)
		p.Collection_section()
	}
	{
		p.SetState(308)
		p.Match(SQLParserSET)
	}
	{
		p.SetState(309)
		p.Property_section()
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserCOMMA {
		{
			p.SetState(310)
			p.Match(SQLParserCOMMA)
		}
		{
			p.SetState(311)
			p.Property_section()
		}

		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserWHERE {
		{
			p.SetState(317)
			p.Where_section()
		}

	}

	return localctx
}

// IProperty_sectionContext is an interface to support dynamic dispatch.
type IProperty_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Property() IPropertyContext
	SINGLE_EQ() antlr.TerminalNode
	Expression_literal() IExpression_literalContext

	// IsProperty_sectionContext differentiates from other interfaces.
	IsProperty_sectionContext()
}

type Property_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProperty_sectionContext() *Property_sectionContext {
	var p = new(Property_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_property_section
	return p
}

func (*Property_sectionContext) IsProperty_sectionContext() {}

func NewProperty_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Property_sectionContext {
	var p = new(Property_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_property_section

	return p
}

func (s *Property_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Property_sectionContext) Property() IPropertyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *Property_sectionContext) SINGLE_EQ() antlr.TerminalNode {
	return s.GetToken(SQLParserSINGLE_EQ, 0)
}

func (s *Property_sectionContext) Expression_literal() IExpression_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_literalContext)
}

func (s *Property_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Property_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Property_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterProperty_section(s)
	}
}

func (s *Property_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitProperty_section(s)
	}
}

func (p *SQLParser) Property_section() (localctx IProperty_sectionContext) {
	this := p
	_ = this

	localctx = NewProperty_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SQLParserRULE_property_section)

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
		p.SetState(320)
		p.Property()
	}
	{
		p.SetState(321)
		p.Match(SQLParserSINGLE_EQ)
	}
	{
		p.SetState(322)
		p.Expression_literal()
	}

	return localctx
}

// IDeleteStmtContext is an interface to support dynamic dispatch.
type IDeleteStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DELETE() antlr.TerminalNode
	FROM() antlr.TerminalNode
	Collection_section() ICollection_sectionContext
	Sync_operator() ISync_operatorContext
	Where_section() IWhere_sectionContext

	// IsDeleteStmtContext differentiates from other interfaces.
	IsDeleteStmtContext()
}

type DeleteStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteStmtContext() *DeleteStmtContext {
	var p = new(DeleteStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_deleteStmt
	return p
}

func (*DeleteStmtContext) IsDeleteStmtContext() {}

func NewDeleteStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteStmtContext {
	var p = new(DeleteStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_deleteStmt

	return p
}

func (s *DeleteStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteStmtContext) DELETE() antlr.TerminalNode {
	return s.GetToken(SQLParserDELETE, 0)
}

func (s *DeleteStmtContext) FROM() antlr.TerminalNode {
	return s.GetToken(SQLParserFROM, 0)
}

func (s *DeleteStmtContext) Collection_section() ICollection_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollection_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollection_sectionContext)
}

func (s *DeleteStmtContext) Sync_operator() ISync_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISync_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISync_operatorContext)
}

func (s *DeleteStmtContext) Where_section() IWhere_sectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhere_sectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhere_sectionContext)
}

func (s *DeleteStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterDeleteStmt(s)
	}
}

func (s *DeleteStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitDeleteStmt(s)
	}
}

func (p *SQLParser) DeleteStmt() (localctx IDeleteStmtContext) {
	this := p
	_ = this

	localctx = NewDeleteStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SQLParserRULE_deleteStmt)
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
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserASYNC || _la == SQLParserSYNC {
		{
			p.SetState(324)
			p.Sync_operator()
		}

	}
	{
		p.SetState(327)
		p.Match(SQLParserDELETE)
	}
	{
		p.SetState(328)
		p.Match(SQLParserFROM)
	}
	{
		p.SetState(329)
		p.Collection_section()
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserWHERE {
		{
			p.SetState(330)
			p.Where_section()
		}

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression_list() IExpression_listContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Expression_list() IExpression_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_listContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SQLParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SQLParserRULE_expression)

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
		p.SetState(333)
		p.Expression_list()
	}

	return localctx
}

// IExpression_listContext is an interface to support dynamic dispatch.
type IExpression_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression_literal() IExpression_literalContext
	Expression_function() IExpression_functionContext
	AllExpression_binary_operator() []IExpression_binary_operatorContext
	Expression_binary_operator(i int) IExpression_binary_operatorContext
	AllExpression_logic_operator() []IExpression_logic_operatorContext
	Expression_logic_operator(i int) IExpression_logic_operatorContext
	AllExpression_array() []IExpression_arrayContext
	Expression_array(i int) IExpression_arrayContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllExpression_dictionary() []IExpression_dictionaryContext
	Expression_dictionary(i int) IExpression_dictionaryContext
	Unreserved_keyword() IUnreserved_keywordContext

	// IsExpression_listContext differentiates from other interfaces.
	IsExpression_listContext()
}

type Expression_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_listContext() *Expression_listContext {
	var p = new(Expression_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_expression_list
	return p
}

func (*Expression_listContext) IsExpression_listContext() {}

func NewExpression_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_listContext {
	var p = new(Expression_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_expression_list

	return p
}

func (s *Expression_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_listContext) Expression_literal() IExpression_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_literalContext)
}

func (s *Expression_listContext) Expression_function() IExpression_functionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_functionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_functionContext)
}

func (s *Expression_listContext) AllExpression_binary_operator() []IExpression_binary_operatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpression_binary_operatorContext); ok {
			len++
		}
	}

	tst := make([]IExpression_binary_operatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpression_binary_operatorContext); ok {
			tst[i] = t.(IExpression_binary_operatorContext)
			i++
		}
	}

	return tst
}

func (s *Expression_listContext) Expression_binary_operator(i int) IExpression_binary_operatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_binary_operatorContext); ok {
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

	return t.(IExpression_binary_operatorContext)
}

func (s *Expression_listContext) AllExpression_logic_operator() []IExpression_logic_operatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpression_logic_operatorContext); ok {
			len++
		}
	}

	tst := make([]IExpression_logic_operatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpression_logic_operatorContext); ok {
			tst[i] = t.(IExpression_logic_operatorContext)
			i++
		}
	}

	return tst
}

func (s *Expression_listContext) Expression_logic_operator(i int) IExpression_logic_operatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_logic_operatorContext); ok {
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

	return t.(IExpression_logic_operatorContext)
}

func (s *Expression_listContext) AllExpression_array() []IExpression_arrayContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpression_arrayContext); ok {
			len++
		}
	}

	tst := make([]IExpression_arrayContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpression_arrayContext); ok {
			tst[i] = t.(IExpression_arrayContext)
			i++
		}
	}

	return tst
}

func (s *Expression_listContext) Expression_array(i int) IExpression_arrayContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_arrayContext); ok {
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

	return t.(IExpression_arrayContext)
}

func (s *Expression_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLParserCOMMA)
}

func (s *Expression_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *Expression_listContext) AllExpression_dictionary() []IExpression_dictionaryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpression_dictionaryContext); ok {
			len++
		}
	}

	tst := make([]IExpression_dictionaryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpression_dictionaryContext); ok {
			tst[i] = t.(IExpression_dictionaryContext)
			i++
		}
	}

	return tst
}

func (s *Expression_listContext) Expression_dictionary(i int) IExpression_dictionaryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_dictionaryContext); ok {
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

	return t.(IExpression_dictionaryContext)
}

func (s *Expression_listContext) Unreserved_keyword() IUnreserved_keywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnreserved_keywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnreserved_keywordContext)
}

func (s *Expression_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterExpression_list(s)
	}
}

func (s *Expression_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitExpression_list(s)
	}
}

func (p *SQLParser) Expression_list() (localctx IExpression_listContext) {
	this := p
	_ = this

	localctx = NewExpression_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SQLParserRULE_expression_list)
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

	p.SetState(380)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(335)
			p.Expression_literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(336)
			p.Expression_function()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(337)
			p.Expression_binary_operator()
		}
		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLParserAND || _la == SQLParserOR {
			{
				p.SetState(338)
				p.Expression_logic_operator()
			}
			{
				p.SetState(339)
				p.Expression_binary_operator()
			}

			p.SetState(345)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(346)
			p.Match(SQLParserT__0)
		}
		{
			p.SetState(347)
			p.Expression_array()
		}
		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLParserCOMMA {
			{
				p.SetState(348)
				p.Match(SQLParserCOMMA)
			}
			{
				p.SetState(349)
				p.Expression_array()
			}

			p.SetState(354)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(355)
			p.Match(SQLParserT__1)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(357)
			p.Match(SQLParserT__2)
		}

		{
			p.SetState(358)
			p.Expression_dictionary()
		}

		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLParserCOMMA {
			{
				p.SetState(359)
				p.Match(SQLParserCOMMA)
			}
			{
				p.SetState(360)
				p.Expression_dictionary()
			}

			p.SetState(365)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(366)
			p.Match(SQLParserT__3)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(368)
			p.Match(SQLParserT__4)
		}
		{
			p.SetState(369)
			p.Expression_array()
		}
		p.SetState(374)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLParserCOMMA {
			{
				p.SetState(370)
				p.Match(SQLParserCOMMA)
			}
			{
				p.SetState(371)
				p.Expression_array()
			}

			p.SetState(376)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(377)
			p.Match(SQLParserT__5)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(379)
			p.Unreserved_keyword()
		}

	}

	return localctx
}

// IExpression_literalContext is an interface to support dynamic dispatch.
type IExpression_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression_literal_value() IExpression_literal_valueContext

	// IsExpression_literalContext differentiates from other interfaces.
	IsExpression_literalContext()
}

type Expression_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_literalContext() *Expression_literalContext {
	var p = new(Expression_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_expression_literal
	return p
}

func (*Expression_literalContext) IsExpression_literalContext() {}

func NewExpression_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_literalContext {
	var p = new(Expression_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_expression_literal

	return p
}

func (s *Expression_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_literalContext) Expression_literal_value() IExpression_literal_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_literal_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_literal_valueContext)
}

func (s *Expression_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterExpression_literal(s)
	}
}

func (s *Expression_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitExpression_literal(s)
	}
}

func (p *SQLParser) Expression_literal() (localctx IExpression_literalContext) {
	this := p
	_ = this

	localctx = NewExpression_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SQLParserRULE_expression_literal)

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
		p.SetState(382)
		p.Expression_literal_value()
	}

	return localctx
}

// IExpression_literal_valueContext is an interface to support dynamic dispatch.
type IExpression_literal_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Property_literal() IProperty_literalContext
	Integer_literal() IInteger_literalContext
	Real_literal() IReal_literalContext
	String_literal() IString_literalContext
	True_literal() ITrue_literalContext
	False_literal() IFalse_literalContext
	NIL() antlr.TerminalNode
	CURRENT_TIME() antlr.TerminalNode
	CURRENT_DATE() antlr.TerminalNode
	CURRENT_TIMESTAMP() antlr.TerminalNode

	// IsExpression_literal_valueContext differentiates from other interfaces.
	IsExpression_literal_valueContext()
}

type Expression_literal_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_literal_valueContext() *Expression_literal_valueContext {
	var p = new(Expression_literal_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_expression_literal_value
	return p
}

func (*Expression_literal_valueContext) IsExpression_literal_valueContext() {}

func NewExpression_literal_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_literal_valueContext {
	var p = new(Expression_literal_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_expression_literal_value

	return p
}

func (s *Expression_literal_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_literal_valueContext) Property_literal() IProperty_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProperty_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProperty_literalContext)
}

func (s *Expression_literal_valueContext) Integer_literal() IInteger_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInteger_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInteger_literalContext)
}

func (s *Expression_literal_valueContext) Real_literal() IReal_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReal_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReal_literalContext)
}

func (s *Expression_literal_valueContext) String_literal() IString_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *Expression_literal_valueContext) True_literal() ITrue_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITrue_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITrue_literalContext)
}

func (s *Expression_literal_valueContext) False_literal() IFalse_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFalse_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFalse_literalContext)
}

func (s *Expression_literal_valueContext) NIL() antlr.TerminalNode {
	return s.GetToken(SQLParserNIL, 0)
}

func (s *Expression_literal_valueContext) CURRENT_TIME() antlr.TerminalNode {
	return s.GetToken(SQLParserCURRENT_TIME, 0)
}

func (s *Expression_literal_valueContext) CURRENT_DATE() antlr.TerminalNode {
	return s.GetToken(SQLParserCURRENT_DATE, 0)
}

func (s *Expression_literal_valueContext) CURRENT_TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(SQLParserCURRENT_TIMESTAMP, 0)
}

func (s *Expression_literal_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_literal_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_literal_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterExpression_literal_value(s)
	}
}

func (s *Expression_literal_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitExpression_literal_value(s)
	}
}

func (p *SQLParser) Expression_literal_value() (localctx IExpression_literal_valueContext) {
	this := p
	_ = this

	localctx = NewExpression_literal_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SQLParserRULE_expression_literal_value)

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

	p.SetState(394)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(384)
			p.Property_literal()
		}

	case SQLParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(385)
			p.Integer_literal()
		}

	case SQLParserFLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(386)
			p.Real_literal()
		}

	case SQLParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(387)
			p.String_literal()
		}

	case SQLParserTRUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(388)
			p.True_literal()
		}

	case SQLParserFALSE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(389)
			p.False_literal()
		}

	case SQLParserNIL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(390)
			p.Match(SQLParserNIL)
		}

	case SQLParserCURRENT_TIME:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(391)
			p.Match(SQLParserCURRENT_TIME)
		}

	case SQLParserCURRENT_DATE:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(392)
			p.Match(SQLParserCURRENT_DATE)
		}

	case SQLParserCURRENT_TIMESTAMP:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(393)
			p.Match(SQLParserCURRENT_TIMESTAMP)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpression_dictionaryContext is an interface to support dynamic dispatch.
type IExpression_dictionaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Expression_literal() IExpression_literalContext

	// IsExpression_dictionaryContext differentiates from other interfaces.
	IsExpression_dictionaryContext()
}

type Expression_dictionaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_dictionaryContext() *Expression_dictionaryContext {
	var p = new(Expression_dictionaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_expression_dictionary
	return p
}

func (*Expression_dictionaryContext) IsExpression_dictionaryContext() {}

func NewExpression_dictionaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_dictionaryContext {
	var p = new(Expression_dictionaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_expression_dictionary

	return p
}

func (s *Expression_dictionaryContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_dictionaryContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Expression_dictionaryContext) Expression_literal() IExpression_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_literalContext)
}

func (s *Expression_dictionaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_dictionaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_dictionaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterExpression_dictionary(s)
	}
}

func (s *Expression_dictionaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitExpression_dictionary(s)
	}
}

func (p *SQLParser) Expression_dictionary() (localctx IExpression_dictionaryContext) {
	this := p
	_ = this

	localctx = NewExpression_dictionaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SQLParserRULE_expression_dictionary)

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
		p.SetState(396)
		p.Name()
	}
	{
		p.SetState(397)
		p.Match(SQLParserT__6)
	}
	{
		p.SetState(398)
		p.Expression_literal()
	}

	return localctx
}

// IDictionary_literalContext is an interface to support dynamic dispatch.
type IDictionary_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Expression_literal() IExpression_literalContext

	// IsDictionary_literalContext differentiates from other interfaces.
	IsDictionary_literalContext()
}

type Dictionary_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictionary_literalContext() *Dictionary_literalContext {
	var p = new(Dictionary_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_dictionary_literal
	return p
}

func (*Dictionary_literalContext) IsDictionary_literalContext() {}

func NewDictionary_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dictionary_literalContext {
	var p = new(Dictionary_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_dictionary_literal

	return p
}

func (s *Dictionary_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Dictionary_literalContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Dictionary_literalContext) Expression_literal() IExpression_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_literalContext)
}

func (s *Dictionary_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dictionary_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dictionary_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterDictionary_literal(s)
	}
}

func (s *Dictionary_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitDictionary_literal(s)
	}
}

func (p *SQLParser) Dictionary_literal() (localctx IDictionary_literalContext) {
	this := p
	_ = this

	localctx = NewDictionary_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SQLParserRULE_dictionary_literal)

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
		p.SetState(400)
		p.Name()
	}
	{
		p.SetState(401)
		p.Match(SQLParserT__6)
	}
	{
		p.SetState(402)
		p.Expression_literal()
	}

	return localctx
}

// IExpression_arrayContext is an interface to support dynamic dispatch.
type IExpression_arrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression_literal() IExpression_literalContext

	// IsExpression_arrayContext differentiates from other interfaces.
	IsExpression_arrayContext()
}

type Expression_arrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_arrayContext() *Expression_arrayContext {
	var p = new(Expression_arrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_expression_array
	return p
}

func (*Expression_arrayContext) IsExpression_arrayContext() {}

func NewExpression_arrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_arrayContext {
	var p = new(Expression_arrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_expression_array

	return p
}

func (s *Expression_arrayContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_arrayContext) Expression_literal() IExpression_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_literalContext)
}

func (s *Expression_arrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_arrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_arrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterExpression_array(s)
	}
}

func (s *Expression_arrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitExpression_array(s)
	}
}

func (p *SQLParser) Expression_array() (localctx IExpression_arrayContext) {
	this := p
	_ = this

	localctx = NewExpression_arrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SQLParserRULE_expression_array)

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
		p.SetState(404)
		p.Expression_literal()
	}

	return localctx
}

// IArray_literalContext is an interface to support dynamic dispatch.
type IArray_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression_literal() IExpression_literalContext

	// IsArray_literalContext differentiates from other interfaces.
	IsArray_literalContext()
}

type Array_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_literalContext() *Array_literalContext {
	var p = new(Array_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_array_literal
	return p
}

func (*Array_literalContext) IsArray_literalContext() {}

func NewArray_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_literalContext {
	var p = new(Array_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_array_literal

	return p
}

func (s *Array_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_literalContext) Expression_literal() IExpression_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_literalContext)
}

func (s *Array_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterArray_literal(s)
	}
}

func (s *Array_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitArray_literal(s)
	}
}

func (p *SQLParser) Array_literal() (localctx IArray_literalContext) {
	this := p
	_ = this

	localctx = NewArray_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SQLParserRULE_array_literal)

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
		p.SetState(406)
		p.Expression_literal()
	}

	return localctx
}

// IExpression_logic_operatorContext is an interface to support dynamic dispatch.
type IExpression_logic_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Logical_operator() ILogical_operatorContext

	// IsExpression_logic_operatorContext differentiates from other interfaces.
	IsExpression_logic_operatorContext()
}

type Expression_logic_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_logic_operatorContext() *Expression_logic_operatorContext {
	var p = new(Expression_logic_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_expression_logic_operator
	return p
}

func (*Expression_logic_operatorContext) IsExpression_logic_operatorContext() {}

func NewExpression_logic_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_logic_operatorContext {
	var p = new(Expression_logic_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_expression_logic_operator

	return p
}

func (s *Expression_logic_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_logic_operatorContext) Logical_operator() ILogical_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogical_operatorContext)
}

func (s *Expression_logic_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_logic_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_logic_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterExpression_logic_operator(s)
	}
}

func (s *Expression_logic_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitExpression_logic_operator(s)
	}
}

func (p *SQLParser) Expression_logic_operator() (localctx IExpression_logic_operatorContext) {
	this := p
	_ = this

	localctx = NewExpression_logic_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SQLParserRULE_expression_logic_operator)

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
		p.SetState(408)
		p.Logical_operator()
	}

	return localctx
}

// IExpression_binary_operatorContext is an interface to support dynamic dispatch.
type IExpression_binary_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression_operator() IExpression_operatorContext

	// IsExpression_binary_operatorContext differentiates from other interfaces.
	IsExpression_binary_operatorContext()
}

type Expression_binary_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_binary_operatorContext() *Expression_binary_operatorContext {
	var p = new(Expression_binary_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_expression_binary_operator
	return p
}

func (*Expression_binary_operatorContext) IsExpression_binary_operatorContext() {}

func NewExpression_binary_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_binary_operatorContext {
	var p = new(Expression_binary_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_expression_binary_operator

	return p
}

func (s *Expression_binary_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_binary_operatorContext) Expression_operator() IExpression_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_operatorContext)
}

func (s *Expression_binary_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_binary_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_binary_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterExpression_binary_operator(s)
	}
}

func (s *Expression_binary_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitExpression_binary_operator(s)
	}
}

func (p *SQLParser) Expression_binary_operator() (localctx IExpression_binary_operatorContext) {
	this := p
	_ = this

	localctx = NewExpression_binary_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SQLParserRULE_expression_binary_operator)

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
		p.SetState(410)
		p.Expression_operator()
	}

	return localctx
}

// IExpression_functionContext is an interface to support dynamic dispatch.
type IExpression_functionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Function_value() IFunction_valueContext

	// IsExpression_functionContext differentiates from other interfaces.
	IsExpression_functionContext()
}

type Expression_functionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_functionContext() *Expression_functionContext {
	var p = new(Expression_functionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_expression_function
	return p
}

func (*Expression_functionContext) IsExpression_functionContext() {}

func NewExpression_functionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_functionContext {
	var p = new(Expression_functionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_expression_function

	return p
}

func (s *Expression_functionContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_functionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLParserIDENTIFIER, 0)
}

func (s *Expression_functionContext) Function_value() IFunction_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_valueContext)
}

func (s *Expression_functionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_functionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_functionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterExpression_function(s)
	}
}

func (s *Expression_functionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitExpression_function(s)
	}
}

func (p *SQLParser) Expression_function() (localctx IExpression_functionContext) {
	this := p
	_ = this

	localctx = NewExpression_functionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SQLParserRULE_expression_function)
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
		p.SetState(412)
		p.Match(SQLParserIDENTIFIER)
	}
	{
		p.SetState(413)
		p.Match(SQLParserT__0)
	}
	p.SetState(415)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9221683185115463382) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&61) != 0) {
		{
			p.SetState(414)
			p.Function_value()
		}

	}
	{
		p.SetState(417)
		p.Match(SQLParserT__1)
	}

	return localctx
}

// IFunction_nameContext is an interface to support dynamic dispatch.
type IFunction_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsFunction_nameContext differentiates from other interfaces.
	IsFunction_nameContext()
}

type Function_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_nameContext() *Function_nameContext {
	var p = new(Function_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_function_name
	return p
}

func (*Function_nameContext) IsFunction_nameContext() {}

func NewFunction_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_nameContext {
	var p = new(Function_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_function_name

	return p
}

func (s *Function_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLParserIDENTIFIER, 0)
}

func (s *Function_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterFunction_name(s)
	}
}

func (s *Function_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitFunction_name(s)
	}
}

func (p *SQLParser) Function_name() (localctx IFunction_nameContext) {
	this := p
	_ = this

	localctx = NewFunction_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SQLParserRULE_function_name)

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
		p.SetState(419)
		p.Match(SQLParserIDENTIFIER)
	}

	return localctx
}

// IFunction_valueContext is an interface to support dynamic dispatch.
type IFunction_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	ASTERISK() antlr.TerminalNode

	// IsFunction_valueContext differentiates from other interfaces.
	IsFunction_valueContext()
}

type Function_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_valueContext() *Function_valueContext {
	var p = new(Function_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_function_value
	return p
}

func (*Function_valueContext) IsFunction_valueContext() {}

func NewFunction_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_valueContext {
	var p = new(Function_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_function_value

	return p
}

func (s *Function_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_valueContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Function_valueContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *Function_valueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLParserCOMMA)
}

func (s *Function_valueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *Function_valueContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(SQLParserASTERISK, 0)
}

func (s *Function_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterFunction_value(s)
	}
}

func (s *Function_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitFunction_value(s)
	}
}

func (p *SQLParser) Function_value() (localctx IFunction_valueContext) {
	this := p
	_ = this

	localctx = NewFunction_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SQLParserRULE_function_value)
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

	p.SetState(430)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLParserT__0, SQLParserT__2, SQLParserT__4, SQLParserCURRENT_DATE, SQLParserCURRENT_TIME, SQLParserCURRENT_TIMESTAMP, SQLParserNIL, SQLParserOFFSET, SQLParserTRUE, SQLParserFALSE, SQLParserIDENTIFIER, SQLParserNUMBER, SQLParserFLOAT, SQLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(421)
			p.Expression()
		}
		p.SetState(426)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLParserCOMMA {
			{
				p.SetState(422)
				p.Match(SQLParserCOMMA)
			}
			{
				p.SetState(423)
				p.Expression()
			}

			p.SetState(428)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case SQLParserASTERISK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(429)
			p.Match(SQLParserASTERISK)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpression_operatorContext is an interface to support dynamic dispatch.
type IExpression_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression_literal() []IExpression_literalContext
	Expression_literal(i int) IExpression_literalContext
	Binary_operator() IBinary_operatorContext

	// IsExpression_operatorContext differentiates from other interfaces.
	IsExpression_operatorContext()
}

type Expression_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_operatorContext() *Expression_operatorContext {
	var p = new(Expression_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_expression_operator
	return p
}

func (*Expression_operatorContext) IsExpression_operatorContext() {}

func NewExpression_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_operatorContext {
	var p = new(Expression_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_expression_operator

	return p
}

func (s *Expression_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_operatorContext) AllExpression_literal() []IExpression_literalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpression_literalContext); ok {
			len++
		}
	}

	tst := make([]IExpression_literalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpression_literalContext); ok {
			tst[i] = t.(IExpression_literalContext)
			i++
		}
	}

	return tst
}

func (s *Expression_operatorContext) Expression_literal(i int) IExpression_literalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_literalContext); ok {
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

	return t.(IExpression_literalContext)
}

func (s *Expression_operatorContext) Binary_operator() IBinary_operatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinary_operatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinary_operatorContext)
}

func (s *Expression_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterExpression_operator(s)
	}
}

func (s *Expression_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitExpression_operator(s)
	}
}

func (p *SQLParser) Expression_operator() (localctx IExpression_operatorContext) {
	this := p
	_ = this

	localctx = NewExpression_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, SQLParserRULE_expression_operator)

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
		p.SetState(432)
		p.Expression_literal()
	}
	{
		p.SetState(433)
		p.Binary_operator()
	}
	{
		p.SetState(434)
		p.Expression_literal()
	}

	return localctx
}

// IBinary_operatorContext is an interface to support dynamic dispatch.
type IBinary_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SINGLE_EQ() antlr.TerminalNode
	DOUBLE_EQ() antlr.TerminalNode
	OP_LT() antlr.TerminalNode
	LE() antlr.TerminalNode
	GT() antlr.TerminalNode
	GE() antlr.TerminalNode
	NOTEQ() antlr.TerminalNode

	// IsBinary_operatorContext differentiates from other interfaces.
	IsBinary_operatorContext()
}

type Binary_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinary_operatorContext() *Binary_operatorContext {
	var p = new(Binary_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_binary_operator
	return p
}

func (*Binary_operatorContext) IsBinary_operatorContext() {}

func NewBinary_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_operatorContext {
	var p = new(Binary_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_binary_operator

	return p
}

func (s *Binary_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Binary_operatorContext) SINGLE_EQ() antlr.TerminalNode {
	return s.GetToken(SQLParserSINGLE_EQ, 0)
}

func (s *Binary_operatorContext) DOUBLE_EQ() antlr.TerminalNode {
	return s.GetToken(SQLParserDOUBLE_EQ, 0)
}

func (s *Binary_operatorContext) OP_LT() antlr.TerminalNode {
	return s.GetToken(SQLParserOP_LT, 0)
}

func (s *Binary_operatorContext) LE() antlr.TerminalNode {
	return s.GetToken(SQLParserLE, 0)
}

func (s *Binary_operatorContext) GT() antlr.TerminalNode {
	return s.GetToken(SQLParserGT, 0)
}

func (s *Binary_operatorContext) GE() antlr.TerminalNode {
	return s.GetToken(SQLParserGE, 0)
}

func (s *Binary_operatorContext) NOTEQ() antlr.TerminalNode {
	return s.GetToken(SQLParserNOTEQ, 0)
}

func (s *Binary_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Binary_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterBinary_operator(s)
	}
}

func (s *Binary_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitBinary_operator(s)
	}
}

func (p *SQLParser) Binary_operator() (localctx IBinary_operatorContext) {
	this := p
	_ = this

	localctx = NewBinary_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, SQLParserRULE_binary_operator)
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
		p.SetState(436)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&65024) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILogical_operatorContext is an interface to support dynamic dispatch.
type ILogical_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsLogical_operatorContext differentiates from other interfaces.
	IsLogical_operatorContext()
}

type Logical_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogical_operatorContext() *Logical_operatorContext {
	var p = new(Logical_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_logical_operator
	return p
}

func (*Logical_operatorContext) IsLogical_operatorContext() {}

func NewLogical_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logical_operatorContext {
	var p = new(Logical_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_logical_operator

	return p
}

func (s *Logical_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Logical_operatorContext) AND() antlr.TerminalNode {
	return s.GetToken(SQLParserAND, 0)
}

func (s *Logical_operatorContext) OR() antlr.TerminalNode {
	return s.GetToken(SQLParserOR, 0)
}

func (s *Logical_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logical_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Logical_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterLogical_operator(s)
	}
}

func (s *Logical_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitLogical_operator(s)
	}
}

func (p *SQLParser) Logical_operator() (localctx ILogical_operatorContext) {
	this := p
	_ = this

	localctx = NewLogical_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, SQLParserRULE_logical_operator)
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
		p.SetState(438)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SQLParserAND || _la == SQLParserOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IProperty_literalContext is an interface to support dynamic dispatch.
type IProperty_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsProperty_literalContext differentiates from other interfaces.
	IsProperty_literalContext()
}

type Property_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProperty_literalContext() *Property_literalContext {
	var p = new(Property_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_property_literal
	return p
}

func (*Property_literalContext) IsProperty_literalContext() {}

func NewProperty_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Property_literalContext {
	var p = new(Property_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_property_literal

	return p
}

func (s *Property_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Property_literalContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLParserIDENTIFIER, 0)
}

func (s *Property_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Property_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Property_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterProperty_literal(s)
	}
}

func (s *Property_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitProperty_literal(s)
	}
}

func (p *SQLParser) Property_literal() (localctx IProperty_literalContext) {
	this := p
	_ = this

	localctx = NewProperty_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, SQLParserRULE_property_literal)

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
		p.SetState(440)
		p.Match(SQLParserIDENTIFIER)
	}

	return localctx
}

// IInteger_literalContext is an interface to support dynamic dispatch.
type IInteger_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsInteger_literalContext differentiates from other interfaces.
	IsInteger_literalContext()
}

type Integer_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInteger_literalContext() *Integer_literalContext {
	var p = new(Integer_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_integer_literal
	return p
}

func (*Integer_literalContext) IsInteger_literalContext() {}

func NewInteger_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integer_literalContext {
	var p = new(Integer_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_integer_literal

	return p
}

func (s *Integer_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Integer_literalContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SQLParserNUMBER, 0)
}

func (s *Integer_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integer_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterInteger_literal(s)
	}
}

func (s *Integer_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitInteger_literal(s)
	}
}

func (p *SQLParser) Integer_literal() (localctx IInteger_literalContext) {
	this := p
	_ = this

	localctx = NewInteger_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, SQLParserRULE_integer_literal)

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
		p.SetState(442)
		p.Match(SQLParserNUMBER)
	}

	return localctx
}

// IReal_literalContext is an interface to support dynamic dispatch.
type IReal_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOAT() antlr.TerminalNode

	// IsReal_literalContext differentiates from other interfaces.
	IsReal_literalContext()
}

type Real_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReal_literalContext() *Real_literalContext {
	var p = new(Real_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_real_literal
	return p
}

func (*Real_literalContext) IsReal_literalContext() {}

func NewReal_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Real_literalContext {
	var p = new(Real_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_real_literal

	return p
}

func (s *Real_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Real_literalContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SQLParserFLOAT, 0)
}

func (s *Real_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Real_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Real_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterReal_literal(s)
	}
}

func (s *Real_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitReal_literal(s)
	}
}

func (p *SQLParser) Real_literal() (localctx IReal_literalContext) {
	this := p
	_ = this

	localctx = NewReal_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, SQLParserRULE_real_literal)

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
		p.SetState(444)
		p.Match(SQLParserFLOAT)
	}

	return localctx
}

// IString_literalContext is an interface to support dynamic dispatch.
type IString_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsString_literalContext differentiates from other interfaces.
	IsString_literalContext()
}

type String_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_literalContext() *String_literalContext {
	var p = new(String_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_string_literal
	return p
}

func (*String_literalContext) IsString_literalContext() {}

func NewString_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_literalContext {
	var p = new(String_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_string_literal

	return p
}

func (s *String_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *String_literalContext) STRING() antlr.TerminalNode {
	return s.GetToken(SQLParserSTRING, 0)
}

func (s *String_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterString_literal(s)
	}
}

func (s *String_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitString_literal(s)
	}
}

func (p *SQLParser) String_literal() (localctx IString_literalContext) {
	this := p
	_ = this

	localctx = NewString_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, SQLParserRULE_string_literal)

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
		p.SetState(446)
		p.Match(SQLParserSTRING)
	}

	return localctx
}

// ITrue_literalContext is an interface to support dynamic dispatch.
type ITrue_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRUE() antlr.TerminalNode

	// IsTrue_literalContext differentiates from other interfaces.
	IsTrue_literalContext()
}

type True_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTrue_literalContext() *True_literalContext {
	var p = new(True_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_true_literal
	return p
}

func (*True_literalContext) IsTrue_literalContext() {}

func NewTrue_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *True_literalContext {
	var p = new(True_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_true_literal

	return p
}

func (s *True_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *True_literalContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SQLParserTRUE, 0)
}

func (s *True_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *True_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *True_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterTrue_literal(s)
	}
}

func (s *True_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitTrue_literal(s)
	}
}

func (p *SQLParser) True_literal() (localctx ITrue_literalContext) {
	this := p
	_ = this

	localctx = NewTrue_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, SQLParserRULE_true_literal)

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
		p.SetState(448)
		p.Match(SQLParserTRUE)
	}

	return localctx
}

// IFalse_literalContext is an interface to support dynamic dispatch.
type IFalse_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FALSE() antlr.TerminalNode

	// IsFalse_literalContext differentiates from other interfaces.
	IsFalse_literalContext()
}

type False_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFalse_literalContext() *False_literalContext {
	var p = new(False_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_false_literal
	return p
}

func (*False_literalContext) IsFalse_literalContext() {}

func NewFalse_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *False_literalContext {
	var p = new(False_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_false_literal

	return p
}

func (s *False_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *False_literalContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SQLParserFALSE, 0)
}

func (s *False_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *False_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *False_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterFalse_literal(s)
	}
}

func (s *False_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitFalse_literal(s)
	}
}

func (p *SQLParser) False_literal() (localctx IFalse_literalContext) {
	this := p
	_ = this

	localctx = NewFalse_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, SQLParserRULE_false_literal)

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
		p.SetState(450)
		p.Match(SQLParserFALSE)
	}

	return localctx
}

// ISync_operatorContext is an interface to support dynamic dispatch.
type ISync_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYNC() antlr.TerminalNode
	ASYNC() antlr.TerminalNode

	// IsSync_operatorContext differentiates from other interfaces.
	IsSync_operatorContext()
}

type Sync_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySync_operatorContext() *Sync_operatorContext {
	var p = new(Sync_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_sync_operator
	return p
}

func (*Sync_operatorContext) IsSync_operatorContext() {}

func NewSync_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sync_operatorContext {
	var p = new(Sync_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_sync_operator

	return p
}

func (s *Sync_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Sync_operatorContext) SYNC() antlr.TerminalNode {
	return s.GetToken(SQLParserSYNC, 0)
}

func (s *Sync_operatorContext) ASYNC() antlr.TerminalNode {
	return s.GetToken(SQLParserASYNC, 0)
}

func (s *Sync_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sync_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sync_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterSync_operator(s)
	}
}

func (s *Sync_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitSync_operator(s)
	}
}

func (p *SQLParser) Sync_operator() (localctx ISync_operatorContext) {
	this := p
	_ = this

	localctx = NewSync_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, SQLParserRULE_sync_operator)
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
		p.SetState(452)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SQLParserASYNC || _la == SQLParserSYNC) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICompound_operatorContext is an interface to support dynamic dispatch.
type ICompound_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UNION() antlr.TerminalNode
	ALL() antlr.TerminalNode
	INTERSECT() antlr.TerminalNode
	EXCEPT() antlr.TerminalNode

	// IsCompound_operatorContext differentiates from other interfaces.
	IsCompound_operatorContext()
}

type Compound_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompound_operatorContext() *Compound_operatorContext {
	var p = new(Compound_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_compound_operator
	return p
}

func (*Compound_operatorContext) IsCompound_operatorContext() {}

func NewCompound_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Compound_operatorContext {
	var p = new(Compound_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_compound_operator

	return p
}

func (s *Compound_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Compound_operatorContext) UNION() antlr.TerminalNode {
	return s.GetToken(SQLParserUNION, 0)
}

func (s *Compound_operatorContext) ALL() antlr.TerminalNode {
	return s.GetToken(SQLParserALL, 0)
}

func (s *Compound_operatorContext) INTERSECT() antlr.TerminalNode {
	return s.GetToken(SQLParserINTERSECT, 0)
}

func (s *Compound_operatorContext) EXCEPT() antlr.TerminalNode {
	return s.GetToken(SQLParserEXCEPT, 0)
}

func (s *Compound_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Compound_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Compound_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterCompound_operator(s)
	}
}

func (s *Compound_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitCompound_operator(s)
	}
}

func (p *SQLParser) Compound_operator() (localctx ICompound_operatorContext) {
	this := p
	_ = this

	localctx = NewCompound_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, SQLParserRULE_compound_operator)
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

	p.SetState(460)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLParserUNION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(454)
			p.Match(SQLParserUNION)
		}
		p.SetState(456)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLParserALL {
			{
				p.SetState(455)
				p.Match(SQLParserALL)
			}

		}

	case SQLParserINTERSECT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(458)
			p.Match(SQLParserINTERSECT)
		}

	case SQLParserEXCEPT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(459)
			p.Match(SQLParserEXCEPT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICondition_operatorContext is an interface to support dynamic dispatch.
type ICondition_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SINGLE_EQ() antlr.TerminalNode
	DOUBLE_EQ() antlr.TerminalNode
	OP_LT() antlr.TerminalNode
	LE() antlr.TerminalNode
	GT() antlr.TerminalNode
	GE() antlr.TerminalNode
	NOTEQ() antlr.TerminalNode

	// IsCondition_operatorContext differentiates from other interfaces.
	IsCondition_operatorContext()
}

type Condition_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondition_operatorContext() *Condition_operatorContext {
	var p = new(Condition_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_condition_operator
	return p
}

func (*Condition_operatorContext) IsCondition_operatorContext() {}

func NewCondition_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Condition_operatorContext {
	var p = new(Condition_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_condition_operator

	return p
}

func (s *Condition_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Condition_operatorContext) SINGLE_EQ() antlr.TerminalNode {
	return s.GetToken(SQLParserSINGLE_EQ, 0)
}

func (s *Condition_operatorContext) DOUBLE_EQ() antlr.TerminalNode {
	return s.GetToken(SQLParserDOUBLE_EQ, 0)
}

func (s *Condition_operatorContext) OP_LT() antlr.TerminalNode {
	return s.GetToken(SQLParserOP_LT, 0)
}

func (s *Condition_operatorContext) LE() antlr.TerminalNode {
	return s.GetToken(SQLParserLE, 0)
}

func (s *Condition_operatorContext) GT() antlr.TerminalNode {
	return s.GetToken(SQLParserGT, 0)
}

func (s *Condition_operatorContext) GE() antlr.TerminalNode {
	return s.GetToken(SQLParserGE, 0)
}

func (s *Condition_operatorContext) NOTEQ() antlr.TerminalNode {
	return s.GetToken(SQLParserNOTEQ, 0)
}

func (s *Condition_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Condition_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Condition_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterCondition_operator(s)
	}
}

func (s *Condition_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitCondition_operator(s)
	}
}

func (p *SQLParser) Condition_operator() (localctx ICondition_operatorContext) {
	this := p
	_ = this

	localctx = NewCondition_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, SQLParserRULE_condition_operator)
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
		p.SetState(462)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&65024) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPropertyContext is an interface to support dynamic dispatch.
type IPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsPropertyContext differentiates from other interfaces.
	IsPropertyContext()
}

type PropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyContext() *PropertyContext {
	var p = new(PropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_property
	return p
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLParserIDENTIFIER, 0)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterProperty(s)
	}
}

func (s *PropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitProperty(s)
	}
}

func (p *SQLParser) Property() (localctx IPropertyContext) {
	this := p
	_ = this

	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, SQLParserRULE_property)

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
		p.SetState(464)
		p.Match(SQLParserIDENTIFIER)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLParserIDENTIFIER, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *SQLParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, SQLParserRULE_value)

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
		p.SetState(466)
		p.Match(SQLParserIDENTIFIER)
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Unreserved_keyword() IUnreserved_keywordContext

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLParserIDENTIFIER, 0)
}

func (s *NameContext) Unreserved_keyword() IUnreserved_keywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnreserved_keywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnreserved_keywordContext)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *SQLParser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, SQLParserRULE_name)

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

	p.SetState(471)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(468)
			p.Match(SQLParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(469)
			p.Unreserved_keyword()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)

	}

	return localctx
}

// ICollection_sectionContext is an interface to support dynamic dispatch.
type ICollection_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CollectionName() ICollectionNameContext

	// IsCollection_sectionContext differentiates from other interfaces.
	IsCollection_sectionContext()
}

type Collection_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollection_sectionContext() *Collection_sectionContext {
	var p = new(Collection_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_collection_section
	return p
}

func (*Collection_sectionContext) IsCollection_sectionContext() {}

func NewCollection_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Collection_sectionContext {
	var p = new(Collection_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_collection_section

	return p
}

func (s *Collection_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Collection_sectionContext) CollectionName() ICollectionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollectionNameContext)
}

func (s *Collection_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Collection_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Collection_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterCollection_section(s)
	}
}

func (s *Collection_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitCollection_section(s)
	}
}

func (p *SQLParser) Collection_section() (localctx ICollection_sectionContext) {
	this := p
	_ = this

	localctx = NewCollection_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, SQLParserRULE_collection_section)

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
		p.SetState(473)
		p.CollectionName()
	}

	return localctx
}

// ICollectionNameContext is an interface to support dynamic dispatch.
type ICollectionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsCollectionNameContext differentiates from other interfaces.
	IsCollectionNameContext()
}

type CollectionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollectionNameContext() *CollectionNameContext {
	var p = new(CollectionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_collectionName
	return p
}

func (*CollectionNameContext) IsCollectionNameContext() {}

func NewCollectionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CollectionNameContext {
	var p = new(CollectionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_collectionName

	return p
}

func (s *CollectionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *CollectionNameContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *CollectionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CollectionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterCollectionName(s)
	}
}

func (s *CollectionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitCollectionName(s)
	}
}

func (p *SQLParser) CollectionName() (localctx ICollectionNameContext) {
	this := p
	_ = this

	localctx = NewCollectionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, SQLParserRULE_collectionName)

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
		p.SetState(475)
		p.Name()
	}

	return localctx
}

// IDatabaseNameContext is an interface to support dynamic dispatch.
type IDatabaseNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsDatabaseNameContext differentiates from other interfaces.
	IsDatabaseNameContext()
}

type DatabaseNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatabaseNameContext() *DatabaseNameContext {
	var p = new(DatabaseNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_databaseName
	return p
}

func (*DatabaseNameContext) IsDatabaseNameContext() {}

func NewDatabaseNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatabaseNameContext {
	var p = new(DatabaseNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_databaseName

	return p
}

func (s *DatabaseNameContext) GetParser() antlr.Parser { return s.parser }

func (s *DatabaseNameContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DatabaseNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatabaseNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatabaseNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterDatabaseName(s)
	}
}

func (s *DatabaseNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitDatabaseName(s)
	}
}

func (p *SQLParser) DatabaseName() (localctx IDatabaseNameContext) {
	this := p
	_ = this

	localctx = NewDatabaseNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, SQLParserRULE_databaseName)

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
		p.SetState(477)
		p.Name()
	}

	return localctx
}

// IColumnContext is an interface to support dynamic dispatch.
type IColumnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	AS() antlr.TerminalNode
	Name() INameContext

	// IsColumnContext differentiates from other interfaces.
	IsColumnContext()
}

type ColumnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnContext() *ColumnContext {
	var p = new(ColumnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_column
	return p
}

func (*ColumnContext) IsColumnContext() {}

func NewColumnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnContext {
	var p = new(ColumnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_column

	return p
}

func (s *ColumnContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ColumnContext) AS() antlr.TerminalNode {
	return s.GetToken(SQLParserAS, 0)
}

func (s *ColumnContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ColumnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterColumn(s)
	}
}

func (s *ColumnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitColumn(s)
	}
}

func (p *SQLParser) Column() (localctx IColumnContext) {
	this := p
	_ = this

	localctx = NewColumnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, SQLParserRULE_column)
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
		p.SetState(479)
		p.Expression()
	}

	p.SetState(482)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserAS {
		{
			p.SetState(480)
			p.Match(SQLParserAS)
		}
		{
			p.SetState(481)
			p.Name()
		}

	}

	return localctx
}

// IIndex_sectionContext is an interface to support dynamic dispatch.
type IIndex_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Index_name() IIndex_nameContext

	// IsIndex_sectionContext differentiates from other interfaces.
	IsIndex_sectionContext()
}

type Index_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndex_sectionContext() *Index_sectionContext {
	var p = new(Index_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_index_section
	return p
}

func (*Index_sectionContext) IsIndex_sectionContext() {}

func NewIndex_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_sectionContext {
	var p = new(Index_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_index_section

	return p
}

func (s *Index_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_sectionContext) Index_name() IIndex_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndex_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndex_nameContext)
}

func (s *Index_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Index_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterIndex_section(s)
	}
}

func (s *Index_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitIndex_section(s)
	}
}

func (p *SQLParser) Index_section() (localctx IIndex_sectionContext) {
	this := p
	_ = this

	localctx = NewIndex_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, SQLParserRULE_index_section)

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
		p.SetState(484)
		p.Index_name()
	}

	return localctx
}

// IIndex_nameContext is an interface to support dynamic dispatch.
type IIndex_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	String_literal() IString_literalContext

	// IsIndex_nameContext differentiates from other interfaces.
	IsIndex_nameContext()
}

type Index_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndex_nameContext() *Index_nameContext {
	var p = new(Index_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_index_name
	return p
}

func (*Index_nameContext) IsIndex_nameContext() {}

func NewIndex_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_nameContext {
	var p = new(Index_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_index_name

	return p
}

func (s *Index_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLParserIDENTIFIER, 0)
}

func (s *Index_nameContext) String_literal() IString_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *Index_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Index_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterIndex_name(s)
	}
}

func (s *Index_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitIndex_name(s)
	}
}

func (p *SQLParser) Index_name() (localctx IIndex_nameContext) {
	this := p
	_ = this

	localctx = NewIndex_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, SQLParserRULE_index_name)

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

	p.SetState(488)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(486)
			p.Match(SQLParserIDENTIFIER)
		}

	case SQLParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(487)
			p.String_literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IWhere_sectionContext is an interface to support dynamic dispatch.
type IWhere_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHERE() antlr.TerminalNode
	Expression() IExpressionContext

	// IsWhere_sectionContext differentiates from other interfaces.
	IsWhere_sectionContext()
}

type Where_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhere_sectionContext() *Where_sectionContext {
	var p = new(Where_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_where_section
	return p
}

func (*Where_sectionContext) IsWhere_sectionContext() {}

func NewWhere_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Where_sectionContext {
	var p = new(Where_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_where_section

	return p
}

func (s *Where_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Where_sectionContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SQLParserWHERE, 0)
}

func (s *Where_sectionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Where_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Where_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Where_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterWhere_section(s)
	}
}

func (s *Where_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitWhere_section(s)
	}
}

func (p *SQLParser) Where_section() (localctx IWhere_sectionContext) {
	this := p
	_ = this

	localctx = NewWhere_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, SQLParserRULE_where_section)

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
		p.SetState(490)
		p.Match(SQLParserWHERE)
	}
	{
		p.SetState(491)
		p.Expression()
	}

	return localctx
}

// IUnreserved_keywordContext is an interface to support dynamic dispatch.
type IUnreserved_keywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OFFSET() antlr.TerminalNode

	// IsUnreserved_keywordContext differentiates from other interfaces.
	IsUnreserved_keywordContext()
}

type Unreserved_keywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnreserved_keywordContext() *Unreserved_keywordContext {
	var p = new(Unreserved_keywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_unreserved_keyword
	return p
}

func (*Unreserved_keywordContext) IsUnreserved_keywordContext() {}

func NewUnreserved_keywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unreserved_keywordContext {
	var p = new(Unreserved_keywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_unreserved_keyword

	return p
}

func (s *Unreserved_keywordContext) GetParser() antlr.Parser { return s.parser }

func (s *Unreserved_keywordContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(SQLParserOFFSET, 0)
}

func (s *Unreserved_keywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unreserved_keywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unreserved_keywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterUnreserved_keyword(s)
	}
}

func (s *Unreserved_keywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitUnreserved_keyword(s)
	}
}

func (p *SQLParser) Unreserved_keyword() (localctx IUnreserved_keywordContext) {
	this := p
	_ = this

	localctx = NewUnreserved_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, SQLParserRULE_unreserved_keyword)

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
		p.SetState(493)
		p.Match(SQLParserOFFSET)
	}

	return localctx
}
