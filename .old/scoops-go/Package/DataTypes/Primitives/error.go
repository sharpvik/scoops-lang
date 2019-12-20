package primitives

import (
    "bufio"
    "errors"
    "fmt"
)



/*
 * Error consists of two important parts: tag and message.
 * Error tag is like a type of error -- it allows us to track what exactly went
 * wrong. Error message is pretty self-explanatory -- we use it to give some
 * details about the error.
 */
type Error struct {
    Tag     string
    Message error
}


func NewError(tag, message string) *Error {
    return &Error{tag, errors.New(message)}
}


func FromError(tag string, message error) *Error {
    return &Error{tag, message}
}


func (e *Error) Clone() *Error {
    return FromError(e.Tag, e.Message)
}


func (e *Error) Print(w *bufio.Writer) {
    w.WriteString( e.ToGoString() )
    w.WriteByte('\n')
    w.Flush()
}


func (e *Error) ToGoString() string {
    return fmt.Sprintf("%s: %s", e.Tag, e.Message)
}


func (e *Error) ToGoError() error {
    return errors.New( e.ToGoString() )
}


func (e *Error) Type() string {
    return "err"
}