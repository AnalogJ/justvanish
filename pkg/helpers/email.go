package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/analogj/justvanish/pkg/models"
	"github.com/xhit/go-simple-mail/v2"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func EmailSend(smtpConfig *models.SmtpConfig, userConfig *models.UserConfig, orgConfig *models.OrganizationConfig,
	regulationType string, actionType string, emailContent string) error {

	if len(smtpConfig.Username) == 0 || len(smtpConfig.Password) == 0 {
		return errors.New("SMTP Username & Password required")
	}

	server := mail.NewSMTPClient()

	// SMTP Server
	server.Host = smtpConfig.Hostname
	server.Port = smtpConfig.Port
	server.Username = smtpConfig.Username
	server.Password = smtpConfig.Password
	server.Encryption = mail.EncryptionSTARTTLS

	// Variable to keep alive connection
	server.KeepAlive = false

	// Timeout for connect to SMTP Server
	server.ConnectTimeout = 10 * time.Second

	// Timeout for send the data and wait respond
	server.SendTimeout = 10 * time.Second

	// SMTP client
	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}

	// New email simple html with inline and CC
	email := mail.NewMSG()
	email.SetFrom(smtpConfig.Username).
		SetSubject(fmt.Sprintf("%s %s - %s %s", strings.ToUpper(regulationType), strings.ToUpper(actionType), userConfig.FirstName, userConfig.LastName)).
		SetBody(mail.TextPlain, emailContent)

	// Add Receiver email addresses
	for _, emailInfo := range orgConfig.Contact.Email {
		email.AddTo(emailInfo.Address)
	}

	// always check error after send
	if email.Error != nil {
		return email.Error
	}

	// Call Send and pass the client
	err = email.Send(smtpClient)
	if err != nil {
		return err
	} else {
		fmt.Println(fmt.Sprintf("Email Sent Successfully! (%s)", orgConfig.OrganizationName))
	}
	return nil
}

type etherealAccount struct {
	Username string `json:"user"`
	Password string `json:"pass"`
	Website  string `json:"web"`
	Status   string `json:"status"`

	Imap struct {
		Host   string `json:"host"`
		Port   int    `json:"port"`
		Secure bool   `json:"secure"`
	} `json:"imap"`
	Pop3 struct {
		Host   string `json:"host"`
		Port   int    `json:"port"`
		Secure bool   `json:"secure"`
	} `json:"pop3"`
	Smtp struct {
		Host   string `json:"host"`
		Port   int    `json:"port"`
		Secure bool   `json:"secure"`
	} `json:"smtp"`
}

// generate a smtpConfig for https://ethereal.email/
// https://github.com/nodemailer/nodemailer/blob/3491486281ea2e2cba9a07d4df14d136f6ebb153/lib/nodemailer.js#L58-L124
// Ethereal is a fake SMTP service, mostly aimed at Nodemailer and EmailEngine users (but not limited to).
// It's a completely free anti-transactional email service where messages never get delivered.
func EmailTestSmtpConfig() (*models.SmtpConfig, error) {

	values := map[string]string{"requestor": "justvanish", "version": "0.0.1"}
	jsonData, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post("https://api.nodemailer.com/user", "application/json",
		bytes.NewBuffer(jsonData))

	if err != nil {
		return nil, err
	}

	var res etherealAccount
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	if res.Status != "success" {
		return nil, errors.New("could not create test account on https://ethereal.email ")
	}

	file, _ := json.MarshalIndent(res, "", " ")
	_ = ioutil.WriteFile("ethereal.credentials.json", file, 0644)

	return &models.SmtpConfig{
		Hostname: res.Smtp.Host,
		Port:     res.Smtp.Port,
		Username: res.Username,
		Password: res.Password,
	}, nil
}
