package main

import (
	"fmt"
	"testing"
)

func TestRender(t *testing.T) {
	render := NewRender()

	s, err := render.RenderString(`{{"hello"|upper}}`)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(s)
}
