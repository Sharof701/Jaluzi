package api

import (
	"jaluzi/api/handler"
	"jaluzi/config"
	"jaluzi/pkg/logger"
	"jaluzi/storage"

	_ "jaluzi/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageI, logger logger.LoggerI) {
	h := handler.NewHandler(cfg, storage, logger)
	r.Use(customCORSMiddleware())
	v1 := r.Group("/jaluzi/api/v1")

	v1.POST("/admin", h.CreateAdmin)
	v1.GET("/admin/:id", h.GetByIdAdmin)
	v1.GET("/admin", h.GetListAdmin)
	v1.PUT("/admin/:id", h.UpdateAdmin)
	v1.DELETE("/admin/:id", h.DeleteAdmin)

	v1.POST("/product", h.CreateProduct)
	v1.GET("/product/:id", h.GetByIdProduct)
	v1.GET("/product", h.GetListProduct)
	v1.PUT("/product/:id", h.UpdateProduct)
	v1.DELETE("/product/:id", h.DeleteProduct)

	v1.POST("upload-files", h.UploadFiles)
	v1.DELETE("delete-file", h.DeleteFile)

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
