// Code generated from SQL.g4 by ANTLR 4.7.2. DO NOT EDIT.

package antlr

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 70, 655,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4,
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86,
	9, 86, 4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 4, 90, 9, 90, 4, 91, 9,
	91, 4, 92, 9, 92, 4, 93, 9, 93, 4, 94, 9, 94, 4, 95, 9, 95, 4, 96, 9, 96,
	4, 97, 9, 97, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 235, 10, 17,
	3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 241, 10, 18, 3, 19, 3, 19, 3, 20, 3,
	20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25,
	3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3,
	31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36,
	3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3,
	41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46,
	3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3,
	48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 51,
	3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3,
	53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54,
	3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3,
	55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56,
	3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3,
	57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57,
	3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58, 3,
	58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60,
	3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 3,
	61, 3, 61, 3, 62, 3, 62, 3, 62, 3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 3, 63,
	3, 63, 3, 63, 3, 63, 3, 64, 3, 64, 3, 64, 3, 64, 3, 64, 3, 64, 3, 64, 3,
	64, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66,
	3, 66, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 68, 3, 68, 3,
	68, 3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3, 70, 3, 70, 3, 70, 3, 70,
	3, 70, 3, 70, 3, 70, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3,
	71, 3, 71, 3, 71, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 73, 3, 73, 3, 73,
	3, 74, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74, 3, 75, 3, 75, 3, 75, 3, 75, 3,
	75, 3, 76, 3, 76, 3, 76, 3, 76, 3, 76, 3, 76, 3, 76, 3, 77, 3, 77, 3, 77,
	3, 77, 3, 77, 3, 77, 3, 77, 3, 77, 3, 78, 3, 78, 3, 78, 3, 78, 3, 78, 3,
	78, 3, 79, 3, 79, 3, 79, 3, 79, 3, 79, 3, 79, 3, 79, 3, 80, 3, 80, 3, 80,
	3, 80, 3, 81, 3, 81, 3, 81, 3, 81, 3, 81, 3, 82, 3, 82, 3, 82, 3, 82, 3,
	82, 3, 83, 3, 83, 3, 83, 3, 83, 3, 84, 3, 84, 3, 84, 3, 84, 3, 84, 3, 84,
	3, 85, 3, 85, 3, 85, 3, 85, 3, 85, 3, 85, 3, 85, 3, 86, 3, 86, 3, 86, 3,
	86, 3, 86, 3, 86, 3, 87, 3, 87, 3, 87, 3, 87, 3, 87, 3, 87, 3, 88, 3, 88,
	3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 89, 3, 89, 3, 89, 3, 89, 3, 89, 3,
	90, 3, 90, 3, 90, 3, 90, 3, 90, 3, 90, 3, 91, 3, 91, 3, 92, 3, 92, 7, 92,
	592, 10, 92, 12, 92, 14, 92, 595, 11, 92, 3, 93, 6, 93, 598, 10, 93, 13,
	93, 14, 93, 599, 3, 94, 6, 94, 603, 10, 94, 13, 94, 14, 94, 604, 3, 94,
	3, 94, 7, 94, 609, 10, 94, 12, 94, 14, 94, 612, 11, 94, 3, 94, 5, 94, 615,
	10, 94, 3, 94, 3, 94, 6, 94, 619, 10, 94, 13, 94, 14, 94, 620, 3, 94, 5,
	94, 624, 10, 94, 3, 94, 6, 94, 627, 10, 94, 13, 94, 14, 94, 628, 3, 94,
	5, 94, 632, 10, 94, 3, 95, 3, 95, 3, 95, 3, 95, 7, 95, 638, 10, 95, 12,
	95, 14, 95, 641, 11, 95, 3, 95, 3, 95, 3, 96, 3, 96, 5, 96, 647, 10, 96,
	3, 96, 6, 96, 650, 10, 96, 13, 96, 14, 96, 651, 3, 97, 3, 97, 2, 2, 98,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41,
	2, 43, 2, 45, 2, 47, 2, 49, 2, 51, 2, 53, 2, 55, 2, 57, 2, 59, 2, 61, 2,
	63, 2, 65, 2, 67, 2, 69, 2, 71, 2, 73, 2, 75, 2, 77, 2, 79, 2, 81, 2, 83,
	2, 85, 2, 87, 2, 89, 2, 91, 2, 93, 22, 95, 23, 97, 24, 99, 25, 101, 26,
	103, 27, 105, 28, 107, 29, 109, 30, 111, 31, 113, 32, 115, 33, 117, 34,
	119, 35, 121, 36, 123, 37, 125, 38, 127, 39, 129, 40, 131, 41, 133, 42,
	135, 43, 137, 44, 139, 45, 141, 46, 143, 47, 145, 48, 147, 49, 149, 50,
	151, 51, 153, 52, 155, 53, 157, 54, 159, 55, 161, 56, 163, 57, 165, 58,
	167, 59, 169, 60, 171, 61, 173, 62, 175, 63, 177, 64, 179, 65, 181, 66,
	183, 67, 185, 68, 187, 69, 189, 70, 191, 2, 193, 2, 3, 2, 34, 4, 2, 67,
	67, 99, 99, 4, 2, 68, 68, 100, 100, 4, 2, 69, 69, 101, 101, 4, 2, 70, 70,
	102, 102, 4, 2, 71, 71, 103, 103, 4, 2, 72, 72, 104, 104, 4, 2, 73, 73,
	105, 105, 4, 2, 74, 74, 106, 106, 4, 2, 75, 75, 107, 107, 4, 2, 76, 76,
	108, 108, 4, 2, 77, 77, 109, 109, 4, 2, 78, 78, 110, 110, 4, 2, 79, 79,
	111, 111, 4, 2, 80, 80, 112, 112, 4, 2, 81, 81, 113, 113, 4, 2, 82, 82,
	114, 114, 4, 2, 83, 83, 115, 115, 4, 2, 84, 84, 116, 116, 4, 2, 85, 85,
	117, 117, 4, 2, 86, 86, 118, 118, 4, 2, 87, 87, 119, 119, 4, 2, 88, 88,
	120, 120, 4, 2, 89, 89, 121, 121, 4, 2, 90, 90, 122, 122, 4, 2, 91, 91,
	123, 123, 4, 2, 92, 92, 124, 124, 5, 2, 11, 12, 15, 15, 34, 34, 5, 2, 67,
	92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 41, 41,
	4, 2, 45, 45, 47, 47, 5, 2, 50, 59, 67, 72, 99, 104, 2, 642, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3,
	2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27,
	3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2,
	35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2,
	2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2,
	2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109,
	3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2,
	2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3,
	2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2,
	131, 3, 2, 2, 2, 2, 133, 3, 2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137, 3, 2,
	2, 2, 2, 139, 3, 2, 2, 2, 2, 141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2, 2, 145,
	3, 2, 2, 2, 2, 147, 3, 2, 2, 2, 2, 149, 3, 2, 2, 2, 2, 151, 3, 2, 2, 2,
	2, 153, 3, 2, 2, 2, 2, 155, 3, 2, 2, 2, 2, 157, 3, 2, 2, 2, 2, 159, 3,
	2, 2, 2, 2, 161, 3, 2, 2, 2, 2, 163, 3, 2, 2, 2, 2, 165, 3, 2, 2, 2, 2,
	167, 3, 2, 2, 2, 2, 169, 3, 2, 2, 2, 2, 171, 3, 2, 2, 2, 2, 173, 3, 2,
	2, 2, 2, 175, 3, 2, 2, 2, 2, 177, 3, 2, 2, 2, 2, 179, 3, 2, 2, 2, 2, 181,
	3, 2, 2, 2, 2, 183, 3, 2, 2, 2, 2, 185, 3, 2, 2, 2, 2, 187, 3, 2, 2, 2,
	2, 189, 3, 2, 2, 2, 3, 195, 3, 2, 2, 2, 5, 197, 3, 2, 2, 2, 7, 199, 3,
	2, 2, 2, 9, 201, 3, 2, 2, 2, 11, 203, 3, 2, 2, 2, 13, 205, 3, 2, 2, 2,
	15, 207, 3, 2, 2, 2, 17, 209, 3, 2, 2, 2, 19, 211, 3, 2, 2, 2, 21, 213,
	3, 2, 2, 2, 23, 216, 3, 2, 2, 2, 25, 218, 3, 2, 2, 2, 27, 221, 3, 2, 2,
	2, 29, 223, 3, 2, 2, 2, 31, 226, 3, 2, 2, 2, 33, 234, 3, 2, 2, 2, 35, 240,
	3, 2, 2, 2, 37, 242, 3, 2, 2, 2, 39, 244, 3, 2, 2, 2, 41, 246, 3, 2, 2,
	2, 43, 248, 3, 2, 2, 2, 45, 250, 3, 2, 2, 2, 47, 252, 3, 2, 2, 2, 49, 254,
	3, 2, 2, 2, 51, 256, 3, 2, 2, 2, 53, 258, 3, 2, 2, 2, 55, 260, 3, 2, 2,
	2, 57, 262, 3, 2, 2, 2, 59, 264, 3, 2, 2, 2, 61, 266, 3, 2, 2, 2, 63, 268,
	3, 2, 2, 2, 65, 270, 3, 2, 2, 2, 67, 272, 3, 2, 2, 2, 69, 274, 3, 2, 2,
	2, 71, 276, 3, 2, 2, 2, 73, 278, 3, 2, 2, 2, 75, 280, 3, 2, 2, 2, 77, 282,
	3, 2, 2, 2, 79, 284, 3, 2, 2, 2, 81, 286, 3, 2, 2, 2, 83, 288, 3, 2, 2,
	2, 85, 290, 3, 2, 2, 2, 87, 292, 3, 2, 2, 2, 89, 294, 3, 2, 2, 2, 91, 296,
	3, 2, 2, 2, 93, 298, 3, 2, 2, 2, 95, 302, 3, 2, 2, 2, 97, 311, 3, 2, 2,
	2, 99, 314, 3, 2, 2, 2, 101, 318, 3, 2, 2, 2, 103, 324, 3, 2, 2, 2, 105,
	327, 3, 2, 2, 2, 107, 334, 3, 2, 2, 2, 109, 345, 3, 2, 2, 2, 111, 358,
	3, 2, 2, 2, 113, 371, 3, 2, 2, 2, 115, 389, 3, 2, 2, 2, 117, 394, 3, 2,
	2, 2, 119, 401, 3, 2, 2, 2, 121, 410, 3, 2, 2, 2, 123, 415, 3, 2, 2, 2,
	125, 420, 3, 2, 2, 2, 127, 427, 3, 2, 2, 2, 129, 435, 3, 2, 2, 2, 131,
	440, 3, 2, 2, 2, 133, 446, 3, 2, 2, 2, 135, 453, 3, 2, 2, 2, 137, 456,
	3, 2, 2, 2, 139, 462, 3, 2, 2, 2, 141, 469, 3, 2, 2, 2, 143, 479, 3, 2,
	2, 2, 145, 484, 3, 2, 2, 2, 147, 487, 3, 2, 2, 2, 149, 493, 3, 2, 2, 2,
	151, 498, 3, 2, 2, 2, 153, 505, 3, 2, 2, 2, 155, 513, 3, 2, 2, 2, 157,
	519, 3, 2, 2, 2, 159, 526, 3, 2, 2, 2, 161, 530, 3, 2, 2, 2, 163, 535,
	3, 2, 2, 2, 165, 540, 3, 2, 2, 2, 167, 544, 3, 2, 2, 2, 169, 550, 3, 2,
	2, 2, 171, 557, 3, 2, 2, 2, 173, 563, 3, 2, 2, 2, 175, 569, 3, 2, 2, 2,
	177, 576, 3, 2, 2, 2, 179, 581, 3, 2, 2, 2, 181, 587, 3, 2, 2, 2, 183,
	589, 3, 2, 2, 2, 185, 597, 3, 2, 2, 2, 187, 631, 3, 2, 2, 2, 189, 633,
	3, 2, 2, 2, 191, 644, 3, 2, 2, 2, 193, 653, 3, 2, 2, 2, 195, 196, 7, 42,
	2, 2, 196, 4, 3, 2, 2, 2, 197, 198, 7, 43, 2, 2, 198, 6, 3, 2, 2, 2, 199,
	200, 7, 125, 2, 2, 200, 8, 3, 2, 2, 2, 201, 202, 7, 127, 2, 2, 202, 10,
	3, 2, 2, 2, 203, 204, 7, 93, 2, 2, 204, 12, 3, 2, 2, 2, 205, 206, 7, 95,
	2, 2, 206, 14, 3, 2, 2, 2, 207, 208, 7, 60, 2, 2, 208, 16, 3, 2, 2, 2,
	209, 210, 7, 44, 2, 2, 210, 18, 3, 2, 2, 2, 211, 212, 7, 63, 2, 2, 212,
	20, 3, 2, 2, 2, 213, 214, 7, 63, 2, 2, 214, 215, 7, 63, 2, 2, 215, 22,
	3, 2, 2, 2, 216, 217, 7, 62, 2, 2, 217, 24, 3, 2, 2, 2, 218, 219, 7, 62,
	2, 2, 219, 220, 7, 63, 2, 2, 220, 26, 3, 2, 2, 2, 221, 222, 7, 64, 2, 2,
	222, 28, 3, 2, 2, 2, 223, 224, 7, 64, 2, 2, 224, 225, 7, 63, 2, 2, 225,
	30, 3, 2, 2, 2, 226, 227, 7, 35, 2, 2, 227, 228, 7, 63, 2, 2, 228, 32,
	3, 2, 2, 2, 229, 230, 5, 41, 21, 2, 230, 231, 5, 67, 34, 2, 231, 232, 5,
	47, 24, 2, 232, 235, 3, 2, 2, 2, 233, 235, 7, 40, 2, 2, 234, 229, 3, 2,
	2, 2, 234, 233, 3, 2, 2, 2, 235, 34, 3, 2, 2, 2, 236, 237, 5, 69, 35, 2,
	237, 238, 5, 75, 38, 2, 238, 241, 3, 2, 2, 2, 239, 241, 7, 126, 2, 2, 240,
	236, 3, 2, 2, 2, 240, 239, 3, 2, 2, 2, 241, 36, 3, 2, 2, 2, 242, 243, 7,
	46, 2, 2, 243, 38, 3, 2, 2, 2, 244, 245, 7, 61, 2, 2, 245, 40, 3, 2, 2,
	2, 246, 247, 9, 2, 2, 2, 247, 42, 3, 2, 2, 2, 248, 249, 9, 3, 2, 2, 249,
	44, 3, 2, 2, 2, 250, 251, 9, 4, 2, 2, 251, 46, 3, 2, 2, 2, 252, 253, 9,
	5, 2, 2, 253, 48, 3, 2, 2, 2, 254, 255, 9, 6, 2, 2, 255, 50, 3, 2, 2, 2,
	256, 257, 9, 7, 2, 2, 257, 52, 3, 2, 2, 2, 258, 259, 9, 8, 2, 2, 259, 54,
	3, 2, 2, 2, 260, 261, 9, 9, 2, 2, 261, 56, 3, 2, 2, 2, 262, 263, 9, 10,
	2, 2, 263, 58, 3, 2, 2, 2, 264, 265, 9, 11, 2, 2, 265, 60, 3, 2, 2, 2,
	266, 267, 9, 12, 2, 2, 267, 62, 3, 2, 2, 2, 268, 269, 9, 13, 2, 2, 269,
	64, 3, 2, 2, 2, 270, 271, 9, 14, 2, 2, 271, 66, 3, 2, 2, 2, 272, 273, 9,
	15, 2, 2, 273, 68, 3, 2, 2, 2, 274, 275, 9, 16, 2, 2, 275, 70, 3, 2, 2,
	2, 276, 277, 9, 17, 2, 2, 277, 72, 3, 2, 2, 2, 278, 279, 9, 18, 2, 2, 279,
	74, 3, 2, 2, 2, 280, 281, 9, 19, 2, 2, 281, 76, 3, 2, 2, 2, 282, 283, 9,
	20, 2, 2, 283, 78, 3, 2, 2, 2, 284, 285, 9, 21, 2, 2, 285, 80, 3, 2, 2,
	2, 286, 287, 9, 22, 2, 2, 287, 82, 3, 2, 2, 2, 288, 289, 9, 23, 2, 2, 289,
	84, 3, 2, 2, 2, 290, 291, 9, 24, 2, 2, 291, 86, 3, 2, 2, 2, 292, 293, 9,
	25, 2, 2, 293, 88, 3, 2, 2, 2, 294, 295, 9, 26, 2, 2, 295, 90, 3, 2, 2,
	2, 296, 297, 9, 27, 2, 2, 297, 92, 3, 2, 2, 2, 298, 299, 5, 41, 21, 2,
	299, 300, 5, 63, 32, 2, 300, 301, 5, 63, 32, 2, 301, 94, 3, 2, 2, 2, 302,
	303, 5, 41, 21, 2, 303, 304, 5, 67, 34, 2, 304, 305, 5, 45, 23, 2, 305,
	306, 5, 49, 25, 2, 306, 307, 5, 77, 39, 2, 307, 308, 5, 79, 40, 2, 308,
	309, 5, 69, 35, 2, 309, 310, 5, 75, 38, 2, 310, 96, 3, 2, 2, 2, 311, 312,
	5, 41, 21, 2, 312, 313, 5, 77, 39, 2, 313, 98, 3, 2, 2, 2, 314, 315, 5,
	41, 21, 2, 315, 316, 5, 77, 39, 2, 316, 317, 5, 45, 23, 2, 317, 100, 3,
	2, 2, 2, 318, 319, 5, 41, 21, 2, 319, 320, 5, 77, 39, 2, 320, 321, 5, 89,
	45, 2, 321, 322, 5, 67, 34, 2, 322, 323, 5, 45, 23, 2, 323, 102, 3, 2,
	2, 2, 324, 325, 5, 43, 22, 2, 325, 326, 5, 89, 45, 2, 326, 104, 3, 2, 2,
	2, 327, 328, 5, 45, 23, 2, 328, 329, 5, 75, 38, 2, 329, 330, 5, 49, 25,
	2, 330, 331, 5, 41, 21, 2, 331, 332, 5, 79, 40, 2, 332, 333, 5, 49, 25,
	2, 333, 106, 3, 2, 2, 2, 334, 335, 5, 45, 23, 2, 335, 336, 5, 69, 35, 2,
	336, 337, 5, 63, 32, 2, 337, 338, 5, 63, 32, 2, 338, 339, 5, 49, 25, 2,
	339, 340, 5, 45, 23, 2, 340, 341, 5, 79, 40, 2, 341, 342, 5, 57, 29, 2,
	342, 343, 5, 69, 35, 2, 343, 344, 5, 67, 34, 2, 344, 108, 3, 2, 2, 2, 345,
	346, 5, 45, 23, 2, 346, 347, 5, 81, 41, 2, 347, 348, 5, 75, 38, 2, 348,
	349, 5, 75, 38, 2, 349, 350, 5, 49, 25, 2, 350, 351, 5, 67, 34, 2, 351,
	352, 5, 79, 40, 2, 352, 353, 7, 97, 2, 2, 353, 354, 5, 47, 24, 2, 354,
	355, 5, 41, 21, 2, 355, 356, 5, 79, 40, 2, 356, 357, 5, 49, 25, 2, 357,
	110, 3, 2, 2, 2, 358, 359, 5, 45, 23, 2, 359, 360, 5, 81, 41, 2, 360, 361,
	5, 75, 38, 2, 361, 362, 5, 75, 38, 2, 362, 363, 5, 49, 25, 2, 363, 364,
	5, 67, 34, 2, 364, 365, 5, 79, 40, 2, 365, 366, 7, 97, 2, 2, 366, 367,
	5, 79, 40, 2, 367, 368, 5, 57, 29, 2, 368, 369, 5, 65, 33, 2, 369, 370,
	5, 49, 25, 2, 370, 112, 3, 2, 2, 2, 371, 372, 5, 45, 23, 2, 372, 373, 5,
	81, 41, 2, 373, 374, 5, 75, 38, 2, 374, 375, 5, 75, 38, 2, 375, 376, 5,
	49, 25, 2, 376, 377, 5, 67, 34, 2, 377, 378, 5, 79, 40, 2, 378, 379, 7,
	97, 2, 2, 379, 380, 5, 79, 40, 2, 380, 381, 5, 57, 29, 2, 381, 382, 5,
	65, 33, 2, 382, 383, 5, 49, 25, 2, 383, 384, 5, 77, 39, 2, 384, 385, 5,
	79, 40, 2, 385, 386, 5, 41, 21, 2, 386, 387, 5, 65, 33, 2, 387, 388, 5,
	71, 36, 2, 388, 114, 3, 2, 2, 2, 389, 390, 5, 47, 24, 2, 390, 391, 5, 49,
	25, 2, 391, 392, 5, 77, 39, 2, 392, 393, 5, 45, 23, 2, 393, 116, 3, 2,
	2, 2, 394, 395, 5, 47, 24, 2, 395, 396, 5, 49, 25, 2, 396, 397, 5, 63,
	32, 2, 397, 398, 5, 49, 25, 2, 398, 399, 5, 79, 40, 2, 399, 400, 5, 49,
	25, 2, 400, 118, 3, 2, 2, 2, 401, 402, 5, 47, 24, 2, 402, 403, 5, 57, 29,
	2, 403, 404, 5, 77, 39, 2, 404, 405, 5, 79, 40, 2, 405, 406, 5, 57, 29,
	2, 406, 407, 5, 67, 34, 2, 407, 408, 5, 45, 23, 2, 408, 409, 5, 79, 40,
	2, 409, 120, 3, 2, 2, 2, 410, 411, 5, 47, 24, 2, 411, 412, 5, 75, 38, 2,
	412, 413, 5, 69, 35, 2, 413, 414, 5, 71, 36, 2, 414, 122, 3, 2, 2, 2, 415,
	416, 5, 49, 25, 2, 416, 417, 5, 41, 21, 2, 417, 418, 5, 45, 23, 2, 418,
	419, 5, 55, 28, 2, 419, 124, 3, 2, 2, 2, 420, 421, 5, 49, 25, 2, 421, 422,
	5, 87, 44, 2, 422, 423, 5, 45, 23, 2, 423, 424, 5, 49, 25, 2, 424, 425,
	5, 71, 36, 2, 425, 426, 5, 79, 40, 2, 426, 126, 3, 2, 2, 2, 427, 428, 5,
	51, 26, 2, 428, 429, 5, 63, 32, 2, 429, 430, 5, 41, 21, 2, 430, 431, 5,
	79, 40, 2, 431, 432, 5, 79, 40, 2, 432, 433, 5, 49, 25, 2, 433, 434, 5,
	67, 34, 2, 434, 128, 3, 2, 2, 2, 435, 436, 5, 51, 26, 2, 436, 437, 5, 75,
	38, 2, 437, 438, 5, 69, 35, 2, 438, 439, 5, 65, 33, 2, 439, 130, 3, 2,
	2, 2, 440, 441, 5, 53, 27, 2, 441, 442, 5, 75, 38, 2, 442, 443, 5, 69,
	35, 2, 443, 444, 5, 81, 41, 2, 444, 445, 5, 71, 36, 2, 445, 132, 3, 2,
	2, 2, 446, 447, 5, 55, 28, 2, 447, 448, 5, 41, 21, 2, 448, 449, 5, 83,
	42, 2, 449, 450, 5, 57, 29, 2, 450, 451, 5, 67, 34, 2, 451, 452, 5, 53,
	27, 2, 452, 134, 3, 2, 2, 2, 453, 454, 5, 57, 29, 2, 454, 455, 5, 67, 34,
	2, 455, 136, 3, 2, 2, 2, 456, 457, 5, 57, 29, 2, 457, 458, 5, 67, 34, 2,
	458, 459, 5, 47, 24, 2, 459, 460, 5, 49, 25, 2, 460, 461, 5, 87, 44, 2,
	461, 138, 3, 2, 2, 2, 462, 463, 5, 57, 29, 2, 463, 464, 5, 67, 34, 2, 464,
	465, 5, 77, 39, 2, 465, 466, 5, 49, 25, 2, 466, 467, 5, 75, 38, 2, 467,
	468, 5, 79, 40, 2, 468, 140, 3, 2, 2, 2, 469, 470, 5, 57, 29, 2, 470, 471,
	5, 67, 34, 2, 471, 472, 5, 79, 40, 2, 472, 473, 5, 49, 25, 2, 473, 474,
	5, 75, 38, 2, 474, 475, 5, 77, 39, 2, 475, 476, 5, 49, 25, 2, 476, 477,
	5, 45, 23, 2, 477, 478, 5, 79, 40, 2, 478, 142, 3, 2, 2, 2, 479, 480, 5,
	57, 29, 2, 480, 481, 5, 67, 34, 2, 481, 482, 5, 79, 40, 2, 482, 483, 5,
	69, 35, 2, 483, 144, 3, 2, 2, 2, 484, 485, 5, 57, 29, 2, 485, 486, 5, 77,
	39, 2, 486, 146, 3, 2, 2, 2, 487, 488, 5, 63, 32, 2, 488, 489, 5, 57, 29,
	2, 489, 490, 5, 65, 33, 2, 490, 491, 5, 57, 29, 2, 491, 492, 5, 79, 40,
	2, 492, 148, 3, 2, 2, 2, 493, 494, 5, 67, 34, 2, 494, 495, 5, 81, 41, 2,
	495, 496, 5, 63, 32, 2, 496, 497, 5, 63, 32, 2, 497, 150, 3, 2, 2, 2, 498,
	499, 5, 69, 35, 2, 499, 500, 5, 51, 26, 2, 500, 501, 5, 51, 26, 2, 501,
	502, 5, 77, 39, 2, 502, 503, 5, 49, 25, 2, 503, 504, 5, 79, 40, 2, 504,
	152, 3, 2, 2, 2, 505, 506, 5, 69, 35, 2, 506, 507, 5, 71, 36, 2, 507, 508,
	5, 79, 40, 2, 508, 509, 5, 57, 29, 2, 509, 510, 5, 69, 35, 2, 510, 511,
	5, 67, 34, 2, 511, 512, 5, 77, 39, 2, 512, 154, 3, 2, 2, 2, 513, 514, 5,
	69, 35, 2, 514, 515, 5, 75, 38, 2, 515, 516, 5, 47, 24, 2, 516, 517, 5,
	49, 25, 2, 517, 518, 5, 75, 38, 2, 518, 156, 3, 2, 2, 2, 519, 520, 5, 77,
	39, 2, 520, 521, 5, 49, 25, 2, 521, 522, 5, 63, 32, 2, 522, 523, 5, 49,
	25, 2, 523, 524, 5, 45, 23, 2, 524, 525, 5, 79, 40, 2, 525, 158, 3, 2,
	2, 2, 526, 527, 5, 77, 39, 2, 527, 528, 5, 49, 25, 2, 528, 529, 5, 79,
	40, 2, 529, 160, 3, 2, 2, 2, 530, 531, 5, 77, 39, 2, 531, 532, 5, 55, 28,
	2, 532, 533, 5, 69, 35, 2, 533, 534, 5, 85, 43, 2, 534, 162, 3, 2, 2, 2,
	535, 536, 5, 77, 39, 2, 536, 537, 5, 89, 45, 2, 537, 538, 5, 67, 34, 2,
	538, 539, 5, 45, 23, 2, 539, 164, 3, 2, 2, 2, 540, 541, 5, 81, 41, 2, 541,
	542, 5, 77, 39, 2, 542, 543, 5, 49, 25, 2, 543, 166, 3, 2, 2, 2, 544, 545,
	5, 81, 41, 2, 545, 546, 5, 67, 34, 2, 546, 547, 5, 57, 29, 2, 547, 548,
	5, 69, 35, 2, 548, 549, 5, 67, 34, 2, 549, 168, 3, 2, 2, 2, 550, 551, 5,
	81, 41, 2, 551, 552, 5, 71, 36, 2, 552, 553, 5, 47, 24, 2, 553, 554, 5,
	41, 21, 2, 554, 555, 5, 79, 40, 2, 555, 556, 5, 49, 25, 2, 556, 170, 3,
	2, 2, 2, 557, 558, 5, 85, 43, 2, 558, 559, 5, 55, 28, 2, 559, 560, 5, 49,
	25, 2, 560, 561, 5, 75, 38, 2, 561, 562, 5, 49, 25, 2, 562, 172, 3, 2,
	2, 2, 563, 564, 5, 83, 42, 2, 564, 565, 5, 41, 21, 2, 565, 566, 5, 63,
	32, 2, 566, 567, 5, 81, 41, 2, 567, 568, 5, 49, 25, 2, 568, 174, 3, 2,
	2, 2, 569, 570, 5, 83, 42, 2, 570, 571, 5, 41, 21, 2, 571, 572, 5, 63,
	32, 2, 572, 573, 5, 81, 41, 2, 573, 574, 5, 49, 25, 2, 574, 575, 5, 77,
	39, 2, 575, 176, 3, 2, 2, 2, 576, 577, 5, 79, 40, 2, 577, 578, 5, 75, 38,
	2, 578, 579, 5, 81, 41, 2, 579, 580, 5, 49, 25, 2, 580, 178, 3, 2, 2, 2,
	581, 582, 5, 51, 26, 2, 582, 583, 5, 41, 21, 2, 583, 584, 5, 63, 32, 2,
	584, 585, 5, 77, 39, 2, 585, 586, 5, 49, 25, 2, 586, 180, 3, 2, 2, 2, 587,
	588, 9, 28, 2, 2, 588, 182, 3, 2, 2, 2, 589, 593, 9, 29, 2, 2, 590, 592,
	9, 30, 2, 2, 591, 590, 3, 2, 2, 2, 592, 595, 3, 2, 2, 2, 593, 591, 3, 2,
	2, 2, 593, 594, 3, 2, 2, 2, 594, 184, 3, 2, 2, 2, 595, 593, 3, 2, 2, 2,
	596, 598, 4, 50, 59, 2, 597, 596, 3, 2, 2, 2, 598, 599, 3, 2, 2, 2, 599,
	597, 3, 2, 2, 2, 599, 600, 3, 2, 2, 2, 600, 186, 3, 2, 2, 2, 601, 603,
	4, 50, 59, 2, 602, 601, 3, 2, 2, 2, 603, 604, 3, 2, 2, 2, 604, 602, 3,
	2, 2, 2, 604, 605, 3, 2, 2, 2, 605, 606, 3, 2, 2, 2, 606, 610, 7, 48, 2,
	2, 607, 609, 4, 50, 59, 2, 608, 607, 3, 2, 2, 2, 609, 612, 3, 2, 2, 2,
	610, 608, 3, 2, 2, 2, 610, 611, 3, 2, 2, 2, 611, 614, 3, 2, 2, 2, 612,
	610, 3, 2, 2, 2, 613, 615, 5, 191, 96, 2, 614, 613, 3, 2, 2, 2, 614, 615,
	3, 2, 2, 2, 615, 632, 3, 2, 2, 2, 616, 618, 7, 48, 2, 2, 617, 619, 4, 50,
	59, 2, 618, 617, 3, 2, 2, 2, 619, 620, 3, 2, 2, 2, 620, 618, 3, 2, 2, 2,
	620, 621, 3, 2, 2, 2, 621, 623, 3, 2, 2, 2, 622, 624, 5, 191, 96, 2, 623,
	622, 3, 2, 2, 2, 623, 624, 3, 2, 2, 2, 624, 632, 3, 2, 2, 2, 625, 627,
	4, 50, 59, 2, 626, 625, 3, 2, 2, 2, 627, 628, 3, 2, 2, 2, 628, 626, 3,
	2, 2, 2, 628, 629, 3, 2, 2, 2, 629, 630, 3, 2, 2, 2, 630, 632, 5, 191,
	96, 2, 631, 602, 3, 2, 2, 2, 631, 616, 3, 2, 2, 2, 631, 626, 3, 2, 2, 2,
	632, 188, 3, 2, 2, 2, 633, 639, 7, 41, 2, 2, 634, 638, 10, 31, 2, 2, 635,
	636, 7, 41, 2, 2, 636, 638, 7, 41, 2, 2, 637, 634, 3, 2, 2, 2, 637, 635,
	3, 2, 2, 2, 638, 641, 3, 2, 2, 2, 639, 637, 3, 2, 2, 2, 639, 640, 3, 2,
	2, 2, 640, 642, 3, 2, 2, 2, 641, 639, 3, 2, 2, 2, 642, 643, 7, 41, 2, 2,
	643, 190, 3, 2, 2, 2, 644, 646, 9, 6, 2, 2, 645, 647, 9, 32, 2, 2, 646,
	645, 3, 2, 2, 2, 646, 647, 3, 2, 2, 2, 647, 649, 3, 2, 2, 2, 648, 650,
	4, 50, 59, 2, 649, 648, 3, 2, 2, 2, 650, 651, 3, 2, 2, 2, 651, 649, 3,
	2, 2, 2, 651, 652, 3, 2, 2, 2, 652, 192, 3, 2, 2, 2, 653, 654, 9, 33, 2,
	2, 654, 194, 3, 2, 2, 2, 18, 2, 234, 240, 593, 599, 604, 610, 614, 620,
	623, 628, 631, 637, 639, 646, 651, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'{'", "'}'", "'['", "']'", "':'", "'*'", "'='", "'=='",
	"'<'", "'<='", "'>'", "'>='", "'!='", "", "", "','", "';'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "ASTERISK", "SINGLE_EQ", "DOUBLE_EQ", "OP_LT",
	"LE", "GT", "GE", "NOTEQ", "AND", "OR", "COMMA", "SEMICOLON", "ALL", "ANCESTOR",
	"AS", "ASC", "ASYNC", "BY", "CREATE", "COLLECTION", "CURRENT_DATE", "CURRENT_TIME",
	"CURRENT_TIMESTAMP", "DESC", "DELETE", "DISTINCT", "DROP", "EACH", "EXCEPT",
	"FLATTEN", "FROM", "GROUP", "HAVING", "IN", "COLLECTION_INDEX", "INSERT",
	"INTERSECT", "INTO", "IS", "LIMIT", "NIL", "OFFSET", "OPTIONS", "ORDER",
	"SELECT", "SET", "SHOW", "SYNC", "USE", "UNION", "UPDATE", "WHERE", "VALUE",
	"VALUES", "TRUE", "FALSE", "WS", "IDENTIFIER", "NUMBER", "FLOAT", "STRING",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "ASTERISK", "SINGLE_EQ",
	"DOUBLE_EQ", "OP_LT", "LE", "GT", "GE", "NOTEQ", "AND", "OR", "COMMA",
	"SEMICOLON", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
	"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "ALL",
	"ANCESTOR", "AS", "ASC", "ASYNC", "BY", "CREATE", "COLLECTION", "CURRENT_DATE",
	"CURRENT_TIME", "CURRENT_TIMESTAMP", "DESC", "DELETE", "DISTINCT", "DROP",
	"EACH", "EXCEPT", "FLATTEN", "FROM", "GROUP", "HAVING", "IN", "COLLECTION_INDEX",
	"INSERT", "INTERSECT", "INTO", "IS", "LIMIT", "NIL", "OFFSET", "OPTIONS",
	"ORDER", "SELECT", "SET", "SHOW", "SYNC", "USE", "UNION", "UPDATE", "WHERE",
	"VALUE", "VALUES", "TRUE", "FALSE", "WS", "IDENTIFIER", "NUMBER", "FLOAT",
	"STRING", "EXPONENT", "HEX_DIGIT",
}

type SQLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewSQLLexer(input antlr.CharStream) *SQLLexer {

	l := new(SQLLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "SQL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SQLLexer tokens.
const (
	SQLLexerT__0              = 1
	SQLLexerT__1              = 2
	SQLLexerT__2              = 3
	SQLLexerT__3              = 4
	SQLLexerT__4              = 5
	SQLLexerT__5              = 6
	SQLLexerT__6              = 7
	SQLLexerASTERISK          = 8
	SQLLexerSINGLE_EQ         = 9
	SQLLexerDOUBLE_EQ         = 10
	SQLLexerOP_LT             = 11
	SQLLexerLE                = 12
	SQLLexerGT                = 13
	SQLLexerGE                = 14
	SQLLexerNOTEQ             = 15
	SQLLexerAND               = 16
	SQLLexerOR                = 17
	SQLLexerCOMMA             = 18
	SQLLexerSEMICOLON         = 19
	SQLLexerALL               = 20
	SQLLexerANCESTOR          = 21
	SQLLexerAS                = 22
	SQLLexerASC               = 23
	SQLLexerASYNC             = 24
	SQLLexerBY                = 25
	SQLLexerCREATE            = 26
	SQLLexerCOLLECTION        = 27
	SQLLexerCURRENT_DATE      = 28
	SQLLexerCURRENT_TIME      = 29
	SQLLexerCURRENT_TIMESTAMP = 30
	SQLLexerDESC              = 31
	SQLLexerDELETE            = 32
	SQLLexerDISTINCT          = 33
	SQLLexerDROP              = 34
	SQLLexerEACH              = 35
	SQLLexerEXCEPT            = 36
	SQLLexerFLATTEN           = 37
	SQLLexerFROM              = 38
	SQLLexerGROUP             = 39
	SQLLexerHAVING            = 40
	SQLLexerIN                = 41
	SQLLexerCOLLECTION_INDEX  = 42
	SQLLexerINSERT            = 43
	SQLLexerINTERSECT         = 44
	SQLLexerINTO              = 45
	SQLLexerIS                = 46
	SQLLexerLIMIT             = 47
	SQLLexerNIL               = 48
	SQLLexerOFFSET            = 49
	SQLLexerOPTIONS           = 50
	SQLLexerORDER             = 51
	SQLLexerSELECT            = 52
	SQLLexerSET               = 53
	SQLLexerSHOW              = 54
	SQLLexerSYNC              = 55
	SQLLexerUSE               = 56
	SQLLexerUNION             = 57
	SQLLexerUPDATE            = 58
	SQLLexerWHERE             = 59
	SQLLexerVALUE             = 60
	SQLLexerVALUES            = 61
	SQLLexerTRUE              = 62
	SQLLexerFALSE             = 63
	SQLLexerWS                = 64
	SQLLexerIDENTIFIER        = 65
	SQLLexerNUMBER            = 66
	SQLLexerFLOAT             = 67
	SQLLexerSTRING            = 68
)
