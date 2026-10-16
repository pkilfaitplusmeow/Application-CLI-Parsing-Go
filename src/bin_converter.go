package src

func BinConverter(fileText string, i int, pow int) int {
	var bin int = 0
	if (fileText[i] == '0') || (fileText[i] == '1') {
		a := fileText[i] - '0'
		bin += int(a) * Pow(2, pow)
	}
	return bin
}
