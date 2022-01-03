package main

import (
	"fmt"
	"net/url"

	"github.com/walkergriggs/okapi"
)

const BASE_URL = "https://swapi.dev"

func main() {
	// Create a new client
	client, _ := okapi.NewClient(&okapi.Config{
		Address: BASE_URL,
	})

	// Use the wrapped to get a single character
	res, err := People(client).Get("R2-D2")
	if err != nil {
		panic(err)
	}

	r2 := res.People[0]
	fmt.Printf("R2-D2 was made in %s and was featured in %d films:\n", r2.BirthYear, len(r2.Films))

	// Use client directly to fetch individual films
	for i := range r2.Films {
		url, err := url.Parse(r2.Films[i])
		if err != nil {
			panic(err)
		}

		var res FilmsResult
		if _, err := client.Get(url.Path, &okapi.QueryOptions{Out: &res}); err != nil {
			panic(err)
		}

		fmt.Printf(" ... Episode %d: %s\n", res.Episode, res.Title)
	}
}
