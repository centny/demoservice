package demoapi

import "github.com/codingeasygo/web"

func Handle(pre string, mux *web.SessionMux) {
	mux.FilterFunc("^/usr/.*$", LoginAccessF)

	mux.HandleFunc("/pub/login", LoginH)
	mux.HandleFunc("/usr/addArticle", AddArticleH)
	mux.HandleFunc("/usr/searchArticle", SearchArticleH)
}
