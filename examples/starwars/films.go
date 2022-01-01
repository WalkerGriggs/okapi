package main

import "github.com/walkergriggs/okapi"

type FilmsClient struct {
	client *okapi.Client
}

func Films(client *okapi.Client) *FilmsClient {
	return &FilmsClient{
		client: client,
	}
}

func (c *FilmsClient) Get(name string) (*FilmsResponse, error) {
	var res FilmsResponse

	opts := &okapi.QueryOptions{
		Out: &res,
		Params: map[string]string{
			"search": name,
		},
	}

	if err := c.client.Get("/api/films/", opts); err != nil {
		return nil, err
	}
	return &res, nil
}

// vvv Full type definition (see: https://swapi.dev/documentation#films) vvv

type (
	FilmsResult struct {
		// The title of this film
		Title string `json:"title"`

		// The episode number of this film.
		Episode int `json:"episode_id"`

		// The opening paragraphs at the beginning of this film.
		OpeningCrawl string `json:"opening_crawl"`

		// The name of the director of this film.
		Director string `json:"director"`

		// The name(s) of the producer(s) of this film. Comma separated.
		Producer string `json:"producer"`

		// The ISO 8601 date format of film release at original creator country.
		ReleaseDate string `json:"release_date"`

		// An array of species resource URLs that are in this film.
		Species []string `json:"species"`

		// An array of starship resource URLs that are in this film.
		Starships []string `json:"starships"`

		// An array of vehicle resource URLs that are in this film.
		Vehicles []string `json:"vehicles"`

		// An array of people resource URLs that are in this film.
		Characters []string `json:"characters"`

		// An array of planet resource URLs that are in this film.
		Planets []string `json:"planets"`

		// The hypermedia URL of this resource.
		URL string `json:"url"`

		// The ISO 8601 date format of the time that this resource was created.
		Created string `json:"created"`

		// The ISO 8601 date format of the time that this resource was edited.
		Edited string `json:"edited"`
	}

	FilmsResponse struct {
		// The number of films returned across all pages.
		Count int `json:"count"`

		// The URL to the next result page.
		Next string `json:"next"`

		// The URL to the previous result page.
		Previous string `json:"previous"`

		// The PeopleResults for the current page.
		Films []FilmsResult `json:"results"`
	}
)
