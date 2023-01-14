package context

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

type GlobalContext struct {
	echo.Context
	ResponseContext echo.Map
	Session         session.Session
	AccessToken     jwt.Claims
	RefreshToken    jwt.Claims
}

func InitContext(c echo.Context) *GlobalContext {
	g := &GlobalContext{Context: c}
	g.Session = session.Default(g)
	g.ResponseContext = echo.Map{}
	g.AccessToken, g.RefreshToken = jwt.MapClaims{}, jwt.MapClaims{}
	return g
}

func ContextHandler(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		g := &GlobalContext{Context: c}
		g.Session = session.Default(g)

		return h(g)
	}
}

func (g *GlobalContext) GetCookie(name string) string {
	for _, c := range g.Cookies() {
		if c.Name == name {
			return c.Value
		}
	}
	return ""
}

func (g *GlobalContext) DeleteCookie(key string) {
	cookie := new(http.Cookie)
	cookie.Name = key
	cookie.Path = "/"
	cookie.Value = ""
	cookie.Expires = time.Unix(0, 0)
	// cookie.Domain = g.Domain
	cookie.HttpOnly = true
	g.Context.SetCookie(cookie)
}

func (g *GlobalContext) WriteCookie(key string, value string) {
	g.DeleteCookie(key)
	cookie := new(http.Cookie)
	cookie.Name = key
	cookie.Path = "/"
	cookie.Value = value
	cookie.Expires = time.Now().Add(24 * time.Hour)
	// cookie.Domain = g.Domain
	cookie.HttpOnly = true
	g.Context.SetCookie(cookie)
}

func (g *GlobalContext) ReadCookie(key string) string {
	cookie, err := g.Context.Cookie(key)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func (g *GlobalContext) PageParam() int {
	page, _ := strconv.Atoi(g.QueryParam("page"))
	if page < 1 {
		return 1
	}
	return page
}

func (g *GlobalContext) LimitParam() int {
	limit, _ := strconv.Atoi(g.QueryParam("limit"))
	if limit < 1 {
		return 10
	}

	return limit
}

func (g *GlobalContext) SearchParam() string {
	return g.QueryParam("search")
}

func (g *GlobalContext) SortParam() string {
	return g.QueryParam("sort")
}

func (g *GlobalContext) NextURL() string {
	val := g.Session.Get("NextUrlCookie")
	url, ok := val.(string)
	if !ok || url == "" {
		return ""
	}
	g.Session.Delete("NextUrlCookie")
	g.Session.Save()
	return url
}

func (g *GlobalContext) RenderTemplate(httpStatus int, path string, c map[string]interface{}) error {

	context := echo.Map{}
	for key, value := range c {
		context[key] = value
	}

	err := g.Render(httpStatus, path, context)
	if err != nil {
		fmt.Println("template error:  ", err.Error())
	}
	return err
}
