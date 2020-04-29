package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"simple-crud/conf"
	"simple-crud/db"
	"simple-crud/internal/users"
	"simple-crud/models"

	nestedFormatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
)

func main() {
	conf.Init()
	r := setupRouter()
	_ = r.Run(fmt.Sprintf("0.0.0.0:%d", conf.C.Port)) // listen and serve on 0.0.0.0:8080
}

func setupRouter() *gin.Engine {
	initLog()
	db.Init()
	if conf.C.Reset {
		models.CreateTable()
	}
	if conf.C.IsProduction {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(static.ServeRoot("/", conf.C.DistDir))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	g := r.Group("/api")
	{
		g.GET("/users", users.List)
		g.PUT("/users", users.Add)
		g.DELETE("/users", users.Delete)
	}
	return r
}

func initLog() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	var formatter = &zt_formatter.ZtFormatter{
		Formatter: nestedFormatter.Formatter{
			HideKeys:        true,
			TimestampFormat: " ",
			TrimMessages:    false,
		},
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}
	logrus.SetFormatter(formatter)
}
