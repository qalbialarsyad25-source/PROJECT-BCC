package rest

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(app *gin.Engine, v1 *V1) {
	api := app.Group("/api/v1")
	{

		auth := api.Group("/auth")
		{
			auth.GET("/google/login", v1.LoginGoogle)
			auth.GET("/google/callback", v1.CallbackGoogle)
		}

		anak := api.Group("/anak")
		{
			anak.GET("", v1.IMiddleware.Authentication, v1.GetDataAnak)
			anak.POST("", v1.IMiddleware.Authentication, v1.IMiddleware.Authorization("admin", "user", "dokter"), v1.CreateDataAnak)
			anak.DELETE("/:id", v1.DeleteDataAnak)
			anak.PATCH("/:id", v1.EditDataAnak)

			log := anak.Group("/:id/log")
			{
				log.GET("", v1.IMiddleware.Authentication, v1.GetLog)
				log.POST("", v1.IMiddleware.Authentication, v1.CreateLog)
				log.DELETE("/:logId", v1.DeleteLog)
			}

		}

		makanan := api.Group("/makanan")
		{
			makanan.GET("", v1.GetMakanan)
		}
	}

}
