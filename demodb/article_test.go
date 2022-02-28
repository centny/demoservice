package demodb

import (
	"testing"

	"github.com/centny/demoservice/deps/go/pgx"
	"github.com/codingeasygo/util/converter"
)

func TestArticle(t *testing.T) {
	userID := int64(1000)

	//test add
	article := &Article{
		UserID:      userID,
		Ttile:       "test article",
		Description: converter.StringPtr("article descr"),
		Status:      ArticleStatusNormal,
	}
	err := AddArticle(article)
	if err != nil {
		t.Error(err)
		return
	}
	if article.TID < 1 {
		t.Errorf("the article is %v", article.TID)
		return
	}

	//test update
	article.Ttile = "updated title"
	article.Description = converter.StringPtr("updated desc")
	err = UpdateArticle(article)
	if err != nil {
		t.Error(err)
		return
	}

	//test find
	findArticle, err := FindArticleByID(article.TID)
	if err != nil {
		t.Error(err)
		return
	}
	if findArticle.Ttile != article.Ttile {
		t.Errorf("title is not updated")
		return
	}

	//test search
	articleList, total, err := SearchArticle(article.UserID, "title", 0, 10)
	if err != nil {
		t.Error(err)
		return
	}
	if len(articleList) != 1 || total != 1 {
		t.Errorf("data error")
		return
	}

	//
	//test error
	pgx.MockerStart()
	defer pgx.MockerStop()

	pgx.MockerSet("Pool.Query", 1)
	_, _, err = SearchArticle(article.UserID, "title", 0, 10)
	if err == nil {
		t.Error(err)
		return
	}
	pgx.MockerClear()

	pgx.MockerSet("Rows.Scan", 1)
	_, _, err = SearchArticle(article.UserID, "title", 0, 10)
	if err == nil {
		t.Error(err)
		return
	}
	pgx.MockerClear()

	//
	//test delete
	article.Status = ArticleStatusDeleted
	err = UpdateArticle(article)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = FindArticleByID(article.TID)
	if err == nil {
		t.Error(err)
		return
	}
}
