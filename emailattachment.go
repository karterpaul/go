//go get github.com/scorredoira/email
//https://github.com/scorredoira/email

package main 
import (
    "log"
    "net/mail"
    "net/smtp"
    "github.com/scorredoira/email"
)

func main(){
  Example()
}

func Example() {
    // compose the message
    m := email.NewMessage("Hi", "this is the body")
    m.From = mail.Address{Name: "From", Address: "abc@gmail.com"}
    m.To = []string{"cdf@gmail.com"}

    // add attachments
    if err := m.Attach("email.go"); err != nil {
        log.Fatal(err)
    }

    // send it
    auth := smtp.PlainAuth("", "abc@gmail.com", "pws", "smtp.gmail.com")
    if err := email.Send("smtp.gmail.com:587", auth, m); err != nil {
        log.Fatal(err)
    }
}







