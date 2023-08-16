package main

import (
	"errors"
	"fmt"
)

type Error struct {
	Message string
	Code    int
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Is(tgt error) bool {
	// This is never printed - this method never excutes for some reason
	fmt.Println("compared!")
	target, ok := tgt.(*Error)
	if !ok {
		return false
	}
	return e.Code == target.Code
}

var NotFoundError *Error = &Error{Code: 404, Message: "The page was not found"}

func NewError(errorType *Error, message string) error {
	rc := *errorType
	rc.Message = message

	return &rc
}

func FetchImage() error {
	return NewError(NotFoundError, "That image is gone")
}

func main() {
  err := FetchImage()
  fmt.Println(errors.Is(err, NotFoundError))
}
