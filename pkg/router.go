package pkg

import (
	"net/http"
)

// InitRoutes init the handlers
func InitRoutes() {
	for _, route := range routes() {
		path, handler := BuildHandler(route)
		AddRoute(path, handler)
	}
}

/***** Handlers *****/

// routes: a list of routes
func routes() []Handler {
	routes := []Handler{
		home(),
		about(),
		articles(),
		articlesList(),
	}
	return routes
}

// HomeHandler renders the home page
func home() Handler {
	handler := func(writer http.ResponseWriter, request *http.Request) {
		StaticRenderer(&writer, "home", nil)
	}
	return NewHandler("/", handler)
}

// AboutHandler renders the about page
func about() Handler {
	handler := func(writer http.ResponseWriter, request *http.Request) {
		StaticRenderer(&writer, "about", nil)
	}
	return NewHandler("/about", handler)
}

// ArticlesHandler renders the articles page
func articles() Handler {
	handler := func(writer http.ResponseWriter, request *http.Request) {
		StaticRenderer(&writer, "articles", ArticlesMock())
	}
	return NewHandler("/articles", handler)
}
