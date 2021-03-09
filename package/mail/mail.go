// hello62
package main

import (
	"log"

	"github.com/alexcesaro/mail/gomail"
)

func main() {
	msg := gomail.NewMessage()
	msg.SetAddressHeader("From", "accunix@accubot.com.tw", "andy")
	msg.SetHeader("To", "andy841201@gmail.com")
	msg.SetHeader("Subject", "Hello!")
	msg.SetBody("text/plain", "Hello!")
	msg.AddAlternative("text/html", "Hello <b></b>!")
	// if err := msg.Attach("p1.jpg"); err != nil {
	// 	log.Println(err)
	// 	return
	// }

	m := gomail.NewMailer("mail.accubot.com.tw", "accunix@accubot.com.tw", "96ac2731-e232-4b3d-bb9d-845c95f0acda", 587)
	if err := m.Send(msg); err != nil {
		log.Println(err)
	}
}
