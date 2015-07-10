package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/revel/revel"
)

type Result struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func OK() Result {
	return Result{true, "OK", nil}
}

func ERROR(msg string) Result {
	return Result{false, msg, nil}
}

func renderOk(c *revel.Controller, data interface{}) revel.Result {
	result := Result{true, "ok", data}
	return c.RenderJson(result)
}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {

	return c.RenderTemplate("index.html")
}

func readJson(req *http.Request, v interface{}) error {
	defer req.Body.Close()
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&v)
	revel.ERROR.Printf("Error while paring JSON-POST: %v", err)
	return err
}
