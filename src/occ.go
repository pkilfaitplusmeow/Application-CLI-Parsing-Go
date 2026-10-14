package src

func Occ(str string, i int) int {
	i += 2
	switch str[i] {
	case 'h': //hexa
		return 1
	case 'b': //binaire
		return 2
	case 'u': //uppercase
		return 3
	case 'l': //lowercase
		return 4
	case 'c': //cap
		return 5
	default:
		return -1
	}
}
