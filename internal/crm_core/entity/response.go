package entity

type CustomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type CustomResponseWithData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
