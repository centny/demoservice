package demoapi

import (
	"github.com/centny/demoservice/demodb"
	"github.com/codingeasygo/util/xmap"
	"github.com/codingeasygo/web"
	"github.com/jackc/pgx/v4"
	"sxbastudio.com/base/go/define"
)

func LoginAccessF(w *web.Session) web.Result {
	userID := w.Int64Def(0, "user_id")
	if userID < 1 {
		return w.SendJSON(xmap.M{
			"code":    define.NotAccess,
			"message": "not login",
		})
	}
	return web.Continue
}

func LoginH(w *web.Session) web.Result {
	var username, password string
	err := w.ValidFormat(`
		username,R|S,L:0~255;
		password,R|S,L:0~255;
	`, &username, &password)
	if err != nil {
		return w.SendJSON(xmap.M{
			"code":    define.ArgsInvalid,
			"message": err.Error(),
		})
	}
	user, err := demodb.FindUserByUsrPass(username, password)
	if err != nil {
		code := define.ServerError
		if err == pgx.ErrNoRows {
			code = define.NotFound
		}
		return w.SendJSON(xmap.M{
			"code":    code,
			"message": err.Error(),
		})
	}
	w.SetValue("user_id", user.TID)
	w.Flush()
	return w.SendJSON(xmap.M{
		"code":       define.Success,
		"user":       user,
		"session_id": w.ID(),
	})
}
