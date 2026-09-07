package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	nombreSecret := 7

	// Permet à Go de servir index.html et styles.css
	http.Handle("/", http.FileServer(http.Dir(".")))

	// Gestion du jeu
	http.HandleFunc("/deviner", func(w http.ResponseWriter, r *http.Request) {

		choix, _ := strconv.Atoi(r.FormValue("choix"))

		if choix < 1 || choix > 10 {
			fmt.Fprint(w, "Le nombre doit être entre 1 et 10.")
			return
		}

		if choix == nombreSecret {
			fmt.Fprint(w, "Bravo ! Tu as deviné le nombre secret.")
		} else if choix < nombreSecret {
			fmt.Fprint(w, "Le nombre secret est plus grand.")
		} else {
			fmt.Fprint(w, "Le nombre secret est plus petit.")
		}
	})

	fmt.Println("Serveur lancé sur http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}