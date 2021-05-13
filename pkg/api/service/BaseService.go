package service

import validation "github.com/go-ozzo/ozzo-validation"

// BaseService : Use Struct Embedding ,
// this will containt ht eerror emitter and some functionality around validation
type BaseService struct {
	er serviceErrorTemplate
}

// Error : Get the error Emitter
func (b *BaseService) Error() *serviceErrorTemplate {
	return &b.er
}

// ValidateInput : Validate user input against the rules
func (b *BaseService) ValidateInput(input interface{}, fieldRules []*validation.FieldRules) error {
	err := validation.ValidateStruct(
		input,
		fieldRules...,
	)
	if err != nil {
		return b.Error().CustomError(400, err.Error())
	}
	return nil
}
