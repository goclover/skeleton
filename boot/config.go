package boot

import (
	"github.com/BurntSushi/toml"
)

var cs = &config{}

type config struct {
	AppName    string     `toml:"AppName"`
	AppMode    string     `toml:"AppMode"`
	AppIDC     string     `toml:"AppIDC"`
	HTTPServer httpServer `toml:"HTTPServer"`
}

type httpServer struct {
	Addr string `toml:"Addr"`
}

func MustLoadConfig(c string) *config {
	if _, err := toml.DecodeFile(c, cs); err != nil {
		panic(err.Error())
	}
	return cs
}

func Config() *config {
	return cs
}
