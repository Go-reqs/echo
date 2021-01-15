package hellolog

import (
	"github.com/go-yaml/yaml"
	"github.com/labstack/echo"
	"os"
)

type Config struct {
	Greet  string `yaml:"greet"`
	Format string `yaml:"format"`
}

type Outer struct {
	Config Config `yaml:"log/hellolog"`
}

func InitWithYaml(ypath string, e *echo.Echo) error {
	f, err := os.Open(ypath)
	if err != nil {
		return err
	}
	outercfg := &Outer{}
	yaml.NewDecoder(f).Decode(outercfg)
	cfg := &outercfg.Config
	if cfg.Format != "" {
		e.Logger.SetHeader(cfg.Format)
	} else {
		e.Logger.SetHeader(cfg.Greet + " ${level} ${message}")
	}
	e.Logger.SetLevel(0)
	return nil
}
