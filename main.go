package main

import (
	"fmt"
	"initial/src"
	"os"
)

func main() {

	fileName := os.Args[1:] // recup le file name

	//error si fichier vide (todo)

	text, err := os.ReadFile(fileName[0]) // lit le text dans le fichier

	for err != nil { // si le fichier est inexistant
		fmt.Printf("file not found\n")
		//fmt.Printf("%s\n", err)
		// fmt.Scanf("%s", fileName)

	}
	src.Edits(string(text)) //demande a la fonction d'edit (ptetre juste rediriger vers l'edit en question)
}
