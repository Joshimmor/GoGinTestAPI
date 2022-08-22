package Comp

import (
	"net/http"
	// "github.com/Joshimmor/GoGinTestAPI/Models"
	"github.com/gin-gonic/gin"
)

func GetOutput(c *gin.Context){
	c.JSON(http.StatusCreated,"Hello")
}