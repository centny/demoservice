package demoapi

import (
	"github.com/centny/demoservice/demodb"
	"github.com/centny/demoservice/demoupgrade"
	"github.com/centny/demoservice/deps/go/pgx"
	"github.com/codingeasygo/util/xhttp"
	"github.com/codingeasygo/web/httptest"
)

var ts *httptest.Server

func init() {
	err := pgx.Bootstrap("postgresql://dev:123@pg.loc:5432/demoservice")
	if err != nil {
		panic(err)
	}
	demodb.Pool = pgx.C
	_, err = demodb.Pool().Exec(demoupgrade.DROP)
	if err != nil {
		panic(err)
	}
	_, err = demodb.CheckDb()
	if err != nil {
		panic(err)
	}

	ts = httptest.NewMuxServer()
	Handle("", ts.Mux)
	xhttp.EnableCookie()
}
