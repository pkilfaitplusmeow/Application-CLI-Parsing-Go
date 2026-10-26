package src

func CorrectOccurence(fileText string, i int, occurence int, fileSize int) []byte {
	i--
	pow := 0
	var hexa int = 0
	bin := 0
	var cestjustepourlereturn []byte // temporaire

	//dans le cas 1 (hex) retourne en arriere jusqu'au dernier espace et convertie en decimale
	if occurence == 1 {
		for fileText[i] != ' ' && !IsPonctuation([]byte(fileText), i) {
			hexa += HexaCoverter(fileText, i, pow)
			pow++ //puissance du 16
			i--
		}
		deci_Val := DecimalToString(hexa)
		deci_Val[0] = fileText[i]
		return deci_Val
	}
	//dans le cas 2 (bin) ...
	if occurence == 2 {
		for fileText[i] != ' ' && !IsPonctuation([]byte(fileText), i) {
			bin += BinConverter(fileText, i, pow)
			pow++ //puissance du 2
			i--
		}
		deci_Val := DecimalToString(bin)
		deci_Val[0] = fileText[i]
		return deci_Val
	}
	//dans le cas 3 (up)
	if occurence == 3 {
		j := 0
		k := i
		for fileText[k] != ' ' && !IsPonctuation([]byte(fileText), k) {
			k--
		}

		str := make([]byte, i-(k-1)) // size error

		if IsPonctuation([]byte(fileText), k) {
			str[j] = fileText[k]
		}
		k++
		j++

		for fileText[k] != ' ' {
			if fileText[k] >= 'a' && fileText[k] <= 'z' {
				str[j] = fileText[k] - 32
				j++
				k++
			} else {
				str[j] = fileText[k]
				j++
				k++
			}
		}
		return str
	}

	//dans le cas 4 (low)
	if occurence == 4 {
		j := 0
		k := i
		for fileText[k] != ' ' && !IsPonctuation([]byte(fileText), k) {
			k--
		}

		str := make([]byte, i-(k-1)) // size error

		if IsPonctuation([]byte(fileText), k) {
			str[j] = fileText[k]
		}
		k++
		j++

		for fileText[k] != ' ' {
			if fileText[k] >= 'A' && fileText[k] <= 'Z' {
				str[j] = fileText[k] + 32
				j++
				k++
			} else {
				str[j] = fileText[k]
				j++
				k++
			}
		}
		return str
	}

	//dans le cas 5 (cap)
	if occurence == 5 {
		j := 0
		k := i

		for fileText[k] != ' ' && !IsPonctuation([]byte(fileText), k) {
			k--
		}
		str := make([]byte, i-(k-1))

		if IsPonctuation([]byte(fileText), k) {
			str[j] = fileText[k]
		}
		k++
		j++

		if fileText[k+1] >= 'a' && fileText[k+1] <= 'z' {
			str[j] = fileText[k] - 32 // passage en maj de la lettre a l'indice k + 1 (apres l'espace)
		}

		j++
		k++
		for fileText[k] != ' ' {
			str[j] = fileText[k] //transfer le reste des lettres
			k++
			j++
		}
		return str
	}
	return cestjustepourlereturn // return d'erreur (todo)
}
