package conf

import (
	"fmt"

	"github.com/koding/multiconfig"
)

type Server struct {
	Reset    bool `default:"false"`
	Port     int  `default:"8080"`
	Database struct {
		Host     string `default:"127.0.0.1"`
		Port     int    `default:"3306"`
		User     string `default:"root"`
		Password string
		DB       string `default:"simple_crud"`
	}
}

var C *Server

func Init() {
	m := multiconfig.NewWithPath("build/config.toml") // supports TOML, JSON and YAML

	C = new(Server)

	m.MustLoad(C) // Panic's if there is any error
	fmt.Println("conf: %+v", C)

}
