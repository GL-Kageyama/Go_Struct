package main

import (
	"fmt"
)

// Common struct for supervisor names
type Directer struct {
	Name    string
	Country string
}

// Struct for movie titles
type Movie struct {
	Directer
	Title    string
	Year     int
}

// Struct for award titles
type Prise struct {
	Directer
	Title    string
	Year     int
}

// main function
func main() {

	// Steven Spielberg's work
	movie1 := Movie{
		Directer: Directer{"Steven Spielberg", "USA"},
		Title: "Saving Private Ryan",
		Year: 1998,
	}
	prise1 := Prise{
		Directer: Directer{"Steven Spielberg", "USA"},
		Title: "Academy Award for Directing",
		Year: 1998,
	}

	// Akira Kurosawa's work
	movie2 := Movie{
		Directer: Directer{"Akira Kurosawa", "JAPAN"},
		Title: "Seven Samurai",
		Year: 1954,
	}
	prise2 := Prise{
		Directer: Directer{"Akira Kurosawa", "JAPAN"},
		Title: "Leone d'Argento",
		Year: 1954,
	}

	// Roman Polanski's work
	movie3 := Movie{
		Directer: Directer{"Roman Polanski", "Poland"},
		Title: "The Pianist",
		Year: 2002,
	}
	prise3 := Prise{
		Directer: Directer{"Roman Polanski", "Poland"},
		Title: "Golden Palmae",
		Year: 2002,
	}

	// Struct Output
	fmt.Println(movie1, prise1)
	fmt.Println(movie2, prise2)
	fmt.Println(movie3, prise3)
}

// =================================
//           Output Sample
// =================================
// ~ $ go build main.go 
// ~ $ ./main 
// {{Steven Spielberg USA} Saving Private Ryan 1998} {{Steven Spielberg USA} Academy Award for Directing 1998}
// {{Akira Kurosawa JAPAN} Seven Samurai 1954} {{Akira Kurosawa JAPAN} Leone d'Argento 1954}
// {{Roman Polanski Poland} The Pianist 2002} {{Roman Polanski Poland} Golden Palmae 2002}
