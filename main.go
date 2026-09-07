package main

import "fmt"

func main() {
    question := "Devine un nombre entre 1 et 10"
    fmt.Println(question)
    nombreSecret := 7
    fmt.Println("Entre ton nombre :")
    var choix int
    fmt.Scan(&choix)
    for choix < 1 || choix > 10 {
        fmt.Println("Le nombre doit être entre 1 et 10. Entre un nouveau nombre :")
        fmt.Scan(&choix)
    }
    if choix == nombreSecret {
        fmt.Println("Bravo ! Tu as deviné le nombre secret.")
    }else if choix < nombreSecret {
        fmt.Println("Le nombre secret est plus grand que ton choix.")
    } else {
        fmt.Println("Le nombre secret est plus petit que ton choix.")
    }
}