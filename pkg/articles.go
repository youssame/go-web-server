package pkg

import (
	"net/http"
)

// All renders the list of articles
func articlesList() Handler {
	handler := func(writer http.ResponseWriter, request *http.Request) {
		ComponentRenderer(&writer, "articles", ArticlesMock())
	}
	path := ApiUrl + "/articles"
	return NewHandler(path, handler)
}

type Article struct {
	Title   string
	Content string
	Link    string
	Author  string
	Date    string
}

type ArticleData struct {
	Articles []Article
}

// ArticlesMock to be deleted after
func ArticlesMock() ArticleData {
	articles := []Article{
		{
			Title:   "Understanding Go's Compilation Process",
			Content: "An in-depth look into how Go compiles code and the trade-offs involved.",
			Link:    "https://example.com/go-compilation",
			Author:  "John Doe",
			Date:    "2024-10-10",
		},
		{
			Title:   "Stream Processing with Apache Flink in Java",
			Content: "A gentle introduction to stream processing and Flink's role in real-time data analysis.",
			Link:    "https://example.com/flink-stream-processing",
			Author:  "Jane Smith",
			Date:    "2024-11-01",
		},
		{
			Title:   "Best Practices for Onboarding in Software Teams",
			Content: "Guidelines and steps to effectively onboard new engineers in development teams.",
			Link:    "https://example.com/onboarding-practices",
			Author:  "Ahmed El-Mansouri",
			Date:    "2024-09-15",
		},
		{
			Title:   "An Introduction to JIT and AOT Compilation",
			Content: "A comparison of Just-In-Time and Ahead-Of-Time compilation and their use cases.",
			Link:    "https://example.com/jit-vs-aot",
			Author:  "Sara Lee",
			Date:    "2024-10-05",
		},
		{
			Title:   "Building a Simple Web Server with Go",
			Content: "A step-by-step guide to creating a basic web server in Golang.",
			Link:    "https://example.com/go-web-server",
			Author:  "Tom Nguyen",
			Date:    "2024-10-25",
		},
		{
			Title:   "How to Use Docker for CI/CD in JavaScript Projects",
			Content: "A tutorial on setting up CI/CD for JavaScript apps using Docker.",
			Link:    "https://example.com/docker-ci-cd-js",
			Author:  "Emily Tran",
			Date:    "2024-11-02",
		},
		{
			Title:   "Solr Basics: Full-Text Search and Indexing",
			Content: "An introductory article on using Solr for text search and indexing in applications.",
			Link:    "https://example.com/solr-basics",
			Author:  "Ryan H.",
			Date:    "2024-10-21",
		},
	}
	return ArticleData{
		articles,
	}
}
