package main

import (
	"fmt"
	"initial/src"
	"os"
)

func edits(fileText string) {
	var occurence int

	j := -1
	fileSize := len(fileText) - 1 // recup la taille du fichier
	str := make([]byte, fileSize) // string de recup
	for i := 0; i <= fileSize; i++ {
		if fileText[i] == '(' {
			x := 1
			j--
			occurence = src.Occ(fileText, i)                                      // revoie le type de modif a faire
			corrected := src.CorrectOccurence(fileText, i-1, occurence, fileSize) // renvoie un bytes corrigé
			//si erreur a gerer (todo)
			for str[j] != ' ' {
				j--
			}
			for x < len(corrected) {
				j++
				str[j] = corrected[x]
				x++
			}
			for (fileText[i-1] != ')') && (i != fileSize) {
				i++
			}
		}
		j++
		str[j] = fileText[i]
	}
	os.WriteFile("output.txt", str, 0644) //ecrit les characters dans le fichier output.txt avec les perms d'ecrire et lire
	if str[j] == ')' {                    // si le fichier finis avec une modification
		os.Truncate("output.txt", int64(j))
	} else {
		os.Truncate("output.txt", int64(j+1)) //evite tout les espaces a la fin du fichier en limitant la taille
	}
	fmt.Printf("\n┌———————————————————————————┐\n|file successfully modified!|\n└———————————————————————————┘\n")
}

func main() {

	fileName := os.Args[1:] // recup le file name

	//error si fichier vide (todo)

	text, err := os.ReadFile(fileName[0]) // lit le text dans le fichier

	for err != nil { // si le fichier est inexistant
		fmt.Printf("file not found\n")
		//fmt.Printf("%s\n", err)
		// fmt.Scanf("%s", fileName)

	}
	edits(string(text)) //demande a la fonction d'edit (ptetre juste rediriger vers l'edit en question)
}
