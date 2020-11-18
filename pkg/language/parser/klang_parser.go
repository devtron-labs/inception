// Generated from Klang.g4 by ANTLR 4.7.

package parser // Klang
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 72, 522,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 87, 10, 3,
	12, 3, 14, 3, 90, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 108, 10, 4, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 120, 10, 5,
	3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 152, 10, 9, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 164, 10, 10, 3,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 5, 11, 178, 10, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 5, 12, 219, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 5, 13, 230, 10, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 5, 14, 271, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 286, 10,
	17, 12, 17, 14, 17, 289, 11, 17, 3, 17, 3, 17, 5, 17, 293, 10, 17, 3, 18,
	3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 303, 10, 19, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 6, 22, 320, 10, 22, 13, 22, 14, 22, 321, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 6, 22, 333,
	10, 22, 13, 22, 14, 22, 334, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 6, 22,
	342, 10, 22, 13, 22, 14, 22, 343, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 6,
	22, 351, 10, 22, 13, 22, 14, 22, 352, 5, 22, 355, 10, 22, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 5, 23, 362, 10, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 5, 25, 380, 10, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 5, 26, 389, 10, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 5, 27, 399, 10, 27, 3, 28, 3, 28, 5, 28, 403, 10, 28, 3, 29, 3,
	29, 5, 29, 407, 10, 29, 3, 30, 3, 30, 3, 31, 3, 31, 5, 31, 413, 10, 31,
	3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	5, 35, 436, 10, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 7, 35, 459, 10, 35, 12, 35, 14, 35, 462, 11, 35, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	5, 36, 475, 10, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 7, 38, 483,
	10, 38, 12, 38, 14, 38, 486, 11, 38, 3, 38, 3, 38, 3, 38, 3, 38, 5, 38,
	492, 10, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 7,
	40, 502, 10, 40, 12, 40, 14, 40, 505, 11, 40, 3, 40, 3, 40, 3, 40, 3, 40,
	5, 40, 511, 10, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 5,
	41, 520, 10, 41, 3, 41, 2, 3, 68, 42, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
	58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 2, 8, 4, 2, 65, 65, 68,
	69, 3, 2, 17, 19, 3, 2, 15, 16, 3, 2, 11, 14, 3, 2, 9, 10, 3, 2, 29, 30,
	2, 562, 2, 82, 3, 2, 2, 2, 4, 88, 3, 2, 2, 2, 6, 107, 3, 2, 2, 2, 8, 119,
	3, 2, 2, 2, 10, 121, 3, 2, 2, 2, 12, 124, 3, 2, 2, 2, 14, 134, 3, 2, 2,
	2, 16, 142, 3, 2, 2, 2, 18, 156, 3, 2, 2, 2, 20, 168, 3, 2, 2, 2, 22, 218,
	3, 2, 2, 2, 24, 220, 3, 2, 2, 2, 26, 270, 3, 2, 2, 2, 28, 272, 3, 2, 2,
	2, 30, 276, 3, 2, 2, 2, 32, 280, 3, 2, 2, 2, 34, 294, 3, 2, 2, 2, 36, 302,
	3, 2, 2, 2, 38, 304, 3, 2, 2, 2, 40, 308, 3, 2, 2, 2, 42, 354, 3, 2, 2,
	2, 44, 356, 3, 2, 2, 2, 46, 365, 3, 2, 2, 2, 48, 372, 3, 2, 2, 2, 50, 383,
	3, 2, 2, 2, 52, 398, 3, 2, 2, 2, 54, 402, 3, 2, 2, 2, 56, 406, 3, 2, 2,
	2, 58, 408, 3, 2, 2, 2, 60, 412, 3, 2, 2, 2, 62, 414, 3, 2, 2, 2, 64, 416,
	3, 2, 2, 2, 66, 420, 3, 2, 2, 2, 68, 435, 3, 2, 2, 2, 70, 474, 3, 2, 2,
	2, 72, 476, 3, 2, 2, 2, 74, 491, 3, 2, 2, 2, 76, 493, 3, 2, 2, 2, 78, 510,
	3, 2, 2, 2, 80, 519, 3, 2, 2, 2, 82, 83, 5, 4, 3, 2, 83, 84, 7, 2, 2, 3,
	84, 3, 3, 2, 2, 2, 85, 87, 5, 6, 4, 2, 86, 85, 3, 2, 2, 2, 87, 90, 3, 2,
	2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 5, 3, 2, 2, 2, 90, 88,
	3, 2, 2, 2, 91, 108, 5, 8, 5, 2, 92, 108, 5, 12, 7, 2, 93, 108, 5, 14,
	8, 2, 94, 108, 5, 16, 9, 2, 95, 108, 5, 18, 10, 2, 96, 108, 5, 22, 12,
	2, 97, 108, 5, 20, 11, 2, 98, 108, 5, 26, 14, 2, 99, 108, 5, 24, 13, 2,
	100, 108, 5, 32, 17, 2, 101, 108, 5, 38, 20, 2, 102, 108, 5, 28, 15, 2,
	103, 108, 5, 30, 16, 2, 104, 108, 5, 40, 21, 2, 105, 106, 7, 72, 2, 2,
	106, 108, 8, 4, 1, 2, 107, 91, 3, 2, 2, 2, 107, 92, 3, 2, 2, 2, 107, 93,
	3, 2, 2, 2, 107, 94, 3, 2, 2, 2, 107, 95, 3, 2, 2, 2, 107, 96, 3, 2, 2,
	2, 107, 97, 3, 2, 2, 2, 107, 98, 3, 2, 2, 2, 107, 99, 3, 2, 2, 2, 107,
	100, 3, 2, 2, 2, 107, 101, 3, 2, 2, 2, 107, 102, 3, 2, 2, 2, 107, 103,
	3, 2, 2, 2, 107, 104, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 108, 7, 3, 2, 2,
	2, 109, 110, 7, 65, 2, 2, 110, 111, 7, 23, 2, 2, 111, 112, 5, 68, 35, 2,
	112, 113, 7, 22, 2, 2, 113, 120, 3, 2, 2, 2, 114, 115, 7, 65, 2, 2, 115,
	116, 7, 23, 2, 2, 116, 117, 5, 50, 26, 2, 117, 118, 7, 22, 2, 2, 118, 120,
	3, 2, 2, 2, 119, 109, 3, 2, 2, 2, 119, 114, 3, 2, 2, 2, 120, 9, 3, 2, 2,
	2, 121, 122, 7, 59, 2, 2, 122, 123, 5, 58, 30, 2, 123, 11, 3, 2, 2, 2,
	124, 125, 7, 50, 2, 2, 125, 126, 7, 24, 2, 2, 126, 127, 7, 65, 2, 2, 127,
	128, 7, 28, 2, 2, 128, 129, 5, 58, 30, 2, 129, 130, 7, 28, 2, 2, 130, 131,
	5, 68, 35, 2, 131, 132, 7, 25, 2, 2, 132, 133, 7, 22, 2, 2, 133, 13, 3,
	2, 2, 2, 134, 135, 7, 51, 2, 2, 135, 136, 7, 24, 2, 2, 136, 137, 7, 65,
	2, 2, 137, 138, 7, 28, 2, 2, 138, 139, 5, 58, 30, 2, 139, 140, 7, 25, 2,
	2, 140, 141, 7, 22, 2, 2, 141, 15, 3, 2, 2, 2, 142, 143, 7, 53, 2, 2, 143,
	144, 7, 24, 2, 2, 144, 145, 7, 65, 2, 2, 145, 146, 7, 28, 2, 2, 146, 147,
	5, 58, 30, 2, 147, 148, 7, 28, 2, 2, 148, 151, 5, 68, 35, 2, 149, 150,
	7, 28, 2, 2, 150, 152, 7, 66, 2, 2, 151, 149, 3, 2, 2, 2, 151, 152, 3,
	2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 154, 7, 25, 2, 2, 154, 155, 7, 22,
	2, 2, 155, 17, 3, 2, 2, 2, 156, 157, 7, 54, 2, 2, 157, 158, 7, 24, 2, 2,
	158, 159, 7, 65, 2, 2, 159, 160, 7, 28, 2, 2, 160, 163, 5, 58, 30, 2, 161,
	162, 7, 28, 2, 2, 162, 164, 7, 66, 2, 2, 163, 161, 3, 2, 2, 2, 163, 164,
	3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 166, 7, 25, 2, 2, 166, 167, 7, 22,
	2, 2, 167, 19, 3, 2, 2, 2, 168, 169, 7, 55, 2, 2, 169, 170, 7, 24, 2, 2,
	170, 171, 7, 65, 2, 2, 171, 172, 7, 28, 2, 2, 172, 173, 5, 58, 30, 2, 173,
	174, 7, 28, 2, 2, 174, 177, 5, 68, 35, 2, 175, 176, 7, 28, 2, 2, 176, 178,
	5, 58, 30, 2, 177, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 179, 3,
	2, 2, 2, 179, 180, 7, 25, 2, 2, 180, 181, 7, 22, 2, 2, 181, 21, 3, 2, 2,
	2, 182, 183, 7, 56, 2, 2, 183, 184, 7, 24, 2, 2, 184, 185, 7, 65, 2, 2,
	185, 186, 7, 28, 2, 2, 186, 187, 5, 64, 33, 2, 187, 188, 7, 28, 2, 2, 188,
	189, 5, 66, 34, 2, 189, 190, 7, 25, 2, 2, 190, 191, 7, 22, 2, 2, 191, 219,
	3, 2, 2, 2, 192, 193, 7, 56, 2, 2, 193, 194, 7, 24, 2, 2, 194, 195, 7,
	65, 2, 2, 195, 196, 7, 28, 2, 2, 196, 197, 5, 66, 34, 2, 197, 198, 7, 28,
	2, 2, 198, 199, 5, 64, 33, 2, 199, 200, 7, 25, 2, 2, 200, 201, 7, 22, 2,
	2, 201, 219, 3, 2, 2, 2, 202, 203, 7, 56, 2, 2, 203, 204, 7, 24, 2, 2,
	204, 205, 7, 65, 2, 2, 205, 206, 7, 28, 2, 2, 206, 207, 5, 64, 33, 2, 207,
	208, 7, 25, 2, 2, 208, 209, 7, 22, 2, 2, 209, 219, 3, 2, 2, 2, 210, 211,
	7, 56, 2, 2, 211, 212, 7, 24, 2, 2, 212, 213, 7, 65, 2, 2, 213, 214, 7,
	28, 2, 2, 214, 215, 5, 66, 34, 2, 215, 216, 7, 25, 2, 2, 216, 217, 7, 22,
	2, 2, 217, 219, 3, 2, 2, 2, 218, 182, 3, 2, 2, 2, 218, 192, 3, 2, 2, 2,
	218, 202, 3, 2, 2, 2, 218, 210, 3, 2, 2, 2, 219, 23, 3, 2, 2, 2, 220, 221,
	7, 57, 2, 2, 221, 222, 7, 24, 2, 2, 222, 223, 7, 65, 2, 2, 223, 224, 7,
	28, 2, 2, 224, 225, 5, 58, 30, 2, 225, 226, 7, 28, 2, 2, 226, 229, 5, 68,
	35, 2, 227, 228, 7, 28, 2, 2, 228, 230, 5, 58, 30, 2, 229, 227, 3, 2, 2,
	2, 229, 230, 3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 232, 7, 25, 2, 2, 232,
	233, 7, 22, 2, 2, 233, 25, 3, 2, 2, 2, 234, 235, 7, 58, 2, 2, 235, 236,
	7, 24, 2, 2, 236, 237, 7, 65, 2, 2, 237, 238, 7, 28, 2, 2, 238, 239, 5,
	64, 33, 2, 239, 240, 7, 28, 2, 2, 240, 241, 5, 66, 34, 2, 241, 242, 7,
	25, 2, 2, 242, 243, 7, 22, 2, 2, 243, 271, 3, 2, 2, 2, 244, 245, 7, 58,
	2, 2, 245, 246, 7, 24, 2, 2, 246, 247, 7, 65, 2, 2, 247, 248, 7, 28, 2,
	2, 248, 249, 5, 66, 34, 2, 249, 250, 7, 28, 2, 2, 250, 251, 5, 64, 33,
	2, 251, 252, 7, 25, 2, 2, 252, 253, 7, 22, 2, 2, 253, 271, 3, 2, 2, 2,
	254, 255, 7, 58, 2, 2, 255, 256, 7, 24, 2, 2, 256, 257, 7, 65, 2, 2, 257,
	258, 7, 28, 2, 2, 258, 259, 5, 64, 33, 2, 259, 260, 7, 25, 2, 2, 260, 261,
	7, 22, 2, 2, 261, 271, 3, 2, 2, 2, 262, 263, 7, 58, 2, 2, 263, 264, 7,
	24, 2, 2, 264, 265, 7, 65, 2, 2, 265, 266, 7, 28, 2, 2, 266, 267, 5, 66,
	34, 2, 267, 268, 7, 25, 2, 2, 268, 269, 7, 22, 2, 2, 269, 271, 3, 2, 2,
	2, 270, 234, 3, 2, 2, 2, 270, 244, 3, 2, 2, 2, 270, 254, 3, 2, 2, 2, 270,
	262, 3, 2, 2, 2, 271, 27, 3, 2, 2, 2, 272, 273, 7, 61, 2, 2, 273, 274,
	7, 66, 2, 2, 274, 275, 7, 22, 2, 2, 275, 29, 3, 2, 2, 2, 276, 277, 7, 48,
	2, 2, 277, 278, 7, 66, 2, 2, 278, 279, 7, 22, 2, 2, 279, 31, 3, 2, 2, 2,
	280, 281, 7, 32, 2, 2, 281, 287, 5, 34, 18, 2, 282, 283, 7, 33, 2, 2, 283,
	284, 7, 32, 2, 2, 284, 286, 5, 34, 18, 2, 285, 282, 3, 2, 2, 2, 286, 289,
	3, 2, 2, 2, 287, 285, 3, 2, 2, 2, 287, 288, 3, 2, 2, 2, 288, 292, 3, 2,
	2, 2, 289, 287, 3, 2, 2, 2, 290, 291, 7, 33, 2, 2, 291, 293, 5, 36, 19,
	2, 292, 290, 3, 2, 2, 2, 292, 293, 3, 2, 2, 2, 293, 33, 3, 2, 2, 2, 294,
	295, 5, 68, 35, 2, 295, 296, 5, 36, 19, 2, 296, 35, 3, 2, 2, 2, 297, 298,
	7, 26, 2, 2, 298, 299, 5, 4, 3, 2, 299, 300, 7, 27, 2, 2, 300, 303, 3,
	2, 2, 2, 301, 303, 5, 6, 4, 2, 302, 297, 3, 2, 2, 2, 302, 301, 3, 2, 2,
	2, 303, 37, 3, 2, 2, 2, 304, 305, 7, 34, 2, 2, 305, 306, 5, 68, 35, 2,
	306, 307, 5, 36, 19, 2, 307, 39, 3, 2, 2, 2, 308, 309, 7, 35, 2, 2, 309,
	310, 5, 68, 35, 2, 310, 311, 7, 22, 2, 2, 311, 41, 3, 2, 2, 2, 312, 313,
	7, 36, 2, 2, 313, 319, 7, 37, 2, 2, 314, 315, 7, 42, 2, 2, 315, 320, 5,
	54, 28, 2, 316, 320, 5, 58, 30, 2, 317, 318, 7, 45, 2, 2, 318, 320, 5,
	62, 32, 2, 319, 314, 3, 2, 2, 2, 319, 316, 3, 2, 2, 2, 319, 317, 3, 2,
	2, 2, 320, 321, 3, 2, 2, 2, 321, 319, 3, 2, 2, 2, 321, 322, 3, 2, 2, 2,
	322, 355, 3, 2, 2, 2, 323, 324, 7, 36, 2, 2, 324, 332, 7, 38, 2, 2, 325,
	326, 7, 42, 2, 2, 326, 333, 5, 54, 28, 2, 327, 333, 5, 60, 31, 2, 328,
	329, 7, 43, 2, 2, 329, 333, 5, 56, 29, 2, 330, 331, 7, 44, 2, 2, 331, 333,
	5, 58, 30, 2, 332, 325, 3, 2, 2, 2, 332, 327, 3, 2, 2, 2, 332, 328, 3,
	2, 2, 2, 332, 330, 3, 2, 2, 2, 333, 334, 3, 2, 2, 2, 334, 332, 3, 2, 2,
	2, 334, 335, 3, 2, 2, 2, 335, 355, 3, 2, 2, 2, 336, 337, 7, 36, 2, 2, 337,
	341, 7, 39, 2, 2, 338, 339, 7, 42, 2, 2, 339, 342, 5, 54, 28, 2, 340, 342,
	5, 60, 31, 2, 341, 338, 3, 2, 2, 2, 341, 340, 3, 2, 2, 2, 342, 343, 3,
	2, 2, 2, 343, 341, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344, 355, 3, 2, 2,
	2, 345, 346, 7, 36, 2, 2, 346, 350, 7, 41, 2, 2, 347, 348, 7, 42, 2, 2,
	348, 351, 5, 54, 28, 2, 349, 351, 5, 60, 31, 2, 350, 347, 3, 2, 2, 2, 350,
	349, 3, 2, 2, 2, 351, 352, 3, 2, 2, 2, 352, 350, 3, 2, 2, 2, 352, 353,
	3, 2, 2, 2, 353, 355, 3, 2, 2, 2, 354, 312, 3, 2, 2, 2, 354, 323, 3, 2,
	2, 2, 354, 336, 3, 2, 2, 2, 354, 345, 3, 2, 2, 2, 355, 43, 3, 2, 2, 2,
	356, 357, 7, 60, 2, 2, 357, 358, 7, 24, 2, 2, 358, 361, 5, 58, 30, 2, 359,
	360, 7, 28, 2, 2, 360, 362, 5, 58, 30, 2, 361, 359, 3, 2, 2, 2, 361, 362,
	3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363, 364, 7, 25, 2, 2, 364, 45, 3, 2,
	2, 2, 365, 366, 7, 49, 2, 2, 366, 367, 7, 24, 2, 2, 367, 368, 7, 65, 2,
	2, 368, 369, 7, 28, 2, 2, 369, 370, 5, 58, 30, 2, 370, 371, 7, 25, 2, 2,
	371, 47, 3, 2, 2, 2, 372, 373, 7, 52, 2, 2, 373, 374, 7, 24, 2, 2, 374,
	375, 7, 65, 2, 2, 375, 376, 7, 28, 2, 2, 376, 379, 5, 58, 30, 2, 377, 378,
	7, 28, 2, 2, 378, 380, 7, 66, 2, 2, 379, 377, 3, 2, 2, 2, 379, 380, 3,
	2, 2, 2, 380, 381, 3, 2, 2, 2, 381, 382, 7, 25, 2, 2, 382, 49, 3, 2, 2,
	2, 383, 384, 7, 47, 2, 2, 384, 385, 7, 24, 2, 2, 385, 388, 5, 58, 30, 2,
	386, 387, 7, 28, 2, 2, 387, 389, 7, 69, 2, 2, 388, 386, 3, 2, 2, 2, 388,
	389, 3, 2, 2, 2, 389, 390, 3, 2, 2, 2, 390, 391, 7, 25, 2, 2, 391, 51,
	3, 2, 2, 2, 392, 393, 7, 62, 2, 2, 393, 394, 7, 69, 2, 2, 394, 399, 7,
	22, 2, 2, 395, 396, 7, 62, 2, 2, 396, 397, 7, 68, 2, 2, 397, 399, 7, 22,
	2, 2, 398, 392, 3, 2, 2, 2, 398, 395, 3, 2, 2, 2, 399, 53, 3, 2, 2, 2,
	400, 403, 5, 58, 30, 2, 401, 403, 7, 67, 2, 2, 402, 400, 3, 2, 2, 2, 402,
	401, 3, 2, 2, 2, 403, 55, 3, 2, 2, 2, 404, 407, 7, 67, 2, 2, 405, 407,
	5, 58, 30, 2, 406, 404, 3, 2, 2, 2, 406, 405, 3, 2, 2, 2, 407, 57, 3, 2,
	2, 2, 408, 409, 9, 2, 2, 2, 409, 59, 3, 2, 2, 2, 410, 413, 7, 67, 2, 2,
	411, 413, 5, 58, 30, 2, 412, 410, 3, 2, 2, 2, 412, 411, 3, 2, 2, 2, 413,
	61, 3, 2, 2, 2, 414, 415, 5, 58, 30, 2, 415, 63, 3, 2, 2, 2, 416, 417,
	7, 63, 2, 2, 417, 418, 7, 23, 2, 2, 418, 419, 5, 58, 30, 2, 419, 65, 3,
	2, 2, 2, 420, 421, 7, 64, 2, 2, 421, 422, 7, 23, 2, 2, 422, 423, 5, 58,
	30, 2, 423, 67, 3, 2, 2, 2, 424, 425, 8, 35, 1, 2, 425, 426, 7, 16, 2,
	2, 426, 436, 5, 68, 35, 16, 427, 428, 7, 21, 2, 2, 428, 436, 5, 68, 35,
	15, 429, 436, 5, 42, 22, 2, 430, 436, 5, 46, 24, 2, 431, 436, 5, 48, 25,
	2, 432, 436, 5, 10, 6, 2, 433, 436, 5, 44, 23, 2, 434, 436, 5, 70, 36,
	2, 435, 424, 3, 2, 2, 2, 435, 427, 3, 2, 2, 2, 435, 429, 3, 2, 2, 2, 435,
	430, 3, 2, 2, 2, 435, 431, 3, 2, 2, 2, 435, 432, 3, 2, 2, 2, 435, 433,
	3, 2, 2, 2, 435, 434, 3, 2, 2, 2, 436, 460, 3, 2, 2, 2, 437, 438, 12, 17,
	2, 2, 438, 439, 7, 20, 2, 2, 439, 459, 5, 68, 35, 17, 440, 441, 12, 14,
	2, 2, 441, 442, 9, 3, 2, 2, 442, 459, 5, 68, 35, 15, 443, 444, 12, 13,
	2, 2, 444, 445, 9, 4, 2, 2, 445, 459, 5, 68, 35, 14, 446, 447, 12, 12,
	2, 2, 447, 448, 9, 5, 2, 2, 448, 459, 5, 68, 35, 13, 449, 450, 12, 11,
	2, 2, 450, 451, 9, 6, 2, 2, 451, 459, 5, 68, 35, 12, 452, 453, 12, 10,
	2, 2, 453, 454, 7, 8, 2, 2, 454, 459, 5, 68, 35, 11, 455, 456, 12, 9, 2,
	2, 456, 457, 7, 7, 2, 2, 457, 459, 5, 68, 35, 10, 458, 437, 3, 2, 2, 2,
	458, 440, 3, 2, 2, 2, 458, 443, 3, 2, 2, 2, 458, 446, 3, 2, 2, 2, 458,
	449, 3, 2, 2, 2, 458, 452, 3, 2, 2, 2, 458, 455, 3, 2, 2, 2, 459, 462,
	3, 2, 2, 2, 460, 458, 3, 2, 2, 2, 460, 461, 3, 2, 2, 2, 461, 69, 3, 2,
	2, 2, 462, 460, 3, 2, 2, 2, 463, 464, 7, 24, 2, 2, 464, 465, 5, 68, 35,
	2, 465, 466, 7, 25, 2, 2, 466, 475, 3, 2, 2, 2, 467, 475, 7, 66, 2, 2,
	468, 475, 9, 7, 2, 2, 469, 475, 7, 68, 2, 2, 470, 475, 7, 65, 2, 2, 471,
	475, 7, 69, 2, 2, 472, 475, 5, 72, 37, 2, 473, 475, 7, 31, 2, 2, 474, 463,
	3, 2, 2, 2, 474, 467, 3, 2, 2, 2, 474, 468, 3, 2, 2, 2, 474, 469, 3, 2,
	2, 2, 474, 470, 3, 2, 2, 2, 474, 471, 3, 2, 2, 2, 474, 472, 3, 2, 2, 2,
	474, 473, 3, 2, 2, 2, 475, 71, 3, 2, 2, 2, 476, 477, 5, 80, 41, 2, 477,
	73, 3, 2, 2, 2, 478, 479, 7, 26, 2, 2, 479, 484, 5, 76, 39, 2, 480, 481,
	7, 28, 2, 2, 481, 483, 5, 76, 39, 2, 482, 480, 3, 2, 2, 2, 483, 486, 3,
	2, 2, 2, 484, 482, 3, 2, 2, 2, 484, 485, 3, 2, 2, 2, 485, 487, 3, 2, 2,
	2, 486, 484, 3, 2, 2, 2, 487, 488, 7, 27, 2, 2, 488, 492, 3, 2, 2, 2, 489,
	490, 7, 26, 2, 2, 490, 492, 7, 27, 2, 2, 491, 478, 3, 2, 2, 2, 491, 489,
	3, 2, 2, 2, 492, 75, 3, 2, 2, 2, 493, 494, 7, 69, 2, 2, 494, 495, 7, 3,
	2, 2, 495, 496, 5, 80, 41, 2, 496, 77, 3, 2, 2, 2, 497, 498, 7, 4, 2, 2,
	498, 503, 5, 80, 41, 2, 499, 500, 7, 28, 2, 2, 500, 502, 5, 80, 41, 2,
	501, 499, 3, 2, 2, 2, 502, 505, 3, 2, 2, 2, 503, 501, 3, 2, 2, 2, 503,
	504, 3, 2, 2, 2, 504, 506, 3, 2, 2, 2, 505, 503, 3, 2, 2, 2, 506, 507,
	7, 5, 2, 2, 507, 511, 3, 2, 2, 2, 508, 509, 7, 4, 2, 2, 509, 511, 7, 5,
	2, 2, 510, 497, 3, 2, 2, 2, 510, 508, 3, 2, 2, 2, 511, 79, 3, 2, 2, 2,
	512, 520, 7, 69, 2, 2, 513, 520, 7, 66, 2, 2, 514, 520, 5, 74, 38, 2, 515,
	520, 5, 78, 40, 2, 516, 520, 7, 29, 2, 2, 517, 520, 7, 30, 2, 2, 518, 520,
	7, 6, 2, 2, 519, 512, 3, 2, 2, 2, 519, 513, 3, 2, 2, 2, 519, 514, 3, 2,
	2, 2, 519, 515, 3, 2, 2, 2, 519, 516, 3, 2, 2, 2, 519, 517, 3, 2, 2, 2,
	519, 518, 3, 2, 2, 2, 520, 81, 3, 2, 2, 2, 39, 88, 107, 119, 151, 163,
	177, 218, 229, 270, 287, 292, 302, 319, 321, 332, 334, 341, 343, 350, 352,
	354, 361, 379, 388, 398, 402, 406, 412, 435, 458, 460, 474, 484, 491, 503,
	510, 519,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
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
var symbolicNames = []string{
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

var ruleNames = []string{
	"parse", "block", "stat", "assignment", "shell_script", "json_edit_fn",
	"json_delete_fn", "yaml_edit_fn", "yaml_delete_fn", "kube_json_edit_fn",
	"kube_json_delete_fn", "kube_yaml_edit_fn", "kube_yaml_delete_fn", "sleep_fn",
	"exit_fn", "if_stat", "condition_block", "stat_block", "while_stat", "log",
	"kubectl_command", "download_fn", "json_select_fn", "yaml_select_fn", "load_fn",
	"stepInfo", "ns", "patch_type", "string_or_id", "resource", "kubernetes_object_config",
	"filter", "pattern", "expr", "atom", "json", "obj", "pair", "arr", "value",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type KlangParser struct {
	*antlr.BaseParser
}

func NewKlangParser(input antlr.TokenStream) *KlangParser {
	this := new(KlangParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Klang.g4"

	return this
}

// KlangParser tokens.
const (
	KlangParserEOF            = antlr.TokenEOF
	KlangParserT__0           = 1
	KlangParserT__1           = 2
	KlangParserT__2           = 3
	KlangParserT__3           = 4
	KlangParserOR             = 5
	KlangParserAND            = 6
	KlangParserEQ             = 7
	KlangParserNEQ            = 8
	KlangParserGT             = 9
	KlangParserLT             = 10
	KlangParserGTEQ           = 11
	KlangParserLTEQ           = 12
	KlangParserPLUS           = 13
	KlangParserMINUS          = 14
	KlangParserMULT           = 15
	KlangParserDIV            = 16
	KlangParserMOD            = 17
	KlangParserPOW            = 18
	KlangParserNOT            = 19
	KlangParserSCOL           = 20
	KlangParserASSIGN         = 21
	KlangParserOPAR           = 22
	KlangParserCPAR           = 23
	KlangParserOBRACE         = 24
	KlangParserCBRACE         = 25
	KlangParserCOMMA          = 26
	KlangParserTRUE           = 27
	KlangParserFALSE          = 28
	KlangParserNIL            = 29
	KlangParserIF             = 30
	KlangParserELSE           = 31
	KlangParserWHILE          = 32
	KlangParserLOG            = 33
	KlangParserKUBECTL        = 34
	KlangParserAPPLY          = 35
	KlangParserPATCH          = 36
	KlangParserGET            = 37
	KlangParserREPLACE        = 38
	KlangParserDELETE         = 39
	KlangParserNAMESPACE      = 40
	KlangParserPATCHTYPE      = 41
	KlangParserPATCHLOAD      = 42
	KlangParserUPDATELOAD     = 43
	KlangParserJSONPATH       = 44
	KlangParserLOAD           = 45
	KlangParserEXIT           = 46
	KlangParserJSONSELECT     = 47
	KlangParserJSONEDIT       = 48
	KlangParserJSONDELETE     = 49
	KlangParserYAMLSELECT     = 50
	KlangParserYAMLEDIT       = 51
	KlangParserYAMLDELETE     = 52
	KlangParserKUBEJSONEDIT   = 53
	KlangParserKUBEJSONDELETE = 54
	KlangParserKUBEYAMLEDIT   = 55
	KlangParserKUBEYAMLDELETE = 56
	KlangParserSHELLSCRIPT    = 57
	KlangParserDOWNLOAD       = 58
	KlangParserSLEEP          = 59
	KlangParserSTEPINFO       = 60
	KlangParserFILTER         = 61
	KlangParserPATTERN        = 62
	KlangParserID             = 63
	KlangParserNUMBER         = 64
	KlangParserPATH           = 65
	KlangParserRAW_STRING_LIT = 66
	KlangParserSTRING         = 67
	KlangParserCOMMENT        = 68
	KlangParserSPACE          = 69
	KlangParserOTHER          = 70
)

// KlangParser rules.
const (
	KlangParserRULE_parse                    = 0
	KlangParserRULE_block                    = 1
	KlangParserRULE_stat                     = 2
	KlangParserRULE_assignment               = 3
	KlangParserRULE_shell_script             = 4
	KlangParserRULE_json_edit_fn             = 5
	KlangParserRULE_json_delete_fn           = 6
	KlangParserRULE_yaml_edit_fn             = 7
	KlangParserRULE_yaml_delete_fn           = 8
	KlangParserRULE_kube_json_edit_fn        = 9
	KlangParserRULE_kube_json_delete_fn      = 10
	KlangParserRULE_kube_yaml_edit_fn        = 11
	KlangParserRULE_kube_yaml_delete_fn      = 12
	KlangParserRULE_sleep_fn                 = 13
	KlangParserRULE_exit_fn                  = 14
	KlangParserRULE_if_stat                  = 15
	KlangParserRULE_condition_block          = 16
	KlangParserRULE_stat_block               = 17
	KlangParserRULE_while_stat               = 18
	KlangParserRULE_log                      = 19
	KlangParserRULE_kubectl_command          = 20
	KlangParserRULE_download_fn              = 21
	KlangParserRULE_json_select_fn           = 22
	KlangParserRULE_yaml_select_fn           = 23
	KlangParserRULE_load_fn                  = 24
	KlangParserRULE_stepInfo                 = 25
	KlangParserRULE_ns                       = 26
	KlangParserRULE_patch_type               = 27
	KlangParserRULE_string_or_id             = 28
	KlangParserRULE_resource                 = 29
	KlangParserRULE_kubernetes_object_config = 30
	KlangParserRULE_filter                   = 31
	KlangParserRULE_pattern                  = 32
	KlangParserRULE_expr                     = 33
	KlangParserRULE_atom                     = 34
	KlangParserRULE_json                     = 35
	KlangParserRULE_obj                      = 36
	KlangParserRULE_pair                     = 37
	KlangParserRULE_arr                      = 38
	KlangParserRULE_value                    = 39
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(KlangParserEOF, 0)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitParse(s)
	}
}

func (p *KlangParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, KlangParserRULE_parse)

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
		p.SetState(80)
		p.Block()
	}
	{
		p.SetState(81)
		p.Match(KlangParserEOF)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStat() []IStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatContext)(nil)).Elem())
	var tst = make([]IStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatContext)
		}
	}

	return tst
}

