// Code generated from SQL.g4 by ANTLR 4.7.2. DO NOT EDIT.

package antlr // SQL
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 70, 470,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9,
	60, 4, 61, 9, 61, 3, 2, 3, 2, 3, 2, 7, 2, 126, 10, 2, 12, 2, 14, 2, 129,
	11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3,
	141, 10, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 5, 6, 154, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 5, 10, 170, 10, 10, 3, 10, 3, 10,
	3, 10, 5, 10, 175, 10, 10, 3, 10, 5, 10, 178, 10, 10, 3, 10, 5, 10, 181,
	10, 10, 3, 10, 5, 10, 184, 10, 10, 3, 10, 5, 10, 187, 10, 10, 3, 10, 5,
	10, 190, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 196, 10, 11, 12, 11,
	14, 11, 199, 11, 11, 5, 11, 201, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 7, 12, 208, 10, 12, 12, 12, 14, 12, 211, 11, 12, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 222, 10, 15, 12, 15, 14,
	15, 225, 11, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	7, 17, 235, 10, 17, 12, 17, 14, 17, 238, 11, 17, 3, 18, 3, 18, 5, 18, 242,
	10, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 250, 10, 20, 3,
	20, 3, 20, 3, 21, 3, 21, 3, 22, 5, 22, 257, 10, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 5, 22, 263, 10, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 7,
	23, 271, 10, 23, 12, 23, 14, 23, 274, 11, 23, 3, 23, 3, 23, 3, 24, 3, 24,
	3, 24, 3, 24, 5, 24, 282, 10, 24, 3, 25, 5, 25, 285, 10, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 7, 25, 293, 10, 25, 12, 25, 14, 25, 296,
	11, 25, 3, 25, 5, 25, 299, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 5,
	27, 306, 10, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 312, 10, 27, 3, 28,
	3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 7, 29, 322, 10, 29, 12,
	29, 14, 29, 325, 11, 29, 3, 29, 3, 29, 3, 29, 3, 29, 7, 29, 331, 10, 29,
	12, 29, 14, 29, 334, 11, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	7, 29, 342, 10, 29, 12, 29, 14, 29, 345, 11, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 7, 29, 353, 10, 29, 12, 29, 14, 29, 356, 11, 29, 3, 29,
	3, 29, 5, 29, 360, 10, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 374, 10, 31, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3,
	36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 5, 38, 395, 10, 38, 3, 38,
	3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 7, 40, 404, 10, 40, 12, 40, 14,
	40, 407, 11, 40, 3, 40, 5, 40, 410, 10, 40, 3, 41, 3, 41, 3, 41, 3, 41,
	3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3,
	47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51, 5, 51,
	436, 10, 51, 3, 51, 3, 51, 5, 51, 440, 10, 51, 3, 52, 3, 52, 3, 53, 3,
	53, 3, 54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 57, 3, 57, 5, 57, 454,
	10, 57, 3, 58, 3, 58, 3, 58, 5, 58, 459, 10, 58, 3, 59, 3, 59, 3, 60, 3,
	60, 5, 60, 465, 10, 60, 3, 61, 3, 61, 3, 61, 3, 61, 2, 2, 62, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
	80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112,
	114, 116, 118, 120, 2, 7, 4, 2, 22, 22, 35, 35, 4, 2, 25, 25, 33, 33, 3,
	2, 11, 17, 3, 2, 18, 19, 4, 2, 26, 26, 57, 57, 2, 470, 2, 122, 3, 2, 2,
	2, 4, 140, 3, 2, 2, 2, 6, 142, 3, 2, 2, 2, 8, 145, 3, 2, 2, 2, 10, 148,
	3, 2, 2, 2, 12, 155, 3, 2, 2, 2, 14, 159, 3, 2, 2, 2, 16, 163, 3, 2, 2,
	2, 18, 167, 3, 2, 2, 2, 20, 200, 3, 2, 2, 2, 22, 202, 3, 2, 2, 2, 24, 212,
	3, 2, 2, 2, 26, 214, 3, 2, 2, 2, 28, 216, 3, 2, 2, 2, 30, 226, 3, 2, 2,
	2, 32, 229, 3, 2, 2, 2, 34, 239, 3, 2, 2, 2, 36, 243, 3, 2, 2, 2, 38, 245,
	3, 2, 2, 2, 40, 253, 3, 2, 2, 2, 42, 256, 3, 2, 2, 2, 44, 266, 3, 2, 2,
	2, 46, 281, 3, 2, 2, 2, 48, 284, 3, 2, 2, 2, 50, 300, 3, 2, 2, 2, 52, 305,
	3, 2, 2, 2, 54, 313, 3, 2, 2, 2, 56, 359, 3, 2, 2, 2, 58, 361, 3, 2, 2,
	2, 60, 373, 3, 2, 2, 2, 62, 375, 3, 2, 2, 2, 64, 379, 3, 2, 2, 2, 66, 383,
	3, 2, 2, 2, 68, 385, 3, 2, 2, 2, 70, 387, 3, 2, 2, 2, 72, 389, 3, 2, 2,
	2, 74, 391, 3, 2, 2, 2, 76, 398, 3, 2, 2, 2, 78, 409, 3, 2, 2, 2, 80, 411,
	3, 2, 2, 2, 82, 415, 3, 2, 2, 2, 84, 417, 3, 2, 2, 2, 86, 419, 3, 2, 2,
	2, 88, 421, 3, 2, 2, 2, 90, 423, 3, 2, 2, 2, 92, 425, 3, 2, 2, 2, 94, 427,
	3, 2, 2, 2, 96, 429, 3, 2, 2, 2, 98, 431, 3, 2, 2, 2, 100, 439, 3, 2, 2,
	2, 102, 441, 3, 2, 2, 2, 104, 443, 3, 2, 2, 2, 106, 445, 3, 2, 2, 2, 108,
	447, 3, 2, 2, 2, 110, 449, 3, 2, 2, 2, 112, 453, 3, 2, 2, 2, 114, 455,
	3, 2, 2, 2, 116, 460, 3, 2, 2, 2, 118, 464, 3, 2, 2, 2, 120, 466, 3, 2,
	2, 2, 122, 127, 5, 4, 3, 2, 123, 124, 7, 21, 2, 2, 124, 126, 5, 4, 3, 2,
	125, 123, 3, 2, 2, 2, 126, 129, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 127,
	128, 3, 2, 2, 2, 128, 3, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 130, 141, 5,
	6, 4, 2, 131, 141, 5, 8, 5, 2, 132, 141, 5, 10, 6, 2, 133, 141, 5, 14,
	8, 2, 134, 141, 5, 12, 7, 2, 135, 141, 5, 16, 9, 2, 136, 141, 5, 18, 10,
	2, 137, 141, 5, 42, 22, 2, 138, 141, 5, 48, 25, 2, 139, 141, 5, 52, 27,
	2, 140, 130, 3, 2, 2, 2, 140, 131, 3, 2, 2, 2, 140, 132, 3, 2, 2, 2, 140,
	133, 3, 2, 2, 2, 140, 134, 3, 2, 2, 2, 140, 135, 3, 2, 2, 2, 140, 136,
	3, 2, 2, 2, 140, 137, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 140, 139, 3, 2,
	2, 2, 141, 5, 3, 2, 2, 2, 142, 143, 7, 56, 2, 2, 143, 144, 5, 110, 56,
	2, 144, 7, 3, 2, 2, 2, 145, 146, 7, 58, 2, 2, 146, 147, 5, 110, 56, 2,
	147, 9, 3, 2, 2, 2, 148, 149, 7, 28, 2, 2, 149, 150, 7, 29, 2, 2, 150,
	153, 5, 110, 56, 2, 151, 152, 7, 52, 2, 2, 152, 154, 5, 54, 28, 2, 153,
	151, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 11, 3, 2, 2, 2, 155, 156, 7,
	36, 2, 2, 156, 157, 7, 29, 2, 2, 157, 158, 5, 110, 56, 2, 158, 13, 3, 2,
	2, 2, 159, 160, 7, 28, 2, 2, 160, 161, 7, 44, 2, 2, 161, 162, 5, 116, 59,
	2, 162, 15, 3, 2, 2, 2, 163, 164, 7, 36, 2, 2, 164, 165, 7, 44, 2, 2, 165,
	166, 5, 116, 59, 2, 166, 17, 3, 2, 2, 2, 167, 169, 7, 54, 2, 2, 168, 170,
	9, 2, 2, 2, 169, 168, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 171, 3, 2,
	2, 2, 171, 172, 5, 20, 11, 2, 172, 174, 5, 22, 12, 2, 173, 175, 5, 120,
	61, 2, 174, 173, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 177, 3, 2, 2, 2,
	176, 178, 5, 28, 15, 2, 177, 176, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178,
	180, 3, 2, 2, 2, 179, 181, 5, 30, 16, 2, 180, 179, 3, 2, 2, 2, 180, 181,
	3, 2, 2, 2, 181, 183, 3, 2, 2, 2, 182, 184, 5, 32, 17, 2, 183, 182, 3,
	2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 186, 3, 2, 2, 2, 185, 187, 5, 38, 20,
	2, 186, 185, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 189, 3, 2, 2, 2, 188,
	190, 5, 40, 21, 2, 189, 188, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 19,
	3, 2, 2, 2, 191, 201, 7, 10, 2, 2, 192, 197, 5, 114, 58, 2, 193, 194, 7,
	20, 2, 2, 194, 196, 5, 114, 58, 2, 195, 193, 3, 2, 2, 2, 196, 199, 3, 2,
	2, 2, 197, 195, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 201, 3, 2, 2, 2,
	199, 197, 3, 2, 2, 2, 200, 191, 3, 2, 2, 2, 200, 192, 3, 2, 2, 2, 201,
	21, 3, 2, 2, 2, 202, 203, 7, 40, 2, 2, 203, 204, 5, 24, 13, 2, 204, 209,
	3, 2, 2, 2, 205, 206, 7, 20, 2, 2, 206, 208, 5, 24, 13, 2, 207, 205, 3,
	2, 2, 2, 208, 211, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2,
	2, 210, 23, 3, 2, 2, 2, 211, 209, 3, 2, 2, 2, 212, 213, 5, 26, 14, 2, 213,
	25, 3, 2, 2, 2, 214, 215, 5, 112, 57, 2, 215, 27, 3, 2, 2, 2, 216, 217,
	7, 41, 2, 2, 217, 218, 7, 27, 2, 2, 218, 223, 5, 54, 28, 2, 219, 220, 7,
	20, 2, 2, 220, 222, 5, 54, 28, 2, 221, 219, 3, 2, 2, 2, 222, 225, 3, 2,
	2, 2, 223, 221, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 29, 3, 2, 2, 2,
	225, 223, 3, 2, 2, 2, 226, 227, 7, 42, 2, 2, 227, 228, 5, 54, 28, 2, 228,
	31, 3, 2, 2, 2, 229, 230, 7, 53, 2, 2, 230, 231, 7, 27, 2, 2, 231, 236,
	5, 34, 18, 2, 232, 233, 7, 20, 2, 2, 233, 235, 5, 34, 18, 2, 234, 232,
	3, 2, 2, 2, 235, 238, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 236, 237, 3, 2,
	2, 2, 237, 33, 3, 2, 2, 2, 238, 236, 3, 2, 2, 2, 239, 241, 5, 104, 53,
	2, 240, 242, 5, 36, 19, 2, 241, 240, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2,
	242, 35, 3, 2, 2, 2, 243, 244, 9, 3, 2, 2, 244, 37, 3, 2, 2, 2, 245, 249,
	7, 49, 2, 2, 246, 247, 5, 58, 30, 2, 247, 248, 7, 20, 2, 2, 248, 250, 3,
	2, 2, 2, 249, 246, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250, 251, 3, 2, 2,
	2, 251, 252, 5, 58, 30, 2, 252, 39, 3, 2, 2, 2, 253, 254, 7, 51, 2, 2,
	254, 41, 3, 2, 2, 2, 255, 257, 5, 98, 50, 2, 256, 255, 3, 2, 2, 2, 256,
	257, 3, 2, 2, 2, 257, 258, 3, 2, 2, 2, 258, 259, 7, 45, 2, 2, 259, 260,
	7, 47, 2, 2, 260, 262, 5, 110, 56, 2, 261, 263, 5, 44, 23, 2, 262, 261,
	3, 2, 2, 2, 262, 263, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 265, 5, 46,
	24, 2, 265, 43, 3, 2, 2, 2, 266, 267, 7, 3, 2, 2, 267, 272, 5, 114, 58,
	2, 268, 269, 7, 20, 2, 2, 269, 271, 5, 114, 58, 2, 270, 268, 3, 2, 2, 2,
	271, 274, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273,
	275, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2, 275, 276, 7, 4, 2, 2, 276, 45, 3,
	2, 2, 2, 277, 278, 7, 62, 2, 2, 278, 282, 5, 54, 28, 2, 279, 280, 7, 63,
	2, 2, 280, 282, 5, 54, 28, 2, 281, 277, 3, 2, 2, 2, 281, 279, 3, 2, 2,
	2, 282, 47, 3, 2, 2, 2, 283, 285, 5, 98, 50, 2, 284, 283, 3, 2, 2, 2, 284,
	285, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 287, 7, 60, 2, 2, 287, 288,
	5, 110, 56, 2, 288, 289, 7, 55, 2, 2, 289, 294, 5, 50, 26, 2, 290, 291,
	7, 20, 2, 2, 291, 293, 5, 50, 26, 2, 292, 290, 3, 2, 2, 2, 293, 296, 3,
	2, 2, 2, 294, 292, 3, 2, 2, 2, 294, 295, 3, 2, 2, 2, 295, 298, 3, 2, 2,
	2, 296, 294, 3, 2, 2, 2, 297, 299, 5, 120, 61, 2, 298, 297, 3, 2, 2, 2,
	298, 299, 3, 2, 2, 2, 299, 49, 3, 2, 2, 2, 300, 301, 5, 104, 53, 2, 301,
	302, 7, 11, 2, 2, 302, 303, 5, 58, 30, 2, 303, 51, 3, 2, 2, 2, 304, 306,
	5, 98, 50, 2, 305, 304, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 307, 3,
	2, 2, 2, 307, 308, 7, 34, 2, 2, 308, 309, 7, 40, 2, 2, 309, 311, 5, 110,
	56, 2, 310, 312, 5, 120, 61, 2, 311, 310, 3, 2, 2, 2, 311, 312, 3, 2, 2,
	2, 312, 53, 3, 2, 2, 2, 313, 314, 5, 56, 29, 2, 314, 55, 3, 2, 2, 2, 315,
	360, 5, 58, 30, 2, 316, 360, 5, 74, 38, 2, 317, 323, 5, 72, 37, 2, 318,
	319, 5, 70, 36, 2, 319, 320, 5, 72, 37, 2, 320, 322, 3, 2, 2, 2, 321, 318,
	3, 2, 2, 2, 322, 325, 3, 2, 2, 2, 323, 321, 3, 2, 2, 2, 323, 324, 3, 2,
	2, 2, 324, 360, 3, 2, 2, 2, 325, 323, 3, 2, 2, 2, 326, 327, 7, 3, 2, 2,
	327, 332, 5, 66, 34, 2, 328, 329, 7, 20, 2, 2, 329, 331, 5, 66, 34, 2,
	330, 328, 3, 2, 2, 2, 331, 334, 3, 2, 2, 2, 332, 330, 3, 2, 2, 2, 332,
	333, 3, 2, 2, 2, 333, 335, 3, 2, 2, 2, 334, 332, 3, 2, 2, 2, 335, 336,
	7, 4, 2, 2, 336, 360, 3, 2, 2, 2, 337, 338, 7, 5, 2, 2, 338, 343, 5, 62,
	32, 2, 339, 340, 7, 20, 2, 2, 340, 342, 5, 62, 32, 2, 341, 339, 3, 2, 2,
	2, 342, 345, 3, 2, 2, 2, 343, 341, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344,
	346, 3, 2, 2, 2, 345, 343, 3, 2, 2, 2, 346, 347, 7, 6, 2, 2, 347, 360,
	3, 2, 2, 2, 348, 349, 7, 7, 2, 2, 349, 354, 5, 66, 34, 2, 350, 351, 7,
	20, 2, 2, 351, 353, 5, 66, 34, 2, 352, 350, 3, 2, 2, 2, 353, 356, 3, 2,
	2, 2, 354, 352, 3, 2, 2, 2, 354, 355, 3, 2, 2, 2, 355, 357, 3, 2, 2, 2,
	356, 354, 3, 2, 2, 2, 357, 358, 7, 8, 2, 2, 358, 360, 3, 2, 2, 2, 359,
	315, 3, 2, 2, 2, 359, 316, 3, 2, 2, 2, 359, 317, 3, 2, 2, 2, 359, 326,
	3, 2, 2, 2, 359, 337, 3, 2, 2, 2, 359, 348, 3, 2, 2, 2, 360, 57, 3, 2,
	2, 2, 361, 362, 5, 60, 31, 2, 362, 59, 3, 2, 2, 2, 363, 374, 5, 86, 44,
	2, 364, 374, 5, 88, 45, 2, 365, 374, 5, 90, 46, 2, 366, 374, 5, 92, 47,
	2, 367, 374, 5, 94, 48, 2, 368, 374, 5, 96, 49, 2, 369, 374, 7, 50, 2,
	2, 370, 374, 7, 31, 2, 2, 371, 374, 7, 30, 2, 2, 372, 374, 7, 32, 2, 2,
	373, 363, 3, 2, 2, 2, 373, 364, 3, 2, 2, 2, 373, 365, 3, 2, 2, 2, 373,
	366, 3, 2, 2, 2, 373, 367, 3, 2, 2, 2, 373, 368, 3, 2, 2, 2, 373, 369,
	3, 2, 2, 2, 373, 370, 3, 2, 2, 2, 373, 371, 3, 2, 2, 2, 373, 372, 3, 2,
	2, 2, 374, 61, 3, 2, 2, 2, 375, 376, 5, 108, 55, 2, 376, 377, 7, 9, 2,
	2, 377, 378, 5, 58, 30, 2, 378, 63, 3, 2, 2, 2, 379, 380, 5, 108, 55, 2,
	380, 381, 7, 9, 2, 2, 381, 382, 5, 58, 30, 2, 382, 65, 3, 2, 2, 2, 383,
	384, 5, 58, 30, 2, 384, 67, 3, 2, 2, 2, 385, 386, 5, 58, 30, 2, 386, 69,
	3, 2, 2, 2, 387, 388, 5, 84, 43, 2, 388, 71, 3, 2, 2, 2, 389, 390, 5, 80,
	41, 2, 390, 73, 3, 2, 2, 2, 391, 392, 7, 67, 2, 2, 392, 394, 7, 3, 2, 2,
	393, 395, 5, 78, 40, 2, 394, 393, 3, 2, 2, 2, 394, 395, 3, 2, 2, 2, 395,
	396, 3, 2, 2, 2, 396, 397, 7, 4, 2, 2, 397, 75, 3, 2, 2, 2, 398, 399, 7,
	67, 2, 2, 399, 77, 3, 2, 2, 2, 400, 405, 5, 54, 28, 2, 401, 402, 7, 20,
	2, 2, 402, 404, 5, 54, 28, 2, 403, 401, 3, 2, 2, 2, 404, 407, 3, 2, 2,
	2, 405, 403, 3, 2, 2, 2, 405, 406, 3, 2, 2, 2, 406, 410, 3, 2, 2, 2, 407,
	405, 3, 2, 2, 2, 408, 410, 7, 10, 2, 2, 409, 400, 3, 2, 2, 2, 409, 408,
	3, 2, 2, 2, 410, 79, 3, 2, 2, 2, 411, 412, 5, 58, 30, 2, 412, 413, 5, 82,
	42, 2, 413, 414, 5, 58, 30, 2, 414, 81, 3, 2, 2, 2, 415, 416, 9, 4, 2,
	2, 416, 83, 3, 2, 2, 2, 417, 418, 9, 5, 2, 2, 418, 85, 3, 2, 2, 2, 419,
	420, 7, 67, 2, 2, 420, 87, 3, 2, 2, 2, 421, 422, 7, 68, 2, 2, 422, 89,
	3, 2, 2, 2, 423, 424, 7, 69, 2, 2, 424, 91, 3, 2, 2, 2, 425, 426, 7, 70,
	2, 2, 426, 93, 3, 2, 2, 2, 427, 428, 7, 64, 2, 2, 428, 95, 3, 2, 2, 2,
	429, 430, 7, 65, 2, 2, 430, 97, 3, 2, 2, 2, 431, 432, 9, 6, 2, 2, 432,
	99, 3, 2, 2, 2, 433, 435, 7, 59, 2, 2, 434, 436, 7, 22, 2, 2, 435, 434,
	3, 2, 2, 2, 435, 436, 3, 2, 2, 2, 436, 440, 3, 2, 2, 2, 437, 440, 7, 46,
	2, 2, 438, 440, 7, 38, 2, 2, 439, 433, 3, 2, 2, 2, 439, 437, 3, 2, 2, 2,
	439, 438, 3, 2, 2, 2, 440, 101, 3, 2, 2, 2, 441, 442, 9, 4, 2, 2, 442,
	103, 3, 2, 2, 2, 443, 444, 7, 67, 2, 2, 444, 105, 3, 2, 2, 2, 445, 446,
	7, 67, 2, 2, 446, 107, 3, 2, 2, 2, 447, 448, 7, 67, 2, 2, 448, 109, 3,
	2, 2, 2, 449, 450, 5, 112, 57, 2, 450, 111, 3, 2, 2, 2, 451, 454, 7, 67,
	2, 2, 452, 454, 5, 92, 47, 2, 453, 451, 3, 2, 2, 2, 453, 452, 3, 2, 2,
	2, 454, 113, 3, 2, 2, 2, 455, 458, 5, 54, 28, 2, 456, 457, 7, 24, 2, 2,
	457, 459, 5, 108, 55, 2, 458, 456, 3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459,
	115, 3, 2, 2, 2, 460, 461, 5, 118, 60, 2, 461, 117, 3, 2, 2, 2, 462, 465,
	7, 67, 2, 2, 463, 465, 5, 92, 47, 2, 464, 462, 3, 2, 2, 2, 464, 463, 3,
	2, 2, 2, 465, 119, 3, 2, 2, 2, 466, 467, 7, 61, 2, 2, 467, 468, 5, 54,
	28, 2, 468, 121, 3, 2, 2, 2, 42, 127, 140, 153, 169, 174, 177, 180, 183,
	186, 189, 197, 200, 209, 223, 236, 241, 249, 256, 262, 272, 281, 284, 294,
	298, 305, 311, 323, 332, 343, 354, 359, 373, 394, 405, 409, 435, 439, 453,
	458, 464,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'{'", "'}'", "'['", "']'", "':'", "'*'", "'='", "'=='",
	"'<'", "'<='", "'>'", "'>='", "'!='", "", "", "','", "';'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "ASTERISK", "SINGLE_EQ", "DOUBLE_EQ", "OP_LT",
	"LE", "GT", "GE", "NOTEQ", "AND", "OR", "COMMA", "SEMICOLON", "ALL", "ANCESTOR",
	"AS", "ASC", "ASYNC", "BY", "CREATE", "COLLECTION", "CURRENT_DATE", "CURRENT_TIME",
	"CURRENT_TIMESTAMP", "DESC", "DELETE", "DISTINCT", "DROP", "EACH", "EXCEPT",
	"FLATTEN", "FROM", "GROUP", "HAVING", "IN", "COLLECTION_INDEX", "INSERT",
	"INTERSECT", "INTO", "IS", "LIMIT", "NIL", "OFFSET", "OPTIONS", "ORDER",
	"SELECT", "SET", "SHOW", "SYNC", "USE", "UNION", "UPDATE", "WHERE", "VALUE",
	"VALUES", "TRUE", "FALSE", "WS", "IDENTIFIER", "NUMBER", "FLOAT", "STRING",
}

