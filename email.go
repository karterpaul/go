package main

import (
	"log"
	"net/smtp"
        "fmt"
)

func main() {
       for i := 0; i < 3; i++ {
	go send("hello there")
        log.Printf("mail count:%d",i); 
        }
  var input string
  fmt.Scanln(&input)


}

func send(body string) {
	from := "@gmail.com"
	pass := ""
	to := "anymail@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	
	log.Print("sent, test")
}
