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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 64, 379,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 3, 2,
	3, 2, 3, 2, 3, 3, 7, 3, 71, 10, 3, 12, 3, 14, 3, 74, 11, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 87, 10, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 99, 10,
	5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 131, 10, 9, 3, 9, 3, 9, 3,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 143, 10, 10,
	3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 7, 12, 159, 10, 12, 12, 12, 14, 12, 162, 11, 12,
	3, 12, 3, 12, 5, 12, 166, 10, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 5, 14, 176, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 6,
	17, 193, 10, 17, 13, 17, 14, 17, 194, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 6, 17, 206, 10, 17, 13, 17, 14, 17, 207, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 6, 17, 215, 10, 17, 13, 17, 14, 17, 216,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 6, 17, 224, 10, 17, 13, 17, 14, 17,
	225, 5, 17, 228, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 235,
	10, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 253, 10, 20, 3,
	20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 262, 10, 21, 3, 21,
	3, 21, 3, 22, 3, 22, 5, 22, 268, 10, 22, 3, 23, 3, 23, 5, 23, 272, 10,
	23, 3, 24, 3, 24, 3, 25, 3, 25, 5, 25, 278, 10, 25, 3, 26, 3, 26, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5,
	27, 293, 10, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 7, 27, 316, 10, 27, 12, 27, 14, 27, 319, 11, 27, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5,
	28, 332, 10, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 7, 30, 340,
	10, 30, 12, 30, 14, 30, 343, 11, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30,
	349, 10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 7,
	32, 359, 10, 32, 12, 32, 14, 32, 362, 11, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	5, 32, 368, 10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 5,
	33, 377, 10, 33, 3, 33, 2, 3, 52, 34, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
	58, 60, 62, 64, 2, 8, 4, 2, 57, 57, 60, 61, 3, 2, 17, 19, 3, 2, 15, 16,
	3, 2, 11, 14, 3, 2, 9, 10, 3, 2, 29, 30, 2, 413, 2, 66, 3, 2, 2, 2, 4,
	72, 3, 2, 2, 2, 6, 86, 3, 2, 2, 2, 8, 98, 3, 2, 2, 2, 10, 100, 3, 2, 2,
	2, 12, 103, 3, 2, 2, 2, 14, 113, 3, 2, 2, 2, 16, 121, 3, 2, 2, 2, 18, 135,
	3, 2, 2, 2, 20, 147, 3, 2, 2, 2, 22, 153, 3, 2, 2, 2, 24, 167, 3, 2, 2,
	2, 26, 175, 3, 2, 2, 2, 28, 177, 3, 2, 2, 2, 30, 181, 3, 2, 2, 2, 32, 227,
	3, 2, 2, 2, 34, 229, 3, 2, 2, 2, 36, 238, 3, 2, 2, 2, 38, 245, 3, 2, 2,
	2, 40, 256, 3, 2, 2, 2, 42, 267, 3, 2, 2, 2, 44, 271, 3, 2, 2, 2, 46, 273,
	3, 2, 2, 2, 48, 277, 3, 2, 2, 2, 50, 279, 3, 2, 2, 2, 52, 292, 3, 2, 2,
	2, 54, 331, 3, 2, 2, 2, 56, 333, 3, 2, 2, 2, 58, 348, 3, 2, 2, 2, 60, 350,
	3, 2, 2, 2, 62, 367, 3, 2, 2, 2, 64, 376, 3, 2, 2, 2, 66, 67, 5, 4, 3,
	2, 67, 68, 7, 2, 2, 3, 68, 3, 3, 2, 2, 2, 69, 71, 5, 6, 4, 2, 70, 69, 3,
	2, 2, 2, 71, 74, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73,
	5, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 75, 87, 5, 8, 5, 2, 76, 87, 5, 12, 7,
	2, 77, 87, 5, 14, 8, 2, 78, 87, 5, 16, 9, 2, 79, 87, 5, 18, 10, 2, 80,
	87, 5, 22, 12, 2, 81, 87, 5, 28, 15, 2, 82, 87, 5, 20, 11, 2, 83, 87, 5,
	30, 16, 2, 84, 85, 7, 64, 2, 2, 85, 87, 8, 4, 1, 2, 86, 75, 3, 2, 2, 2,
	86, 76, 3, 2, 2, 2, 86, 77, 3, 2, 2, 2, 86, 78, 3, 2, 2, 2, 86, 79, 3,
	2, 2, 2, 86, 80, 3, 2, 2, 2, 86, 81, 3, 2, 2, 2, 86, 82, 3, 2, 2, 2, 86,
	83, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 87, 7, 3, 2, 2, 2, 88, 89, 7, 57, 2,
	2, 89, 90, 7, 23, 2, 2, 90, 91, 5, 52, 27, 2, 91, 92, 7, 22, 2, 2, 92,
	99, 3, 2, 2, 2, 93, 94, 7, 57, 2, 2, 94, 95, 7, 23, 2, 2, 95, 96, 5, 40,
	21, 2, 96, 97, 7, 22, 2, 2, 97, 99, 3, 2, 2, 2, 98, 88, 3, 2, 2, 2, 98,
	93, 3, 2, 2, 2, 99, 9, 3, 2, 2, 2, 100, 101, 7, 54, 2, 2, 101, 102, 5,
	46, 24, 2, 102, 11, 3, 2, 2, 2, 103, 104, 7, 49, 2, 2, 104, 105, 7, 24,
	2, 2, 105, 106, 7, 57, 2, 2, 106, 107, 7, 28, 2, 2, 107, 108, 5, 46, 24,
	2, 108, 109, 7, 28, 2, 2, 109, 110, 5, 52, 27, 2, 110, 111, 7, 25, 2, 2,
	111, 112, 7, 22, 2, 2, 112, 13, 3, 2, 2, 2, 113, 114, 7, 50, 2, 2, 114,
	115, 7, 24, 2, 2, 115, 116, 7, 57, 2, 2, 116, 117, 7, 28, 2, 2, 117, 118,
	5, 46, 24, 2, 118, 119, 7, 25, 2, 2, 119, 120, 7, 22, 2, 2, 120, 15, 3,
	2, 2, 2, 121, 122, 7, 52, 2, 2, 122, 123, 7, 24, 2, 2, 123, 124, 7, 57,
	2, 2, 124, 125, 7, 28, 2, 2, 125, 126, 5, 46, 24, 2, 126, 127, 7, 28, 2,
	2, 127, 130, 5, 52, 27, 2, 128, 129, 7, 28, 2, 2, 129, 131, 7, 58, 2, 2,
	130, 128, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132,
	133, 7, 25, 2, 2, 133, 134, 7, 22, 2, 2, 134, 17, 3, 2, 2, 2, 135, 136,
	7, 53, 2, 2, 136, 137, 7, 24, 2, 2, 137, 138, 7, 57, 2, 2, 138, 139, 7,
	28, 2, 2, 139, 142, 5, 46, 24, 2, 140, 141, 7, 28, 2, 2, 141, 143, 7, 58,
	2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2,
	144, 145, 7, 25, 2, 2, 145, 146, 7, 22, 2, 2, 146, 19, 3, 2, 2, 2, 147,
	148, 7, 56, 2, 2, 148, 149, 7, 24, 2, 2, 149, 150, 7, 58, 2, 2, 150, 151,
	7, 25, 2, 2, 151, 152, 7, 22, 2, 2, 152, 21, 3, 2, 2, 2, 153, 154, 7, 32,
	2, 2, 154, 160, 5, 24, 13, 2, 155, 156, 7, 33, 2, 2, 156, 157, 7, 32, 2,
	2, 157, 159, 5, 24, 13, 2, 158, 155, 3, 2, 2, 2, 159, 162, 3, 2, 2, 2,
	160, 158, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 165, 3, 2, 2, 2, 162,
	160, 3, 2, 2, 2, 163, 164, 7, 33, 2, 2, 164, 166, 5, 26, 14, 2, 165, 163,
	3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 23, 3, 2, 2, 2, 167, 168, 5, 52,
	27, 2, 168, 169, 5, 26, 14, 2, 169, 25, 3, 2, 2, 2, 170, 171, 7, 26, 2,
	2, 171, 172, 5, 4, 3, 2, 172, 173, 7, 27, 2, 2, 173, 176, 3, 2, 2, 2, 174,
	176, 5, 6, 4, 2, 175, 170, 3, 2, 2, 2, 175, 174, 3, 2, 2, 2, 176, 27, 3,
	2, 2, 2, 177, 178, 7, 34, 2, 2, 178, 179, 5, 52, 27, 2, 179, 180, 5, 26,
	14, 2, 180, 29, 3, 2, 2, 2, 181, 182, 7, 35, 2, 2, 182, 183, 5, 52, 27,
	2, 183, 184, 7, 22, 2, 2, 184, 31, 3, 2, 2, 2, 185, 186, 7, 36, 2, 2, 186,
	192, 7, 37, 2, 2, 187, 188, 7, 42, 2, 2, 188, 193, 5, 42, 22, 2, 189, 193,
	5, 46, 24, 2, 190, 191, 7, 45, 2, 2, 191, 193, 5, 50, 26, 2, 192, 187,
	3, 2, 2, 2, 192, 189, 3, 2, 2, 2, 192, 190, 3, 2, 2, 2, 193, 194, 3, 2,
	2, 2, 194, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 228, 3, 2, 2, 2,
	196, 197, 7, 36, 2, 2, 197, 205, 7, 38, 2, 2, 198, 199, 7, 42, 2, 2, 199,
	206, 5, 42, 22, 2, 200, 206, 5, 48, 25, 2, 201, 202, 7, 43, 2, 2, 202,
	206, 5, 44, 23, 2, 203, 204, 7, 44, 2, 2, 204, 206, 5, 46, 24, 2, 205,
	198, 3, 2, 2, 2, 205, 200, 3, 2, 2, 2, 205, 201, 3, 2, 2, 2, 205, 203,
	3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 205, 3, 2, 2, 2, 207, 208, 3, 2,
	2, 2, 208, 228, 3, 2, 2, 2, 209, 210, 7, 36, 2, 2, 210, 214, 7, 39, 2,
	2, 211, 212, 7, 42, 2, 2, 212, 215, 5, 42, 22, 2, 213, 215, 5, 48, 25,
	2, 214, 211, 3, 2, 2, 2, 214, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216,
	214, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 228, 3, 2, 2, 2, 218, 219,
	7, 36, 2, 2, 219, 223, 7, 41, 2, 2, 220, 221, 7, 42, 2, 2, 221, 224, 5,
	42, 22, 2, 222, 224, 5, 48, 25, 2, 223, 220, 3, 2, 2, 2, 223, 222, 3, 2,
	2, 2, 224, 225, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2,
	226, 228, 3, 2, 2, 2, 227, 185, 3, 2, 2, 2, 227, 196, 3, 2, 2, 2, 227,
	209, 3, 2, 2, 2, 227, 218, 3, 2, 2, 2, 228, 33, 3, 2, 2, 2, 229, 230, 7,
	55, 2, 2, 230, 231, 7, 24, 2, 2, 231, 234, 5, 46, 24, 2, 232, 233, 7, 28,
	2, 2, 233, 235, 5, 46, 24, 2, 234, 232, 3, 2, 2, 2, 234, 235, 3, 2, 2,
	2, 235, 236, 3, 2, 2, 2, 236, 237, 7, 25, 2, 2, 237, 35, 3, 2, 2, 2, 238,
	239, 7, 48, 2, 2, 239, 240, 7, 24, 2, 2, 240, 241, 7, 57, 2, 2, 241, 242,
	7, 28, 2, 2, 242, 243, 5, 46, 24, 2, 243, 244, 7, 25, 2, 2, 244, 37, 3,
	2, 2, 2, 245, 246, 7, 51, 2, 2, 246, 247, 7, 24, 2, 2, 247, 248, 7, 57,
	2, 2, 248, 249, 7, 28, 2, 2, 249, 252, 5, 46, 24, 2, 250, 251, 7, 28, 2,
	2, 251, 253, 7, 58, 2, 2, 252, 250, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253,
	254, 3, 2, 2, 2, 254, 255, 7, 25, 2, 2, 255, 39, 3, 2, 2, 2, 256, 257,
	7, 47, 2, 2, 257, 258, 7, 24, 2, 2, 258, 261, 5, 46, 24, 2, 259, 260, 7,
	28, 2, 2, 260, 262, 7, 61, 2, 2, 261, 259, 3, 2, 2, 2, 261, 262, 3, 2,
	2, 2, 262, 263, 3, 2, 2, 2, 263, 264, 7, 25, 2, 2, 264, 41, 3, 2, 2, 2,
	265, 268, 5, 46, 24, 2, 266, 268, 7, 59, 2, 2, 267, 265, 3, 2, 2, 2, 267,
	266, 3, 2, 2, 2, 268, 43, 3, 2, 2, 2, 269, 272, 7, 59, 2, 2, 270, 272,
	5, 46, 24, 2, 271, 269, 3, 2, 2, 2, 271, 270, 3, 2, 2, 2, 272, 45, 3, 2,
	2, 2, 273, 274, 9, 2, 2, 2, 274, 47, 3, 2, 2, 2, 275, 278, 7, 59, 2, 2,
	276, 278, 5, 46, 24, 2, 277, 275, 3, 2, 2, 2, 277, 276, 3, 2, 2, 2, 278,
	49, 3, 2, 2, 2, 279, 280, 5, 46, 24, 2, 280, 51, 3, 2, 2, 2, 281, 282,
	8, 27, 1, 2, 282, 283, 7, 16, 2, 2, 283, 293, 5, 52, 27, 16, 284, 285,
	7, 21, 2, 2, 285, 293, 5, 52, 27, 15, 286, 293, 5, 32, 17, 2, 287, 293,
	5, 36, 19, 2, 288, 293, 5, 38, 20, 2, 289, 293, 5, 10, 6, 2, 290, 293,
	5, 34, 18, 2, 291, 293, 5, 54, 28, 2, 292, 281, 3, 2, 2, 2, 292, 284, 3,
	2, 2, 2, 292, 286, 3, 2, 2, 2, 292, 287, 3, 2, 2, 2, 292, 288, 3, 2, 2,
	2, 292, 289, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2, 292, 291, 3, 2, 2, 2, 293,
	317, 3, 2, 2, 2, 294, 295, 12, 17, 2, 2, 295, 296, 7, 20, 2, 2, 296, 316,
	5, 52, 27, 17, 297, 298, 12, 14, 2, 2, 298, 299, 9, 3, 2, 2, 299, 316,
	5, 52, 27, 15, 300, 301, 12, 13, 2, 2, 301, 302, 9, 4, 2, 2, 302, 316,
	5, 52, 27, 14, 303, 304, 12, 12, 2, 2, 304, 305, 9, 5, 2, 2, 305, 316,
	5, 52, 27, 13, 306, 307, 12, 11, 2, 2, 307, 308, 9, 6, 2, 2, 308, 316,
	5, 52, 27, 12, 309, 310, 12, 10, 2, 2, 310, 311, 7, 8, 2, 2, 311, 316,
	5, 52, 27, 11, 312, 313, 12, 9, 2, 2, 313, 314, 7, 7, 2, 2, 314, 316, 5,
	52, 27, 10, 315, 294, 3, 2, 2, 2, 315, 297, 3, 2, 2, 2, 315, 300, 3, 2,
	2, 2, 315, 303, 3, 2, 2, 2, 315, 306, 3, 2, 2, 2, 315, 309, 3, 2, 2, 2,
	315, 312, 3, 2, 2, 2, 316, 319, 3, 2, 2, 2, 317, 315, 3, 2, 2, 2, 317,
	318, 3, 2, 2, 2, 318, 53, 3, 2, 2, 2, 319, 317, 3, 2, 2, 2, 320, 321, 7,
	24, 2, 2, 321, 322, 5, 52, 27, 2, 322, 323, 7, 25, 2, 2, 323, 332, 3, 2,
	2, 2, 324, 332, 7, 58, 2, 2, 325, 332, 9, 7, 2, 2, 326, 332, 7, 60, 2,
	2, 327, 332, 7, 57, 2, 2, 328, 332, 7, 61, 2, 2, 329, 332, 5, 56, 29, 2,
	330, 332, 7, 31, 2, 2, 331, 320, 3, 2, 2, 2, 331, 324, 3, 2, 2, 2, 331,
	325, 3, 2, 2, 2, 331, 326, 3, 2, 2, 2, 331, 327, 3, 2, 2, 2, 331, 328,
	3, 2, 2, 2, 331, 329, 3, 2, 2, 2, 331, 330, 3, 2, 2, 2, 332, 55, 3, 2,
	2, 2, 333, 334, 5, 64, 33, 2, 334, 57, 3, 2, 2, 2, 335, 336, 7, 26, 2,
	2, 336, 341, 5, 60, 31, 2, 337, 338, 7, 28, 2, 2, 338, 340, 5, 60, 31,
	2, 339, 337, 3, 2, 2, 2, 340, 343, 3, 2, 2, 2, 341, 339, 3, 2, 2, 2, 341,
	342, 3, 2, 2, 2, 342, 344, 3, 2, 2, 2, 343, 341, 3, 2, 2, 2, 344, 345,
	7, 27, 2, 2, 345, 349, 3, 2, 2, 2, 346, 347, 7, 26, 2, 2, 347, 349, 7,
	27, 2, 2, 348, 335, 3, 2, 2, 2, 348, 346, 3, 2, 2, 2, 349, 59, 3, 2, 2,
	2, 350, 351, 7, 61, 2, 2, 351, 352, 7, 3, 2, 2, 352, 353, 5, 64, 33, 2,
	353, 61, 3, 2, 2, 2, 354, 355, 7, 4, 2, 2, 355, 360, 5, 64, 33, 2, 356,
	357, 7, 28, 2, 2, 357, 359, 5, 64, 33, 2, 358, 356, 3, 2, 2, 2, 359, 362,
	3, 2, 2, 2, 360, 358, 3, 2, 2, 2, 360, 361, 3, 2, 2, 2, 361, 363, 3, 2,
	2, 2, 362, 360, 3, 2, 2, 2, 363, 364, 7, 5, 2, 2, 364, 368, 3, 2, 2, 2,
	365, 366, 7, 4, 2, 2, 366, 368, 7, 5, 2, 2, 367, 354, 3, 2, 2, 2, 367,
	365, 3, 2, 2, 2, 368, 63, 3, 2, 2, 2, 369, 377, 7, 61, 2, 2, 370, 377,
	7, 58, 2, 2, 371, 377, 5, 58, 30, 2, 372, 377, 5, 62, 32, 2, 373, 377,
	7, 29, 2, 2, 374, 377, 7, 30, 2, 2, 375, 377, 7, 6, 2, 2, 376, 369, 3,
	2, 2, 2, 376, 370, 3, 2, 2, 2, 376, 371, 3, 2, 2, 2, 376, 372, 3, 2, 2,
	2, 376, 373, 3, 2, 2, 2, 376, 374, 3, 2, 2, 2, 376, 375, 3, 2, 2, 2, 377,
	65, 3, 2, 2, 2, 34, 72, 86, 98, 130, 142, 160, 165, 175, 192, 194, 205,
	207, 214, 216, 223, 225, 227, 234, 252, 261, 267, 271, 277, 292, 315, 317,
	331, 341, 348, 360, 367, 376,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':'", "'['", "']'", "'null'", "'||'", "'&&'", "'=='", "'!='", "'>'",
	"'<'", "'>='", "'<='", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'!'",
	"';'", "'='", "'('", "')'", "'{'", "'}'", "','", "'true'", "'false'", "'nil'",
	"'if'", "'else'", "'while'", "'log'", "'kubectl'", "'apply'", "'patch'",
	"'get'", "'replace'", "'delete'", "'-n'", "'--type'", "'-p'", "'-u'", "'-jsonpath'",
	"'load'", "'jsonSelect'", "'jsonEdit'", "'jsonDelete'", "'yamlSelect'",
	"'yamlEdit'", "'yamlDelete'", "'shellScript'", "'download'", "'sleep'",
}
var symbolicNames = []string{
	"", "", "", "", "", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ",
	"PLUS", "MINUS", "MULT", "DIV", "MOD", "POW", "NOT", "SCOL", "ASSIGN",
	"OPAR", "CPAR", "OBRACE", "CBRACE", "COMMA", "TRUE", "FALSE", "NIL", "IF",
	"ELSE", "WHILE", "LOG", "KUBECTL", "APPLY", "PATCH", "GET", "REPLACE",
	"DELETE", "NAMESPACE", "PATCHTYPE", "PATCHLOAD", "UPDATELOAD", "JSONPATH",
	"LOAD", "JSONSELECT", "JSONEDIT", "JSONDELETE", "YAMLSELECT", "YAMLEDIT",
	"YAMLDELETE", "SHELLSCRIPT", "DOWNLOAD", "SLEEP", "ID", "NUMBER", "PATH",
	"RAW_STRING_LIT", "STRING", "COMMENT", "SPACE", "OTHER",
}

