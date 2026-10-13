package main

import (
	"fmt"
	"os"
)

func correct(fileText string) {

}

func main() {

	file := os.Args[1:]               // recup le file name
	text, err := os.ReadFile(file[0]) // lit le text dans le fichier
	if err != nil {                   // si le fichier est inexistant
		fmt.Printf("file not found")
	}
	fmt.Printf("%s\n", text)
}
