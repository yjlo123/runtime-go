package main

import "fyne.io/fyne"

var font = &fyne.StaticResource{
	StaticName: "kongtext.ttf",
	StaticContent: []byte{
		0, 1, 0, 0, 0, 13, 0, 128, 0, 3, 0, 80, 71, 68, 69, 70, 1, 52, 1, 224, 0, 0, 37, 168, 0, 0, 0, 42, 71, 83, 85, 66, 172, 247, 174, 150, 0, 0, 37, 212, 0, 0, 0, 108, 79, 83, 47, 50, 81, 218, 192, 117, 0, 0, 1, 88, 0, 0, 0, 86, 99, 109, 97, 112, 61, 25, 70, 207, 0, 0, 5, 208, 0, 0, 1, 214, 103, 97, 115, 112, 255, 255, 0, 3, 0, 0, 1, 176, 0, 0, 0, 8, 103, 108, 121, 102, 219, 46, 13, 38, 0, 0, 11, 28, 0, 0, 26, 140, 104, 101, 97, 100, 219, 161, 59, 52, 0, 0, 0, 220, 0, 0, 0, 54, 104, 104, 101, 97, 7, 222, 3, 132, 0, 0, 1, 20, 0, 0, 0, 36, 104, 109, 116, 120, 43, 0, 37, 128, 0, 0, 9, 96, 0, 0, 1, 188, 108, 111, 99, 97, 192, 8, 198, 253, 0, 0, 7, 168, 0, 0, 1, 184, 109, 97, 120, 112, 1, 33, 0, 45, 0, 0, 1, 56, 0, 0, 0, 32, 110, 97, 109, 101, 46, 242, 34, 58, 0, 0, 1, 184, 0, 0, 4, 21, 112, 111, 115, 116, 94, 198, 250, 150, 0, 0, 38, 64, 0, 0, 1, 232, 0, 1, 0, 0, 0, 1, 2, 143, 196, 61, 194, 0, 95, 15, 60, 245, 0, 3, 4, 0, 0, 0, 0, 0, 188, 68, 250, 84, 0, 0, 0, 0, 188, 68, 250, 84, 0, 0, 255, 128, 4, 0, 3, 128, 0, 0, 0, 8, 0, 2, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 3, 128, 255, 128, 0, 92, 4, 0, 0, 0, 0, 0, 4, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 1, 0, 0, 0, 219, 0, 42, 0, 4, 0, 0, 0, 0, 0, 2, 0, 0, 0, 1, 0, 1, 0, 0, 0, 64, 0, 0, 0, 0, 0, 0, 0, 1, 4, 0, 1, 144, 0, 5, 0, 8, 0, 204, 0, 204, 0, 0, 0, 204, 0, 204, 0, 204, 0, 0, 0, 204, 0, 51, 1, 9, 0, 0, 2, 0, 5, 9, 0, 0, 0, 0, 0, 0, 128, 0, 0, 47, 0, 0, 0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 80, 102, 69, 100, 0, 64, 0, 12, 33, 34, 3, 128, 255, 128, 0, 92, 3, 128, 0, 128, 0, 0, 0, 1, 192, 212, 0, 0, 0, 0, 0, 0, 0, 1, 255, 255, 0, 2, 0, 0, 0, 28, 1, 86, 0, 0, 0, 3, 0, 0, 0, 0, 0, 74, 0, 0, 0, 0, 0, 3, 0, 0, 0, 1, 0, 16, 0, 76, 0, 0, 0, 3, 0, 0, 0, 2, 0, 14, 0, 94, 0, 0, 0, 3, 0, 0, 0, 3, 0, 76, 0, 110, 0, 0, 0, 3, 0, 0, 0, 4, 0, 32, 0, 188, 0, 0, 0, 3, 0, 0, 0, 5, 0, 26, 0, 222, 0, 0, 0, 3, 0, 0, 0, 6, 0, 30, 0, 250, 0, 1, 0, 0, 0, 0, 0, 0, 0, 37, 1, 26, 0, 1, 0, 0, 0, 0, 0, 1, 0, 8, 1, 64, 0, 1, 0, 0, 0, 0, 0, 2, 0, 7, 1, 73, 0, 1, 0, 0, 0, 0, 0, 3, 0, 38, 1, 81, 0, 1, 0, 0, 0, 0, 0, 4, 0, 16, 1, 120, 0, 1, 0, 0, 0, 0, 0, 5, 0, 13, 1, 137, 0, 1, 0, 0, 0, 0, 0, 6, 0, 15, 1, 151, 0, 3, 0, 1, 4, 0, 0, 0, 0, 74, 1, 167, 0, 3, 0, 1, 4, 0, 0, 1, 0, 16, 1, 243, 0, 3, 0, 1, 4, 0, 0, 2, 0, 14, 2, 5, 0, 3, 0, 1, 4, 0, 0, 3, 0, 76, 2, 21, 0, 3, 0, 1, 4, 0, 0, 4, 0, 32, 2, 99, 0, 3, 0, 1, 4, 0, 0, 5, 0, 24, 2, 133, 0, 3, 0, 1, 4, 0, 0, 6, 0, 30, 2, 159, 0, 3, 0, 1, 4, 9, 0, 0, 0, 74, 0, 0, 0, 3, 0, 1, 4, 9, 0, 1, 0, 16, 0, 76, 0, 3, 0, 1, 4, 9, 0, 2, 0, 14, 0, 94, 0, 3, 0, 1, 4, 9, 0, 3, 0, 76, 0, 110, 0, 3, 0, 1, 4, 9, 0, 4, 0, 32, 0, 188, 0, 3, 0, 1, 4, 9, 0, 5, 0, 26, 0, 222, 0, 3, 0, 1, 4, 9, 0, 6, 0, 30, 0, 250, 0, 84, 0, 114, 0, 117, 0, 101, 0, 84, 0, 121, 0, 112, 0, 101, 0, 32, 0, 99, 0, 111, 0, 110, 0, 118, 0, 101, 0, 114, 0, 115, 0, 105, 0, 111, 0, 110, 0, 32, 0, 169, 0, 32, 0, 50, 0, 48, 0, 48, 0, 51, 0, 32, 0, 99, 0, 111, 0, 100, 0, 101, 0, 109, 0, 97, 0, 110, 0, 51, 0, 56, 0, 46, 0, 0, 0, 75, 0, 111, 0, 110, 0, 103, 0, 116, 0, 101, 0, 120, 0, 116, 0, 0, 0, 82, 0, 101, 0, 103, 0, 117, 0, 108, 0, 97, 0, 114, 0, 0, 0, 80, 0, 102, 0, 97, 0, 69, 0, 100, 0, 105, 0, 116, 0, 32, 0, 58, 0, 32, 0, 75, 0, 111, 0, 110, 0, 103, 0, 116, 0, 101, 0, 120, 0, 116, 0, 32, 0, 82, 0, 101, 0, 103, 0, 117, 0, 108, 0, 97, 0, 114, 0, 32, 0, 58, 0, 32, 0, 51, 0, 48, 0, 45, 0, 54, 0, 45, 0, 50, 0, 48, 0, 48, 0, 51, 0, 0, 0, 75, 0, 111, 0, 110, 0, 103, 0, 116, 0, 101, 0, 120, 0, 116, 0, 32, 0, 82, 0, 101, 0, 103, 0, 117, 0, 108, 0, 97, 0, 114, 0, 0, 0, 86, 0, 101, 0, 114, 0, 115, 0, 105, 0, 111, 0, 110, 0, 32, 0, 49, 0, 46, 0, 48, 0, 49, 0, 32, 0, 0, 0, 75, 0, 111, 0, 110, 0, 103, 0, 116, 0, 101, 0, 120, 0, 116, 0, 82, 0, 101, 0, 103, 0, 117, 0, 108, 0, 97, 0, 114, 0, 0, 84, 114, 117, 101, 84, 121, 112, 101, 32, 99, 111, 110, 118, 101, 114, 115, 105, 111, 110, 32, 169, 32, 50, 48, 48, 51, 32, 99, 111, 100, 101, 109, 97, 110, 51, 56, 46, 0, 75, 111, 110, 103, 116, 101, 120, 116, 0, 82, 101, 103, 117, 108, 97, 114, 0, 80, 102, 97, 69, 100, 105, 116, 32, 58, 32, 75, 111, 110, 103, 116, 101, 120, 116, 32, 82, 101, 103, 117, 108, 97, 114, 32, 58, 32, 51, 48, 45, 54, 45, 50, 48, 48, 51, 0, 75, 111, 110, 103, 116, 101, 120, 116, 32, 82, 101, 103, 117, 108, 97, 114, 0, 86, 101, 114, 115, 105, 111, 110, 32, 49, 46, 48, 49, 32, 0, 75, 111, 110, 103, 116, 101, 120, 116, 82, 101, 103, 117, 108, 97, 114, 0, 0, 84, 0, 114, 0, 117, 0, 101, 0, 84, 0, 121, 0, 112, 0, 101, 0, 32, 0, 99, 0, 111, 0, 110, 0, 118, 0, 101, 0, 114, 0, 115, 0, 105, 0, 111, 0, 110, 0, 32, 0, 169, 0, 32, 0, 50, 0, 48, 0, 48, 0, 51, 0, 32, 0, 99, 0, 111, 0, 100, 0, 101, 0, 109, 0, 97, 0, 110, 0, 51, 0, 56, 0, 46, 0, 0, 0, 75, 0, 111, 0, 110, 0, 103, 0, 116, 0, 101, 0, 120, 0, 116, 0, 0, 0, 82, 0, 101, 0, 103, 0, 117, 0, 108, 0, 97, 0, 114, 0, 0, 0, 80, 0, 102, 0, 97, 0, 69, 0, 100, 0, 105, 0, 116, 0, 32, 0, 58, 0, 32, 0, 75, 0, 111, 0, 110, 0, 103, 0, 116, 0, 101, 0, 120, 0, 116, 0, 32, 0, 82, 0, 101, 0, 103, 0, 117, 0, 108, 0, 97, 0, 114, 0, 32, 0, 58, 0, 32, 0, 51, 0, 48, 0, 45, 0, 54, 0, 45, 0, 50, 0, 48, 0, 48, 0, 51, 0, 0, 0, 75, 0, 111, 0, 110, 0, 103, 0, 116, 0, 101, 0, 120, 0, 116, 0, 32, 0, 82, 0, 101, 0, 103, 0, 117, 0, 108, 0, 97, 0, 114, 0, 0, 0, 86, 0, 101, 0, 114, 0, 115, 0, 105, 0, 111, 0, 110, 0, 32, 0, 49, 0, 46, 0, 48, 0, 32, 0, 0, 0, 75, 0, 111, 0, 110, 0, 103, 0, 116, 0, 101, 0, 120, 0, 116, 0, 82, 0, 101, 0, 103, 0, 117, 0, 108, 0, 97, 0, 114, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 3, 0, 0, 0, 28, 0, 1, 0, 0, 0, 0, 0, 208, 0, 3, 0, 1, 0, 0, 0, 28, 0, 4, 0, 180, 0, 0, 0, 40, 0, 32, 0, 4, 0, 8, 0, 0, 0, 13, 0, 126, 0, 255, 1, 83, 1, 97, 1, 120, 1, 146, 2, 198, 2, 220, 32, 20, 32, 26, 32, 30, 32, 34, 32, 38, 32, 48, 32, 58, 32, 172, 33, 34, 255, 255, 0, 0, 0, 0, 0, 12, 0, 32, 0, 160, 1, 82, 1, 96, 1, 120, 1, 146, 2, 198, 2, 220, 32, 19, 32, 24, 32, 28, 32, 32, 32, 38, 32, 48, 32, 57, 32, 172, 33, 34, 255, 255, 0, 1, 0, 0, 255, 227, 255, 194, 255, 112, 255, 100, 255, 78, 255, 53, 254, 2, 253, 237, 224, 183, 224, 180, 224, 179, 224, 178, 224, 175, 224, 166, 224, 158, 224, 45, 223, 184, 0, 1, 0, 0, 0, 38, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 2, 0, 0, 1, 6, 0, 0, 1, 146, 178, 0, 0, 0, 0, 159, 1, 3, 0, 160, 192, 1, 0, 0, 0, 0, 0, 0, 0, 127, 126, 123, 128, 117, 116, 104, 111, 1, 0, 0, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 0, 134, 135, 137, 139, 147, 152, 158, 163, 162, 164, 166, 165, 167, 169, 171, 170, 172, 173, 175, 174, 176, 177, 179, 181, 180, 182, 184, 183, 188, 187, 189, 190, 0, 114, 100, 101, 105, 0, 120, 161, 112, 107, 0, 118, 106, 0, 136, 154, 0, 115, 0, 0, 103, 119, 0, 0, 0, 0, 0, 108, 124, 0, 168, 186, 129, 99, 110, 0, 0, 0, 0, 109, 125, 0, 98, 130, 133, 151, 0, 0, 0, 0, 0, 0, 0, 0, 185, 0, 193, 0, 0, 0, 0, 0, 0, 0, 0, 121, 0, 0, 0, 132, 140, 131, 141, 138, 143, 144, 145, 142, 149, 150, 0, 148, 156, 157, 155, 0, 0, 0, 113, 0, 0, 0, 122, 0, 0, 0, 0, 0, 0, 0, 0, 12, 0, 12, 0, 12, 0, 12, 0, 32, 0, 52, 0, 96, 0, 137, 0, 188, 0, 246, 1, 8, 1, 31, 1, 54, 1, 86, 1, 111, 1, 129, 1, 142, 1, 156, 1, 184, 1, 221, 1, 244, 2, 24, 2, 65, 2, 92, 2, 125, 2, 160, 2, 184, 2, 228, 3, 6, 3, 27, 3, 49, 3, 79, 3, 100, 3, 130, 3, 165, 3, 198, 3, 229, 4, 10, 4, 43, 4, 69, 4, 93, 4, 114, 4, 149, 4, 175, 4, 198, 4, 226, 5, 6, 5, 22, 5, 54, 5, 82, 5, 111, 5, 140, 5, 177, 5, 211, 5, 253, 6, 16, 6, 40, 6, 68, 6, 100, 6, 143, 6, 172, 6, 205, 6, 223, 6, 250, 7, 13, 7, 35, 7, 48, 7, 66, 7, 98, 7, 127, 7, 160, 7, 189, 7, 222, 7, 249, 8, 25, 8, 50, 8, 70, 8, 99, 8, 133, 8, 152, 8, 182, 8, 204, 8, 233, 9, 6, 9, 38, 9, 65, 9, 97, 9, 129, 9, 157, 9, 185, 9, 216, 9, 249, 10, 22, 10, 52, 10, 84, 10, 98, 10, 130, 10, 160, 10, 160, 10, 160, 10, 160, 10, 160, 10, 160, 10, 160, 10, 160, 10, 160, 10, 160, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 217, 10, 252, 11, 38, 11, 38, 11, 85, 11, 85, 11, 129, 11, 129, 11, 129, 11, 160, 11, 201, 11, 242, 12, 31, 12, 31, 12, 31, 12, 31, 12, 60, 12, 86, 12, 86, 12, 86, 12, 86, 12, 86, 12, 131, 12, 131, 12, 172, 12, 172, 12, 172, 12, 212, 12, 212, 13, 1, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 35, 13, 68, 4, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 1, 128, 1, 0, 0, 128, 0, 128, 0, 0, 0, 0, 1, 128, 1, 128, 1, 128, 0, 128, 0, 128, 1, 128, 0, 128, 1, 128, 0, 128, 0, 128, 1, 0, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 1, 128, 1, 128, 1, 0, 1, 0, 1, 0, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 1, 0, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 1, 128, 0, 128, 1, 128, 1, 0, 0, 128, 1, 0, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 1, 0, 0, 128, 0, 128, 1, 128, 0, 128, 0, 128, 1, 0, 0, 0, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 0, 128, 1, 128, 1, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 1, 0, 0, 0, 1, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 128, 0, 128, 0, 128, 0, 128, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 128, 0, 0, 0, 0, 0, 128, 0, 0, 0, 128, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 3, 51, 2, 205, 0, 3, 0, 0, 49, 17, 33, 17, 3, 51, 2, 205, 253, 51, 0, 2, 1, 128, 0, 0, 2, 128, 3, 0, 0, 3, 0, 7, 0, 0, 1, 33, 17, 33, 21, 33, 21, 33, 1, 128, 1, 0, 255, 0, 1, 0, 255, 0, 3, 0, 254, 0, 128, 128, 0, 2, 1, 0, 1, 128, 3, 0, 3, 0, 0, 3, 0, 7, 0, 0, 1, 51, 17, 35, 1, 51, 17, 35, 1, 0, 128, 128, 1, 128, 128, 128, 3, 0, 254, 128, 1, 128, 254, 128, 0, 2, 0, 128, 0, 128, 3, 128, 3, 0, 0, 27, 0, 31, 0, 0, 1, 51, 21, 33, 53, 51, 21, 51, 21, 35, 21, 51, 21, 35, 21, 35, 53, 33, 21, 35, 53, 35, 53, 51, 53, 35, 53, 51, 23, 21, 33, 53, 1, 0, 128, 1, 0, 128, 128, 128, 128, 128, 128, 255, 0, 128, 128, 128, 128, 128, 128, 1, 0, 3, 0, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 128, 0, 27, 0, 0, 1, 33, 21, 33, 21, 33, 21, 33, 21, 51, 21, 35, 21, 35, 21, 33, 53, 33, 53, 33, 53, 33, 53, 35, 53, 51, 53, 51, 1, 128, 1, 0, 1, 0, 253, 128, 2, 0, 128, 128, 128, 255, 0, 255, 0, 2, 128, 254, 0, 128, 128, 128, 3, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 0, 0, 3, 0, 0, 0, 0, 3, 128, 3, 128, 0, 3, 0, 29, 0, 33, 0, 0, 19, 33, 17, 33, 1, 51, 17, 35, 21, 35, 21, 35, 21, 35, 21, 35, 21, 33, 53, 51, 53, 51, 53, 51, 53, 51, 53, 51, 53, 51, 3, 33, 17, 33, 128, 1, 0, 255, 0, 2, 128, 128, 128, 128, 128, 128, 128, 255, 0, 128, 128, 128, 128, 128, 128, 128, 1, 0, 255, 0, 3, 128, 255, 0, 1, 0, 255, 0, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 254, 0, 255, 0, 0, 0, 4, 0, 0, 0, 0, 4, 0, 3, 0, 0, 25, 0, 29, 0, 33, 0, 41, 0, 0, 1, 33, 21, 51, 21, 33, 21, 35, 17, 51, 21, 33, 53, 35, 21, 33, 53, 35, 53, 51, 53, 51, 53, 35, 53, 59, 1, 21, 51, 53, 23, 21, 51, 53, 5, 21, 35, 21, 33, 53, 35, 53, 1, 0, 1, 128, 128, 1, 0, 128, 128, 255, 0, 128, 254, 0, 128, 128, 128, 128, 128, 128, 128, 128, 128, 254, 128, 128, 1, 128, 128, 3, 0, 128, 128, 128, 255, 0, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 0, 1, 1, 128, 1, 128, 3, 0, 3, 0, 0, 7, 0, 0, 1, 33, 17, 35, 21, 33, 53, 51, 2, 0, 1, 0, 128, 255, 0, 128, 3, 0, 255, 0, 128, 128, 0, 1, 1, 128, 0, 0, 3, 0, 3, 0, 0, 11, 0, 0, 1, 33, 21, 35, 17, 51, 21, 33, 53, 35, 17, 51, 2, 0, 1, 0, 128, 128, 255, 0, 128, 128, 3, 0, 128, 254, 0, 128, 128, 2, 0, 0, 0, 1, 1, 128, 0, 0, 3, 0, 3, 0, 0, 11, 0, 0, 1, 33, 21, 51, 17, 35, 21, 33, 53, 51, 17, 35, 1, 128, 1, 0, 128, 128, 255, 0, 128, 128, 3, 0, 128, 254, 0, 128, 128, 2, 0, 0, 0, 1, 0, 128, 0, 128, 3, 0, 3, 0, 0, 19, 0, 0, 19, 33, 21, 51, 53, 33, 17, 35, 21, 51, 17, 33, 53, 35, 21, 33, 17, 51, 53, 35, 128, 1, 0, 128, 1, 0, 128, 128, 255, 0, 128, 255, 0, 128, 128, 3, 0, 128, 128, 255, 0, 128, 255, 0, 128, 128, 1, 0, 128, 0, 0, 1, 0, 128, 0, 128, 3, 128, 3, 0, 0, 11, 0, 0, 1, 33, 17, 33, 21, 33, 17, 33, 17, 33, 53, 33, 1, 128, 1, 0, 1, 0, 255, 0, 255, 0, 255, 0, 1, 0, 3, 0, 255, 0, 128, 255, 0, 1, 0, 128, 0, 1, 1, 128, 0, 0, 3, 0, 1, 128, 0, 7, 0, 0, 1, 33, 17, 35, 21, 33, 53, 51, 2, 0, 1, 0, 128, 255, 0, 128, 1, 128, 255, 0, 128, 128, 0, 1, 0, 128, 1, 128, 3, 128, 2, 0, 0, 3, 0, 0, 19, 33, 21, 33, 128, 3, 0, 253, 0, 2, 0, 128, 0, 1, 1, 128, 0, 0, 2, 128, 1, 0, 0, 3, 0, 0, 1, 33, 17, 33, 1, 128, 1, 0, 255, 0, 1, 0, 255, 0, 0, 1, 0, 128, 0, 128, 3, 0, 3, 0, 0, 17, 0, 0, 1, 33, 21, 35, 21, 35, 21, 35, 21, 35, 21, 35, 17, 51, 53, 51, 53, 51, 2, 0, 1, 0, 128, 128, 128, 128, 128, 128, 128, 128, 3, 0, 128, 128, 128, 128, 128, 1, 0, 128, 128, 0, 0, 2, 0, 128, 0, 0, 3, 128, 3, 0, 0, 11, 0, 23, 0, 0, 1, 33, 21, 51, 17, 35, 21, 33, 53, 35, 17, 59, 1, 17, 51, 21, 35, 21, 33, 17, 35, 53, 51, 53, 1, 0, 2, 0, 128, 128, 254, 0, 128, 128, 128, 128, 128, 1, 0, 128, 128, 3, 0, 128, 254, 0, 128, 128, 2, 0, 255, 0, 128, 128, 1, 0, 128, 128, 0, 1, 1, 0, 0, 0, 3, 0, 3, 0, 0, 11, 0, 0, 1, 33, 17, 51, 21, 33, 53, 51, 17, 35, 53, 51, 1, 128, 1, 0, 128, 254, 0, 128, 128, 128, 3, 0, 253, 128, 128, 128, 1, 128, 128, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 21, 0, 0, 1, 33, 21, 51, 17, 35, 21, 33, 21, 33, 21, 33, 17, 51, 53, 33, 17, 33, 21, 33, 53, 51, 1, 0, 2, 0, 128, 128, 254, 128, 2, 0, 253, 0, 128, 1, 128, 255, 0, 255, 0, 128, 3, 0, 128, 255, 0, 128, 128, 128, 1, 0, 128, 1, 0, 128, 128, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 27, 0, 0, 1, 33, 21, 51, 21, 35, 21, 51, 17, 35, 21, 33, 53, 35, 53, 33, 21, 33, 17, 35, 53, 51, 53, 33, 21, 33, 53, 51, 1, 0, 2, 0, 128, 128, 128, 128, 254, 0, 128, 1, 0, 1, 0, 128, 128, 255, 0, 255, 0, 128, 3, 0, 128, 128, 128, 255, 0, 128, 128, 128, 128, 1, 0, 128, 128, 128, 128, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 13, 0, 0, 19, 33, 17, 51, 17, 33, 17, 51, 21, 35, 17, 33, 17, 33, 128, 1, 0, 128, 1, 0, 128, 128, 255, 0, 254, 128, 3, 0, 254, 128, 1, 0, 255, 0, 128, 255, 0, 1, 0, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 19, 0, 0, 19, 33, 21, 33, 21, 33, 21, 51, 17, 35, 21, 33, 53, 35, 53, 33, 21, 33, 17, 33, 128, 3, 0, 253, 128, 2, 0, 128, 128, 254, 0, 128, 1, 0, 1, 0, 254, 0, 3, 0, 128, 128, 128, 255, 0, 128, 128, 128, 128, 1, 0, 0, 0, 2, 0, 128, 0, 0, 3, 128, 3, 0, 0, 15, 0, 19, 0, 0, 1, 33, 21, 33, 21, 33, 21, 51, 17, 35, 21, 33, 53, 35, 17, 51, 19, 17, 33, 17, 1, 0, 2, 0, 254, 128, 1, 128, 128, 128, 254, 0, 128, 128, 128, 1, 0, 3, 0, 128, 128, 128, 255, 0, 128, 128, 2, 0, 255, 0, 255, 0, 1, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 11, 0, 0, 19, 33, 21, 35, 17, 35, 17, 33, 17, 51, 17, 33, 128, 3, 0, 128, 128, 255, 0, 128, 254, 128, 3, 0, 128, 255, 0, 254, 128, 1, 128, 1, 0, 0, 0, 3, 0, 128, 0, 0, 3, 128, 3, 0, 0, 19, 0, 23, 0, 27, 0, 0, 1, 33, 21, 51, 21, 35, 21, 51, 17, 35, 21, 33, 53, 35, 17, 51, 53, 35, 53, 59, 1, 21, 33, 53, 1, 17, 33, 17, 1, 0, 2, 0, 128, 128, 128, 128, 254, 0, 128, 128, 128, 128, 128, 1, 0, 255, 0, 1, 0, 3, 0, 128, 128, 128, 255, 0, 128, 128, 1, 0, 128, 128, 128, 128, 255, 0, 255, 0, 1, 0, 0, 2, 0, 128, 0, 0, 3, 128, 3, 0, 0, 15, 0, 19, 0, 0, 1, 33, 21, 51, 17, 35, 21, 33, 53, 33, 53, 33, 53, 35, 17, 59, 1, 17, 33, 17, 1, 0, 2, 0, 128, 128, 254, 0, 1, 128, 254, 128, 128, 128, 128, 1, 0, 3, 0, 128, 254, 0, 128, 128, 128, 128, 1, 0, 255, 0, 1, 0, 0, 2, 1, 128, 0, 128, 2, 128, 3, 0, 0, 3, 0, 7, 0, 0, 1, 33, 17, 33, 21, 33, 17, 33, 1, 128, 1, 0, 255, 0, 1, 0, 255, 0, 3, 0, 255, 0, 128, 255, 0, 0, 0, 2, 1, 128, 0, 0, 2, 128, 3, 0, 0, 3, 0, 9, 0, 0, 1, 33, 17, 33, 21, 33, 17, 35, 21, 35, 1, 128, 1, 0, 255, 0, 1, 0, 128, 128, 3, 0, 255, 0, 128, 255, 0, 128, 0, 1, 1, 0, 0, 128, 3, 0, 3, 0, 0, 19, 0, 0, 1, 33, 21, 35, 21, 35, 21, 51, 21, 51, 21, 33, 53, 35, 53, 35, 53, 51, 53, 51, 2, 0, 1, 0, 128, 128, 128, 128, 255, 0, 128, 128, 128, 128, 3, 0, 128, 128, 128, 128, 128, 128, 128, 128, 128, 0, 0, 2, 1, 0, 0, 128, 3, 128, 3, 0, 0, 3, 0, 7, 0, 0, 1, 33, 17, 33, 21, 33, 17, 33, 1, 0, 2, 128, 253, 128, 2, 128, 253, 128, 3, 0, 255, 0, 128, 255, 0, 0, 0, 1, 1, 0, 0, 128, 3, 0, 3, 0, 0, 19, 0, 0, 1, 33, 21, 51, 21, 51, 21, 35, 21, 35, 21, 33, 53, 51, 53, 51, 53, 35, 53, 35, 1, 0, 1, 0, 128, 128, 128, 128, 255, 0, 128, 128, 128, 128, 3, 0, 128, 128, 128, 128, 128, 128, 128, 128, 128, 0, 0, 2, 0, 128, 0, 0, 3, 128, 3, 0, 0, 15, 0, 19, 0, 0, 1, 33, 21, 51, 17, 35, 21, 33, 53, 33, 17, 33, 21, 33, 53, 51, 19, 33, 21, 33, 1, 0, 2, 0, 128, 128, 254, 128, 1, 0, 255, 0, 255, 0, 128, 128, 1, 0, 255, 0, 3, 0, 128, 255, 0, 128, 128, 1, 0, 128, 128, 254, 0, 128, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 19, 0, 0, 1, 33, 21, 51, 17, 35, 21, 33, 17, 51, 53, 33, 17, 33, 21, 33, 53, 35, 17, 51, 1, 0, 2, 0, 128, 128, 255, 0, 128, 255, 0, 2, 0, 253, 128, 128, 128, 3, 0, 128, 255, 0, 128, 1, 0, 128, 254, 0, 128, 128, 2, 0, 0, 2, 0, 128, 0, 0, 3, 128, 3, 0, 0, 11, 0, 15, 0, 0, 1, 33, 21, 51, 17, 33, 17, 33, 17, 33, 17, 59, 1, 17, 33, 17, 1, 0, 2, 0, 128, 255, 0, 255, 0, 255, 0, 128, 128, 1, 0, 3, 0, 128, 253, 128, 1, 0, 255, 0, 2, 128, 255, 0, 1, 0, 0, 3, 0, 128, 0, 0, 3, 128, 3, 0, 0, 11, 0, 15, 0, 19, 0, 0, 19, 33, 21, 51, 21, 35, 21, 51, 17, 35, 21, 33, 1, 21, 33, 53, 1, 17, 33, 17, 128, 2, 128, 128, 128, 128, 128, 253, 128, 1, 0, 1, 0, 255, 0, 1, 0, 3, 0, 128, 128, 128, 255, 0, 128, 2, 128, 128, 128, 255, 0, 255, 0, 1, 0, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 19, 0, 0, 1, 33, 21, 51, 21, 33, 53, 33, 17, 33, 53, 33, 21, 35, 21, 33, 53, 35, 17, 51, 1, 0, 2, 0, 128, 255, 0, 255, 0, 1, 0, 1, 0, 128, 254, 0, 128, 128, 3, 0, 128, 128, 128, 254, 0, 128, 128, 128, 128, 2, 0, 0, 0, 2, 0, 128, 0, 0, 3, 128, 3, 0, 0, 7, 0, 11, 0, 0, 19, 33, 21, 51, 17, 35, 21, 33, 1, 17, 33, 17, 128, 2, 128, 128, 128, 253, 128, 1, 0, 1, 0, 3, 0, 128, 254, 0, 128, 2, 128, 254, 0, 2, 0, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 11, 0, 0, 19, 33, 21, 33, 21, 33, 21, 33, 17, 33, 21, 33, 128, 3, 0, 254, 0, 1, 0, 255, 0, 2, 0, 253, 0, 3, 0, 128, 128, 128, 255, 0, 128, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 9, 0, 0, 19, 33, 21, 33, 21, 33, 21, 33, 17, 33, 128, 3, 0, 254, 0, 1, 0, 255, 0, 255, 0, 3, 0, 128, 128, 128, 254, 128, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 21, 0, 0, 1, 33, 21, 51, 21, 33, 53, 33, 17, 33, 53, 35, 53, 33, 17, 35, 21, 33, 53, 35, 17, 51, 1, 0, 2, 0, 128, 255, 0, 255, 0, 1, 0, 128, 1, 128, 128, 254, 0, 128, 128, 3, 0, 128, 128, 128, 254, 0, 128, 128, 255, 0, 128, 128, 2, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 11, 0, 0, 19, 33, 17, 33, 17, 33, 17, 33, 17, 33, 17, 33, 128, 1, 0, 1, 0, 1, 0, 255, 0, 255, 0, 255, 0, 3, 0, 255, 0, 1, 0, 253, 0, 1, 128, 254, 128, 0, 0, 1, 1, 0, 0, 0, 3, 0, 3, 0, 0, 11, 0, 0, 1, 33, 21, 35, 17, 51, 21, 33, 53, 51, 17, 35, 1, 0, 2, 0, 128, 128, 254, 0, 128, 128, 3, 0, 128, 254, 0, 128, 128, 2, 0, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 15, 0, 0, 1, 33, 21, 35, 17, 35, 21, 33, 53, 35, 53, 33, 21, 51, 17, 33, 1, 0, 2, 128, 128, 128, 254, 128, 128, 1, 0, 128, 255, 0, 3, 0, 128, 254, 0, 128, 128, 128, 128, 2, 0, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 23, 0, 0, 19, 33, 17, 51, 53, 51, 53, 33, 21, 35, 21, 35, 17, 51, 21, 51, 21, 33, 53, 35, 53, 35, 17, 33, 128, 1, 0, 128, 128, 1, 0, 128, 128, 128, 128, 255, 0, 128, 128, 255, 0, 3, 0, 255, 0, 128, 128, 128, 128, 255, 0, 128, 128, 128, 128, 255, 0, 0, 0, 1, 0, 128, 0, 0, 3, 0, 3, 0, 0, 5, 0, 0, 19, 33, 17, 33, 21, 33, 128, 1, 0, 1, 128, 253, 128, 3, 0, 253, 128, 128, 0, 1, 0, 128, 0, 0, 4, 0, 3, 0, 0, 19, 0, 0, 19, 33, 21, 51, 21, 51, 53, 51, 53, 33, 17, 33, 17, 35, 21, 35, 53, 35, 17, 33, 128, 1, 0, 128, 128, 128, 1, 0, 255, 0, 128, 128, 128, 255, 0, 3, 0, 128, 128, 128, 128, 253, 0, 1, 128, 128, 128, 254, 128, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 15, 0, 0, 19, 33, 21, 51, 21, 51, 17, 33, 17, 33, 53, 35, 53, 35, 17, 33, 128, 1, 0, 128, 128, 1, 0, 255, 0, 128, 128, 255, 0, 3, 0, 128, 128, 1, 0, 253, 0, 128, 128, 255, 0, 0, 0, 2, 0, 128, 0, 0, 3, 128, 3, 0, 0, 11, 0, 15, 0, 0, 1, 33, 21, 51, 17, 35, 21, 33, 53, 35, 17, 59, 1, 17, 33, 17, 1, 0, 2, 0, 128, 128, 254, 0, 128, 128, 128, 1, 0, 3, 0, 128, 254, 0, 128, 128, 2, 0, 254, 0, 2, 0, 0, 2, 0, 128, 0, 0, 3, 128, 3, 0, 0, 9, 0, 13, 0, 0, 19, 33, 21, 51, 17, 35, 21, 33, 17, 33, 1, 17, 33, 17, 128, 2, 128, 128, 128, 254, 128, 255, 0, 1, 0, 1, 0, 3, 0, 128, 255, 0, 128, 255, 0, 2, 128, 255, 0, 1, 0, 0, 0, 2, 0, 128, 0, 0, 3, 128, 3, 0, 0, 17, 0, 23, 0, 0, 1, 33, 21, 51, 17, 35, 21, 51, 21, 35, 53, 35, 21, 33, 53, 35, 17, 59, 1, 17, 51, 53, 51, 17, 1, 0, 2, 0, 128, 128, 128, 128, 128, 254, 128, 128, 128, 128, 128, 128, 3, 0, 128, 254, 128, 128, 128, 128, 128, 128, 2, 0, 254, 0, 128, 1, 128, 0, 0, 2, 0, 128, 0, 0, 3, 128, 3, 0, 0, 13, 0, 17, 0, 0, 19, 33, 21, 51, 17, 35, 21, 51, 17, 33, 17, 33, 17, 33, 1, 17, 33, 17, 128, 2, 128, 128, 128, 128, 255, 0, 255, 0, 255, 0, 1, 0, 1, 0, 3, 0, 128, 255, 0, 128, 255, 0, 1, 0, 255, 0, 2, 128, 255, 0, 1, 0, 0, 2, 0, 128, 0, 0, 3, 128, 3, 0, 0, 23, 0, 27, 0, 0, 1, 33, 21, 51, 21, 35, 21, 51, 17, 35, 21, 33, 53, 35, 53, 33, 21, 33, 17, 33, 53, 35, 53, 59, 1, 21, 33, 53, 1, 0, 2, 0, 128, 128, 128, 128, 254, 0, 128, 1, 0, 1, 0, 254, 128, 128, 128, 128, 1, 128, 3, 0, 128, 128, 128, 255, 0, 128, 128, 128, 128, 1, 0, 128, 128, 128, 128, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 7, 0, 0, 19, 33, 21, 33, 17, 33, 17, 33, 128, 3, 0, 255, 0, 255, 0, 255, 0, 3, 0, 128, 253, 128, 2, 128, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 11, 0, 0, 19, 33, 17, 33, 17, 33, 17, 35, 21, 33, 53, 35, 128, 1, 0, 1, 0, 1, 0, 128, 254, 0, 128, 3, 0, 253, 128, 2, 128, 253, 128, 128, 128, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 15, 0, 0, 19, 33, 17, 33, 17, 33, 17, 35, 21, 35, 21, 33, 53, 35, 53, 35, 128, 1, 0, 1, 0, 1, 0, 128, 128, 255, 0, 128, 128, 3, 0, 254, 0, 2, 0, 254, 0, 128, 128, 128, 128, 0, 0, 1, 0, 128, 0, 0, 4, 0, 3, 0, 0, 19, 0, 0, 19, 33, 17, 51, 53, 51, 21, 51, 17, 33, 17, 33, 53, 35, 53, 35, 21, 35, 21, 33, 128, 1, 0, 128, 128, 128, 1, 0, 255, 0, 128, 128, 128, 255, 0, 3, 0, 254, 128, 128, 128, 1, 128, 253, 0, 128, 128, 128, 128, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 31, 0, 0, 19, 33, 21, 33, 53, 33, 21, 35, 21, 35, 21, 51, 21, 51, 17, 35, 53, 35, 53, 33, 21, 35, 21, 35, 17, 51, 53, 51, 53, 35, 53, 35, 128, 1, 0, 1, 0, 1, 0, 128, 128, 128, 128, 128, 128, 255, 0, 128, 128, 128, 128, 128, 128, 3, 0, 128, 128, 128, 128, 128, 128, 255, 0, 128, 128, 128, 128, 1, 0, 128, 128, 128, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 15, 0, 0, 19, 33, 17, 33, 17, 33, 17, 35, 21, 35, 17, 33, 17, 35, 53, 35, 128, 1, 0, 1, 0, 1, 0, 128, 128, 255, 0, 128, 128, 3, 0, 255, 0, 1, 0, 255, 0, 128, 254, 128, 1, 128, 128, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 21, 0, 0, 19, 33, 21, 35, 21, 35, 21, 35, 21, 35, 21, 33, 21, 33, 17, 51, 53, 51, 53, 51, 53, 33, 128, 3, 0, 128, 128, 128, 128, 2, 0, 253, 0, 128, 128, 128, 254, 128, 3, 0, 128, 128, 128, 128, 128, 128, 1, 0, 128, 128, 128, 0, 0, 1, 1, 128, 0, 0, 3, 0, 3, 0, 0, 7, 0, 0, 1, 33, 21, 35, 17, 51, 21, 33, 1, 128, 1, 128, 128, 128, 254, 128, 3, 0, 128, 254, 0, 128, 0, 1, 0, 128, 0, 128, 3, 0, 3, 0, 0, 17, 0, 0, 19, 33, 21, 51, 21, 51, 21, 51, 17, 35, 53, 35, 53, 35, 53, 35, 53, 35, 128, 1, 0, 128, 128, 128, 128, 128, 128, 128, 128, 3, 0, 128, 128, 128, 255, 0, 128, 128, 128, 128, 0, 1, 1, 128, 0, 0, 3, 0, 3, 0, 0, 7, 0, 0, 1, 33, 17, 33, 53, 51, 17, 35, 1, 128, 1, 128, 254, 128, 128, 128, 3, 0, 253, 0, 128, 2, 0, 0, 0, 1, 1, 0, 2, 128, 3, 128, 3, 128, 0, 11, 0, 0, 1, 33, 21, 51, 21, 33, 53, 35, 21, 33, 53, 51, 1, 128, 1, 128, 128, 255, 0, 128, 255, 0, 128, 3, 128, 128, 128, 128, 128, 128, 0, 1, 0, 128, 0, 0, 3, 128, 0, 128, 0, 3, 0, 0, 55, 33, 21, 33, 128, 3, 0, 253, 0, 128, 128, 0, 0, 1, 1, 0, 1, 128, 2, 128, 3, 0, 0, 7, 0, 0, 1, 33, 17, 51, 21, 33, 53, 35, 1, 0, 1, 0, 128, 255, 0, 128, 3, 0, 255, 0, 128, 128, 0, 2, 0, 128, 0, 0, 4, 0, 2, 128, 0, 13, 0, 17, 0, 0, 1, 33, 17, 51, 21, 33, 53, 35, 21, 33, 53, 35, 17, 59, 1, 17, 33, 17, 1, 0, 2, 128, 128, 255, 0, 128, 254, 128, 128, 128, 128, 1, 0, 2, 128, 254, 0, 128, 128, 128, 128, 1, 128, 254, 128, 1, 128, 0, 0, 2, 0, 128, 0, 0, 3, 128, 3, 128, 0, 9, 0, 13, 0, 0, 19, 33, 17, 33, 21, 51, 17, 35, 21, 33, 1, 17, 33, 17, 128, 1, 0, 1, 128, 128, 128, 253, 128, 1, 0, 1, 0, 3, 128, 255, 0, 128, 254, 128, 128, 2, 0, 254, 128, 1, 128, 0, 0, 1, 0, 128, 0, 0, 3, 128, 2, 128, 0, 19, 0, 0, 1, 33, 21, 51, 21, 33, 53, 33, 17, 33, 53, 33, 21, 35, 21, 33, 53, 35, 17, 51, 1, 0, 2, 0, 128, 255, 0, 255, 0, 1, 0, 1, 0, 128, 254, 0, 128, 128, 2, 128, 128, 128, 128, 254, 128, 128, 128, 128, 128, 1, 128, 0, 0, 2, 0, 128, 0, 0, 3, 128, 3, 128, 0, 9, 0, 13, 0, 0, 1, 33, 17, 33, 53, 35, 17, 51, 53, 33, 5, 17, 33, 17, 2, 128, 1, 0, 253, 128, 128, 128, 1, 128, 255, 0, 1, 0, 3, 128, 252, 128, 128, 1, 128, 128, 128, 254, 128, 1, 128, 0, 0, 2, 0, 128, 0, 0, 3, 128, 2, 128, 0, 15, 0, 19, 0, 0, 1, 33, 21, 51, 21, 35, 21, 33, 21, 33, 21, 33, 53, 35, 17, 59, 1, 21, 33, 53, 1, 0, 2, 0, 128, 128, 254, 128, 2, 0, 253, 128, 128, 128, 128, 1, 0, 2, 128, 128, 128, 128, 128, 128, 128, 1, 128, 128, 128, 0, 0, 1, 1, 0, 0, 0, 3, 128, 3, 0, 0, 15, 0, 0, 1, 33, 21, 33, 21, 51, 21, 35, 17, 33, 17, 35, 53, 51, 53, 51, 2, 0, 1, 128, 255, 0, 128, 128, 255, 0, 128, 128, 128, 3, 0, 128, 128, 128, 254, 128, 1, 128, 128, 128, 0, 2, 0, 128, 255, 128, 3, 128, 2, 128, 0, 13, 0, 17, 0, 0, 1, 33, 17, 35, 21, 33, 53, 33, 53, 33, 53, 35, 17, 59, 1, 17, 33, 17, 1, 0, 2, 128, 128, 254, 0, 1, 128, 254, 128, 128, 128, 128, 1, 0, 2, 128, 253, 128, 128, 128, 128, 128, 1, 0, 255, 0, 1, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 128, 0, 11, 0, 0, 19, 33, 17, 33, 21, 51, 17, 33, 17, 33, 17, 33, 128, 1, 0, 1, 128, 128, 255, 0, 255, 0, 255, 0, 3, 128, 255, 0, 128, 254, 0, 2, 0, 254, 0, 0, 0, 2, 1, 128, 0, 0, 2, 128, 3, 128, 0, 3, 0, 7, 0, 0, 1, 33, 21, 33, 21, 33, 17, 33, 1, 128, 1, 0, 255, 0, 1, 0, 255, 0, 3, 128, 128, 128, 253, 128, 0, 2, 0, 128, 255, 128, 3, 0, 3, 128, 0, 3, 0, 15, 0, 0, 1, 33, 21, 33, 21, 33, 17, 35, 21, 33, 53, 35, 53, 33, 21, 51, 2, 0, 1, 0, 255, 0, 1, 0, 128, 254, 128, 128, 1, 0, 128, 3, 128, 128, 128, 253, 128, 128, 128, 128, 128, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 128, 0, 21, 0, 0, 19, 33, 17, 33, 53, 33, 21, 35, 21, 35, 21, 51, 21, 51, 21, 33, 53, 35, 53, 35, 17, 33, 128, 1, 0, 1, 0, 1, 0, 128, 128, 128, 128, 255, 0, 128, 128, 255, 0, 3, 128, 254, 128, 128, 128, 128, 128, 128, 128, 128, 128, 255, 0, 0, 0, 1, 1, 0, 0, 0, 3, 0, 3, 128, 0, 7, 0, 0, 1, 33, 17, 33, 21, 33, 53, 35, 1, 0, 1, 0, 1, 0, 254, 128, 128, 3, 128, 253, 0, 128, 128, 0, 0, 1, 0, 0, 0, 0, 3, 128, 2, 128, 0, 17, 0, 0, 17, 33, 21, 51, 53, 33, 21, 51, 17, 33, 17, 35, 17, 35, 17, 35, 17, 33, 1, 128, 128, 1, 0, 128, 255, 0, 128, 128, 128, 255, 0, 2, 128, 128, 128, 128, 254, 0, 2, 0, 254, 0, 2, 0, 254, 0, 0, 1, 0, 128, 0, 0, 3, 128, 2, 128, 0, 9, 0, 0, 19, 33, 21, 51, 17, 33, 17, 33, 17, 33, 128, 2, 128, 128, 255, 0, 255, 0, 255, 0, 2, 128, 128, 254, 0, 2, 0, 254, 0, 0, 0, 2, 0, 128, 0, 0, 3, 128, 2, 128, 0, 11, 0, 15, 0, 0, 1, 33, 21, 51, 17, 35, 21, 33, 53, 35, 17, 59, 1, 17, 33, 17, 1, 0, 2, 0, 128, 128, 254, 0, 128, 128, 128, 1, 0, 2, 128, 128, 254, 128, 128, 128, 1, 128, 254, 128, 1, 128, 0, 2, 0, 128, 255, 128, 3, 128, 2, 128, 0, 9, 0, 13, 0, 0, 19, 33, 21, 51, 17, 35, 21, 33, 17, 33, 1, 17, 33, 17, 128, 2, 128, 128, 128, 254, 128, 255, 0, 1, 0, 1, 0, 2, 128, 128, 255, 0, 128, 255, 0, 2, 128, 255, 0, 1, 0, 0, 0, 2, 0, 128, 255, 128, 4, 0, 2, 128, 0, 13, 0, 17, 0, 0, 1, 33, 17, 51, 21, 35, 21, 33, 17, 33, 53, 35, 17, 59, 1, 17, 33, 17, 1, 0, 2, 128, 128, 128, 255, 0, 254, 128, 128, 128, 128, 1, 0, 2, 128, 254, 0, 128, 128, 1, 0, 128, 1, 0, 255, 0, 1, 0, 0, 1, 0, 128, 0, 0, 3, 128, 2, 128, 0, 15, 0, 0, 19, 33, 21, 51, 53, 33, 21, 51, 21, 33, 53, 35, 21, 35, 17, 33, 128, 1, 0, 128, 1, 0, 128, 255, 0, 128, 128, 255, 0, 2, 128, 128, 128, 128, 128, 128, 128, 254, 128, 0, 0, 1, 0, 128, 0, 0, 3, 128, 2, 128, 0, 19, 0, 0, 1, 33, 21, 33, 21, 33, 21, 51, 21, 35, 21, 33, 53, 33, 53, 33, 53, 35, 53, 51, 1, 0, 2, 0, 254, 128, 1, 128, 128, 128, 253, 128, 2, 0, 254, 128, 128, 128, 2, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 0, 0, 1, 0, 128, 0, 0, 3, 128, 3, 0, 0, 19, 0, 0, 1, 33, 21, 33, 21, 33, 17, 51, 53, 33, 21, 35, 21, 33, 53, 35, 17, 35, 53, 51, 1, 0, 1, 0, 1, 0, 255, 0, 128, 1, 0, 128, 254, 128, 128, 128, 128, 3, 0, 128, 128, 254, 128, 128, 128, 128, 128, 1, 128, 128, 0, 1, 0, 128, 0, 0, 4, 0, 2, 128, 0, 15, 0, 0, 19, 33, 17, 33, 17, 33, 17, 51, 21, 33, 53, 35, 21, 33, 53, 35, 128, 1, 0, 1, 0, 1, 0, 128, 255, 0, 128, 254, 128, 128, 2, 128, 254, 0, 2, 0, 254, 0, 128, 128, 128, 128, 0, 1, 0, 128, 0, 0, 3, 128, 2, 128, 0, 15, 0, 0, 19, 33, 17, 33, 17, 33, 17, 35, 21, 35, 21, 33, 53, 35, 53, 35, 128, 1, 0, 1, 0, 1, 0, 128, 128, 255, 0, 128, 128, 2, 128, 254, 128, 1, 128, 254, 128, 128, 128, 128, 128, 0, 0, 1, 0, 128, 0, 0, 4, 0, 2, 128, 0, 17, 0, 0, 19, 33, 17, 51, 17, 51, 17, 51, 17, 33, 17, 33, 53, 35, 21, 33, 53, 35, 128, 1, 0, 128, 128, 128, 1, 0, 254, 128, 128, 255, 0, 128, 2, 128, 254, 0, 2, 0, 254, 0, 2, 0, 253, 128, 128, 128, 128, 0, 0, 1, 0, 128, 0, 0, 3, 128, 2, 128, 0, 19, 0, 0, 19, 33, 21, 33, 53, 33, 17, 35, 21, 51, 17, 33, 53, 33, 21, 33, 17, 51, 53, 35, 128, 1, 0, 1, 0, 1, 0, 128, 128, 255, 0, 255, 0, 255, 0, 128, 128, 2, 128, 128, 128, 255, 0, 128, 255, 0, 128, 128, 1, 0, 128, 0, 0, 1, 0, 128, 255, 128, 3, 128, 2, 128, 0, 15, 0, 0, 19, 33, 17, 33, 17, 33, 17, 35, 21, 33, 53, 33, 53, 33, 53, 35, 128, 1, 0, 1, 0, 1, 0, 128, 254, 0, 1, 128, 254, 128, 128, 2, 128, 254, 128, 1, 128, 253, 128, 128, 128, 128, 128, 0, 0, 1, 0, 128, 0, 0, 3, 128, 2, 128, 0, 19, 0, 0, 19, 33, 21, 35, 21, 35, 21, 35, 21, 33, 21, 33, 53, 51, 53, 51, 53, 51, 53, 33, 128, 3, 0, 128, 128, 128, 1, 128, 253, 0, 128, 128, 128, 254, 128, 2, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 0, 1, 0, 128, 0, 0, 3, 0, 3, 0, 0, 19, 0, 0, 1, 33, 21, 35, 21, 35, 21, 51, 17, 51, 21, 33, 53, 35, 17, 33, 53, 33, 53, 51, 2, 0, 1, 0, 128, 128, 128, 128, 255, 0, 128, 255, 0, 1, 0, 128, 3, 0, 128, 128, 128, 255, 0, 128, 128, 1, 0, 128, 128, 0, 0, 1, 1, 128, 0, 0, 2, 128, 3, 0, 0, 3, 0, 0, 1, 33, 17, 33, 1, 128, 1, 0, 255, 0, 3, 0, 253, 0, 0, 1, 1, 0, 0, 0, 3, 128, 3, 0, 0, 19, 0, 0, 1, 33, 21, 51, 21, 33, 21, 33, 17, 35, 21, 33, 53, 51, 17, 51, 53, 35, 53, 35, 1, 0, 1, 0, 128, 1, 0, 255, 0, 128, 255, 0, 128, 128, 128, 128, 3, 0, 128, 128, 128, 255, 0, 128, 128, 1, 0, 128, 128, 0, 0, 1, 0, 128, 1, 0, 3, 128, 2, 128, 0, 19, 0, 0, 1, 33, 21, 51, 21, 51, 53, 51, 21, 35, 21, 33, 53, 35, 53, 35, 21, 35, 53, 51, 1, 0, 1, 0, 128, 128, 128, 128, 255, 0, 128, 128, 128, 128, 2, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 0, 0, 3, 0, 0, 0, 0, 3, 128, 3, 128, 0, 19, 0, 31, 0, 39, 0, 0, 1, 33, 21, 51, 21, 51, 17, 35, 21, 35, 21, 33, 53, 35, 53, 35, 17, 51, 53, 51, 49, 21, 35, 17, 51, 21, 33, 53, 51, 17, 35, 53, 5, 33, 21, 33, 21, 33, 21, 33, 1, 0, 1, 128, 128, 128, 128, 128, 254, 128, 128, 128, 128, 128, 128, 128, 1, 128, 128, 128, 254, 128, 1, 128, 255, 0, 1, 0, 254, 128, 3, 128, 128, 128, 254, 128, 128, 128, 128, 128, 1, 128, 128, 128, 254, 128, 128, 128, 1, 128, 128, 128, 128, 128, 128, 0, 0, 1, 0, 128, 255, 128, 3, 128, 3, 0, 0, 21, 0, 0, 1, 33, 21, 51, 17, 51, 17, 35, 21, 33, 53, 51, 17, 35, 53, 51, 53, 33, 17, 33, 17, 51, 1, 0, 1, 128, 128, 128, 128, 255, 0, 128, 128, 128, 255, 0, 255, 0, 128, 3, 0, 128, 255, 0, 255, 0, 128, 128, 1, 0, 128, 128, 253, 0, 3, 0, 0, 3, 1, 0, 0, 0, 4, 0, 3, 128, 0, 7, 0, 21, 0, 25, 0, 0, 1, 33, 21, 51, 21, 33, 53, 35, 17, 33, 17, 51, 21, 33, 53, 35, 21, 33, 53, 35, 17, 59, 1, 17, 51, 17, 1, 128, 1, 0, 128, 255, 0, 128, 2, 0, 128, 255, 0, 128, 255, 0, 128, 128, 128, 128, 3, 128, 128, 128, 128, 255, 0, 254, 128, 128, 128, 128, 128, 1, 0, 255, 0, 1, 0, 0, 0, 3, 1, 0, 0, 0, 4, 0, 3, 128, 0, 11, 0, 25, 0, 29, 0, 0, 1, 33, 21, 51, 21, 33, 53, 35, 21, 33, 53, 51, 3, 33, 17, 51, 21, 33, 53, 35, 21, 33, 53, 35, 17, 59, 1, 17, 51, 17, 2, 0, 1, 128, 128, 255, 0, 128, 255, 0, 128, 128, 2, 0, 128, 255, 0, 128, 255, 0, 128, 128, 128, 128, 3, 128, 128, 128, 128, 128, 128, 255, 0, 254, 128, 128, 128, 128, 128, 1, 0, 255, 0, 1, 0, 0, 0, 4, 0, 128, 0, 0, 4, 0, 3, 128, 0, 3, 0, 7, 0, 21, 0, 25, 0, 0, 1, 51, 21, 35, 37, 51, 21, 35, 5, 33, 17, 51, 21, 33, 53, 35, 21, 33, 53, 35, 17, 59, 1, 17, 33, 17, 1, 128, 128, 128, 1, 0, 128, 128, 254, 128, 2, 128, 128, 255, 0, 128, 254, 128, 128, 128, 128, 1, 0, 3, 128, 128, 128, 128, 128, 254, 0, 128, 128, 128, 128, 1, 128, 254, 128, 1, 128, 0, 0, 1, 0, 128, 255, 128, 3, 0, 2, 128, 0, 17, 0, 0, 1, 33, 21, 33, 17, 33, 21, 35, 17, 33, 53, 51, 53, 33, 53, 35, 17, 51, 1, 0, 2, 0, 254, 128, 1, 128, 128, 255, 0, 128, 255, 0, 128, 128, 2, 128, 128, 255, 0, 128, 255, 0, 128, 128, 128, 1, 0, 0, 0, 3, 0, 128, 0, 0, 3, 128, 3, 128, 0, 7, 0, 21, 0, 25, 0, 0, 1, 33, 21, 51, 21, 33, 53, 35, 17, 33, 21, 51, 21, 33, 21, 33, 21, 33, 53, 35, 17, 59, 1, 21, 51, 53, 1, 0, 1, 0, 128, 255, 0, 128, 1, 128, 128, 255, 0, 1, 128, 253, 128, 128, 128, 128, 128, 3, 128, 128, 128, 128, 255, 0, 128, 128, 128, 128, 128, 1, 0, 128, 128, 0, 0, 3, 0, 128, 0, 0, 3, 128, 3, 128, 0, 7, 0, 21, 0, 25, 0, 0, 1, 33, 21, 35, 21, 33, 53, 51, 3, 33, 21, 51, 21, 33, 21, 33, 21, 33, 53, 35, 17, 59, 1, 21, 51, 53, 1, 128, 1, 0, 128, 255, 0, 128, 128, 1, 128, 128, 255, 0, 1, 128, 253, 128, 128, 128, 128, 128, 3, 128, 128, 128, 128, 255, 0, 128, 128, 128, 128, 128, 1, 0, 128, 128, 0, 3, 0, 128, 0, 0, 3, 128, 3, 128, 0, 11, 0, 25, 0, 29, 0, 0, 1, 33, 21, 51, 21, 33, 53, 35, 21, 33, 53, 51, 17, 33, 21, 51, 21, 33, 21, 33, 21, 33, 53, 35, 17, 59, 1, 21, 51, 53, 1, 0, 1, 128, 128, 255, 0, 128, 255, 0, 128, 1, 128, 128, 255, 0, 1, 128, 253, 128, 128, 128, 128, 128, 3, 128, 128, 128, 128, 128, 128, 255, 0, 128, 128, 128, 128, 128, 1, 0, 128, 128, 0, 2, 1, 0, 0, 0, 3, 128, 3, 128, 0, 11, 0, 15, 0, 0, 1, 33, 21, 51, 21, 33, 53, 35, 21, 33, 53, 51, 17, 33, 17, 33, 1, 128, 1, 128, 128, 255, 0, 128, 255, 0, 128, 1, 0, 255, 0, 3, 128, 128, 128, 128, 128, 128, 255, 0, 254, 0, 0, 3, 1, 0, 0, 0, 3, 0, 3, 128, 0, 3, 0, 7, 0, 11, 0, 0, 1, 51, 21, 35, 37, 51, 21, 35, 5, 33, 17, 33, 1, 0, 128, 128, 1, 128, 128, 128, 255, 0, 1, 0, 255, 0, 3, 128, 128, 128, 128, 128, 253, 128, 0, 3, 0, 128, 0, 0, 3, 128, 3, 128, 0, 11, 0, 23, 0, 27, 0, 0, 1, 33, 21, 51, 21, 33, 53, 35, 21, 33, 53, 51, 3, 33, 21, 51, 17, 35, 21, 33, 53, 35, 17, 59, 1, 17, 33, 17, 1, 128, 1, 128, 128, 255, 0, 128, 255, 0, 128, 128, 2, 0, 128, 128, 254, 0, 128, 128, 128, 1, 0, 3, 128, 128, 128, 128, 128, 128, 255, 0, 128, 255, 0, 128, 128, 1, 0, 255, 0, 1, 0, 0, 0, 4, 0, 128, 0, 0, 3, 128, 3, 128, 0, 3, 0, 7, 0, 19, 0, 23, 0, 0, 1, 51, 21, 35, 37, 51, 21, 35, 5, 33, 21, 51, 17, 35, 21, 33, 53, 35, 17, 59, 1, 17, 33, 17, 1, 0, 128, 128, 1, 128, 128, 128, 254, 128, 2, 0, 128, 128, 254, 0, 128, 128, 128, 1, 0, 3, 128, 128, 128, 128, 128, 128, 254, 128, 128, 128, 1, 128, 254, 128, 1, 128, 0, 2, 0, 128, 0, 0, 4, 0, 3, 128, 0, 7, 0, 23, 0, 0, 1, 33, 21, 51, 21, 33, 53, 35, 3, 33, 17, 33, 17, 33, 17, 51, 21, 33, 53, 35, 21, 33, 53, 35, 1, 0, 1, 0, 128, 255, 0, 128, 128, 1, 0, 1, 0, 1, 0, 128, 255, 0, 128, 254, 128, 128, 3, 128, 128, 128, 128, 255, 0, 254, 128, 1, 128, 254, 128, 128, 128, 128, 128, 0, 0, 2, 0, 128, 0, 0, 4, 0, 3, 128, 0, 11, 0, 27, 0, 0, 1, 33, 21, 51, 21, 33, 53, 35, 21, 33, 53, 51, 1, 33, 17, 33, 17, 33, 17, 51, 21, 33, 53, 35, 21, 33, 53, 35, 1, 128, 1, 128, 128, 255, 0, 128, 255, 0, 128, 255, 0, 1, 0, 1, 0, 1, 0, 128, 255, 0, 128, 254, 128, 128, 3, 128, 128, 128, 128, 128, 128, 255, 0, 254, 128, 1, 128, 254, 128, 128, 128, 128, 128, 0, 0, 3, 0, 128, 0, 0, 3, 128, 3, 128, 0, 3, 0, 7, 0, 17, 0, 0, 1, 51, 21, 35, 37, 51, 21, 35, 5, 33, 17, 33, 17, 33, 17, 33, 53, 35, 1, 0, 128, 128, 1, 128, 128, 128, 254, 0, 1, 0, 1, 0, 1, 0, 253, 128, 128, 3, 128, 128, 128, 128, 128, 254, 0, 2, 0, 253, 128, 128, 0, 2, 0, 0, 1, 0, 4, 0, 3, 128, 0, 7, 0, 19, 0, 0, 17, 33, 21, 35, 17, 35, 17, 35, 5, 33, 17, 35, 17, 35, 21, 35, 53, 35, 17, 35, 1, 128, 128, 128, 128, 1, 128, 2, 128, 128, 128, 128, 128, 128, 3, 128, 128, 255, 0, 1, 0, 128, 254, 128, 1, 0, 128, 128, 255, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 12, 0, 0, 0, 34, 0, 0, 0, 2, 0, 3, 0, 0, 0, 125, 0, 1, 0, 126, 0, 128, 0, 2, 0, 129, 0, 218, 0, 1, 0, 4, 0, 0, 0, 2, 0, 0, 0, 0, 0, 1, 0, 0, 0, 10, 0, 30, 0, 44, 0, 1, 68, 70, 76, 84, 0, 8, 0, 4, 0, 0, 0, 0, 255, 255, 0, 1, 0, 0, 0, 1, 102, 114, 97, 99, 0, 8, 0, 0, 0, 1, 0, 0, 0, 1, 0, 4, 0, 4, 0, 0, 0, 1, 0, 8, 0, 1, 0, 44, 0, 2, 0, 10, 0, 32, 0, 2, 0, 6, 0, 14, 0, 127, 0, 3, 0, 18, 0, 21, 0, 126, 0, 3, 0, 18, 0, 23, 0, 1, 0, 4, 0, 128, 0, 3, 0, 18, 0, 23, 0, 1, 0, 2, 0, 20, 0, 22, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 219, 0, 0, 0, 1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 6, 0, 7, 0, 8, 0, 9, 0, 10, 0, 11, 0, 12, 0, 13, 0, 14, 0, 15, 0, 16, 0, 17, 0, 18, 0, 19, 0, 20, 0, 21, 0, 22, 0, 23, 0, 24, 0, 25, 0, 26, 0, 27, 0, 28, 0, 29, 0, 30, 0, 31, 0, 32, 0, 33, 0, 34, 0, 35, 0, 36, 0, 37, 0, 38, 0, 39, 0, 40, 0, 41, 0, 42, 0, 43, 0, 44, 0, 45, 0, 46, 0, 47, 0, 48, 0, 49, 0, 50, 0, 51, 0, 52, 0, 53, 0, 54, 0, 55, 0, 56, 0, 57, 0, 58, 0, 59, 0, 60, 0, 61, 0, 62, 0, 63, 0, 64, 0, 65, 0, 66, 0, 67, 0, 68, 0, 69, 0, 70, 0, 71, 0, 72, 0, 73, 0, 74, 0, 75, 0, 76, 0, 77, 0, 78, 0, 79, 0, 80, 0, 81, 0, 82, 0, 83, 0, 84, 0, 85, 0, 86, 0, 87, 0, 88, 0, 89, 0, 90, 0, 91, 0, 92, 0, 93, 0, 94, 0, 95, 0, 96, 0, 97, 0, 172, 0, 163, 0, 132, 0, 133, 0, 189, 0, 150, 0, 232, 0, 134, 0, 142, 0, 139, 0, 157, 0, 169, 0, 164, 1, 2, 0, 138, 0, 218, 0, 131, 0, 147, 0, 242, 0, 243, 0, 141, 0, 151, 0, 136, 0, 195, 0, 222, 0, 241, 0, 158, 0, 170, 0, 245, 0, 244, 0, 246, 0, 162, 0, 173, 0, 201, 0, 199, 0, 174, 0, 98, 0, 99, 0, 144, 0, 100, 0, 203, 0, 101, 0, 200, 0, 202, 0, 207, 0, 204, 0, 205, 0, 206, 0, 233, 0, 102, 0, 211, 0, 208, 0, 209, 0, 175, 0, 103, 0, 240, 0, 145, 0, 214, 0, 212, 0, 213, 0, 104, 0, 235, 0, 237, 0, 137, 0, 106, 0, 105, 0, 107, 0, 109, 0, 108, 0, 110, 0, 160, 0, 111, 0, 113, 0, 112, 0, 114, 0, 115, 0, 117, 0, 116, 0, 118, 0, 119, 0, 234, 0, 120, 0, 122, 0, 121, 0, 123, 0, 125, 0, 124, 0, 184, 0, 161, 0, 127, 0, 126, 0, 128, 0, 129, 0, 236, 0, 238, 0, 186, 0, 176, 0, 177, 0, 228, 0, 229, 0, 187, 0, 166, 0, 216, 0, 217, 0, 178, 0, 179, 0, 182, 0, 183, 0, 196, 0, 180, 0, 181, 0, 197, 0, 130, 0, 194, 0, 135, 0, 171, 0, 198, 0, 190, 0, 191, 1, 3, 0, 140, 10, 115, 111, 102, 116, 104, 121, 112, 104, 101, 110, 4, 69, 117, 114, 111}}
