package models

type ResponseModel struct {
	RetCode string `json:"retcode"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponsetokenModel struct {
	RetCode string `json:"retcode"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Token   any    `json:"token"`
}

type ResponseWoModel struct {
	RetCode string `json:"retcode"`
	Message string `json:"message"`
}
