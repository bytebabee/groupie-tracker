package models

type Artists struct {
	ID              int                 `json:"id"`
	Image           string              `json:"image"`
	Name            string              `json:"name"`
	Members         []string            `json:"members"`
	CreationDate    int                 `json:"creationDate"`
	FirstAlbum      string              `json:"firstAlbum"`
	LocationsUrl    string              `json:"locations"`
	ConcertDatesUrl string              `json:"concertDates"`
	RelationsUrl    string              `json:"relations"`
	Locations       []string            `json:"-"`
	ConcertDates    []string            `json:"-"`
	Relations       map[string][]string `json:"-"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Dates struct {
	ID           int      `json:"id"`
	ConcertDates []string `json:"dates"`
}

type Relations struct {
	ID        int                 `json:"id"`
	Relations map[string][]string `json:"datesLocations"`
}
