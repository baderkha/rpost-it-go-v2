package mail

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sesv2"
	"github.com/aws/aws-sdk-go/service/sesv2/sesv2iface"
	"github.com/davecgh/go-spew/spew"
)

type SesMail struct {
	driver    sesv2iface.SESV2API
	FromEmail string `json:"from"`
}

func (s *SesMail) SendEmail(mailInput *SendMailInput) bool {

	_, err := s.driver.SendEmail(&sesv2.SendEmailInput{
		Destination: &sesv2.Destination{
			CcAddresses: mailInput.CC,
			ToAddresses: mailInput.To,
		},
		FromEmailAddress: &s.FromEmail,
		Content: &sesv2.EmailContent{
			Simple: &sesv2.Message{
				Body: &sesv2.Body{
					Text: &sesv2.Content{
						Data:    aws.String(mailInput.Body),
						Charset: aws.String("UTF-8"),
					},
				},
				Subject: &sesv2.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(mailInput.Subject),
				},
			},
		},
	},
	)
	spew.Dump(mailInput)
	spew.Dump(err)
	return err == nil
}
