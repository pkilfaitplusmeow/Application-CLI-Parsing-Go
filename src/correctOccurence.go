package src

func CorrectOccurence(fileText string, i int, occurence int, fileSize int) []byte {
	i--
	pow := 0
	var hexa int = 0
	bin := 0
	var cestjustepourlereturn []byte // temporaire

	//dans le cas 1 (hex) retourne en arriere jusqu'au dernier espace et convertie en decimale
	if occurence == 1 {
		for fileText[i] != ' ' {
			hexa += HexaCoverter(fileText, i, pow)
			pow++ //puissance du 16
			i--
		}
		deci_Val := DecimalToString(hexa)
		return deci_Val
	}
	//dans le cas 2 (dec) ...
	if occurence == 2 {
		for fileText[i] != ' ' {
			bin += BinConverter(fileText, i, pow)
			pow++ //puissance du 2
			i--
		}
		deci_Val := DecimalToString(bin)
		return deci_Val
	}
	return cestjustepourlereturn
}
