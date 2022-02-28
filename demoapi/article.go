package demoapi

import (
	"github.com/centny/demoservice/demodb"
	"github.com/codingeasygo/util/attrvalid"
	"github.com/codingeasygo/util/xmap"
	"github.com/codingeasygo/web"
	"sxbastudio.com/base/go/define"
)

func AddArticleH(w *web.Session) web.Result {
	article := &demodb.Article{}
	_, err := w.RecvJSON(article)
	if err != nil {
		return w.SendJSON(xmap.M{
			"code":    define.ArgsInvalid,
			"message": err.Error(),
		})
	}
	err = attrvalid.Valid(`
		title,R|S,L:0~255;
		description,O|S,L:0~1024;
	`, article)
	if err != nil {
		return w.SendJSON(xmap.M{
			"code":    define.ArgsInvalid,
			"message": err.Error(),
		})
	}
	userID := w.Int64Def(0, "user_id")
	article.UserID = userID
	article.Status = demodb.ArticleStatusNormal
	err = demodb.AddArticle(article)
	if err != nil {
		return w.SendJSON(xmap.M{
			"code":    define.ServerError,
			"message": err.Error(),
		})
	}
	return w.SendJSON(xmap.M{
		"code":    define.Success,
		"article": article,
	})
}

func SearchArticleH(w *web.Session) web.Result {
	var key string
	var skip, limit int
	err := w.ValidFormat(`
		key,O|S,L:0~255;
		skip,O|I,R:-1;
		limit,O|I,R:0;
	`, &key, &skip, &limit)
	if err != nil {
		return w.SendJSON(xmap.M{
			"code":    define.ArgsInvalid,
			"message": err.Error(),
		})
	}
	userID := w.Int64Def(0, "user_id")
	articleList, total, err := demodb.SearchArticle(userID, key, skip, limit)
	if err != nil {
		return w.SendJSON(xmap.M{
			"code":    define.ServerError,
			"message": err.Error(),
		})
	}
	return w.SendJSON(xmap.M{
		"code":     define.Success,
		"articles": articleList,
		"total":    total,
	})
}