func (s *BlockContext) Stat(i int) IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *KlangParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, KlangParserRULE_block)
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
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la-30)&-(0x1f+1)) == 0 && ((1<<uint((_la-30)))&((1<<(KlangParserIF-30))|(1<<(KlangParserWHILE-30))|(1<<(KlangParserLOG-30))|(1<<(KlangParserEXIT-30))|(1<<(KlangParserJSONEDIT-30))|(1<<(KlangParserJSONDELETE-30))|(1<<(KlangParserYAMLEDIT-30))|(1<<(KlangParserYAMLDELETE-30))|(1<<(KlangParserKUBEJSONEDIT-30))|(1<<(KlangParserKUBEJSONDELETE-30))|(1<<(KlangParserKUBEYAMLEDIT-30))|(1<<(KlangParserKUBEYAMLDELETE-30))|(1<<(KlangParserSLEEP-30)))) != 0) || _la == KlangParserID || _la == KlangParserOTHER {
		{
			p.SetState(83)
			p.Stat()
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_OTHER returns the _OTHER token.
	Get_OTHER() antlr.Token

	// Set_OTHER sets the _OTHER token.
	Set_OTHER(antlr.Token)

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_OTHER antlr.Token
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) Get_OTHER() antlr.Token { return s._OTHER }

func (s *StatContext) Set_OTHER(v antlr.Token) { s._OTHER = v }

func (s *StatContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatContext) Json_edit_fn() IJson_edit_fnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJson_edit_fnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJson_edit_fnContext)
}

