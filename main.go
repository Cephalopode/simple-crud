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
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
)

func main() {
	conf.Init()
	initLog()
	db.Init()
	if conf.C.Reset {
		models.CreateTable()
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/users", users.List)
	r.PUT("/users", users.Add)
	r.DELETE("users", users.Delete)
	_ = r.Run() // listen and serve on 0.0.0.0:8080
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
