package controllers

import "github.com/revel/revel"

type HelloController struct {
	*revel.Controller
}

func (c HelloController) Get() revel.Result {
	hello := map[string]string{"key": "hello world"}
	return c.RenderJSON(hello)
}
