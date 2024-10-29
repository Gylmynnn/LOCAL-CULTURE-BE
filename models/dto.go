package models

type Responsetype struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}



