package src

func Occ(str string, i int) int {
	i++
	/*
		regarde les 2 premiere lettre apres une parenthese
		2 lettres pour eviter les erreurs
	*/
	switch {
	case str[i] == 'h' && str[i+1] == 'e': //hexa
		return 1
	case str[i] == 'b' && str[i+1] == 'i': //binaire
		return 2
	case str[i] == 'u' && str[i+1] == 'p': //uppercase
		return 3
	case str[i] == 'l' && str[i+1] == 'o': //lowercase
		return 4
	case str[i] == 'c' && str[i+1] == 'a': //cap
		return 5
	default:
		return -1
	}
}
