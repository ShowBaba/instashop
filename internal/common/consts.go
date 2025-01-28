package common

import "net/http"

type RestErr struct {
	Message    string `json:"message"`
	Success    bool   `json:"success"`
	StatusCode int    `json:"code"`
}

func (r *RestErr) BadRequest(message string) *RestErr {
	return &RestErr{
		Message:    message,
		Success:    false,
		StatusCode: http.StatusBadRequest,
	}
}

func (r *RestErr) NotFound(message string) *RestErr {
	return &RestErr{
		Message:    message,
		Success:    false,
		StatusCode: http.StatusNotFound,
	}
}

func (r *RestErr) ServerError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		Success:    false,
		StatusCode: http.StatusInternalServerError,
	}
}

func (r *RestErr) RequestNotAllowed(message string) *RestErr {
	return &RestErr{
		Message:    message,
		Success:    false,
		StatusCode: http.StatusForbidden,
	}
}

func NewRestErr() *RestErr {
	return &RestErr{}
}

const (
	ErrBadRequest                = "bad request"
	ErrSomethingWentWrong        = "something went wrong, please try again"
	ErrEmailAlreadyInUse         = "email already in use"
	ErrUserWithEmailNotFound     = "user with email not found"
	ErrInvalidPassword           = "invalid password"
	ErrMissingAuthTokenInHeader  = "missing auth token in header"
	ErrInvalidAuthToken          = "invalid auth token"
	ErrFailToParseReqBody        = "failed to parse request body"
	ErrInsufficientFunds         = "insufficient funds"
	ErrInvalidTransactionType    = "invalid transaction type"
	ErrInsufficientStock         = "insufficient stock for product"
	ErrProductNotFound           = "product not found"
	ErrInvalidOrder              = "invalid order"
	ErrCanOnlyCancelPendingOrder = "only orders in Pending status can be canceled"
	ErrOrderNotFound             = "order not found"
	ErrUnauthorized              = "unauthorized"
)