func (s *StatContext) Json_delete_fn() IJson_delete_fnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJson_delete_fnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJson_delete_fnContext)
}

func (s *StatContext) Yaml_edit_fn() IYaml_edit_fnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYaml_edit_fnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYaml_edit_fnContext)
}

func (s *StatContext) Yaml_delete_fn() IYaml_delete_fnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYaml_delete_fnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYaml_delete_fnContext)
}

func (s *StatContext) Kube_json_delete_fn() IKube_json_delete_fnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKube_json_delete_fnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKube_json_delete_fnContext)
}

func (s *StatContext) Kube_json_edit_fn() IKube_json_edit_fnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKube_json_edit_fnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKube_json_edit_fnContext)
}

func (s *StatContext) Kube_yaml_delete_fn() IKube_yaml_delete_fnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKube_yaml_delete_fnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKube_yaml_delete_fnContext)
}

func (s *StatContext) Kube_yaml_edit_fn() IKube_yaml_edit_fnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKube_yaml_edit_fnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKube_yaml_edit_fnContext)
}

func (s *StatContext) If_stat() IIf_statContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_statContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_statContext)
}

func (s *StatContext) While_stat() IWhile_statContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhile_statContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhile_statContext)
}

func (s *StatContext) Sleep_fn() ISleep_fnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISleep_fnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISleep_fnContext)
}

func (s *StatContext) Exit_fn() IExit_fnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExit_fnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExit_fnContext)
}

func (s *StatContext) Log() ILogContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogContext)
}

func (s *StatContext) OTHER() antlr.TerminalNode {
	return s.GetToken(KlangParserOTHER, 0)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitStat(s)
	}
}

func (p *KlangParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, KlangParserRULE_stat)

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

	p.SetState(105)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)
			p.Assignment()
		}

	case KlangParserJSONEDIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.Json_edit_fn()
		}

	case KlangParserJSONDELETE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(91)
			p.Json_delete_fn()
		}

	case KlangParserYAMLEDIT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(92)
			p.Yaml_edit_fn()
		}

	case KlangParserYAMLDELETE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(93)
			p.Yaml_delete_fn()
		}

	case KlangParserKUBEJSONDELETE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(94)
			p.Kube_json_delete_fn()
		}

	case KlangParserKUBEJSONEDIT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(95)
			p.Kube_json_edit_fn()
		}

	case KlangParserKUBEYAMLDELETE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(96)
			p.Kube_yaml_delete_fn()
		}

	case KlangParserKUBEYAMLEDIT:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(97)
			p.Kube_yaml_edit_fn()
		}

	case KlangParserIF:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(98)
			p.If_stat()
		}

	case KlangParserWHILE:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(99)
			p.While_stat()
		}

	case KlangParserSLEEP:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(100)
			p.Sleep_fn()
		}

	case KlangParserEXIT:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(101)
			p.Exit_fn()
		}

	case KlangParserLOG:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(102)
			p.Log()
		}

	case KlangParserOTHER:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(103)

			var _m = p.Match(KlangParserOTHER)

			localctx.(*StatContext)._OTHER = _m
		}
		fmt.Println("unknown char: " + (func() string {
			if localctx.(*StatContext).Get_OTHER() == nil {
				return ""
			} else {
				return localctx.(*StatContext).Get_OTHER().GetText()
			}
		}()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(KlangParserID, 0)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(KlangParserASSIGN, 0)
}

func (s *AssignmentContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignmentContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KlangParserSCOL, 0)
}

func (s *AssignmentContext) Load_fn() ILoad_fnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoad_fnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoad_fnContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *KlangParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, KlangParserRULE_assignment)

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

	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.Match(KlangParserID)
		}
		{
			p.SetState(108)
			p.Match(KlangParserASSIGN)
		}
		{
			p.SetState(109)
			p.expr(0)
		}
		{
			p.SetState(110)
			p.Match(KlangParserSCOL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Match(KlangParserID)
		}
		{
			p.SetState(113)
			p.Match(KlangParserASSIGN)
		}
		{
			p.SetState(114)
			p.Load_fn()
		}
		{
			p.SetState(115)
			p.Match(KlangParserSCOL)
		}

	}

	return localctx
}

// IShell_scriptContext is an interface to support dynamic dispatch.
type IShell_scriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShell_scriptContext differentiates from other interfaces.
	IsShell_scriptContext()
}

type Shell_scriptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShell_scriptContext() *Shell_scriptContext {
	var p = new(Shell_scriptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_shell_script
	return p
}

func (*Shell_scriptContext) IsShell_scriptContext() {}

func NewShell_scriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Shell_scriptContext {
	var p = new(Shell_scriptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_shell_script

	return p
}

func (s *Shell_scriptContext) GetParser() antlr.Parser { return s.parser }

func (s *Shell_scriptContext) SHELLSCRIPT() antlr.TerminalNode {
	return s.GetToken(KlangParserSHELLSCRIPT, 0)
}

func (s *Shell_scriptContext) String_or_id() IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *Shell_scriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Shell_scriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Shell_scriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterShell_script(s)
	}
}

func (s *Shell_scriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitShell_script(s)
	}
}

func (p *KlangParser) Shell_script() (localctx IShell_scriptContext) {
	localctx = NewShell_scriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, KlangParserRULE_shell_script)

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
		p.SetState(119)
		p.Match(KlangParserSHELLSCRIPT)
	}
	{
		p.SetState(120)
		p.String_or_id()
	}

	return localctx
}

// IJson_edit_fnContext is an interface to support dynamic dispatch.
type IJson_edit_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJson_edit_fnContext differentiates from other interfaces.
	IsJson_edit_fnContext()
}

type Json_edit_fnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJson_edit_fnContext() *Json_edit_fnContext {
	var p = new(Json_edit_fnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_json_edit_fn
	return p
}

func (*Json_edit_fnContext) IsJson_edit_fnContext() {}

func NewJson_edit_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Json_edit_fnContext {
	var p = new(Json_edit_fnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_json_edit_fn

	return p
}

func (s *Json_edit_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Json_edit_fnContext) JSONEDIT() antlr.TerminalNode {
	return s.GetToken(KlangParserJSONEDIT, 0)
}

func (s *Json_edit_fnContext) OPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserOPAR, 0)
}

func (s *Json_edit_fnContext) ID() antlr.TerminalNode {
	return s.GetToken(KlangParserID, 0)
}

func (s *Json_edit_fnContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KlangParserCOMMA)
}

func (s *Json_edit_fnContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KlangParserCOMMA, i)
}

func (s *Json_edit_fnContext) String_or_id() IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *Json_edit_fnContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Json_edit_fnContext) CPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserCPAR, 0)
}

func (s *Json_edit_fnContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KlangParserSCOL, 0)
}

func (s *Json_edit_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Json_edit_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Json_edit_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterJson_edit_fn(s)
	}
}

func (s *Json_edit_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitJson_edit_fn(s)
	}
}

func (p *KlangParser) Json_edit_fn() (localctx IJson_edit_fnContext) {
	localctx = NewJson_edit_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, KlangParserRULE_json_edit_fn)

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
		p.SetState(122)
		p.Match(KlangParserJSONEDIT)
	}
	{
		p.SetState(123)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(124)
		p.Match(KlangParserID)
	}
	{
		p.SetState(125)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(126)
		p.String_or_id()
	}
	{
		p.SetState(127)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(128)
		p.expr(0)
	}
	{
		p.SetState(129)
		p.Match(KlangParserCPAR)
	}
	{
		p.SetState(130)
		p.Match(KlangParserSCOL)
	}

	return localctx
}

// IJson_delete_fnContext is an interface to support dynamic dispatch.
type IJson_delete_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJson_delete_fnContext differentiates from other interfaces.
	IsJson_delete_fnContext()
}

type Json_delete_fnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJson_delete_fnContext() *Json_delete_fnContext {
	var p = new(Json_delete_fnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_json_delete_fn
	return p
}

func (*Json_delete_fnContext) IsJson_delete_fnContext() {}

func NewJson_delete_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Json_delete_fnContext {
	var p = new(Json_delete_fnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_json_delete_fn

	return p
}

func (s *Json_delete_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Json_delete_fnContext) JSONDELETE() antlr.TerminalNode {
	return s.GetToken(KlangParserJSONDELETE, 0)
}

func (s *Json_delete_fnContext) OPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserOPAR, 0)
}

func (s *Json_delete_fnContext) ID() antlr.TerminalNode {
	return s.GetToken(KlangParserID, 0)
}

func (s *Json_delete_fnContext) COMMA() antlr.TerminalNode {
	return s.GetToken(KlangParserCOMMA, 0)
}

func (s *Json_delete_fnContext) String_or_id() IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *Json_delete_fnContext) CPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserCPAR, 0)
}

func (s *Json_delete_fnContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KlangParserSCOL, 0)
}

func (s *Json_delete_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Json_delete_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Json_delete_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterJson_delete_fn(s)
	}
}

func (s *Json_delete_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitJson_delete_fn(s)
	}
}

func (p *KlangParser) Json_delete_fn() (localctx IJson_delete_fnContext) {
	localctx = NewJson_delete_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, KlangParserRULE_json_delete_fn)

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
		p.SetState(132)
		p.Match(KlangParserJSONDELETE)
	}
	{
		p.SetState(133)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(134)
		p.Match(KlangParserID)
	}
	{
		p.SetState(135)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(136)
		p.String_or_id()
	}
	{
		p.SetState(137)
		p.Match(KlangParserCPAR)
	}
	{
		p.SetState(138)
		p.Match(KlangParserSCOL)
	}

	return localctx
}

// IYaml_edit_fnContext is an interface to support dynamic dispatch.
type IYaml_edit_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsYaml_edit_fnContext differentiates from other interfaces.
	IsYaml_edit_fnContext()
}

type Yaml_edit_fnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYaml_edit_fnContext() *Yaml_edit_fnContext {
	var p = new(Yaml_edit_fnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_yaml_edit_fn
	return p
}

func (*Yaml_edit_fnContext) IsYaml_edit_fnContext() {}

func NewYaml_edit_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Yaml_edit_fnContext {
	var p = new(Yaml_edit_fnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_yaml_edit_fn

	return p
}

func (s *Yaml_edit_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Yaml_edit_fnContext) YAMLEDIT() antlr.TerminalNode {
	return s.GetToken(KlangParserYAMLEDIT, 0)
}

func (s *Yaml_edit_fnContext) OPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserOPAR, 0)
}

