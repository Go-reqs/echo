package rotatelog

import (
	"errors"
	"github.com/go-yaml/yaml"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/natefinch/lumberjack"
	"os"
)

type Config struct {
	Prefix    string            `yaml:"Prefix"`
	Level     int               `yaml:"Level"`
	Header    string            `yaml:"Header"`
	RotateLog lumberjack.Logger `yaml:"RotateLog"`
}

type Outer struct {
	Config Config `yaml:"log/rotatelog"`
}

func InitWithYaml(ypath string, e *echo.Echo) error {
	f, err := os.Open(ypath)
	if err != nil {
		return err
	}
	outercfg := &Outer{}
	yaml.NewDecoder(f).Decode(outercfg)
	cfg := &outercfg.Config
	if cfg.RotateLog.Filename == "" {
		return errors.New("RotateLog.Filename must be setted")
	}
	if cfg.Header != "" {
		e.Logger.SetHeader(cfg.Header)
	}
	if cfg.Level != 0 {
		e.Logger.SetLevel(log.Lvl(cfg.Level))
	}
	if cfg.Prefix != "" {
		e.Logger.SetPrefix(cfg.Prefix)
	}
	e.Logger.SetOutput(&cfg.RotateLog)
	return nil
}
