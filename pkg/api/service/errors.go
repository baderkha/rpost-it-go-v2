package service

import "fmt"

// serviceErrorTemplate : reusable erroring for the service to
//						  proxy the controller with correct http statuses
//						  encoded in the strings of the function errors
type serviceErrorTemplate struct {
	model string
}

// NotFoundResourceGeneric : error not found 404
func (e *serviceErrorTemplate) NotFoundResourceGeneric() error {
	return fmt.Errorf("404,`%s` resource was not found. Check your query", e.model)
}

// NotFoundResourceReason : error not found 404
func (e *serviceErrorTemplate) NotFoundResourceReason(reason string) error {
	return fmt.Errorf("404,`%s` resource was not found. Check your query. Detailed reason => %s", e.model, reason)
}

// InternalError : 500 db error
func (e *serviceErrorTemplate) InternalError() error {
	return fmt.Errorf("500,Fatal internal error. This is due to the database or store . Please contact the website admin")
}

// UserInputError : missing or incorrect field 400
func (e *serviceErrorTemplate) UserInputError(fields string, reason string) error {
	return fmt.Errorf("400,Bad or Missing Input for the following field `%s` from the request for the following reason `%s`", fields, reason)
}

// CustomError : your own custom error with cusstom status
func (e *serviceErrorTemplate) CustomError(statusCode uint, fmtSmt string, args ...interface{}) error {
	combinedStmt := fmt.Sprint(statusCode) + "," + fmtSmt
	return fmt.Errorf(combinedStmt, args...)
}

// UnAuthorized : use this for accounts that don't have access to do stuff
func (e *serviceErrorTemplate) UnAuthorized() error {
	return fmt.Errorf("401,The action on this resource `%s` is unauthorized for this request. Ensure you have the access privelage for this resource", e.model)
}

// ImATeaPot : figure out where to throw this ?
func (e *serviceErrorTemplate) ImATeaPot() error {
	return fmt.Errorf("418, I'm a Teapot !!!")
}
