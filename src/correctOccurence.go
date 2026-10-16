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
	//dans le cas 2 (bin) ...
	if occurence == 2 {
		for fileText[i] != ' ' {
			bin += BinConverter(fileText, i, pow)
			pow++ //puissance du 2
			i--
		}
		deci_Val := DecimalToString(bin)
		return deci_Val
	}
	//dans le cas 3 (up)
	if occurence == 3 {
		j := -1
		k := i
		for fileText[k] != ' ' {
			k--
			j++
		}

		str := make([]byte, i-k)
		for fileText[i] != ' ' {
			if fileText[i] >= 'a' && fileText[i] <= 'z' {
				str[j] = fileText[i] - 32
				j--
				i--
			} else {
				str[j] = fileText[i]
				j--
				i--
			}
		}
		return str
	}

	//dans le cas 4 (low)
	if occurence == 4 {
		j := -1
		k := i
		for fileText[k] != ' ' {
			k--
			j++
		}

		str := make([]byte, i-k)
		for fileText[i] != ' ' {
			if fileText[i] >= 'A' && fileText[i] <= 'Z' {
				str[j] = fileText[i] + 32
				j--
				i--
			} else {
				str[j] = fileText[i]
				j--
				i--
			}
		}
		return str
	}
	return cestjustepourlereturn
}
