package models

type Runner struct {
	ID           string    `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Age          int       `json:"age,omitempty"`
	IsActive     bool      `json:"is_active"`
	Country      string    `json:"country"`
	PersonalBest string    `json:"personal_best,omitempty"`
	SeasonBest   string    `json:"season_best,omitempty"`
	Results      []*Result `json:"results,omitempty"`
}

type Result struct {
	ID         string `json:"id"`
	RunnerID   string `json:"runner_id"`
	RaceResult string `json:"race_result"`
	Location   string `json:"location"`
	Position   int    `json:"position"`
	Year       int    `json:"year"`
}

// Custom errors that will be wrapped on all errors
type ResponseError struct {
	Message string `json:"message"`
	Status  int    `json:"-"`
}
