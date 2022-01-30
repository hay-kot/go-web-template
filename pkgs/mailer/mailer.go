package mailer

import (
	"encoding/base64"
	"fmt"
	"mime"
	"net/smtp"
	"strconv"
)

type Mailer struct {
	Host     string `conf:""`
	Port     int    `conf:""`
	Username string `conf:""`
	Password string `conf:""`
	From     string `conf:""`
}

func (m *Mailer) server() string {
	return m.Host + ":" + strconv.Itoa(m.Port)
}

func (m *Mailer) Send(msg *Message) error {
	server := m.server()

	header := make(map[string]string)
	header["From"] = msg.From.String()
	header["To"] = msg.To.String()
	header["Subject"] = mime.QEncoding.Encode("UTF-8", msg.Subject)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(msg.Body))

	return smtp.SendMail(
		server,
		smtp.PlainAuth("", m.Username, m.Password, m.Host),
		m.From,
		[]string{msg.To.Address},
		[]byte(message),
	)
}
