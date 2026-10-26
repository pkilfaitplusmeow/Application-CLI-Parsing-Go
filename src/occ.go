package src

func Occ(str string, i int) (int, int) {
	i++
	var b int
	/*
		regarde les 2 premiere lettre apres une parenthese
		2 lettres pour eviter les erreurs
	*/
	k := i
	for str[k] != ')' {
		if str[k] == ',' {
			for str[k+2] >= '0' && str[k+2] <= '9' {
				b = (b * 10) + int(str[k+2]) - 48
				k++
			}
		}
		k++
	}
	if b == 0 {
		b = 1
	}
	switch {
	case str[i] == 'h' && str[i+1] == 'e': //hexa
		return 1, b
	case str[i] == 'b' && str[i+1] == 'i': //binaire
		return 2, b
	case str[i] == 'u' && str[i+1] == 'p': //uppercase
		return 3, b
	case str[i] == 'l' && str[i+1] == 'o': //lowercase
		return 4, b
	case str[i] == 'c' && str[i+1] == 'a': //cap
		return 5, b
	default:
		return -1, b
	}
}
