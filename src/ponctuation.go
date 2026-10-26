package src

func Ponctuation(str []byte, fileSize int) []byte {
	i := 0
	for i < len(str) {
		switch str[i] {
		case ',', '?', ';', '.', ':', '!':
			if str[i-1] == ' ' {
				str[i-1] = str[i]
				str[i] = ' '

			}
		}
		i++
	}
	return str
}

func IsPonctuation(str []byte, i int) bool {
	switch str[i] {
	case ',', '?', ';', '.', ':', '!':
		return true
	default:
		return false
	}
}
