package main

import (
	"encoding/gob"
	"time"

	"github.com/Centny/rediscache"
	"github.com/centny/demoservice/demoapi"
	"github.com/centny/demoservice/demodb"
	"github.com/centny/demoservice/deps/go/pgx"
	"github.com/centny/demoservice/deps/go/session"
	"github.com/codingeasygo/util/xmap"
	"github.com/codingeasygo/web"
)

func main() {
	rediscache.InitRedisPool("redis:6379?db=3")
	err := pgx.Bootstrap("postgresql://dev:123@pg.loc:5432/demoservice")
	if err != nil {
		panic(err)
	}
	demodb.Pool = pgx.C
	_, err = demodb.CheckDb()
	if err != nil {
		panic(err)
	}

	gob.Register(xmap.M{})
	sb := session.NewDbSessionBuilder()
	sb.Redis = rediscache.C
	sb.ShowLog = false
	web.Shared.Builder = sb
	web.Shared.ShowSlow = 100 * time.Millisecond
	// web.Shared.Builder = web.NewMemSessionBuilder("", "", "session_id", 60*time.Second)
	demoapi.Handle("", web.Shared)
	web.ListenAndServe(":8080")
}
