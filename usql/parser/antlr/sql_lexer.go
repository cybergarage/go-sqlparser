// Code generated from SQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 70, 661,
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
	4, 97, 9, 97, 4, 98, 9, 98, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17,
	237, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 243, 10, 18, 3, 19, 3,
	19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24,
	3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3,
	30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35,
	3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3,
	40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45,
	3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3,
	48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50,
	3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3,
	53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54,
	3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3,
	55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 56,
	3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3,
	56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57,
	3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 58, 3,
	58, 3, 58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59,
	3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 61, 3,
	61, 3, 61, 3, 61, 3, 61, 3, 62, 3, 62, 3, 62, 3, 62, 3, 62, 3, 63, 3, 63,
	3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 64, 3, 64, 3, 64, 3, 64, 3, 64, 3,
	64, 3, 64, 3, 64, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 66, 3, 66, 3, 66,
	3, 66, 3, 66, 3, 66, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3,
	68, 3, 68, 3, 68, 3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3, 70, 3, 70,
	3, 70, 3, 70, 3, 70, 3, 70, 3, 70, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3,
	71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 73,
	3, 73, 3, 73, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74, 3, 75, 3, 75, 3,
	75, 3, 75, 3, 75, 3, 76, 3, 76, 3, 76, 3, 76, 3, 76, 3, 76, 3, 76, 3, 77,
	3, 77, 3, 77, 3, 77, 3, 77, 3, 77, 3, 77, 3, 77, 3, 78, 3, 78, 3, 78, 3,
	78, 3, 78, 3, 78, 3, 79, 3, 79, 3, 79, 3, 79, 3, 79, 3, 79, 3, 79, 3, 80,
	3, 80, 3, 80, 3, 80, 3, 81, 3, 81, 3, 81, 3, 81, 3, 81, 3, 82, 3, 82, 3,
	82, 3, 82, 3, 82, 3, 83, 3, 83, 3, 83, 3, 83, 3, 84, 3, 84, 3, 84, 3, 84,
	3, 84, 3, 84, 3, 85, 3, 85, 3, 85, 3, 85, 3, 85, 3, 85, 3, 85, 3, 86, 3,
	86, 3, 86, 3, 86, 3, 86, 3, 86, 3, 87, 3, 87, 3, 87, 3, 87, 3, 87, 3, 87,
	3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 89, 3, 89, 3, 89, 3,
	89, 3, 89, 3, 90, 3, 90, 3, 90, 3, 90, 3, 90, 3, 90, 3, 91, 3, 91, 3, 92,
	3, 92, 7, 92, 594, 10, 92, 12, 92, 14, 92, 597, 11, 92, 3, 93, 6, 93, 600,
	10, 93, 13, 93, 14, 93, 601, 3, 94, 6, 94, 605, 10, 94, 13, 94, 14, 94,
	606, 3, 94, 3, 94, 7, 94, 611, 10, 94, 12, 94, 14, 94, 614, 11, 94, 3,
	94, 5, 94, 617, 10, 94, 3, 94, 3, 94, 6, 94, 621, 10, 94, 13, 94, 14, 94,
	622, 3, 94, 5, 94, 626, 10, 94, 3, 94, 6, 94, 629, 10, 94, 13, 94, 14,
	94, 630, 3, 94, 5, 94, 634, 10, 94, 3, 95, 3, 95, 3, 95, 7, 95, 639, 10,
	95, 12, 95, 14, 95, 642, 11, 95, 3, 95, 3, 95, 3, 96, 3, 96, 3, 96, 5,
	96, 649, 10, 96, 3, 97, 3, 97, 5, 97, 653, 10, 97, 3, 97, 6, 97, 656, 10,
	97, 13, 97, 14, 97, 657, 3, 98, 3, 98, 2, 2, 99, 3, 3, 5, 4, 7, 5, 9, 6,
	11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29,
	16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 2, 43, 2, 45, 2, 47, 2,
	49, 2, 51, 2, 53, 2, 55, 2, 57, 2, 59, 2, 61, 2, 63, 2, 65, 2, 67, 2, 69,
	2, 71, 2, 73, 2, 75, 2, 77, 2, 79, 2, 81, 2, 83, 2, 85, 2, 87, 2, 89, 2,
	91, 2, 93, 22, 95, 23, 97, 24, 99, 25, 101, 26, 103, 27, 105, 28, 107,
	29, 109, 30, 111, 31, 113, 32, 115, 33, 117, 34, 119, 35, 121, 36, 123,
	37, 125, 38, 127, 39, 129, 40, 131, 41, 133, 42, 135, 43, 137, 44, 139,
	45, 141, 46, 143, 47, 145, 48, 147, 49, 149, 50, 151, 51, 153, 52, 155,
	53, 157, 54, 159, 55, 161, 56, 163, 57, 165, 58, 167, 59, 169, 60, 171,
	61, 173, 62, 175, 63, 177, 64, 179, 65, 181, 66, 183, 67, 185, 68, 187,
	69, 189, 70, 191, 2, 193, 2, 195, 2, 3, 2, 35, 4, 2, 67, 67, 99, 99, 4,
	2, 68, 68, 100, 100, 4, 2, 69, 69, 101, 101, 4, 2, 70, 70, 102, 102, 4,
	2, 71, 71, 103, 103, 4, 2, 72, 72, 104, 104, 4, 2, 73, 73, 105, 105, 4,
	2, 74, 74, 106, 106, 4, 2, 75, 75, 107, 107, 4, 2, 76, 76, 108, 108, 4,
	2, 77, 77, 109, 109, 4, 2, 78, 78, 110, 110, 4, 2, 79, 79, 111, 111, 4,
	2, 80, 80, 112, 112, 4, 2, 81, 81, 113, 113, 4, 2, 82, 82, 114, 114, 4,
	2, 83, 83, 115, 115, 4, 2, 84, 84, 116, 116, 4, 2, 85, 85, 117, 117, 4,
	2, 86, 86, 118, 118, 4, 2, 87, 87, 119, 119, 4, 2, 88, 88, 120, 120, 4,
	2, 89, 89, 121, 121, 4, 2, 90, 90, 122, 122, 4, 2, 91, 91, 123, 123, 4,
	2, 92, 92, 124, 124, 5, 2, 11, 12, 15, 15, 34, 34, 6, 2, 49, 49, 67, 92,
	97, 97, 99, 124, 6, 2, 47, 59, 67, 92, 97, 97, 99, 124, 4, 2, 36, 36, 94,
	94, 4, 2, 41, 41, 94, 94, 4, 2, 45, 45, 47, 47, 5, 2, 50, 59, 67, 72, 99,
	104, 2, 648, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2,
	2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2,
	2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2,
	2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3,
	2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2,
	107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2,
	2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121,
	3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127, 3, 2, 2, 2,
	2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2, 133, 3, 2, 2, 2, 2, 135, 3,
	2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3, 2, 2, 2, 2, 141, 3, 2, 2, 2, 2,
	143, 3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147, 3, 2, 2, 2, 2, 149, 3, 2,
	2, 2, 2, 151, 3, 2, 2, 2, 2, 153, 3, 2, 2, 2, 2, 155, 3, 2, 2, 2, 2, 157,
	3, 2, 2, 2, 2, 159, 3, 2, 2, 2, 2, 161, 3, 2, 2, 2, 2, 163, 3, 2, 2, 2,
	2, 165, 3, 2, 2, 2, 2, 167, 3, 2, 2, 2, 2, 169, 3, 2, 2, 2, 2, 171, 3,
	2, 2, 2, 2, 173, 3, 2, 2, 2, 2, 175, 3, 2, 2, 2, 2, 177, 3, 2, 2, 2, 2,
	179, 3, 2, 2, 2, 2, 181, 3, 2, 2, 2, 2, 183, 3, 2, 2, 2, 2, 185, 3, 2,
	2, 2, 2, 187, 3, 2, 2, 2, 2, 189, 3, 2, 2, 2, 3, 197, 3, 2, 2, 2, 5, 199,
	3, 2, 2, 2, 7, 201, 3, 2, 2, 2, 9, 203, 3, 2, 2, 2, 11, 205, 3, 2, 2, 2,
	13, 207, 3, 2, 2, 2, 15, 209, 3, 2, 2, 2, 17, 211, 3, 2, 2, 2, 19, 213,
	3, 2, 2, 2, 21, 215, 3, 2, 2, 2, 23, 218, 3, 2, 2, 2, 25, 220, 3, 2, 2,
	2, 27, 223, 3, 2, 2, 2, 29, 225, 3, 2, 2, 2, 31, 228, 3, 2, 2, 2, 33, 236,
	3, 2, 2, 2, 35, 242, 3, 2, 2, 2, 37, 244, 3, 2, 2, 2, 39, 246, 3, 2, 2,
	2, 41, 248, 3, 2, 2, 2, 43, 250, 3, 2, 2, 2, 45, 252, 3, 2, 2, 2, 47, 254,
	3, 2, 2, 2, 49, 256, 3, 2, 2, 2, 51, 258, 3, 2, 2, 2, 53, 260, 3, 2, 2,
	2, 55, 262, 3, 2, 2, 2, 57, 264, 3, 2, 2, 2, 59, 266, 3, 2, 2, 2, 61, 268,
	3, 2, 2, 2, 63, 270, 3, 2, 2, 2, 65, 272, 3, 2, 2, 2, 67, 274, 3, 2, 2,
	2, 69, 276, 3, 2, 2, 2, 71, 278, 3, 2, 2, 2, 73, 280, 3, 2, 2, 2, 75, 282,
	3, 2, 2, 2, 77, 284, 3, 2, 2, 2, 79, 286, 3, 2, 2, 2, 81, 288, 3, 2, 2,
	2, 83, 290, 3, 2, 2, 2, 85, 292, 3, 2, 2, 2, 87, 294, 3, 2, 2, 2, 89, 296,
	3, 2, 2, 2, 91, 298, 3, 2, 2, 2, 93, 300, 3, 2, 2, 2, 95, 304, 3, 2, 2,
	2, 97, 313, 3, 2, 2, 2, 99, 316, 3, 2, 2, 2, 101, 320, 3, 2, 2, 2, 103,
	326, 3, 2, 2, 2, 105, 329, 3, 2, 2, 2, 107, 336, 3, 2, 2, 2, 109, 347,
	3, 2, 2, 2, 111, 360, 3, 2, 2, 2, 113, 373, 3, 2, 2, 2, 115, 391, 3, 2,
	2, 2, 117, 396, 3, 2, 2, 2, 119, 403, 3, 2, 2, 2, 121, 412, 3, 2, 2, 2,
	123, 417, 3, 2, 2, 2, 125, 422, 3, 2, 2, 2, 127, 429, 3, 2, 2, 2, 129,
	437, 3, 2, 2, 2, 131, 442, 3, 2, 2, 2, 133, 448, 3, 2, 2, 2, 135, 455,
	3, 2, 2, 2, 137, 458, 3, 2, 2, 2, 139, 464, 3, 2, 2, 2, 141, 471, 3, 2,
	2, 2, 143, 481, 3, 2, 2, 2, 145, 486, 3, 2, 2, 2, 147, 489, 3, 2, 2, 2,
	149, 495, 3, 2, 2, 2, 151, 500, 3, 2, 2, 2, 153, 507, 3, 2, 2, 2, 155,
	515, 3, 2, 2, 2, 157, 521, 3, 2, 2, 2, 159, 528, 3, 2, 2, 2, 161, 532,
	3, 2, 2, 2, 163, 537, 3, 2, 2, 2, 165, 542, 3, 2, 2, 2, 167, 546, 3, 2,
	2, 2, 169, 552, 3, 2, 2, 2, 171, 559, 3, 2, 2, 2, 173, 565, 3, 2, 2, 2,
	175, 571, 3, 2, 2, 2, 177, 578, 3, 2, 2, 2, 179, 583, 3, 2, 2, 2, 181,
	589, 3, 2, 2, 2, 183, 591, 3, 2, 2, 2, 185, 599, 3, 2, 2, 2, 187, 633,
	3, 2, 2, 2, 189, 635, 3, 2, 2, 2, 191, 645, 3, 2, 2, 2, 193, 650, 3, 2,
	2, 2, 195, 659, 3, 2, 2, 2, 197, 198, 7, 42, 2, 2, 198, 4, 3, 2, 2, 2,
	199, 200, 7, 43, 2, 2, 200, 6, 3, 2, 2, 2, 201, 202, 7, 125, 2, 2, 202,
	8, 3, 2, 2, 2, 203, 204, 7, 127, 2, 2, 204, 10, 3, 2, 2, 2, 205, 206, 7,
	93, 2, 2, 206, 12, 3, 2, 2, 2, 207, 208, 7, 95, 2, 2, 208, 14, 3, 2, 2,
	2, 209, 210, 7, 60, 2, 2, 210, 16, 3, 2, 2, 2, 211, 212, 7, 44, 2, 2, 212,
	18, 3, 2, 2, 2, 213, 214, 7, 63, 2, 2, 214, 20, 3, 2, 2, 2, 215, 216, 7,
	63, 2, 2, 216, 217, 7, 63, 2, 2, 217, 22, 3, 2, 2, 2, 218, 219, 7, 62,
	2, 2, 219, 24, 3, 2, 2, 2, 220, 221, 7, 62, 2, 2, 221, 222, 7, 63, 2, 2,
	222, 26, 3, 2, 2, 2, 223, 224, 7, 64, 2, 2, 224, 28, 3, 2, 2, 2, 225, 226,
	7, 64, 2, 2, 226, 227, 7, 63, 2, 2, 227, 30, 3, 2, 2, 2, 228, 229, 7, 35,
	2, 2, 229, 230, 7, 63, 2, 2, 230, 32, 3, 2, 2, 2, 231, 232, 5, 41, 21,
	2, 232, 233, 5, 67, 34, 2, 233, 234, 5, 47, 24, 2, 234, 237, 3, 2, 2, 2,
	235, 237, 7, 40, 2, 2, 236, 231, 3, 2, 2, 2, 236, 235, 3, 2, 2, 2, 237,
	34, 3, 2, 2, 2, 238, 239, 5, 69, 35, 2, 239, 240, 5, 75, 38, 2, 240, 243,
	3, 2, 2, 2, 241, 243, 7, 126, 2, 2, 242, 238, 3, 2, 2, 2, 242, 241, 3,
	2, 2, 2, 243, 36, 3, 2, 2, 2, 244, 245, 7, 46, 2, 2, 245, 38, 3, 2, 2,
	2, 246, 247, 7, 61, 2, 2, 247, 40, 3, 2, 2, 2, 248, 249, 9, 2, 2, 2, 249,
	42, 3, 2, 2, 2, 250, 251, 9, 3, 2, 2, 251, 44, 3, 2, 2, 2, 252, 253, 9,
	4, 2, 2, 253, 46, 3, 2, 2, 2, 254, 255, 9, 5, 2, 2, 255, 48, 3, 2, 2, 2,
	256, 257, 9, 6, 2, 2, 257, 50, 3, 2, 2, 2, 258, 259, 9, 7, 2, 2, 259, 52,
	3, 2, 2, 2, 260, 261, 9, 8, 2, 2, 261, 54, 3, 2, 2, 2, 262, 263, 9, 9,
	2, 2, 263, 56, 3, 2, 2, 2, 264, 265, 9, 10, 2, 2, 265, 58, 3, 2, 2, 2,
	266, 267, 9, 11, 2, 2, 267, 60, 3, 2, 2, 2, 268, 269, 9, 12, 2, 2, 269,
	62, 3, 2, 2, 2, 270, 271, 9, 13, 2, 2, 271, 64, 3, 2, 2, 2, 272, 273, 9,
	14, 2, 2, 273, 66, 3, 2, 2, 2, 274, 275, 9, 15, 2, 2, 275, 68, 3, 2, 2,
	2, 276, 277, 9, 16, 2, 2, 277, 70, 3, 2, 2, 2, 278, 279, 9, 17, 2, 2, 279,
	72, 3, 2, 2, 2, 280, 281, 9, 18, 2, 2, 281, 74, 3, 2, 2, 2, 282, 283, 9,
	19, 2, 2, 283, 76, 3, 2, 2, 2, 284, 285, 9, 20, 2, 2, 285, 78, 3, 2, 2,
	2, 286, 287, 9, 21, 2, 2, 287, 80, 3, 2, 2, 2, 288, 289, 9, 22, 2, 2, 289,
	82, 3, 2, 2, 2, 290, 291, 9, 23, 2, 2, 291, 84, 3, 2, 2, 2, 292, 293, 9,
	24, 2, 2, 293, 86, 3, 2, 2, 2, 294, 295, 9, 25, 2, 2, 295, 88, 3, 2, 2,
	2, 296, 297, 9, 26, 2, 2, 297, 90, 3, 2, 2, 2, 298, 299, 9, 27, 2, 2, 299,
	92, 3, 2, 2, 2, 300, 301, 5, 41, 21, 2, 301, 302, 5, 63, 32, 2, 302, 303,
	5, 63, 32, 2, 303, 94, 3, 2, 2, 2, 304, 305, 5, 41, 21, 2, 305, 306, 5,
	67, 34, 2, 306, 307, 5, 45, 23, 2, 307, 308, 5, 49, 25, 2, 308, 309, 5,
	77, 39, 2, 309, 310, 5, 79, 40, 2, 310, 311, 5, 69, 35, 2, 311, 312, 5,
	75, 38, 2, 312, 96, 3, 2, 2, 2, 313, 314, 5, 41, 21, 2, 314, 315, 5, 77,
	39, 2, 315, 98, 3, 2, 2, 2, 316, 317, 5, 41, 21, 2, 317, 318, 5, 77, 39,
	2, 318, 319, 5, 45, 23, 2, 319, 100, 3, 2, 2, 2, 320, 321, 5, 41, 21, 2,
	321, 322, 5, 77, 39, 2, 322, 323, 5, 89, 45, 2, 323, 324, 5, 67, 34, 2,
	324, 325, 5, 45, 23, 2, 325, 102, 3, 2, 2, 2, 326, 327, 5, 43, 22, 2, 327,
	328, 5, 89, 45, 2, 328, 104, 3, 2, 2, 2, 329, 330, 5, 45, 23, 2, 330, 331,
	5, 75, 38, 2, 331, 332, 5, 49, 25, 2, 332, 333, 5, 41, 21, 2, 333, 334,
	5, 79, 40, 2, 334, 335, 5, 49, 25, 2, 335, 106, 3, 2, 2, 2, 336, 337, 5,
	45, 23, 2, 337, 338, 5, 69, 35, 2, 338, 339, 5, 63, 32, 2, 339, 340, 5,
	63, 32, 2, 340, 341, 5, 49, 25, 2, 341, 342, 5, 45, 23, 2, 342, 343, 5,
	79, 40, 2, 343, 344, 5, 57, 29, 2, 344, 345, 5, 69, 35, 2, 345, 346, 5,
	67, 34, 2, 346, 108, 3, 2, 2, 2, 347, 348, 5, 45, 23, 2, 348, 349, 5, 81,
	41, 2, 349, 350, 5, 75, 38, 2, 350, 351, 5, 75, 38, 2, 351, 352, 5, 49,
	25, 2, 352, 353, 5, 67, 34, 2, 353, 354, 5, 79, 40, 2, 354, 355, 7, 97,
	2, 2, 355, 356, 5, 47, 24, 2, 356, 357, 5, 41, 21, 2, 357, 358, 5, 79,
	40, 2, 358, 359, 5, 49, 25, 2, 359, 110, 3, 2, 2, 2, 360, 361, 5, 45, 23,
	2, 361, 362, 5, 81, 41, 2, 362, 363, 5, 75, 38, 2, 363, 364, 5, 75, 38,
	2, 364, 365, 5, 49, 25, 2, 365, 366, 5, 67, 34, 2, 366, 367, 5, 79, 40,
	2, 367, 368, 7, 97, 2, 2, 368, 369, 5, 79, 40, 2, 369, 370, 5, 57, 29,
	2, 370, 371, 5, 65, 33, 2, 371, 372, 5, 49, 25, 2, 372, 112, 3, 2, 2, 2,
	373, 374, 5, 45, 23, 2, 374, 375, 5, 81, 41, 2, 375, 376, 5, 75, 38, 2,
	376, 377, 5, 75, 38, 2, 377, 378, 5, 49, 25, 2, 378, 379, 5, 67, 34, 2,
	379, 380, 5, 79, 40, 2, 380, 381, 7, 97, 2, 2, 381, 382, 5, 79, 40, 2,
	382, 383, 5, 57, 29, 2, 383, 384, 5, 65, 33, 2, 384, 385, 5, 49, 25, 2,
	385, 386, 5, 77, 39, 2, 386, 387, 5, 79, 40, 2, 387, 388, 5, 41, 21, 2,
	388, 389, 5, 65, 33, 2, 389, 390, 5, 71, 36, 2, 390, 114, 3, 2, 2, 2, 391,
	392, 5, 47, 24, 2, 392, 393, 5, 49, 25, 2, 393, 394, 5, 77, 39, 2, 394,
	395, 5, 45, 23, 2, 395, 116, 3, 2, 2, 2, 396, 397, 5, 47, 24, 2, 397, 398,
	5, 49, 25, 2, 398, 399, 5, 63, 32, 2, 399, 400, 5, 49, 25, 2, 400, 401,
	5, 79, 40, 2, 401, 402, 5, 49, 25, 2, 402, 118, 3, 2, 2, 2, 403, 404, 5,
	47, 24, 2, 404, 405, 5, 57, 29, 2, 405, 406, 5, 77, 39, 2, 406, 407, 5,
	79, 40, 2, 407, 408, 5, 57, 29, 2, 408, 409, 5, 67, 34, 2, 409, 410, 5,
	45, 23, 2, 410, 411, 5, 79, 40, 2, 411, 120, 3, 2, 2, 2, 412, 413, 5, 47,
	24, 2, 413, 414, 5, 75, 38, 2, 414, 415, 5, 69, 35, 2, 415, 416, 5, 71,
	36, 2, 416, 122, 3, 2, 2, 2, 417, 418, 5, 49, 25, 2, 418, 419, 5, 41, 21,
	2, 419, 420, 5, 45, 23, 2, 420, 421, 5, 55, 28, 2, 421, 124, 3, 2, 2, 2,
	422, 423, 5, 49, 25, 2, 423, 424, 5, 87, 44, 2, 424, 425, 5, 45, 23, 2,
	425, 426, 5, 49, 25, 2, 426, 427, 5, 71, 36, 2, 427, 428, 5, 79, 40, 2,
	428, 126, 3, 2, 2, 2, 429, 430, 5, 51, 26, 2, 430, 431, 5, 63, 32, 2, 431,
	432, 5, 41, 21, 2, 432, 433, 5, 79, 40, 2, 433, 434, 5, 79, 40, 2, 434,
	435, 5, 49, 25, 2, 435, 436, 5, 67, 34, 2, 436, 128, 3, 2, 2, 2, 437, 438,
	5, 51, 26, 2, 438, 439, 5, 75, 38, 2, 439, 440, 5, 69, 35, 2, 440, 441,
	5, 65, 33, 2, 441, 130, 3, 2, 2, 2, 442, 443, 5, 53, 27, 2, 443, 444, 5,
	75, 38, 2, 444, 445, 5, 69, 35, 2, 445, 446, 5, 81, 41, 2, 446, 447, 5,
	71, 36, 2, 447, 132, 3, 2, 2, 2, 448, 449, 5, 55, 28, 2, 449, 450, 5, 41,
	21, 2, 450, 451, 5, 83, 42, 2, 451, 452, 5, 57, 29, 2, 452, 453, 5, 67,
	34, 2, 453, 454, 5, 53, 27, 2, 454, 134, 3, 2, 2, 2, 455, 456, 5, 57, 29,
	2, 456, 457, 5, 67, 34, 2, 457, 136, 3, 2, 2, 2, 458, 459, 5, 57, 29, 2,
	459, 460, 5, 67, 34, 2, 460, 461, 5, 47, 24, 2, 461, 462, 5, 49, 25, 2,
	462, 463, 5, 87, 44, 2, 463, 138, 3, 2, 2, 2, 464, 465, 5, 57, 29, 2, 465,
	466, 5, 67, 34, 2, 466, 467, 5, 77, 39, 2, 467, 468, 5, 49, 25, 2, 468,
	469, 5, 75, 38, 2, 469, 470, 5, 79, 40, 2, 470, 140, 3, 2, 2, 2, 471, 472,
	5, 57, 29, 2, 472, 473, 5, 67, 34, 2, 473, 474, 5, 79, 40, 2, 474, 475,
	5, 49, 25, 2, 475, 476, 5, 75, 38, 2, 476, 477, 5, 77, 39, 2, 477, 478,
	5, 49, 25, 2, 478, 479, 5, 45, 23, 2, 479, 480, 5, 79, 40, 2, 480, 142,
	3, 2, 2, 2, 481, 482, 5, 57, 29, 2, 482, 483, 5, 67, 34, 2, 483, 484, 5,
	79, 40, 2, 484, 485, 5, 69, 35, 2, 485, 144, 3, 2, 2, 2, 486, 487, 5, 57,
	29, 2, 487, 488, 5, 77, 39, 2, 488, 146, 3, 2, 2, 2, 489, 490, 5, 63, 32,
	2, 490, 491, 5, 57, 29, 2, 491, 492, 5, 65, 33, 2, 492, 493, 5, 57, 29,
	2, 493, 494, 5, 79, 40, 2, 494, 148, 3, 2, 2, 2, 495, 496, 5, 67, 34, 2,
	496, 497, 5, 81, 41, 2, 497, 498, 5, 63, 32, 2, 498, 499, 5, 63, 32, 2,
	499, 150, 3, 2, 2, 2, 500, 501, 5, 69, 35, 2, 501, 502, 5, 51, 26, 2, 502,
	503, 5, 51, 26, 2, 503, 504, 5, 77, 39, 2, 504, 505, 5, 49, 25, 2, 505,
	506, 5, 79, 40, 2, 506, 152, 3, 2, 2, 2, 507, 508, 5, 69, 35, 2, 508, 509,
	5, 71, 36, 2, 509, 510, 5, 79, 40, 2, 510, 511, 5, 57, 29, 2, 511, 512,
	5, 69, 35, 2, 512, 513, 5, 67, 34, 2, 513, 514, 5, 77, 39, 2, 514, 154,
	3, 2, 2, 2, 515, 516, 5, 69, 35, 2, 516, 517, 5, 75, 38, 2, 517, 518, 5,
	47, 24, 2, 518, 519, 5, 49, 25, 2, 519, 520, 5, 75, 38, 2, 520, 156, 3,
	2, 2, 2, 521, 522, 5, 77, 39, 2, 522, 523, 5, 49, 25, 2, 523, 524, 5, 63,
	32, 2, 524, 525, 5, 49, 25, 2, 525, 526, 5, 45, 23, 2, 526, 527, 5, 79,
	40, 2, 527, 158, 3, 2, 2, 2, 528, 529, 5, 77, 39, 2, 529, 530, 5, 49, 25,
	2, 530, 531, 5, 79, 40, 2, 531, 160, 3, 2, 2, 2, 532, 533, 5, 77, 39, 2,
	533, 534, 5, 55, 28, 2, 534, 535, 5, 69, 35, 2, 535, 536, 5, 85, 43, 2,
	536, 162, 3, 2, 2, 2, 537, 538, 5, 77, 39, 2, 538, 539, 5, 89, 45, 2, 539,
	540, 5, 67, 34, 2, 540, 541, 5, 45, 23, 2, 541, 164, 3, 2, 2, 2, 542, 543,
	5, 81, 41, 2, 543, 544, 5, 77, 39, 2, 544, 545, 5, 49, 25, 2, 545, 166,
	3, 2, 2, 2, 546, 547, 5, 81, 41, 2, 547, 548, 5, 67, 34, 2, 548, 549, 5,
	57, 29, 2, 549, 550, 5, 69, 35, 2, 550, 551, 5, 67, 34, 2, 551, 168, 3,
	2, 2, 2, 552, 553, 5, 81, 41, 2, 553, 554, 5, 71, 36, 2, 554, 555, 5, 47,
	24, 2, 555, 556, 5, 41, 21, 2, 556, 557, 5, 79, 40, 2, 557, 558, 5, 49,
	25, 2, 558, 170, 3, 2, 2, 2, 559, 560, 5, 85, 43, 2, 560, 561, 5, 55, 28,
	2, 561, 562, 5, 49, 25, 2, 562, 563, 5, 75, 38, 2, 563, 564, 5, 49, 25,
	2, 564, 172, 3, 2, 2, 2, 565, 566, 5, 83, 42, 2, 566, 567, 5, 41, 21, 2,
	567, 568, 5, 63, 32, 2, 568, 569, 5, 81, 41, 2, 569, 570, 5, 49, 25, 2,
	570, 174, 3, 2, 2, 2, 571, 572, 5, 83, 42, 2, 572, 573, 5, 41, 21, 2, 573,
	574, 5, 63, 32, 2, 574, 575, 5, 81, 41, 2, 575, 576, 5, 49, 25, 2, 576,
	577, 5, 77, 39, 2, 577, 176, 3, 2, 2, 2, 578, 579, 5, 79, 40, 2, 579, 580,
	5, 75, 38, 2, 580, 581, 5, 81, 41, 2, 581, 582, 5, 49, 25, 2, 582, 178,
	3, 2, 2, 2, 583, 584, 5, 51, 26, 2, 584, 585, 5, 41, 21, 2, 585, 586, 5,
	63, 32, 2, 586, 587, 5, 77, 39, 2, 587, 588, 5, 49, 25, 2, 588, 180, 3,
	2, 2, 2, 589, 590, 9, 28, 2, 2, 590, 182, 3, 2, 2, 2, 591, 595, 9, 29,
	2, 2, 592, 594, 9, 30, 2, 2, 593, 592, 3, 2, 2, 2, 594, 597, 3, 2, 2, 2,
	595, 593, 3, 2, 2, 2, 595, 596, 3, 2, 2, 2, 596, 184, 3, 2, 2, 2, 597,
	595, 3, 2, 2, 2, 598, 600, 4, 50, 59, 2, 599, 598, 3, 2, 2, 2, 600, 601,
	3, 2, 2, 2, 601, 599, 3, 2, 2, 2, 601, 602, 3, 2, 2, 2, 602, 186, 3, 2,
	2, 2, 603, 605, 4, 50, 59, 2, 604, 603, 3, 2, 2, 2, 605, 606, 3, 2, 2,
	2, 606, 604, 3, 2, 2, 2, 606, 607, 3, 2, 2, 2, 607, 608, 3, 2, 2, 2, 608,
	612, 7, 48, 2, 2, 609, 611, 4, 50, 59, 2, 610, 609, 3, 2, 2, 2, 611, 614,
	3, 2, 2, 2, 612, 610, 3, 2, 2, 2, 612, 613, 3, 2, 2, 2, 613, 616, 3, 2,
	2, 2, 614, 612, 3, 2, 2, 2, 615, 617, 5, 193, 97, 2, 616, 615, 3, 2, 2,
	2, 616, 617, 3, 2, 2, 2, 617, 634, 3, 2, 2, 2, 618, 620, 7, 48, 2, 2, 619,
	621, 4, 50, 59, 2, 620, 619, 3, 2, 2, 2, 621, 622, 3, 2, 2, 2, 622, 620,
	3, 2, 2, 2, 622, 623, 3, 2, 2, 2, 623, 625, 3, 2, 2, 2, 624, 626, 5, 193,
	97, 2, 625, 624, 3, 2, 2, 2, 625, 626, 3, 2, 2, 2, 626, 634, 3, 2, 2, 2,
	627, 629, 4, 50, 59, 2, 628, 627, 3, 2, 2, 2, 629, 630, 3, 2, 2, 2, 630,
	628, 3, 2, 2, 2, 630, 631, 3, 2, 2, 2, 631, 632, 3, 2, 2, 2, 632, 634,
	5, 193, 97, 2, 633, 604, 3, 2, 2, 2, 633, 618, 3, 2, 2, 2, 633, 628, 3,
	2, 2, 2, 634, 188, 3, 2, 2, 2, 635, 640, 7, 36, 2, 2, 636, 639, 5, 191,
	96, 2, 637, 639, 10, 31, 2, 2, 638, 636, 3, 2, 2, 2, 638, 637, 3, 2, 2,
	2, 639, 642, 3, 2, 2, 2, 640, 638, 3, 2, 2, 2, 640, 641, 3, 2, 2, 2, 641,
	643, 3, 2, 2, 2, 642, 640, 3, 2, 2, 2, 643, 644, 7, 36, 2, 2, 644, 190,
	3, 2, 2, 2, 645, 648, 7, 94, 2, 2, 646, 649, 3, 2, 2, 2, 647, 649, 9, 32,
	2, 2, 648, 646, 3, 2, 2, 2, 648, 647, 3, 2, 2, 2, 649, 192, 3, 2, 2, 2,
	650, 652, 9, 6, 2, 2, 651, 653, 9, 33, 2, 2, 652, 651, 3, 2, 2, 2, 652,
	653, 3, 2, 2, 2, 653, 655, 3, 2, 2, 2, 654, 656, 4, 50, 59, 2, 655, 654,
	3, 2, 2, 2, 656, 657, 3, 2, 2, 2, 657, 655, 3, 2, 2, 2, 657, 658, 3, 2,
	2, 2, 658, 194, 3, 2, 2, 2, 659, 660, 9, 34, 2, 2, 660, 196, 3, 2, 2, 2,
	19, 2, 236, 242, 595, 601, 606, 612, 616, 622, 625, 630, 633, 638, 640,
	648, 652, 657, 2,
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
	"VALUES", "TRUE", "FALSE", "WS", "ID", "NUMBER", "FLOAT", "STRING",
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
	"VALUE", "VALUES", "TRUE", "FALSE", "WS", "ID", "NUMBER", "FLOAT", "STRING",
	"EscapeSequence", "EXPONENT", "HEX_DIGIT",
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
	SQLLexerID                = 65
	SQLLexerNUMBER            = 66
	SQLLexerFLOAT             = 67
	SQLLexerSTRING            = 68
)
