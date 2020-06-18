package main

import (
	"html/template"
	"io"
)

/// Process
type Process struct {
	in  io.Reader
	out io.Writer

	inputString  string
	outputString string

	tpl  *template.Template
	data interface{}
}