var ruleNames = []string{
	"queries", "statement", "showStmt", "useStmt", "create_collectionStmt",
	"drop_collectionStmt", "create_indexStmt", "drop_indexStmt", "selectStmt",
	"columns", "from_section", "table_name", "data_source", "grouping_section",
	"having_section", "sorting_section", "sorting_item", "sorting_specification",
	"limit_section", "offset_section", "insertStmt", "insert_columns_section",
	"insert_values_section", "updateStmt", "property_section", "deleteStmt",
	"expression", "expression_list", "expression_literal", "expression_literal_value",
	"expression_dictionary", "dictionary_literal", "expression_array", "array_literal",
	"expression_logic_operator", "expression_binary_operator", "expression_function",
	"function_name", "function_value", "expression_operator", "binary_operator",
	"logical_operator", "property_literal", "integer_literal", "real_literal",
	"string_literal", "true_literal", "false_literal", "sync_operator", "compound_operator",
	"condition_operator", "property", "value", "name", "collection_section",
	"collection_name", "column", "index_section", "index_name", "where_section",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SQLParser struct {
	*antlr.BaseParser
}

func NewSQLParser(input antlr.TokenStream) *SQLParser {
	this := new(SQLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
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
	SQLParserDESC              = 31
	SQLParserDELETE            = 32
	SQLParserDISTINCT          = 33
	SQLParserDROP              = 34
	SQLParserEACH              = 35
	SQLParserEXCEPT            = 36
	SQLParserFLATTEN           = 37
	SQLParserFROM              = 38
	SQLParserGROUP             = 39
	SQLParserHAVING            = 40
	SQLParserIN                = 41
	SQLParserCOLLECTION_INDEX  = 42
	SQLParserINSERT            = 43
	SQLParserINTERSECT         = 44
	SQLParserINTO              = 45
	SQLParserIS                = 46
	SQLParserLIMIT             = 47
	SQLParserNIL               = 48
	SQLParserOFFSET            = 49
	SQLParserOPTIONS           = 50
	SQLParserORDER             = 51
	SQLParserSELECT            = 52
	SQLParserSET               = 53
	SQLParserSHOW              = 54
	SQLParserSYNC              = 55
	SQLParserUSE               = 56
	SQLParserUNION             = 57
	SQLParserUPDATE            = 58
	SQLParserWHERE             = 59
	SQLParserVALUE             = 60
	SQLParserVALUES            = 61
	SQLParserTRUE              = 62
	SQLParserFALSE             = 63
	SQLParserWS                = 64
	SQLParserIDENTIFIER        = 65
	SQLParserNUMBER            = 66
	SQLParserFLOAT             = 67
	SQLParserSTRING            = 68
)

// SQLParser rules.
const (
	SQLParserRULE_queries                    = 0
	SQLParserRULE_statement                  = 1
	SQLParserRULE_showStmt                   = 2
	SQLParserRULE_useStmt                    = 3
	SQLParserRULE_create_collectionStmt      = 4
	SQLParserRULE_drop_collectionStmt        = 5
	SQLParserRULE_create_indexStmt           = 6
	SQLParserRULE_drop_indexStmt             = 7
	SQLParserRULE_selectStmt                 = 8
	SQLParserRULE_columns                    = 9
	SQLParserRULE_from_section               = 10
	SQLParserRULE_table_name                 = 11
	SQLParserRULE_data_source                = 12
	SQLParserRULE_grouping_section           = 13
	SQLParserRULE_having_section             = 14
	SQLParserRULE_sorting_section            = 15
	SQLParserRULE_sorting_item               = 16
	SQLParserRULE_sorting_specification      = 17
	SQLParserRULE_limit_section              = 18
	SQLParserRULE_offset_section             = 19
	SQLParserRULE_insertStmt                 = 20
	SQLParserRULE_insert_columns_section     = 21
	SQLParserRULE_insert_values_section      = 22
	SQLParserRULE_updateStmt                 = 23
	SQLParserRULE_property_section           = 24
	SQLParserRULE_deleteStmt                 = 25
	SQLParserRULE_expression                 = 26
	SQLParserRULE_expression_list            = 27
	SQLParserRULE_expression_literal         = 28
	SQLParserRULE_expression_literal_value   = 29
	SQLParserRULE_expression_dictionary      = 30
	SQLParserRULE_dictionary_literal         = 31
	SQLParserRULE_expression_array           = 32
	SQLParserRULE_array_literal              = 33
	SQLParserRULE_expression_logic_operator  = 34
	SQLParserRULE_expression_binary_operator = 35
	SQLParserRULE_expression_function        = 36
	SQLParserRULE_function_name              = 37
	SQLParserRULE_function_value             = 38
	SQLParserRULE_expression_operator        = 39
	SQLParserRULE_binary_operator            = 40
	SQLParserRULE_logical_operator           = 41
	SQLParserRULE_property_literal           = 42
	SQLParserRULE_integer_literal            = 43
	SQLParserRULE_real_literal               = 44
	SQLParserRULE_string_literal             = 45
	SQLParserRULE_true_literal               = 46
	SQLParserRULE_false_literal              = 47
	SQLParserRULE_sync_operator              = 48
	SQLParserRULE_compound_operator          = 49
	SQLParserRULE_condition_operator         = 50
	SQLParserRULE_property                   = 51
	SQLParserRULE_value                      = 52
	SQLParserRULE_name                       = 53
	SQLParserRULE_collection_section         = 54
	SQLParserRULE_collection_name            = 55
	SQLParserRULE_column                     = 56
	SQLParserRULE_index_section              = 57
	SQLParserRULE_index_name                 = 58
	SQLParserRULE_where_section              = 59
)

// IQueriesContext is an interface to support dynamic dispatch.
type IQueriesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *QueriesContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Statement()
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserSEMICOLON {
		{
			p.SetState(121)
			p.Match(SQLParserSEMICOLON)
		}
		{
			p.SetState(122)
			p.Statement()
		}

		p.SetState(127)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShowStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShowStmtContext)
}

func (s *StatementContext) UseStmt() IUseStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUseStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUseStmtContext)
}

