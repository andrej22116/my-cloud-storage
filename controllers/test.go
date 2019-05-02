package controllers

import "github.com/kataras/iris/mvc"

type IndexController struct{}

func (c *IndexController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1> <a href=\"/help\"> U can go to HELP!</a>",
	}
}

type HelpController struct {
}

func (c *HelpController) GetHelp() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Help page!</h1> <a href=\"/help/ping\"> test ping </a>",
	}
}

func (c *HelpController) GetHelpPing() string {
	return "pong"
}

/*
func (testController *TestController) Put() {}

func (testController *TestController) Delete() {}
*/
