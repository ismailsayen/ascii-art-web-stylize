package asciiartweb

import (
	"bufio"
	"log"
	"os"
)

func AsciiArt(sentence, banner string) string {
	// Ouvre le fichier contenant les définitions des symboles.
	file, err := os.Open("./Files/" + banner + ".txt")
	if err != nil {
		return "Banner not found Only: standard | shadow | thinkertoy"
	}
	defer file.Close() // Assure la fermeture du fichier après usage

	// Vérifie que le mot n'est pas vide
	if len(sentence) == 0 {
		return "Enter a sentence"
	}

	scanner := bufio.NewScanner(file) // Initialise un scanner pour lire le fichier ligne par ligne

	count := 0
	symbole := []string{}
	symboles := [][]string{}

	// Boucle à travers chaque ligne du fichier texte
	for scanner.Scan() {
		symbole = append(symbole, scanner.Text())
		count++

		// Chaque symbole ASCII dans ce format est constitué de 9 lignes
		if count == 9 {
			symboles = append(symboles, symbole) // Ajoute le symbole complet à la liste de symboles
			symbole = []string{}
			count = 0
		}
	}
	// Vérifie si le fichier contient 94 symboles.
	if len(symboles) < 95 {
		log.Fatal("Please make sure all characters are present in the file.")
	}
	// retourner le resulta
	return PrintWords(Split(sentence), symboles)
}
