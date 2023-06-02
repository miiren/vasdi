package router

import (
	"github.com/gin-gonic/gin"
	"runtime/debug"

	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type GinWebRun struct {
	Srv         *http.Server
	AppBasePath string
}

func (g *GinWebRun) New(appBasePath string) *GinWebRun {
	g.Srv = &http.Server{
		Addr:           ":8080",
		Handler:        routers(appBasePath),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return g
}

func (g *GinWebRun) Run() error {
	go func() {
		defer func() {
			if p := recover(); p != nil {
				fmt.Println("recover:", p)
			}
		}()
		//这里还是有些问题的 外层无法获取内层是否err
		gin.SetMode(gin.ReleaseMode)
		// service connections
		if err := g.Srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := g.Srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
	return nil
}

func routers(appBasePath string) *gin.Engine {
	r := gin.New()
	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	//处理异常
	r.NoRoute(handleNotFound)
	r.NoMethod(handleNotFound)
	//限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	//设置静态资源路径
	r.Static("/static/", "resources/static/")
	//设置ico
	r.StaticFile("/favicon.ico", "resources/static/favicon.ico")
	//设置html路径
	r.LoadHTMLGlob(appBasePath + "view/html/**/*")
	r.GET("/", Index) //home
	r.GET("/testTimeout", Timeout)
	r.GET("/testPanic", TestPanic)
	return r
}

func Index(c *gin.Context) {
	//c.JSON(200, data)
	c.JSON(200, gin.H{"data": "这是一个接口系统"})
}
func Timeout(c *gin.Context) {
	fmt.Println(time.Now().String())
	x := <-time.After(5 * time.Second)
	c.JSON(200, gin.H{
		"msg": x,
	})
}

func TestPanic(c *gin.Context) {
	panic("An unexpected error happen!")
}

// 404
func handleNotFound(c *gin.Context) {
	log.Printf("handle not found: %v , stack: %v", c.Request.RequestURI, string(debug.Stack()))
	c.HTML(http.StatusNotFound, "error/404.html", nil)
}