func (s *StatementContext) Create_collectionStmt() ICreate_collectionStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreate_collectionStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreate_collectionStmtContext)
}

func (s *StatementContext) Create_indexStmt() ICreate_indexStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreate_indexStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreate_indexStmtContext)
}

func (s *StatementContext) Drop_collectionStmt() IDrop_collectionStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDrop_collectionStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDrop_collectionStmtContext)
}

func (s *StatementContext) Drop_indexStmt() IDrop_indexStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDrop_indexStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDrop_indexStmtContext)
}

func (s *StatementContext) SelectStmt() ISelectStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStmtContext)
}

func (s *StatementContext) InsertStmt() IInsertStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInsertStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInsertStmtContext)
}

func (s *StatementContext) UpdateStmt() IUpdateStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdateStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdateStmtContext)
}

func (s *StatementContext) DeleteStmt() IDeleteStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeleteStmtContext)(nil)).Elem(), 0)

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

	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.ShowStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.UseStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.Create_collectionStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(131)
			p.Create_indexStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(132)
			p.Drop_collectionStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(133)
			p.Drop_indexStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(134)
			p.SelectStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(135)
			p.InsertStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(136)
			p.UpdateStmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(137)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICollection_sectionContext)(nil)).Elem(), 0)

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
		p.SetState(140)
		p.Match(SQLParserSHOW)
	}
	{
		p.SetState(141)
		p.Collection_section()
	}

	return localctx
}

