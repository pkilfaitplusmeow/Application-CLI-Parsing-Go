package main

import (
	"fmt"
	"initial/src"
	"os"
)

func edits(fileText string) {
	var occurence int

	j := 0
	fileSize := len(fileText) - 1 // recup la taille du fichier
	str := make([]byte, fileSize) // string de recup
	for i := 0; i <= fileSize; i++ {
		if fileText[i] == '(' {
			x := 1
			j--
			occurence = src.Occ(fileText, i) // revoie le type de modif a faire
			deciVal := src.CorrectOccurence(fileText, i-1, occurence, fileSize)

			for str[j] != ' ' {
				j--

			}
			for x < len(deciVal) {
				j++
				str[j] = deciVal[x]
				x++
			}
			for fileText[i-1] != ')' {
				i++
			}
		}
		j++
		str[j] = fileText[i]
	}
	fmt.Printf("%s %d", str, j) //jdois ecrire dans le fichier a cette étape.
	os.WriteFile("output.txt", str, 0777)
}

func main() {

	fileName := os.Args[1:]               // recup le file name
	text, err := os.ReadFile(fileName[0]) // lit le text dans le fichier
	if err != nil {                       // si le fichier est inexistant
		fmt.Printf("file not found\n")
		panic(err)
	}
	edits(string(text)) //demande a la fonction d'edit (ptetre juste rediriger vers l'edit en question)
}
