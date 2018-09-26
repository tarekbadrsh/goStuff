package numbercompression

//Ref https://stackoverflow.com/questions/5901153/compress-large-integers-into-smallest-possible-string

var defaultEncodeDic = map[int64]string{
	0:  "A",
	1:  "B",
	2:  "C",
	3:  "D",
	4:  "E",
	5:  "F",
	6:  "G",
	7:  "H",
	8:  "J",
	9:  "K",
	10: "M",
	11: "N",
	12: "P",
	13: "Q",
	14: "R",
	15: "S",
	16: "T",
	17: "U",
	18: "V",
	19: "W",
	20: "X",
	21: "Y",
	22: "Z",
	23: "a",
	24: "b",
	25: "c",
	26: "d",
	27: "e",
	28: "f",
	29: "g",
	30: "h",
	31: "j",
	32: "k",
	33: "m",
	34: "n",
	35: "p",
	36: "q",
	37: "r",
	38: "s",
	39: "t",
	40: "u",
	41: "v",
	42: "x",
	43: "y",
	44: "z",
	45: "2",
	46: "3",
	47: "4",
}

var defaultDecodeDic = map[string]int64{
	"A": 0,
	"B": 1,
	"C": 2,
	"D": 3,
	"E": 4,
	"F": 5,
	"G": 6,
	"H": 7,
	"J": 8,
	"K": 9,
	"M": 10,
	"N": 11,
	"P": 12,
	"Q": 13,
	"R": 14,
	"S": 15,
	"T": 16,
	"U": 17,
	"V": 18,
	"W": 19,
	"X": 20,
	"Y": 21,
	"Z": 22,
	"a": 23,
	"b": 24,
	"c": 25,
	"d": 26,
	"e": 27,
	"f": 28,
	"g": 29,
	"h": 30,
	"j": 31,
	"k": 32,
	"m": 33,
	"n": 34,
	"p": 35,
	"q": 36,
	"r": 37,
	"s": 38,
	"t": 39,
	"u": 40,
	"v": 41,
	"x": 42,
	"y": 43,
	"z": 44,
	"2": 45,
	"3": 46,
	"4": 47,
}

// CompresNubmer :
func CompresNubmer(input int64, dic map[int64]string) string {
	b := int64(len(dic))
	res := ""
	if input == 0 {
		return dic[0]
	}
	for input > 0 {
		val := input % b
		input = input / b
		res += dic[val]
	}
	return res
}

// UncompresNubmer :
func UncompresNubmer(encoded string, dic map[string]int64) int64 {
	b := int64(len(dic))
	res := int64(0)
	for i := int64(len(encoded)) - 1; i >= 0; i-- {
		ch := string([]rune(encoded)[i])
		val := dic[ch]
		res = (res * b) + val
	}
	return res
}

// CompresNubmerDefault :
func CompresNubmerDefault(input int64) string {
	return CompresNubmer(input, defaultEncodeDic)
}

// UncompresNubmerDefault :
func UncompresNubmerDefault(input string) int64 {
	return UncompresNubmer(input, defaultDecodeDic)
}