// IUseStmtContext is an interface to support dynamic dispatch.
type IUseStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICollection_sectionContext)(nil)).Elem(), 0)

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
		p.SetState(143)
		p.Match(SQLParserUSE)
	}
	{
		p.SetState(144)
		p.Collection_section()
	}

	return localctx
}

// ICreate_collectionStmtContext is an interface to support dynamic dispatch.
type ICreate_collectionStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreate_collectionStmtContext differentiates from other interfaces.
	IsCreate_collectionStmtContext()
}

type Create_collectionStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreate_collectionStmtContext() *Create_collectionStmtContext {
	var p = new(Create_collectionStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_create_collectionStmt
	return p
}

func (*Create_collectionStmtContext) IsCreate_collectionStmtContext() {}

func NewCreate_collectionStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Create_collectionStmtContext {
	var p = new(Create_collectionStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_create_collectionStmt

	return p
}

func (s *Create_collectionStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Create_collectionStmtContext) CREATE() antlr.TerminalNode {
	return s.GetToken(SQLParserCREATE, 0)
}

func (s *Create_collectionStmtContext) COLLECTION() antlr.TerminalNode {
	return s.GetToken(SQLParserCOLLECTION, 0)
}

func (s *Create_collectionStmtContext) Collection_section() ICollection_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICollection_sectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICollection_sectionContext)
}

