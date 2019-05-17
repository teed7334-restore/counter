package controllers

import "github.com/gin-gonic/gin"

//Home 首頁
func Home(response *gin.Context) {
	response.String(200, "")
}
