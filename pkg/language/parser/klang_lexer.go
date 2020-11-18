// Generated from Klang.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 72, 611,
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
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3,
	26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3,
	31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3,
	39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40,
	3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3,
	42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3,
	46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48,
	3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3,
	49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50,
	3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3,
	51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52,
	3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3,
	53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54,
	3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3,
	55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55,
	3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3,
	56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57,
	3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 58, 3,
	58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58,
	3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 60, 3,
	60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61,
	3, 61, 3, 61, 3, 61, 3, 62, 3, 62, 3, 62, 3, 62, 3, 62, 3, 62, 3, 62, 3,
	63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 64, 3, 64, 7, 64,
	500, 10, 64, 12, 64, 14, 64, 503, 11, 64, 3, 65, 5, 65, 506, 10, 65, 3,
	65, 3, 65, 3, 65, 6, 65, 511, 10, 65, 13, 65, 14, 65, 512, 5, 65, 515,
	10, 65, 3, 65, 5, 65, 518, 10, 65, 3, 66, 3, 66, 5, 66, 522, 10, 66, 3,
	66, 3, 66, 3, 67, 3, 67, 3, 67, 7, 67, 529, 10, 67, 12, 67, 14, 67, 532,
	11, 67, 5, 67, 534, 10, 67, 3, 68, 3, 68, 7, 68, 538, 10, 68, 12, 68, 14,
	68, 541, 11, 68, 3, 69, 3, 69, 7, 69, 545, 10, 69, 12, 69, 14, 69, 548,
	11, 69, 3, 69, 3, 69, 3, 70, 3, 70, 3, 70, 7, 70, 555, 10, 70, 12, 70,
	14, 70, 558, 11, 70, 3, 70, 3, 70, 3, 70, 3, 70, 7, 70, 564, 10, 70, 12,
	70, 14, 70, 567, 11, 70, 3, 70, 5, 70, 570, 10, 70, 3, 71, 3, 71, 3, 71,
	5, 71, 575, 10, 71, 3, 72, 3, 72, 3, 73, 3, 73, 3, 73, 5, 73, 582, 10,
	73, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74, 3, 75, 3, 75, 3, 76, 3, 76,
	3, 77, 3, 77, 7, 77, 596, 10, 77, 12, 77, 14, 77, 599, 11, 77, 3, 77, 3,
	77, 3, 78, 6, 78, 604, 10, 78, 13, 78, 14, 78, 605, 3, 78, 3, 78, 3, 79,
	3, 79, 2, 2, 80, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55,
	29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73,
	38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91,
	47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55,
	109, 56, 111, 57, 113, 58, 115, 59, 117, 60, 119, 61, 121, 62, 123, 63,
	125, 64, 127, 65, 129, 66, 131, 2, 133, 2, 135, 67, 137, 68, 139, 69, 141,
	2, 143, 2, 145, 2, 147, 2, 149, 2, 151, 2, 153, 70, 155, 71, 157, 72, 3,
	2, 18, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99,
	124, 3, 2, 50, 59, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45, 47, 47, 3, 2,
	51, 59, 4, 2, 67, 92, 99, 124, 7, 2, 47, 59, 67, 92, 94, 94, 97, 97, 99,
	124, 3, 2, 98, 98, 10, 2, 41, 41, 49, 49, 94, 94, 100, 100, 104, 104, 112,
	112, 116, 116, 118, 118, 5, 2, 2, 33, 41, 41, 94, 94, 10, 2, 36, 36, 49,
	49, 94, 94, 100, 100, 104, 104, 112, 112, 116, 116, 118, 118, 5, 2, 50,
	59, 67, 72, 99, 104, 5, 2, 2, 33, 36, 36, 94, 94, 4, 2, 12, 12, 15, 15,
	5, 2, 11, 12, 15, 15, 34, 34, 2, 621, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29,
	3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2,
	37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2,
	2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2,
	2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2,
	2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3,
	2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75,
	3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2,
	83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2,
	2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2,
	2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3,
	2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2,
	113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2,
	2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127,
	3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2,
	2, 139, 3, 2, 2, 2, 2, 153, 3, 2, 2, 2, 2, 155, 3, 2, 2, 2, 2, 157, 3,
	2, 2, 2, 3, 159, 3, 2, 2, 2, 5, 161, 3, 2, 2, 2, 7, 163, 3, 2, 2, 2, 9,
	165, 3, 2, 2, 2, 11, 170, 3, 2, 2, 2, 13, 173, 3, 2, 2, 2, 15, 176, 3,
	2, 2, 2, 17, 179, 3, 2, 2, 2, 19, 182, 3, 2, 2, 2, 21, 184, 3, 2, 2, 2,
	23, 186, 3, 2, 2, 2, 25, 189, 3, 2, 2, 2, 27, 192, 3, 2, 2, 2, 29, 194,
	3, 2, 2, 2, 31, 196, 3, 2, 2, 2, 33, 198, 3, 2, 2, 2, 35, 200, 3, 2, 2,
	2, 37, 202, 3, 2, 2, 2, 39, 204, 3, 2, 2, 2, 41, 206, 3, 2, 2, 2, 43, 208,
	3, 2, 2, 2, 45, 210, 3, 2, 2, 2, 47, 212, 3, 2, 2, 2, 49, 214, 3, 2, 2,
	2, 51, 216, 3, 2, 2, 2, 53, 218, 3, 2, 2, 2, 55, 220, 3, 2, 2, 2, 57, 225,
	3, 2, 2, 2, 59, 231, 3, 2, 2, 2, 61, 235, 3, 2, 2, 2, 63, 238, 3, 2, 2,
	2, 65, 243, 3, 2, 2, 2, 67, 249, 3, 2, 2, 2, 69, 253, 3, 2, 2, 2, 71, 261,
	3, 2, 2, 2, 73, 267, 3, 2, 2, 2, 75, 273, 3, 2, 2, 2, 77, 277, 3, 2, 2,
	2, 79, 285, 3, 2, 2, 2, 81, 292, 3, 2, 2, 2, 83, 295, 3, 2, 2, 2, 85, 302,
	3, 2, 2, 2, 87, 305, 3, 2, 2, 2, 89, 308, 3, 2, 2, 2, 91, 318, 3, 2, 2,
	2, 93, 323, 3, 2, 2, 2, 95, 328, 3, 2, 2, 2, 97, 339, 3, 2, 2, 2, 99, 348,
	3, 2, 2, 2, 101, 359, 3, 2, 2, 2, 103, 370, 3, 2, 2, 2, 105, 379, 3, 2,
	2, 2, 107, 390, 3, 2, 2, 2, 109, 403, 3, 2, 2, 2, 111, 418, 3, 2, 2, 2,
	113, 431, 3, 2, 2, 2, 115, 446, 3, 2, 2, 2, 117, 458, 3, 2, 2, 2, 119,
	467, 3, 2, 2, 2, 121, 473, 3, 2, 2, 2, 123, 482, 3, 2, 2, 2, 125, 489,
	3, 2, 2, 2, 127, 497, 3, 2, 2, 2, 129, 505, 3, 2, 2, 2, 131, 519, 3, 2,
	2, 2, 133, 533, 3, 2, 2, 2, 135, 535, 3, 2, 2, 2, 137, 542, 3, 2, 2, 2,
	139, 569, 3, 2, 2, 2, 141, 571, 3, 2, 2, 2, 143, 576, 3, 2, 2, 2, 145,
	578, 3, 2, 2, 2, 147, 583, 3, 2, 2, 2, 149, 589, 3, 2, 2, 2, 151, 591,
	3, 2, 2, 2, 153, 593, 3, 2, 2, 2, 155, 603, 3, 2, 2, 2, 157, 609, 3, 2,
	2, 2, 159, 160, 7, 60, 2, 2, 160, 4, 3, 2, 2, 2, 161, 162, 7, 93, 2, 2,
	162, 6, 3, 2, 2, 2, 163, 164, 7, 95, 2, 2, 164, 8, 3, 2, 2, 2, 165, 166,
	7, 112, 2, 2, 166, 167, 7, 119, 2, 2, 167, 168, 7, 110, 2, 2, 168, 169,
	7, 110, 2, 2, 169, 10, 3, 2, 2, 2, 170, 171, 7, 126, 2, 2, 171, 172, 7,
	126, 2, 2, 172, 12, 3, 2, 2, 2, 173, 174, 7, 40, 2, 2, 174, 175, 7, 40,
	2, 2, 175, 14, 3, 2, 2, 2, 176, 177, 7, 63, 2, 2, 177, 178, 7, 63, 2, 2,
	178, 16, 3, 2, 2, 2, 179, 180, 7, 35, 2, 2, 180, 181, 7, 63, 2, 2, 181,
	18, 3, 2, 2, 2, 182, 183, 7, 64, 2, 2, 183, 20, 3, 2, 2, 2, 184, 185, 7,
	62, 2, 2, 185, 22, 3, 2, 2, 2, 186, 187, 7, 64, 2, 2, 187, 188, 7, 63,
	2, 2, 188, 24, 3, 2, 2, 2, 189, 190, 7, 62, 2, 2, 190, 191, 7, 63, 2, 2,
	191, 26, 3, 2, 2, 2, 192, 193, 7, 45, 2, 2, 193, 28, 3, 2, 2, 2, 194, 195,
	7, 47, 2, 2, 195, 30, 3, 2, 2, 2, 196, 197, 7, 44, 2, 2, 197, 32, 3, 2,
	2, 2, 198, 199, 7, 49, 2, 2, 199, 34, 3, 2, 2, 2, 200, 201, 7, 39, 2, 2,
	201, 36, 3, 2, 2, 2, 202, 203, 7, 96, 2, 2, 203, 38, 3, 2, 2, 2, 204, 205,
	7, 35, 2, 2, 205, 40, 3, 2, 2, 2, 206, 207, 7, 61, 2, 2, 207, 42, 3, 2,
	2, 2, 208, 209, 7, 63, 2, 2, 209, 44, 3, 2, 2, 2, 210, 211, 7, 42, 2, 2,
	211, 46, 3, 2, 2, 2, 212, 213, 7, 43, 2, 2, 213, 48, 3, 2, 2, 2, 214, 215,
	7, 125, 2, 2, 215, 50, 3, 2, 2, 2, 216, 217, 7, 127, 2, 2, 217, 52, 3,
	2, 2, 2, 218, 219, 7, 46, 2, 2, 219, 54, 3, 2, 2, 2, 220, 221, 7, 118,
	2, 2, 221, 222, 7, 116, 2, 2, 222, 223, 7, 119, 2, 2, 223, 224, 7, 103,
	2, 2, 224, 56, 3, 2, 2, 2, 225, 226, 7, 104, 2, 2, 226, 227, 7, 99, 2,
	2, 227, 228, 7, 110, 2, 2, 228, 229, 7, 117, 2, 2, 229, 230, 7, 103, 2,
	2, 230, 58, 3, 2, 2, 2, 231, 232, 7, 112, 2, 2, 232, 233, 7, 107, 2, 2,
	233, 234, 7, 110, 2, 2, 234, 60, 3, 2, 2, 2, 235, 236, 7, 107, 2, 2, 236,
	237, 7, 104, 2, 2, 237, 62, 3, 2, 2, 2, 238, 239, 7, 103, 2, 2, 239, 240,
	7, 110, 2, 2, 240, 241, 7, 117, 2, 2, 241, 242, 7, 103, 2, 2, 242, 64,
	3, 2, 2, 2, 243, 244, 7, 121, 2, 2, 244, 245, 7, 106, 2, 2, 245, 246, 7,
	107, 2, 2, 246, 247, 7, 110, 2, 2, 247, 248, 7, 103, 2, 2, 248, 66, 3,
	2, 2, 2, 249, 250, 7, 110, 2, 2, 250, 251, 7, 113, 2, 2, 251, 252, 7, 105,
	2, 2, 252, 68, 3, 2, 2, 2, 253, 254, 7, 109, 2, 2, 254, 255, 7, 119, 2,
	2, 255, 256, 7, 100, 2, 2, 256, 257, 7, 103, 2, 2, 257, 258, 7, 101, 2,
	2, 258, 259, 7, 118, 2, 2, 259, 260, 7, 110, 2, 2, 260, 70, 3, 2, 2, 2,
	261, 262, 7, 99, 2, 2, 262, 263, 7, 114, 2, 2, 263, 264, 7, 114, 2, 2,
	264, 265, 7, 110, 2, 2, 265, 266, 7, 123, 2, 2, 266, 72, 3, 2, 2, 2, 267,
	268, 7, 114, 2, 2, 268, 269, 7, 99, 2, 2, 269, 270, 7, 118, 2, 2, 270,
	271, 7, 101, 2, 2, 271, 272, 7, 106, 2, 2, 272, 74, 3, 2, 2, 2, 273, 274,
	7, 105, 2, 2, 274, 275, 7, 103, 2, 2, 275, 276, 7, 118, 2, 2, 276, 76,
	3, 2, 2, 2, 277, 278, 7, 116, 2, 2, 278, 279, 7, 103, 2, 2, 279, 280, 7,
	114, 2, 2, 280, 281, 7, 110, 2, 2, 281, 282, 7, 99, 2, 2, 282, 283, 7,
	101, 2, 2, 283, 284, 7, 103, 2, 2, 284, 78, 3, 2, 2, 2, 285, 286, 7, 102,
	2, 2, 286, 287, 7, 103, 2, 2, 287, 288, 7, 110, 2, 2, 288, 289, 7, 103,
	2, 2, 289, 290, 7, 118, 2, 2, 290, 291, 7, 103, 2, 2, 291, 80, 3, 2, 2,
	2, 292, 293, 7, 47, 2, 2, 293, 294, 7, 112, 2, 2, 294, 82, 3, 2, 2, 2,
	295, 296, 7, 47, 2, 2, 296, 297, 7, 47, 2, 2, 297, 298, 7, 118, 2, 2, 298,
	299, 7, 123, 2, 2, 299, 300, 7, 114, 2, 2, 300, 301, 7, 103, 2, 2, 301,
	84, 3, 2, 2, 2, 302, 303, 7, 47, 2, 2, 303, 304, 7, 114, 2, 2, 304, 86,
	3, 2, 2, 2, 305, 306, 7, 47, 2, 2, 306, 307, 7, 119, 2, 2, 307, 88, 3,
	2, 2, 2, 308, 309, 7, 47, 2, 2, 309, 310, 7, 108, 2, 2, 310, 311, 7, 117,
	2, 2, 311, 312, 7, 113, 2, 2, 312, 313, 7, 112, 2, 2, 313, 314, 7, 114,
	2, 2, 314, 315, 7, 99, 2, 2, 315, 316, 7, 118, 2, 2, 316, 317, 7, 106,
	2, 2, 317, 90, 3, 2, 2, 2, 318, 319, 7, 110, 2, 2, 319, 320, 7, 113, 2,
	2, 320, 321, 7, 99, 2, 2, 321, 322, 7, 102, 2, 2, 322, 92, 3, 2, 2, 2,
	323, 324, 7, 103, 2, 2, 324, 325, 7, 122, 2, 2, 325, 326, 7, 107, 2, 2,
	326, 327, 7, 118, 2, 2, 327, 94, 3, 2, 2, 2, 328, 329, 7, 108, 2, 2, 329,
	330, 7, 117, 2, 2, 330, 331, 7, 113, 2, 2, 331, 332, 7, 112, 2, 2, 332,
	333, 7, 85, 2, 2, 333, 334, 7, 103, 2, 2, 334, 335, 7, 110, 2, 2, 335,
	336, 7, 103, 2, 2, 336, 337, 7, 101, 2, 2, 337, 338, 7, 118, 2, 2, 338,
	96, 3, 2, 2, 2, 339, 340, 7, 108, 2, 2, 340, 341, 7, 117, 2, 2, 341, 342,
	7, 113, 2, 2, 342, 343, 7, 112, 2, 2, 343, 344, 7, 71, 2, 2, 344, 345,
	7, 102, 2, 2, 345, 346, 7, 107, 2, 2, 346, 347, 7, 118, 2, 2, 347, 98,
	3, 2, 2, 2, 348, 349, 7, 108, 2, 2, 349, 350, 7, 117, 2, 2, 350, 351, 7,
	113, 2, 2, 351, 352, 7, 112, 2, 2, 352, 353, 7, 70, 2, 2, 353, 354, 7,
	103, 2, 2, 354, 355, 7, 110, 2, 2, 355, 356, 7, 103, 2, 2, 356, 357, 7,
	118, 2, 2, 357, 358, 7, 103, 2, 2, 358, 100, 3, 2, 2, 2, 359, 360, 7, 123,
	2, 2, 360, 361, 7, 99, 2, 2, 361, 362, 7, 111, 2, 2, 362, 363, 7, 110,
	2, 2, 363, 364, 7, 85, 2, 2, 364, 365, 7, 103, 2, 2, 365, 366, 7, 110,
	2, 2, 366, 367, 7, 103, 2, 2, 367, 368, 7, 101, 2, 2, 368, 369, 7, 118,
	2, 2, 369, 102, 3, 2, 2, 2, 370, 371, 7, 123, 2, 2, 371, 372, 7, 99, 2,
	2, 372, 373, 7, 111, 2, 2, 373, 374, 7, 110, 2, 2, 374, 375, 7, 71, 2,
	2, 375, 376, 7, 102, 2, 2, 376, 377, 7, 107, 2, 2, 377, 378, 7, 118, 2,
	2, 378, 104, 3, 2, 2, 2, 379, 380, 7, 123, 2, 2, 380, 381, 7, 99, 2, 2,
	381, 382, 7, 111, 2, 2, 382, 383, 7, 110, 2, 2, 383, 384, 7, 70, 2, 2,
	384, 385, 7, 103, 2, 2, 385, 386, 7, 110, 2, 2, 386, 387, 7, 103, 2, 2,
	387, 388, 7, 118, 2, 2, 388, 389, 7, 103, 2, 2, 389, 106, 3, 2, 2, 2, 390,
	391, 7, 109, 2, 2, 391, 392, 7, 119, 2, 2, 392, 393, 7, 100, 2, 2, 393,
	394, 7, 103, 2, 2, 394, 395, 7, 76, 2, 2, 395, 396, 7, 117, 2, 2, 396,
	397, 7, 113, 2, 2, 397, 398, 7, 112, 2, 2, 398, 399, 7, 71, 2, 2, 399,
	400, 7, 102, 2, 2, 400, 401, 7, 107, 2, 2, 401, 402, 7, 118, 2, 2, 402,
	108, 3, 2, 2, 2, 403, 404, 7, 109, 2, 2, 404, 405, 7, 119, 2, 2, 405, 406,
	7, 100, 2, 2, 406, 407, 7, 103, 2, 2, 407, 408, 7, 76, 2, 2, 408, 409,
	7, 117, 2, 2, 409, 410, 7, 113, 2, 2, 410, 411, 7, 112, 2, 2, 411, 412,
	7, 70, 2, 2, 412, 413, 7, 103, 2, 2, 413, 414, 7, 110, 2, 2, 414, 415,
	7, 103, 2, 2, 415, 416, 7, 118, 2, 2, 416, 417, 7, 103, 2, 2, 417, 110,
	3, 2, 2, 2, 418, 419, 7, 109, 2, 2, 419, 420, 7, 119, 2, 2, 420, 421, 7,
	100, 2, 2, 421, 422, 7, 103, 2, 2, 422, 423, 7, 91, 2, 2, 423, 424, 7,
	99, 2, 2, 424, 425, 7, 111, 2, 2, 425, 426, 7, 110, 2, 2, 426, 427, 7,
	71, 2, 2, 427, 428, 7, 102, 2, 2, 428, 429, 7, 107, 2, 2, 429, 430, 7,
	118, 2, 2, 430, 112, 3, 2, 2, 2, 431, 432, 7, 109, 2, 2, 432, 433, 7, 119,
	2, 2, 433, 434, 7, 100, 2, 2, 434, 435, 7, 103, 2, 2, 435, 436, 7, 91,
	2, 2, 436, 437, 7, 99, 2, 2, 437, 438, 7, 111, 2, 2, 438, 439, 7, 110,
	2, 2, 439, 440, 7, 70, 2, 2, 440, 441, 7, 103, 2, 2, 441, 442, 7, 110,
	2, 2, 442, 443, 7, 103, 2, 2, 443, 444, 7, 118, 2, 2, 444, 445, 7, 103,
	2, 2, 445, 114, 3, 2, 2, 2, 446, 447, 7, 117, 2, 2, 447, 448, 7, 106, 2,
	2, 448, 449, 7, 103, 2, 2, 449, 450, 7, 110, 2, 2, 450, 451, 7, 110, 2,
	2, 451, 452, 7, 85, 2, 2, 452, 453, 7, 101, 2, 2, 453, 454, 7, 116, 2,
	2, 454, 455, 7, 107, 2, 2, 455, 456, 7, 114, 2, 2, 456, 457, 7, 118, 2,
	2, 457, 116, 3, 2, 2, 2, 458, 459, 7, 102, 2, 2, 459, 460, 7, 113, 2, 2,
	460, 461, 7, 121, 2, 2, 461, 462, 7, 112, 2, 2, 462, 463, 7, 110, 2, 2,
	463, 464, 7, 113, 2, 2, 464, 465, 7, 99, 2, 2, 465, 466, 7, 102, 2, 2,
	466, 118, 3, 2, 2, 2, 467, 468, 7, 117, 2, 2, 468, 469, 7, 110, 2, 2, 469,
	470, 7, 103, 2, 2, 470, 471, 7, 103, 2, 2, 471, 472, 7, 114, 2, 2, 472,
	120, 3, 2, 2, 2, 473, 474, 7, 117, 2, 2, 474, 475, 7, 118, 2, 2, 475, 476,
	7, 103, 2, 2, 476, 477, 7, 114, 2, 2, 477, 478, 7, 75, 2, 2, 478, 479,
	7, 112, 2, 2, 479, 480, 7, 104, 2, 2, 480, 481, 7, 113, 2, 2, 481, 122,
	3, 2, 2, 2, 482, 483, 7, 104, 2, 2, 483, 484, 7, 107, 2, 2, 484, 485, 7,
	110, 2, 2, 485, 486, 7, 118, 2, 2, 486, 487, 7, 103, 2, 2, 487, 488, 7,
	116, 2, 2, 488, 124, 3, 2, 2, 2, 489, 490, 7, 114, 2, 2, 490, 491, 7, 99,
	2, 2, 491, 492, 7, 118, 2, 2, 492, 493, 7, 118, 2, 2, 493, 494, 7, 103,
	2, 2, 494, 495, 7, 116, 2, 2, 495, 496, 7, 112, 2, 2, 496, 126, 3, 2, 2,
	2, 497, 501, 9, 2, 2, 2, 498, 500, 9, 3, 2, 2, 499, 498, 3, 2, 2, 2, 500,
	503, 3, 2, 2, 2, 501, 499, 3, 2, 2, 2, 501, 502, 3, 2, 2, 2, 502, 128,
	3, 2, 2, 2, 503, 501, 3, 2, 2, 2, 504, 506, 7, 47, 2, 2, 505, 504, 3, 2,
	2, 2, 505, 506, 3, 2, 2, 2, 506, 507, 3, 2, 2, 2, 507, 514, 5, 133, 67,
	2, 508, 510, 7, 48, 2, 2, 509, 511, 9, 4, 2, 2, 510, 509, 3, 2, 2, 2, 511,
	512, 3, 2, 2, 2, 512, 510, 3, 2, 2, 2, 512, 513, 3, 2, 2, 2, 513, 515,
	3, 2, 2, 2, 514, 508, 3, 2, 2, 2, 514, 515, 3, 2, 2, 2, 515, 517, 3, 2,
	2, 2, 516, 518, 5, 131, 66, 2, 517, 516, 3, 2, 2, 2, 517, 518, 3, 2, 2,
	2, 518, 130, 3, 2, 2, 2, 519, 521, 9, 5, 2, 2, 520, 522, 9, 6, 2, 2, 521,
	520, 3, 2, 2, 2, 521, 522, 3, 2, 2, 2, 522, 523, 3, 2, 2, 2, 523, 524,
	5, 133, 67, 2, 524, 132, 3, 2, 2, 2, 525, 534, 7, 50, 2, 2, 526, 530, 9,
	7, 2, 2, 527, 529, 9, 4, 2, 2, 528, 527, 3, 2, 2, 2, 529, 532, 3, 2, 2,
	2, 530, 528, 3, 2, 2, 2, 530, 531, 3, 2, 2, 2, 531, 534, 3, 2, 2, 2, 532,
	530, 3, 2, 2, 2, 533, 525, 3, 2, 2, 2, 533, 526, 3, 2, 2, 2, 534, 134,
	3, 2, 2, 2, 535, 539, 9, 8, 2, 2, 536, 538, 9, 9, 2, 2, 537, 536, 3, 2,
	2, 2, 538, 541, 3, 2, 2, 2, 539, 537, 3, 2, 2, 2, 539, 540, 3, 2, 2, 2,
	540, 136, 3, 2, 2, 2, 541, 539, 3, 2, 2, 2, 542, 546, 7, 98, 2, 2, 543,
	545, 10, 10, 2, 2, 544, 543, 3, 2, 2, 2, 545, 548, 3, 2, 2, 2, 546, 544,
	3, 2, 2, 2, 546, 547, 3, 2, 2, 2, 547, 549, 3, 2, 2, 2, 548, 546, 3, 2,
	2, 2, 549, 550, 7, 98, 2, 2, 550, 138, 3, 2, 2, 2, 551, 556, 7, 36, 2,
	2, 552, 555, 5, 145, 73, 2, 553, 555, 5, 151, 76, 2, 554, 552, 3, 2, 2,
	2, 554, 553, 3, 2, 2, 2, 555, 558, 3, 2, 2, 2, 556, 554, 3, 2, 2, 2, 556,
	557, 3, 2, 2, 2, 557, 559, 3, 2, 2, 2, 558, 556, 3, 2, 2, 2, 559, 570,
	7, 36, 2, 2, 560, 565, 7, 41, 2, 2, 561, 564, 5, 141, 71, 2, 562, 564,
	5, 143, 72, 2, 563, 561, 3, 2, 2, 2, 563, 562, 3, 2, 2, 2, 564, 567, 3,
	2, 2, 2, 565, 563, 3, 2, 2, 2, 565, 566, 3, 2, 2, 2, 566, 568, 3, 2, 2,
	2, 567, 565, 3, 2, 2, 2, 568, 570, 7, 41, 2, 2, 569, 551, 3, 2, 2, 2, 569,
	560, 3, 2, 2, 2, 570, 140, 3, 2, 2, 2, 571, 574, 7, 94, 2, 2, 572, 575,
	9, 11, 2, 2, 573, 575, 5, 147, 74, 2, 574, 572, 3, 2, 2, 2, 574, 573, 3,
	2, 2, 2, 575, 142, 3, 2, 2, 2, 576, 577, 10, 12, 2, 2, 577, 144, 3, 2,
	2, 2, 578, 581, 7, 94, 2, 2, 579, 582, 9, 13, 2, 2, 580, 582, 5, 147, 74,
	2, 581, 579, 3, 2, 2, 2, 581, 580, 3, 2, 2, 2, 582, 146, 3, 2, 2, 2, 583,
	584, 7, 119, 2, 2, 584, 585, 5, 149, 75, 2, 585, 586, 5, 149, 75, 2, 586,
	587, 5, 149, 75, 2, 587, 588, 5, 149, 75, 2, 588, 148, 3, 2, 2, 2, 589,
	590, 9, 14, 2, 2, 590, 150, 3, 2, 2, 2, 591, 592, 10, 15, 2, 2, 592, 152,
	3, 2, 2, 2, 593, 597, 7, 37, 2, 2, 594, 596, 10, 16, 2, 2, 595, 594, 3,
	2, 2, 2, 596, 599, 3, 2, 2, 2, 597, 595, 3, 2, 2, 2, 597, 598, 3, 2, 2,
	2, 598, 600, 3, 2, 2, 2, 599, 597, 3, 2, 2, 2, 600, 601, 8, 77, 2, 2, 601,
	154, 3, 2, 2, 2, 602, 604, 9, 17, 2, 2, 603, 602, 3, 2, 2, 2, 604, 605,
	3, 2, 2, 2, 605, 603, 3, 2, 2, 2, 605, 606, 3, 2, 2, 2, 606, 607, 3, 2,
	2, 2, 607, 608, 8, 78, 2, 2, 608, 156, 3, 2, 2, 2, 609, 610, 11, 2, 2,
	2, 610, 158, 3, 2, 2, 2, 22, 2, 501, 505, 512, 514, 517, 521, 530, 533,
	539, 546, 554, 556, 563, 565, 569, 574, 581, 597, 605, 3, 8, 2, 2,
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
	"", "':'", "'['", "']'", "'null'", "'||'", "'&&'", "'=='", "'!='", "'>'",
	"'<'", "'>='", "'<='", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'!'",
	"';'", "'='", "'('", "')'", "'{'", "'}'", "','", "'true'", "'false'", "'nil'",
	"'if'", "'else'", "'while'", "'log'", "'kubectl'", "'apply'", "'patch'",
	"'get'", "'replace'", "'delete'", "'-n'", "'--type'", "'-p'", "'-u'", "'-jsonpath'",
	"'load'", "'exit'", "'jsonSelect'", "'jsonEdit'", "'jsonDelete'", "'yamlSelect'",
	"'yamlEdit'", "'yamlDelete'", "'kubeJsonEdit'", "'kubeJsonDelete'", "'kubeYamlEdit'",
	"'kubeYamlDelete'", "'shellScript'", "'download'", "'sleep'", "'stepInfo'",
	"'filter'", "'pattern'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ",
	"PLUS", "MINUS", "MULT", "DIV", "MOD", "POW", "NOT", "SCOL", "ASSIGN",
	"OPAR", "CPAR", "OBRACE", "CBRACE", "COMMA", "TRUE", "FALSE", "NIL", "IF",
	"ELSE", "WHILE", "LOG", "KUBECTL", "APPLY", "PATCH", "GET", "REPLACE",
	"DELETE", "NAMESPACE", "PATCHTYPE", "PATCHLOAD", "UPDATELOAD", "JSONPATH",
	"LOAD", "EXIT", "JSONSELECT", "JSONEDIT", "JSONDELETE", "YAMLSELECT", "YAMLEDIT",
	"YAMLDELETE", "KUBEJSONEDIT", "KUBEJSONDELETE", "KUBEYAMLEDIT", "KUBEYAMLDELETE",
	"SHELLSCRIPT", "DOWNLOAD", "SLEEP", "STEPINFO", "FILTER", "PATTERN", "ID",
	"NUMBER", "PATH", "RAW_STRING_LIT", "STRING", "COMMENT", "SPACE", "OTHER",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ",
	"LTEQ", "PLUS", "MINUS", "MULT", "DIV", "MOD", "POW", "NOT", "SCOL", "ASSIGN",
	"OPAR", "CPAR", "OBRACE", "CBRACE", "COMMA", "TRUE", "FALSE", "NIL", "IF",
	"ELSE", "WHILE", "LOG", "KUBECTL", "APPLY", "PATCH", "GET", "REPLACE",
	"DELETE", "NAMESPACE", "PATCHTYPE", "PATCHLOAD", "UPDATELOAD", "JSONPATH",
	"LOAD", "EXIT", "JSONSELECT", "JSONEDIT", "JSONDELETE", "YAMLSELECT", "YAMLEDIT",
	"YAMLDELETE", "KUBEJSONEDIT", "KUBEJSONDELETE", "KUBEYAMLEDIT", "KUBEYAMLDELETE",
	"SHELLSCRIPT", "DOWNLOAD", "SLEEP", "STEPINFO", "FILTER", "PATTERN", "ID",
	"NUMBER", "EXP", "INT", "PATH", "RAW_STRING_LIT", "STRING", "ESCQUOTE",
	"SAFECODEPOINTQUOTE", "ESC", "UNICODE", "HEX", "SAFECODEPOINT", "COMMENT",
	"SPACE", "OTHER",
}

