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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 63, 370,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 3, 2, 3, 2, 3, 2,
	3, 3, 7, 3, 69, 10, 3, 12, 3, 14, 3, 72, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 84, 10, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 96, 10, 5, 3, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 128, 10, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 140, 10, 10, 3, 10, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 150, 10, 11, 12, 11, 14,
	11, 153, 11, 11, 3, 11, 3, 11, 5, 11, 157, 10, 11, 3, 12, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 167, 10, 13, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 6, 16, 184, 10, 16, 13, 16, 14, 16, 185, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 6, 16, 197, 10, 16, 13, 16,
	14, 16, 198, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 6, 16, 206, 10, 16, 13,
	16, 14, 16, 207, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 6, 16, 215, 10, 16,
	13, 16, 14, 16, 216, 5, 16, 219, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 5, 17, 226, 10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 244,
	10, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 253, 10,
	20, 3, 20, 3, 20, 3, 21, 3, 21, 5, 21, 259, 10, 21, 3, 22, 3, 22, 5, 22,
	263, 10, 22, 3, 23, 3, 23, 3, 24, 3, 24, 5, 24, 269, 10, 24, 3, 25, 3,
	25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 5, 26, 284, 10, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 7, 26, 307, 10, 26, 12, 26, 14, 26, 310, 11,
	26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 5, 27, 323, 10, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 7,
	29, 331, 10, 29, 12, 29, 14, 29, 334, 11, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	5, 29, 340, 10, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3,
	31, 7, 31, 350, 10, 31, 12, 31, 14, 31, 353, 11, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 5, 31, 359, 10, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 5, 32, 368, 10, 32, 3, 32, 2, 3, 50, 33, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
	54, 56, 58, 60, 62, 2, 8, 4, 2, 56, 56, 59, 60, 3, 2, 17, 19, 3, 2, 15,
	16, 3, 2, 11, 14, 3, 2, 9, 10, 3, 2, 29, 30, 2, 404, 2, 64, 3, 2, 2, 2,
	4, 70, 3, 2, 2, 2, 6, 83, 3, 2, 2, 2, 8, 95, 3, 2, 2, 2, 10, 97, 3, 2,
	2, 2, 12, 100, 3, 2, 2, 2, 14, 110, 3, 2, 2, 2, 16, 118, 3, 2, 2, 2, 18,
	132, 3, 2, 2, 2, 20, 144, 3, 2, 2, 2, 22, 158, 3, 2, 2, 2, 24, 166, 3,
	2, 2, 2, 26, 168, 3, 2, 2, 2, 28, 172, 3, 2, 2, 2, 30, 218, 3, 2, 2, 2,
	32, 220, 3, 2, 2, 2, 34, 229, 3, 2, 2, 2, 36, 236, 3, 2, 2, 2, 38, 247,
	3, 2, 2, 2, 40, 258, 3, 2, 2, 2, 42, 262, 3, 2, 2, 2, 44, 264, 3, 2, 2,
	2, 46, 268, 3, 2, 2, 2, 48, 270, 3, 2, 2, 2, 50, 283, 3, 2, 2, 2, 52, 322,
	3, 2, 2, 2, 54, 324, 3, 2, 2, 2, 56, 339, 3, 2, 2, 2, 58, 341, 3, 2, 2,
	2, 60, 358, 3, 2, 2, 2, 62, 367, 3, 2, 2, 2, 64, 65, 5, 4, 3, 2, 65, 66,
	7, 2, 2, 3, 66, 3, 3, 2, 2, 2, 67, 69, 5, 6, 4, 2, 68, 67, 3, 2, 2, 2,
	69, 72, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 5, 3, 2,
	2, 2, 72, 70, 3, 2, 2, 2, 73, 84, 5, 8, 5, 2, 74, 84, 5, 12, 7, 2, 75,
	84, 5, 14, 8, 2, 76, 84, 5, 16, 9, 2, 77, 84, 5, 18, 10, 2, 78, 84, 5,
	20, 11, 2, 79, 84, 5, 26, 14, 2, 80, 84, 5, 28, 15, 2, 81, 82, 7, 63, 2,
	2, 82, 84, 8, 4, 1, 2, 83, 73, 3, 2, 2, 2, 83, 74, 3, 2, 2, 2, 83, 75,
	3, 2, 2, 2, 83, 76, 3, 2, 2, 2, 83, 77, 3, 2, 2, 2, 83, 78, 3, 2, 2, 2,
	83, 79, 3, 2, 2, 2, 83, 80, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 7, 3, 2,
	2, 2, 85, 86, 7, 56, 2, 2, 86, 87, 7, 23, 2, 2, 87, 88, 5, 50, 26, 2, 88,
	89, 7, 22, 2, 2, 89, 96, 3, 2, 2, 2, 90, 91, 7, 56, 2, 2, 91, 92, 7, 23,
	2, 2, 92, 93, 5, 38, 20, 2, 93, 94, 7, 22, 2, 2, 94, 96, 3, 2, 2, 2, 95,
	85, 3, 2, 2, 2, 95, 90, 3, 2, 2, 2, 96, 9, 3, 2, 2, 2, 97, 98, 7, 54, 2,
	2, 98, 99, 5, 44, 23, 2, 99, 11, 3, 2, 2, 2, 100, 101, 7, 49, 2, 2, 101,
	102, 7, 24, 2, 2, 102, 103, 7, 56, 2, 2, 103, 104, 7, 28, 2, 2, 104, 105,
	5, 44, 23, 2, 105, 106, 7, 28, 2, 2, 106, 107, 5, 50, 26, 2, 107, 108,
	7, 25, 2, 2, 108, 109, 7, 22, 2, 2, 109, 13, 3, 2, 2, 2, 110, 111, 7, 50,
	2, 2, 111, 112, 7, 24, 2, 2, 112, 113, 7, 56, 2, 2, 113, 114, 7, 28, 2,
	2, 114, 115, 5, 44, 23, 2, 115, 116, 7, 25, 2, 2, 116, 117, 7, 22, 2, 2,
	117, 15, 3, 2, 2, 2, 118, 119, 7, 52, 2, 2, 119, 120, 7, 24, 2, 2, 120,
	121, 7, 56, 2, 2, 121, 122, 7, 28, 2, 2, 122, 123, 5, 44, 23, 2, 123, 124,
	7, 28, 2, 2, 124, 127, 5, 50, 26, 2, 125, 126, 7, 28, 2, 2, 126, 128, 7,
	57, 2, 2, 127, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 129, 3, 2, 2,
	2, 129, 130, 7, 25, 2, 2, 130, 131, 7, 22, 2, 2, 131, 17, 3, 2, 2, 2, 132,
	133, 7, 53, 2, 2, 133, 134, 7, 24, 2, 2, 134, 135, 7, 56, 2, 2, 135, 136,
	7, 28, 2, 2, 136, 139, 5, 44, 23, 2, 137, 138, 7, 28, 2, 2, 138, 140, 7,
	57, 2, 2, 139, 137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 141, 3, 2, 2,
	2, 141, 142, 7, 25, 2, 2, 142, 143, 7, 22, 2, 2, 143, 19, 3, 2, 2, 2, 144,
	145, 7, 32, 2, 2, 145, 151, 5, 22, 12, 2, 146, 147, 7, 33, 2, 2, 147, 148,
	7, 32, 2, 2, 148, 150, 5, 22, 12, 2, 149, 146, 3, 2, 2, 2, 150, 153, 3,
	2, 2, 2, 151, 149, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 156, 3, 2, 2,
	2, 153, 151, 3, 2, 2, 2, 154, 155, 7, 33, 2, 2, 155, 157, 5, 24, 13, 2,
	156, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 21, 3, 2, 2, 2, 158, 159,
	5, 50, 26, 2, 159, 160, 5, 24, 13, 2, 160, 23, 3, 2, 2, 2, 161, 162, 7,
	26, 2, 2, 162, 163, 5, 4, 3, 2, 163, 164, 7, 27, 2, 2, 164, 167, 3, 2,
	2, 2, 165, 167, 5, 6, 4, 2, 166, 161, 3, 2, 2, 2, 166, 165, 3, 2, 2, 2,
	167, 25, 3, 2, 2, 2, 168, 169, 7, 34, 2, 2, 169, 170, 5, 50, 26, 2, 170,
	171, 5, 24, 13, 2, 171, 27, 3, 2, 2, 2, 172, 173, 7, 35, 2, 2, 173, 174,
	5, 50, 26, 2, 174, 175, 7, 22, 2, 2, 175, 29, 3, 2, 2, 2, 176, 177, 7,
	36, 2, 2, 177, 183, 7, 37, 2, 2, 178, 179, 7, 42, 2, 2, 179, 184, 5, 40,
	21, 2, 180, 184, 5, 44, 23, 2, 181, 182, 7, 45, 2, 2, 182, 184, 5, 48,
	25, 2, 183, 178, 3, 2, 2, 2, 183, 180, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2,
	184, 185, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186,
	219, 3, 2, 2, 2, 187, 188, 7, 36, 2, 2, 188, 196, 7, 38, 2, 2, 189, 190,
	7, 42, 2, 2, 190, 197, 5, 40, 21, 2, 191, 197, 5, 46, 24, 2, 192, 193,
	7, 43, 2, 2, 193, 197, 5, 42, 22, 2, 194, 195, 7, 44, 2, 2, 195, 197, 5,
	44, 23, 2, 196, 189, 3, 2, 2, 2, 196, 191, 3, 2, 2, 2, 196, 192, 3, 2,
	2, 2, 196, 194, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 196, 3, 2, 2, 2,
	198, 199, 3, 2, 2, 2, 199, 219, 3, 2, 2, 2, 200, 201, 7, 36, 2, 2, 201,
	205, 7, 39, 2, 2, 202, 203, 7, 42, 2, 2, 203, 206, 5, 40, 21, 2, 204, 206,
	5, 46, 24, 2, 205, 202, 3, 2, 2, 2, 205, 204, 3, 2, 2, 2, 206, 207, 3,
	2, 2, 2, 207, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 219, 3, 2, 2,
	2, 209, 210, 7, 36, 2, 2, 210, 214, 7, 41, 2, 2, 211, 212, 7, 42, 2, 2,
	212, 215, 5, 40, 21, 2, 213, 215, 5, 46, 24, 2, 214, 211, 3, 2, 2, 2, 214,
	213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 216, 217,
	3, 2, 2, 2, 217, 219, 3, 2, 2, 2, 218, 176, 3, 2, 2, 2, 218, 187, 3, 2,
	2, 2, 218, 200, 3, 2, 2, 2, 218, 209, 3, 2, 2, 2, 219, 31, 3, 2, 2, 2,
	220, 221, 7, 55, 2, 2, 221, 222, 7, 24, 2, 2, 222, 225, 5, 44, 23, 2, 223,
	224, 7, 28, 2, 2, 224, 226, 5, 44, 23, 2, 225, 223, 3, 2, 2, 2, 225, 226,
	3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 228, 7, 25, 2, 2, 228, 33, 3, 2,
	2, 2, 229, 230, 7, 48, 2, 2, 230, 231, 7, 24, 2, 2, 231, 232, 7, 56, 2,
	2, 232, 233, 7, 28, 2, 2, 233, 234, 5, 44, 23, 2, 234, 235, 7, 25, 2, 2,
	235, 35, 3, 2, 2, 2, 236, 237, 7, 51, 2, 2, 237, 238, 7, 24, 2, 2, 238,
	239, 7, 56, 2, 2, 239, 240, 7, 28, 2, 2, 240, 243, 5, 44, 23, 2, 241, 242,
	7, 28, 2, 2, 242, 244, 7, 57, 2, 2, 243, 241, 3, 2, 2, 2, 243, 244, 3,
	2, 2, 2, 244, 245, 3, 2, 2, 2, 245, 246, 7, 25, 2, 2, 246, 37, 3, 2, 2,
	2, 247, 248, 7, 47, 2, 2, 248, 249, 7, 24, 2, 2, 249, 252, 5, 44, 23, 2,
	250, 251, 7, 28, 2, 2, 251, 253, 7, 60, 2, 2, 252, 250, 3, 2, 2, 2, 252,
	253, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254, 255, 7, 25, 2, 2, 255, 39,
	3, 2, 2, 2, 256, 259, 5, 44, 23, 2, 257, 259, 7, 58, 2, 2, 258, 256, 3,
	2, 2, 2, 258, 257, 3, 2, 2, 2, 259, 41, 3, 2, 2, 2, 260, 263, 7, 58, 2,
	2, 261, 263, 5, 44, 23, 2, 262, 260, 3, 2, 2, 2, 262, 261, 3, 2, 2, 2,
	263, 43, 3, 2, 2, 2, 264, 265, 9, 2, 2, 2, 265, 45, 3, 2, 2, 2, 266, 269,
	7, 58, 2, 2, 267, 269, 5, 44, 23, 2, 268, 266, 3, 2, 2, 2, 268, 267, 3,
	2, 2, 2, 269, 47, 3, 2, 2, 2, 270, 271, 5, 44, 23, 2, 271, 49, 3, 2, 2,
	2, 272, 273, 8, 26, 1, 2, 273, 274, 7, 16, 2, 2, 274, 284, 5, 50, 26, 16,
	275, 276, 7, 21, 2, 2, 276, 284, 5, 50, 26, 15, 277, 284, 5, 30, 16, 2,
	278, 284, 5, 34, 18, 2, 279, 284, 5, 36, 19, 2, 280, 284, 5, 10, 6, 2,
	281, 284, 5, 32, 17, 2, 282, 284, 5, 52, 27, 2, 283, 272, 3, 2, 2, 2, 283,
	275, 3, 2, 2, 2, 283, 277, 3, 2, 2, 2, 283, 278, 3, 2, 2, 2, 283, 279,
	3, 2, 2, 2, 283, 280, 3, 2, 2, 2, 283, 281, 3, 2, 2, 2, 283, 282, 3, 2,
	2, 2, 284, 308, 3, 2, 2, 2, 285, 286, 12, 17, 2, 2, 286, 287, 7, 20, 2,
	2, 287, 307, 5, 50, 26, 17, 288, 289, 12, 14, 2, 2, 289, 290, 9, 3, 2,
	2, 290, 307, 5, 50, 26, 15, 291, 292, 12, 13, 2, 2, 292, 293, 9, 4, 2,
	2, 293, 307, 5, 50, 26, 14, 294, 295, 12, 12, 2, 2, 295, 296, 9, 5, 2,
	2, 296, 307, 5, 50, 26, 13, 297, 298, 12, 11, 2, 2, 298, 299, 9, 6, 2,
	2, 299, 307, 5, 50, 26, 12, 300, 301, 12, 10, 2, 2, 301, 302, 7, 8, 2,
	2, 302, 307, 5, 50, 26, 11, 303, 304, 12, 9, 2, 2, 304, 305, 7, 7, 2, 2,
	305, 307, 5, 50, 26, 10, 306, 285, 3, 2, 2, 2, 306, 288, 3, 2, 2, 2, 306,
	291, 3, 2, 2, 2, 306, 294, 3, 2, 2, 2, 306, 297, 3, 2, 2, 2, 306, 300,
	3, 2, 2, 2, 306, 303, 3, 2, 2, 2, 307, 310, 3, 2, 2, 2, 308, 306, 3, 2,
	2, 2, 308, 309, 3, 2, 2, 2, 309, 51, 3, 2, 2, 2, 310, 308, 3, 2, 2, 2,
	311, 312, 7, 24, 2, 2, 312, 313, 5, 50, 26, 2, 313, 314, 7, 25, 2, 2, 314,
	323, 3, 2, 2, 2, 315, 323, 7, 57, 2, 2, 316, 323, 9, 7, 2, 2, 317, 323,
	7, 59, 2, 2, 318, 323, 7, 56, 2, 2, 319, 323, 7, 60, 2, 2, 320, 323, 5,
	54, 28, 2, 321, 323, 7, 31, 2, 2, 322, 311, 3, 2, 2, 2, 322, 315, 3, 2,
	2, 2, 322, 316, 3, 2, 2, 2, 322, 317, 3, 2, 2, 2, 322, 318, 3, 2, 2, 2,
	322, 319, 3, 2, 2, 2, 322, 320, 3, 2, 2, 2, 322, 321, 3, 2, 2, 2, 323,
	53, 3, 2, 2, 2, 324, 325, 5, 62, 32, 2, 325, 55, 3, 2, 2, 2, 326, 327,
	7, 26, 2, 2, 327, 332, 5, 58, 30, 2, 328, 329, 7, 28, 2, 2, 329, 331, 5,
	58, 30, 2, 330, 328, 3, 2, 2, 2, 331, 334, 3, 2, 2, 2, 332, 330, 3, 2,
	2, 2, 332, 333, 3, 2, 2, 2, 333, 335, 3, 2, 2, 2, 334, 332, 3, 2, 2, 2,
	335, 336, 7, 27, 2, 2, 336, 340, 3, 2, 2, 2, 337, 338, 7, 26, 2, 2, 338,
	340, 7, 27, 2, 2, 339, 326, 3, 2, 2, 2, 339, 337, 3, 2, 2, 2, 340, 57,
	3, 2, 2, 2, 341, 342, 7, 60, 2, 2, 342, 343, 7, 3, 2, 2, 343, 344, 5, 62,
	32, 2, 344, 59, 3, 2, 2, 2, 345, 346, 7, 4, 2, 2, 346, 351, 5, 62, 32,
	2, 347, 348, 7, 28, 2, 2, 348, 350, 5, 62, 32, 2, 349, 347, 3, 2, 2, 2,
	350, 353, 3, 2, 2, 2, 351, 349, 3, 2, 2, 2, 351, 352, 3, 2, 2, 2, 352,
	354, 3, 2, 2, 2, 353, 351, 3, 2, 2, 2, 354, 355, 7, 5, 2, 2, 355, 359,
	3, 2, 2, 2, 356, 357, 7, 4, 2, 2, 357, 359, 7, 5, 2, 2, 358, 345, 3, 2,
	2, 2, 358, 356, 3, 2, 2, 2, 359, 61, 3, 2, 2, 2, 360, 368, 7, 60, 2, 2,
	361, 368, 7, 57, 2, 2, 362, 368, 5, 56, 29, 2, 363, 368, 5, 60, 31, 2,
	364, 368, 7, 29, 2, 2, 365, 368, 7, 30, 2, 2, 366, 368, 7, 6, 2, 2, 367,
	360, 3, 2, 2, 2, 367, 361, 3, 2, 2, 2, 367, 362, 3, 2, 2, 2, 367, 363,
	3, 2, 2, 2, 367, 364, 3, 2, 2, 2, 367, 365, 3, 2, 2, 2, 367, 366, 3, 2,
	2, 2, 368, 63, 3, 2, 2, 2, 34, 70, 83, 95, 127, 139, 151, 156, 166, 183,
	185, 196, 198, 205, 207, 214, 216, 218, 225, 243, 252, 258, 262, 268, 283,
	306, 308, 322, 332, 339, 351, 358, 367,
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
	"'yamlEdit'", "'yamlDelete'", "'shellScript'", "'download'",
}
var symbolicNames = []string{
	"", "", "", "", "", "OR", "AND", "EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ",
	"PLUS", "MINUS", "MULT", "DIV", "MOD", "POW", "NOT", "SCOL", "ASSIGN",
	"OPAR", "CPAR", "OBRACE", "CBRACE", "COMMA", "TRUE", "FALSE", "NIL", "IF",
	"ELSE", "WHILE", "LOG", "KUBECTL", "APPLY", "PATCH", "GET", "REPLACE",
	"DELETE", "NAMESPACE", "PATCHTYPE", "PATCHLOAD", "UPDATELOAD", "JSONPATH",
	"LOAD", "JSONSELECT", "JSONEDIT", "JSONDELETE", "YAMLSELECT", "YAMLEDIT",
	"YAMLDELETE", "SHELLSCRIPT", "DOWNLOAD", "ID", "NUMBER", "PATH", "RAW_STRING_LIT",
	"STRING", "COMMENT", "SPACE", "OTHER",
}

