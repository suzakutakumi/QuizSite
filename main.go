package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var startFlg bool

func main() {
	startFlg = false

	router := gin.Default()
	router.LoadHTMLGlob("docs/*.html")
	router.GET("/", index)
	router.GET("/Respondent", Respondent)
	router.POST("/GetAnswer", GetAnswer)
	router.GET("/SendAnswer", SendAnswer)
	router.POST("/Start", Start)
	router.GET("/CommonStart", CommonStart)
	router.Run("0.0.0.0:8080")
}

type Answer struct {
	Name       string `json:"name"`
	Content    string `json:"content"`
	AnswerTime int    `json:"time"`
}

var answers []Answer

func Start(ctx *gin.Context) {
	startFlg = true
	answers = []Answer{}
	ctx.Status(http.StatusOK)
}

func CommonStart(ctx *gin.Context) {
	msg := "NG"
	if startFlg {
		msg = "OK"
	}
	ctx.String(http.StatusOK, msg)
}

func index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
func Respondent(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "respondent.html", gin.H{})
}

func GetAnswer(ctx *gin.Context) {
	startFlg = false
	var newAnswer Answer
	ctx.BindJSON(&newAnswer)
	answers = append(answers, newAnswer)
	ctx.String(http.StatusOK, "")
}

func SendAnswer(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, answers)
}
