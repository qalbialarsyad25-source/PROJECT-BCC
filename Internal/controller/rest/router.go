package rest

import (
	websocket "bcc-geazy/internal/controller/delivery"

	"github.com/gin-gonic/gin"
)

func NewRouter(app *gin.Engine, v1 *V1, wsManager *websocket.WSManager) {
	wsHandler := websocket.NewWSHandler(wsManager)
	api := app.Group("/api/v1")
	{

		auth := api.Group("/auth")
		{
			auth.POST("/register", v1.Register)
			auth.POST("/login", v1.Login)
			auth.GET("/google/login", v1.LoginGoogle)
			auth.GET("/google/callback", v1.CallbackGoogle)
			auth.POST("/forgot-password", v1.ForgotPassword)
			auth.POST("/reset-password", v1.ResetPassword)
		}

		user := api.Group("/user")
		{
			user.GET("/profile", v1.IMiddleware.Authentication, v1.GetProfile)
			user.POST("/upload", v1.IMiddleware.Authentication, v1.UploadFotoUser)
		}

		anak := api.Group("/anak")
		{
			anak.GET("", v1.IMiddleware.Authentication, v1.GetDataAnak)
			anak.POST("", v1.IMiddleware.Authentication, v1.IMiddleware.Authorization("admin", "user"), v1.CreateDataAnak)
			anak.DELETE("/:id", v1.DeleteDataAnak)
			anak.PATCH("/:id", v1.EditDataAnak)

			nutrisi := anak.Group("/nutrisi")
			{
				nutrisi.GET("/:id", v1.IMiddleware.Authentication, v1.GetNutrisiHarian)
			}
			option := anak.Group("/option")
			{
				option.GET("/gender", v1.IMiddleware.Authentication, v1.GetGenderOptions)
				option.GET("/ke", v1.IMiddleware.Authentication, v1.GetAnakKeOptions)
				option.GET("/golongan", v1.IMiddleware.Authentication, v1.GetGolonganOption)
			}

			log := anak.Group("/:id/log")
			{
				log.GET("", v1.IMiddleware.Authentication, v1.GetLog)
				log.POST("", v1.IMiddleware.Authentication, v1.CreateLog)
				log.DELETE("/:logId", v1.DeleteLog)
			}

		}

		informasi := api.Group("/informasi")
		{
			informasi.GET("", v1.IMiddleware.Authentication, v1.GetInformasi)
			informasi.POST("", v1.IMiddleware.Authentication, v1.IMiddleware.Authorization("admin"), v1.CreateInformasi)
			informasi.DELETE("/:id", v1.IMiddleware.Authentication, v1.IMiddleware.Authorization("admin"), v1.DeleteInformasi)
			informasi.PATCH("/:id", v1.IMiddleware.Authentication, v1.IMiddleware.Authorization("admin"), v1.EditInformasi)
		}

		logInformasi := api.Group("/log-informasi")
		{
			logInformasi.GET("", v1.IMiddleware.Authentication, v1.GetLogInformasi)
			logInformasi.POST("", v1.IMiddleware.Authentication, v1.CreateLogInformasi)
			logInformasi.DELETE("/:id", v1.IMiddleware.Authentication, v1.DeleteLogInformasi)
		}

		makanan := api.Group("/makanan")
		{
			makanan.GET("", v1.IMiddleware.Authentication, v1.GetMakanan)
			makanan.POST("", v1.IMiddleware.Authentication, v1.IMiddleware.Authorization("admin"), v1.CreateMakanan)
			makanan.PATCH("/:id", v1.IMiddleware.Authentication, v1.IMiddleware.Authorization("admin"), v1.EditMakanan)
			makanan.DELETE("/:id", v1.IMiddleware.Authentication, v1.IMiddleware.Authorization("admin"), v1.DeleteMakanan)
		}

		konsultasi := api.Group("/konsultasi")
		{
			konsultasi.GET("", v1.IMiddleware.Authentication, v1.GetKonsultasi)
			konsultasi.POST("", v1.IMiddleware.Authentication, v1.IMiddleware.Authorization("admin", "user", "dokter"), v1.CreateKonsultasi)
		}

		dokter := api.Group("/dokter")
		{
			dokter.GET("", v1.IMiddleware.Authentication, v1.GetDokter)
			dokter.GET("/:id", v1.IMiddleware.Authentication, v1.GetDetailDokter)
			dokter.POST("", v1.IMiddleware.Authentication, v1.IMiddleware.Authorization("admin"), v1.CreateDokter)
			dokter.DELETE("/:id", v1.IMiddleware.Authentication, v1.IMiddleware.Authorization("admin"), v1.DeleteDokter)
			dokter.PATCH("/:id", v1.IMiddleware.Authentication, v1.IMiddleware.Authorization("admin"), v1.EditDokter)
		}

		notifikasi := api.Group("/notifikasi")
		{
			notifikasi.GET("", v1.GetNotifikasi)
			notifikasi.POST("", v1.CreateNotifikasi)
			notifikasi.DELETE("/:id", v1.DeleteNotifikasi)
			notifikasi.PATCH("/:id", v1.EditNotifikasi)
		}

	}

	app.GET("/ws/chat", wsHandler.HandleWS)

}
