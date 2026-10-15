package src

func HexaCoverter(fileText string, i int, pow int) int {
	var hexa int = 0
	if fileText[i] >= 'A' && fileText[i] <= 'F' {
		a := fileText[i] - 55         // valeur de la lettre en hexa ('A' = 65 et A(hex) = 10)
		hexa += int(a) * Pow(16, pow) // hexa = a * (16^dig)
	}

	if fileText[i] >= '0' && fileText[i] <= '9' {
		a := fileText[i] - 48         // valeur du int en hexa ('0' = 48 et 0(hex) = 0)
		hexa += int(a) * Pow(16, pow) // hexa = a * (16^dig)
	}
	return hexa
}
