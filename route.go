package go_simple_http

// Route ...
type Route struct {
	Method string
	Path   string
	Params string
	Exp    string
	F      func(Request, Response)
}

// Exec ...
func (route *Route) Exec() {

}
