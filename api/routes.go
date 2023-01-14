package api

import (
	"micro_services/identity/api/public"
	"micro_services/identity/server/config"

	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	gp := e.Group(config.CFG.App)

	public.Routes(gp)
}
