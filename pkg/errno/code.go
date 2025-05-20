// Package errno defines system error code and error type.
package errno

var (

	// Code format
	// 1: 1 server error, 2 client error
	// 00: cgi/function code
	// 00: error code

	// Common errors
	OK = &Errno{Code: 0, Message: "OK"}
	// Client error
	InvalidIP = &Errno{Code: 20003, Message: "Invalid IP address"}
	// Special Error
	SampleError = &Errno{Code: 383838, Message: "Just Example"}
)