func (s *Create_collectionStmtContext) OPTIONS() antlr.TerminalNode {
	return s.GetToken(SQLParserOPTIONS, 0)
}

func (s *Create_collectionStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Create_collectionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Create_collectionStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Create_collectionStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterCreate_collectionStmt(s)
	}
}

func (s *Create_collectionStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitCreate_collectionStmt(s)
	}
}

func (p *SQLParser) Create_collectionStmt() (localctx ICreate_collectionStmtContext) {
	localctx = NewCreate_collectionStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SQLParserRULE_create_collectionStmt)
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
		p.SetState(146)
		p.Match(SQLParserCREATE)
	}
	{
		p.SetState(147)
		p.Match(SQLParserCOLLECTION)
	}
	{
		p.SetState(148)
		p.Collection_section()
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserOPTIONS {
		{
			p.SetState(149)
			p.Match(SQLParserOPTIONS)
		}
		{
			p.SetState(150)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICollection_sectionContext)(nil)).Elem(), 0)

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
	localctx = NewDrop_collectionStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SQLParserRULE_drop_collectionStmt)

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
		p.SetState(153)
		p.Match(SQLParserDROP)
	}
	{
		p.SetState(154)
		p.Match(SQLParserCOLLECTION)
	}
	{
		p.SetState(155)
		p.Collection_section()
	}

	return localctx
}

// ICreate_indexStmtContext is an interface to support dynamic dispatch.
type ICreate_indexStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndex_sectionContext)(nil)).Elem(), 0)

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
	localctx = NewCreate_indexStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SQLParserRULE_create_indexStmt)

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
		p.Match(SQLParserCREATE)
	}
	{
		p.SetState(158)
		p.Match(SQLParserCOLLECTION_INDEX)
	}
	{
		p.SetState(159)
		p.Index_section()
	}

	return localctx
}

// IDrop_indexStmtContext is an interface to support dynamic dispatch.
type IDrop_indexStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndex_sectionContext)(nil)).Elem(), 0)

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
	localctx = NewDrop_indexStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SQLParserRULE_drop_indexStmt)

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
		p.SetState(161)
		p.Match(SQLParserDROP)
	}
	{
		p.SetState(162)
		p.Match(SQLParserCOLLECTION_INDEX)
	}
	{
		p.SetState(163)
		p.Index_section()
	}

	return localctx
}

// ISelectStmtContext is an interface to support dynamic dispatch.
type ISelectStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnsContext)
}

func (s *SelectStmtContext) From_section() IFrom_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFrom_sectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFrom_sectionContext)
}

func (s *SelectStmtContext) Where_section() IWhere_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhere_sectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhere_sectionContext)
}

func (s *SelectStmtContext) Grouping_section() IGrouping_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGrouping_sectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGrouping_sectionContext)
}

func (s *SelectStmtContext) Having_section() IHaving_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHaving_sectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHaving_sectionContext)
}

func (s *SelectStmtContext) Sorting_section() ISorting_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISorting_sectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISorting_sectionContext)
}

func (s *SelectStmtContext) Limit_section() ILimit_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimit_sectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimit_sectionContext)
}

func (s *SelectStmtContext) Offset_section() IOffset_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOffset_sectionContext)(nil)).Elem(), 0)

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
	localctx = NewSelectStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SQLParserRULE_selectStmt)
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
		p.SetState(165)
		p.Match(SQLParserSELECT)
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserALL || _la == SQLParserDISTINCT {
		{
			p.SetState(166)
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
		p.SetState(169)
		p.Columns()
	}
	{
		p.SetState(170)
		p.From_section()
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserWHERE {
		{
			p.SetState(171)
			p.Where_section()
		}

	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserGROUP {
		{
			p.SetState(174)
			p.Grouping_section()
		}

	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserHAVING {
		{
			p.SetState(177)
			p.Having_section()
		}

	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserORDER {
		{
			p.SetState(180)
			p.Sorting_section()
		}

	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserLIMIT {
		{
			p.SetState(183)
			p.Limit_section()
		}

	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserOFFSET {
		{
			p.SetState(186)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumnContext)(nil)).Elem())
	var tst = make([]IColumnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumnContext)
		}
	}

	return tst
}

func (s *ColumnsContext) Column(i int) IColumnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnContext)(nil)).Elem(), i)

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
	localctx = NewColumnsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SQLParserRULE_columns)
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

	p.SetState(198)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLParserASTERISK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Match(SQLParserASTERISK)
		}

	case SQLParserT__0, SQLParserT__2, SQLParserT__4, SQLParserCURRENT_DATE, SQLParserCURRENT_TIME, SQLParserCURRENT_TIMESTAMP, SQLParserNIL, SQLParserTRUE, SQLParserFALSE, SQLParserIDENTIFIER, SQLParserNUMBER, SQLParserFLOAT, SQLParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Column()
		}
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLParserCOMMA {
			{
				p.SetState(191)
				p.Match(SQLParserCOMMA)
			}
			{
				p.SetState(192)
				p.Column()
			}

			p.SetState(197)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITable_nameContext)(nil)).Elem())
	var tst = make([]ITable_nameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITable_nameContext)
		}
	}

	return tst
}

func (s *From_sectionContext) Table_name(i int) ITable_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_nameContext)(nil)).Elem(), i)

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
	localctx = NewFrom_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SQLParserRULE_from_section)
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
		p.SetState(200)
		p.Match(SQLParserFROM)
	}
	{
		p.SetState(201)
		p.Table_name()
	}

	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserCOMMA {
		{
			p.SetState(203)
			p.Match(SQLParserCOMMA)
		}
		{
			p.SetState(204)
			p.Table_name()
		}

		p.SetState(209)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IData_sourceContext)(nil)).Elem(), 0)

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
	localctx = NewTable_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SQLParserRULE_table_name)

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
		p.SetState(210)
		p.Data_source()
	}

	return localctx
}

// IData_sourceContext is an interface to support dynamic dispatch.
type IData_sourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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

func (s *Data_sourceContext) Collection_name() ICollection_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICollection_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICollection_nameContext)
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
	localctx = NewData_sourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SQLParserRULE_data_source)

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
		p.SetState(212)
		p.Collection_name()
	}

	return localctx
}

