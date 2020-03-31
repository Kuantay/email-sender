package model

//EmailSender is
type EmailSender struct {
	Email struct {
		Address string `json:"address"`
	} `json:"email"`
}
