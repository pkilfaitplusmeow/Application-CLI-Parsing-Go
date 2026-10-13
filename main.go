package main

import (
	"fmt"
	"os"
)

func edits(fileText string) {
	lastchar := fileText[len(fileText)-1]
	fmt.Printf("|%d|", lastchar)
	fmt.Printf("|%c|", fileText[lastchar])

}

func main() {

	file := os.Args[1:]               // recup le file name
	text, err := os.ReadFile(file[0]) // lit le text dans le fichier
	if err != nil {                   // si le fichier est inexistant
		fmt.Printf("file not found")
	}
	edits(string(text))
	fmt.Printf("\n%s\n", text)
}
