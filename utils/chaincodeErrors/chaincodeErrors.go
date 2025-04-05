package chaincodeErrors

import (
	"fmt"
	"runtime/debug"
)

type ChaincodeError interface {
	error
	Error() string
}

const wrapperString = "Error: %s.\nDebugStack: %s"

type MarshallingError struct {
	FuncName     string
	Message      string
	WrappedError error
	Stack        []byte
}

func NewMarshallingError(funcName string, structName string, err error) *MarshallingError {
	const msgTemplate = "[%s] Não foi possível converter a struct [%s]."
	return &MarshallingError{
		FuncName:     funcName,
		Message:      msgTemplate,
		WrappedError: err,
		Stack:        debug.Stack(),
	}
}

func (err *MarshallingError) Error() string {
	return fmt.Sprintf(wrapperString, err.FuncName, err.Stack)
}

func (err *MarshallingError) Unwrap() error {
	return err.WrappedError
}

//**************************************************************************

type ValidationError struct {
	FuncName     string
	Message      string
	WrappedError error
	Stack        []byte
}

func NewValidationError(funcName string, errID string, extendedMsg string, err error) *ValidationError {

	const msgTemplate = "[%s] Erro de validação: %s"
	var errorMsg = ""
	if extendedMsg == "" {
		errorMsg = fmt.Sprintf(msgTemplate, funcName)
	} else {
		errorMsg = fmt.Sprintf(msgTemplate, funcName, extendedMsg)
	}

	return &ValidationError{
		FuncName:     funcName,
		Message:      errorMsg,
		WrappedError: err,
		Stack:        debug.Stack(),
	}
}

func (err *ValidationError) Error() string {
	return err.Message
}

func (err *ValidationError) Unwrap() error {
	return err.WrappedError
}

//**************************************************************************

type ReadWorldStateError struct {
	FuncName     string
	Message      string
	WrappedError error
	Stack        []byte
}

func NewReadWorldStateError(funcName string, err error) *ReadWorldStateError {
	const msgTemplate = "[%s] Não foi possivel ler o World State"
	return &ReadWorldStateError{
		FuncName:     funcName,
		WrappedError: err,
		Stack:        debug.Stack(),
	}
}

func (err *ReadWorldStateError) Error() string {
	return err.Message
}

func (err *ReadWorldStateError) Unwrap() error {
	return err.WrappedError
}

//**************************************************************************

type WriteWorldStateError struct {
	FuncName     string
	Message      string
	WrappedError error
	Stack        []byte
}

func NewWriteWorldStateError(funcName string, err error) *WriteWorldStateError {
	const msgTemplate = "[%s] Não foi possivel escrever no World State"
	return &WriteWorldStateError{
		FuncName:     funcName,
		Message:      fmt.Sprintf(msgTemplate, funcName),
		WrappedError: err,
		Stack:        debug.Stack(),
	}
}

func (err *WriteWorldStateError) Error() string {
	return err.Message
}

func (err *WriteWorldStateError) Unwrap() error {
	return err.WrappedError
}

//**************************************************************************

type UpdateWorldStateError struct {
	FuncName     string
	Message      string
	WrappedError error
	Stack        []byte
}

func NewUpdateWorldStateError(funcName string, err error) *UpdateWorldStateError {
	const msgTemplate = "[%s] Não foi possivel fazer update no World State"
	return &UpdateWorldStateError{
		FuncName:     funcName,
		Message:      fmt.Sprintf(msgTemplate, funcName),
		WrappedError: err,
		Stack:        debug.Stack(),
	}
}

func (err *UpdateWorldStateError) Error() string {
	return err.Message
}

func (err *UpdateWorldStateError) Unwrap() error {
	return err.WrappedError
}

//**************************************************************************

type GenericError struct {
	FuncName     string
	Message      string
	WrappedError error
	Stack        []byte
}

func NewGenericError(funcName string, err error) *GenericError {
	const msgTemplate = "[%s] Não foi possivel fazer update no World State"
	return &GenericError{
		FuncName:     funcName,
		Message:      fmt.Sprintf(msgTemplate, funcName),
		WrappedError: err,
		Stack:        debug.Stack(),
	}
}

func (err *GenericError) Error() string {
	return err.Message
}

func (err *GenericError) Unwrap() error {
	return err.WrappedError
}
