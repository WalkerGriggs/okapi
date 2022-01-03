package main

import "github.com/walkergriggs/okapi"

type PeopleClient struct {
	client *okapi.Client
}

func People(client *okapi.Client) *PeopleClient {
	return &PeopleClient{
		client: client,
	}
}

func (c *PeopleClient) Get(name string) (*PeopleResponse, error) {
	var res PeopleResponse

	opts := &okapi.QueryOptions{
		Out: &res,
		Params: map[string]string{
			"search": name,
		},
	}

	if _, err := c.client.Get("/api/people/", opts); err != nil {
		return nil, err
	}
	return &res, nil
}

// vvv Full type definition (see: https://swapi.dev/documentation#people) vvv

type (
	PeopleResult struct {
		// The name of this person.
		Name string `json:"name"`

		// The birth year of the person, using the in-universe standard of BBY or
		// ABY - Before the Battle of Yavin or After the Battle of Yavin. The Battle
		// of Yavin is a battle that occurs at the end of Star Wars episode IV: A
		// New Hope.
		BirthYear string `json:"birth_year"`

		// The eye color of this person. Will be "unknown" if not known or "n/a" if
		// the person does not have an eye.
		EyeColor string `json:"eye_color"`

		// The gender of this person. Either "Male", "Female" or "unknown", "n/a" if
		// the person does not have a gender.
		Gender string `json:"gender"`

		// The hair color of this person. Will be "unknown" if not known or "n/a" if
		// the person does not have hair.
		HairColor string `json:"gender"`

		// The height of the person in centimeters.
		Height string `json:"height"`

		// The mass of the person in kilograms.
		Mass string `json:"mass"`

		// The skin color of this person.
		SkinColor string `json:"skin_color"`

		// The URL of a planet resource, a planet that this person was born on or
		// inhabits.
		Homeworld string `json:"homeworld"`

		// An array of film resource URLs that this person has been in.
		Films []string `json:"films"`

		// An array of species resource URLs that this person belongs to.
		Species []string `json:"species"`

		// An array of starship resource URLs that this person has piloted.
		Starships []string `json:"starships"`

		// An array of vehicle resource URLs that this person has piloted.
		Vehicles []string `json:"vehicles"`

		// The hypermedia URL of this resource.
		URL string `json:"url"`

		// The ISO 8601 date format of the time that this resource was created.
		Created string `json:"created"`

		// The ISO 8601 date format of the time that this resource was edited.
		Edited string `json:"edited"`
	}

	PeopleResponse struct {
		// The number of people returned across all pages.
		Count int `json:"count"`

		// The URL to the next result page.
		Next string `json:"next"`

		// The URL to the previous result page.
		Previous string `json:"previous"`

		// The PeopleResults for the current page.
		People []PeopleResult `json:"results"`
	}
)
