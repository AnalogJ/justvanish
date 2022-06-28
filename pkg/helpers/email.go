package helpers

import (
	"errors"
	"fmt"
	"github.com/analogj/justvanish/pkg/models"
	"github.com/xhit/go-simple-mail/v2"
	"strings"
	"time"
)

func EmailSend(smtpConfig *models.SmtpConfig, userConfig *models.UserConfig, orgConfig *models.OrganizationConfig,
	regulationType string, actionType string, emailContent string) (error){

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
	if err != nil{
		return err
	}

	// New email simple html with inline and CC
	email := mail.NewMSG()
	email.SetFrom(smtpConfig.Username).
		SetSubject(fmt.Sprintf("%s %s - %s %s", strings.ToUpper(regulationType), strings.ToUpper(actionType), userConfig.FirstName, userConfig.LastName)).
		SetBody(mail.TextPlain, emailContent)

	// Add Receiver email addresses
	//
	email.AddTo("jason@thesparktree.com")

	// always check error after send
	if email.Error != nil{
		return email.Error
	}

	// Call Send and pass the client
	err = email.Send(smtpClient)
	if err != nil {
		return err
	} else {
		fmt.Println("Email Sent Successfully!")
	}
	return nil
}