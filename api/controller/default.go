package controller

import (
	"api/model"
	"fmt"
	"net/http"
)

type Default struct {
	*Base
}

func NewDefault() *Default {
	c := &Default{
		Base: &Base{
			Routes: model.Routes{},
			Prefix: "/",
		},
	}
	c.Routes = append(c.Routes, model.Route{
		Name:    "Index",
		Method:  "GET",
		Pattern: c.Prefix,
		Func:    c.index,
	})
	return c
}

func (d *Default) index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