func (s *Yaml_edit_fnContext) ID() antlr.TerminalNode {
	return s.GetToken(KlangParserID, 0)
}

func (s *Yaml_edit_fnContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KlangParserCOMMA)
}

func (s *Yaml_edit_fnContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KlangParserCOMMA, i)
}

func (s *Yaml_edit_fnContext) String_or_id() IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *Yaml_edit_fnContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Yaml_edit_fnContext) CPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserCPAR, 0)
}

func (s *Yaml_edit_fnContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KlangParserSCOL, 0)
}

func (s *Yaml_edit_fnContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(KlangParserNUMBER, 0)
}

func (s *Yaml_edit_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Yaml_edit_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Yaml_edit_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterYaml_edit_fn(s)
	}
}

func (s *Yaml_edit_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitYaml_edit_fn(s)
	}
}

func (p *KlangParser) Yaml_edit_fn() (localctx IYaml_edit_fnContext) {
	localctx = NewYaml_edit_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, KlangParserRULE_yaml_edit_fn)
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
		p.SetState(140)
		p.Match(KlangParserYAMLEDIT)
	}
	{
		p.SetState(141)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(142)
		p.Match(KlangParserID)
	}
	{
		p.SetState(143)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(144)
		p.String_or_id()
	}
	{
		p.SetState(145)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(146)
		p.expr(0)
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KlangParserCOMMA {
		{
			p.SetState(147)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(148)
			p.Match(KlangParserNUMBER)
		}

	}
	{
		p.SetState(151)
		p.Match(KlangParserCPAR)
	}
	{
		p.SetState(152)
		p.Match(KlangParserSCOL)
	}

	return localctx
}

// IYaml_delete_fnContext is an interface to support dynamic dispatch.
type IYaml_delete_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsYaml_delete_fnContext differentiates from other interfaces.
	IsYaml_delete_fnContext()
}

type Yaml_delete_fnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYaml_delete_fnContext() *Yaml_delete_fnContext {
	var p = new(Yaml_delete_fnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_yaml_delete_fn
	return p
}

func (*Yaml_delete_fnContext) IsYaml_delete_fnContext() {}

func NewYaml_delete_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Yaml_delete_fnContext {
	var p = new(Yaml_delete_fnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_yaml_delete_fn

	return p
}

func (s *Yaml_delete_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Yaml_delete_fnContext) YAMLDELETE() antlr.TerminalNode {
	return s.GetToken(KlangParserYAMLDELETE, 0)
}

func (s *Yaml_delete_fnContext) OPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserOPAR, 0)
}

func (s *Yaml_delete_fnContext) ID() antlr.TerminalNode {
	return s.GetToken(KlangParserID, 0)
}

func (s *Yaml_delete_fnContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KlangParserCOMMA)
}

func (s *Yaml_delete_fnContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KlangParserCOMMA, i)
}

func (s *Yaml_delete_fnContext) String_or_id() IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *Yaml_delete_fnContext) CPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserCPAR, 0)
}

func (s *Yaml_delete_fnContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KlangParserSCOL, 0)
}

func (s *Yaml_delete_fnContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(KlangParserNUMBER, 0)
}

func (s *Yaml_delete_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Yaml_delete_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Yaml_delete_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterYaml_delete_fn(s)
	}
}

func (s *Yaml_delete_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitYaml_delete_fn(s)
	}
}

func (p *KlangParser) Yaml_delete_fn() (localctx IYaml_delete_fnContext) {
	localctx = NewYaml_delete_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, KlangParserRULE_yaml_delete_fn)
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
		p.SetState(154)
		p.Match(KlangParserYAMLDELETE)
	}
	{
		p.SetState(155)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(156)
		p.Match(KlangParserID)
	}
	{
		p.SetState(157)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(158)
		p.String_or_id()
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KlangParserCOMMA {
		{
			p.SetState(159)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(160)
			p.Match(KlangParserNUMBER)
		}

	}
	{
		p.SetState(163)
		p.Match(KlangParserCPAR)
	}
	{
		p.SetState(164)
		p.Match(KlangParserSCOL)
	}

	return localctx
}

// IKube_json_edit_fnContext is an interface to support dynamic dispatch.
type IKube_json_edit_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKube_json_edit_fnContext differentiates from other interfaces.
	IsKube_json_edit_fnContext()
}

type Kube_json_edit_fnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKube_json_edit_fnContext() *Kube_json_edit_fnContext {
	var p = new(Kube_json_edit_fnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_kube_json_edit_fn
	return p
}

func (*Kube_json_edit_fnContext) IsKube_json_edit_fnContext() {}

func NewKube_json_edit_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Kube_json_edit_fnContext {
	var p = new(Kube_json_edit_fnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_kube_json_edit_fn

	return p
}

func (s *Kube_json_edit_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Kube_json_edit_fnContext) KUBEJSONEDIT() antlr.TerminalNode {
	return s.GetToken(KlangParserKUBEJSONEDIT, 0)
}

func (s *Kube_json_edit_fnContext) OPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserOPAR, 0)
}

func (s *Kube_json_edit_fnContext) ID() antlr.TerminalNode {
	return s.GetToken(KlangParserID, 0)
}

func (s *Kube_json_edit_fnContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KlangParserCOMMA)
}

func (s *Kube_json_edit_fnContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KlangParserCOMMA, i)
}

func (s *Kube_json_edit_fnContext) AllString_or_id() []IString_or_idContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IString_or_idContext)(nil)).Elem())
	var tst = make([]IString_or_idContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IString_or_idContext)
		}
	}

	return tst
}

func (s *Kube_json_edit_fnContext) String_or_id(i int) IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *Kube_json_edit_fnContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Kube_json_edit_fnContext) CPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserCPAR, 0)
}

func (s *Kube_json_edit_fnContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KlangParserSCOL, 0)
}

func (s *Kube_json_edit_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Kube_json_edit_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Kube_json_edit_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterKube_json_edit_fn(s)
	}
}

func (s *Kube_json_edit_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitKube_json_edit_fn(s)
	}
}

func (p *KlangParser) Kube_json_edit_fn() (localctx IKube_json_edit_fnContext) {
	localctx = NewKube_json_edit_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, KlangParserRULE_kube_json_edit_fn)
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
		p.SetState(166)
		p.Match(KlangParserKUBEJSONEDIT)
	}
	{
		p.SetState(167)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(168)
		p.Match(KlangParserID)
	}
	{
		p.SetState(169)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(170)
		p.String_or_id()
	}
	{
		p.SetState(171)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(172)
		p.expr(0)
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KlangParserCOMMA {
		{
			p.SetState(173)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(174)
			p.String_or_id()
		}

	}
	{
		p.SetState(177)
		p.Match(KlangParserCPAR)
	}
	{
		p.SetState(178)
		p.Match(KlangParserSCOL)
	}

	return localctx
}

// IKube_json_delete_fnContext is an interface to support dynamic dispatch.
type IKube_json_delete_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKube_json_delete_fnContext differentiates from other interfaces.
	IsKube_json_delete_fnContext()
}

type Kube_json_delete_fnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKube_json_delete_fnContext() *Kube_json_delete_fnContext {
	var p = new(Kube_json_delete_fnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_kube_json_delete_fn
	return p
}

func (*Kube_json_delete_fnContext) IsKube_json_delete_fnContext() {}

func NewKube_json_delete_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Kube_json_delete_fnContext {
	var p = new(Kube_json_delete_fnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_kube_json_delete_fn

	return p
}

func (s *Kube_json_delete_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Kube_json_delete_fnContext) KUBEJSONDELETE() antlr.TerminalNode {
	return s.GetToken(KlangParserKUBEJSONDELETE, 0)
}

func (s *Kube_json_delete_fnContext) OPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserOPAR, 0)
}

func (s *Kube_json_delete_fnContext) ID() antlr.TerminalNode {
	return s.GetToken(KlangParserID, 0)
}

func (s *Kube_json_delete_fnContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KlangParserCOMMA)
}

func (s *Kube_json_delete_fnContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KlangParserCOMMA, i)
}

func (s *Kube_json_delete_fnContext) Filter() IFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *Kube_json_delete_fnContext) Pattern() IPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatternContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *Kube_json_delete_fnContext) CPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserCPAR, 0)
}

func (s *Kube_json_delete_fnContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KlangParserSCOL, 0)
}

func (s *Kube_json_delete_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Kube_json_delete_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Kube_json_delete_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterKube_json_delete_fn(s)
	}
}

func (s *Kube_json_delete_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitKube_json_delete_fn(s)
	}
}

func (p *KlangParser) Kube_json_delete_fn() (localctx IKube_json_delete_fnContext) {
	localctx = NewKube_json_delete_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, KlangParserRULE_kube_json_delete_fn)

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

	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Match(KlangParserKUBEJSONDELETE)
		}
		{
			p.SetState(181)
			p.Match(KlangParserOPAR)
		}
		{
			p.SetState(182)
			p.Match(KlangParserID)
		}
		{
			p.SetState(183)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(184)
			p.Filter()
		}
		{
			p.SetState(185)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(186)
			p.Pattern()
		}
		{
			p.SetState(187)
			p.Match(KlangParserCPAR)
		}
		{
			p.SetState(188)
			p.Match(KlangParserSCOL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Match(KlangParserKUBEJSONDELETE)
		}
		{
			p.SetState(191)
			p.Match(KlangParserOPAR)
		}
		{
			p.SetState(192)
			p.Match(KlangParserID)
		}
		{
			p.SetState(193)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(194)
			p.Pattern()
		}
		{
			p.SetState(195)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(196)
			p.Filter()
		}
		{
			p.SetState(197)
			p.Match(KlangParserCPAR)
		}
		{
			p.SetState(198)
			p.Match(KlangParserSCOL)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(200)
			p.Match(KlangParserKUBEJSONDELETE)
		}
		{
			p.SetState(201)
			p.Match(KlangParserOPAR)
		}
		{
			p.SetState(202)
			p.Match(KlangParserID)
		}
		{
			p.SetState(203)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(204)
			p.Filter()
		}
		{
			p.SetState(205)
			p.Match(KlangParserCPAR)
		}
		{
			p.SetState(206)
			p.Match(KlangParserSCOL)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(208)
			p.Match(KlangParserKUBEJSONDELETE)
		}
		{
			p.SetState(209)
			p.Match(KlangParserOPAR)
		}
		{
			p.SetState(210)
			p.Match(KlangParserID)
		}
		{
			p.SetState(211)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(212)
			p.Pattern()
		}
		{
			p.SetState(213)
			p.Match(KlangParserCPAR)
		}
		{
			p.SetState(214)
			p.Match(KlangParserSCOL)
		}

	}

	return localctx
}

// IKube_yaml_edit_fnContext is an interface to support dynamic dispatch.
type IKube_yaml_edit_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKube_yaml_edit_fnContext differentiates from other interfaces.
	IsKube_yaml_edit_fnContext()
}

type Kube_yaml_edit_fnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKube_yaml_edit_fnContext() *Kube_yaml_edit_fnContext {
	var p = new(Kube_yaml_edit_fnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_kube_yaml_edit_fn
	return p
}

func (*Kube_yaml_edit_fnContext) IsKube_yaml_edit_fnContext() {}

func NewKube_yaml_edit_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Kube_yaml_edit_fnContext {
	var p = new(Kube_yaml_edit_fnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_kube_yaml_edit_fn

	return p
}

func (s *Kube_yaml_edit_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Kube_yaml_edit_fnContext) KUBEYAMLEDIT() antlr.TerminalNode {
	return s.GetToken(KlangParserKUBEYAMLEDIT, 0)
}

func (s *Kube_yaml_edit_fnContext) OPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserOPAR, 0)
}

func (s *Kube_yaml_edit_fnContext) ID() antlr.TerminalNode {
	return s.GetToken(KlangParserID, 0)
}

func (s *Kube_yaml_edit_fnContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KlangParserCOMMA)
}

func (s *Kube_yaml_edit_fnContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KlangParserCOMMA, i)
}

func (s *Kube_yaml_edit_fnContext) AllString_or_id() []IString_or_idContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IString_or_idContext)(nil)).Elem())
	var tst = make([]IString_or_idContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IString_or_idContext)
		}
	}

	return tst
}

func (s *Kube_yaml_edit_fnContext) String_or_id(i int) IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *Kube_yaml_edit_fnContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Kube_yaml_edit_fnContext) CPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserCPAR, 0)
}

func (s *Kube_yaml_edit_fnContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KlangParserSCOL, 0)
}

func (s *Kube_yaml_edit_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Kube_yaml_edit_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Kube_yaml_edit_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterKube_yaml_edit_fn(s)
	}
}

func (s *Kube_yaml_edit_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitKube_yaml_edit_fn(s)
	}
}

func (p *KlangParser) Kube_yaml_edit_fn() (localctx IKube_yaml_edit_fnContext) {
	localctx = NewKube_yaml_edit_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, KlangParserRULE_kube_yaml_edit_fn)
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
		p.SetState(218)
		p.Match(KlangParserKUBEYAMLEDIT)
	}
	{
		p.SetState(219)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(220)
		p.Match(KlangParserID)
	}
	{
		p.SetState(221)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(222)
		p.String_or_id()
	}
	{
		p.SetState(223)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(224)
		p.expr(0)
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KlangParserCOMMA {
		{
			p.SetState(225)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(226)
			p.String_or_id()
		}

	}
	{
		p.SetState(229)
		p.Match(KlangParserCPAR)
	}
	{
		p.SetState(230)
		p.Match(KlangParserSCOL)
	}

	return localctx
}

// IKube_yaml_delete_fnContext is an interface to support dynamic dispatch.
type IKube_yaml_delete_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKube_yaml_delete_fnContext differentiates from other interfaces.
	IsKube_yaml_delete_fnContext()
}

type Kube_yaml_delete_fnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKube_yaml_delete_fnContext() *Kube_yaml_delete_fnContext {
	var p = new(Kube_yaml_delete_fnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_kube_yaml_delete_fn
	return p
}

func (*Kube_yaml_delete_fnContext) IsKube_yaml_delete_fnContext() {}

func NewKube_yaml_delete_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Kube_yaml_delete_fnContext {
	var p = new(Kube_yaml_delete_fnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_kube_yaml_delete_fn

	return p
}

func (s *Kube_yaml_delete_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Kube_yaml_delete_fnContext) KUBEYAMLDELETE() antlr.TerminalNode {
	return s.GetToken(KlangParserKUBEYAMLDELETE, 0)
}

func (s *Kube_yaml_delete_fnContext) OPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserOPAR, 0)
}

func (s *Kube_yaml_delete_fnContext) ID() antlr.TerminalNode {
	return s.GetToken(KlangParserID, 0)
}

