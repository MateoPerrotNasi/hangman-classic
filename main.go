package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
	"fmt"
	"os"
)

type Hangman struct {
	word string
	hiddenWord []string
	attempt int
	position string
	tried []string
}

func main() {
	var h1 Hangman
	h1.attempt = 10
	Rules()
	h1.ReadFile()
	h1.CreateHidden()
	h1.Reveal()
	for i:= 0; i < 9999; i++ {
		h1.HangmanPositions()
		h1.IsWin()
		h1.IsLoose()
		h1.PlayerTurn()
	}
}

func Rules() {
	fmt.Println("Bienvenue sur le meilleur jeu du pendu du monde.")
	fmt.Println("Voici les règles:\n")
	fmt.Println("Dans ce jeu, le but est de deviner un mot choisis par l'ordinateur.")
	fmt.Println("Pour cela tu écris une lettre et l'ordinateur te dis si elle est présente dans le mot.")
	fmt.Println("Tu peux même écrire un mot en entier si tu penses avoir trouvé.")
	fmt.Println("Mais fait attention car chaque erreur rapprochera José le bonhomme un peu plus de son dernier souffle.")
	fmt.Println("Bonne chance, et puisse la chance te sourire (mais surtout à José).\n")

}

func (h *Hangman) ReadFile() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(38)
	data, err := ioutil.ReadFile("words.txt")
	content := string(data)
	words := strings.Split(content, "\n")
	if err != nil {
		fmt.Println(err)

	} else {
		for i := 0; i < n; i++ {
			h.word = words[i]
		}
	}
	// fmt.Println(h.word) // c'est le cheat code pour voir le mot
}

func (h *Hangman) CreateHidden() {
	for i := 0; i < len(h.word); i++ {
		h.hiddenWord = append(h.hiddenWord, "_")
	}
}

func (h *Hangman) Reveal() {
	var randomLetter int
	n := len(h.word) / 2 - 1
	if n > 0 {
		for i := 1 ; i <= n ; i ++ {
			randomLetter = rand.Intn(len(h.word))
			for i, letter := range h.word {
				if i == randomLetter && h.hiddenWord[i] == "_" {
					h.hiddenWord[i] = string(letter)
				}
			}
		}
	}
}

func (h *Hangman) HangmanPositions() {
	data, err := ioutil.ReadFile("hangman.txt")
	content := string(data)
	positions := strings.Split(content, "\n\n")
	var n int = 10-h.attempt

	if err != nil {
		fmt.Println(err)

	} else {
		for i := 0 ; i < n ; i++ {
			h.position = positions[i]
		}
		if n > 0 {
			fmt.Println(positions[n-1])
		}
	}

}

func (h *Hangman) PlayerTurn() {
	var s string
	IsTried := false
	found := false
	fmt.Println("Voici le mot à deviner :",h.hiddenWord)
	if h.attempt > 1 {
		fmt.Println("Il vous reste", h.attempt ,"essais.")
	} else {
		fmt.Println("Il vous reste", h.attempt ,"essai.")
	}
	if h.tried != nil {
		fmt.Println("Vous avez déjà essayé:",h.tried)
	}
	fmt.Println("Entrez un mot ou une lettre.")
	fmt.Scanln(&s)
	strings.ToLower(s)
	if len(s) > 1 {
		if s == h.word {
			for i, letter := range h.word {
				h.hiddenWord[i] = string(letter)
			}
			h.IsWin()
		} else {
			fmt.Println("Ce n'était pas le bon mot.")
			h.tried = append(h.tried, s)
			h.attempt -= 1
		}
	} else {
		for _, word := range h.tried {
			if s == word {
				fmt.Println("Vous avez déjà essayé cette lettre, tentez autre chose:")
				IsTried = true
				fmt.Scanln(&s)
				h.tried = append(h.tried, s)
				h.attempt -= 1
			}
		}
		if !IsTried {
			for i, letter := range h.word {
				if s == string(letter) {
					h.hiddenWord[i] = string(letter) 
					found = true
				}
			}
			if !found {
				fmt.Println("Votre lettre n'est pas présente dans le mot.")
				h.tried = append(h.tried, s)
				h.attempt -= 1
			} else {
				h.tried = append(h.tried, s)
				fmt.Println("Bravo, vous avez deviné une lettre !")
			}
		}
	}
}

func (h *Hangman) IsTried(s string) {
	for _, word := range h.tried {
		if s == word {
			fmt.Println("Vous avez déjà essayé ", word)
			fmt.Println("Entrez un mot ou une lettre.")
			fmt.Scanln(&s)
			break
		}
	}
}

func (h *Hangman) IsWin() {
	match := false
	Loop:
	for i, letter := range h.word {
		if h.hiddenWord[i] == string(letter) {
			match = true
		} else {
			match = false
			break Loop
		}
	}
	if match {
		fmt.Println("Bravo, tu as gagnés ! Le mot était bel est bien", h.word)
		os.Exit(0)
	}
}

func (h *Hangman) IsLoose() {
	if h.attempt == 0 {
		fmt.Println("Tu as perdu, tes chances sont épuisés et José est mort.")
		fmt.Println("Le mot a deviner était:", h.word)
		os.Exit(0)	
	}
}
