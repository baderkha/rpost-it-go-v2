package mail

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sesv2"
)

// SendMailInput : input needed to send mail
type SendMailInput struct {
	// Subject : subject to put in the email
	Subject string
	// CC : cc others ?
	CC []*string
	// To : people interested
	To []*string
	// mail body
	Body string
}

// IMailer : a mailer interface we can generically call
type IMailer interface {
	// SendEmail : send email to client with a message , subject and body , this is raw unformatted
	SendEmail(mailInput *SendMailInput) bool
}

func New(region string, fromEmail string) IMailer {
	awsSession := session.New(&aws.Config{
		Region: aws.String(region),
	})
	sesSession := sesv2.New(awsSession)
	return &SesMail{
		driver:    sesSession,
		FromEmail: fromEmail,
	}
}
