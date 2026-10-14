package main

import (
	"fmt"
	"os"
)

func occ(str string, i int) int {
	i++
	switch str[i] {
	case 'h':
		return 1
	case 'b':
		return 2
	case 'u':
		return 3
	case 'l':
		return 4
	case 'c':
		return 5
	default:
		return -1
	}
}

func edits(fileText string) {
	var occurence int

	j := 0
	fileSize := len(fileText) - 1 // recup la taille du fichier
	str := make([]byte, fileSize) // string de recup
	for i := 0; i <= fileSize; i++ {
		if fileText[i] == '(' {
			occurence = occ(fileText, i)

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
