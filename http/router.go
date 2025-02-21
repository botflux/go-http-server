package http

type Route struct {
	Method  string
	Path    string
	Handler func(request Request) Response
}

type Router struct {
	Routes []Route
}

func (r *Router) Add(method string, path string, handler func(request Request) Response) {
	r.Routes = append(r.Routes, Route{Method: method, Path: path, Handler: handler})
}

func (r *Router) Handle(request Request) Response {
	for _, route := range r.Routes {
		if route.Method == request.Method && route.Path == request.URL {
			return route.Handler(request)
		}
	}

	return Response{
		StatusCode: 404,
		Headers: map[string]string{
			"Content-Type": "text/html; charset=utf-8",
		},
		Body: "<html>" +
			"<body>" +
			"<h1>404 Not Found</h1>" +
			"</body>" +
			"</html>",
	}
}
