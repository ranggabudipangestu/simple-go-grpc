package models

type ResponseWrapper struct {
	Success    bool
	StatusCode int
	Message    string
	Data       interface{}
}
