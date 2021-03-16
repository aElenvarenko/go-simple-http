package go_simple_http

import (
	"net/http"
)

// Config ...
type Config struct {
	Host string
	Port string
}

// App ...
type App struct {
	Config *Config
	Vars   map[string]interface{}
	Server *http.Server
	Router *Router
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}

// NewApp ...
func NewApp(config *Config) *App {
	router := NewRouter()

	return &App{
		Config: config,
		Server: &http.Server{
			Addr:    config.Host + ":" + config.Port,
			Handler: router,
		},
		Router: router,
	}
}

// Mount ...
func (app *App) Mount(path string, router *Router) {
	for routePath, routeData := range router.Routes {
		app.Router.Routes["^"+path+routePath[1:]] = routeData
	}
}

// Static ...
func (app *App) Static(path string, dir string) {
}

// Use ...
func (app *App) Use(path string, f func(Request, Response)) {
	app.Router.Use(path, f)
}

// Get ...
func (app *App) Get(path string, f func(Request, Response)) {
	app.Router.Get(path, f)
}

// Post ...
func (app *App) Post(path string, f func(Request, Response)) {
	app.Router.Post(path, f)
}

// Put ...
func (app *App) Put(path string, f func(Request, Response)) {
	app.Router.Put(path, f)
}

// Patch ...
func (app *App) Patch(path string, f func(Request, Response)) {
	app.Router.Patch(path, f)
}

// Delete ...
func (app *App) Delete(path string, f func(Request, Response)) {
	app.Router.Delete(path, f)
}

// func (app *App) Path(path string, f func(Request, Response)) {
// 	app.Router.Path(path, f)
// }

// Listen ...
func (app *App) Listen() error {
	return app.Server.ListenAndServe()
}
