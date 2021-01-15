package preset

import (
	_ "fmt"
	"github.com/go-yaml/yaml"
	"github.com/labstack/echo"
	"os"

	"github.com/go-reqs/echo/preset/log/hellolog"
	"github.com/go-reqs/echo/preset/log/rotatelog"
)

type Config struct {
	Objs []string `yaml:"Preset"`
}

func getDefaultYamlPath() string {
	var path string = "conf/setting.yaml"
	return path
}

func Init(e *echo.Echo) error {
	y := getDefaultYamlPath()
	return InitWithYaml(y, e)
}

func InitWithYaml(ypath string, e *echo.Echo) error {
	var l logger = GetLogger()
	f, err := os.Open(ypath)
	if err != nil {
		return err
	}
	cfg := &Config{}
	yaml.NewDecoder(f).Decode(cfg)
	l.Logf("Config %s", *cfg)
	for _, obj := range cfg.Objs {
		switch obj {
		case "log/hellolog":
			err = hellolog.InitWithYaml(ypath, e)
		case "log/rotatelog":
			err = rotatelog.InitWithYaml(ypath, e)
		}
	}
	return err
}
