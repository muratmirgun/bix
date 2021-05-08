package bix

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	log "github.com/muratmirgun/logger-go"
)

// Stores routes added by the application.
type Router struct {
	routes []route
}

// Prototype for the handler function.
type Handler func(http.ResponseWriter, *http.Request)

// Describes a single route as a combination of HTTP VERB, regular expression
// path matcher, and the handler function.
type route struct {
	verb    string
	path    *regexp.Regexp
	handler Handler
}

// Some globals to make life easier.
var (
	paramRE = regexp.MustCompile("{(.+?)}")
)

// NewRouter initializes a new HTTP request httprouter.
func NewRouter() *Router {
	return new(Router)
}

// Adds a new route with a handler function. The router structure is also
// returned to allow chaining.
func (router *Router) AddRoute(verb string, path string, handler Handler) *Router {
	log.Info("Adding route", log.String("verb", verb), log.String("path", path))

	// Converts params in the path from "{param}" to a non-greedy regex named
	// match, "(?P<param>.+?)"
	if path != "/" {
		path = strings.TrimRight(path, "/")
		submatches := paramRE.FindAllString(path, -1)
		for _, s := range submatches {
			path = strings.Replace(path, s, "(?P<"+strings.Trim(s, "{}")+">.+?)", 1)
		}
		path = "^" + path + "$"
	}

	// Compile the path regex
	re, err := regexp.Compile(path)
	if err != nil {
		log.Error("Invalid path regex", log.Err(err))
	}

	// Adds the route if no errors occurred the regex compiler.
	var r route
	r.handler = handler
	r.path = re
	r.verb = verb

	router.routes = append(router.routes, r)
	return router
}

// Default global request handler that matches the incoming request with a
// registered handler.
func (router *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for _, r := range router.routes {
		if request.Method == r.verb && r.path.MatchString(request.URL.Path) {
			m := matches(r.path, request.URL.Path)
			for k, v := range request.URL.Query() {
				m[k] = strings.Join(v, "; ")
			}

			r.handler(writer, request.WithContext(context.WithValue(request.Context(), "params", m)))
			return
		}
	}
	log.Warn("Path not found", log.String("path", request.URL.Path))
	writer.WriteHeader(404)
}

// Retrieves query string parameters from the request context.
func GetParams(ctx context.Context) map[string]string {
	return ctx.Value("params").(map[string]string)
}

// Helper that applies the path regex to the incoming path to parse param
// values from it.
func matches(re *regexp.Regexp, s string) map[string]string {
	submatches := re.FindStringSubmatch(s)
	matches := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i > 0 && name != "" {
			matches[name] = submatches[i]
		}
	}
	return matches
}
