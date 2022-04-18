package main

import (
	"encoding/gob"
	"os"
	"time"

	"github.com/Centny/rediscache"
	"github.com/centny/demoservice/demoapi"
	"github.com/centny/demoservice/demodb"
	"github.com/centny/demoservice/deps/go/pgx"
	"github.com/centny/demoservice/deps/go/session"
	"github.com/centny/demoservice/deps/go/xlog"
	"github.com/codingeasygo/util/xmap"
	"github.com/codingeasygo/util/xprop"
	"github.com/codingeasygo/web"
)

func main() {
	confPath := "conf/demoservice.properties"
	if len(os.Args) > 1 {
		confPath = os.Args[1]
	}
	conf := xprop.NewConfig()
	conf.LoadFile(confPath)
	conf.Print()
	rediscache.InitRedisPool(conf.StrDef("", "/server/redis_conn"))
	err := pgx.Bootstrap(conf.StrDef("", "/server/pg_conn"))
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
	xlog.Infof("demo service listen on %v", conf.StrDef("", "/server/listen"))
	web.ListenAndServe(conf.StrDef("", "/server/listen"))
}
