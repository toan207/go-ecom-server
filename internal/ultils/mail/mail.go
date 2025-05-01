package mail

import (
	"fmt"
	"net/smtp"

	"go.uber.org/zap"
	"pawtopia.com/global"
)

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
	msg := "MINE-version: 1.0;\nContent-Type: text/html; charset=UTF-8;\r\n"
	msg += fmt.Sprintf("From: %s <%s>\r\n", mail.From.Name, mail.From.Email)
	msg += fmt.Sprintf("To: %s\r\n", mail.To[0].Email)
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func SendTextEmail(to []string, otp string) error {
	mail := Mail{
		From: EmailAdress{
			Email: global.Config.STMP.From,
			Name:  "Pawtopia",
		},
		To: []EmailAdress{
			{
				Email: to[0],
			},
		},
		Subject: "OTP Verification",
		Body:    fmt.Sprintf("Your OTP is %s. Please enter it to verify account!", otp),
	}

	msg := BuilMessage(mail)

	auth := smtp.PlainAuth("", global.Config.STMP.Username, global.Config.STMP.Password, global.Config.STMP.Host)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", global.Config.STMP.Host, global.Config.STMP.Port), auth, global.Config.STMP.From, to, []byte(msg))
	if err != nil {
		global.Logger.Error("Failed to send email", zap.Error(err))
		return err
	}

	global.Logger.Info("Email sent successfully", zap.String("to", to[0]), zap.String("from", global.Config.STMP.From))
	return nil
}