var ruleNames = []string{
	"parse", "block", "stat", "assignment", "shell_script", "json_edit_fn",
	"json_delete_fn", "yaml_edit_fn", "yaml_delete_fn", "sleep_fn", "if_stat",
	"condition_block", "stat_block", "while_stat", "log", "kubectl_command",
	"download_fn", "json_select_fn", "yaml_select_fn", "load_fn", "ns", "patch_type",
	"string_or_id", "resource", "kubernetes_object_config", "expr", "atom",
	"json", "obj", "pair", "arr", "value",
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
	KlangParserJSONSELECT     = 46
	KlangParserJSONEDIT       = 47
	KlangParserJSONDELETE     = 48
	KlangParserYAMLSELECT     = 49
	KlangParserYAMLEDIT       = 50
	KlangParserYAMLDELETE     = 51
	KlangParserSHELLSCRIPT    = 52
	KlangParserDOWNLOAD       = 53
	KlangParserSLEEP          = 54
	KlangParserID             = 55
	KlangParserNUMBER         = 56
	KlangParserPATH           = 57
	KlangParserRAW_STRING_LIT = 58
	KlangParserSTRING         = 59
	KlangParserCOMMENT        = 60
	KlangParserSPACE          = 61
	KlangParserOTHER          = 62
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
	KlangParserRULE_sleep_fn                 = 9
	KlangParserRULE_if_stat                  = 10
	KlangParserRULE_condition_block          = 11
	KlangParserRULE_stat_block               = 12
	KlangParserRULE_while_stat               = 13
	KlangParserRULE_log                      = 14
	KlangParserRULE_kubectl_command          = 15
	KlangParserRULE_download_fn              = 16
	KlangParserRULE_json_select_fn           = 17
	KlangParserRULE_yaml_select_fn           = 18
	KlangParserRULE_load_fn                  = 19
	KlangParserRULE_ns                       = 20
	KlangParserRULE_patch_type               = 21
	KlangParserRULE_string_or_id             = 22
	KlangParserRULE_resource                 = 23
	KlangParserRULE_kubernetes_object_config = 24
	KlangParserRULE_expr                     = 25
	KlangParserRULE_atom                     = 26
	KlangParserRULE_json                     = 27
	KlangParserRULE_obj                      = 28
	KlangParserRULE_pair                     = 29
	KlangParserRULE_arr                      = 30
	KlangParserRULE_value                    = 31
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
		p.SetState(64)
		p.Block()
	}
	{
		p.SetState(65)
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
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KlangParserIF || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(KlangParserWHILE-32))|(1<<(KlangParserLOG-32))|(1<<(KlangParserJSONEDIT-32))|(1<<(KlangParserJSONDELETE-32))|(1<<(KlangParserYAMLEDIT-32))|(1<<(KlangParserYAMLDELETE-32))|(1<<(KlangParserSLEEP-32))|(1<<(KlangParserID-32))|(1<<(KlangParserOTHER-32)))) != 0) {
		{
			p.SetState(67)
			p.Stat()
		}

		p.SetState(72)
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

	p.SetState(84)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.Assignment()
		}

	case KlangParserJSONEDIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)
			p.Json_edit_fn()
		}

	case KlangParserJSONDELETE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(75)
			p.Json_delete_fn()
		}

	case KlangParserYAMLEDIT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(76)
			p.Yaml_edit_fn()
		}

	case KlangParserYAMLDELETE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(77)
			p.Yaml_delete_fn()
		}

	case KlangParserIF:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(78)
			p.If_stat()
		}

	case KlangParserWHILE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(79)
			p.While_stat()
		}

	case KlangParserSLEEP:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(80)
			p.Sleep_fn()
		}

	case KlangParserLOG:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(81)
			p.Log()
		}

	case KlangParserOTHER:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(82)

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

	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(KlangParserID)
		}
		{
			p.SetState(87)
			p.Match(KlangParserASSIGN)
		}
		{
			p.SetState(88)
			p.expr(0)
		}
		{
			p.SetState(89)
			p.Match(KlangParserSCOL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.Match(KlangParserID)
		}
		{
			p.SetState(92)
			p.Match(KlangParserASSIGN)
		}
		{
			p.SetState(93)
			p.Load_fn()
		}
		{
			p.SetState(94)
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
		p.SetState(98)
		p.Match(KlangParserSHELLSCRIPT)
	}
	{
		p.SetState(99)
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
		p.SetState(101)
		p.Match(KlangParserJSONEDIT)
	}
	{
		p.SetState(102)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(103)
		p.Match(KlangParserID)
	}
	{
		p.SetState(104)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(105)
		p.String_or_id()
	}
	{
		p.SetState(106)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(107)
		p.expr(0)
	}
	{
		p.SetState(108)
		p.Match(KlangParserCPAR)
	}
	{
		p.SetState(109)
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
		p.SetState(111)
		p.Match(KlangParserJSONDELETE)
	}
	{
		p.SetState(112)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(113)
		p.Match(KlangParserID)
	}
	{
		p.SetState(114)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(115)
		p.String_or_id()
	}
	{
		p.SetState(116)
		p.Match(KlangParserCPAR)
	}
	{
		p.SetState(117)
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
		p.SetState(119)
		p.Match(KlangParserYAMLEDIT)
	}
	{
		p.SetState(120)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(121)
		p.Match(KlangParserID)
	}
	{
		p.SetState(122)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(123)
		p.String_or_id()
	}
	{
		p.SetState(124)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(125)
		p.expr(0)
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KlangParserCOMMA {
		{
			p.SetState(126)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(127)
			p.Match(KlangParserNUMBER)
		}

	}
	{
		p.SetState(130)
		p.Match(KlangParserCPAR)
	}
	{
		p.SetState(131)
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
		p.SetState(133)
		p.Match(KlangParserYAMLDELETE)
	}
	{
		p.SetState(134)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(135)
		p.Match(KlangParserID)
	}
	{
		p.SetState(136)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(137)
		p.String_or_id()
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KlangParserCOMMA {
		{
			p.SetState(138)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(139)
			p.Match(KlangParserNUMBER)
		}

	}
	{
		p.SetState(142)
		p.Match(KlangParserCPAR)
	}
	{
		p.SetState(143)
		p.Match(KlangParserSCOL)
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

func (s *Sleep_fnContext) OPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserOPAR, 0)
}

func (s *Sleep_fnContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(KlangParserNUMBER, 0)
}

func (s *Sleep_fnContext) CPAR() antlr.TerminalNode {
	return s.GetToken(KlangParserCPAR, 0)
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
	p.EnterRule(localctx, 18, KlangParserRULE_sleep_fn)

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
		p.SetState(145)
		p.Match(KlangParserSLEEP)
	}
	{
		p.SetState(146)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(147)
		p.Match(KlangParserNUMBER)
	}
	{
		p.SetState(148)
		p.Match(KlangParserCPAR)
	}
	{
		p.SetState(149)
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
	p.EnterRule(localctx, 20, KlangParserRULE_if_stat)

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
		p.SetState(151)
		p.Match(KlangParserIF)
	}
	{
		p.SetState(152)
		p.Condition_block()
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(153)
				p.Match(KlangParserELSE)
			}
			{
				p.SetState(154)
				p.Match(KlangParserIF)
			}
			{
				p.SetState(155)
				p.Condition_block()
			}

		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(161)
			p.Match(KlangParserELSE)
		}
		{
			p.SetState(162)
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
	p.EnterRule(localctx, 22, KlangParserRULE_condition_block)

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
		p.expr(0)
	}
	{
		p.SetState(166)
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
	p.EnterRule(localctx, 24, KlangParserRULE_stat_block)

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

	p.SetState(173)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserOBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(168)
			p.Match(KlangParserOBRACE)
		}
		{
			p.SetState(169)
			p.Block()
		}
		{
			p.SetState(170)
			p.Match(KlangParserCBRACE)
		}

	case KlangParserIF, KlangParserWHILE, KlangParserLOG, KlangParserJSONEDIT, KlangParserJSONDELETE, KlangParserYAMLEDIT, KlangParserYAMLDELETE, KlangParserSLEEP, KlangParserID, KlangParserOTHER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
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
	p.EnterRule(localctx, 26, KlangParserRULE_while_stat)

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
		p.Match(KlangParserWHILE)
	}
	{
		p.SetState(176)
		p.expr(0)
	}
	{
		p.SetState(177)
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
	p.EnterRule(localctx, 28, KlangParserRULE_log)

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
		p.Match(KlangParserLOG)
	}
	{
		p.SetState(180)
		p.expr(0)
	}
	{
		p.SetState(181)
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
	p.EnterRule(localctx, 30, KlangParserRULE_kubectl_command)

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

	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewApplyKubectlCommandContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(183)
			p.Match(KlangParserKUBECTL)
		}
		{
			p.SetState(184)
			p.Match(KlangParserAPPLY)
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(190)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case KlangParserNAMESPACE:
					{
						p.SetState(185)
						p.Match(KlangParserNAMESPACE)
					}
					{
						p.SetState(186)
						p.Ns()
					}

				case KlangParserID, KlangParserRAW_STRING_LIT, KlangParserSTRING:
					{
						p.SetState(187)
						p.String_or_id()
					}

				case KlangParserUPDATELOAD:
					{
						p.SetState(188)
						p.Match(KlangParserUPDATELOAD)
					}
					{
						p.SetState(189)
						p.Kubernetes_object_config()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(192)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
		}

	case 2:
		localctx = NewPatchKubectlCommandContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(194)
			p.Match(KlangParserKUBECTL)
		}
		{
			p.SetState(195)
			p.Match(KlangParserPATCH)
		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(203)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case KlangParserNAMESPACE:
					{
						p.SetState(196)
						p.Match(KlangParserNAMESPACE)
					}
					{
						p.SetState(197)
						p.Ns()
					}

				case KlangParserID, KlangParserPATH, KlangParserRAW_STRING_LIT, KlangParserSTRING:
					{
						p.SetState(198)
						p.Resource()
					}

				case KlangParserPATCHTYPE:
					{
						p.SetState(199)
						p.Match(KlangParserPATCHTYPE)
					}
					{
						p.SetState(200)
						p.Patch_type()
					}

				case KlangParserPATCHLOAD:
					{
						p.SetState(201)
						p.Match(KlangParserPATCHLOAD)
					}
					{
						p.SetState(202)
						p.String_or_id()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(205)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
		}

	case 3:
		localctx = NewGetKubectlCommandContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(207)
			p.Match(KlangParserKUBECTL)
		}
		{
			p.SetState(208)
			p.Match(KlangParserGET)
		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(212)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case KlangParserNAMESPACE:
					{
						p.SetState(209)
						p.Match(KlangParserNAMESPACE)
					}
					{
						p.SetState(210)
						p.Ns()
					}

				case KlangParserID, KlangParserPATH, KlangParserRAW_STRING_LIT, KlangParserSTRING:
					{
						p.SetState(211)
						p.Resource()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(214)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
		}

	case 4:
		localctx = NewDeleteKubectlCommandContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(216)
			p.Match(KlangParserKUBECTL)
		}
		{
			p.SetState(217)
			p.Match(KlangParserDELETE)
		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(221)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case KlangParserNAMESPACE:
					{
						p.SetState(218)
						p.Match(KlangParserNAMESPACE)
					}
					{
						p.SetState(219)
						p.Ns()
					}

				case KlangParserID, KlangParserPATH, KlangParserRAW_STRING_LIT, KlangParserSTRING:
					{
						p.SetState(220)
						p.Resource()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(223)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 32, KlangParserRULE_download_fn)
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
		p.Match(KlangParserDOWNLOAD)
	}
	{
		p.SetState(228)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(229)
		p.String_or_id()
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KlangParserCOMMA {
		{
			p.SetState(230)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(231)
			p.String_or_id()
		}

	}
	{
		p.SetState(234)
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
	p.EnterRule(localctx, 34, KlangParserRULE_json_select_fn)

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
		p.Match(KlangParserJSONSELECT)
	}
	{
		p.SetState(237)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(238)
		p.Match(KlangParserID)
	}
	{
		p.SetState(239)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(240)
		p.String_or_id()
	}
	{
		p.SetState(241)
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
	p.EnterRule(localctx, 36, KlangParserRULE_yaml_select_fn)
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
		p.SetState(243)
		p.Match(KlangParserYAMLSELECT)
	}
	{
		p.SetState(244)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(245)
		p.Match(KlangParserID)
	}
	{
		p.SetState(246)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(247)
		p.String_or_id()
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KlangParserCOMMA {
		{
			p.SetState(248)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(249)
			p.Match(KlangParserNUMBER)
		}

	}
	{
		p.SetState(252)
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
	p.EnterRule(localctx, 38, KlangParserRULE_load_fn)
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
		p.SetState(254)
		p.Match(KlangParserLOAD)
	}
	{
		p.SetState(255)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(256)
		p.String_or_id()
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KlangParserCOMMA {
		{
			p.SetState(257)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(258)
			p.Match(KlangParserSTRING)
		}

	}
	{
		p.SetState(261)
		p.Match(KlangParserCPAR)
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
	p.EnterRule(localctx, 40, KlangParserRULE_ns)

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

	p.SetState(265)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserID, KlangParserRAW_STRING_LIT, KlangParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.String_or_id()
		}

	case KlangParserPATH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(264)
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
	p.EnterRule(localctx, 42, KlangParserRULE_patch_type)

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

	p.SetState(269)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserPATH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(267)
			p.Match(KlangParserPATH)
		}

	case KlangParserID, KlangParserRAW_STRING_LIT, KlangParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(268)
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
	p.EnterRule(localctx, 44, KlangParserRULE_string_or_id)
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
	p.SetState(271)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-55)&-(0x1f+1)) == 0 && ((1<<uint((_la-55)))&((1<<(KlangParserID-55))|(1<<(KlangParserRAW_STRING_LIT-55))|(1<<(KlangParserSTRING-55)))) != 0) {
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
	p.EnterRule(localctx, 46, KlangParserRULE_resource)

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

	p.SetState(275)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserPATH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(273)
			p.Match(KlangParserPATH)
		}

	case KlangParserID, KlangParserRAW_STRING_LIT, KlangParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)
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
	p.EnterRule(localctx, 48, KlangParserRULE_kubernetes_object_config)

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
		p.SetState(277)
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
	_startState := 50
	p.EnterRecursionRule(localctx, 50, KlangParserRULE_expr, _p)
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
	p.SetState(290)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserMINUS:
		localctx = NewUnaryMinusExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(280)
			p.Match(KlangParserMINUS)
		}
		{
			p.SetState(281)
			p.expr(14)
		}

	case KlangParserNOT:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(282)
			p.Match(KlangParserNOT)
		}
		{
			p.SetState(283)
			p.expr(13)
		}

	case KlangParserKUBECTL:
		localctx = NewKubectlExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(284)
			p.Kubectl_command()
		}

	case KlangParserJSONSELECT:
		localctx = NewJsonSelectFnContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(285)
			p.Json_select_fn()
		}

	case KlangParserYAMLSELECT:
		localctx = NewYamlSelectFnContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(286)
			p.Yaml_select_fn()
		}

	case KlangParserSHELLSCRIPT:
		localctx = NewShellScriptContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(287)
			p.Shell_script()
		}

	case KlangParserDOWNLOAD:
		localctx = NewDownloadFnContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(288)
			p.Download_fn()
		}

	case KlangParserT__1, KlangParserT__3, KlangParserOPAR, KlangParserOBRACE, KlangParserTRUE, KlangParserFALSE, KlangParserNIL, KlangParserID, KlangParserNUMBER, KlangParserRAW_STRING_LIT, KlangParserSTRING:
		localctx = NewAtomExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(289)
			p.Atom()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(313)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(293)
					p.Match(KlangParserPOW)
				}
				{
					p.SetState(294)
					p.expr(15)
				}

			case 2:
				localctx = NewMultiplicationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(295)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				p.SetState(296)

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
					p.SetState(297)
					p.expr(13)
				}

			case 3:
				localctx = NewAdditiveExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(298)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				p.SetState(299)

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
					p.SetState(300)
					p.expr(12)
				}

			case 4:
				localctx = NewRelationalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(301)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				p.SetState(302)

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
					p.SetState(303)
					p.expr(11)
				}

			case 5:
				localctx = NewEqualityExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(304)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				p.SetState(305)

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
					p.SetState(306)
					p.expr(10)
				}

			case 6:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(307)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(308)
					p.Match(KlangParserAND)
				}
				{
					p.SetState(309)
					p.expr(9)
				}

			case 7:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(310)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(311)
					p.Match(KlangParserOR)
				}
				{
					p.SetState(312)
					p.expr(8)
				}

			}

		}
		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 52, KlangParserRULE_atom)
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

	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(318)
			p.Match(KlangParserOPAR)
		}
		{
			p.SetState(319)
			p.expr(0)
		}
		{
			p.SetState(320)
			p.Match(KlangParserCPAR)
		}

	case 2:
		localctx = NewNumberAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(322)
			p.Match(KlangParserNUMBER)
		}

	case 3:
		localctx = NewBooleanAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(323)
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
			p.SetState(324)
			p.Match(KlangParserRAW_STRING_LIT)
		}

	case 5:
		localctx = NewIdAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(325)
			p.Match(KlangParserID)
		}

	case 6:
		localctx = NewStringAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(326)
			p.Match(KlangParserSTRING)
		}

	case 7:
		localctx = NewJsonAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(327)
			p.Json()
		}

	case 8:
		localctx = NewNilAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(328)
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
	p.EnterRule(localctx, 54, KlangParserRULE_json)

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
		p.SetState(331)
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
	p.EnterRule(localctx, 56, KlangParserRULE_obj)
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

	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(333)
			p.Match(KlangParserOBRACE)
		}
		{
			p.SetState(334)
			p.Pair()
		}
		p.SetState(339)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == KlangParserCOMMA {
			{
				p.SetState(335)
				p.Match(KlangParserCOMMA)
			}
			{
				p.SetState(336)
				p.Pair()
			}

			p.SetState(341)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(342)
			p.Match(KlangParserCBRACE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(344)
			p.Match(KlangParserOBRACE)
		}
		{
			p.SetState(345)
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
	p.EnterRule(localctx, 58, KlangParserRULE_pair)

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
		p.SetState(348)
		p.Match(KlangParserSTRING)
	}
	{
		p.SetState(349)
		p.Match(KlangParserT__0)
	}
	{
		p.SetState(350)
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
	p.EnterRule(localctx, 60, KlangParserRULE_arr)
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

	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(352)
			p.Match(KlangParserT__1)
		}
		{
			p.SetState(353)
			p.Value()
		}
		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == KlangParserCOMMA {
			{
				p.SetState(354)
				p.Match(KlangParserCOMMA)
			}
			{
				p.SetState(355)
				p.Value()
			}

			p.SetState(360)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(361)
			p.Match(KlangParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(363)
			p.Match(KlangParserT__1)
		}
		{
			p.SetState(364)
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
	p.EnterRule(localctx, 62, KlangParserRULE_value)

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

	p.SetState(374)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(367)
			p.Match(KlangParserSTRING)
		}

	case KlangParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(368)
			p.Match(KlangParserNUMBER)
		}

	case KlangParserOBRACE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(369)
			p.Obj()
		}

	case KlangParserT__1:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(370)
			p.Arr()
		}

	case KlangParserTRUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(371)
			p.Match(KlangParserTRUE)
		}

	case KlangParserFALSE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(372)
			p.Match(KlangParserFALSE)
		}

	case KlangParserT__3:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(373)
			p.Match(KlangParserT__3)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *KlangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 25:
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
