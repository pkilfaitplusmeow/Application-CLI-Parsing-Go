package main

import (
	"fmt"
	"os"
)

func edits(fileText string) {
	fileSize := len(fileText) - 1 // recup la taille du fichier
	for i := 0; i <= fileSize; i++ {
		fmt.Printf("%c", fileText[i])
		//fmt.Printf("%d", i)
	}

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
