package internal

type MyError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
