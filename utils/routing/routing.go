package routing

import (
	"micro_services/identity/constant"
	"micro_services/identity/utils/objective"
	"strings"

	"github.com/labstack/echo"
)

var routesMap map[string]string

func URL(name string, urlParams ...interface{}) string {
	if routesMap == nil { // run this op only once upon the first call.
		routesMap = make(map[string]string)
		for _, route := range constant.EC.Routes() {
			routesMap[route.Name] = route.Path
		}
		// debugger.Pretty("routesMap", routesMap)
	}

	path, found := routesMap[name]
	if !found {
		panic(name + " route not found")
	}

	if len(urlParams) > 0 && strings.Count(path, ":") == len(urlParams) {
		for _, param := range urlParams {
			colonPlace := strings.Index(path, ":")
			slashPlace := strings.Index(path[colonPlace:], "/") + colonPlace
			path = objective.ReplaceAtIndexTo(path, param, colonPlace, slashPlace)
		}
	}
	return path
}

func FindRoute(c echo.Context, userAccess bool) (route echo.Route) {

	for _, r := range c.Echo().Routes() {

		if r.Path == c.Path() {
			route = *r

			if r.Name == "user-access" && userAccess {
				break
			}
		}
	}

	return route
}
