package chaincodeErrors

import (
	"fmt"
	"runtime/debug"
)

type ChaincodeError interface {
	error
	Info() Payload
	Log() string
	Error() string
}

type Payload struct {
	ID      string      `json:"id"`
	Status  int32       `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

const wrapperString = "Error: %s.\nErrorStatus: %d.\nErrorMessage: %s.\nDebugStack: %s"

type MarshallingError struct {
	FuncName     string
	WrappedError error
	ErrorData    Payload
	Stack        []byte
}

func NewMarshallingError(funcName string, structName string, err error) *MarshallingError {
	const msgTemplate = "%s: Não foi possível converter a struct %s."
	return &MarshallingError{
		ErrorData: Payload{
			ID:      "MarshallError",
			Status:  400,
			Message: fmt.Sprintf(msgTemplate, funcName, structName),
			Payload: nil,
		},
		FuncName:     funcName,
		WrappedError: err,
		Stack:        debug.Stack(),
	}
}

func (err *MarshallingError) Info() Payload {
	return err.ErrorData
}

func (err *MarshallingError) Log() string {
	return fmt.Sprintf(wrapperString, err.ErrorData.ID, err.ErrorData.Status, err.ErrorData.Message, err.Stack)
}

func (err *MarshallingError) Error() string {
	return err.ErrorData.Message
}

func (err *MarshallingError) Unwrap() error {
	return err.WrappedError
}

type ValidationError struct {
	FuncName     string
	WrappedError error
	ErrorData    Payload
	Stack        []byte
}

var validationErrors = map[string]string{
	"": "%s:",
}

func NewValidationError(funcName string, errID string, extendedMsg string, err error) *ValidationError {
	var msgTemplate = validationErrors[errID]

	var errorMsg = ""
	if extendedMsg == "" {
		errorMsg = fmt.Sprintf(msgTemplate, funcName)
	} else {
		errorMsg = fmt.Sprintf(msgTemplate, funcName, extendedMsg)
	}

	return &ValidationError{
		ErrorData: Payload{
			ID:      "ValidationError",
			Status:  400,
			Message: errorMsg,
			Payload: nil,
		},
		FuncName:     funcName,
		WrappedError: err,
		Stack:        debug.Stack(),
	}
}

func (err *ValidationError) Info() Payload {
	return err.ErrorData
}

func (err *ValidationError) Log() string {
	return fmt.Sprintf(wrapperString, err.ErrorData.ID, err.ErrorData.Status, err.ErrorData.Message, err.Stack)
}

func (err *ValidationError) Error() string {
	return err.ErrorData.Message
}

func (err *ValidationError) Unwrap() error {
	return err.WrappedError
}
