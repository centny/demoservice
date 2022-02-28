package demodb

import "github.com/codingeasygo/util/xsql"

const (
	UserStatusNormal   = 100
	UserStatusDisabled = 200
	UserStatusDeleted  = -1
)

type User struct {
	TID        int64     `json:"tid"`
	Username   string    `json:"username"`
	Password   *string   `json:"password"`    //the user password encrypted by sha1
	UpdateTime xsql.Time `json:"update_time"` //the last updatet time
	CreateTime xsql.Time `json:"create_time"` //the user create time
	Status     int       `json:"status"`      //the user status, see define by UserStatus*
}

const (
	ArticleStatusNormal  = 100
	ArticleStatusDeleted = -1
)

type Article struct {
	TID         int64     `json:"tid"`
	UserID      int64     `json:"user_id"`
	Ttile       string    `json:"title"`
	Description *string   `json:"description"`
	Image       *string   `json:"image"`
	UpdateTime  xsql.Time `json:"update_time"` //the last updatet time
	CreateTime  xsql.Time `json:"create_time"` //the user create time
	Status      int       `json:"status"`      //the user status, see define by UserStatus*
}
