package src

import "fmt"

func CorrectOccurence(fileText string, i int, occurence int, fileSize int) {
	i--
	pow := 0
	var hexa int = 0
	var a byte
	if occurence == 1 {
		for fileText[i] != ' ' { //arrive au dernier charactere du mot avant les parenthèses
			if fileText[i] >= 'A' && fileText[i] <= 'F' {
				a = fileText[i] - 55          // valeur de la lettre en hexa ('A' = 65 et A(hex) = 10)
				hexa += int(a) * Pow(16, pow) // hexa = a * (16^dig)

				fmt.Printf("char %c\n", a+55)
				fmt.Printf("pow = %d\n", pow)
				fmt.Printf("a value = %d\n", a)
				fmt.Printf("hexa value = %d\n", hexa)
				fmt.Printf("i = %d\n_____\n", i)

			}
			pow++ // pour la puissance du 16
			i--
		}
	}
}
