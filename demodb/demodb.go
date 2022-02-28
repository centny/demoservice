package demodb

import (
	"github.com/centny/demoservice/demoupgrade"
	"github.com/centny/demoservice/deps/go/pgx"
	"github.com/centny/demoservice/deps/go/xlog"
)

//Pool will return database connection pool
var Pool = func() *pgx.Pool {
	panic("db is not initial")
}

//CheckDb will check database if is initial
func CheckDb() (created bool, err error) {
	_, err = Pool().Exec(`select tid from demo_user limit 1`)
	if err != nil {
		xlog.Infof("start generate database...")
		_, err = Pool().Exec(demoupgrade.LATEST)
		created = true
	}
	return
}
