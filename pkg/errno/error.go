package errno

import (
"fmt"
)

// Errno error info provide to client
type Errno struct {
	Code    int
	Message string
	Err     error
}

// Error return error info
func (err Errno) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %v", err.Code, err.Message, err.Err)
}

// New create a errno struct
func New(errno *Errno, err error) *Errno {
	errno.Err = err
	return errno
}

// Add append error message
func (err *Errno) Add(message string) error {
	err.Message += " " + message
	return err
}

// Addf append formatted message
func (err *Errno) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

// DecodeErr decode err into error code and message
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch et := err.(type) {
	case *Errno:
		return et.Code, et.Message
	default:
		return InternalServerError.Code, err.Error()
	}
}
