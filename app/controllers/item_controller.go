package controllers

import (
	"github.com/revel/revel"
	"github.com/tomoyane/revel-performance/app/domains"
	"github.com/tomoyane/revel-performance/app"
)

type JsonResponse struct {
	Response interface{} `json:"items"`
}

type ItemController struct {
	*revel.Controller
}

func (c ItemController) Get() revel.Result {
	var items []domains.Items

	app.Db.Find(&items)

	response := JsonResponse{}
	response.Response = items

	return c.RenderJSON(response)
}