func (s *Kube_yaml_delete_fnContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KlangParserCOMMA)
}

func (s *Kube_yaml_delete_fnContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KlangParserCOMMA, i)
}

func (s *Kube_yaml_delete_fnContext) Filter() IFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *Kube_yaml_delete_fnContext) Pattern() IPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatternContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *Kube_yaml_delete_fnContext) CPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserCPAR, 0)
}

func (s *Kube_yaml_delete_fnContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KlangParserSCOL, 0)
}

func (s *Kube_yaml_delete_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Kube_yaml_delete_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Kube_yaml_delete_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterKube_yaml_delete_fn(s)
	}
}

func (s *Kube_yaml_delete_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitKube_yaml_delete_fn(s)
	}
}

func (p *KlangParser) Kube_yaml_delete_fn() (localctx IKube_yaml_delete_fnContext) {
	localctx = NewKube_yaml_delete_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, KlangParserRULE_kube_yaml_delete_fn)

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

	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			p.Match(KlangParserKUBEYAMLDELETE)
		}
		{
			p.SetState(233)
			p.Match(KlangParserOPAR)
		}
		{
			p.SetState(234)
			p.Match(KlangParserID)
		}
		{
			p.SetState(235)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(236)
			p.Filter()
		}
		{
			p.SetState(237)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(238)
			p.Pattern()
		}
		{
			p.SetState(239)
			p.Match(KlangParserCPAR)
		}
		{
			p.SetState(240)
			p.Match(KlangParserSCOL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)
			p.Match(KlangParserKUBEYAMLDELETE)
		}
		{
			p.SetState(243)
			p.Match(KlangParserOPAR)
		}
		{
			p.SetState(244)
			p.Match(KlangParserID)
		}
		{
			p.SetState(245)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(246)
			p.Pattern()
		}
		{
			p.SetState(247)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(248)
			p.Filter()
		}
		{
			p.SetState(249)
			p.Match(KlangParserCPAR)
		}
		{
			p.SetState(250)
			p.Match(KlangParserSCOL)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(252)
			p.Match(KlangParserKUBEYAMLDELETE)
		}
		{
			p.SetState(253)
			p.Match(KlangParserOPAR)
		}
		{
			p.SetState(254)
			p.Match(KlangParserID)
		}
		{
			p.SetState(255)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(256)
			p.Filter()
		}
		{
			p.SetState(257)
			p.Match(KlangParserCPAR)
		}
		{
			p.SetState(258)
			p.Match(KlangParserSCOL)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(260)
			p.Match(KlangParserKUBEYAMLDELETE)
		}
		{
			p.SetState(261)
			p.Match(KlangParserOPAR)
		}
		{
			p.SetState(262)
			p.Match(KlangParserID)
		}
		{
			p.SetState(263)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(264)
			p.Pattern()
		}
		{
			p.SetState(265)
			p.Match(KlangParserCPAR)
		}
		{
			p.SetState(266)
			p.Match(KlangParserSCOL)
		}

	}

	return localctx
}

// ISleep_fnContext is an interface to support dynamic dispatch.
type ISleep_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSleep_fnContext differentiates from other interfaces.
	IsSleep_fnContext()
}

type Sleep_fnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySleep_fnContext() *Sleep_fnContext {
	var p = new(Sleep_fnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_sleep_fn
	return p
}

func (*Sleep_fnContext) IsSleep_fnContext() {}

func NewSleep_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sleep_fnContext {
	var p = new(Sleep_fnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_sleep_fn

	return p
}

func (s *Sleep_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Sleep_fnContext) SLEEP() antlr.TerminalNode {
	return s.GetToken(KlangParserSLEEP, 0)
}

func (s *Sleep_fnContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(KlangParserNUMBER, 0)
}

func (s *Sleep_fnContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KlangParserSCOL, 0)
}

func (s *Sleep_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sleep_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sleep_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterSleep_fn(s)
	}
}

func (s *Sleep_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitSleep_fn(s)
	}
}

func (p *KlangParser) Sleep_fn() (localctx ISleep_fnContext) {
	localctx = NewSleep_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, KlangParserRULE_sleep_fn)

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
		p.SetState(270)
		p.Match(KlangParserSLEEP)
	}
	{
		p.SetState(271)
		p.Match(KlangParserNUMBER)
	}
	{
		p.SetState(272)
		p.Match(KlangParserSCOL)
	}

	return localctx
}

// IExit_fnContext is an interface to support dynamic dispatch.
type IExit_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExit_fnContext differentiates from other interfaces.
	IsExit_fnContext()
}

type Exit_fnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExit_fnContext() *Exit_fnContext {
	var p = new(Exit_fnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_exit_fn
	return p
}

func (*Exit_fnContext) IsExit_fnContext() {}

func NewExit_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exit_fnContext {
	var p = new(Exit_fnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_exit_fn

	return p
}

func (s *Exit_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Exit_fnContext) EXIT() antlr.TerminalNode {
	return s.GetToken(KlangParserEXIT, 0)
}

func (s *Exit_fnContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(KlangParserNUMBER, 0)
}

func (s *Exit_fnContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KlangParserSCOL, 0)
}

func (s *Exit_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exit_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exit_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterExit_fn(s)
	}
}

func (s *Exit_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitExit_fn(s)
	}
}

func (p *KlangParser) Exit_fn() (localctx IExit_fnContext) {
	localctx = NewExit_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, KlangParserRULE_exit_fn)

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
		p.SetState(274)
		p.Match(KlangParserEXIT)
	}
	{
		p.SetState(275)
		p.Match(KlangParserNUMBER)
	}
	{
		p.SetState(276)
		p.Match(KlangParserSCOL)
	}

	return localctx
}

// IIf_statContext is an interface to support dynamic dispatch.
type IIf_statContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_statContext differentiates from other interfaces.
	IsIf_statContext()
}

type If_statContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_statContext() *If_statContext {
	var p = new(If_statContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_if_stat
	return p
}

func (*If_statContext) IsIf_statContext() {}

func NewIf_statContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statContext {
	var p = new(If_statContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_if_stat

	return p
}

func (s *If_statContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statContext) AllIF() []antlr.TerminalNode {
	return s.GetTokens(KlangParserIF)
}

func (s *If_statContext) IF(i int) antlr.TerminalNode {
	return s.GetToken(KlangParserIF, i)
}

func (s *If_statContext) AllCondition_block() []ICondition_blockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICondition_blockContext)(nil)).Elem())
	var tst = make([]ICondition_blockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICondition_blockContext)
		}
	}

	return tst
}

func (s *If_statContext) Condition_block(i int) ICondition_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondition_blockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICondition_blockContext)
}

func (s *If_statContext) AllELSE() []antlr.TerminalNode {
	return s.GetTokens(KlangParserELSE)
}

func (s *If_statContext) ELSE(i int) antlr.TerminalNode {
	return s.GetToken(KlangParserELSE, i)
}

func (s *If_statContext) Stat_block() IStat_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStat_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStat_blockContext)
}

func (s *If_statContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_statContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterIf_stat(s)
	}
}

func (s *If_statContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitIf_stat(s)
	}
}

func (p *KlangParser) If_stat() (localctx IIf_statContext) {
	localctx = NewIf_statContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, KlangParserRULE_if_stat)

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
		p.SetState(278)
		p.Match(KlangParserIF)
	}
	{
		p.SetState(279)
		p.Condition_block()
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(280)
				p.Match(KlangParserELSE)
			}
			{
				p.SetState(281)
				p.Match(KlangParserIF)
			}
			{
				p.SetState(282)
				p.Condition_block()
			}

		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(288)
			p.Match(KlangParserELSE)
		}
		{
			p.SetState(289)
			p.Stat_block()
		}

	}

	return localctx
}

// ICondition_blockContext is an interface to support dynamic dispatch.
type ICondition_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCondition_blockContext differentiates from other interfaces.
	IsCondition_blockContext()
}

type Condition_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondition_blockContext() *Condition_blockContext {
	var p = new(Condition_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_condition_block
	return p
}

func (*Condition_blockContext) IsCondition_blockContext() {}

func NewCondition_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Condition_blockContext {
	var p = new(Condition_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_condition_block

	return p
}

func (s *Condition_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Condition_blockContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Condition_blockContext) Stat_block() IStat_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStat_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStat_blockContext)
}

func (s *Condition_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Condition_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Condition_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterCondition_block(s)
	}
}

func (s *Condition_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitCondition_block(s)
	}
}

func (p *KlangParser) Condition_block() (localctx ICondition_blockContext) {
	localctx = NewCondition_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, KlangParserRULE_condition_block)

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
		p.SetState(292)
		p.expr(0)
	}
	{
		p.SetState(293)
		p.Stat_block()
	}

	return localctx
}

// IStat_blockContext is an interface to support dynamic dispatch.
type IStat_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStat_blockContext differentiates from other interfaces.
	IsStat_blockContext()
}

type Stat_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStat_blockContext() *Stat_blockContext {
	var p = new(Stat_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_stat_block
	return p
}

func (*Stat_blockContext) IsStat_blockContext() {}

func NewStat_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stat_blockContext {
	var p = new(Stat_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_stat_block

	return p
}

func (s *Stat_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Stat_blockContext) OBRACE() antlr.TerminalNode {
	return s.GetToken(KlangParserOBRACE, 0)
}

func (s *Stat_blockContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Stat_blockContext) CBRACE() antlr.TerminalNode {
	return s.GetToken(KlangParserCBRACE, 0)
}

func (s *Stat_blockContext) Stat() IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *Stat_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stat_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Stat_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterStat_block(s)
	}
}

func (s *Stat_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitStat_block(s)
	}
}

func (p *KlangParser) Stat_block() (localctx IStat_blockContext) {
	localctx = NewStat_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, KlangParserRULE_stat_block)

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

	p.SetState(300)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserOBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(295)
			p.Match(KlangParserOBRACE)
		}
		{
			p.SetState(296)
			p.Block()
		}
		{
			p.SetState(297)
			p.Match(KlangParserCBRACE)
		}

	case KlangParserIF, KlangParserWHILE, KlangParserLOG, KlangParserEXIT, KlangParserJSONEDIT, KlangParserJSONDELETE, KlangParserYAMLEDIT, KlangParserYAMLDELETE, KlangParserKUBEJSONEDIT, KlangParserKUBEJSONDELETE, KlangParserKUBEYAMLEDIT, KlangParserKUBEYAMLDELETE, KlangParserSLEEP, KlangParserID, KlangParserOTHER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(299)
			p.Stat()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IWhile_statContext is an interface to support dynamic dispatch.
type IWhile_statContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhile_statContext differentiates from other interfaces.
	IsWhile_statContext()
}

type While_statContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_statContext() *While_statContext {
	var p = new(While_statContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_while_stat
	return p
}

func (*While_statContext) IsWhile_statContext() {}

func NewWhile_statContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_statContext {
	var p = new(While_statContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_while_stat

	return p
}

func (s *While_statContext) GetParser() antlr.Parser { return s.parser }

func (s *While_statContext) WHILE() antlr.TerminalNode {
	return s.GetToken(KlangParserWHILE, 0)
}

func (s *While_statContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *While_statContext) Stat_block() IStat_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStat_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStat_blockContext)
}

func (s *While_statContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_statContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_statContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterWhile_stat(s)
	}
}

func (s *While_statContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitWhile_stat(s)
	}
}

func (p *KlangParser) While_stat() (localctx IWhile_statContext) {
	localctx = NewWhile_statContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, KlangParserRULE_while_stat)

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
		p.SetState(302)
		p.Match(KlangParserWHILE)
	}
	{
		p.SetState(303)
		p.expr(0)
	}
	{
		p.SetState(304)
		p.Stat_block()
	}

	return localctx
}

// ILogContext is an interface to support dynamic dispatch.
type ILogContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogContext differentiates from other interfaces.
	IsLogContext()
}

type LogContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogContext() *LogContext {
	var p = new(LogContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_log
	return p
}

func (*LogContext) IsLogContext() {}

func NewLogContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogContext {
	var p = new(LogContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_log

	return p
}

func (s *LogContext) GetParser() antlr.Parser { return s.parser }

func (s *LogContext) LOG() antlr.TerminalNode {
	return s.GetToken(KlangParserLOG, 0)
}

func (s *LogContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LogContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KlangParserSCOL, 0)
}

func (s *LogContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterLog(s)
	}
}

func (s *LogContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitLog(s)
	}
}

func (p *KlangParser) Log() (localctx ILogContext) {
	localctx = NewLogContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, KlangParserRULE_log)

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
		p.SetState(306)
		p.Match(KlangParserLOG)
	}
	{
		p.SetState(307)
		p.expr(0)
	}
	{
		p.SetState(308)
		p.Match(KlangParserSCOL)
	}

	return localctx
}

// IKubectl_commandContext is an interface to support dynamic dispatch.
type IKubectl_commandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKubectl_commandContext differentiates from other interfaces.
	IsKubectl_commandContext()
}

type Kubectl_commandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKubectl_commandContext() *Kubectl_commandContext {
	var p = new(Kubectl_commandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_kubectl_command
	return p
}

func (*Kubectl_commandContext) IsKubectl_commandContext() {}

func NewKubectl_commandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Kubectl_commandContext {
	var p = new(Kubectl_commandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_kubectl_command

	return p
}

func (s *Kubectl_commandContext) GetParser() antlr.Parser { return s.parser }

func (s *Kubectl_commandContext) CopyFrom(ctx *Kubectl_commandContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Kubectl_commandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Kubectl_commandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeleteKubectlCommandContext struct {
	*Kubectl_commandContext
}

func NewDeleteKubectlCommandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeleteKubectlCommandContext {
	var p = new(DeleteKubectlCommandContext)

	p.Kubectl_commandContext = NewEmptyKubectl_commandContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Kubectl_commandContext))

	return p
}

func (s *DeleteKubectlCommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteKubectlCommandContext) KUBECTL() antlr.TerminalNode {
	return s.GetToken(KlangParserKUBECTL, 0)
}

func (s *DeleteKubectlCommandContext) DELETE() antlr.TerminalNode {
	return s.GetToken(KlangParserDELETE, 0)
}

func (s *DeleteKubectlCommandContext) AllNAMESPACE() []antlr.TerminalNode {
	return s.GetTokens(KlangParserNAMESPACE)
}

func (s *DeleteKubectlCommandContext) NAMESPACE(i int) antlr.TerminalNode {
	return s.GetToken(KlangParserNAMESPACE, i)
}

func (s *DeleteKubectlCommandContext) AllNs() []INsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INsContext)(nil)).Elem())
	var tst = make([]INsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INsContext)
		}
	}

	return tst
}

func (s *DeleteKubectlCommandContext) Ns(i int) INsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INsContext)
}

func (s *DeleteKubectlCommandContext) AllResource() []IResourceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IResourceContext)(nil)).Elem())
	var tst = make([]IResourceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IResourceContext)
		}
	}

	return tst
}