// IGrouping_sectionContext is an interface to support dynamic dispatch.
type IGrouping_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Grouping_sectionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

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
	localctx = NewGrouping_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SQLParserRULE_grouping_section)
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
		p.SetState(214)
		p.Match(SQLParserGROUP)
	}
	{
		p.SetState(215)
		p.Match(SQLParserBY)
	}
	{
		p.SetState(216)
		p.Expression()
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserCOMMA {
		{
			p.SetState(217)
			p.Match(SQLParserCOMMA)
		}
		{
			p.SetState(218)
			p.Expression()
		}

		p.SetState(223)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

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
	localctx = NewHaving_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SQLParserRULE_having_section)

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
		p.SetState(224)
		p.Match(SQLParserHAVING)
	}
	{
		p.SetState(225)
		p.Expression()
	}

	return localctx
}

// ISorting_sectionContext is an interface to support dynamic dispatch.
type ISorting_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISorting_itemContext)(nil)).Elem())
	var tst = make([]ISorting_itemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISorting_itemContext)
		}
	}

	return tst
}

func (s *Sorting_sectionContext) Sorting_item(i int) ISorting_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISorting_itemContext)(nil)).Elem(), i)

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
	localctx = NewSorting_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SQLParserRULE_sorting_section)
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
		p.SetState(227)
		p.Match(SQLParserORDER)
	}
	{
		p.SetState(228)
		p.Match(SQLParserBY)
	}
	{
		p.SetState(229)
		p.Sorting_item()
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserCOMMA {
		{
			p.SetState(230)
			p.Match(SQLParserCOMMA)
		}
		{
			p.SetState(231)
			p.Sorting_item()
		}

		p.SetState(236)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *Sorting_itemContext) Sorting_specification() ISorting_specificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISorting_specificationContext)(nil)).Elem(), 0)

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
	localctx = NewSorting_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SQLParserRULE_sorting_item)
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
		p.SetState(237)
		p.Property()
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserASC || _la == SQLParserDESC {
		{
			p.SetState(238)
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
	localctx = NewSorting_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SQLParserRULE_sorting_specification)
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
		p.SetState(241)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpression_literalContext)(nil)).Elem())
	var tst = make([]IExpression_literalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpression_literalContext)
		}
	}

	return tst
}

func (s *Limit_sectionContext) Expression_literal(i int) IExpression_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_literalContext)(nil)).Elem(), i)

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
	localctx = NewLimit_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SQLParserRULE_limit_section)

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
		p.SetState(243)
		p.Match(SQLParserLIMIT)
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(244)
			p.Expression_literal()
		}
		{
			p.SetState(245)
			p.Match(SQLParserCOMMA)
		}

	}
	{
		p.SetState(249)
		p.Expression_literal()
	}

	return localctx
}

// IOffset_sectionContext is an interface to support dynamic dispatch.
type IOffset_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	localctx = NewOffset_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SQLParserRULE_offset_section)

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
		p.SetState(251)
		p.Match(SQLParserOFFSET)
	}

	return localctx
}

// IInsertStmtContext is an interface to support dynamic dispatch.
type IInsertStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICollection_sectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICollection_sectionContext)
}

func (s *InsertStmtContext) Insert_values_section() IInsert_values_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInsert_values_sectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInsert_values_sectionContext)
}

func (s *InsertStmtContext) Sync_operator() ISync_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISync_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISync_operatorContext)
}

func (s *InsertStmtContext) Insert_columns_section() IInsert_columns_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInsert_columns_sectionContext)(nil)).Elem(), 0)

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
	localctx = NewInsertStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SQLParserRULE_insertStmt)
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
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserASYNC || _la == SQLParserSYNC {
		{
			p.SetState(253)
			p.Sync_operator()
		}

	}
	{
		p.SetState(256)
		p.Match(SQLParserINSERT)
	}
	{
		p.SetState(257)
		p.Match(SQLParserINTO)
	}
	{
		p.SetState(258)
		p.Collection_section()
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserT__0 {
		{
			p.SetState(259)
			p.Insert_columns_section()
		}

	}
	{
		p.SetState(262)
		p.Insert_values_section()
	}

	return localctx
}

// IInsert_columns_sectionContext is an interface to support dynamic dispatch.
type IInsert_columns_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumnContext)(nil)).Elem())
	var tst = make([]IColumnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumnContext)
		}
	}

	return tst
}

func (s *Insert_columns_sectionContext) Column(i int) IColumnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnContext)(nil)).Elem(), i)

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
	localctx = NewInsert_columns_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SQLParserRULE_insert_columns_section)
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
		p.SetState(264)
		p.Match(SQLParserT__0)
	}
	{
		p.SetState(265)
		p.Column()
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserCOMMA {
		{
			p.SetState(266)
			p.Match(SQLParserCOMMA)
		}
		{
			p.SetState(267)
			p.Column()
		}

		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(273)
		p.Match(SQLParserT__1)
	}

	return localctx
}

// IInsert_values_sectionContext is an interface to support dynamic dispatch.
type IInsert_values_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

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
	localctx = NewInsert_values_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SQLParserRULE_insert_values_section)

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

	p.SetState(279)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLParserVALUE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)
			p.Match(SQLParserVALUE)
		}
		{
			p.SetState(276)
			p.Expression()
		}

	case SQLParserVALUES:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(277)
			p.Match(SQLParserVALUES)
		}
		{
			p.SetState(278)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICollection_sectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICollection_sectionContext)
}

func (s *UpdateStmtContext) SET() antlr.TerminalNode {
	return s.GetToken(SQLParserSET, 0)
}

func (s *UpdateStmtContext) AllProperty_section() []IProperty_sectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProperty_sectionContext)(nil)).Elem())
	var tst = make([]IProperty_sectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProperty_sectionContext)
		}
	}

	return tst
}

func (s *UpdateStmtContext) Property_section(i int) IProperty_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProperty_sectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProperty_sectionContext)
}

func (s *UpdateStmtContext) Sync_operator() ISync_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISync_operatorContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhere_sectionContext)(nil)).Elem(), 0)

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
	localctx = NewUpdateStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SQLParserRULE_updateStmt)
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
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserASYNC || _la == SQLParserSYNC {
		{
			p.SetState(281)
			p.Sync_operator()
		}

	}
	{
		p.SetState(284)
		p.Match(SQLParserUPDATE)
	}
	{
		p.SetState(285)
		p.Collection_section()
	}
	{
		p.SetState(286)
		p.Match(SQLParserSET)
	}
	{
		p.SetState(287)
		p.Property_section()
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
			p.Property_section()
		}

		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserWHERE {
		{
			p.SetState(295)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *Property_sectionContext) SINGLE_EQ() antlr.TerminalNode {
	return s.GetToken(SQLParserSINGLE_EQ, 0)
}

func (s *Property_sectionContext) Expression_literal() IExpression_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_literalContext)(nil)).Elem(), 0)

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
	localctx = NewProperty_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SQLParserRULE_property_section)

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
		p.SetState(298)
		p.Property()
	}
	{
		p.SetState(299)
		p.Match(SQLParserSINGLE_EQ)
	}
	{
		p.SetState(300)
		p.Expression_literal()
	}

	return localctx
}

// IDeleteStmtContext is an interface to support dynamic dispatch.
type IDeleteStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICollection_sectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICollection_sectionContext)
}

func (s *DeleteStmtContext) Sync_operator() ISync_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISync_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISync_operatorContext)
}

func (s *DeleteStmtContext) Where_section() IWhere_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhere_sectionContext)(nil)).Elem(), 0)

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
	localctx = NewDeleteStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SQLParserRULE_deleteStmt)
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
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserASYNC || _la == SQLParserSYNC {
		{
			p.SetState(302)
			p.Sync_operator()
		}

	}
	{
		p.SetState(305)
		p.Match(SQLParserDELETE)
	}
	{
		p.SetState(306)
		p.Match(SQLParserFROM)
	}
	{
		p.SetState(307)
		p.Collection_section()
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserWHERE {
		{
			p.SetState(308)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_listContext)(nil)).Elem(), 0)

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
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SQLParserRULE_expression)

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
		p.SetState(311)
		p.Expression_list()
	}

	return localctx
}

