package src

import (
	"fmt"
	"os"
)

func Edits(fileText string) {
	var occurence int
	var repeat int
	var corrected []byte
	j := -1

	fileSize := len(fileText) - 1   // recup la taille du fichier
	str := make([]byte, fileSize+1) // chaine de char de recup
	for i := 0; i <= fileSize; i++ {
		if fileText[i] == '(' {
			x := 0
			j--
			occurence, repeat = Occ(fileText, i)       // revoie le type de modif a faire et le nombre
			fmt.Printf("(%d %d)\n", occurence, repeat) // debug
			// une boucle qui commmence au mot d'avant "repeat" fois ?(todo)

			corrected = CorrectOccurence(fileText, i-1, occurence, fileSize) // renvoie un bytes corrigé

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
	str = Ponctuation(str, len(str))

	os.WriteFile("output.txt", str, 0644) //ecrit les characters dans le fichier output.txt avec les perms d'ecrire et lire
	if str[j] == ')' {                    // si le fichier finis avec une modification
		os.Truncate("output.txt", int64(j))
	} else {
		os.Truncate("output.txt", int64(j+1)) //evite tout les espaces a la fin du fichier en limitant la taille
	}
	fmt.Printf("\n┌———————————————————————————┐\n|file successfully modified!|\n└———————————————————————————┘\n")
}
