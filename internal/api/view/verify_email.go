package view

import "rpost-it-go/pkg/util/template"

var (
	templateEmail = template.New("web/templates/email/verify.html")
)

type EmailVerificationView struct {
	Name       string `json:"name"`
	VerifyLink string `json:"verifyLink"`
}

var _ IView = &EmailVerificationView{}

func (e *EmailVerificationView) Render() (string, error) {
	return templateEmail.Generate(e)
}