func (s *DeleteKubectlCommandContext) Resource(i int) IResourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResourceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IResourceContext)
}

func (s *DeleteKubectlCommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterDeleteKubectlCommand(s)
	}
}

func (s *DeleteKubectlCommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitDeleteKubectlCommand(s)
	}
}

type PatchKubectlCommandContext struct {
	*Kubectl_commandContext
}

func NewPatchKubectlCommandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatchKubectlCommandContext {
	var p = new(PatchKubectlCommandContext)

	p.Kubectl_commandContext = NewEmptyKubectl_commandContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Kubectl_commandContext))

	return p
}

func (s *PatchKubectlCommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatchKubectlCommandContext) KUBECTL() antlr.TerminalNode {
	return s.GetToken(KlangParserKUBECTL, 0)
}

func (s *PatchKubectlCommandContext) PATCH() antlr.TerminalNode {
	return s.GetToken(KlangParserPATCH, 0)
}

func (s *PatchKubectlCommandContext) AllNAMESPACE() []antlr.TerminalNode {
	return s.GetTokens(KlangParserNAMESPACE)
}

func (s *PatchKubectlCommandContext) NAMESPACE(i int) antlr.TerminalNode {
	return s.GetToken(KlangParserNAMESPACE, i)
}

func (s *PatchKubectlCommandContext) AllNs() []INsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INsContext)(nil)).Elem())
	var tst = make([]INsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INsContext)
		}
	}

	return tst
}

func (s *PatchKubectlCommandContext) Ns(i int) INsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INsContext)
}

func (s *PatchKubectlCommandContext) AllResource() []IResourceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IResourceContext)(nil)).Elem())
	var tst = make([]IResourceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IResourceContext)
		}
	}

	return tst
}

func (s *PatchKubectlCommandContext) Resource(i int) IResourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResourceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IResourceContext)
}

func (s *PatchKubectlCommandContext) AllPATCHTYPE() []antlr.TerminalNode {
	return s.GetTokens(KlangParserPATCHTYPE)
}

func (s *PatchKubectlCommandContext) PATCHTYPE(i int) antlr.TerminalNode {
	return s.GetToken(KlangParserPATCHTYPE, i)
}

func (s *PatchKubectlCommandContext) AllPatch_type() []IPatch_typeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPatch_typeContext)(nil)).Elem())
	var tst = make([]IPatch_typeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPatch_typeContext)
		}
	}

	return tst
}

func (s *PatchKubectlCommandContext) Patch_type(i int) IPatch_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatch_typeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPatch_typeContext)
}

func (s *PatchKubectlCommandContext) AllPATCHLOAD() []antlr.TerminalNode {
	return s.GetTokens(KlangParserPATCHLOAD)
}

func (s *PatchKubectlCommandContext) PATCHLOAD(i int) antlr.TerminalNode {
	return s.GetToken(KlangParserPATCHLOAD, i)
}

func (s *PatchKubectlCommandContext) AllString_or_id() []IString_or_idContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IString_or_idContext)(nil)).Elem())
	var tst = make([]IString_or_idContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IString_or_idContext)
		}
	}

	return tst
}

func (s *PatchKubectlCommandContext) String_or_id(i int) IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *PatchKubectlCommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterPatchKubectlCommand(s)
	}
}

func (s *PatchKubectlCommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitPatchKubectlCommand(s)
	}
}

type ApplyKubectlCommandContext struct {
	*Kubectl_commandContext
}

func NewApplyKubectlCommandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ApplyKubectlCommandContext {
	var p = new(ApplyKubectlCommandContext)

	p.Kubectl_commandContext = NewEmptyKubectl_commandContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Kubectl_commandContext))

	return p
}

func (s *ApplyKubectlCommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApplyKubectlCommandContext) KUBECTL() antlr.TerminalNode {
	return s.GetToken(KlangParserKUBECTL, 0)
}

func (s *ApplyKubectlCommandContext) APPLY() antlr.TerminalNode {
	return s.GetToken(KlangParserAPPLY, 0)
}

func (s *ApplyKubectlCommandContext) AllNAMESPACE() []antlr.TerminalNode {
	return s.GetTokens(KlangParserNAMESPACE)
}

func (s *ApplyKubectlCommandContext) NAMESPACE(i int) antlr.TerminalNode {
	return s.GetToken(KlangParserNAMESPACE, i)
}

func (s *ApplyKubectlCommandContext) AllNs() []INsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INsContext)(nil)).Elem())
	var tst = make([]INsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INsContext)
		}
	}

	return tst
}

func (s *ApplyKubectlCommandContext) Ns(i int) INsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INsContext)
}

func (s *ApplyKubectlCommandContext) AllString_or_id() []IString_or_idContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IString_or_idContext)(nil)).Elem())
	var tst = make([]IString_or_idContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IString_or_idContext)
		}
	}

	return tst
}

func (s *ApplyKubectlCommandContext) String_or_id(i int) IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *ApplyKubectlCommandContext) AllUPDATELOAD() []antlr.TerminalNode {
	return s.GetTokens(KlangParserUPDATELOAD)
}

func (s *ApplyKubectlCommandContext) UPDATELOAD(i int) antlr.TerminalNode {
	return s.GetToken(KlangParserUPDATELOAD, i)
}

func (s *ApplyKubectlCommandContext) AllKubernetes_object_config() []IKubernetes_object_configContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKubernetes_object_configContext)(nil)).Elem())
	var tst = make([]IKubernetes_object_configContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKubernetes_object_configContext)
		}
	}

	return tst
}

func (s *ApplyKubectlCommandContext) Kubernetes_object_config(i int) IKubernetes_object_configContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKubernetes_object_configContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKubernetes_object_configContext)
}

func (s *ApplyKubectlCommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterApplyKubectlCommand(s)
	}
}

func (s *ApplyKubectlCommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitApplyKubectlCommand(s)
	}
}

type GetKubectlCommandContext struct {
	*Kubectl_commandContext
}

func NewGetKubectlCommandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetKubectlCommandContext {
	var p = new(GetKubectlCommandContext)

	p.Kubectl_commandContext = NewEmptyKubectl_commandContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Kubectl_commandContext))

	return p
}

func (s *GetKubectlCommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetKubectlCommandContext) KUBECTL() antlr.TerminalNode {
	return s.GetToken(KlangParserKUBECTL, 0)
}

func (s *GetKubectlCommandContext) GET() antlr.TerminalNode {
	return s.GetToken(KlangParserGET, 0)
}

func (s *GetKubectlCommandContext) AllNAMESPACE() []antlr.TerminalNode {
	return s.GetTokens(KlangParserNAMESPACE)
}

func (s *GetKubectlCommandContext) NAMESPACE(i int) antlr.TerminalNode {
	return s.GetToken(KlangParserNAMESPACE, i)
}

func (s *GetKubectlCommandContext) AllNs() []INsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INsContext)(nil)).Elem())
	var tst = make([]INsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INsContext)
		}
	}

	return tst
}

func (s *GetKubectlCommandContext) Ns(i int) INsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INsContext)
}

func (s *GetKubectlCommandContext) AllResource() []IResourceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IResourceContext)(nil)).Elem())
	var tst = make([]IResourceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IResourceContext)
		}
	}

	return tst
}

func (s *GetKubectlCommandContext) Resource(i int) IResourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResourceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IResourceContext)
}

func (s *GetKubectlCommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterGetKubectlCommand(s)
	}
}

func (s *GetKubectlCommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitGetKubectlCommand(s)
	}
}

func (p *KlangParser) Kubectl_command() (localctx IKubectl_commandContext) {
	localctx = NewKubectl_commandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, KlangParserRULE_kubectl_command)

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

	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		localctx = NewApplyKubectlCommandContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(310)
			p.Match(KlangParserKUBECTL)
		}
		{
			p.SetState(311)
			p.Match(KlangParserAPPLY)
		}
		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(317)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case KlangParserNAMESPACE:
					{
						p.SetState(312)
						p.Match(KlangParserNAMESPACE)
					}
					{
						p.SetState(313)
						p.Ns()
					}

				case KlangParserID, KlangParserRAW_STRING_LIT, KlangParserSTRING:
					{
						p.SetState(314)
						p.String_or_id()
					}

				case KlangParserUPDATELOAD:
					{
						p.SetState(315)
						p.Match(KlangParserUPDATELOAD)
					}
					{
						p.SetState(316)
						p.Kubernetes_object_config()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(319)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
		}

	case 2:
		localctx = NewPatchKubectlCommandContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(321)
			p.Match(KlangParserKUBECTL)
		}
		{
			p.SetState(322)
			p.Match(KlangParserPATCH)
		}
		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(330)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case KlangParserNAMESPACE:
					{
						p.SetState(323)
						p.Match(KlangParserNAMESPACE)
					}
					{
						p.SetState(324)
						p.Ns()
					}

				case KlangParserID, KlangParserPATH, KlangParserRAW_STRING_LIT, KlangParserSTRING:
					{
						p.SetState(325)
						p.Resource()
					}

				case KlangParserPATCHTYPE:
					{
						p.SetState(326)
						p.Match(KlangParserPATCHTYPE)
					}
					{
						p.SetState(327)
						p.Patch_type()
					}

				case KlangParserPATCHLOAD:
					{
						p.SetState(328)
						p.Match(KlangParserPATCHLOAD)
					}
					{
						p.SetState(329)
						p.String_or_id()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(332)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
		}

	case 3:
		localctx = NewGetKubectlCommandContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(334)
			p.Match(KlangParserKUBECTL)
		}
		{
			p.SetState(335)
			p.Match(KlangParserGET)
		}
		p.SetState(339)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(339)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case KlangParserNAMESPACE:
					{
						p.SetState(336)
						p.Match(KlangParserNAMESPACE)
					}
					{
						p.SetState(337)
						p.Ns()
					}

				case KlangParserID, KlangParserPATH, KlangParserRAW_STRING_LIT, KlangParserSTRING:
					{
						p.SetState(338)
						p.Resource()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(341)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
		}

	case 4:
		localctx = NewDeleteKubectlCommandContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(343)
			p.Match(KlangParserKUBECTL)
		}
		{
			p.SetState(344)
			p.Match(KlangParserDELETE)
		}
		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(348)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case KlangParserNAMESPACE:
					{
						p.SetState(345)
						p.Match(KlangParserNAMESPACE)
					}
					{
						p.SetState(346)
						p.Ns()
					}

				case KlangParserID, KlangParserPATH, KlangParserRAW_STRING_LIT, KlangParserSTRING:
					{
						p.SetState(347)
						p.Resource()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(350)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
		}

	}

	return localctx
}

// IDownload_fnContext is an interface to support dynamic dispatch.
type IDownload_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDownload_fnContext differentiates from other interfaces.
	IsDownload_fnContext()
}

type Download_fnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDownload_fnContext() *Download_fnContext {
	var p = new(Download_fnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_download_fn
	return p
}

func (*Download_fnContext) IsDownload_fnContext() {}

func NewDownload_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Download_fnContext {
	var p = new(Download_fnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_download_fn

	return p
}

func (s *Download_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Download_fnContext) DOWNLOAD() antlr.TerminalNode {
	return s.GetToken(KlangParserDOWNLOAD, 0)
}

func (s *Download_fnContext) OPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserOPAR, 0)
}

func (s *Download_fnContext) AllString_or_id() []IString_or_idContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IString_or_idContext)(nil)).Elem())
	var tst = make([]IString_or_idContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IString_or_idContext)
		}
	}

	return tst
}

func (s *Download_fnContext) String_or_id(i int) IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *Download_fnContext) CPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserCPAR, 0)
}

func (s *Download_fnContext) COMMA() antlr.TerminalNode {
	return s.GetToken(KlangParserCOMMA, 0)
}

func (s *Download_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Download_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Download_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterDownload_fn(s)
	}
}

func (s *Download_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitDownload_fn(s)
	}
}

func (p *KlangParser) Download_fn() (localctx IDownload_fnContext) {
	localctx = NewDownload_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, KlangParserRULE_download_fn)
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
		p.SetState(354)
		p.Match(KlangParserDOWNLOAD)
	}
	{
		p.SetState(355)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(356)
		p.String_or_id()
	}
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KlangParserCOMMA {
		{
			p.SetState(357)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(358)
			p.String_or_id()
		}

	}
	{
		p.SetState(361)
		p.Match(KlangParserCPAR)
	}

	return localctx
}

// IJson_select_fnContext is an interface to support dynamic dispatch.
type IJson_select_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJson_select_fnContext differentiates from other interfaces.
	IsJson_select_fnContext()
}

type Json_select_fnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJson_select_fnContext() *Json_select_fnContext {
	var p = new(Json_select_fnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_json_select_fn
	return p
}

func (*Json_select_fnContext) IsJson_select_fnContext() {}

func NewJson_select_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Json_select_fnContext {
	var p = new(Json_select_fnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_json_select_fn

	return p
}

func (s *Json_select_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Json_select_fnContext) JSONSELECT() antlr.TerminalNode {
	return s.GetToken(KlangParserJSONSELECT, 0)
}

func (s *Json_select_fnContext) OPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserOPAR, 0)
}

func (s *Json_select_fnContext) ID() antlr.TerminalNode {
	return s.GetToken(KlangParserID, 0)
}

func (s *Json_select_fnContext) COMMA() antlr.TerminalNode {
	return s.GetToken(KlangParserCOMMA, 0)
}

func (s *Json_select_fnContext) String_or_id() IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *Json_select_fnContext) CPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserCPAR, 0)
}

func (s *Json_select_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Json_select_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Json_select_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterJson_select_fn(s)
	}
}

func (s *Json_select_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitJson_select_fn(s)
	}
}

func (p *KlangParser) Json_select_fn() (localctx IJson_select_fnContext) {
	localctx = NewJson_select_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, KlangParserRULE_json_select_fn)

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
		p.SetState(363)
		p.Match(KlangParserJSONSELECT)
	}
	{
		p.SetState(364)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(365)
		p.Match(KlangParserID)
	}
	{
		p.SetState(366)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(367)
		p.String_or_id()
	}
	{
		p.SetState(368)
		p.Match(KlangParserCPAR)
	}

	return localctx
}

// IYaml_select_fnContext is an interface to support dynamic dispatch.
type IYaml_select_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsYaml_select_fnContext differentiates from other interfaces.
	IsYaml_select_fnContext()
}

type Yaml_select_fnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYaml_select_fnContext() *Yaml_select_fnContext {
	var p = new(Yaml_select_fnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_yaml_select_fn
	return p
}

func (*Yaml_select_fnContext) IsYaml_select_fnContext() {}

func NewYaml_select_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Yaml_select_fnContext {
	var p = new(Yaml_select_fnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_yaml_select_fn

	return p
}

func (s *Yaml_select_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Yaml_select_fnContext) YAMLSELECT() antlr.TerminalNode {
	return s.GetToken(KlangParserYAMLSELECT, 0)
}

func (s *Yaml_select_fnContext) OPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserOPAR, 0)
}

