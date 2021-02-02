package moqup

import (
	"html/template"
	"github.com/lambda-platform/moqup/handler"
	"github.com/lambda-platform/moqup/utils"
	"github.com/labstack/echo/v4"
	lambdaUtils "github.com/lambda-platform/lambda/utils"

)


func Set(e *echo.Echo){

	templates := lambdaUtils.GetTemplates(e)
	AbsolutePath := utils.AbsolutePath()
	templates["moqup.html"] = template.Must(template.ParseFiles(AbsolutePath+"templates/moqup.html"))


	e.GET("/moqup/:id", handler.Moqup)



}

