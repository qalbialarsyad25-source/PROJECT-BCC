package rest

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(app *gin.Engine, v1 *V1) {
	api := app.Group("/api/v1")
	{
		anak := api.Group("/anak")
		{
			anak.GET("", v1.GetDataAnak)
			anak.POST("", v1.CreateDataAnak)
			anak.DELETE("/:id", v1.DeleteDataAnak)
			anak.PATCH(":id", v1.EditDataAnak)
		}
	}
}