func (s *Yaml_select_fnContext) ID() antlr.TerminalNode {
	return s.GetToken(KlangParserID, 0)
}

func (s *Yaml_select_fnContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KlangParserCOMMA)
}

func (s *Yaml_select_fnContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KlangParserCOMMA, i)
}

func (s *Yaml_select_fnContext) String_or_id() IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *Yaml_select_fnContext) CPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserCPAR, 0)
}

func (s *Yaml_select_fnContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(KlangParserNUMBER, 0)
}

func (s *Yaml_select_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Yaml_select_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Yaml_select_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterYaml_select_fn(s)
	}
}

func (s *Yaml_select_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitYaml_select_fn(s)
	}
}

func (p *KlangParser) Yaml_select_fn() (localctx IYaml_select_fnContext) {
	localctx = NewYaml_select_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, KlangParserRULE_yaml_select_fn)
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
		p.SetState(370)
		p.Match(KlangParserYAMLSELECT)
	}
	{
		p.SetState(371)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(372)
		p.Match(KlangParserID)
	}
	{
		p.SetState(373)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(374)
		p.String_or_id()
	}
	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KlangParserCOMMA {
		{
			p.SetState(375)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(376)
			p.Match(KlangParserNUMBER)
		}

	}
	{
		p.SetState(379)
		p.Match(KlangParserCPAR)
	}

	return localctx
}

// ILoad_fnContext is an interface to support dynamic dispatch.
type ILoad_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoad_fnContext differentiates from other interfaces.
	IsLoad_fnContext()
}

type Load_fnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoad_fnContext() *Load_fnContext {
	var p = new(Load_fnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_load_fn
	return p
}

func (*Load_fnContext) IsLoad_fnContext() {}

func NewLoad_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Load_fnContext {
	var p = new(Load_fnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_load_fn

	return p
}

func (s *Load_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Load_fnContext) LOAD() antlr.TerminalNode {
	return s.GetToken(KlangParserLOAD, 0)
}

func (s *Load_fnContext) OPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserOPAR, 0)
}

func (s *Load_fnContext) String_or_id() IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *Load_fnContext) CPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserCPAR, 0)
}

func (s *Load_fnContext) COMMA() antlr.TerminalNode {
	return s.GetToken(KlangParserCOMMA, 0)
}

func (s *Load_fnContext) STRING() antlr.TerminalNode {
	return s.GetToken(KlangParserSTRING, 0)
}

func (s *Load_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Load_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Load_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterLoad_fn(s)
	}
}

func (s *Load_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitLoad_fn(s)
	}
}

func (p *KlangParser) Load_fn() (localctx ILoad_fnContext) {
	localctx = NewLoad_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, KlangParserRULE_load_fn)
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
		p.SetState(381)
		p.Match(KlangParserLOAD)
	}
	{
		p.SetState(382)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(383)
		p.String_or_id()
	}
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KlangParserCOMMA {
		{
			p.SetState(384)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(385)
			p.Match(KlangParserSTRING)
		}

	}
	{
		p.SetState(388)
		p.Match(KlangParserCPAR)
	}

	return localctx
}

// IStepInfoContext is an interface to support dynamic dispatch.
type IStepInfoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStepInfoContext differentiates from other interfaces.
	IsStepInfoContext()
}

type StepInfoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStepInfoContext() *StepInfoContext {
	var p = new(StepInfoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_stepInfo
	return p
}

func (*StepInfoContext) IsStepInfoContext() {}

func NewStepInfoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StepInfoContext {
	var p = new(StepInfoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_stepInfo

	return p
}

func (s *StepInfoContext) GetParser() antlr.Parser { return s.parser }

func (s *StepInfoContext) STEPINFO() antlr.TerminalNode {
	return s.GetToken(KlangParserSTEPINFO, 0)
}

func (s *StepInfoContext) STRING() antlr.TerminalNode {
	return s.GetToken(KlangParserSTRING, 0)
}

func (s *StepInfoContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KlangParserSCOL, 0)
}

func (s *StepInfoContext) RAW_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(KlangParserRAW_STRING_LIT, 0)
}

func (s *StepInfoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StepInfoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StepInfoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterStepInfo(s)
	}
}

func (s *StepInfoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitStepInfo(s)
	}
}

func (p *KlangParser) StepInfo() (localctx IStepInfoContext) {
	localctx = NewStepInfoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, KlangParserRULE_stepInfo)

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

	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(390)
			p.Match(KlangParserSTEPINFO)
		}
		{
			p.SetState(391)
			p.Match(KlangParserSTRING)
		}
		{
			p.SetState(392)
			p.Match(KlangParserSCOL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(393)
			p.Match(KlangParserSTEPINFO)
		}
		{
			p.SetState(394)
			p.Match(KlangParserRAW_STRING_LIT)
		}
		{
			p.SetState(395)
			p.Match(KlangParserSCOL)
		}

	}

	return localctx
}

// INsContext is an interface to support dynamic dispatch.
type INsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNsContext differentiates from other interfaces.
	IsNsContext()
}

type NsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNsContext() *NsContext {
	var p = new(NsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_ns
	return p
}

func (*NsContext) IsNsContext() {}

func NewNsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NsContext {
	var p = new(NsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_ns

	return p
}

func (s *NsContext) GetParser() antlr.Parser { return s.parser }

func (s *NsContext) String_or_id() IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *NsContext) PATH() antlr.TerminalNode {
	return s.GetToken(KlangParserPATH, 0)
}

func (s *NsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterNs(s)
	}
}

func (s *NsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitNs(s)
	}
}

func (p *KlangParser) Ns() (localctx INsContext) {
	localctx = NewNsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, KlangParserRULE_ns)

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

	p.SetState(400)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserID, KlangParserRAW_STRING_LIT, KlangParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(398)
			p.String_or_id()
		}

	case KlangParserPATH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(399)
			p.Match(KlangParserPATH)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPatch_typeContext is an interface to support dynamic dispatch.
type IPatch_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPatch_typeContext differentiates from other interfaces.
	IsPatch_typeContext()
}

type Patch_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatch_typeContext() *Patch_typeContext {
	var p = new(Patch_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_patch_type
	return p
}

func (*Patch_typeContext) IsPatch_typeContext() {}

func NewPatch_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Patch_typeContext {
	var p = new(Patch_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_patch_type

	return p
}

func (s *Patch_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Patch_typeContext) PATH() antlr.TerminalNode {
	return s.GetToken(KlangParserPATH, 0)
}

func (s *Patch_typeContext) String_or_id() IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *Patch_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Patch_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Patch_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterPatch_type(s)
	}
}

func (s *Patch_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitPatch_type(s)
	}
}

func (p *KlangParser) Patch_type() (localctx IPatch_typeContext) {
	localctx = NewPatch_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, KlangParserRULE_patch_type)

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

	p.SetState(404)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserPATH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(402)
			p.Match(KlangParserPATH)
		}

	case KlangParserID, KlangParserRAW_STRING_LIT, KlangParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(403)
			p.String_or_id()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IString_or_idContext is an interface to support dynamic dispatch.
type IString_or_idContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_or_idContext differentiates from other interfaces.
	IsString_or_idContext()
}

type String_or_idContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_or_idContext() *String_or_idContext {
	var p = new(String_or_idContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_string_or_id
	return p
}

func (*String_or_idContext) IsString_or_idContext() {}

func NewString_or_idContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_or_idContext {
	var p = new(String_or_idContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_string_or_id

	return p
}

func (s *String_or_idContext) GetParser() antlr.Parser { return s.parser }

func (s *String_or_idContext) ID() antlr.TerminalNode {
	return s.GetToken(KlangParserID, 0)
}

func (s *String_or_idContext) STRING() antlr.TerminalNode {
	return s.GetToken(KlangParserSTRING, 0)
}

func (s *String_or_idContext) RAW_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(KlangParserRAW_STRING_LIT, 0)
}

func (s *String_or_idContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_or_idContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_or_idContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterString_or_id(s)
	}
}

func (s *String_or_idContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitString_or_id(s)
	}
}

func (p *KlangParser) String_or_id() (localctx IString_or_idContext) {
	localctx = NewString_or_idContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, KlangParserRULE_string_or_id)
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
	p.SetState(406)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-63)&-(0x1f+1)) == 0 && ((1<<uint((_la-63)))&((1<<(KlangParserID-63))|(1<<(KlangParserRAW_STRING_LIT-63))|(1<<(KlangParserSTRING-63)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IResourceContext is an interface to support dynamic dispatch.
type IResourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResourceContext differentiates from other interfaces.
	IsResourceContext()
}

type ResourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResourceContext() *ResourceContext {
	var p = new(ResourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_resource
	return p
}

func (*ResourceContext) IsResourceContext() {}

func NewResourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourceContext {
	var p = new(ResourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_resource

	return p
}

func (s *ResourceContext) GetParser() antlr.Parser { return s.parser }

func (s *ResourceContext) PATH() antlr.TerminalNode {
	return s.GetToken(KlangParserPATH, 0)
}

func (s *ResourceContext) String_or_id() IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *ResourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterResource(s)
	}
}

func (s *ResourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitResource(s)
	}
}

func (p *KlangParser) Resource() (localctx IResourceContext) {
	localctx = NewResourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, KlangParserRULE_resource)

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

	p.SetState(410)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserPATH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(408)
			p.Match(KlangParserPATH)
		}

	case KlangParserID, KlangParserRAW_STRING_LIT, KlangParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(409)
			p.String_or_id()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IKubernetes_object_configContext is an interface to support dynamic dispatch.
type IKubernetes_object_configContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKubernetes_object_configContext differentiates from other interfaces.
	IsKubernetes_object_configContext()
}

type Kubernetes_object_configContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKubernetes_object_configContext() *Kubernetes_object_configContext {
	var p = new(Kubernetes_object_configContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_kubernetes_object_config
	return p
}

func (*Kubernetes_object_configContext) IsKubernetes_object_configContext() {}

func NewKubernetes_object_configContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Kubernetes_object_configContext {
	var p = new(Kubernetes_object_configContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_kubernetes_object_config

	return p
}

func (s *Kubernetes_object_configContext) GetParser() antlr.Parser { return s.parser }

func (s *Kubernetes_object_configContext) String_or_id() IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *Kubernetes_object_configContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Kubernetes_object_configContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Kubernetes_object_configContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterKubernetes_object_config(s)
	}
}

func (s *Kubernetes_object_configContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitKubernetes_object_config(s)
	}
}

func (p *KlangParser) Kubernetes_object_config() (localctx IKubernetes_object_configContext) {
	localctx = NewKubernetes_object_configContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, KlangParserRULE_kubernetes_object_config)

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
		p.String_or_id()
	}

	return localctx
}

// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_filter
	return p
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) FILTER() antlr.TerminalNode {
	return s.GetToken(KlangParserFILTER, 0)
}

func (s *FilterContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(KlangParserASSIGN, 0)
}

func (s *FilterContext) String_or_id() IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterFilter(s)
	}
}

func (s *FilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitFilter(s)
	}
}

func (p *KlangParser) Filter() (localctx IFilterContext) {
	localctx = NewFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, KlangParserRULE_filter)

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
		p.SetState(414)
		p.Match(KlangParserFILTER)
	}
	{
		p.SetState(415)
		p.Match(KlangParserASSIGN)
	}
	{
		p.SetState(416)
		p.String_or_id()
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
	p.RuleIndex = KlangParserRULE_pattern
	return p
}

func (*PatternContext) IsPatternContext() {}

func NewPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternContext {
	var p = new(PatternContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_pattern

	return p
}

func (s *PatternContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternContext) PATTERN() antlr.TerminalNode {
	return s.GetToken(KlangParserPATTERN, 0)
}

func (s *PatternContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(KlangParserASSIGN, 0)
}

func (s *PatternContext) String_or_id() IString_or_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_or_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_or_idContext)
}

func (s *PatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterPattern(s)
	}
}

func (s *PatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitPattern(s)
	}
}

func (p *KlangParser) Pattern() (localctx IPatternContext) {
	localctx = NewPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, KlangParserRULE_pattern)

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
		p.SetState(418)
		p.Match(KlangParserPATTERN)
	}
	{
		p.SetState(419)
		p.Match(KlangParserASSIGN)
	}
	{
		p.SetState(420)
		p.String_or_id()
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
	p.RuleIndex = KlangParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_expr

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

type JsonSelectFnContext struct {
	*ExprContext
}

func NewJsonSelectFnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonSelectFnContext {
	var p = new(JsonSelectFnContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *JsonSelectFnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonSelectFnContext) Json_select_fn() IJson_select_fnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJson_select_fnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJson_select_fnContext)
}

func (s *JsonSelectFnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterJsonSelectFn(s)
	}
}

func (s *JsonSelectFnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitJsonSelectFn(s)
	}
}

type ShellScriptContext struct {
	*ExprContext
}

func NewShellScriptContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShellScriptContext {
	var p = new(ShellScriptContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ShellScriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShellScriptContext) Shell_script() IShell_scriptContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShell_scriptContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShell_scriptContext)
}

func (s *ShellScriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterShellScript(s)
	}
}

func (s *ShellScriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitShellScript(s)
	}
}

type YamlSelectFnContext struct {
	*ExprContext
}

func NewYamlSelectFnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *YamlSelectFnContext {
	var p = new(YamlSelectFnContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *YamlSelectFnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *YamlSelectFnContext) Yaml_select_fn() IYaml_select_fnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYaml_select_fnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYaml_select_fnContext)
}

func (s *YamlSelectFnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterYamlSelectFn(s)
	}
}

func (s *YamlSelectFnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitYamlSelectFn(s)
	}
}

type AtomExprContext struct {
	*ExprContext
}

func NewAtomExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomExprContext {
	var p = new(AtomExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AtomExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomExprContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *AtomExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterAtomExpr(s)
	}
}

func (s *AtomExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitAtomExpr(s)
	}
}

type OrExprContext struct {
	*ExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *OrExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrExprContext) OR() antlr.TerminalNode {
	return s.GetToken(KlangParserOR, 0)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

type AdditiveExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewAdditiveExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExprContext {
	var p = new(AdditiveExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AdditiveExprContext) GetOp() antlr.Token { return s.op }

func (s *AdditiveExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AdditiveExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AdditiveExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AdditiveExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(KlangParserPLUS, 0)
}

func (s *AdditiveExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(KlangParserMINUS, 0)
}

func (s *AdditiveExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitAdditiveExpr(s)
	}
}

type RelationalExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewRelationalExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalExprContext {
	var p = new(RelationalExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RelationalExprContext) GetOp() antlr.Token { return s.op }

func (s *RelationalExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *RelationalExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RelationalExprContext) LTEQ() antlr.TerminalNode {
	return s.GetToken(KlangParserLTEQ, 0)
}

func (s *RelationalExprContext) GTEQ() antlr.TerminalNode {
	return s.GetToken(KlangParserGTEQ, 0)
}

func (s *RelationalExprContext) LT() antlr.TerminalNode {
	return s.GetToken(KlangParserLT, 0)
}

func (s *RelationalExprContext) GT() antlr.TerminalNode {
	return s.GetToken(KlangParserGT, 0)
}

func (s *RelationalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterRelationalExpr(s)
	}
}

func (s *RelationalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitRelationalExpr(s)
	}
}

