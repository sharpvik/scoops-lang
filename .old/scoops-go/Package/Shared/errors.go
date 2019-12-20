package shared

/*
 * This file exists so that it's easier to track error tags that are used in
 * Scoops. Every legitimate error tag is to be declared here and used as a
 * const value from the package 'shared' like so: shared.ErrorTag.
 */
const (
    IndexError      string = "OpcodeError"
    OpcodeError     string = "OpcodeError"
    RuntimeError    string = "RuntimeError"
    TypeError       string = "TypeError"
)
