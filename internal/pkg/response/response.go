package response

type Success struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Status  Status `json:"status"`
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Message struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Message string `json:"message"`
	Errors  any    `json:"errors"`
	Status  Status `json:"status"`
}
