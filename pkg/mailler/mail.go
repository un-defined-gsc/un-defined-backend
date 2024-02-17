package email

import (
	"bytes"
	"context"
	"crypto/tls"
	"html/template"
	"log"

	"strconv"

	"gopkg.in/gomail.v2"
)

type Email struct {
	Address  string // gonderen mail
	Name     string // email gönderen ad
	Host     string // email servisi host
	Port     int    // email servisi port
	Username string // email servisi username
	Password string // email servisi parola
	errchan  chan error
}

// Address   gonderen mail
// Name      email gönderen ad
// Host      email servisi host
// Port      email servisi port
// Username  email servisi username
// Password  email servisi parola
func EmailInit(address, name, host, port, username, password string) *Email {

	portInt, _ := strconv.Atoi(port)
	errchan := make(chan error)
	return &Email{
		Address:  address,
		Name:     name,
		Host:     host,
		Port:     portInt,
		Username: username,
		Password: password,
		errchan:  errchan,
	}
}

func (e *Email) GetErrChan() chan error {
	return e.errchan
}

func (e *Email) WriteStdoutError() {
	for err := range e.errchan {
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func (e *Email) Send(subject string, HTMLbody string, mail string) {
	go func(errchan chan error) {
		m := gomail.NewMessage()
		m.SetHeader("From", m.FormatAddress(e.Address, e.Name))
		m.SetHeader("To", mail)
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", HTMLbody)

		d := gomail.NewDialer(e.Host, e.Port, e.Username, e.Password)

		if err := d.DialAndSend(m); err != nil {
			errchan <- err
			return
		}
	}(e.errchan)
}

func (e *Email) SendWithTemplate(subject string, p interface{}, templatePath string, mail string) error {

	body, err := parseMailTemplate(p, templatePath)
	if err != nil {
		return err
	}
	e.Send(subject, body, mail)
	return nil
}
func (e *Email) SendBulkEmail(to, subject, HTMLbody string, mails []string) {
	go func(errchan chan error) {
		m := gomail.NewMessage()

		m.SetHeaders(map[string][]string{
			"From":    {m.FormatAddress(e.Address, e.Name)},
			"Bcc":     mails,
			"Subject": {subject},
		})
		m.SetHeader("To", to)
		m.SetBody("text/html", HTMLbody)

		d := gomail.NewDialer(e.Host, e.Port, e.Username, e.Password)

		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		if err := d.DialAndSend(m); err != nil {
			errchan <- err
			return
		}
	}(e.errchan)

}

func parseMailTemplate(p interface{}, templatePath string) (string, error) {

	parsedTemplate, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	err = parsedTemplate.Execute(buf, p)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

/*
SendRegisterVerifyMail(ctx context.Context, to, key string) (err error)
SendMails(ctx context.Context, to []string, subject, body string) (err error)
*/
func (e *Email) SendRegisterVerifyMail(ctx context.Context, name, surname, expdate, to, link string) (err error) {
	return e.SendWithTemplate("Register Verify Mail", struct {
		FirstName string
		LastName  string
		Link      string
		ExpDate   string
	}{
		FirstName: name,
		LastName:  surname,
		Link:      link,
		ExpDate:   expdate,
	}, "web/static/templates/register_mail_verify.html", to)

}

func (e *Email) SendMailChangeVerify(ctx context.Context, name, surname, expdate, to, link string) (err error) {
	return e.SendWithTemplate("Register Verify Mail", struct {
		FirstName string
		LastName  string
		Link      string
		ExpDate   string
	}{
		FirstName: name,
		LastName:  surname,
		Link:      link,
		ExpDate:   expdate,
	}, "web/static/templates/mail_verify.html", to)

}

func (e *Email) SendPasswordRecoveryMail(ctx context.Context, name, surname, expdate, to, link string) (err error) {
	return e.SendWithTemplate("Password Recovery", struct {
		FirstName string
		LastName  string
		Link      string
		ExpDate   string
	}{
		FirstName: name,
		LastName:  surname,
		Link:      link,
		ExpDate:   expdate,
	}, "web/static/templates/password_recovery.html", to)

}

func (e *Email) SendFeedbackSuccessMail(ctx context.Context, name, surname, feedback, created_at, to string) (err error) {
	return e.SendWithTemplate("Feedback Received", struct {
		FirstName string
		LastName  string
		Feedback  string
		CreatedAt string
	}{
		FirstName: name,
		LastName:  surname,
		Feedback:  feedback,
		CreatedAt: created_at,
	}, "web/static/templates/feedback_send_success.html", to)

}

func (e *Email) SendNotifyOldMail(ctx context.Context, name, surname, changedate, to string, changed bool) (err error) {

	return e.SendWithTemplate("Email", struct {
		FirstName  string
		LastName   string
		ChangeDate string
		Changed    bool
	}{
		FirstName:  name,
		LastName:   surname,
		Changed:    changed,
		ChangeDate: changedate,
	}, "web/static/templates/notify_old_mail.html", to)

}

func (e *Email) SendMails(ctx context.Context, to []string, subject, body string) (err error) {
	return e.SendWithTemplate(subject, struct {
		Body string
	}{
		Body: body,
	}, "web/static/templates/mail.html", to[0])

	// Bu geçiçi bir çözüm bunu düzeltmek lazım
}