var ruleNames = []string{
	"parse", "block", "stat", "assignment", "shell_script", "json_edit_fn",
	"json_delete_fn", "yaml_edit_fn", "yaml_delete_fn", "if_stat", "condition_block",
	"stat_block", "while_stat", "log", "kubectl_command", "download_fn", "json_select_fn",
	"yaml_select_fn", "load_fn", "ns", "patch_type", "string_or_id", "resource",
	"kubernetes_object_config", "expr", "atom", "json", "obj", "pair", "arr",
	"value",
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
	KlangParserID             = 54
	KlangParserNUMBER         = 55
	KlangParserPATH           = 56
	KlangParserRAW_STRING_LIT = 57
	KlangParserSTRING         = 58
	KlangParserCOMMENT        = 59
	KlangParserSPACE          = 60
	KlangParserOTHER          = 61
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
	KlangParserRULE_if_stat                  = 9
	KlangParserRULE_condition_block          = 10
	KlangParserRULE_stat_block               = 11
	KlangParserRULE_while_stat               = 12
	KlangParserRULE_log                      = 13
	KlangParserRULE_kubectl_command          = 14
	KlangParserRULE_download_fn              = 15
	KlangParserRULE_json_select_fn           = 16
	KlangParserRULE_yaml_select_fn           = 17
	KlangParserRULE_load_fn                  = 18
	KlangParserRULE_ns                       = 19
	KlangParserRULE_patch_type               = 20
	KlangParserRULE_string_or_id             = 21
	KlangParserRULE_resource                 = 22
	KlangParserRULE_kubernetes_object_config = 23
	KlangParserRULE_expr                     = 24
	KlangParserRULE_atom                     = 25
	KlangParserRULE_json                     = 26
	KlangParserRULE_obj                      = 27
	KlangParserRULE_pair                     = 28
	KlangParserRULE_arr                      = 29
	KlangParserRULE_value                    = 30
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
		p.SetState(62)
		p.Block()
	}
	{
		p.SetState(63)
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
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-30)&-(0x1f+1)) == 0 && ((1<<uint((_la-30)))&((1<<(KlangParserIF-30))|(1<<(KlangParserWHILE-30))|(1<<(KlangParserLOG-30))|(1<<(KlangParserJSONEDIT-30))|(1<<(KlangParserJSONDELETE-30))|(1<<(KlangParserYAMLEDIT-30))|(1<<(KlangParserYAMLDELETE-30))|(1<<(KlangParserID-30))|(1<<(KlangParserOTHER-30)))) != 0 {
		{
			p.SetState(65)
			p.Stat()
		}

		p.SetState(70)
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

	p.SetState(81)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.Assignment()
		}

	case KlangParserJSONEDIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.Json_edit_fn()
		}

	case KlangParserJSONDELETE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(73)
			p.Json_delete_fn()
		}

	case KlangParserYAMLEDIT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(74)
			p.Yaml_edit_fn()
		}

	case KlangParserYAMLDELETE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(75)
			p.Yaml_delete_fn()
		}

	case KlangParserIF:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(76)
			p.If_stat()
		}

	case KlangParserWHILE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(77)
			p.While_stat()
		}

	case KlangParserLOG:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(78)
			p.Log()
		}

	case KlangParserOTHER:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(79)

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

	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.Match(KlangParserID)
		}
		{
			p.SetState(84)
			p.Match(KlangParserASSIGN)
		}
		{
			p.SetState(85)
			p.expr(0)
		}
		{
			p.SetState(86)
			p.Match(KlangParserSCOL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(88)
			p.Match(KlangParserID)
		}
		{
			p.SetState(89)
			p.Match(KlangParserASSIGN)
		}
		{
			p.SetState(90)
			p.Load_fn()
		}
		{
			p.SetState(91)
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
		p.SetState(95)
		p.Match(KlangParserSHELLSCRIPT)
	}
	{
		p.SetState(96)
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
		p.SetState(98)
		p.Match(KlangParserJSONEDIT)
	}
	{
		p.SetState(99)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(100)
		p.Match(KlangParserID)
	}
	{
		p.SetState(101)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(102)
		p.String_or_id()
	}
	{
		p.SetState(103)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(104)
		p.expr(0)
	}
	{
		p.SetState(105)
		p.Match(KlangParserCPAR)
	}
	{
		p.SetState(106)
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
		p.SetState(108)
		p.Match(KlangParserJSONDELETE)
	}
	{
		p.SetState(109)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(110)
		p.Match(KlangParserID)
	}
	{
		p.SetState(111)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(112)
		p.String_or_id()
	}
	{
		p.SetState(113)
		p.Match(KlangParserCPAR)
	}
	{
		p.SetState(114)
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
		p.SetState(116)
		p.Match(KlangParserYAMLEDIT)
	}
	{
		p.SetState(117)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(118)
		p.Match(KlangParserID)
	}
	{
		p.SetState(119)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(120)
		p.String_or_id()
	}
	{
		p.SetState(121)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(122)
		p.expr(0)
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KlangParserCOMMA {
		{
			p.SetState(123)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(124)
			p.Match(KlangParserNUMBER)
		}

	}
	{
		p.SetState(127)
		p.Match(KlangParserCPAR)
	}
	{
		p.SetState(128)
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
		p.SetState(130)
		p.Match(KlangParserYAMLDELETE)
	}
	{
		p.SetState(131)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(132)
		p.Match(KlangParserID)
	}
	{
		p.SetState(133)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(134)
		p.String_or_id()
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KlangParserCOMMA {
		{
			p.SetState(135)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(136)
			p.Match(KlangParserNUMBER)
		}

	}
	{
		p.SetState(139)
		p.Match(KlangParserCPAR)
	}
	{
		p.SetState(140)
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
	p.EnterRule(localctx, 18, KlangParserRULE_if_stat)

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
		p.SetState(142)
		p.Match(KlangParserIF)
	}
	{
		p.SetState(143)
		p.Condition_block()
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(144)
				p.Match(KlangParserELSE)
			}
			{
				p.SetState(145)
				p.Match(KlangParserIF)
			}
			{
				p.SetState(146)
				p.Condition_block()
			}

		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(152)
			p.Match(KlangParserELSE)
		}
		{
			p.SetState(153)
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
	p.EnterRule(localctx, 20, KlangParserRULE_condition_block)

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
		p.SetState(156)
		p.expr(0)
	}
	{
		p.SetState(157)
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
	p.EnterRule(localctx, 22, KlangParserRULE_stat_block)

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

	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserOBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)
			p.Match(KlangParserOBRACE)
		}
		{
			p.SetState(160)
			p.Block()
		}
		{
			p.SetState(161)
			p.Match(KlangParserCBRACE)
		}

	case KlangParserIF, KlangParserWHILE, KlangParserLOG, KlangParserJSONEDIT, KlangParserJSONDELETE, KlangParserYAMLEDIT, KlangParserYAMLDELETE, KlangParserID, KlangParserOTHER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)
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
	p.EnterRule(localctx, 24, KlangParserRULE_while_stat)

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
		p.Match(KlangParserWHILE)
	}
	{
		p.SetState(167)
		p.expr(0)
	}
	{
		p.SetState(168)
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
	p.EnterRule(localctx, 26, KlangParserRULE_log)

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
		p.SetState(170)
		p.Match(KlangParserLOG)
	}
	{
		p.SetState(171)
		p.expr(0)
	}
	{
		p.SetState(172)
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
	p.EnterRule(localctx, 28, KlangParserRULE_kubectl_command)

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

	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewApplyKubectlCommandContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(174)
			p.Match(KlangParserKUBECTL)
		}
		{
			p.SetState(175)
			p.Match(KlangParserAPPLY)
		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(181)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case KlangParserNAMESPACE:
					{
						p.SetState(176)
						p.Match(KlangParserNAMESPACE)
					}
					{
						p.SetState(177)
						p.Ns()
					}

				case KlangParserID, KlangParserRAW_STRING_LIT, KlangParserSTRING:
					{
						p.SetState(178)
						p.String_or_id()
					}

				case KlangParserUPDATELOAD:
					{
						p.SetState(179)
						p.Match(KlangParserUPDATELOAD)
					}
					{
						p.SetState(180)
						p.Kubernetes_object_config()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(183)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
		}

	case 2:
		localctx = NewPatchKubectlCommandContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(185)
			p.Match(KlangParserKUBECTL)
		}
		{
			p.SetState(186)
			p.Match(KlangParserPATCH)
		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(194)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case KlangParserNAMESPACE:
					{
						p.SetState(187)
						p.Match(KlangParserNAMESPACE)
					}
					{
						p.SetState(188)
						p.Ns()
					}

				case KlangParserID, KlangParserPATH, KlangParserRAW_STRING_LIT, KlangParserSTRING:
					{
						p.SetState(189)
						p.Resource()
					}

				case KlangParserPATCHTYPE:
					{
						p.SetState(190)
						p.Match(KlangParserPATCHTYPE)
					}
					{
						p.SetState(191)
						p.Patch_type()
					}

				case KlangParserPATCHLOAD:
					{
						p.SetState(192)
						p.Match(KlangParserPATCHLOAD)
					}
					{
						p.SetState(193)
						p.String_or_id()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(196)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
		}

	case 3:
		localctx = NewGetKubectlCommandContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(198)
			p.Match(KlangParserKUBECTL)
		}
		{
			p.SetState(199)
			p.Match(KlangParserGET)
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
						p.SetState(200)
						p.Match(KlangParserNAMESPACE)
					}
					{
						p.SetState(201)
						p.Ns()
					}

				case KlangParserID, KlangParserPATH, KlangParserRAW_STRING_LIT, KlangParserSTRING:
					{
						p.SetState(202)
						p.Resource()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(205)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
		}

	case 4:
		localctx = NewDeleteKubectlCommandContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(207)
			p.Match(KlangParserKUBECTL)
		}
		{
			p.SetState(208)
			p.Match(KlangParserDELETE)
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
	p.EnterRule(localctx, 30, KlangParserRULE_download_fn)
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
		p.Match(KlangParserDOWNLOAD)
	}
	{
		p.SetState(219)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(220)
		p.String_or_id()
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KlangParserCOMMA {
		{
			p.SetState(221)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(222)
			p.String_or_id()
		}

	}
	{
		p.SetState(225)
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
	p.EnterRule(localctx, 32, KlangParserRULE_json_select_fn)

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
		p.Match(KlangParserJSONSELECT)
	}
	{
		p.SetState(228)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(229)
		p.Match(KlangParserID)
	}
	{
		p.SetState(230)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(231)
		p.String_or_id()
	}
	{
		p.SetState(232)
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
	p.EnterRule(localctx, 34, KlangParserRULE_yaml_select_fn)
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
		p.SetState(234)
		p.Match(KlangParserYAMLSELECT)
	}
	{
		p.SetState(235)
		p.Match(KlangParserOPAR)
	}
	{
		p.SetState(236)
		p.Match(KlangParserID)
	}
	{
		p.SetState(237)
		p.Match(KlangParserCOMMA)
	}
	{
		p.SetState(238)
		p.String_or_id()
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KlangParserCOMMA {
		{
			p.SetState(239)
			p.Match(KlangParserCOMMA)
		}
		{
			p.SetState(240)
			p.Match(KlangParserNUMBER)
		}

	}
	{
		p.SetState(243)
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
	p.EnterRule(localctx, 36, KlangParserRULE_load_fn)
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
		p.SetState(245)
		p.Match(KlangParserLOAD)
	}
	{
		p.SetState(246)
		p.Match(KlangParserOPAR)
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
			p.Match(KlangParserSTRING)
		}

	}
	{
		p.SetState(252)
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
	p.EnterRule(localctx, 38, KlangParserRULE_ns)

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

	p.SetState(256)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserID, KlangParserRAW_STRING_LIT, KlangParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(254)
			p.String_or_id()
		}

	case KlangParserPATH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(255)
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
	p.EnterRule(localctx, 40, KlangParserRULE_patch_type)

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

	p.SetState(260)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserPATH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(258)
			p.Match(KlangParserPATH)
		}

	case KlangParserID, KlangParserRAW_STRING_LIT, KlangParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(259)
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
	p.EnterRule(localctx, 42, KlangParserRULE_string_or_id)
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
	p.SetState(262)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-54)&-(0x1f+1)) == 0 && ((1<<uint((_la-54)))&((1<<(KlangParserID-54))|(1<<(KlangParserRAW_STRING_LIT-54))|(1<<(KlangParserSTRING-54)))) != 0) {
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
	p.EnterRule(localctx, 44, KlangParserRULE_resource)

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

	p.SetState(266)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserPATH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)
			p.Match(KlangParserPATH)
		}

	case KlangParserID, KlangParserRAW_STRING_LIT, KlangParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(265)
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
	p.EnterRule(localctx, 46, KlangParserRULE_kubernetes_object_config)

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
		p.SetState(268)
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
	_startState := 48
	p.EnterRecursionRule(localctx, 48, KlangParserRULE_expr, _p)
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
	p.SetState(281)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KlangParserMINUS:
		localctx = NewUnaryMinusExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(271)
			p.Match(KlangParserMINUS)
		}
		{
			p.SetState(272)
			p.expr(14)
		}

	case KlangParserNOT:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(273)
			p.Match(KlangParserNOT)
		}
		{
			p.SetState(274)
			p.expr(13)
		}

	case KlangParserKUBECTL:
		localctx = NewKubectlExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(275)
			p.Kubectl_command()
		}

	case KlangParserJSONSELECT:
		localctx = NewJsonSelectFnContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(276)
			p.Json_select_fn()
		}

	case KlangParserYAMLSELECT:
		localctx = NewYamlSelectFnContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(277)
			p.Yaml_select_fn()
		}

	case KlangParserSHELLSCRIPT:
		localctx = NewShellScriptContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(278)
			p.Shell_script()
		}

	case KlangParserDOWNLOAD:
		localctx = NewDownloadFnContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(279)
			p.Download_fn()
		}

	case KlangParserT__1, KlangParserT__3, KlangParserOPAR, KlangParserOBRACE, KlangParserTRUE, KlangParserFALSE, KlangParserNIL, KlangParserID, KlangParserNUMBER, KlangParserRAW_STRING_LIT, KlangParserSTRING:
		localctx = NewAtomExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(280)
			p.Atom()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(304)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(283)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(284)
					p.Match(KlangParserPOW)
				}
				{
					p.SetState(285)
					p.expr(15)
				}

			case 2:
				localctx = NewMultiplicationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(286)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				p.SetState(287)

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
					p.SetState(288)
					p.expr(13)
				}

			case 3:
				localctx = NewAdditiveExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(289)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				p.SetState(290)

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
					p.SetState(291)
					p.expr(12)
				}

			case 4:
				localctx = NewRelationalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				p.SetState(293)

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
					p.SetState(294)
					p.expr(11)
				}

			case 5:
				localctx = NewEqualityExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(295)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				p.SetState(296)

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
					p.SetState(297)
					p.expr(10)
				}

			case 6:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(298)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(299)
					p.Match(KlangParserAND)
				}
				{
					p.SetState(300)
					p.expr(9)
				}

			case 7:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, KlangParserRULE_expr)
				p.SetState(301)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(302)
					p.Match(KlangParserOR)
				}
				{
					p.SetState(303)
					p.expr(8)
				}

			}

		}
		p.SetState(308)
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
	p.EnterRule(localctx, 50, KlangParserRULE_atom)
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

	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(309)
			p.Match(KlangParserOPAR)
		}
		{
			p.SetState(310)
			p.expr(0)
		}
		{
			p.SetState(311)
			p.Match(KlangParserCPAR)
		}

	case 2:
		localctx = NewNumberAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(313)
			p.Match(KlangParserNUMBER)
		}

	case 3:
		localctx = NewBooleanAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(314)
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
			p.SetState(315)
			p.Match(KlangParserRAW_STRING_LIT)
		}

	case 5:
		localctx = NewIdAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(316)
			p.Match(KlangParserID)
		}

	case 6:
		localctx = NewStringAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(317)
			p.Match(KlangParserSTRING)
		}

	case 7:
		localctx = NewJsonAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(318)
			p.Json()
		}

	case 8:
		localctx = NewNilAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(319)
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
	p.EnterRule(localctx, 52, KlangParserRULE_json)

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
		p.SetState(322)
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
	p.EnterRule(localctx, 54, KlangParserRULE_obj)
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

	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(324)
			p.Match(KlangParserOBRACE)
		}
		{
			p.SetState(325)
			p.Pair()
		}
		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == KlangParserCOMMA {
			{
				p.SetState(326)
				p.Match(KlangParserCOMMA)
			}
			{
				p.SetState(327)
				p.Pair()
			}

			p.SetState(332)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(333)
			p.Match(KlangParserCBRACE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(335)
			p.Match(KlangParserOBRACE)
		}
		{
			p.SetState(336)
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
	p.EnterRule(localctx, 56, KlangParserRULE_pair)

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
		p.SetState(339)
		p.Match(KlangParserSTRING)
	}
	{
		p.SetState(340)
		p.Match(KlangParserT__0)
	}
	{
		p.SetState(341)
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
	p.EnterRule(localctx, 58, KlangParserRULE_arr)
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

	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(343)
			p.Match(KlangParserT__1)
		}
		{
			p.SetState(344)
			p.Value()
		}
		p.SetState(349)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == KlangParserCOMMA {
			{
				p.SetState(345)
				p.Match(KlangParserCOMMA)
			}
			{
				p.SetState(346)
				p.Value()
			}

			p.SetState(351)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(352)
			p.Match(KlangParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(354)
			p.Match(KlangParserT__1)
		}
		{
			p.SetState(355)
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
	p.EnterRule(localctx, 60, KlangParserRULE_value)

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

	switch p.GetTokenStream().LA(1) {
	case KlangParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(358)
			p.Match(KlangParserSTRING)
		}

	case KlangParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(359)
			p.Match(KlangParserNUMBER)
		}

	case KlangParserOBRACE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(360)
			p.Obj()
		}

	case KlangParserT__1:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(361)
			p.Arr()
		}

	case KlangParserTRUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(362)
			p.Match(KlangParserTRUE)
		}

	case KlangParserFALSE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(363)
			p.Match(KlangParserFALSE)
		}

	case KlangParserT__3:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(364)
			p.Match(KlangParserT__3)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *KlangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 24:
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
