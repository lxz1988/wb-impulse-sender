package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wb-impulse-sender/app/http/controller"
	"wb-impulse-sender/app/http/middleware"
)

func Bootstrap() {
	//router := gin.Default()
	router := gin.New()
	router.Use(middleware.LoggerHandler, middleware.RecoverHandler)
	//注册路由
	register(router)
	//监听
	listen(router)
}

//注册路由
func register(router *gin.Engine)  {
	router.GET("/pond/number", controller.NewNumber().GetNumber)
}

//监听路由
func listen(router *gin.Engine)  {
	srv := &http.Server{
		Addr:    ":8991",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("listen error:", err)
		}
	}()
}
