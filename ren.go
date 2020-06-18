package main

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"html/template"
)

type Render struct {
	tpl  *template.Template
	data map[string]interface{}
}

func NewRender() *Render {
	return &Render{
		tpl:  template.New("base").Funcs(sprig.FuncMap()),
		data: make(map[string]interface{}),
	}
}

func (r *Render) RenderString(content string) (string, error) {
	t, err := r.tpl.Parse(content)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBufferString("")
	err = t.Execute(buf, r.data)
	return buf.String(), err
}
