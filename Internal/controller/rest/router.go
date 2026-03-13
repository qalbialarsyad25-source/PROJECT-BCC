package rest 

import (
	"github.com/gin-gonic/gin"
)

func NewRouter (app *gin.Engine, v1 *V1){
	api := app.Group("/api/v1"){
		
	}
}