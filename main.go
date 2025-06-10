package main

import "fmt"

type Album struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

// Variable global (declaración válida fuera de funciones)
var albums = []Album{
	{ID: 1, Title: "El Mal Querer", Artist: "Rosalía", Year: 2018},
	{ID: 2, Title: "Colores", Artist: "J Balvin", Year: 2020},
}

func main() {
	fmt.Println(albums)
}
