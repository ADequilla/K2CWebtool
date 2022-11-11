package models

type ResponseModel struct {
	RetCode string `json:"retcode"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWoModel struct {
	RetCode string `json:"retcode"`
	Message string `json:"message"`
}
