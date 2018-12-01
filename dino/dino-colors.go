package dino

import "image/color"

type colorGroup [5]color.RGBA

var dinoPalette = []colorGroup{
	colorGroup{
		color.RGBA{83, 13, 14, 255},
		color.RGBA{184, 166, 169, 255},
		color.RGBA{233, 214, 192, 255},
		color.RGBA{199, 147, 99, 255},
		color.RGBA{59, 96, 60, 255},
	},
	colorGroup{
		color.RGBA{136, 74, 53, 255},
		color.RGBA{100, 107, 95, 255},
		color.RGBA{124, 139, 137, 255},
		color.RGBA{79, 120, 131, 255},
		color.RGBA{83, 164, 193, 255},
	},
	colorGroup{
		color.RGBA{208, 201, 189, 255},
		color.RGBA{178, 168, 165, 255},
		color.RGBA{130, 144, 148, 255},
		color.RGBA{44, 62, 60, 255},
		color.RGBA{54, 44, 44, 255},
	},
	colorGroup{
		color.RGBA{4, 49, 75, 255},
		color.RGBA{4, 139, 138, 255},
		color.RGBA{179, 181, 118, 255},
		color.RGBA{246, 191, 35, 255},
		color.RGBA{222, 119, 92, 255},
	},
	colorGroup{
		color.RGBA{55, 68, 87, 255},
		color.RGBA{139, 158, 167, 255},
		color.RGBA{185, 222, 219, 255},
		color.RGBA{140, 140, 121, 255},
		color.RGBA{241, 201, 139, 255},
	},
	colorGroup{
		color.RGBA{246, 239, 187, 255},
		color.RGBA{221, 166, 150, 255},
		color.RGBA{167, 56, 104, 255},
		color.RGBA{85, 60, 73, 255},
		color.RGBA{72, 35, 44, 255},
	},
	colorGroup{
		color.RGBA{106, 28, 48, 255},
		color.RGBA{252, 67, 73, 255},
		color.RGBA{217, 218, 220, 255},
		color.RGBA{148, 220, 224, 255},
		color.RGBA{121, 163, 139, 255},
	},
	colorGroup{
		color.RGBA{140, 87, 131, 255},
		color.RGBA{119, 144, 123, 255},
		color.RGBA{140, 202, 168, 255},
		color.RGBA{245, 229, 159, 255},
		color.RGBA{231, 139, 102, 255},
	},
	colorGroup{
		color.RGBA{141, 68, 116, 255},
		color.RGBA{122, 59, 101, 255},
		color.RGBA{168, 107, 74, 255},
		color.RGBA{213, 166, 112, 255},
		color.RGBA{163, 69, 38, 255},
	},
	colorGroup{
		color.RGBA{46, 45, 37, 255},
		color.RGBA{103, 59, 56, 255},
		color.RGBA{248, 95, 40, 255},
		color.RGBA{241, 194, 37, 255},
		color.RGBA{211, 191, 123, 255},
	},
	colorGroup{
		color.RGBA{81, 82, 90, 255},
		color.RGBA{133, 136, 140, 255},
		color.RGBA{197, 167, 158, 255},
		color.RGBA{243, 225, 196, 255},
		color.RGBA{211, 140, 135, 255},
	},
	colorGroup{
		color.RGBA{67, 39, 82, 255},
		color.RGBA{159, 37, 89, 255},
		color.RGBA{240, 194, 56, 255},
		color.RGBA{135, 214, 176, 255},
		color.RGBA{15, 152, 145, 255},
	},
	colorGroup{
		color.RGBA{50, 51, 67, 255},
		color.RGBA{82, 75, 73, 255},
		color.RGBA{97, 153, 128, 255},
		color.RGBA{180, 208, 147, 255},
		color.RGBA{249, 246, 214, 255},
	},
	colorGroup{
		color.RGBA{58, 48, 57, 255},
		color.RGBA{143, 106, 103, 255},
		color.RGBA{244, 242, 242, 255},
		color.RGBA{193, 215, 144, 255},
		color.RGBA{155, 188, 40, 255},
	},
	colorGroup{
		color.RGBA{245, 65, 100, 255},
		color.RGBA{212, 115, 70, 255},
		color.RGBA{202, 195, 156, 255},
		color.RGBA{200, 204, 173, 255},
		color.RGBA{159, 194, 186, 255},
	},
	colorGroup{
		color.RGBA{60, 62, 66, 255},
		color.RGBA{160, 168, 168, 255},
		color.RGBA{216, 183, 145, 255},
		color.RGBA{124, 84, 79, 255},
		color.RGBA{45, 35, 46, 255},
	},
	colorGroup{
		color.RGBA{25, 37, 49, 255},
		color.RGBA{43, 97, 106, 255},
		color.RGBA{144, 167, 154, 255},
		color.RGBA{241, 223, 192, 255},
		color.RGBA{230, 118, 106, 255},
	},
	colorGroup{
		color.RGBA{82, 41, 24, 255},
		color.RGBA{193, 147, 86, 255},
		color.RGBA{198, 164, 138, 255},
		color.RGBA{92, 105, 94, 255},
		color.RGBA{56, 57, 87, 255},
	},
	colorGroup{
		color.RGBA{35, 44, 59, 255},
		color.RGBA{56, 60, 83, 255},
		color.RGBA{134, 193, 167, 255},
		color.RGBA{118, 184, 67, 255},
		color.RGBA{254, 244, 78, 255},
	},
	colorGroup{
		color.RGBA{34, 21, 32, 255},
		color.RGBA{132, 21, 36, 255},
		color.RGBA{179, 98, 83, 255},
		color.RGBA{166, 151, 111, 255},
		color.RGBA{211, 200, 117, 255},
	},
	colorGroup{
		color.RGBA{213, 55, 48, 255},
		color.RGBA{242, 246, 211, 255},
		color.RGBA{184, 208, 148, 255},
		color.RGBA{87, 144, 102, 255},
		color.RGBA{54, 70, 79, 255},
	},
	colorGroup{
		color.RGBA{182, 148, 139, 255},
		color.RGBA{236, 216, 142, 255},
		color.RGBA{188, 149, 94, 255},
		color.RGBA{107, 129, 81, 255},
		color.RGBA{66, 58, 63, 255},
	},
	colorGroup{
		color.RGBA{63, 30, 13, 255},
		color.RGBA{189, 105, 30, 255},
		color.RGBA{114, 111, 53, 255},
		color.RGBA{92, 80, 19, 255},
		color.RGBA{84, 21, 14, 255},
	},
	colorGroup{
		color.RGBA{48, 65, 72, 255},
		color.RGBA{78, 172, 166, 255},
		color.RGBA{71, 199, 217, 255},
		color.RGBA{237, 222, 182, 255},
		color.RGBA{245, 90, 48, 255},
	},
	colorGroup{
		color.RGBA{170, 196, 114, 255},
		color.RGBA{87, 198, 119, 255},
		color.RGBA{31, 189, 155, 255},
		color.RGBA{226, 36, 62, 255},
		color.RGBA{185, 51, 97, 255},
	},
	colorGroup{
		color.RGBA{242, 152, 35, 255},
		color.RGBA{220, 227, 189, 255},
		color.RGBA{48, 114, 108, 255},
		color.RGBA{19, 69, 82, 255},
		color.RGBA{14, 17, 17, 255},
	},
	colorGroup{
		color.RGBA{6, 86, 143, 255},
		color.RGBA{10, 134, 176, 255},
		color.RGBA{66, 158, 167, 255},
		color.RGBA{161, 207, 159, 255},
		color.RGBA{221, 211, 124, 255},
	},
	colorGroup{
		color.RGBA{252, 10, 13, 255},
		color.RGBA{216, 91, 51, 255},
		color.RGBA{209, 142, 37, 255},
		color.RGBA{248, 191, 9, 255},
		color.RGBA{55, 108, 47, 255},
	},
	colorGroup{
		color.RGBA{38, 61, 95, 255},
		color.RGBA{27, 183, 201, 255},
		color.RGBA{207, 210, 202, 255},
		color.RGBA{192, 217, 195, 255},
		color.RGBA{244, 196, 137, 255},
	},
	colorGroup{
		color.RGBA{220, 201, 72, 255},
		color.RGBA{215, 143, 68, 255},
		color.RGBA{116, 100, 101, 255},
		color.RGBA{65, 125, 108, 255},
		color.RGBA{49, 76, 83, 255},
	},
	colorGroup{
		color.RGBA{19, 27, 27, 255},
		color.RGBA{46, 29, 20, 255},
		color.RGBA{128, 110, 55, 255},
		color.RGBA{216, 187, 102, 255},
		color.RGBA{254, 253, 252, 255},
	},
	colorGroup{
		color.RGBA{41, 53, 54, 255},
		color.RGBA{104, 154, 133, 255},
		color.RGBA{217, 205, 174, 255},
		color.RGBA{253, 251, 245, 255},
		color.RGBA{108, 40, 32, 255},
	},
	colorGroup{
		color.RGBA{72, 50, 53, 255},
		color.RGBA{160, 140, 142, 255},
		color.RGBA{197, 191, 170, 255},
		color.RGBA{217, 189, 180, 255},
		color.RGBA{192, 72, 73, 255},
	},
	colorGroup{
		color.RGBA{86, 69, 60, 255},
		color.RGBA{121, 138, 134, 255},
		color.RGBA{219, 242, 236, 255},
		color.RGBA{220, 229, 207, 255},
		color.RGBA{207, 240, 209, 255},
	},
	colorGroup{
		color.RGBA{37, 26, 63, 255},
		color.RGBA{156, 97, 10, 255},
		color.RGBA{253, 121, 11, 255},
		color.RGBA{245, 171, 12, 255},
		color.RGBA{173, 221, 109, 255},
	},
	colorGroup{
		color.RGBA{43, 43, 45, 255},
		color.RGBA{107, 177, 89, 255},
		color.RGBA{188, 144, 89, 255},
		color.RGBA{227, 197, 74, 255},
		color.RGBA{209, 92, 50, 255},
	},
	colorGroup{
		color.RGBA{22, 45, 63, 255},
		color.RGBA{37, 111, 91, 255},
		color.RGBA{61, 143, 118, 255},
		color.RGBA{179, 197, 137, 255},
		color.RGBA{212, 222, 155, 255},
	},
	colorGroup{
		color.RGBA{234, 220, 115, 255},
		color.RGBA{203, 181, 85, 255},
		color.RGBA{83, 88, 49, 255},
		color.RGBA{35, 43, 41, 255},
		color.RGBA{22, 32, 37, 255},
	},
	colorGroup{
		color.RGBA{91, 8, 33, 255},
		color.RGBA{118, 50, 60, 255},
		color.RGBA{167, 139, 120, 255},
		color.RGBA{171, 151, 114, 255},
		color.RGBA{159, 137, 92, 255},
	},
	colorGroup{
		color.RGBA{239, 188, 151, 255},
		color.RGBA{54, 175, 238, 255},
		color.RGBA{139, 210, 194, 255},
		color.RGBA{189, 209, 101, 255},
		color.RGBA{80, 102, 39, 255},
	},
	colorGroup{
		color.RGBA{49, 115, 127, 255},
		color.RGBA{125, 199, 180, 255},
		color.RGBA{164, 248, 249, 255},
		color.RGBA{237, 200, 126, 255},
		color.RGBA{205, 82, 87, 255},
	},
	colorGroup{
		color.RGBA{194, 203, 136, 255},
		color.RGBA{140, 192, 149, 255},
		color.RGBA{42, 107, 124, 255},
		color.RGBA{83, 59, 28, 255},
		color.RGBA{117, 83, 38, 255},
	},
	colorGroup{
		color.RGBA{252, 117, 113, 255},
		color.RGBA{248, 208, 148, 255},
		color.RGBA{170, 134, 54, 255},
		color.RGBA{61, 59, 56, 255},
		color.RGBA{94, 211, 183, 255},
	},
	colorGroup{
		color.RGBA{15, 28, 48, 255},
		color.RGBA{15, 43, 66, 255},
		color.RGBA{77, 110, 124, 255},
		color.RGBA{192, 179, 153, 255},
		color.RGBA{245, 238, 211, 255},
	},
	colorGroup{
		color.RGBA{51, 63, 67, 255},
		color.RGBA{48, 65, 97, 255},
		color.RGBA{90, 168, 169, 255},
		color.RGBA{71, 165, 170, 255},
		color.RGBA{255, 255, 253, 255},
	},
	colorGroup{
		color.RGBA{58, 50, 65, 255},
		color.RGBA{171, 61, 94, 255},
		color.RGBA{240, 113, 97, 255},
		color.RGBA{240, 177, 78, 255},
		color.RGBA{177, 199, 91, 255},
	},
	colorGroup{
		color.RGBA{59, 51, 75, 255},
		color.RGBA{253, 251, 243, 255},
		color.RGBA{194, 239, 232, 255},
		color.RGBA{139, 179, 166, 255},
		color.RGBA{113, 115, 107, 255},
	},
	colorGroup{
		color.RGBA{29, 45, 48, 255},
		color.RGBA{45, 129, 161, 255},
		color.RGBA{176, 205, 214, 255},
		color.RGBA{245, 227, 208, 255},
		color.RGBA{231, 118, 96, 255},
	},
	colorGroup{
		color.RGBA{129, 125, 124, 255},
		color.RGBA{192, 186, 185, 255},
		color.RGBA{253, 249, 246, 255},
		color.RGBA{183, 190, 138, 255},
		color.RGBA{144, 119, 68, 255},
	},
	colorGroup{
		color.RGBA{44, 95, 98, 255},
		color.RGBA{46, 114, 126, 255},
		color.RGBA{62, 66, 63, 255},
		color.RGBA{254, 249, 240, 255},
		color.RGBA{156, 84, 82, 255},
	},
	colorGroup{
		color.RGBA{232, 182, 61, 255},
		color.RGBA{229, 194, 31, 255},
		color.RGBA{120, 181, 99, 255},
		color.RGBA{14, 63, 79, 255},
		color.RGBA{20, 32, 31, 255},
	},
	colorGroup{
		color.RGBA{23, 27, 89, 255},
		color.RGBA{5, 113, 156, 255},
		color.RGBA{169, 231, 89, 255},
		color.RGBA{249, 234, 169, 255},
		color.RGBA{118, 119, 120, 255},
	},
	colorGroup{
		color.RGBA{234, 42, 30, 255},
		color.RGBA{252, 214, 78, 255},
		color.RGBA{186, 201, 136, 255},
		color.RGBA{85, 123, 103, 255},
		color.RGBA{40, 42, 48, 255},
	},
	colorGroup{
		color.RGBA{178, 54, 51, 255},
		color.RGBA{252, 250, 187, 255},
		color.RGBA{133, 181, 119, 255},
		color.RGBA{41, 93, 86, 255},
		color.RGBA{31, 43, 57, 255},
	},
	colorGroup{
		color.RGBA{10, 17, 64, 255},
		color.RGBA{11, 40, 102, 255},
		color.RGBA{8, 14, 241, 255},
		color.RGBA{202, 225, 230, 255},
		color.RGBA{199, 181, 20, 255},
	},
	colorGroup{
		color.RGBA{217, 104, 49, 255},
		color.RGBA{239, 191, 83, 255},
		color.RGBA{160, 166, 91, 255},
		color.RGBA{84, 153, 136, 255},
		color.RGBA{70, 42, 68, 255},
	},
	colorGroup{
		color.RGBA{182, 58, 64, 255},
		color.RGBA{95, 169, 162, 255},
		color.RGBA{141, 215, 183, 255},
		color.RGBA{18, 211, 158, 255},
		color.RGBA{95, 200, 121, 255},
	},
	colorGroup{
		color.RGBA{25, 39, 49, 255},
		color.RGBA{50, 72, 88, 255},
		color.RGBA{72, 95, 116, 255},
		color.RGBA{144, 183, 196, 255},
		color.RGBA{220, 226, 223, 255},
	},
	colorGroup{
		color.RGBA{175, 219, 77, 255},
		color.RGBA{188, 183, 111, 255},
		color.RGBA{118, 134, 145, 255},
		color.RGBA{99, 94, 142, 255},
		color.RGBA{63, 60, 85, 255},
	},
	colorGroup{
		color.RGBA{38, 81, 60, 255},
		color.RGBA{91, 114, 144, 255},
		color.RGBA{213, 225, 139, 255},
		color.RGBA{201, 180, 143, 255},
		color.RGBA{139, 119, 96, 255},
	},
	colorGroup{
		color.RGBA{191, 205, 92, 255},
		color.RGBA{244, 186, 39, 255},
		color.RGBA{175, 185, 92, 255},
		color.RGBA{153, 195, 165, 255},
		color.RGBA{74, 199, 208, 255},
	},
	colorGroup{
		color.RGBA{39, 38, 50, 255},
		color.RGBA{53, 65, 73, 255},
		color.RGBA{92, 94, 88, 255},
		color.RGBA{138, 167, 147, 255},
		color.RGBA{242, 228, 178, 255},
	},
	colorGroup{
		color.RGBA{169, 83, 81, 255},
		color.RGBA{190, 157, 109, 255},
		color.RGBA{205, 212, 171, 255},
		color.RGBA{98, 188, 154, 255},
		color.RGBA{72, 147, 136, 255},
	},
	colorGroup{
		color.RGBA{69, 71, 78, 255},
		color.RGBA{169, 168, 157, 255},
		color.RGBA{213, 222, 239, 255},
		color.RGBA{144, 188, 217, 255},
		color.RGBA{74, 86, 126, 255},
	},
	colorGroup{
		color.RGBA{11, 41, 36, 255},
		color.RGBA{30, 130, 101, 255},
		color.RGBA{249, 222, 167, 255},
		color.RGBA{113, 10, 14, 255},
		color.RGBA{71, 21, 34, 255},
	},
	colorGroup{
		color.RGBA{169, 227, 131, 255},
		color.RGBA{168, 203, 75, 255},
		color.RGBA{126, 149, 103, 255},
		color.RGBA{62, 102, 84, 255},
		color.RGBA{51, 53, 46, 255},
	},
	colorGroup{
		color.RGBA{241, 232, 181, 255},
		color.RGBA{221, 180, 134, 255},
		color.RGBA{152, 120, 79, 255},
		color.RGBA{87, 86, 57, 255},
		color.RGBA{61, 46, 31, 255},
	},
	colorGroup{
		color.RGBA{78, 55, 53, 255},
		color.RGBA{134, 148, 142, 255},
		color.RGBA{138, 184, 142, 255},
		color.RGBA{220, 196, 172, 255},
		color.RGBA{210, 83, 73, 255},
	},
	colorGroup{
		color.RGBA{35, 49, 58, 255},
		color.RGBA{66, 160, 156, 255},
		color.RGBA{131, 165, 147, 255},
		color.RGBA{222, 216, 127, 255},
		color.RGBA{158, 179, 137, 255},
	},
	colorGroup{
		color.RGBA{50, 95, 99, 255},
		color.RGBA{142, 190, 178, 255},
		color.RGBA{233, 228, 185, 255},
		color.RGBA{246, 182, 98, 255},
		color.RGBA{243, 102, 67, 255},
	},
	colorGroup{
		color.RGBA{169, 70, 64, 255},
		color.RGBA{154, 176, 138, 255},
		color.RGBA{252, 249, 217, 255},
		color.RGBA{102, 151, 119, 255},
		color.RGBA{25, 35, 38, 255},
	},
	colorGroup{
		color.RGBA{76, 62, 71, 255},
		color.RGBA{163, 154, 140, 255},
		color.RGBA{174, 184, 104, 255},
		color.RGBA{230, 215, 178, 255},
		color.RGBA{219, 209, 141, 255},
	},
	colorGroup{
		color.RGBA{19, 92, 71, 255},
		color.RGBA{111, 120, 120, 255},
		color.RGBA{228, 212, 80, 255},
		color.RGBA{247, 187, 143, 255},
		color.RGBA{191, 14, 91, 255},
	},
	colorGroup{
		color.RGBA{204, 101, 33, 255},
		color.RGBA{220, 138, 71, 255},
		color.RGBA{148, 156, 142, 255},
		color.RGBA{153, 175, 153, 255},
		color.RGBA{125, 127, 68, 255},
	},
	colorGroup{
		color.RGBA{29, 32, 50, 255},
		color.RGBA{27, 69, 70, 255},
		color.RGBA{115, 100, 95, 255},
		color.RGBA{244, 134, 107, 255},
		color.RGBA{243, 179, 163, 255},
	},
	colorGroup{
		color.RGBA{14, 40, 55, 255},
		color.RGBA{139, 141, 127, 255},
		color.RGBA{237, 201, 157, 255},
		color.RGBA{242, 183, 76, 255},
		color.RGBA{194, 182, 106, 255},
	},
	colorGroup{
		color.RGBA{65, 68, 57, 255},
		color.RGBA{145, 161, 125, 255},
		color.RGBA{240, 212, 144, 255},
		color.RGBA{212, 157, 106, 255},
		color.RGBA{75, 69, 54, 255},
	},
	colorGroup{
		color.RGBA{58, 63, 62, 255},
		color.RGBA{129, 143, 98, 255},
		color.RGBA{208, 209, 150, 255},
		color.RGBA{201, 205, 140, 255},
		color.RGBA{178, 171, 135, 255},
	},
}
