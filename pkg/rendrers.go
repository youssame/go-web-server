package pkg

import (
	"fmt"
	"go-webserver/templates"
	"html/template"
	"net/http"
)

// Handler a type for handlers
type Handler struct {
	pattern string
	handler func(http.ResponseWriter, *http.Request)
}

// ComponentRenderer a renderer for components
func ComponentRenderer(w *http.ResponseWriter, component string, data any) {
	pageContent := templates.GetComponent(component)
	t, err := template.New(component).Parse(pageContent)
	if err != nil {
		panic(err)
	}
	err = t.Execute(*w, data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// StaticRenderer Static page renderer
func StaticRenderer(w *http.ResponseWriter, page string, data any) {
	pageContent := templates.GetPage(page)
	t, err := template.New(page).Parse(pageContent)
	if err != nil {
		panic(err)
	}
	err = t.Execute(*w, data)
	if err != nil {
		panic(err)
	}
}

func NewHandler(path string, handler func(http.ResponseWriter, *http.Request)) Handler {
	return Handler{
		pattern: path,
		handler: handler,
	}
}

func BuildHandler(handler Handler) (string, func(http.ResponseWriter, *http.Request)) {
	return handler.pattern, handler.handler
}
