package main

import (
	"github.com/Joshimmor/GoGinTestAPI/Computation/Comp"
    "github.com/gin-gonic/gin"
)
func main() {
    router := gin.Default()
    Comp := router.Group("/Comp")
	{
		Comp.GET("/",Comp.GetOutput)
	}

    router.Run("localhost:8080")
}
