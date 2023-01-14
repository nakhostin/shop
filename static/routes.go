package static

import (
	"micro_services/identity/server/config"

	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	e.File("/", config.CFG.StaticFilePath+"wellcome.html")
	e.File("/notfound-404", config.CFG.StaticFilePath+"notfound.html").Name = "NotFound"
}
