package src

import (
	"fmt"
)

func CorrectOccurence(fileText string, i int, occurence int, fileSize int) {
	i--
	pow := 0
	var hexa int = 0

	//dans le cas 1 (hex) retourne en arriere jusqu'au dernier espace et convertie en decimale
	if occurence == 1 {
		for fileText[i] != ' ' {
			hexa += HexaCoverter(fileText, i, pow)
			pow++ //puissance du 16
			i--
		}
		fmt.Printf("hexa value = %d\n", hexa)
		fmt.Printf("%s", string(hexa))
	}
	//dans le cas 2 (dec) ...
}
