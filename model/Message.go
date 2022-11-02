package model

type Message struct {
	Id  int    `json:"id""`
	Cmd int    `json:"cmd""`
	Msg string `json:"msg""`
}