type KlangLexer struct {
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

func NewKlangLexer(input antlr.CharStream) *KlangLexer {

	l := new(KlangLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Klang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// KlangLexer tokens.
const (
	KlangLexerT__0           = 1
	KlangLexerT__1           = 2
	KlangLexerT__2           = 3
	KlangLexerT__3           = 4
	KlangLexerOR             = 5
	KlangLexerAND            = 6
	KlangLexerEQ             = 7
	KlangLexerNEQ            = 8
	KlangLexerGT             = 9
	KlangLexerLT             = 10
	KlangLexerGTEQ           = 11
	KlangLexerLTEQ           = 12
	KlangLexerPLUS           = 13
	KlangLexerMINUS          = 14
	KlangLexerMULT           = 15
	KlangLexerDIV            = 16
	KlangLexerMOD            = 17
	KlangLexerPOW            = 18
	KlangLexerNOT            = 19
	KlangLexerSCOL           = 20
	KlangLexerASSIGN         = 21
	KlangLexerOPAR           = 22
	KlangLexerCPAR           = 23
	KlangLexerOBRACE         = 24
	KlangLexerCBRACE         = 25
	KlangLexerCOMMA          = 26
	KlangLexerTRUE           = 27
	KlangLexerFALSE          = 28
	KlangLexerNIL            = 29
	KlangLexerIF             = 30
	KlangLexerELSE           = 31
	KlangLexerWHILE          = 32
	KlangLexerLOG            = 33
	KlangLexerKUBECTL        = 34
	KlangLexerAPPLY          = 35
	KlangLexerPATCH          = 36
	KlangLexerGET            = 37
	KlangLexerREPLACE        = 38
	KlangLexerDELETE         = 39
	KlangLexerNAMESPACE      = 40
	KlangLexerPATCHTYPE      = 41
	KlangLexerPATCHLOAD      = 42
	KlangLexerUPDATELOAD     = 43
	KlangLexerJSONPATH       = 44
	KlangLexerLOAD           = 45
	KlangLexerEXIT           = 46
	KlangLexerJSONSELECT     = 47
	KlangLexerJSONEDIT       = 48
	KlangLexerJSONDELETE     = 49
	KlangLexerYAMLSELECT     = 50
	KlangLexerYAMLEDIT       = 51
	KlangLexerYAMLDELETE     = 52
	KlangLexerKUBEJSONEDIT   = 53
	KlangLexerKUBEJSONDELETE = 54
	KlangLexerKUBEYAMLEDIT   = 55
	KlangLexerKUBEYAMLDELETE = 56
	KlangLexerSHELLSCRIPT    = 57
	KlangLexerDOWNLOAD       = 58
	KlangLexerSLEEP          = 59
	KlangLexerSTEPINFO       = 60
	KlangLexerFILTER         = 61
	KlangLexerPATTERN        = 62
	KlangLexerID             = 63
	KlangLexerNUMBER         = 64
	KlangLexerPATH           = 65
	KlangLexerRAW_STRING_LIT = 66
	KlangLexerSTRING         = 67
	KlangLexerCOMMENT        = 68
	KlangLexerSPACE          = 69
	KlangLexerOTHER          = 70
)