// IExpression_listContext is an interface to support dynamic dispatch.
type IExpression_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_literalContext)
}

func (s *Expression_listContext) Expression_function() IExpression_functionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_functionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_functionContext)
}

func (s *Expression_listContext) AllExpression_binary_operator() []IExpression_binary_operatorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpression_binary_operatorContext)(nil)).Elem())
	var tst = make([]IExpression_binary_operatorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpression_binary_operatorContext)
		}
	}

	return tst
}

func (s *Expression_listContext) Expression_binary_operator(i int) IExpression_binary_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_binary_operatorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpression_binary_operatorContext)
}

func (s *Expression_listContext) AllExpression_logic_operator() []IExpression_logic_operatorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpression_logic_operatorContext)(nil)).Elem())
	var tst = make([]IExpression_logic_operatorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpression_logic_operatorContext)
		}
	}

	return tst
}

func (s *Expression_listContext) Expression_logic_operator(i int) IExpression_logic_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_logic_operatorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpression_logic_operatorContext)
}

func (s *Expression_listContext) AllExpression_array() []IExpression_arrayContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpression_arrayContext)(nil)).Elem())
	var tst = make([]IExpression_arrayContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpression_arrayContext)
		}
	}

	return tst
}

func (s *Expression_listContext) Expression_array(i int) IExpression_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_arrayContext)(nil)).Elem(), i)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpression_dictionaryContext)(nil)).Elem())
	var tst = make([]IExpression_dictionaryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpression_dictionaryContext)
		}
	}

	return tst
}

func (s *Expression_listContext) Expression_dictionary(i int) IExpression_dictionaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_dictionaryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpression_dictionaryContext)
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
	localctx = NewExpression_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SQLParserRULE_expression_list)
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

	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(313)
			p.Expression_literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(314)
			p.Expression_function()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(315)
			p.Expression_binary_operator()
		}
		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLParserAND || _la == SQLParserOR {
			{
				p.SetState(316)
				p.Expression_logic_operator()
			}
			{
				p.SetState(317)
				p.Expression_binary_operator()
			}

			p.SetState(323)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(324)
			p.Match(SQLParserT__0)
		}
		{
			p.SetState(325)
			p.Expression_array()
		}
		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLParserCOMMA {
			{
				p.SetState(326)
				p.Match(SQLParserCOMMA)
			}
			{
				p.SetState(327)
				p.Expression_array()
			}

			p.SetState(332)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(333)
			p.Match(SQLParserT__1)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(335)
			p.Match(SQLParserT__2)
		}

		{
			p.SetState(336)
			p.Expression_dictionary()
		}

		p.SetState(341)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLParserCOMMA {
			{
				p.SetState(337)
				p.Match(SQLParserCOMMA)
			}
			{
				p.SetState(338)
				p.Expression_dictionary()
			}

			p.SetState(343)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(344)
			p.Match(SQLParserT__3)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(346)
			p.Match(SQLParserT__4)
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
			p.Match(SQLParserT__5)
		}

	}

	return localctx
}

// IExpression_literalContext is an interface to support dynamic dispatch.
type IExpression_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_literal_valueContext)(nil)).Elem(), 0)

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
	localctx = NewExpression_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SQLParserRULE_expression_literal)

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
		p.SetState(359)
		p.Expression_literal_value()
	}

	return localctx
}

// IExpression_literal_valueContext is an interface to support dynamic dispatch.
type IExpression_literal_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProperty_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProperty_literalContext)
}

func (s *Expression_literal_valueContext) Integer_literal() IInteger_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_literalContext)
}

func (s *Expression_literal_valueContext) Real_literal() IReal_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReal_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReal_literalContext)
}

func (s *Expression_literal_valueContext) String_literal() IString_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *Expression_literal_valueContext) True_literal() ITrue_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrue_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrue_literalContext)
}

func (s *Expression_literal_valueContext) False_literal() IFalse_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFalse_literalContext)(nil)).Elem(), 0)

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
	localctx = NewExpression_literal_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SQLParserRULE_expression_literal_value)

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

	p.SetState(371)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(361)
			p.Property_literal()
		}

	case SQLParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(362)
			p.Integer_literal()
		}

	case SQLParserFLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(363)
			p.Real_literal()
		}

	case SQLParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(364)
			p.String_literal()
		}

	case SQLParserTRUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(365)
			p.True_literal()
		}

	case SQLParserFALSE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(366)
			p.False_literal()
		}

	case SQLParserNIL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(367)
			p.Match(SQLParserNIL)
		}

	case SQLParserCURRENT_TIME:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(368)
			p.Match(SQLParserCURRENT_TIME)
		}

	case SQLParserCURRENT_DATE:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(369)
			p.Match(SQLParserCURRENT_DATE)
		}

	case SQLParserCURRENT_TIMESTAMP:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(370)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Expression_dictionaryContext) Expression_literal() IExpression_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_literalContext)(nil)).Elem(), 0)

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
	localctx = NewExpression_dictionaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SQLParserRULE_expression_dictionary)

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
		p.SetState(373)
		p.Name()
	}
	{
		p.SetState(374)
		p.Match(SQLParserT__6)
	}
	{
		p.SetState(375)
		p.Expression_literal()
	}

	return localctx
}

// IDictionary_literalContext is an interface to support dynamic dispatch.
type IDictionary_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Dictionary_literalContext) Expression_literal() IExpression_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_literalContext)(nil)).Elem(), 0)

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
	localctx = NewDictionary_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SQLParserRULE_dictionary_literal)

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
		p.SetState(377)
		p.Name()
	}
	{
		p.SetState(378)
		p.Match(SQLParserT__6)
	}
	{
		p.SetState(379)
		p.Expression_literal()
	}

	return localctx
}

// IExpression_arrayContext is an interface to support dynamic dispatch.
type IExpression_arrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_literalContext)(nil)).Elem(), 0)

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
	localctx = NewExpression_arrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SQLParserRULE_expression_array)

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
		p.SetState(381)
		p.Expression_literal()
	}

	return localctx
}

// IArray_literalContext is an interface to support dynamic dispatch.
type IArray_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_literalContext)(nil)).Elem(), 0)

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
	localctx = NewArray_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SQLParserRULE_array_literal)

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
		p.SetState(383)
		p.Expression_literal()
	}

	return localctx
}

// IExpression_logic_operatorContext is an interface to support dynamic dispatch.
type IExpression_logic_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_operatorContext)(nil)).Elem(), 0)

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
	localctx = NewExpression_logic_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SQLParserRULE_expression_logic_operator)

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
		p.SetState(385)
		p.Logical_operator()
	}

	return localctx
}

// IExpression_binary_operatorContext is an interface to support dynamic dispatch.
type IExpression_binary_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_operatorContext)(nil)).Elem(), 0)

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
	localctx = NewExpression_binary_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SQLParserRULE_expression_binary_operator)

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
		p.SetState(387)
		p.Expression_operator()
	}

	return localctx
}

// IExpression_functionContext is an interface to support dynamic dispatch.
type IExpression_functionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_valueContext)(nil)).Elem(), 0)

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
	localctx = NewExpression_functionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SQLParserRULE_expression_function)
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
		p.SetState(389)
		p.Match(SQLParserIDENTIFIER)
	}
	{
		p.SetState(390)
		p.Match(SQLParserT__0)
	}
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SQLParserT__0)|(1<<SQLParserT__2)|(1<<SQLParserT__4)|(1<<SQLParserASTERISK)|(1<<SQLParserCURRENT_DATE)|(1<<SQLParserCURRENT_TIME)|(1<<SQLParserCURRENT_TIMESTAMP))) != 0) || (((_la-48)&-(0x1f+1)) == 0 && ((1<<uint((_la-48)))&((1<<(SQLParserNIL-48))|(1<<(SQLParserTRUE-48))|(1<<(SQLParserFALSE-48))|(1<<(SQLParserIDENTIFIER-48))|(1<<(SQLParserNUMBER-48))|(1<<(SQLParserFLOAT-48))|(1<<(SQLParserSTRING-48)))) != 0) {
		{
			p.SetState(391)
			p.Function_value()
		}

	}
	{
		p.SetState(394)
		p.Match(SQLParserT__1)
	}

	return localctx
}

