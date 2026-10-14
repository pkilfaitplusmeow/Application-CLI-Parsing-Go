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
		if fileText[i+1] == '(' {
			occurence = src.Occ(fileText, i)
			src.CorrectOccurence(fileText, i, occurence, fileSize)
			for fileText[i-1] != ')' {
				i++
			}
		}
		j++
		str[j] = fileText[i]
	}
	fmt.Printf("\n----\n%s\n----\n", fileText)
	fmt.Printf("\n----\n%s\n----\n", str)
	fmt.Printf("%d", occurence)
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
