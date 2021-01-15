package preset

import (
	"github.com/labstack/echo"
	"testing"
)

func TestGetDefaultYamlPath(t *testing.T) {
	t.Logf("%s", getDefaultYamlPath())
	// t.Fatal("not implemented")
}

func TestInit(t *testing.T) {
	SetLogger(t)
	Init(echo.New())
	// t.Fatal("not implemented")
}
