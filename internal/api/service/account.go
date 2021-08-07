package service

import (
	"errors"
	"fmt"
	"rpost-it-go/internal/api/repo"
	"rpost-it-go/pkg/util/crypto"
	"rpost-it-go/pkg/util/jwt"
	"rpost-it-go/pkg/util/mail"
	"strings"
	"time"
	"unicode"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/badoux/checkmail"
	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	PasswordMinLength            = 8
	JWTVerificationExpiryMinutes = 5
	DefaultVerificationMessage   = "Account Has Been Verified Thank you !"
)

// AccountLoginJSON : models a login request
type AccountLoginJSON struct {
	AccountId string `json:"accountId"`
	Password  string `json:"password"`
}

// Account : Account Service
type Account struct {
	repo   repo.IAccountRepo
	er     serviceErrorTemplate
	hasher crypto.Hasher // password hasher
	BaseService
	mailer           mail.IMailer
	hashSecret       string
	jwtHasher        jwt.HS256
	verificationLink string
}

// isValidPassword : does checks on what we consider a good password
func (a *Account) isValidPassword(password string) error {
	if len(strings.Split(password, "")) < PasswordMinLength {
		return fmt.Errorf("expected the password to have %d characters", PasswordMinLength)
	}
	// taken from and thanks to https://gist.github.com/fearblackcat/d0199d6a47d60b067a4d4be173b0ef97
	// code by  fearblackcat
next:
	for name, classes := range map[string][]*unicode.RangeTable{
		"upper case": {unicode.Upper, unicode.Title},
		"lower case": {unicode.Lower},
		"numeric":    {unicode.Number, unicode.Digit},
		"special":    {unicode.Space, unicode.Symbol, unicode.Punct, unicode.Mark},
	} {
		for _, r := range password {
			if unicode.IsOneOf(classes, r) {
				continue next
			}
		}
		return fmt.Errorf("password must have at least one %s character", name)
	}
	return nil

}

func (a *Account) sendEmailVerification(acc *repo.Account) bool {
	token := a.getVerificationToken(acc)
	link := a.verificationLink + "/" + acc.ID + "?token=" + token
	// send verification email
	return a.mailer.SendEmail(&mail.SendMailInput{
		Subject: `Verify your Account Creation for post realm`,
		To:      aws.StringSlice([]string{acc.Email}),
		Body: fmt.Sprintf(`Hello %s ,
		You are recieving this email in order to verify your account creation with postrealm please click on this link %s in order to verify . Do not share this link with anyone !`,
			acc.ID,
			link,
		),
	})
}

func (a *Account) VerifyAccountCreation(accountId string, token string) (status string, err error) {
	if accountId == "" || token == "" {
		return "", a.Error().UnAuthorized()
	}

	acc := a.repo.FindById(accountId)
	if acc == nil || acc.ID != accountId {
		return "", a.Error().UnAuthorized()
	} else if acc.IsVerified {
		return DefaultVerificationMessage, nil
	}

	isValid, claim := a.jwtHasher.ValidateWebToken(token, a.getSecret(acc))
	if !isValid || claim.Subject != acc.ID {
		return "", a.Error().UnAuthorized()
	}
	acc.IsVerified = true
	_, _ = a.repo.Update(acc)
	return DefaultVerificationMessage, nil

}

func (a *Account) getSecret(acc *repo.Account) string {
	return fmt.Sprintf("%s:%s", a.hashSecret, acc.Password)
}

func (a *Account) getVerificationToken(acc *repo.Account) string {
	secret := a.getSecret(acc)
	token, _, _ := a.jwtHasher.GenerateWebToken(acc.ID, JWTVerificationExpiryMinutes, secret)
	return token
}

func (a *Account) isValidDateOfBirth(date time.Time) error {
	if date.Year() < 1900 {
		return errors.New("how are you even alive ?")
	} else if date.Year() > time.Now().Year() {
		return errors.New("so you're from the future ?")
	} else if time.Now().Year()-date.Year() < 13 {
		return errors.New("you're too young to be on this website")
	}
	return nil
}

func (a *Account) isValidEmail(email string) error {
	return checkmail.ValidateFormat(email)
}

func (a *Account) isAccountExist(accountIdUsername string) bool {
	acc := a.repo.FindById(accountIdUsername)
	return acc != nil && acc.ID == accountIdUsername
}

// GetByIdPublic : this is viewable for the public , and safe to consume
func (a *Account) GetByIdPublic(idHandle string) (*repo.AccountView, error) {
	acc := a.repo.FindById(idHandle)
	if acc == nil {
		return nil, a.er.NotFoundResourceGeneric()
	}
	return acc.GenerateView(), nil
}

// GetByIdFuzzy : do a fuzzy quick search for users by the id
func (a *Account) GetByIdFuzzy(id string) *[]repo.AccountView {
	results := a.repo.FindByLikeId(id)
	if results == nil {
		emptyResults := make([]repo.Account, 0)
		results = &emptyResults
	}
	accounts := repo.Accounts(*results)
	return accounts.GenerateView()
}

// Authenticate : Login to account and return status if logged in
func (a *Account) Authenticate(req *AccountLoginJSON) (*repo.AccountView, error) {
	err := a.ValidateInput(req, []*validation.FieldRules{
		validation.Field(&req.AccountId, validation.Required),
		validation.Field(&req.Password, validation.Required),
	})
	if err != nil {
		return nil, err
	}

	requestedAccount := a.repo.FindById(req.AccountId)
	if requestedAccount == nil || requestedAccount.ID != req.AccountId {
		return nil, a.Error().UnAuthorized()
	}

	if !a.hasher.CompareHash(req.Password, requestedAccount.Password) {
		return nil, a.Error().UnAuthorized()
	}
	return requestedAccount.GenerateView(), nil

}

func (a *Account) ValidateCreate(acc *repo.Account) error {
	// guard checks
	if a.isAccountExist(acc.ID) {
		return a.er.UserInputError("id", "this id already exists")
	}
	err := a.isValidPassword(acc.Password)
	if err != nil {
		return a.er.UserInputError("password", err.Error())
	}
	err = a.isValidDateOfBirth(acc.DateOfBirth)
	if err != nil {
		return a.er.UserInputError("dateOfBirth", err.Error())
	}
	err = a.isValidEmail(acc.Email)
	if err != nil {
		return a.er.UserInputError("email", err.Error())
	}
	return nil
}

func (a *Account) Create(acc *repo.Account) (*repo.AccountView, error) {

	err := a.ValidateCreate(acc)
	if err != nil {
		return nil, err
	}
	// change password to hashed one
	acc.Password = a.hasher.HashPassword(acc.Password)
	acc.IsVerified = false

	// create the user now
	acc, err = a.repo.Create(acc)
	if err != nil {
		return nil, a.er.InternalError()
	}

	isSuccess := a.sendEmailVerification(acc)
	if !isSuccess {
		return nil, a.Error().CustomError(500, "Could not send the Verification email , we are having trouble . Try again later")
	}

	return acc.GenerateView(), nil
}

func (a *Account) Update(acc *repo.Account) (*repo.AccountView, error) {
	return nil, nil
}

func (a *Account) Delete(id string) error {
	if !a.isAccountExist(id) {
		return a.er.NotFoundResourceGeneric()
	}
	err := a.repo.Delete(id)
	if err != nil {
		return a.er.InternalError()
	}
	return nil
}