type NotExprContext struct {
	*ExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(KlangParserNOT, 0)
}

func (s *NotExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

type UnaryMinusExprContext struct {
	*ExprContext
}

func NewUnaryMinusExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusExprContext {
	var p = new(UnaryMinusExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *UnaryMinusExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(KlangParserMINUS, 0)
}

func (s *UnaryMinusExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnaryMinusExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterUnaryMinusExpr(s)
	}
}

func (s *UnaryMinusExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitUnaryMinusExpr(s)
	}
}

type MultiplicationExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewMultiplicationExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicationExprContext {
	var p = new(MultiplicationExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MultiplicationExprContext) GetOp() antlr.Token { return s.op }

func (s *MultiplicationExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultiplicationExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MultiplicationExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MultiplicationExprContext) MULT() antlr.TerminalNode {
	return s.GetToken(KlangParserMULT, 0)
}

func (s *MultiplicationExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(KlangParserDIV, 0)
}

func (s *MultiplicationExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(KlangParserMOD, 0)
}

func (s *MultiplicationExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterMultiplicationExpr(s)
	}
}

func (s *MultiplicationExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitMultiplicationExpr(s)
	}
}

type DownloadFnContext struct {
	*ExprContext
}

func NewDownloadFnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DownloadFnContext {
	var p = new(DownloadFnContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DownloadFnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DownloadFnContext) Download_fn() IDownload_fnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDownload_fnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDownload_fnContext)
}

func (s *DownloadFnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterDownloadFn(s)
	}
}

func (s *DownloadFnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitDownloadFn(s)
	}
}

type PowExprContext struct {
	*ExprContext
}

func NewPowExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowExprContext {
	var p = new(PowExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PowExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PowExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PowExprContext) POW() antlr.TerminalNode {
	return s.GetToken(KlangParserPOW, 0)
}

func (s *PowExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterPowExpr(s)
	}
}

func (s *PowExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitPowExpr(s)
	}
}

type KubectlExprContext struct {
	*ExprContext
}

func NewKubectlExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KubectlExprContext {
	var p = new(KubectlExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *KubectlExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KubectlExprContext) Kubectl_command() IKubectl_commandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKubectl_commandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKubectl_commandContext)
}

func (s *KubectlExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterKubectlExpr(s)
	}
}

func (s *KubectlExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitKubectlExpr(s)
	}
}

type EqualityExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewEqualityExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExprContext {
	var p = new(EqualityExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EqualityExprContext) GetOp() antlr.Token { return s.op }

func (s *EqualityExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualityExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *EqualityExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqualityExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(KlangParserEQ, 0)
}

func (s *EqualityExprContext) NEQ() antlr.TerminalNode {
	return s.GetToken(KlangParserNEQ, 0)
}

func (s *EqualityExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterEqualityExpr(s)
	}
}

func (s *EqualityExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitEqualityExpr(s)
	}
}

type AndExprContext struct {
	*ExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AndExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(KlangParserAND, 0)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (p *KlangParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *KlangParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 66
	p.EnterRecursionRule(localctx, 66, KlangParserRULE_expr, _p)
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
	p.SetState(433)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserMINUS:
		localctx = NewUnaryMinusExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(423)
			p.Match(KlangParserMINUS)
		}
		{
			p.SetState(424)
			p.expr(14)
		}

	case KlangParserNOT:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(425)
			p.Match(KlangParserNOT)
		}
		{
			p.SetState(426)
			p.expr(13)
		}

	case KlangParserKUBECTL:
		localctx = NewKubectlExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(427)
			p.Kubectl_command()
		}

	case KlangParserJSONSELECT:
		localctx = NewJsonSelectFnContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(428)
			p.Json_select_fn()
		}

	case KlangParserYAMLSELECT:
		localctx = NewYamlSelectFnContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(429)
			p.Yaml_select_fn()
		}

	case KlangParserSHELLSCRIPT:
		localctx = NewShellScriptContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(430)
			p.Shell_script()
		}

	case KlangParserDOWNLOAD:
		localctx = NewDownloadFnContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(431)
			p.Download_fn()
		}

	case KlangParserT__1, KlangParserT__3, KlangParserOPAR, KlangParserOBRACE, KlangParserTRUE, KlangParserFALSE, KlangParserNIL, KlangParserID, KlangParserNUMBER, KlangParserRAW_STRING_LIT, KlangParserSTRING:
		localctx = NewAtomExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(432)
			p.Atom()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(456)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(435)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(436)
					p.Match(KlangParserPOW)
				}
				{
					p.SetState(437)
					p.expr(15)
				}

			case 2:
				localctx = NewMultiplicationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(438)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				p.SetState(439)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*MultiplicationExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<KlangParserMULT)|(1<<KlangParserDIV)|(1<<KlangParserMOD))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*MultiplicationExprContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(440)
					p.expr(13)
				}

			case 3:
				localctx = NewAdditiveExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(441)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				p.SetState(442)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*AdditiveExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == KlangParserPLUS || _la == KlangParserMINUS) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*AdditiveExprContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(443)
					p.expr(12)
				}

			case 4:
				localctx = NewRelationalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(444)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				p.SetState(445)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*RelationalExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<KlangParserGT)|(1<<KlangParserLT)|(1<<KlangParserGTEQ)|(1<<KlangParserLTEQ))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*RelationalExprContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(446)
					p.expr(11)
				}

			case 5:
				localctx = NewEqualityExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(447)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				p.SetState(448)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*EqualityExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == KlangParserEQ || _la == KlangParserNEQ) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*EqualityExprContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(449)
					p.expr(10)
				}

			case 6:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(450)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(451)
					p.Match(KlangParserAND)
				}
				{
					p.SetState(452)
					p.expr(9)
				}

			case 7:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(453)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(454)
					p.Match(KlangParserOR)
				}
				{
					p.SetState(455)
					p.expr(8)
				}

			}

		}
		p.SetState(460)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) CopyFrom(ctx *AtomContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParExprContext struct {
	*AtomContext
}

func NewParExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParExprContext {
	var p = new(ParExprContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *ParExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExprContext) OPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserOPAR, 0)
}

func (s *ParExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParExprContext) CPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserCPAR, 0)
}

func (s *ParExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterParExpr(s)
	}
}

func (s *ParExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitParExpr(s)
	}
}

type BooleanAtomContext struct {
	*AtomContext
}

func NewBooleanAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanAtomContext {
	var p = new(BooleanAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *BooleanAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanAtomContext) TRUE() antlr.TerminalNode {
	return s.GetToken(KlangParserTRUE, 0)
}

func (s *BooleanAtomContext) FALSE() antlr.TerminalNode {
	return s.GetToken(KlangParserFALSE, 0)
}

func (s *BooleanAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterBooleanAtom(s)
	}
}

func (s *BooleanAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitBooleanAtom(s)
	}
}

type IdAtomContext struct {
	*AtomContext
}

func NewIdAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdAtomContext {
	var p = new(IdAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *IdAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdAtomContext) ID() antlr.TerminalNode {
	return s.GetToken(KlangParserID, 0)
}

func (s *IdAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterIdAtom(s)
	}
}

func (s *IdAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitIdAtom(s)
	}
}

type StringAtomContext struct {
	*AtomContext
}

func NewStringAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringAtomContext {
	var p = new(StringAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *StringAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringAtomContext) STRING() antlr.TerminalNode {
	return s.GetToken(KlangParserSTRING, 0)
}

func (s *StringAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterStringAtom(s)
	}
}

func (s *StringAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitStringAtom(s)
	}
}

type RawStringAtomContext struct {
	*AtomContext
}

func NewRawStringAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RawStringAtomContext {
	var p = new(RawStringAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *RawStringAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RawStringAtomContext) RAW_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(KlangParserRAW_STRING_LIT, 0)
}

func (s *RawStringAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterRawStringAtom(s)
	}
}

func (s *RawStringAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitRawStringAtom(s)
	}
}

type JsonAtomContext struct {
	*AtomContext
}

func NewJsonAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonAtomContext {
	var p = new(JsonAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *JsonAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonAtomContext) Json() IJsonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

func (s *JsonAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterJsonAtom(s)
	}
}

func (s *JsonAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitJsonAtom(s)
	}
}

type NilAtomContext struct {
	*AtomContext
}

func NewNilAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilAtomContext {
	var p = new(NilAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *NilAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilAtomContext) NIL() antlr.TerminalNode {
	return s.GetToken(KlangParserNIL, 0)
}

func (s *NilAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterNilAtom(s)
	}
}

func (s *NilAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitNilAtom(s)
	}
}

type NumberAtomContext struct {
	*AtomContext
}

func NewNumberAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberAtomContext {
	var p = new(NumberAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *NumberAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberAtomContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(KlangParserNUMBER, 0)
}

func (s *NumberAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterNumberAtom(s)
	}
}

func (s *NumberAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitNumberAtom(s)
	}
}

func (p *KlangParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, KlangParserRULE_atom)
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

	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(461)
			p.Match(KlangParserOPAR)
		}
		{
			p.SetState(462)
			p.expr(0)
		}
		{
			p.SetState(463)
			p.Match(KlangParserCPAR)
		}

	case 2:
		localctx = NewNumberAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(465)
			p.Match(KlangParserNUMBER)
		}

	case 3:
		localctx = NewBooleanAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(466)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KlangParserTRUE || _la == KlangParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case 4:
		localctx = NewRawStringAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(467)
			p.Match(KlangParserRAW_STRING_LIT)
		}

	case 5:
		localctx = NewIdAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(468)
			p.Match(KlangParserID)
		}

	case 6:
		localctx = NewStringAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(469)
			p.Match(KlangParserSTRING)
		}

	case 7:
		localctx = NewJsonAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(470)
			p.Json()
		}

	case 8:
		localctx = NewNilAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(471)
			p.Match(KlangParserNIL)
		}

	}

	return localctx
}

// IJsonContext is an interface to support dynamic dispatch.
type IJsonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonContext differentiates from other interfaces.
	IsJsonContext()
}

type JsonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonContext() *JsonContext {
	var p = new(JsonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_json
	return p
}

func (*JsonContext) IsJsonContext() {}

func NewJsonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonContext {
	var p = new(JsonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_json

	return p
}

func (s *JsonContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *JsonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterJson(s)
	}
}

func (s *JsonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitJson(s)
	}
}

func (p *KlangParser) Json() (localctx IJsonContext) {
	localctx = NewJsonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, KlangParserRULE_json)

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
		p.SetState(474)
		p.Value()
	}

	return localctx
}

// IObjContext is an interface to support dynamic dispatch.
type IObjContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjContext differentiates from other interfaces.
	IsObjContext()
}

type ObjContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjContext() *ObjContext {
	var p = new(ObjContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_obj
	return p
}

func (*ObjContext) IsObjContext() {}

func NewObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjContext {
	var p = new(ObjContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_obj

	return p
}

func (s *ObjContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjContext) AllPair() []IPairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPairContext)(nil)).Elem())
	var tst = make([]IPairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPairContext)
		}
	}

	return tst
}

func (s *ObjContext) Pair(i int) IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *ObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterObj(s)
	}
}

func (s *ObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitObj(s)
	}
}

func (p *KlangParser) Obj() (localctx IObjContext) {
	localctx = NewObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, KlangParserRULE_obj)
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

	p.SetState(489)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(476)
			p.Match(KlangParserOBRACE)
		}
		{
			p.SetState(477)
			p.Pair()
		}
		p.SetState(482)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == KlangParserCOMMA {
			{
				p.SetState(478)
				p.Match(KlangParserCOMMA)
			}
			{
				p.SetState(479)
				p.Pair()
			}

			p.SetState(484)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(485)
			p.Match(KlangParserCBRACE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(487)
			p.Match(KlangParserOBRACE)
		}
		{
			p.SetState(488)
			p.Match(KlangParserCBRACE)
		}

	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) STRING() antlr.TerminalNode {
	return s.GetToken(KlangParserSTRING, 0)
}

func (s *PairContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitPair(s)
	}
}

func (p *KlangParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, KlangParserRULE_pair)

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
		p.SetState(491)
		p.Match(KlangParserSTRING)
	}
	{
		p.SetState(492)
		p.Match(KlangParserT__0)
	}
	{
		p.SetState(493)
		p.Value()
	}

	return localctx
}

// IArrContext is an interface to support dynamic dispatch.
type IArrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrContext differentiates from other interfaces.
	IsArrContext()
}

type ArrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrContext() *ArrContext {
	var p = new(ArrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlangParserRULE_arr
	return p
}

func (*ArrContext) IsArrContext() {}

func NewArrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrContext {
	var p = new(ArrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_arr

	return p
}

func (s *ArrContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ArrContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterArr(s)
	}
}

func (s *ArrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitArr(s)
	}
}

func (p *KlangParser) Arr() (localctx IArrContext) {
	localctx = NewArrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, KlangParserRULE_arr)
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

	p.SetState(508)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(495)
			p.Match(KlangParserT__1)
		}
		{
			p.SetState(496)
			p.Value()
		}
		p.SetState(501)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == KlangParserCOMMA {
			{
				p.SetState(497)
				p.Match(KlangParserCOMMA)
			}
			{
				p.SetState(498)
				p.Value()
			}

			p.SetState(503)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(504)
			p.Match(KlangParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(506)
			p.Match(KlangParserT__1)
		}
		{
			p.SetState(507)
			p.Match(KlangParserT__2)
		}

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
	p.RuleIndex = KlangParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlangParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(KlangParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(KlangParserNUMBER, 0)
}

func (s *ValueContext) Obj() IObjContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjContext)
}

func (s *ValueContext) Arr() IArrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlangListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *KlangParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, KlangParserRULE_value)

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

	p.SetState(517)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(510)
			p.Match(KlangParserSTRING)
		}

	case KlangParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(511)
			p.Match(KlangParserNUMBER)
		}

	case KlangParserOBRACE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(512)
			p.Obj()
		}

	case KlangParserT__1:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(513)
			p.Arr()
		}

	case KlangParserTRUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(514)
			p.Match(KlangParserTRUE)
		}

	case KlangParserFALSE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(515)
			p.Match(KlangParserFALSE)
		}

	case KlangParserT__3:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(516)
			p.Match(KlangParserT__3)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *KlangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 33:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *KlangParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
