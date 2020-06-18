package main

import (
	"html/template"
	"io"
)

/// Process
type Process struct {
	src  string
	in   io.Reader
	dest string
	out  io.Writer

	inputString  string
	outputString string

	tpl     *template.Template
	data    interface{}
	closers []io.Closer
}
