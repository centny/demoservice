package demodb

import (
	"fmt"
	"strings"

	"github.com/codingeasygo/util/xsql"
)

func AddArticle(article *Article) (err error) {
	article.CreateTime = xsql.TimeNow()
	article.UpdateTime = xsql.TimeNow()
	insertSQL := `
		insert into demo_article(user_id,title,description,create_time,update_time,status) values ($1,$2,$3,$4,$5,$6)
		returning tid
	`
	insertArgs := []interface{}{article.UserID, article.Ttile, article.Description, article.CreateTime, article.UpdateTime, article.Status}
	err = Pool().QueryRow(insertSQL, insertArgs...).Scan(&article.TID)
	return
}

func UpdateArticle(article *Article) (err error) {
	article.UpdateTime = xsql.TimeNow()
	sets := []string{}
	args := []interface{}{}
	if len(article.Ttile) > 0 {
		args = append(args, article.Ttile)
		sets = append(sets, fmt.Sprintf("title=$%v", len(args)))
	}
	if article.Description != nil && len(*article.Description) > 0 {
		args = append(args, article.Description)
		sets = append(sets, fmt.Sprintf("description=$%v", len(args)))
	}
	if article.Status != 0 {
		args = append(args, article.Status)
		sets = append(sets, fmt.Sprintf("status=$%v", len(args)))
	}
	args = append(args, article.UpdateTime)
	sets = append(sets, fmt.Sprintf("update_time=$%v", len(args)))
	//
	args = append(args, article.TID, article.UserID)
	updateSQL := fmt.Sprintf(`update demo_article set %v where tid=$%v and user_id=$%v`, strings.Join(sets, ","), len(args)-1, len(args))
	err = Pool().ExecRow(updateSQL, args...)
	return
}

func FindArticleByID(articleID int64) (article *Article, err error) {
	querySQL := `
		select tid,user_id,title,description,create_time,update_time,status from demo_article where tid=$1 and status=$2
	`
	queryArgs := []interface{}{articleID, ArticleStatusNormal}
	article = &Article{}
	err = Pool().QueryRow(querySQL, queryArgs...).Scan(
		&article.TID, &article.UserID, &article.Ttile, &article.Description,
		&article.CreateTime, &article.UpdateTime, &article.Status,
	)
	return
}

func SearchArticle(userID int64, key string, skip, limit int) (articleList []*Article, total int64, err error) {
	where := []string{"user_id=$1", "status=$2"}
	args := []interface{}{userID, ArticleStatusNormal}
	if key != "" {
		args = append(args, "%"+key+"%")
		where = append(where, fmt.Sprintf("title like $%d ", len(args)))
	}
	querySQL := `
		select tid,user_id,title,description,create_time,update_time,status from demo_article
	`
	countSQL := `select count(*) from demo_article`
	if len(where) > 0 {
		querySQL += " where " + strings.Join(where, " and ")
		countSQL += " where " + strings.Join(where, " and ")
	}
	querySQL += " order by update_time desc "
	if limit > 0 {
		querySQL += fmt.Sprintf(" limit %v ", limit)
	}
	if skip >= 0 {
		querySQL += fmt.Sprintf(" offset %v ", skip)
	}
	rows, err := Pool().Query(querySQL, args...)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		article := &Article{}
		err = rows.Scan(
			&article.TID, &article.UserID, &article.Ttile, &article.Description,
			&article.CreateTime, &article.UpdateTime, &article.Status,
		)
		if err != nil {
			break
		}
		articleList = append(articleList, article)
	}
	if err != nil {
		return
	}
	err = Pool().QueryRow(countSQL, args...).Scan(&total)
	return
}
