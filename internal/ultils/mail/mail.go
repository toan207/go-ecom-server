package mail

import "fmt"

type EmailAdress struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Mail struct {
	From    EmailAdress
	To      []EmailAdress
	Subject string
	Body    string
}

func BuilMessage(mail Mail) string {
	msg := "MINE-version: 1.0;\nContent-Type: text/html; charset=UTF-8;\n;\r\n"
	msg += fmt.Sprintf("From: %s <%s>\n", mail.From.Name, mail.From.Email)
	msg += fmt.Sprintf("To: %s <%s>\n", mail.To[0].Name, mail.To[0].Email)
	msg += fmt.Sprintf("Subject: %s\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func SendTextEmail(to []string, from string, otp string) error {
	mail := Mail{
		From: EmailAdress{
			Email: from,
			Name:  "pawtopia",
		},
		To: []EmailAdress{
			{
				Email: to[0],
				Name:  "pawtopia",
			},
		},
		Subject: "OTP Verification",
		Body:    fmt.Sprintf("Your OTP is %s. Please enter it to verify account!", otp),
	}

	msg := BuilMessage(mail)
	return fmt.Errorf("Send email to %s with message: %s", to[0], msg)
}
