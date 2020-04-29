package conf

import (
	"fmt"

	"github.com/koding/multiconfig"
)

type Server struct {
	IsProduction bool   `default:"false"`
	Reset        bool   `default:"false"`
	Port         int    `default:"8080"`
	DistDir      string `default:"frontend/dist"`
	Database     struct {
		Host     string `default:"127.0.0.1"`
		Port     int    `default:"3306"`
		User     string `default:"root"`
		Password string
		DB       string `default:"simple_crud"`
		Debug    bool
	}
}

var C *Server

func Init() {
	m := multiconfig.NewWithPath("build/config.toml") // supports TOML, JSON and YAML

	C = new(Server)

	m.MustLoad(C) // Panic's if there is any error
	fmt.Printf("configuration: %+v\n\n", C)

}
