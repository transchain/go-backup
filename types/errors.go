package types

import (
    "strings"
)

type CustomError struct {
    Message    string
    StackTrace string
}

func NewCustomError(message string, stackTrace string) *CustomError {
    return &CustomError{
        Message:    message,
        StackTrace: stackTrace,
    }
}

func (ce *CustomError) GetHtmlStackTrace() string {
    return strings.ReplaceAll(ce.StackTrace, "\n", "<br>")
}

type ProjectsCustomErrors map[string][]*CustomError

func (pce ProjectsCustomErrors) AddErrors(projectName string, customError ...*CustomError) {
    pce[projectName] = append(pce[projectName], customError...)
}
