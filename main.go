package main

import (
	"net/http"

	"github.com/AkezhanOb1/email-sender/services/business/owner"
)

func main() {
	go func() {
		owner.SendEmailToBusinessOwner()
	}()

	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":3030", nil); err != nil {
		panic(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	return
}
