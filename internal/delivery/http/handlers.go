package http

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (a *App) getMain(context *gin.Context) {
	result, _ := context.Cookie("result")
	context.SetCookie("result", "", -1, "/", "", false, true)

	context.HTML(200, "main.html", gin.H{
		"result": result,
	})

}

func (a *App) getSettings(context *gin.Context) {
	idInstance := context.Query("idInstance")
	ApiTokenInstance := context.Query("ApiTokenInstance")

	result, err := a.greenService.GetSettings(idInstance, ApiTokenInstance)
	if err != nil {
		context.SetCookie("result", err.Error(), 60, "/", "", false, true)
		context.Redirect(302, "/")
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		context.SetCookie("result", err.Error(), 60, "/", "", false, true)
		context.Redirect(302, "/")
	}

	context.SetCookie("result", string(jsonData), 60, "/", "", false, true)
	context.Redirect(302, "/")

}

func (a *App) getStateInstance(context *gin.Context) {
	idInstance := context.Query("idInstance")
	ApiTokenInstance := context.Query("ApiTokenInstance")

	result, err := a.greenService.GetStateInstance(idInstance, ApiTokenInstance)
	if err != nil {
		context.SetCookie("result", err.Error(), 60, "/", "", false, true)
		context.Redirect(302, "/")
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		context.SetCookie("result", err.Error(), 60, "/", "", false, true)
		context.Redirect(302, "/")
	}

	context.SetCookie("result", string(jsonData), 60, "/", "", false, true)
	context.Redirect(302, "/")
}

func (a *App) sendMessage(context *gin.Context) {
	idInstance := context.PostForm("idInstance")
	ApiTokenInstance := context.PostForm("ApiTokenInstance")

	chatId := context.PostForm("chatId")
	message := context.PostForm("message")

	result, err := a.greenService.SendMessage(idInstance, ApiTokenInstance, chatId, message)
	if err != nil {
		context.SetCookie("result", err.Error(), 60, "/", "", false, true)
		context.Redirect(302, "/")
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		context.SetCookie("result", err.Error(), 60, "/", "", false, true)
		context.Redirect(302, "/")
	}

	context.SetCookie("result", string(jsonData), 60, "/", "", false, true)
	context.Redirect(302, "/")
}

func (a *App) sendFileByUrl(context *gin.Context) {
	idInstance := context.PostForm("idInstance")
	ApiTokenInstance := context.PostForm("ApiTokenInstance")

	chatId := context.PostForm("chatId2")
	fileUrl := context.PostForm("fileUrl")

	result, err := a.greenService.SendFileByUrlResponse(idInstance, ApiTokenInstance, chatId, fileUrl)
	fmt.Println(result)
	fmt.Println("_________")
	fmt.Println(err)
	if err != nil {
		context.SetCookie("result", err.Error(), 60, "/", "", false, true)
		context.Redirect(302, "/")
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	fmt.Println(jsonData)
	fmt.Println("_________")
	fmt.Println(err)
	if err != nil {
		context.SetCookie("result", err.Error(), 60, "/", "", false, true)
		context.Redirect(302, "/")
	}

	context.SetCookie("result", string(jsonData), 60, "/", "", false, true)
	context.Redirect(302, "/")
}
