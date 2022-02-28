package demodb

func FindUserByUsrPass(username, password string) (user *User, err error) {
	querySQL := `
		select tid,username,create_time,update_time,status from demo_user where username=$1 and password=$2 and status=$3
	`
	queryArgs := []interface{}{username, password, UserStatusNormal}
	user = &User{}
	err = Pool().QueryRow(querySQL, queryArgs...).Scan(&user.TID, &user.Username, &user.CreateTime, &user.UpdateTime, &user.Status)
	return
}
