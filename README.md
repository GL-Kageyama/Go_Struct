# Go Struct

## Struct


## Code
```Go
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

	// Struct Output
	fmt.Println(movie1, prise1)
}
```

## Output Sample
$ go build main.go  
$ ./main  
{{Steven Spielberg USA} Saving Private Ryan 1998} {{Steven Spielberg USA} Academy Award for Directing 1998}  
