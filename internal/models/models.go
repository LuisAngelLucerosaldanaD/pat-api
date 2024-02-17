package models

type Response struct {
	Error bool   `json:"error"`
	Code  int    `json:"code"`
	Type  string `json:"type"`
	Msg   string `json:"msg"`
	Data  string `json:"data"`
}
