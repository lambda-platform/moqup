package moqup

import (
	"html/template"
	"github.com/lambda-platform/moqup/handler"
	"github.com/lambda-platform/moqup/utils"
	"github.com/labstack/echo/v4"
	lambdaUtils "github.com/lambda-platform/lambda/utils"
	templateUtils "github.com/lambda-platform/moqup/utils"
)


func Set(e *echo.Echo){

	templates := lambdaUtils.GetTemplates(e)
	AbsolutePath := utils.AbsolutePath()
	TemplatePath := templateUtils.AbsolutePath()
	templates["moqup.html"] = template.Must(template.ParseFiles(AbsolutePath+"templates/moqup.html"))

	templates["adminmodule.html"] = template.Must(template.ParseFiles(
		TemplatePath+"views/paper.html",
	))
	e.GET("/pages/moqup/:id", handler.Moqup)



}

