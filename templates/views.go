package templates

type views struct {
	pages      map[string]string
	components map[string]string
}

// pages app navigation pages
var pages = map[string]string{
	"home":     "",
	"about":    "",
	"articles": "",
}

// components app navigation pages
var components = map[string]string{
	"articles": "",
}
var vs = views{
	pages:      pages,
	components: components,
}
var viewsLoaded = false

// GetViews get app views
func GetViews() views {
	if !viewsLoaded {
		initViews()
	}
	return vs
}

// GetPage get app page
func GetPage(k string) string {
	return GetViews().pages[k]
}

// GetComponent get app page
func GetComponent(k string) string {
	return GetViews().components[k]
}
func initViews() {
	vs.pages["home"] = masterPages(home)
	vs.pages["about"] = masterPages(about)
	vs.pages["articles"] = masterPages(articlesPage)
	vs.components["articles"] = articles
	viewsLoaded = false
}

func masterPages(page string) string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Blog</title>
    <script src="https://unpkg.com/htmx.org@2.0.3"></script>
</head>
<body>
    <div class="container">
        ` + page + `
    </div>
</body>
</html>`
}

var home = `<h1>Home</h1>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Aperiam consequuntur cupiditate doloribus esse et expedita ipsum itaque, minima nobis non officia optio pariatur quisquam quod sequi sit suscipit tempora unde?</p>`

var about = `<h1>About</h1>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Aperiam consequuntur cupiditate doloribus esse et expedita ipsum itaque, minima nobis non officia optio pariatur quisquam quod sequi sit suscipit tempora unde?</p>`

var article = `<article>
    <h2>{{.Title}}</h2>
    <p>{{.Content}}</p>
    <a href="{{.Link}}">Read</a>
    <footer>
      <p>
        Posted on
        <time datetime="{{.Date}}">{{.Date}}</time>
        by {{.Author}}.
      </p>
    </footer>
  </article>`
var articlesPage = `<h1>Articles</h1>
<aside>
` + articles + `
</aside>`
var articles = ` {{range .Articles}}
` + article + `
  {{end}}`
