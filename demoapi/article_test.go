package demoapi

import (
	"testing"

	"github.com/centny/demoservice/demodb"
	"sxbastudio.com/base/go/define"
)

func TestArticle(t *testing.T) {
	login, err := ts.GetMap("/pub/login?username=%v&password=%v", "admin", "123")
	if err != nil || login.Int64("code") != define.Success {
		t.Errorf("err:%v,code:%v", err, login)
		return
	}
	if login.Int64("user/tid") == 0 {
		t.Error("last_login_time err")
		return
	}
	article := &demodb.Article{
		Ttile: "test",
	}
	addArticle, err := ts.PostJSONMap(article, "/usr/addArticle")
	if err != nil || addArticle.Int64("code") != define.Success {
		t.Errorf("err:%v,addArticle:%v", err, addArticle)
		return
	}

	searchArticle, err := ts.GetMap("/usr/searchArticle")
	if err != nil || searchArticle.Int64("code") != define.Success {
		t.Errorf("err:%v,searchArticle:%v", err, searchArticle)
		return
	}
	articles := searchArticle.ArrayMapDef(nil, "/articles")
	if len(articles) != 1 {
		t.Errorf("err:%v,searchArticle:%v", err, searchArticle)
		return
	}
	total := searchArticle.Int64Def(0, "/total")
	if total != 1 {
		t.Errorf("err:%v,searchArticle:%v", err, searchArticle)
		return
	}
}
