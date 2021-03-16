package go_simple_http

import (
	"net/http"
	"regexp"
	"strings"
)

// Router ...
type Router struct {
	Routes map[string]map[string]Route
}

// NewRouter ...
func NewRouter() *Router {
	return &Router{
		Routes: make(map[string]map[string]Route),
	}
}

// makeExp ...
func (router *Router) makeExp(path string) string {
	expStr := `(:[a-z]+)`
	expRepStr := `([0-9a-z]+)`
	paramExp := regexp.MustCompile(expStr)

	if paramExp.MatchString(path) {
		path = paramExp.ReplaceAllLiteralString(path, expRepStr)
	}

	return "^" + path + "$"
}

// makeRoute ...
func (router *Router) makeRoute(path string, method string, f func(Request, Response)) {
	exp := router.makeExp(path)

	if _, ok := router.Routes[path]; !ok {
		router.Routes[exp] = make(map[string]Route)
	}

	router.Routes[router.makeExp(path)][method] = Route{
		Method: method,
		Path:   path,
		Exp:    exp,
		F:      f,
	}
}

// parseParams ...
// func (router *Router) parseParams(url string) {}

// ServeHTTP ...
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for expStr, methods := range router.Routes {
		exp := regexp.MustCompile(expStr)
		if exp.MatchString(strings.TrimRight(r.URL.Path, "/")) {
			if route, ok := methods[r.Method]; ok {
				route.F(NewRequest(r), NewResponse(w))
			}
			break
		}
	}
}

// Static ...
func (router *Router) Static(path string, dir string) {

}

// Use ...
func (router *Router) Use(path string, f func(Request, Response)) {
	router.makeRoute(path, "ALL", f)
}

// Get ...
func (router *Router) Get(path string, f func(Request, Response)) {
	router.makeRoute(path, http.MethodGet, f)
}

// Post ...
func (router *Router) Post(path string, f func(Request, Response)) {
	router.makeRoute(path, http.MethodPost, f)
}

// Put ...
func (router *Router) Put(path string, f func(Request, Response)) {
	router.makeRoute(path, http.MethodPut, f)
}

// Patch ...
func (router *Router) Patch(path string, f func(Request, Response)) {
	router.makeRoute(path, http.MethodPatch, f)
}

// Delete ...
func (router *Router) Delete(path string, f func(Request, Response)) {
	router.makeRoute(path, http.MethodDelete, f)
}
