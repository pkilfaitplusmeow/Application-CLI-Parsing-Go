package src

func DecimalToString(deci int) []byte {
	i := 0
	var a int
	temp := deci
	for temp != 0 {
		i++
		temp /= 10
	}
	str := make([]byte, i+1)

	for i != 0 {
		a = deci % 10
		deci /= 10
		str[i] = byte(a + '0')
		i--
	}
	return str
}