// IFunction_nameContext is an interface to support dynamic dispatch.
type IFunction_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	localctx = NewFunction_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SQLParserRULE_function_name)

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
		p.Match(SQLParserIDENTIFIER)
	}

	return localctx
}

// IFunction_valueContext is an interface to support dynamic dispatch.
type IFunction_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Function_valueContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

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
	localctx = NewFunction_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SQLParserRULE_function_value)
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

	p.SetState(407)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLParserT__0, SQLParserT__2, SQLParserT__4, SQLParserCURRENT_DATE, SQLParserCURRENT_TIME, SQLParserCURRENT_TIMESTAMP, SQLParserNIL, SQLParserTRUE, SQLParserFALSE, SQLParserIDENTIFIER, SQLParserNUMBER, SQLParserFLOAT, SQLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(398)
			p.Expression()
		}
		p.SetState(403)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLParserCOMMA {
			{
				p.SetState(399)
				p.Match(SQLParserCOMMA)
			}
			{
				p.SetState(400)
				p.Expression()
			}

			p.SetState(405)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case SQLParserASTERISK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(406)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpression_literalContext)(nil)).Elem())
	var tst = make([]IExpression_literalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpression_literalContext)
		}
	}

	return tst
}

func (s *Expression_operatorContext) Expression_literal(i int) IExpression_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_literalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpression_literalContext)
}

func (s *Expression_operatorContext) Binary_operator() IBinary_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_operatorContext)(nil)).Elem(), 0)

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
	localctx = NewExpression_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SQLParserRULE_expression_operator)

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
		p.SetState(409)
		p.Expression_literal()
	}
	{
		p.SetState(410)
		p.Binary_operator()
	}
	{
		p.SetState(411)
		p.Expression_literal()
	}

	return localctx
}

// IBinary_operatorContext is an interface to support dynamic dispatch.
type IBinary_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	localctx = NewBinary_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SQLParserRULE_binary_operator)
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
		p.SetState(413)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SQLParserSINGLE_EQ)|(1<<SQLParserDOUBLE_EQ)|(1<<SQLParserOP_LT)|(1<<SQLParserLE)|(1<<SQLParserGT)|(1<<SQLParserGE)|(1<<SQLParserNOTEQ))) != 0) {
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
	localctx = NewLogical_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, SQLParserRULE_logical_operator)
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
		p.SetState(415)
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
	localctx = NewProperty_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, SQLParserRULE_property_literal)

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
		p.SetState(417)
		p.Match(SQLParserIDENTIFIER)
	}

	return localctx
}

// IInteger_literalContext is an interface to support dynamic dispatch.
type IInteger_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	localctx = NewInteger_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, SQLParserRULE_integer_literal)

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
		p.Match(SQLParserNUMBER)
	}

	return localctx
}

// IReal_literalContext is an interface to support dynamic dispatch.
type IReal_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	localctx = NewReal_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, SQLParserRULE_real_literal)

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
		p.SetState(421)
		p.Match(SQLParserFLOAT)
	}

	return localctx
}

// IString_literalContext is an interface to support dynamic dispatch.
type IString_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	localctx = NewString_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, SQLParserRULE_string_literal)

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
		p.SetState(423)
		p.Match(SQLParserSTRING)
	}

	return localctx
}

// ITrue_literalContext is an interface to support dynamic dispatch.
type ITrue_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	localctx = NewTrue_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, SQLParserRULE_true_literal)

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
		p.SetState(425)
		p.Match(SQLParserTRUE)
	}

	return localctx
}

// IFalse_literalContext is an interface to support dynamic dispatch.
type IFalse_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	localctx = NewFalse_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, SQLParserRULE_false_literal)

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
		p.SetState(427)
		p.Match(SQLParserFALSE)
	}

	return localctx
}

// ISync_operatorContext is an interface to support dynamic dispatch.
type ISync_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	localctx = NewSync_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, SQLParserRULE_sync_operator)
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
		p.SetState(429)
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
	localctx = NewCompound_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, SQLParserRULE_compound_operator)
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

	p.SetState(437)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLParserUNION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(431)
			p.Match(SQLParserUNION)
		}
		p.SetState(433)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLParserALL {
			{
				p.SetState(432)
				p.Match(SQLParserALL)
			}

		}

	case SQLParserINTERSECT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(435)
			p.Match(SQLParserINTERSECT)
		}

	case SQLParserEXCEPT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(436)
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
	localctx = NewCondition_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, SQLParserRULE_condition_operator)
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
		p.SetState(439)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SQLParserSINGLE_EQ)|(1<<SQLParserDOUBLE_EQ)|(1<<SQLParserOP_LT)|(1<<SQLParserLE)|(1<<SQLParserGT)|(1<<SQLParserGE)|(1<<SQLParserNOTEQ))) != 0) {
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
	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, SQLParserRULE_property)

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
		p.SetState(441)
		p.Match(SQLParserIDENTIFIER)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, SQLParserRULE_value)

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
		p.SetState(443)
		p.Match(SQLParserIDENTIFIER)
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, SQLParserRULE_name)

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
		p.SetState(445)
		p.Match(SQLParserIDENTIFIER)
	}

	return localctx
}

// ICollection_sectionContext is an interface to support dynamic dispatch.
type ICollection_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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

func (s *Collection_sectionContext) Collection_name() ICollection_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICollection_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICollection_nameContext)
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
	localctx = NewCollection_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, SQLParserRULE_collection_section)

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
		p.SetState(447)
		p.Collection_name()
	}

	return localctx
}

// ICollection_nameContext is an interface to support dynamic dispatch.
type ICollection_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCollection_nameContext differentiates from other interfaces.
	IsCollection_nameContext()
}

type Collection_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollection_nameContext() *Collection_nameContext {
	var p = new(Collection_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_collection_name
	return p
}

func (*Collection_nameContext) IsCollection_nameContext() {}

func NewCollection_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Collection_nameContext {
	var p = new(Collection_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_collection_name

	return p
}

func (s *Collection_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Collection_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLParserIDENTIFIER, 0)
}

func (s *Collection_nameContext) String_literal() IString_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *Collection_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Collection_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Collection_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterCollection_name(s)
	}
}

func (s *Collection_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitCollection_name(s)
	}
}

func (p *SQLParser) Collection_name() (localctx ICollection_nameContext) {
	localctx = NewCollection_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, SQLParserRULE_collection_name)

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

	p.SetState(451)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(449)
			p.Match(SQLParserIDENTIFIER)
		}

	case SQLParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(450)
			p.String_literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IColumnContext is an interface to support dynamic dispatch.
type IColumnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ColumnContext) AS() antlr.TerminalNode {
	return s.GetToken(SQLParserAS, 0)
}

func (s *ColumnContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

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
	localctx = NewColumnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, SQLParserRULE_column)
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
		p.SetState(453)
		p.Expression()
	}

	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserAS {
		{
			p.SetState(454)
			p.Match(SQLParserAS)
		}
		{
			p.SetState(455)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndex_nameContext)(nil)).Elem(), 0)

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
	localctx = NewIndex_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, SQLParserRULE_index_section)

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
		p.SetState(458)
		p.Index_name()
	}

	return localctx
}

// IIndex_nameContext is an interface to support dynamic dispatch.
type IIndex_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_literalContext)(nil)).Elem(), 0)

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
	localctx = NewIndex_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, SQLParserRULE_index_name)

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

	p.SetState(462)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(460)
			p.Match(SQLParserIDENTIFIER)
		}

	case SQLParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(461)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

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
	localctx = NewWhere_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, SQLParserRULE_where_section)

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
		p.Match(SQLParserWHERE)
	}
	{
		p.SetState(465)
		p.Expression()
	}

	return localctx
}
