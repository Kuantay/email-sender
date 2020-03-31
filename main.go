package main

import (
	"log"
	"net/http"
	"github.com/AkezhanOb1/email-sender/services/business/owner"
)

func main() {
	owner.SendEmailToBusinessOwner()
	http.ListenAndServe(":8080", nil)
}