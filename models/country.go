package models

type Country struct {
    Name string
    Capital string
    Population int64
    Continent string
    Description string
}

var Countries = []Country{
    {
	Name: "Japan",
	Capital: "Tokyo",
	Population: 125700000,
	Continent: "Asia",
	Description: "Uwu",
    },
    {
	Name: "Brazil",
	Capital: "Brasilia",
	Population: 2143000000,
	Continent: "South America",
	Description: "Brazil mentioned",
    },
    {
	Name: "France",
	Capital: "Paris",
	Population: 67390000,
	Continent: "Europe",
	Description: "Fois gras",
    },
}
