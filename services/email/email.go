package email

import (
	"bytes"
	"log"
	"text/template"

	"gopkg.in/gomail.v2"
)

//WelcomeEmail is
func WelcomeEmail(email string) error {
	d := gomail.NewPlainDialer("smtp.yandex.ru", 465, "akezhanyesbolatov", "Yfehsp2203")

	m := gomail.NewMessage()
	m.SetHeader("From", "akezhanyesbolatov@yandex.ru")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Регистрация на сервисе cactus")

	tmpl, err := template.ParseFiles("./static/email.html")
	if err != nil {
		log.Println(err)
		return err
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, nil)
	if err != nil {
		return nil
	}

	m.SetBody("text/html", b.String())

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
