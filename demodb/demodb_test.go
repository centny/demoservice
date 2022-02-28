package demodb

import (
	"github.com/centny/demoservice/demoupgrade"
	"github.com/centny/demoservice/deps/go/pgx"
)

func init() {
	err := pgx.Bootstrap("postgresql://dev:123@pg.loc:5432/demoservice")
	if err != nil {
		panic(err)
	}
	Pool = pgx.C
	_, err = Pool().Exec(demoupgrade.DROP)
	if err != nil {
		panic(err)
	}
	_, err = CheckDb()
	if err != nil {
		panic(err)
	}
}
