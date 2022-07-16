package report

import (
	"crypto/tls"
	gomail "gopkg.in/mail.v2"
)

type Email struct {
	from string
	to   string

	smtpHost string
	smtpPort int
	pass     string

	html bool

	rep Reporter
}

func NewEmail(from string, to string, smtpHost string, smtpPort int, pass string, html bool) Reporter {

	email := &Email{from: from, to: to, smtpHost: smtpHost, smtpPort: smtpPort, pass: pass, html: html}
	email.rep = NewBaseReport("", email)

	return email.rep
}

func (e *Email) Write(message []byte) (n int, err error) {

	m := gomail.NewMessage()

	//Set E-Mail sender
	m.SetHeader("From", e.from)

	// Set E-Mail receivers
	m.SetHeader("To", e.to)

	// Set E-Mail subject
	m.SetHeader("Subject", stripTags(e.rep.Title()))

	// Set E-Mail body. You can set plain text or html with text/html
	var contentType = "text/plain"
	if e.html {
		contentType = "text/html"
	}
	m.SetBody(contentType, string(message))

	// Settings for SMTP server
	d := gomail.NewDialer(e.smtpHost, e.smtpPort, e.from, e.pass)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		return 0, err
	}

	return len(message), nil

}
