package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var game_over bool = false

func test(grille [9]int) bool {
	var compteur1, compteur2 int

	// Tester les lignes
	for j := 0; j < 9; j += 3 {
		for i := 0; i < 3; i++ {
			if grille[i+j] == 1 {
				compteur1++
			}
			if grille[i+j] == 2 {
				compteur2++
			}
		}
		if compteur1 == 3 || compteur2 == 3 {
			return true
		}
		compteur1, compteur2 = 0, 0
	}

	// Tester les colonnes
	for j := 0; j < 3; j++ {
		for i := 0; i < 9; i += 3 {
			if grille[i+j] == 1 {
				compteur1++
			}
			if grille[i+j] == 2 {
				compteur2++
			}
		}
		if compteur1 == 3 || compteur2 == 3 {
			return true
		}
		compteur1, compteur2 = 0, 0
	}

	// Tester la première diagonale
	for i := 0; i < 9; i += 4 {
		if grille[i] == 1 {
			compteur1++
		}
		if grille[i] == 2 {
			compteur2++
		}
	}
	if compteur1 == 3 || compteur2 == 3 {
		return true
	}
	compteur1, compteur2 = 0, 0
	// Tester la deuxième diagonale
	for i := 2; i < 7; i += 2 {
		if grille[i] == 1 {
			compteur1++
		}
		if grille[i] == 2 {
			compteur2++
		}
	}
	if compteur1 == 3 || compteur2 == 3 {
		return true
	}

	return false
}

func main() {
	const cases int = 9

	var grille [cases]int = [cases]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	var tour int = 1

	for !game_over {
		fmt.Println(grille[0:3])
		fmt.Println("--------")
		fmt.Println(grille[3:6])
		fmt.Println("--------")
		fmt.Println(grille[6:])

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Choisissez une case entre 1 et 9")

		scanner.Scan()
		nb, _ := strconv.Atoi(scanner.Text())

		for nb > 9 || nb < 1 {
			fmt.Println("Erreur : Choisissez une case entre 1 et 9")
			scanner.Scan()
			nb, _ = strconv.Atoi(scanner.Text())
		}

		var reste int = tour % 2
		if grille[nb-1] == 0 {
			if reste == 0 {
				grille[nb-1] = 1
			} else {
				grille[nb-1] = 2
			}
		} else {
			fmt.Println("La case est déjà occupée, veuillez en choisir une autre.")
			continue
		}

		tour++
		var resultat bool = test(grille)
		if resultat {
			game_over = true
			fmt.Println("Le joueur", reste+1, "a gagné la partie")
		}
		var nulle int = 0
		for i := 0; i < len(grille); i++ {
			if grille[i] != 0 {
				nulle++
			}
		}
		if nulle == len(grille) {
			game_over = true
			fmt.Println("Match nul")
		}
	}
}
