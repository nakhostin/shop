package static

import (
	"micro_services/identity/context"
	"micro_services/identity/utils/routing"
	"net/http"

	"github.com/labstack/echo"
)

func NotFound(c echo.Context) error {
	g := c.(*context.GlobalContext)
	return g.Redirect(http.StatusFound, routing.URL("NotFound"))
}
