package models

// Custom errors that will be wrapped on all errors
type ResponseError struct {
	Message string `json:"message"`
	Status  int    `json:"-"`
}
